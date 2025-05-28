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
