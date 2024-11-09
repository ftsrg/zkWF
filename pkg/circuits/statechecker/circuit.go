package statechecker

import (
	"errors"
	"fmt"

	ecc_twisted "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	gnark_mimc "github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gnark/std/signature/eddsa"
	"github.com/ftsrg/zkWF/pkg/circuits/gmimc"
	"github.com/ftsrg/zkWF/pkg/circuits/hkdf"
	"github.com/ftsrg/zkWF/pkg/circuits/lifecycle"
	"github.com/ftsrg/zkWF/pkg/circuits/mimc"
	"github.com/ftsrg/zkWF/pkg/circuits/proofofownership"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
	"github.com/ftsrg/zkWF/pkg/model"
)

type Circuit struct {
	Model            *model.BPMNGraph
	VariableMapping  map[string]int
	ParticipantIndex int // Index of the participant in the participants array. It is easier to use this than the public key
	State_curr       State
	State_new        State
	HashCurr         frontend.Variable `gnark:",public"`
	HashNew          frontend.Variable `gnark:",public"`
	Keys             KeyPair
	Signature        eddsa.Signature `gnark:",public"`
	Key              []frontend.Variable
	Encrypted        []frontend.Variable `gnark:",public"`
	Deposit          frontend.Variable   `gnark:",public"`
	Withdrawal       frontend.Variable   `gnark:",public"`
}

type KeyPair struct {
	PublicKey  eddsa.PublicKey `gnark:",public"`
	PrivateKey [2]frontend.Variable
}

type State struct {
	States    []frontend.Variable
	Variables []frontend.Variable
	Messages  []frontend.Variable
	Balances  []frontend.Variable
	Radomness frontend.Variable
}

