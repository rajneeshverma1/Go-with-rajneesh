package app

import "fmt"

// Day2Basics demonstrates Day 2 Go fundamentals
func Day2Basics() string {
	var result string

	// 1. Variables and Data Types
	result += "=== Variables and Data Types ===\n"
	var name string = "Rajneesh"
	age := 21
	isLearning := true
	gpa := 9.5

	result += fmt.Sprintf("Name: %s\n", name)
	result += fmt.Sprintf("Age: %d\n", age)
	result += fmt.Sprintf("Is Learning: %t\n", isLearning)
	result += fmt.Sprintf("GPA: %.1f\n\n", gpa)

	// 2. If/Else Statements
	result += "=== If/Else Statements ===\n"
	if age >= 18 {
		result += "You are an adult\n"
	} else {
		result += "You are a minor\n"
	}

	if gpa >= 9.0 {
		result += "Grade: A+\n"
	} else if gpa >= 8.0 {
		result += "Grade: A\n"
	} else {
		result += "Grade: B\n"
	}
	result += "\n"

	// 3. Switch Statement
	result += "=== Switch Statement ===\n"
	day := "Monday"
	switch day {
	case "Monday":
		result += "Start of the week\n"
	case "Friday":
		result += "Almost weekend\n"
	case "Saturday", "Sunday":
		result += "Weekend!\n"
	default:
		result += "Midweek day\n"
	}
	result += "\n"

	// 4. For Loops
	result += "=== For Loops ===\n"

	// Basic for loop
	result += "Counting 1 to 5: "
	for i := 1; i <= 5; i++ {
		if i < 5 {
			result += fmt.Sprintf("%d, ", i)
		} else {
			result += fmt.Sprintf("%d\n", i)
		}
	}

	// While-style for loop
	result += "\nWhile-style loop: "
	count := 1
	for count <= 3 {
		result += fmt.Sprintf("%d ", count)
		count++
	}
	result += "\n\n"

	// 5. Functions
	result += "=== Functions ===\n"
	sum := add(10, 20)
	result += fmt.Sprintf("Sum of 10 and 20: %d\n", sum)

	product := multiply(5, 6)
	result += fmt.Sprintf("Product of 5 and 6: %d\n", product)

	// Function with multiple return values
	quotient, remainder := divide(17, 5)
	result += fmt.Sprintf("17 / 5 = %d with remainder %d\n\n", quotient, remainder)

	// 6. Arrays and Slices
	result += "=== Arrays and Slices ===\n"

	// Array (fixed size)
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	result += fmt.Sprintf("Array: %v\n", numbers)

	// Slice (dynamic size)
	fruits := []string{"Apple", "Banana", "Orange"}
	result += fmt.Sprintf("Slice: %v\n", fruits)

	// Append to slice
	fruits = append(fruits, "Mango")
	result += fmt.Sprintf("After append: %v\n", fruits)

	// Slice operations
	result += fmt.Sprintf("First fruit: %s\n", fruits[0])
	result += fmt.Sprintf("Slice length: %d\n", len(fruits))
	result += fmt.Sprintf("Slice [1:3]: %v\n\n", fruits[1:3])

	// 7. Maps
	result += "=== Maps ===\n"
	student := map[string]string{
		"name":    "Rajneesh",
		"course":  "Go Programming",
		"status":  "Learning",
	}

	result += fmt.Sprintf("Student: %v\n", student)
	result += fmt.Sprintf("Name: %s\n", student["name"])
	result += fmt.Sprintf("Course: %s\n\n", student["course"])

	// 8. Range (iterating over slices and maps)
	result += "=== Range ===\n"
	result += "Iterating over fruits:\n"
	for index, fruit := range fruits {
		result += fmt.Sprintf("  %d: %s\n", index, fruit)
	}

	result += "\nIterating over map:\n"
	for key, value := range student {
		result += fmt.Sprintf("  %s: %s\n", key, value)
	}

	return result
}

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
