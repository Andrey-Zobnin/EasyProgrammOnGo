package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [4]string

	for i := 0; i < 4; i++ {
		scanner.Scan()
		lines[i] = scanner.Text()
	}

	fmt.Println(lines[1] + lines[0] + lines[2] + lines[0] + lines[3])
}
