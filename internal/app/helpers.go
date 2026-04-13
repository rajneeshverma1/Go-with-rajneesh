package app

// add returns the sum of two integers
func add(a int, b int) int {
	return a + b
}

// multiply returns the product of two integers
func multiply(a int, b int) int {
	return a * b
}

// divide returns quotient and remainder
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
