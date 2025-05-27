package main

import ( 
	"fmt"
	"os"
	"bufio"
	"strings"
)
var lines [3]string

func main() {
    scan := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3 i++ {
		scan.Scan()
		lines[i] = scan.Text()
	}
	for _, line := range lines {
		fmat.Println(line)
	}
}