// 🧪 Ejemplo de Implementación: Tests para BookValidator
// ===================================================
// Este archivo muestra cómo implementar los tests del Ejercicio 1
// Usa este como guía para completar todos los ejercicios

package main

import (
	"testing"
)

// =============================================================================
// 📝 EJERCICIO 1: Tests para BookValidator
// =============================================================================

func TestBookValidator_ValidateBook(t *testing.T) {
	validator := &BookValidator{}

	tests := []struct {
		name           string
		book           *Book
		expectedErrors []string
	}{
		{
			name: "valid book",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "978-0134190440", // Valid ISBN-13
				Year:     2020,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{},
		},
		{
			name: "empty title",
			book: &Book{
				Title:    "",
				Author:   "John Doe",
				ISBN:     "978-0134190440",
				Year:     2020,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"title cannot be empty"},
		},
		{
			name: "empty author",
			book: &Book{
				Title:    "Go Programming",
				Author:   "",
				ISBN:     "978-0134190440",
				Year:     2020,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"author cannot be empty"},
		},
		{
			name: "invalid ISBN",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "invalid-isbn",
				Year:     2020,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"invalid ISBN format"},
		},
		{
			name: "invalid year - too old",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "978-0134190440",
				Year:     999,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"invalid year"},
		},
		{
			name: "invalid year - future",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "978-0134190440",
				Year:     2030,
				Pages:    350,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"invalid year"},
		},
		{
			name: "invalid pages",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "978-0134190440",
				Year:     2020,
				Pages:    -10,
				Genre:    "Technology",
				Language: "English",
			},
			expectedErrors: []string{"pages must be positive"},
		},
		{
			name: "empty genre",
			book: &Book{
				Title:    "Go Programming",
				Author:   "John Doe",
				ISBN:     "978-0134190440",
				Year:     2020,
				Pages:    350,
				Genre:    "",
				Language: "English",
			},
			expectedErrors: []string{"genre cannot be empty"},
		},
		{
			name: "multiple errors",
			book: &Book{
				Title:    "",
				Author:   "",
				ISBN:     "invalid",
				Year:     999,
				Pages:    -1,
				Genre:    "",
				Language: "English",
			},
			expectedErrors: []string{
				"title cannot be empty",
				"author cannot be empty",
				"invalid ISBN format",
				"invalid year",
				"pages must be positive",
				"genre cannot be empty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := validator.ValidateBook(tt.book)

			// Verificar que tenemos el número correcto de errores
			if len(errors) != len(tt.expectedErrors) {
				t.Errorf("Expected %d errors, got %d: %v",
					len(tt.expectedErrors), len(errors), errors)
				return
			}

			// Verificar que cada error esperado está presente
			for _, expectedError := range tt.expectedErrors {
				found := false
				for _, actualError := range errors {
					if actualError == expectedError {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected error '%s' not found in: %v",
						expectedError, errors)
				}
			}
		})
	}
}

func TestBookValidator_ValidateISBN(t *testing.T) {
	validator := &BookValidator{}

	t.Run("ISBN-10 valid", func(t *testing.T) {
		tests := []struct {
			name string
			isbn string
			want bool
		}{
			{"standard ISBN-10", "0134190440", true},
			{"ISBN-10 with X", "013419044X", true},
			{"ISBN-10 with dashes", "0-13-419044-0", true},
			{"ISBN-10 with spaces", "0 13 419044 0", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := validator.ValidateISBN(tt.isbn)
				if result != tt.want {
					t.Errorf("ValidateISBN(%q) = %v; want %v",
						tt.isbn, result, tt.want)
				}
			})
		}
	})

	t.Run("ISBN-13 valid", func(t *testing.T) {
		tests := []struct {
			name string
			isbn string
			want bool
		}{
			{"standard ISBN-13", "9780134190440", true},
			{"ISBN-13 with dashes", "978-0-13-419044-0", true},
			{"ISBN-13 with spaces", "978 0 13 419044 0", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := validator.ValidateISBN(tt.isbn)
				if result != tt.want {
					t.Errorf("ValidateISBN(%q) = %v; want %v",
						tt.isbn, result, tt.want)
				}
			})
		}
	})

	t.Run("Invalid ISBNs", func(t *testing.T) {
		tests := []struct {
			name string
			isbn string
			want bool
		}{
			{"too short", "123456789", false},
			{"too long", "12345678901234", false},
			{"invalid characters", "abcdefghij", false},
			{"wrong checksum ISBN-10", "0134190441", false},
			{"wrong checksum ISBN-13", "9780134190441", false},
			{"empty string", "", false},
			{"only dashes", "----------", false},
			{"mixed length with dashes", "978-013419044", false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := validator.ValidateISBN(tt.isbn)
				if result != tt.want {
					t.Errorf("ValidateISBN(%q) = %v; want %v",
						tt.isbn, result, tt.want)
				}
			})
		}
	})
}

