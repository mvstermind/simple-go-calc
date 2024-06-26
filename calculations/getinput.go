package calculations

import (
	"bufio"
	"fmt"
	"os"
)

func GetDigits() (float64, error) {
	fmt.Printf("Type a thing to calculate: ")
	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Extract numbers and math operator from input then calculate output
	return DoMaths(ExctractInput(input)), err

}
