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
	}{
		{10, 2, 5},
		//{10, 0, 0},
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)
		if err != nil {
			t.Fatalf("Divide(%d, %d) FAILED: %v", tt.a, tt.b, err)
		}
		if result != tt.expected {
			t.Fatalf("Divide(%d, %d) results in %d, but expected %d", tt.a, tt.b, result, tt.expected)
		}
		t.Logf("Divide(%d, %d) PASS ", tt.a, tt.b)
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
