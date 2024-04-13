package calculations

import (
	"log"
	"strconv"
	"strings"
)

func ExctractInput(input string) (int, int, string) {
	mathString := strings.Split(input, " ")

	firstDigit, err := strconv.Atoi(mathString[0])
	if err != nil {
		log.Fatal(err)
	}

	secDigit, err := strconv.Atoi(strings.TrimSpace(mathString[2]))
	if err != nil {
		log.Fatal(err)
	}

	mathOperator := mathString[1]

	return firstDigit, secDigit, mathOperator
}
