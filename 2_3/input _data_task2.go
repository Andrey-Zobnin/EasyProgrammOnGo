package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReaderSize(os.Stdin)
	metaData, _ := read.ReadString('\n')
	metaData = strings.TrimSpace(metaData)
	// output
	fmt.Println(metaData)

}
