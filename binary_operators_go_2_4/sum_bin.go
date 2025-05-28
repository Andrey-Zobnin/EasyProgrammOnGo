package main

//import (
//"fmt"
//"os"
// )

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
