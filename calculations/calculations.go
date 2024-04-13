package calculations

import (
	"fmt"
	"strconv"
)

var (
	FirstNumString  string
	SecondNumString string
	MathSign        string
)

func GetDigits() (int, error) {
	fmt.Printf("Type first number: ")
	_, err := fmt.Scanf("%v", &FirstNumString)
	if err != nil {
		return 0, err
	}

	FirstNum, err := strconv.Atoi(FirstNumString)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Type second number: ")
	_, errr := fmt.Scanf("%v", &SecondNumString)
	if errr != nil {
		return 0, errr
	}

	SecondNum, err := strconv.Atoi(SecondNumString)
	if err != nil {
		return 0, err
	}

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

	result := strconv.Itoa(Output)

	return "Result: " + result
}
