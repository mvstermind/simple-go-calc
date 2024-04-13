package calculations

import (
	"log"
	"strconv"
	"strings"
)

func ExctractInput(input string) (float64, float64, string) {
	mathString := strings.Split(input, " ")

	firstDigit, err := strconv.ParseFloat(mathString[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	mathString[2] = strings.TrimSpace(mathString[2])
	secDigit, err := strconv.ParseFloat(mathString[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	mathOperator := mathString[1]

	return firstDigit, secDigit, mathOperator
}
