package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	meta := make([]strings, 0, 3)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		meta = append(lines, scanner.Text())
	}

	for i := len(meta) - 1; i >= 0; i-- {
		fmt.Println(meta[i])
	}
}
