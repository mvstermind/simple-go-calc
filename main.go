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
	numToString := strconv.FormatFloat(num, 'f', 2, 64)
	numToString = strings.TrimSuffix(numToString, ".00")
	fmt.Println(numToString)

}
