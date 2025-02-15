const { Engine } = require('bpmn-engine');
const EventEmitter = require('events');
const fs = require('fs');
const { get } = require('http');
const { exec } = require('child_process');
const { exit } = require('process');

const path = '../models/leasing-payment.bpmn';
// Read BPMN definition
const definition = fs.readFileSync(path, 'utf8');

let participants = readParticipantNames();

let elementMap = {};
let activities = [];
let messages = [];
let state = [];

// Create an EventEmitter as the listener
const listener = new EventEmitter();

// Store transitions as pairs
const transitions = [];
let lastActivity = null;
let lastActivity2 = null;
let lastActivityIndex = -1;

const COMPLETING = 3;
const COMPLETED = 10;
const READY = 1;
const INACTIVE = 0;

listener.on('activity.end', (api) => {
  console.log(`Activity ended: ${api.id}`);
  if (isExecutable(api)) {
    let activity = getActivityById(api.id);

    let index = activities.indexOf(activity);

    let currentState = state.slice();
    let newState = state.slice();
    if (lastActivityIndex != -1) {
      currentState[lastActivityIndex] = COMPLETING;
      newState[lastActivityIndex] = COMPLETED;
    }
    if (lastActivity2 != null) {
      let lastActivityObj = elementMap[lastActivity2];
      lastActivityObj.outbound.forEach((edge) => {
        element = elementMap[edge.targetId];
        if (isExecutable(element)) {
          let targetIndex = activities.indexOf(element);
          if (state[targetIndex] == INACTIVE) {
            newState[targetIndex] = READY;
          }
        }
      });
    }
    if (!wasStartLast()) {
      executeZkWF(currentState, newState,lastActivityIndex);
    }
    if (index != -1) {
      //newState[index] = READY;
      lastActivityIndex = index;
      state = newState;
    }
    
  } else if (api.type == 'bpmn:EndEvent') {
    currentState = state.slice();
    currentState[lastActivityIndex] = COMPLETING;
    newState = state.slice();
    newState[lastActivityIndex] = COMPLETED;
    //console.log('State transition:', state, '->', newState);
    executeZkWF(currentState, newState,lastActivityIndex);
  }



  lastActivity = api.id; // Track the last activity that ended
  lastActivity2 = api.id;
});

listener.on('activity.start', (api) => {
  console.log(`Activity started: ${api.id}`);
  if (lastActivity) {
    transitions.push([lastActivity, api.id]); // Create a transition pair
    lastActivity = null; // Reset after recording
  }
});

const engine = new Engine({
  name: 'Transition Tracker',
  source: definition,
});

engine.getDefinitions().then((definitions) => {
  console.log('Loaded', definitions[0].id);
  console.log('The definition comes with process', definitions[0].getProcesses()[0].id);
  

  let processes = definitions[0].getProcesses();
  processes.forEach((process) => {

    console.log('Process', process.id);
    
    activities = activities.concat(process.getActivities());
  });

  activities.forEach((activity) => {
    console.log(activity.id, activity.type);
    if (activity.type == 'bpmn:IntermediateThrowEvent') {
      // Get the message id
      let message = activity.eventDefinitions[0];
      console.log('Message:', message.reference.id);
      messages.push(message.reference.id);
    }

    elementMap[activity.id] = activity;
  });


  startEvents = activities.filter((activity) => activity.type == 'bpmn:StartEvent');
  layerMap = layerize(startEvents, elementMap);

  activities = activities.filter((activity) => isExecutable(activity));
  // Sort activities by layer, then by id
  activities = activities.sort((a, b) => {
    if (layerMap[a.id] == layerMap[b.id]) {
      return a.id > b.id ? 1 : -1;
    }
    return layerMap[a.id] - layerMap[b.id];
  });
  activities.forEach((activity) => {
    console.log(layerMap[activity.id]+":"+activity.id);
  });


  for (let i = 0; i < activities.length; i++) {
    state.push(0);
  }
});
async function main() {
  await compileZkWF();
  await setupZkWF();

  engine.execute({
    listener,
    variables: {}, // Add any process variables if needed
  }, (err) => {
    if (err) throw err;

    console.log('Execution completed.');
    console.log('Generated transitions:');
    console.log(transitions);
  });
}
main();

function getActivityById(id) {
  return activities.find((activity) => activity.id == id);
}

function isExecutable(activity) {
  return activity.type == 'bpmn:Task' || activity.type == 'bpmn:IntermediateCatchEvent' || activity.type == 'bpmn:IntermediateThrowEvent';
}

function wasStartLast() {
  if (lastActivity2 == null) {
    return false;
  }

  let lastActivityObj = elementMap[lastActivity2];

  return lastActivityObj.type == 'bpmn:StartEvent';
}

