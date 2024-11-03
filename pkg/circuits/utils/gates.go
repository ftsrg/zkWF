package utils

import "github.com/consensys/gnark/frontend"

func IsEqual(api frontend.API, a, b frontend.Variable) frontend.Variable {
	return api.IsZero(api.Sub(a, b))
}

func Not(api frontend.API, a frontend.Variable) frontend.Variable {
	api.AssertIsBoolean(a)
	return api.Sub(1, a)
}

func LessThan(api frontend.API, a frontend.Variable, b frontend.Variable) frontend.Variable {
	return IsEqual(api, api.Cmp(a, b), -1)
}

func LessEqThan(api frontend.API, a frontend.Variable, b frontend.Variable) frontend.Variable {
	return Not(api, GreaterThan(api, a, b))
}

func GreaterThan(api frontend.API, a frontend.Variable, b frontend.Variable) frontend.Variable {
	return IsEqual(api, api.Cmp(a, b), 1)
}

func GreaterEqThan(api frontend.API, a frontend.Variable, b frontend.Variable) frontend.Variable {
	return Not(api, LessThan(api, a, b))
}
