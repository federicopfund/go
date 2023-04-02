package calculator


var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"


func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) float64 {
	if b == 0 {
		return float64(0)
	}
	return float64(a) / float64(b)
}


func internalSum(number int) int {
    return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
    return number1 + number2
}