// fmt.Println(name, " - лучшая книга!")
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	name, _ := read.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name, "- лучшая книга!")
}
