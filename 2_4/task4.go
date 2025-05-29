package main

import "fmt"

func main() {
	var num0, num1, num2 int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&num0, &num1, &num2)
	}
	result := num0 * num1 * num2
	fmt.Println(result)
}
