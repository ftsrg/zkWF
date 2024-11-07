# zkWF2

This is the second iteration of the zkWF (Zero Knowledge Workflow) project. The goal of this project is to create a workflow engine like system that can orchestrate business processes in a confidential and secure manner on a public blockchain. 

## Project Structure

The project is divided into the following components:
- cmds: Contains the command line interface for the project
   - zkWF: The main command line interface for the project
- pkg: Contains the core logic of the project
   - common: Contains common utility functions
   - circuits: Contains the zkSNARK circuits for the project
      - expressions: This handles the boolean expressions in the workflows
      - gmimc: This is where the GMiMC encryption circuit is defined
      - hkdf: This is where the HKDF is implemented
      - lifecycle: The lifecycle state machine circuit
      - utils: Contains utility functions for the circuits
      - statechechker: The main zkSNARK circuit for the project
  - contracts: contains the smart contract interactions for the project
  - crypto: Contains the cryptographic functions for the project that are not part of the zkSNARK circuits, e.g. to generate inputs, keys
  - model: Handles the parsing of the BPMN models
  - powersoftau: Contains the powers of tau loading functions. This is used as the trusted setup for the zkSNARK circuits
  - web3: Contains the web3 interactions for the project
  - zkp: Contains the zero knowledge proof functions for the project
- solidity: Contains the solidity contracts for the project
- models: Contains the BPMN models and their corresponding JSON test cases
- editor: Contains the BPMN editor for the project that has the capabilities to use the extended attributes and payment tasks

## Installation

### Prerequisites

- Go
- Node.js
- Solc

### Steps

1. Clone the repository

```bash
git clone -b zkwf2 https://github.com/zkWF/zkWF2.git
```

2. Compile the zkWF command line interface

```bash
go build -o bin/zkWF ./cmds/zkWF
```

3. Install the npm packages

```bash
cd editor && npm install
```

4. Start the editor

```bash
npm start
```

## Usage

### zkWF Command Line Interface

The zkWF command line interface can be used to interact with the project. The following commands are available:

```
zkWF is a zero-knowledge workflow system

Usage:
  zkwf [command]

Available Commands:
  compile      Compile a BPMN file into a zero-knowledge circuit
  completion   Generate the autocompletion script for the specified shell
  deploy-ecdh  Deploy the ECDH contract with predefined public keys
  fill-inputs  Fill inputs for the zkWF circuit
  generate-key Generate a new eddsa key pair
  help         Help about any command
  prove        Prove a statement using a given circuit and witness
  setup        Setup a zero-knowledge circuit
  sign         Sign a given input
  witness      Generate a witness for a given input

Flags:
  -h, --help   help for zkwf
```