func (circuit Circuit) Define(api frontend.API) error {
	N := len(circuit.State_curr.States)
	if len(circuit.State_curr.States) != len(circuit.State_new.States) && len(circuit.State_curr.States) != N {
		return errors.New("state_curr and state_new must have the same length")
	}

	executables := circuit.Model.GetExecutableNodes() // Executable = Task, IntermediateCatchEvent, IntermediateThrowEvent; Always one input and one output
	if len(executables) != N {
		return fmt.Errorf("number of executable nodes must be %d, got %d", N, len(executables))
	}

	// Proof of ownership
	proofofownership.ProofOfOwnership(api, circuit.Keys.PublicKey, circuit.Keys.PrivateKey)
	api.AssertIsEqual(circuit.Model.Participants[circuit.ParticipantIndex].PublicKey[0], circuit.Keys.PublicKey.A.X)
	api.AssertIsEqual(circuit.Model.Participants[circuit.ParticipantIndex].PublicKey[1], circuit.Keys.PublicKey.A.Y) // Just to make sure, that the right index is used

	// Randomness check
	api.AssertIsDifferent(circuit.State_curr.Radomness, circuit.State_new.Radomness)

	same := make([]frontend.Variable, N)
	complated := make([]frontend.Variable, N)
	stateChanged := make([]frontend.Variable, N) // State changed from curr to new, but not completed
	for i, curr := range circuit.State_curr.States {
		new := circuit.State_new.States[i]
		same[i] = utils.IsEqual(api, curr, new)
		complated[i] = api.Select(same[i], common.FALSE, utils.IsEqual(api, new, common.STATE_COMPLETED))

		lifecycleCheck := lifecycle.LifecycleCheck(api, curr, new)
		api.AssertIsEqual(api.Select(same[i], common.TRUE, lifecycleCheck), common.TRUE)

		stateChanged[i] = api.Select(same[i], common.FALSE, api.Select(complated[i], common.FALSE, common.TRUE))
	}

	comp_sum := make([]frontend.Variable, N)
	diff_sum := make([]frontend.Variable, N)

	comp_sum[0] = complated[0]
	diff_sum[0] = utils.Not(api, same[0])
	for i := 1; i < N; i++ {
		comp_sum[i] = api.Add(comp_sum[i-1], complated[i])
		diff_sum[i] = api.Add(diff_sum[i-1], utils.Not(api, same[i]))
	}
	api.Println("comp_sum: ", comp_sum[N-1])
	api.Println("diff_sum: ", diff_sum[N-1])
	noChange := utils.IsEqual(api, diff_sum[N-1], 0)
	api.Println("No change: ", noChange)
	isActivated := utils.IsEqual(api, comp_sum[N-1], 1)
	api.Println("Is activated: ", isActivated)
	justOne := api.Select(isActivated, 0, utils.IsEqual(api, diff_sum[N-1], 1))
	api.Println("Just one: ", justOne)
	validChanges := api.Select(noChange, common.TRUE, api.Xor(isActivated, justOne))

	api.AssertIsEqual(validChanges, common.TRUE)

	readyInTheRightTime := make([]frontend.Variable, N)
	completedAndTokenForwarded := make([]frontend.Variable, N)
	//Check if the task is ready and activated by a completed task
	var involvedWithPayment []frontend.Variable = make([]frontend.Variable, len(circuit.State_curr.Balances))
	for i := 0; i < len(involvedWithPayment); i++ {
		involvedWithPayment[i] = common.FALSE
	}

	for _, task := range executables {
		index := indexOf(executables, task)
		incomingNode := task.Incoming[0].SourceRef
		gotReady := api.And(api.Select(utils.IsEqual(api, circuit.State_new.States[index], common.STATE_READY), common.TRUE, common.FALSE), stateChanged[index])
		changed := utils.Not(api, same[index])
		pubKeySameX := utils.IsEqual(api, circuit.Keys.PublicKey.A.X, task.Owner.PublicKey[0])
		pubKeySameY := utils.IsEqual(api, circuit.Keys.PublicKey.A.Y, task.Owner.PublicKey[1])
		pubKeySame := api.And(pubKeySameX, pubKeySameY)
		api.AssertIsEqual(api.Select(changed, pubKeySame, common.TRUE), common.TRUE)

		// Stuff got completed
		switch task.Type {
		case model.INTERMEDIATE_THROW_EVENT:
			messageId := circuit.Model.MessageMap[task.ID]
			messageSame := utils.IsEqual(api, circuit.State_curr.Messages[messageId], circuit.State_new.Messages[messageId])
			messageUpdated := utils.Not(api, messageSame)
			api.AssertIsEqual(api.Select(complated[index], messageUpdated, common.TRUE), common.TRUE)
		case model.INTERMEDIATE_CATCH_EVENT:
			// TODO: Check if the message is sent by the THROW event (check if the event is completed)
		case model.PAYMENT_TASK:
			newBalance := circuit.State_new.Balances[task.Payment.Receiver]
			expectedBalance := api.Add(circuit.State_curr.Balances[task.Payment.Receiver], task.Payment.Amount)
			api.AssertIsEqual(api.Select(complated[index], utils.IsEqual(api, newBalance, expectedBalance), common.TRUE), common.TRUE)
			newBalance = circuit.State_new.Balances[task.Owner.ID]
			expectedBalance = api.Sub(circuit.State_curr.Balances[task.Owner.ID], task.Payment.Amount)
			api.AssertIsEqual(api.Select(complated[index], utils.IsEqual(api, newBalance, expectedBalance), common.TRUE), common.TRUE)
			involvedWithPayment[task.Owner.ID] = api.Select(complated[index], common.TRUE, involvedWithPayment[task.Owner.ID])
			involvedWithPayment[task.Payment.Receiver] = api.Select(complated[index], common.TRUE, involvedWithPayment[task.Payment.Receiver])
		}

		switch incomingNode.Type {
		case model.START_EVENT:
			readyInTheRightTime[index] = 0
		case model.PARALLEL_GATEWAY:
			/*
				In case of a ParallelGateway, the task is ready if all incoming nodes are completed
			*/
			incomingNodesCompleted := allPreviousNodesCompleted(api, circuit, task)
			readyInTheRightTime[index] = api.And(gotReady, incomingNodesCompleted)
		case model.EXCLUSIVE_GATEWAY, model.INCLUSIVE_GATEWAY:
			/*
				In case of an ExclusiveGateway, the task is ready if the incoming node (one) is completed
			*/
			incomingNodesOneCompleted := atLeastOnePreviousNodeCompleted(api, circuit, incomingNode)
			readyInTheRightTime[index] = api.And(gotReady, incomingNodesOneCompleted)
		default:
			incomingNodeIndex := indexOf(executables, incomingNode)

			readyInTheRightTime[index] = api.And(gotReady, complated[incomingNodeIndex])
		}

		outNode := task.Outgoing[0].TargetRef
		switch outNode.Type {
		case model.END_EVENT:
			completedAndTokenForwarded[index] = complated[index]
		case model.PARALLEL_GATEWAY:
			/*
				In case of a ParallelGateway, the task is completed if all outgoing nodes are ready when all the "paired" tasks are completed
				If one pair is not completed, the task is completed but without activating the next task
			*/

			pairs := task.GetPairs()
			pairsCompleted := make([]frontend.Variable, len(pairs))
			for i, pair := range pairs {
				pairIndex := indexOf(executables, pair)
				if i == 0 {
					pairsCompleted[i] = utils.IsEqual(api, circuit.State_new.States[pairIndex], common.STATE_COMPLETED)
				} else {
					pairsCompleted[i] = api.And(pairsCompleted[i-1], utils.IsEqual(api, circuit.State_new.States[pairIndex], common.STATE_COMPLETED))
				}
			}

			outgoingNodesGotReady := parallelGatewayActivated(api, circuit, same, outNode)

			/*completedAndTokenForwarded[index] = api.Select(complated[index], api.Select(pairsCompleted[len(pairs)-1], outgoingNodesGotReady[len(outgoingNodes)-1], utils.Not(api, outgoingNodesGotReady[len(outgoingNodes)-1])), common.FALSE)*/
			if len(pairs) > 0 {
				completedAndTokenForwarded[index] = api.Select(complated[index], api.Select(pairsCompleted[len(pairs)-1], outgoingNodesGotReady, utils.Not(api, outgoingNodesGotReady)), common.FALSE)
				api.Println("Task: ", task.ID, " Completed: ", complated[index], " Pairs completed: ", pairsCompleted[len(pairs)-1], " Outgoing nodes ready: ", outgoingNodesGotReady, " Completed and token forwarded: ", completedAndTokenForwarded[index])
			} else {
				completedAndTokenForwarded[index] = api.Select(complated[index], outgoingNodesGotReady, common.FALSE)
			}
		case model.EXCLUSIVE_GATEWAY:
			/*
				In case of an ExclusiveGateway, the task is completed if one of the outgoing nodes is ready when the incoming node is completed AND
				the expression on that edge is evaluated to true
			*/
			outgoingNodesGotReady := exclusiveGatewayActivated(api, circuit, same, outNode)
			completedAndTokenForwarded[index] = api.Select(complated[index], outgoingNodesGotReady, common.FALSE)
		case model.INCLUSIVE_GATEWAY:
			outgoingNodesGotReady := inclusiveGatewayActivated(api, circuit, same, outNode)
			completedAndTokenForwarded[index] = api.Select(complated[index], outgoingNodesGotReady, common.FALSE)
		default:
			outNodeIndex := indexOf(executables, outNode)
			gotReady := api.And(api.Select(utils.IsEqual(api, circuit.State_new.States[outNodeIndex], common.STATE_READY), common.TRUE, common.FALSE), stateChanged[outNodeIndex])
			completedAndTokenForwarded[index] = api.And(gotReady, complated[index])
		}

		// One and only one of the following conditions must be 1 (true): readyInTheRightTime, completedAndTokenForwarded, stateChanged, same
		api.Println("Task: ", task.ID, " Ready: ", readyInTheRightTime[index], " Completed: ", completedAndTokenForwarded[index], " State changed: ", stateChanged[index], "Same: ", same[index])

		readyOrStateChange := api.Xor(readyInTheRightTime[index], api.And(stateChanged[index], utils.Not(api, gotReady)))
		complatedOrSame := api.Xor(completedAndTokenForwarded[index], same[index])
		api.Println("Task: ", task.ID, " Final: ", api.Xor(readyOrStateChange, complatedOrSame))
		api.AssertIsEqual(api.Xor(readyOrStateChange, complatedOrSame), common.TRUE)
	}

	checkPayments(api, circuit, involvedWithPayment)

	// Check if the hash is correct
	stateCompressed := compressedState(api, circuit.State_curr)

	//_ = circuits.MultiMiMC7(api, 91, hashInput, 0)
	hash := mimc.MultiMiMC5(api, 91, stateCompressed, 0)
	api.Println("Hash: ", hash)
	api.AssertIsEqual(hash, circuit.HashCurr)
	// Check if the hash is correct
	stateCompressed_new := compressedState(api, circuit.State_new)
	api.Println("State compressed curr: ", stateCompressed_new[0])
	api.Println("State compressed length: ", len(stateCompressed_new))

	hash2 := mimc.MultiMiMC5(api, 91, stateCompressed_new, 0)
	api.Println("Hash2: ", hash2)
	api.AssertIsEqual(hash2, circuit.HashNew)

	// Verify signature
	edCurve, err := twistededwards.NewEdCurve(api, ecc_twisted.BN254)
	if err != nil {
		return fmt.Errorf("failed to create twisted edwards curve: %v", err)
	}

	mimc, err := gnark_mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	api.Println("Verifying signature")
	err = eddsa.Verify(edCurve, circuit.Signature, hash2, circuit.Keys.PublicKey, &mimc)
	if err != nil {
		return fmt.Errorf("failed to verify signature: %v", err)
	}
	api.Println("R:", circuit.Signature.R.X, circuit.Signature.R.Y)
	api.Println("S:", circuit.Signature.S)
	api.Println("Balances length: ", len(circuit.State_new.Balances))

	salt := make([]frontend.Variable, 1)
	salt[0] = 0
	info := make([]frontend.Variable, 2)
	info[0] = circuit.State_curr.Radomness
	info[1] = circuit.State_new.Radomness
	ikm := make([]frontend.Variable, 1)
	ikm[0] = circuit.Key[0]

	key_new := hkdf.Hkdf(api, salt, ikm, info, 2)

	api.Println("Key new: ", key_new[0])

	encrypted := gmimc.Encrypt(api, stateCompressed_new, key_new, gmimc.GetGMiMCRounds(len(stateCompressed)))
	for i, e := range encrypted {
		api.Println("Encrypted: ", i, " ", e)
		api.AssertIsEqual(e, circuit.Encrypted[i])
	}

	return nil
}

