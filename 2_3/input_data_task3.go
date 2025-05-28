package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	//for i := 0; i > 0; i++ {}
	_ = scan.Scan()
	meta := scan.Text()
	fmt.Println(meta)
}
