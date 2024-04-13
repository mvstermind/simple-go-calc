package main

import (
	"calculator/calculations"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num, err := calculations.GetDigits()
	if err != nil {
		os.Exit(1)
	}
	// Round number to 2 places after . and print output
	numToString := strconv.FormatFloat(num, 'f', 2, 64)
	numToString = strings.TrimSuffix(numToString, ".00")
	fmt.Println(numToString)

}
