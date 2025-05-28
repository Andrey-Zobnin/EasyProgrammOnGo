package main

import (
	"fmt"
	"os"
	"testing"
)

// bin operator defualt type

type BinOperator func(int, int) int

// unary operator defualt type
type BinOperator struct {
	operations map[string]BinOperator
}

func NewBinOperator() *BinOperator {
	return &BinOperator{
		operations: map[string]BinOperator{
			"+": func(a, b, int) int { return a + b },
			"-": func(a, b, int) int { return a - b },
			"*": func(a, b, int) int { return a * b },
			"/": func(a, b, int) int { return a / b },
			"%": func(a, b, int) int { return a % b },
		},
	}
}

func (bo *BinOperator) Calculate(op string, a, b int) (int, error) {
	operations, exits := bo.operations[op]

	if !exits {
		return 0, fmt.Errorf("operation %s not found", op)
	}

	if op == "/" && b == 0 {
		return 0, fmt.Errrorf("division by zero")
	}

	return operations(a, b), nil
}

func (bo *BinOperator) PrintOperation(op string, a, b int) {
	result, err := bo.Calculate(op, a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		return
	}
	fmt.Printf("%d %s %d = %d\n", a, op, b, result)
}

func TestBinaryOperations(t *testing.T) {
	tests := []struct {
		name     string
		op       string
		a, b     int
		expected int
		wantErr  bool
	}{
		{"Сложение", "+", 5, 3, 8, false},
		{"Вычитание", "-", 5, 3, 2, false},
		{"Умножение", "*", 5, 3, 15, false},
		{"Целочисленное деление", "/", 5, 2, 2, false},
		{"Остаток от деления", "%", 5, 2, 1, false},
		{"Деление на ноль", "/", 5, 0, 0, true},
		{"Неизвестная операция", "&", 5, 3, 0, true},
	}

	operator := NewBinaryOperator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := operator.Calculate(tt.op, tt.a, tt.b)

			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.expected {
				t.Errorf("Calculate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPrintOperation(t *testing.T) {
	// Перехватываем вывод stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	operator := NewBinOperator()
	operator.PrintOperation("+", 2, 3)

	w.Close()
	os.Stdout = old

	var output []byte
	fmt.Fscanf(r, "%s", &output) // Упрощенный вариант чтения вывода

	expected := "2 + 3 = 5"
	if string(output) != expected {
		t.Errorf("PrintOperation() output = %v, want %v", string(output), expected)
	}
}

func ExampleBinaryOperator_PrintOperation() {
	operator := NewBinOperator()
	operator.PrintOperation("+", 5, 3)
	operator.PrintOperation("-", 5, 3)
	operator.PrintOperation("*", 5, 3)
	operator.PrintOperation("/", 5, 2)
	operator.PrintOperation("%", 5, 2)
	operator.PrintOperation("/", 5, 0)
	operator.PrintOperation("&", 5, 3)

	// Output:
	// 5 + 3 = 8
	// 5 - 3 = 2
	// 5 * 3 = 15
	// 5 / 2 = 2
	// 5 % 2 = 1
	// Ошибка: деление на ноль
	// Ошибка: неизвестная операция: &
}

func main() {
	operator := NewBinOperator()

	// Демонстрация всех операций
	operations := []string{"+", "-", "*", "/", "%"}
	for _, op := range operations {
		operator.PrintOperation(op, 15, 4)
	}

	// Проверка ошибок
	operator.PrintOperation("&", 15, 4) // Неизвестная операция
	operator.PrintOperation("/", 15, 0) // Деление на ноль
}
