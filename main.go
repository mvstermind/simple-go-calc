package main

import (
	"calculator/calculations"
	"fmt"
)

var (
	FirstNum  int
	MathSign  string
	SecondNum int
)

func main() {
	firstDigit := calculations.GetDigit1(FirstNum)
	fmt.Printf("%v\n", firstDigit)
}
