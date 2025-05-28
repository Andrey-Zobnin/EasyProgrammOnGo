package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//for i := 0; i > 0; i++ {}
	// _ = scanner.Scan()

	scanner.Scan()
	meta1 := scanner.Text()

	scanner.Scan()
	meta2 := scanner.Text()

	scanner.Scan()
	meta3 := scanner.Text()
	// output in random order
	fmt.Println(meta1)
	fmt.Println(meta3)
	fmt.Println(meta2)

}
