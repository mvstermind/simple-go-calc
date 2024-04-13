package main

import (
	"calculator/calculations"
	"fmt"
)

var (
	FirstNum  int
	SecondNum int
	MathSign  string
)

func main() {
	fmt.Printf("%v\n", calculations.GetDigits(FirstNum, SecondNum, MathSign))
}
