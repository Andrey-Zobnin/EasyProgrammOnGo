package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//for i := 0; i > 0; i++ {}
	_ = scanner.Scan()
	meta := scanner.Text()
	fmt.Println(meta)
}
