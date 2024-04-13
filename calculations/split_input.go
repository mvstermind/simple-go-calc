package calculations

import (
	"strings"
)

func ExctractInput(input string) (float64, float64, string) {
	mathString := strings.Split(input, " ")

	firstDigit := ConvertToFloat(mathString, 0)
	secDigit := ConvertToFloat(mathString, 2)
	mathOperator := mathString[1]

	return firstDigit, secDigit, mathOperator
}
