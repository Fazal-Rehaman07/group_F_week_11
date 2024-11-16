package main

import "testing"

//Running Test-Cases
func TestAdd(t *testing.T) {
	result := Addition(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Addition(2, 3) results in %d, but expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtraction(2, 3)
	expected := -1
	if result != expected {
		t.Errorf("Subtraction(2, 3) results in %d, but expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiplication(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiplication(2, 3) results in %d, but expected %d", result, expected)
	}
}

func TestSquare(t *testing.T) {
	result := Square(5)
	expected := 25
	if result != expected {
		t.Errorf("Square(5) results in %d, but expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
		hasError bool
	}{
		{10, 2, 5, false},
		{10, 0, 0, true},
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)

		// Log the expected error without failing the test
		if tt.hasError && err != nil {
			t.Logf("Expected error for Divide(%d, %d): %v", tt.a, tt.b, err)
		}

		// Check unexpected error state
		if (err != nil) != tt.hasError {
			t.Errorf("Divide(%d, %d) Expected Error = %v; But Got: %v", tt.a, tt.b, tt.hasError, err)
		}

		// Validate the result only if no error is expected
		if !tt.hasError && result != tt.expected {
			t.Errorf("Divide(%d, %d) results in %d, but expected %d", tt.a, tt.b, result, tt.expected)
		}

		// Log success for valid cases
		if !tt.hasError && err == nil {
			t.Logf("Divide(%d, %d) PASS", tt.a, tt.b)
		}
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		a        float64
		expected float64
		hasError bool
	}{
		{25, 5, false},
		{-9, 0, true},
	}

	for _, tt := range tests {
		result, err := SquareRoot(tt.a)

		// Log the expected error
		if tt.hasError && err != nil {
			t.Logf("Expected error for SquareRoot(%f): %v", tt.a, err)
		}

		// Check unexpected error state
		if (err != nil) != tt.hasError {
			t.Errorf("SquareRoot(%f) Expected Error: %v; But Got: %v", tt.a, tt.hasError, err)
		}

		// Validate the result if no error was expected
		if !tt.hasError && result != tt.expected {
			t.Errorf("SquareRoot(%f) results in %f, but expected %f", tt.a, result, tt.expected)
		}

		// Log success for valid cases
		if !tt.hasError && err == nil {
			t.Logf("SquareRoot(%f) PASS", tt.a)
		}
	}
}

//Running Benchmarks
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addition(100, 50)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtraction(100, 50)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplication(24, 52)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(25)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(25, 5)
	}
}

func BenchmarkSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot(625)
	}
}
