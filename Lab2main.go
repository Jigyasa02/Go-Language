package main

import "fmt"

// Function to modify original variable using pointer
func modifyValue(ptr *int) {
	*ptr = 100
}

func main() {

	// Part 1: Pointer Basics
	var num int = 50

	fmt.Println("----- POINTER BASICS -----")

	fmt.Println("Original Value:", num)
	fmt.Println("Memory Address:", &num)
	fmt.Println("Value using Pointer:", *(&num))

	// Part 2: Modify variable using pointer function
	fmt.Println("\n----- POINTER FUNCTION -----")

	fmt.Println("Before Modification:", num)

	modifyValue(&num)

	fmt.Println("After Modification:", num)

	// Part 3: Struct Pointer using new()
	fmt.Println("\n----- STRUCT POINTER -----")

	type Student struct {
		Name  string
		Marks int
	}

	student := new(Student)

	student.Name = "Jigyasa Singh"
	student.Marks = 95

	fmt.Println("Student Name:", student.Name)
	fmt.Println("Student Marks:", student.Marks)

	// Modify struct field
	student.Marks = 98

	fmt.Println("Updated Student Marks:", student.Marks)
}
