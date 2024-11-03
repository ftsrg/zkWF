package common

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Token represents a single parsed unit (variable, operator, or number)
type Token struct {
	Type  string
	Value string
}

// Tokenizer turns an equation string into tokens
func Tokenizer(equation string) ([]Token, error) {
	var tokens []Token
	equation = strings.ReplaceAll(equation, " ", "") // Remove spaces
	i := 0
	for i < len(equation) {
		c := rune(equation[i])
		switch {
		case unicode.IsLetter(c): // Variable
			j := i
			for j < len(equation) && unicode.IsLetter(rune(equation[j])) {
				j++
			}
			tokens = append(tokens, Token{Type: "VAR", Value: equation[i:j]})
			i = j
		case unicode.IsDigit(c): // Number
			j := i
			for j < len(equation) && (unicode.IsDigit(rune(equation[j])) || equation[j] == '.') {
				j++
			}
			tokens = append(tokens, Token{Type: "NUM", Value: equation[i:j]})
			i = j
		case c == '+' || c == '-': // Arithmetic operators
			tokens = append(tokens, Token{Type: "OP", Value: string(c)})
			i++
		case c == '=' || c == '!' || c == '<' || c == '>': // Comparators
			j := i
			if j+1 < len(equation) && equation[j+1] == '=' {
				j++
			}
			tokens = append(tokens, Token{Type: "COMP", Value: equation[i : j+1]})
			i = j + 1
		default:
			return nil, fmt.Errorf("unrecognized character: %c", c)
		}
	}
	return tokens, nil
}

// ExpressionParser evaluates a tokenized equation
func ExpressionParser(tokens []Token, vars map[string]int64) (bool, error) {
	// This is a simple recursive-descent parser implementation
	if len(tokens) < 3 {
		return false, fmt.Errorf("invalid expression")
	}

	// Handle expressions like "a + b < c"
	var left, right int64
	var err error

	// Parse the left side (variable or expression)
	if tokens[0].Type == "VAR" {
		left, err = getVariableValue(tokens[0].Value, vars)
	} else if tokens[0].Type == "NUM" {
		left, err = strconv.ParseInt(tokens[0].Value, 10, 32)
	} else {
		return false, fmt.Errorf("invalid token: %v", tokens[0])
	}

	if err != nil {
		return false, err
	}

	// Parse the right side (after the comparator)
	rightStart := 2
	if len(tokens) > 3 && (tokens[1].Type == "OP" || tokens[2].Type == "OP") {
		rightStart = 3
	}

	if tokens[rightStart].Type == "VAR" {
		right, err = getVariableValue(tokens[rightStart].Value, vars)
	} else if tokens[rightStart].Type == "NUM" {
		right, err = strconv.ParseInt(tokens[rightStart].Value, 10, 64)
	} else {
		return false, fmt.Errorf("invalid token: %v", tokens[rightStart])
	}

	if err != nil {
		return false, err
	}

	// Compare based on the operator
	switch tokens[1].Value {
	case "==":
		return left == right, nil
	case "!=":
		return left != right, nil
	case "<":
		return left < right, nil
	case ">":
		return left > right, nil
	case "<=":
		return left <= right, nil
	case ">=":
		return left >= right, nil
	default:
		return false, fmt.Errorf("unknown operator: %s", tokens[1].Value)
	}
}

// Helper to fetch variable values
func getVariableValue(varName string, vars map[string]int64) (int64, error) {
	if val, exists := vars[varName]; exists {
		return val, nil
	}
	return 0, fmt.Errorf("variable %s not found", varName)
}
