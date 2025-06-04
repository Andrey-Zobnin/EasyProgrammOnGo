package main

import "fmt"

func main() {
	var meta [4]int

	for i := 0; i < 4; i++ {
		fmt.Scan(&meta[i])
	}

	totalCost := (meta[0] + meta[1] + meta[2] + meta[3]) * 3

	fmt.Println(totalCost)
}
