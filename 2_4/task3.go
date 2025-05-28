package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	num = num * num * num
	num = num * num
	fmt.Print(num)

}
