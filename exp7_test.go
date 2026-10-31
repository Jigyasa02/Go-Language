package main

import (
	"strings"
	"testing"
)

// Test division function
func TestDivide(t *testing.T) {
	// Valid input
	result, err := divide(20, 5)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 4 {
		t.Errorf("Expected 4, got %f", result)
	}

	// Error case: division by zero
	_, err = divide(10, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}

	if !strings.Contains(err.Error(), "division by zero") {
		t.Errorf("Unexpected error message: %v", err)
	}
}

// Test square root function
func TestSquareRoot(t *testing.T) {
	// Valid input
	result, err := squareRoot(64)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 8 {
		t.Errorf("Expected 8, got %f", result)
	}

	// Error case: negative number
	_, err = squareRoot(-25)

	if err == nil {
		t.Error("Expected negative square root error")
	}

	if !strings.Contains(err.Error(), "negative number") {
		t.Errorf("Unexpected error message: %v", err)
	}
}

// Test input validation function
func TestValidateInput(t *testing.T) {
	// Valid input
	err := validateInput("Hello Go")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Error case: empty input
	err = validateInput("")

	if err == nil {
		t.Error("Expected empty input error")
	}

	if !strings.Contains(err.Error(), "empty") {
		t.Errorf("Unexpected error message: %v", err)
	}
}
