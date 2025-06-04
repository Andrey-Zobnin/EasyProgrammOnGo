package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 5
	str := strconv.Itoa(num)
	fmt.Println(str + str) // конкатенация вывод 55
}
