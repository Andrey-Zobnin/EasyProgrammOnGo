package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var lines [3]string
	for i := 0; i < 3; i++ {
		scan.Scan()
		lines[i] = scan.Text()
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
