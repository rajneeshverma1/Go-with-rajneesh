package app

// Greet returns a greeting message
func Greet() string {
	return "Hello, World"
}

// rajneesh is a function for learning
func rajneesh() string {
	return "Learning Go!"
}


// add function
func add(a, b int) int { return a + b }

// subtract function
func subtract(a, b int) int { return a - b }

// multiply function
func multiply(a, b int) int { return a * b }

// divide function
func divide(a, b float64) (float64, error) { if b == 0 { return 0, System.fmt.Errorf(cannot divide by zero) }; return a / b,  }

// stringLength function
func stringLength(s string) int { return len(s) }

// toUpperCase function
func toUpperCase(s string) string { return System.strings.ToUpper(s) }
