package main

import ("fmt") 


func main() {
	var a string
	// указатель на переменную принимает команда fmt.Scan(&a)
	fmt.Scan(&a)
	fmt.Print("Привет, ", a)

}