func compressedState(api frontend.API, state State) []frontend.Variable {
	stateCompressed := make([]frontend.Variable, len(state.Variables)+len(state.Messages)+len(state.Balances)+2)
	stateCompressed[0] = utils.CompressToFieldElement(api, state.States)
	stateCompressed[1] = state.Radomness
	i := 2
	for k, v := range state.Variables {
		api.Println("Variable:", v, k)
		stateCompressed[i] = v
		i++
	}

	for _, m := range state.Messages {
		stateCompressed[i] = m
		i++
	}

	for _, b := range state.Balances {
		stateCompressed[i] = b
		i++
	}

	return stateCompressed
}

func checkPayments(api frontend.API, circuit Circuit, involvedWithPayment []frontend.Variable) {
	handleDeposits(api, circuit, involvedWithPayment)
	handleWithdrawals(api, circuit, involvedWithPayment)

	for i, paymentInvolved := range involvedWithPayment {
		same := utils.IsEqual(api, circuit.State_curr.Balances[i], circuit.State_new.Balances[i])
		api.AssertIsEqual(api.Select(paymentInvolved, common.TRUE, same), common.TRUE)
	}
}

func handleDeposits(api frontend.API, circuit Circuit, involvedWithPayment []frontend.Variable) {
	noDeposit := utils.IsEqual(api, circuit.Deposit, 0)
	depositDone := utils.Not(api, noDeposit)
	involvedWithPayment[circuit.ParticipantIndex] = api.Select(depositDone, common.TRUE, involvedWithPayment[circuit.ParticipantIndex])
	expectedBalance := api.Add(circuit.State_curr.Balances[circuit.ParticipantIndex], circuit.Deposit)
	actualBalance := circuit.State_new.Balances[circuit.ParticipantIndex]
	api.AssertIsEqual(api.Select(depositDone, utils.IsEqual(api, actualBalance, expectedBalance), common.TRUE), common.TRUE)
}

func handleWithdrawals(api frontend.API, circuit Circuit, involvedWithPayment []frontend.Variable) {
	noWithdrawal := utils.IsEqual(api, circuit.Withdrawal, 0)
	withdrawalDone := utils.Not(api, noWithdrawal)
	involvedWithPayment[circuit.ParticipantIndex] = api.Select(withdrawalDone, common.TRUE, involvedWithPayment[circuit.ParticipantIndex])
	expectedBalance := api.Sub(circuit.State_curr.Balances[circuit.ParticipantIndex], circuit.Withdrawal)
	actualBalance := circuit.State_new.Balances[circuit.ParticipantIndex]
	api.AssertIsEqual(api.Select(withdrawalDone, utils.IsEqual(api, actualBalance, expectedBalance), common.TRUE), common.TRUE)
}

func indexOf(nodes []*model.Node, node *model.Node) int {
	for i, n := range nodes {
		if n == node {
			return i
		}
	}
	return -1
}
