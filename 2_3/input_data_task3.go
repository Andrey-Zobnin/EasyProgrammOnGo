package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var meta [3]string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		meta[i] = scanner.Text()
	}

	for i := 2; i >= 0; i-- {
		fmt.Println(meta[i])
	}
}