function layerize(startEvents, elementMap) {
  let result = {};
  let queue = [];
  let nextLayer = [];
  let layer = 0;
  let visited = {};

  startEvents.forEach((startEvent) => {
    queue.push(startEvent.id);
    visited[startEvent.id] = true;
  });

  while (queue.length > 0) {
    let nodeID = queue.shift();
    let node = elementMap[nodeID];


    node.outbound.forEach((edge) => {
      let targetID = edge.targetId;
      if (!visited[targetID]) {
        result[targetID] = layer;
        nextLayer.push(targetID);
        visited[targetID] = true;
      }
    });

    if (queue.length == 0) {
      queue = nextLayer;
      nextLayer = [];
      layer++;
    }
  }

  return result;
}

function executeWithPromise(command) {
  return new Promise((resolve, reject) => {
    exec(command, (err, stdout, stderr) => {
      if (err) {
        console.log(stdout);
        reject(err);
        return;
      }
      resolve(stdout);
    });
  });
}

async function compileZkWF() {
  await executeWithPromise('../bin/zkwf compile ' + path).then((stdout) => {
    console.log(stdout);
  });
}

async function fillInputs(stateId) {
  await executeWithPromise(`../bin/zkwf fill-inputs state-${stateId}.json ../key.json`).then((stdout) => {
    console.log(stdout);
  });
}

async function createWitness(stateId) {
  await executeWithPromise(`../bin/zkwf witness ${path} state-${stateId}.json ../key.json --full full-${stateId}.wtns`).then((stdout) => {
    console.log(stdout);
  });
}

async function proveZkWF(stateId) {
  await executeWithPromise(`../bin/zkwf prove circuit.r1cs pk.bin full-${stateId}.wtns -o proof-${stateId}`).then((stdout) => {
    //console.log(stdout);
  });
}

async function setupZkWF() {
  // ./bin/zkfw setup circuit.r1cs
  await executeWithPromise('../bin/zkwf setup circuit.r1cs').then((stdout) => {
    console.log(stdout);
  });
}

function getMessagesMapCurrent(messages) {
  if (messages.length == 0) {
    return {};
  }
  let result = {};
  messages.forEach((message) => {
    result[message] = "0";
  });
  return result;
}

function getMessagesMapNew(messages, completedIndex) {
  if (messages.length == 0) {
    return {};
  }
  let completed = activities[completedIndex];
  

  let result = {};
  messages.forEach((message) => {
    let Value = "0";
    if (completed.type == 'bpmn:IntemediateThrowEvent' && completed.eventDefinitions[0].reference.id == message) {
       Value = "1";
    }
    result[message] = Value;
  });
  return result;
}

async function executeZkWF(currentState, newState,completedId) {
  console.log('State transition:', currentState, '->', newState);

  let encrypted = [0,0,0];
  messages.forEach((message) => {
    encrypted.push(0);
  });

  let balances = {};
  participants.forEach((participant) => {
    balances[participant] = "0";
    encrypted.push(0);
  });

  let data = {
    "State_curr": {
      "States": currentState,
      "Variables": {},
      "Messages": getMessagesMapCurrent(messages),
      "Balances": balances,
      "Radomness": "0"
    },
    "State_new": {
      "States": newState,
      "Variables": {},
      "Messages": getMessagesMapNew(messages, completedId),
      "Balances": balances,
      "Radomness": "0"
    },
    "HashCurr": "0",
    "HashNew": "0",
    "PublicKey": "0",
    "Signature": "0",
    "Key": [0],
    "Encrypted": encrypted
  }

  let id = makeid(10);

  // Save the data to state.json
  fs.writeFileSync(`state-${id}.json`, JSON.stringify(data));
  // Run the zkWf fill inputs ("../bin/zkwf fill-inputs state.json ../key.json")

  await fillInputs(id);
  try {
    await createWitness(id);
  } catch (err) {
    console.log('Proof failed on state', completedId, '!');
    // Print activity information
    let activity = activities[completedId];
    console.log('Activity:', activity.id, activity.type, activity.name);
    console.log(err);
    exit(1);
    return;
  }

  await proveZkWF(id);
}

function makeid(length) {
  let result = '';
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  const charactersLength = characters.length;
  let counter = 0;
  while (counter < length) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
    counter += 1;
  }
  return result;
}

function readParticipantNames() {
  // Find the participants of the process (read from the BPMN file)
  let participants = [];
  definition.split('\n').forEach((line) => {
    trimmed = line.trim();
    if (trimmed.startsWith('<bpmn2:participant') || trimmed.startsWith('bpmn2:participant')) {
      let name = trimmed.split('name="')[1].split('"')[0];
      console.log('Participant:', name);
      participants.push(name);
    }
  });
  return participants;
}