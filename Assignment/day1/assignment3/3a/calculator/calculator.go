package calculator

func Square(n int) int {
	return n * n
}

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) (int, int) {
	return a / b, a % b
}
