package main

import (
	//"calculator/calculations"
	"fmt"
)

var input string

func main() {
	fmt.Print("Type two numbers separated with: ")
	fmt.Scanln(&input)
	fmt.Printf(input)
}
