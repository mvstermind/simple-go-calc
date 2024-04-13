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

	FirstNumString = strings.TrimSpace(FirstNumString)
	FirstNum, err := strconv.Atoi(FirstNumString)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Type second number: ")
	SecondNumString, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	SecondNumString = strings.TrimSpace(SecondNumString)
	SecondNum, err := strconv.Atoi(SecondNumString)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Math operator: ")
	fmt.Scan(&MathSign)

	return doMaths(FirstNum, SecondNum, MathSign), err

}

func doMaths(num1, num2 int, operator string) int {
	var Output int
	switch {
	case operator == "+":
		Output = num1 + num2
	case operator == "-":
		Output = num1 - num2
	case operator == "*":
		Output = num1 * num2
	case operator == "/":
		Output = num1 / num2
	}
	return Output
}
