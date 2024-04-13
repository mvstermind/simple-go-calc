package calculations

func DoMaths(num1, num2 float64, operator string) float64 {
	var output float64
	switch {
	case operator == "+":
		output = num1 + num2
	case operator == "-":
		output = num1 - num2
	case operator == "*":
		output = num1 * num2
	case operator == "/":
		output = num1 / num2
	case operator == "**":
		output = num1 / num2
	}
	return output
}
