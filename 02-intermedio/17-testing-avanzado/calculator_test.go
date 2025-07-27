// 🧪 Tests TDD: Calculadora Científica
// Archivo: calculator_test.go
// Ejecutar con: go test -v -run TestScientificCalculator

package main

import (
	"math"
	"math/rand"
	"testing"
)

// ==========================================
// 🧪 TDD TESTS - CALCULADORA CIENTÍFICA
// ==========================================

func TestScientificCalculator_Creation(t *testing.T) {
	calc := NewScientificCalculator()

	if calc == nil {
		t.Fatal("Calculator should not be nil")
	}

	history := calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("New calculator should have empty history, got %d items", len(history))
	}
}

func TestScientificCalculator_Add(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -2.0, -3.0, -5.0},
		{"mixed signs", -5.0, 3.0, -2.0},
		{"with zero", 0.0, 5.0, 5.0},
		{"decimals", 2.5, 3.7, 6.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.Add(test.a, test.b)
			if math.Abs(result-test.expected) > 0.001 {
				t.Errorf("Add(%.2f, %.2f) = %.2f; want %.2f",
					test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestScientificCalculator_Subtract(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive result", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"zero result", 5.0, 5.0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.Subtract(test.a, test.b)
			if math.Abs(result-test.expected) > 0.001 {
				t.Errorf("Subtract(%.2f, %.2f) = %.2f; want %.2f",
					test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestScientificCalculator_Multiply(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 6.0, 7.0, 42.0},
		{"negative numbers", -3.0, -4.0, 12.0},
		{"mixed signs", -5.0, 3.0, -15.0},
		{"multiply by zero", 5.0, 0.0, 0.0},
		{"multiply by one", 7.0, 1.0, 7.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.Multiply(test.a, test.b)
			if math.Abs(result-test.expected) > 0.001 {
				t.Errorf("Multiply(%.2f, %.2f) = %.2f; want %.2f",
					test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestScientificCalculator_Divide(t *testing.T) {
	calc := NewScientificCalculator()

	t.Run("valid divisions", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     float64
			expected float64
		}{
			{"simple division", 15.0, 3.0, 5.0},
			{"decimal result", 10.0, 3.0, 3.333333},
			{"negative numbers", -12.0, 4.0, -3.0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := calc.Divide(test.a, test.b)
				if err != nil {
					t.Errorf("Divide(%.2f, %.2f) returned error: %v", test.a, test.b, err)
				}
				if math.Abs(result-test.expected) > 0.001 {
					t.Errorf("Divide(%.2f, %.2f) = %.6f; want %.6f",
						test.a, test.b, result, test.expected)
				}
			})
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10.0, 0.0)
		if err == nil {
			t.Error("Divide by zero should return an error")
		}
	})
}

func TestScientificCalculator_Power(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name           string
		base, exponent float64
		expected       float64
	}{
		{"2^3", 2.0, 3.0, 8.0},
		{"5^0", 5.0, 0.0, 1.0},
		{"3^2", 3.0, 2.0, 9.0},
		{"negative base", -2.0, 2.0, 4.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.Power(test.base, test.exponent)
			if math.Abs(result-test.expected) > 0.001 {
				t.Errorf("Power(%.2f, %.2f) = %.2f; want %.2f",
					test.base, test.exponent, result, test.expected)
			}
		})
	}
}

func TestScientificCalculator_SquareRoot(t *testing.T) {
	calc := NewScientificCalculator()

	t.Run("valid square roots", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected float64
		}{
			{"perfect square", 16.0, 4.0},
			{"non-perfect square", 2.0, 1.414213},
			{"zero", 0.0, 0.0},
			{"one", 1.0, 1.0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := calc.SquareRoot(test.input)
				if err != nil {
					t.Errorf("SquareRoot(%.2f) returned error: %v", test.input, err)
				}
				if math.Abs(result-test.expected) > 0.001 {
					t.Errorf("SquareRoot(%.2f) = %.6f; want %.6f",
						test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("negative input", func(t *testing.T) {
		_, err := calc.SquareRoot(-4.0)
		if err == nil {
			t.Error("SquareRoot of negative number should return an error")
		}
	})
}

func TestScientificCalculator_TrigonometricFunctions(t *testing.T) {
	calc := NewScientificCalculator()

	t.Run("sin function", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected float64
		}{
			{"sin(0)", 0.0, 0.0},
			{"sin(π/2)", math.Pi / 2, 1.0},
			{"sin(π)", math.Pi, 0.0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := calc.Sin(test.input)
				if math.Abs(result-test.expected) > 0.001 {
					t.Errorf("Sin(%.6f) = %.6f; want %.6f",
						test.input, result, test.expected)
				}
			})
		}
	})

	t.Run("cos function", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected float64
		}{
			{"cos(0)", 0.0, 1.0},
			{"cos(π/2)", math.Pi / 2, 0.0},
			{"cos(π)", math.Pi, -1.0},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := calc.Cos(test.input)
				if math.Abs(result-test.expected) > 0.001 {
					t.Errorf("Cos(%.6f) = %.6f; want %.6f",
						test.input, result, test.expected)
				}
			})
		}
	})
}

func TestScientificCalculator_History(t *testing.T) {
	calc := NewScientificCalculator()

	// Realizar algunas operaciones
	calc.Add(5.0, 3.0)
	calc.Multiply(2.0, 4.0)
	calc.Divide(10.0, 2.0)

	history := calc.GetHistory()

	if len(history) != 3 {
		t.Errorf("History should have 3 entries, got %d", len(history))
	}

	// Verificar que el historial contiene las operaciones
	expectedOperations := []string{"5.00 + 3.00 = 8.00", "2.00 × 4.00 = 8.00", "10.00 ÷ 2.00 = 5.00"}
	for i, expected := range expectedOperations {
		if history[i] != expected {
			t.Errorf("History[%d] = %s; want %s", i, history[i], expected)
		}
	}

	// Limpiar historial
	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("History should be empty after clear, got %d entries", len(history))
	}
}

