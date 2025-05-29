package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [3]string

	for i := 0; i < 4; i++ {
		scanner.Scan()
		lines[i] = scanner.Text()
	}

	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])
}
