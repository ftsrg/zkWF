package statechecker

import (
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
	"github.com/ftsrg/zkWF/pkg/model"
)

// atLeastOnePreviousNodeCompleted checks if at least one previous node is in the completed state
// It is used to check if a node can be activated. It has to check all the previous nodes in the graph.
// It calls the allPerv function for each previous parrallel gateways and the atLeastOnePreviousNodeCompleted function for each previous exclusive/inclusive gateways.
// Tasks or intermediate events are checked inplace.
func atLeastOnePreviousNodeCompleted(api frontend.API, circuit Circuit, node *model.Node) frontend.Variable {

	completed := make([]frontend.Variable, len(node.Incoming))
	for i, prev := range node.Incoming {
		switch prev.SourceRef.Type {
		case model.PARALLEL_GATEWAY:
			completed[i] = allPreviousNodesCompleted(api, circuit, prev.SourceRef)
		case model.EXCLUSIVE_GATEWAY, model.INCLUSIVE_GATEWAY:
			completed[i] = atLeastOnePreviousNodeCompleted(api, circuit, prev.SourceRef)
		default:
			completed[i] = utils.IsEqual(api, circuit.State_new.States[indexOf(circuit.Model.GetExecutableNodes(), prev.SourceRef)], common.STATE_COMPLETED)
		}
	}

	sum := make([]frontend.Variable, len(completed))
	sum[0] = completed[0]
	for i := 1; i < len(completed); i++ {
		sum[i] = api.Add(sum[i-1], completed[i])
	}
	atLeastOne := utils.Not(api, utils.IsEqual(api, sum[len(sum)-1], 0))

	return atLeastOne
}

// allPreviousNodesCompleted checks if all previous nodes are in the completed state. It is similar to the atLeastOnePrev function, but it checks if all previous nodes are completed.
func allPreviousNodesCompleted(api frontend.API, circuit Circuit, node *model.Node) frontend.Variable {

	completed := make([]frontend.Variable, len(node.Incoming))
	for i, prev := range node.Incoming {
		switch prev.SourceRef.Type {
		case model.PARALLEL_GATEWAY:
			completed[i] = allPreviousNodesCompleted(api, circuit, prev.SourceRef)
		case model.EXCLUSIVE_GATEWAY, model.INCLUSIVE_GATEWAY:
			completed[i] = atLeastOnePreviousNodeCompleted(api, circuit, prev.SourceRef)
		default:
			completed[i] = utils.IsEqual(api, circuit.State_new.States[indexOf(circuit.Model.GetExecutableNodes(), prev.SourceRef)], common.STATE_COMPLETED)
		}
	}

	sum := make([]frontend.Variable, len(completed))
	sum[0] = completed[0]
	for i := 1; i < len(completed); i++ {
		sum[i] = api.Add(sum[i-1], completed[i])
	}

	all := utils.IsEqual(api, sum[len(sum)-1], len(completed))

	return all
}
