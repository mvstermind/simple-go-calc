package calculations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetDigits() (int, error) {
	fmt.Printf("Type a thing to calculate: ")
	in := bufio.NewReader(os.Stdin)

	mathCalculation, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	mathString := strings.Split(mathCalculation, " ")

	FirstNum, err := strconv.Atoi(mathString[0])
	if err != nil {
		return 0, err
	}

	SecondNum, err := strconv.Atoi(strings.TrimSpace(mathString[2]))
	if err != nil {
		return 0, err
	}

	MathSign := mathString[1]

	return DoMaths(FirstNum, SecondNum, MathSign), err

}
