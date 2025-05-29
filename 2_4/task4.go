package main

import (
	"bufio"
	"os"
)

func main() {

	var inputData [3]string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {

		if scanner.Scan() {
			inputData[i] = scanner.Text()

		}
	}
	// TODO create math with string, else refzctor code with zero
	for i := 0; i < 3; i++ {
		// name = strconv.Atoi(inputData[i])
		// fmt.Println(name)
	}

}