// ==========================================
// 🎯 PROPERTY-BASED TESTS
// ==========================================

func TestScientificCalculator_AddCommutativeProperty(t *testing.T) {
	calc := NewScientificCalculator()

	// Property: a + b = b + a (commutative)
	for i := 0; i < 100; i++ {
		a := float64(rand.Intn(2000) - 1000) // -1000 to 1000
		b := float64(rand.Intn(2000) - 1000)

		result1 := calc.Add(a, b)
		calc.ClearHistory() // Clear to avoid history pollution
		result2 := calc.Add(b, a)

		if math.Abs(result1-result2) > 0.001 {
			t.Errorf("Commutative property failed: %.2f + %.2f = %.2f, but %.2f + %.2f = %.2f",
				a, b, result1, b, a, result2)
		}
		calc.ClearHistory()
	}
}

func TestScientificCalculator_MultiplyByZeroProperty(t *testing.T) {
	calc := NewScientificCalculator()

	// Property: any number multiplied by zero equals zero
	for i := 0; i < 50; i++ {
		a := float64(rand.Intn(2000) - 1000)

		result := calc.Multiply(a, 0.0)
		if math.Abs(result) > 0.001 {
			t.Errorf("Multiply by zero property failed: %.2f × 0 = %.2f; want 0", a, result)
		}
		calc.ClearHistory()
	}
}

func TestScientificCalculator_SquareRootProperty(t *testing.T) {
	calc := NewScientificCalculator()

	// Property: sqrt(x^2) = |x| for any real x
	for i := 0; i < 50; i++ {
		x := float64(rand.Intn(100) - 50) // -50 to 50

		squared := calc.Multiply(x, x)
		calc.ClearHistory()

		sqrt, err := calc.SquareRoot(squared)
		if err != nil {
			t.Errorf("Unexpected error in square root: %v", err)
			continue
		}

		expected := math.Abs(x)
		if math.Abs(sqrt-expected) > 0.001 {
			t.Errorf("Square root property failed: sqrt((%.2f)^2) = sqrt(%.2f) = %.2f; want %.2f",
				x, squared, sqrt, expected)
		}
		calc.ClearHistory()
	}
}

// ==========================================
// 🧪 BENCHMARK TESTS
// ==========================================

func BenchmarkScientificCalculator_Add(b *testing.B) {
	calc := NewScientificCalculator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

func BenchmarkScientificCalculator_Multiply(b *testing.B) {
	calc := NewScientificCalculator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Multiply(float64(i), float64(i+1))
	}
}

func BenchmarkScientificCalculator_Power(b *testing.B) {
	calc := NewScientificCalculator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Power(2.0, float64(i%10))
	}
}

func BenchmarkScientificCalculator_Sin(b *testing.B) {
	calc := NewScientificCalculator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Sin(float64(i) * math.Pi / 180) // Convert to radians
	}
}
