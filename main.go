package main

import (
	"calculator/calculations"
	"fmt"
	"os"
)

func main() {
	num, err := calculations.GetDigits()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(num)
}
