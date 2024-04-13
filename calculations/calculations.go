package calculations

import "fmt"

func GetDigits(FirstNum, SecondNum int, MathSign string) int {
	fmt.Printf("Type first number: ")
	fmt.Scan(&FirstNum)

	fmt.Printf("Type second number: ")
	fmt.Scan(&SecondNum)

	fmt.Printf("Math operator: ")
	fmt.Scan(&MathSign)

	var Output int
	switch {
	case MathSign == "+":
		Output = FirstNum + SecondNum
	case MathSign == "-":
		Output = FirstNum - SecondNum
	case MathSign == "*":
		Output = FirstNum * SecondNum
	case MathSign == "/":
		Output = FirstNum / SecondNum
	}

	return Output
}
