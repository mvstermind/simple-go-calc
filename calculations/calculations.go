package calculations

import "fmt"

func GetDigits(FirstNum, SecondNum int) (int, int) {
	fmt.Printf("Type first number: ")
	fmt.Scan(&FirstNum)
	fmt.Printf("Type second number: ")
	fmt.Scan(&SecondNum)
	return FirstNum, SecondNum
}
