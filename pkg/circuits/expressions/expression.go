package expressions

import (
	"fmt"
	"log"
	"strconv"

	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
)

func EvaluateExpression(api frontend.API, expression string, values map[string]frontend.Variable) frontend.Variable {
	tokens, err := common.Tokenizer(expression)
	if err != nil {
		log.Fatalln("Failed to tokenize expression:", err)
	}

	compIndex := -1
	for i, token := range tokens {
		if token.Type == "COMP" {
			compIndex = i
			break
		}
	}

	if compIndex == -1 {
		fmt.Printf("No comparator found in expression: %s %v\n", expression, tokens)
		log.Fatalln("No comparator found in expression")
	}

	leftTokens := tokens[:compIndex]
	rightTokens := tokens[compIndex+1:]

	leftValue, err := calculateValue(api, leftTokens, values)
	if err != nil {
		log.Fatalln("Failed to calculate left value:", err)
	}

	rightValue, err := calculateValue(api, rightTokens, values)
	if err != nil {
		log.Fatalln("Failed to calculate right value:", err)
	}

	switch tokens[compIndex].Value {
	case "==":
		return utils.IsEqual(api, leftValue, rightValue)
	case "!=":
		return utils.Not(api, utils.IsEqual(api, leftValue, rightValue))
	case "<":
		return utils.LessThan(api, leftValue, rightValue)
	case ">":
		return utils.GreaterThan(api, leftValue, rightValue)
	case "<=":
		return utils.LessEqThan(api, leftValue, rightValue)
	case ">=":
		return utils.GreaterEqThan(api, leftValue, rightValue)
	default:
		log.Fatalln("Unknown operator:", tokens[compIndex].Value)
	}

	return 0
}

func calculateValue(api frontend.API, tokens []common.Token, values map[string]frontend.Variable) (frontend.Variable, error) {
	var result frontend.Variable = 0
	var currentValue frontend.Variable = 0
	for _, token := range tokens {
		if token.Type == "VAR" {
			fmt.Println("Token value:", token.Value)
			fmt.Println("Values:", values)
			currentValue = api.Add(currentValue, values[token.Value])
		} else if token.Type == "NUM" {
			// Convert string to frontend.Variable
			num, err := strconv.ParseUint(token.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			currentValue = api.Add(currentValue, num)
		} else if token.Type == "OP" {
			if token.Value == "+" {
				result = api.Add(result, currentValue)
			} else if token.Value == "-" {
				result = api.Sub(result, currentValue)
			} else {
				return nil, nil
			}
			currentValue = 0
		}
	}

	return result, nil
}
