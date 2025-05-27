package main

import ("fmt") 

var a string

func main() {
	// a := ""
	// указатель на переменную принимает команда fmt.Scan(&a)
	fmt.Scan(&a)
	fmt.Print("Привет, ", a)
	
}