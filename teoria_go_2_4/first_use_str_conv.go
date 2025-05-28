package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "10"
	num, _ := strconv.Atoi(str)
	fmt.Println(num * 2) // print 20
}
