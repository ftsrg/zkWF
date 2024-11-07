package statechecker

import (
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/expressions"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
	"github.com/ftsrg/zkWF/pkg/model"
)

func exclusiveGatewayActivated(api frontend.API, circuit Circuit, same []frontend.Variable, node *model.Node) frontend.Variable {
	sum := clusiveSum(api, circuit, same, node)
	return utils.IsEqual(api, sum, 1)
}

func parallelGatewayActivated(api frontend.API, circuit Circuit, same []frontend.Variable, node *model.Node) frontend.Variable {
	activated := make([]frontend.Variable, len(node.Outgoing))
	for i, next := range node.Outgoing {
		switch next.TargetRef.Type {
		case model.PARALLEL_GATEWAY:
			activated[i] = parallelGatewayActivated(api, circuit, same, next.TargetRef)
		case model.EXCLUSIVE_GATEWAY:
			activated[i] = exclusiveGatewayActivated(api, circuit, same, next.TargetRef)
		case model.INCLUSIVE_GATEWAY:
			activated[i] = inclusiveGatewayActivated(api, circuit, same, next.TargetRef)
		case model.END_EVENT:
			activated[i] = common.TRUE

		default:
			index := indexOf(circuit.Model.GetExecutableNodes(), next.TargetRef)
			activated[i] = api.Select(same[index], common.FALSE, utils.IsEqual(api, circuit.State_new.States[index], common.STATE_READY))

		}
	}

	sum := make([]frontend.Variable, len(activated))
	sum[0] = activated[0]
	for i := 1; i < len(activated); i++ {
		sum[i] = api.Add(sum[i-1], activated[i])
	}
	return utils.IsEqual(api, sum[len(sum)-1], len(activated))
}

func inclusiveGatewayActivated(api frontend.API, circuit Circuit, same []frontend.Variable, node *model.Node) frontend.Variable {
	sum := clusiveSum(api, circuit, same, node)

	return utils.Not(api, utils.IsEqual(api, sum, 0))
}

func clusiveSum(api frontend.API, circuit Circuit, same []frontend.Variable, node *model.Node) frontend.Variable {
	activated := make([]frontend.Variable, len(node.Outgoing))
	expressionTrue := make([]frontend.Variable, len(node.Outgoing))
	for i, next := range node.Outgoing {

		if next.Name != "" {
			expressionTrue[i] = expressions.EvaluateExpression(api, next.Name, circuit.State_new.Variables, circuit.VariableMapping)
			api.Println("Expression: ", next.Name, " = ", expressionTrue[i])
		} else {
			expressionTrue[i] = common.TRUE
		}
		var activated_helper frontend.Variable
		switch next.TargetRef.Type {
		case model.PARALLEL_GATEWAY:
			activated_helper = parallelGatewayActivated(api, circuit, same, next.TargetRef)
		case model.EXCLUSIVE_GATEWAY:
			activated_helper = exclusiveGatewayActivated(api, circuit, same, next.TargetRef)
		case model.INCLUSIVE_GATEWAY:
			activated_helper = inclusiveGatewayActivated(api, circuit, same, next.TargetRef)
		default:

			index := indexOf(circuit.Model.GetExecutableNodes(), next.TargetRef)
			activated_helper = api.Select(same[index], common.FALSE, utils.IsEqual(api, circuit.State_new.States[index], common.STATE_READY))
		}

		activated[i] = api.And(expressionTrue[i], activated_helper)
	}

	sum := make([]frontend.Variable, len(activated))
	sum[0] = activated[0]
	for i := 1; i < len(activated); i++ {
		sum[i] = api.Add(sum[i-1], activated[i])
	}

	return sum[len(sum)-1]
}
