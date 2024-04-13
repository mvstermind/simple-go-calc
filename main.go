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
	firstDigit, secondDigit := calculations.GetDigits(FirstNum, SecondNum)
	fmt.Printf("%v, %v\n", firstDigit, secondDigit)
}
