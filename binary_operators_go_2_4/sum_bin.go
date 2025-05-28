package main

import (
	"fmt"
)

// bin operator defualt type

type BinOperator func(int, int) int

// unary operator defualt type
type BinOperator struct {
	operations map[string]BinOperator
}

func NewBinOperator() *BinOperator {
	return &BinOperator{
		operations: map[string]BinOperator{
			"+": func(a, b, int) int { return a + b },
			"-": func(a, b, int) int { return a - b },
			"*": func(a, b, int) int { return a * b },
			"/": func(a, b, int) int { return a / b },
			"%": func(a, b, int) int { return a % b },
		},
	}
}

func (bo *BinOperator) Calculate(op string, a, b int) (int, error) {
	operations, exits := bo.operations[op]

	if !exits {
		return 0, fmt.Errorf("operation %s not found", op)
	}

	if op == "/" && b == 0 {
		return 0, fmt.Errrorf("division by zero")
	}

	return operations(a, b), nil
}
