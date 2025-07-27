package main

import (
	"math"
	"testing"
)

// ========================================
// Tests para Calculadora Científica (TDD)
// ========================================

func TestScientificCalculator_Add(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2.5, 3.5, 6.0},
		{"negative numbers", -2.0, -3.0, -5.0},
		{"mixed numbers", -2.5, 5.0, 2.5},
		{"with zero", 0, 5.0, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%.2f, %.2f) = %.2f; want %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestScientificCalculator_Divide(t *testing.T) {
	calc := NewScientificCalculator()

	t.Run("normal division", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		if err != nil {
			t.Errorf("Divide(10, 2) returned error: %v", err)
		}
		if result != 5.0 {
			t.Errorf("Divide(10, 2) = %.2f; want 5.00", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Error("Divide(10, 0) should return error")
		}
	})
}

func TestScientificCalculator_SquareRoot(t *testing.T) {
	calc := NewScientificCalculator()

	t.Run("positive number", func(t *testing.T) {
		result, err := calc.SquareRoot(9)
		if err != nil {
			t.Errorf("SquareRoot(9) returned error: %v", err)
		}
		if result != 3.0 {
			t.Errorf("SquareRoot(9) = %.2f; want 3.00", result)
		}
	})

	t.Run("negative number", func(t *testing.T) {
		_, err := calc.SquareRoot(-4)
		if err == nil {
			t.Error("SquareRoot(-4) should return error")
		}
	})
}

func TestScientificCalculator_History(t *testing.T) {
	calc := NewScientificCalculator()

	if len(calc.GetHistory()) != 0 {
		t.Error("New calculator should have empty history")
	}

	calc.Add(2, 3)
	calc.Multiply(4, 5)

	history := calc.GetHistory()
	if len(history) != 2 {
		t.Errorf("History should have 2 entries, got %d", len(history))
	}

	calc.ClearHistory()
	if len(calc.GetHistory()) != 0 {
		t.Error("History should be empty after clearing")
	}
}

func TestScientificCalculator_TrigonometricFunctions(t *testing.T) {
	calc := NewScientificCalculator()

	tests := []struct {
		name     string
		function func(float64) float64
		input    float64
		expected float64
		delta    float64
	}{
		{"sin(0)", calc.Sin, 0, 0, 0.001},
		{"sin(π/2)", calc.Sin, math.Pi / 2, 1, 0.001},
		{"cos(0)", calc.Cos, 0, 1, 0.001},
		{"cos(π)", calc.Cos, math.Pi, -1, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input)
			if math.Abs(result-tt.expected) > tt.delta {
				t.Errorf("%s = %.6f; want %.6f (±%.3f)", tt.name, result, tt.expected, tt.delta)
			}
		})
	}
}

// ========================================
// Benchmarks para Calculadora
// ========================================

func BenchmarkCalculator_Add(b *testing.B) {
	calc := NewScientificCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

func BenchmarkCalculator_Power(b *testing.B) {
	calc := NewScientificCalculator()
	for i := 0; i < b.N; i++ {
		calc.Power(2.0, float64(i%10))
	}
}

func BenchmarkCalculator_SquareRoot(b *testing.B) {
	calc := NewScientificCalculator()
	for i := 0; i < b.N; i++ {
		calc.SquareRoot(float64(i + 1))
	}
}
