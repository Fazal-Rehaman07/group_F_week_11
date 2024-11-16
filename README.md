### `README.md`

# Calculator in Go

## **Project Overview**
This project implements a simple Calculator written in GoLang. It provides functions for addition, subtraction, multiplication, division, and squaring a number. The test_calculator.go file includes unit tests and benchmarks to ensure correctness and measure performance.

---

## **Instructions**

### **How to Run the Program**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Fazal-Rehaman07/group_F_week_11.git
   cd group_F_week_11.git
   ```
2. **Run the Program**:
   Use the `go run` command to execute the main file:
   ```bash
   go run calculator.go
   ```
   This will run the Simple Calculator with sample inputs. You can change the inputs inside main function in the calculator.go file

---

### **How to Execute the Tests and Benchmarks**

#### **Run Unit Tests**
To run the unit tests, execute the following command:
```bash
go test -v
```

#### **Run Benchmarks**
Comment line 44 i.e the Divide by zero test case {10, 0, 0} to prevent Failure.  

To run Benchmark tests, run:
```bash
go test -bench=.
```

---

### **Expected Output of the Tests**

#### **Unit Test Output**
Sample output for test file:
```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
=== RUN   TestSquare
--- PASS: TestSquare (0.00s)
=== RUN   TestDivide
    calculator_test.go:55: Divide(10, 2) PASS
    calculator_test.go:50: Divide(10, 0) FAILED: division by zero
--- FAIL: TestDivide (0.00s)
FAIL
exit status 1

```

#### **Benchmark Output**
Sample benchmark output:
```
goos: windows
goarch: amd64
pkg: github.com/Fazal-Rehaman07/group_F_week_11
cpu: 13th Gen Intel(R) Core(TM) i5-13500H
BenchmarkAdd-16                 1000000000               0.1830 ns/op
BenchmarkSubtract-16            1000000000               0.1838 ns/op
BenchmarkMultiply-16            1000000000               0.1850 ns/op
BenchmarkSquare-16              1000000000               0.1835 ns/op
BenchmarkDivide-16              1000000000               0.1937 ns/op
PASS
ok      github.com/Fazal-Rehaman07/group_F_week_11      1.434s

```

---

## **Explanation of Test and Benchmark Design**

### **Test Design**
- **Unit Tests**:
  - Validate the correctness of each function by comparing their output with expected values.
  - Include edge cases, such as division by zero, to ensure error handling works as intended.

- **Structured Test for Division**:
  - Uses a struct to test multiple cases (e.g., valid division and division by zero).
  - Confirms error handling for invalid inputs.

### **Benchmark Design**
- Benchmarks evaluate the performance of the arithmetic functions under repetitive execution.
- Each benchmark function iterates the corresponding operation `b.N` times, where `b.N` is dynamically determined by the testing framework to provide stable performance results.

