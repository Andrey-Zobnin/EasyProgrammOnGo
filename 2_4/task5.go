package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	applesPerStudent := K / N
	fmt.Println(applesPerStudent)
}
