package main

import "fmt"

func main() {
	fmt.Println("Welcome!")
	fmt.Println("Enter filename: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
