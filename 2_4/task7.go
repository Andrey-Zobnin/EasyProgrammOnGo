package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	nums := [3]int{a, a + 1, a + 2}
	for i := 0; i < 3; i++ {
		fmt.Println(nums[0])
		fmt.Println(nums[1])
		fmt.Println(nums[2])
	}

}