// =============================================================================
// 📝 EJERCICIO BONUS: Benchmark para ValidateISBN
// =============================================================================

func BenchmarkBookValidator_ValidateISBN10(b *testing.B) {
	validator := &BookValidator{}
	isbn := "0134190440"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.ValidateISBN(isbn)
	}
}

func BenchmarkBookValidator_ValidateISBN13(b *testing.B) {
	validator := &BookValidator{}
	isbn := "9780134190440"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.ValidateISBN(isbn)
	}
}

func BenchmarkBookValidator_ValidateBook(b *testing.B) {
	validator := &BookValidator{}
	book := &Book{
		Title:    "Go Programming",
		Author:   "John Doe",
		ISBN:     "978-0134190440",
		Year:     2020,
		Pages:    350,
		Genre:    "Technology",
		Language: "English",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.ValidateBook(book)
	}
}

// =============================================================================
// 📝 EJERCICIO AVANZADO: Property-Based Testing
// =============================================================================

func TestBookValidator_ISBN_Properties(t *testing.T) {
	validator := &BookValidator{}

	t.Run("ISBN length property", func(t *testing.T) {
		// Propiedad: Solo ISBNs de 10 o 13 dígitos (sin contar separadores) son válidos
		testCases := []string{
			"123456789",      // 9 dígitos - inválido
			"1234567890",     // 10 dígitos - podría ser válido
			"12345678901",    // 11 dígitos - inválido
			"123456789012",   // 12 dígitos - inválido
			"1234567890123",  // 13 dígitos - podría ser válido
			"12345678901234", // 14 dígitos - inválido
		}

		for _, isbn := range testCases {
			result := validator.ValidateISBN(isbn)
			// Solo longitudes 10 y 13 pueden ser válidas
			if len(isbn) != 10 && len(isbn) != 13 && result {
				t.Errorf("ISBN with length %d should never be valid: %s",
					len(isbn), isbn)
			}
		}
	})

	t.Run("ISBN format property", func(t *testing.T) {
		// Propiedad: Separadores no deben afectar validación
		validISBN10 := "0134190440"

		variations := []string{
			"0134190440",
			"0-134190440",
			"01-34190440",
			"013-4190440",
			"0134-190440",
			"01341-90440",
			"013419-0440",
			"0134190-440",
			"01341904-40",
			"013419044-0",
			"0-1-3-4-1-9-0-4-4-0",
			"0 134190440",
			"0 1 3 4 1 9 0 4 4 0",
		}

		expectedResult := validator.ValidateISBN(validISBN10)

		for _, variation := range variations {
			result := validator.ValidateISBN(variation)
			if result != expectedResult {
				t.Errorf("ISBN variations should have same validity: %s", variation)
			}
		}
	})
}

// =============================================================================
// 📝 HELPER FUNCTIONS PARA TESTING
// =============================================================================

// createValidBook helper para crear libro válido en tests
func createValidBook() *Book {
	return &Book{
		Title:    "Go Programming",
		Author:   "John Doe",
		ISBN:     "978-0134190440",
		Year:     2020,
		Pages:    350,
		Genre:    "Technology",
		Language: "English",
	}
}

// assertNoValidationErrors helper para verificar que no hay errores
func assertNoValidationErrors(t *testing.T, errors []string) {
	t.Helper()
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors, got: %v", errors)
	}
}

// assertHasValidationError helper para verificar error específico
func assertHasValidationError(t *testing.T, errors []string, expectedError string) {
	t.Helper()
	for _, err := range errors {
		if err == expectedError {
			return
		}
	}
	t.Errorf("Expected validation error '%s' not found in: %v",
		expectedError, errors)
}

// =============================================================================
// 📝 EJEMPLO DE USAGE EN TESTS REALES
// =============================================================================

func ExampleBookValidator_ValidateBook() {
	validator := &BookValidator{}

	// Libro válido
	validBook := &Book{
		Title:    "Go Programming",
		Author:   "John Doe",
		ISBN:     "978-0134190440",
		Year:     2020,
		Pages:    350,
		Genre:    "Technology",
		Language: "English",
	}

	errors := validator.ValidateBook(validBook)
	if len(errors) == 0 {
		println("Book is valid!")
	}

	// Libro inválido
	invalidBook := &Book{
		Title:  "", // Error: título vacío
		Author: "John Doe",
		ISBN:   "invalid-isbn", // Error: ISBN inválido
		Year:   2020,
		Pages:  350,
		Genre:  "Technology",
	}

	errors = validator.ValidateBook(invalidBook)
	for _, err := range errors {
		println("Validation error:", err)
	}

	// Output:
	// Book is valid!
	// Validation error: title cannot be empty
	// Validation error: invalid ISBN format
}
