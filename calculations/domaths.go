package calculations

func DoMaths(num1, num2 int, operator string) int {
	var Output int
	switch {
	case operator == "+":
		Output = num1 + num2
	case operator == "-":
		Output = num1 - num2
	case operator == "x":
		Output = num1 * num2
	case operator == "/":
		Output = num1 / num2
	}
	return Output
}
