package calculations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	FirstNumString  string
	SecondNumString string
	MathSign        string
)

func GetDigits() (int, error) {
	fmt.Printf("Type first number: ")
	in := bufio.NewReader(os.Stdin)

	FirstNumString, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	FirstNum, err := strconv.Atoi(FirstNumString)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Type second number: ")
	SecondNumString, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	SecondNum, err := strconv.Atoi(SecondNumString)
	if err != nil {
		return 0, err
	}

	if strings.Contains(FirstNumString, " ") {
		strings.Trim(FirstNumString, " ")
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

	return Output, err
}
