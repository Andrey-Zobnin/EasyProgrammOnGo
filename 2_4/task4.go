package main

import "fmt"

func main() {
	var nums [3]int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&nums[i])
	}
	result := nums[0] * nums[1] * nums[2]
	fmt.Println(result)
}
