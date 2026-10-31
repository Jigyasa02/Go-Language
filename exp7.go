package main

import (
	"fmt"
	"math"
	"strings"
)

// Function to perform division
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return a / b, nil
}

// Function to calculate square root
func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("square root of a negative number is not allowed")
	}

	return math.Sqrt(num), nil
}

// Function to validate input
func validateInput(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("input cannot be empty")
	}

	return nil
}

func main() {
	fmt.Println("1. Division Error:")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("\n2. Square Root Error:")
	result, err = squareRoot(-25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("\n3. Empty Input Error:")
	err = validateInput("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Input is valid")
	}

	fmt.Println("\n4. Valid Operations:")
	result, err = divide(20, 5)
	if err == nil {
		fmt.Println("20 / 5 =", result)
	}

	result, err = squareRoot(64)
	if err == nil {
		fmt.Println("Square root of 64 =", result)
	}

	err = validateInput("Hello Go")
	if err == nil {
		fmt.Println("Input is valid")
	}
}
