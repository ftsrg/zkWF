package expressions

import (
	"fmt"
	"log"
	"strconv"

	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
)

func EvaluateExpression(api frontend.API, expression string, values []frontend.Variable, mapping map[string]int) frontend.Variable {
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

	leftValue, err := calculateValue(api, leftTokens, values, mapping)
	if err != nil {
		log.Fatalln("Failed to calculate left value:", err)
	}
	api.Println("Left value:", leftValue)

	rightValue, err := calculateValue(api, rightTokens, values, mapping)
	if err != nil {
		log.Fatalln("Failed to calculate right value:", err)
	}
	api.Println("Right value:", rightValue)

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

func calculateValue(api frontend.API, tokens []common.Token, values []frontend.Variable, mapping map[string]int) (frontend.Variable, error) {
	var result frontend.Variable = 0
	var currentValue frontend.Variable = 0
	if len(tokens) == 1 {
		switch tokens[0].Type {
		case "VAR":
			return values[mapping[tokens[0].Value]], nil
		case "NUM":
			num, err := strconv.ParseUint(tokens[0].Value, 10, 64)
			if err != nil {
				return nil, err
			}
			return num, nil
		default:
			return 0, nil
		}
	}
	for _, token := range tokens {
		api.Println("Token:", token)
		if token.Type == "VAR" {
			fmt.Println("Token value:", token.Value)
			fmt.Println("Values:", values)
			currentValue = api.Add(currentValue, values[mapping[token.Value]])
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
