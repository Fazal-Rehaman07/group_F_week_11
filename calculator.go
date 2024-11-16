package main

import (
	"errors"
	"fmt"
	"math"
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

// Divide returns the quotient of a / b.
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

// SquareRoot returns the square root of a number
// Returns an error if the input is negative.
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return math.Sqrt(a), nil
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
	sqrt, err := SquareRoot(25)
	if err != nil {
		fmt.Printf("Error Calculating Square Root: %v\n", err)
	} else {
		fmt.Printf("Square Root: %.2f\n", sqrt)
	}

}
