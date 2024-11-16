package main

import (
	"errors"
	"fmt"
)

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	return a * b
}

// Divide returns the quotient of a divided by b.
// Returns an error if b is 0.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Square returns the square of a.
func Square(a int) int {
	return a * a
}

func main() {
	fmt.Println("Addition: ", Addition(2024, 1))
	fmt.Println("Subtraction: ", Subtraction(2, 2))
	fmt.Println("Multiplication: ", Multiplication(2, 8))
	quotient, err := Divide(10, 0)
	if err != nil {
		fmt.Printf("Error Dividing: %v\n", err)
	} else {
		fmt.Printf("Division: %d\n", quotient)
	}
	fmt.Println("Square: ", Square(5))
}
