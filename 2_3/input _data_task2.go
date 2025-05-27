package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReaderSize(os.Stdin)
	var n int
	fmt.Fscan(read, &n)

}
