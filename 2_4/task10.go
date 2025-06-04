package main

import "fmt"

func main() {
	var num int
	if _, err := fmt.Scan(&num); err != nil {
		return
	}
	fmt.Printf("Следующее за числом %d число: %d\n", num, num+1)
	fmt.Printf("Для числа %d предыдущее число: %d\n", num, num-1)

}
