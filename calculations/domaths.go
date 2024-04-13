package calculations

func DoMaths(num1, num2 float64, operator string) float64 {
	var Output float64
	switch {
	case operator == "+":
		Output = num1 + num2
	case operator == "-":
		Output = num1 - num2
	case operator == "*":
		Output = num1 * num2
	case operator == "/":
		Output = num1 / num2
	case operator == "**":
		Output = num1 / num2
	}
	return Output
}
