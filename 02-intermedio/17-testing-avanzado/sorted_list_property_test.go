package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// ========================================
// Property Tests para Lista Ordenada
// ========================================

func TestSortedList_Properties(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Property: Lista siempre está ordenada después de insertar
	t.Run("always sorted after insert", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			list := NewSortedList()

			// Insertar números aleatorios
			for j := 0; j < 20; j++ {
				value := rand.Intn(100)
				list.Insert(value)

				// Verificar que la lista esté ordenada
				slice := list.ToSlice()
				if !isSorted(slice) {
					t.Errorf("Lista no está ordenada después de insertar %d: %v", value, slice)
				}
			}
		}
	})

	// Property: Tamaño aumenta correctamente al insertar elementos únicos
	t.Run("size increases correctly", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			list := NewSortedList()

			for j := 0; j < 10; j++ {
				initialSize := list.Size()
				value := j * 10 // Valores únicos
				list.Insert(value)
				newSize := list.Size()

				if newSize != initialSize+1 {
					t.Errorf("Tamaño debería aumentar de %d a %d, pero es %d", initialSize, initialSize+1, newSize)
				}
			}
		}
	})

	// Property: Contains funciona correctamente
	t.Run("contains works correctly", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			list := NewSortedList()
			values := make(map[int]bool)

			// Insertar valores aleatorios
			for j := 0; j < 15; j++ {
				value := rand.Intn(50)
				list.Insert(value)
				values[value] = true
			}

			// Verificar que Contains funciona para valores insertados
			for value := range values {
				if !list.Contains(value) {
					t.Errorf("Lista debería contener %d", value)
				}
			}

			// Verificar que Contains retorna false para valores no insertados
			for k := 0; k < 10; k++ {
				value := rand.Intn(50) + 100 // Valores fuera del rango insertado
				if list.Contains(value) {
					// Solo falla si realmente no está en la lista
					if !values[value] {
						t.Errorf("Lista no debería contener %d", value)
					}
				}
			}
		}
	})

	// Property: Remove funciona correctamente
	t.Run("remove works correctly", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			list := NewSortedList()
			values := []int{}

			// Insertar valores conocidos
			for j := 0; j < 10; j++ {
				value := j * 2
				list.Insert(value)
				values = append(values, value)
			}

			// Remover algunos valores
			for j := 0; j < 5; j++ {
				value := values[j]
				initialSize := list.Size()
				removed := list.Remove(value)
				newSize := list.Size()

				if !removed {
					t.Errorf("Remove debería retornar true al remover %d", value)
				}

				if newSize != initialSize-1 {
					t.Errorf("Tamaño debería disminuir de %d a %d, pero es %d", initialSize, initialSize-1, newSize)
				}

				if list.Contains(value) {
					t.Errorf("Lista no debería contener %d después de remover", value)
				}

				// Verificar que la lista sigue ordenada
				slice := list.ToSlice()
				if !isSorted(slice) {
					t.Errorf("Lista no está ordenada después de remover %d: %v", value, slice)
				}
			}
		}
	})

	// Property: Min y Max son correctos
	t.Run("min and max are correct", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			list := NewSortedList()

			// Lista vacía
			_, hasMin := list.Min()
			_, hasMax := list.Max()
			if hasMin || hasMax {
				t.Error("Lista vacía no debería tener min o max")
			}

			// Insertar valores
			values := []int{}
			for j := 0; j < 10; j++ {
				value := rand.Intn(100)
				list.Insert(value)
				values = append(values, value)

				min, hasMin := list.Min()
				max, hasMax := list.Max()

				if !hasMin || !hasMax {
					t.Error("Lista no vacía debería tener min y max")
					continue
				}

				// Calcular min y max esperados
				sort.Ints(values)
				expectedMin := values[0]
				expectedMax := values[len(values)-1]

				if min != expectedMin {
					t.Errorf("Min debería ser %d, pero es %d", expectedMin, min)
				}

				if max != expectedMax {
					t.Errorf("Max debería ser %d, pero es %d", expectedMax, max)
				}
			}
		}
	})
}

// Property: Invariante - lista siempre ordenada
func TestSortedList_AlwaysSorted_Invariant(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		list := NewSortedList()

		// Realizar operaciones aleatorias
		for j := 0; j < 50; j++ {
			operation := rand.Intn(3)
			value := rand.Intn(100)

			switch operation {
			case 0: // Insert
				list.Insert(value)
			case 1: // Remove
				list.Remove(value)
			case 2: // Insert duplicado
				list.Insert(value)
				list.Insert(value)
			}

			// Verificar invariante: lista siempre ordenada
			slice := list.ToSlice()
			if !isSorted(slice) {
				t.Errorf("Invariante violado: lista no está ordenada: %v", slice)
			}
		}
	}
}

// Property: Comportamiento idempotente
func TestSortedList_IdempotentOperations(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 50; i++ {
		list := NewSortedList()

		// Insertar algunos valores
		for j := 0; j < 10; j++ {
			value := rand.Intn(50)
			list.Insert(value)
		}

		// Guardar estado
		initialSlice := list.ToSlice()
		initialSize := list.Size()

		// Operaciones idempotentes
		for j := 0; j < 5; j++ {
			// Contains no debería cambiar el estado
			for _, value := range initialSlice {
				list.Contains(value)
			}

			// ToSlice no debería cambiar el estado
			list.ToSlice()

			// Min/Max no deberían cambiar el estado
			list.Min()
			list.Max()

			// Size no debería cambiar el estado
			list.Size()
		}

		// Verificar que el estado no cambió
		finalSlice := list.ToSlice()
		finalSize := list.Size()

		if finalSize != initialSize {
			t.Errorf("Tamaño cambió después de operaciones idempotentes: %d -> %d", initialSize, finalSize)
		}

		if !slicesEqual(initialSlice, finalSlice) {
			t.Errorf("Lista cambió después de operaciones idempotentes: %v -> %v", initialSlice, finalSlice)
		}
	}
}

// Property: Insertar y remover el mismo elemento
func TestSortedList_InsertRemoveCycle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 50; i++ {
		list := NewSortedList()

		// Estado inicial
		for j := 0; j < 5; j++ {
			list.Insert(j * 10)
		}

		initialSlice := list.ToSlice()
		initialSize := list.Size()

		// Ciclo insert -> remove para valores no existentes
		for j := 0; j < 10; j++ {
			value := rand.Intn(50) + 100 // Valor que no está en la lista

			list.Insert(value)
			removed := list.Remove(value)

			if !removed {
				t.Errorf("Remove debería retornar true para valor recién insertado: %d", value)
			}

			if list.Contains(value) {
				t.Errorf("Lista no debería contener valor después de insert->remove: %d", value)
			}
		}

		// Verificar que volvimos al estado inicial
		finalSlice := list.ToSlice()
		finalSize := list.Size()

		if finalSize != initialSize {
			t.Errorf("Tamaño debería volver al inicial después de ciclos insert->remove: %d -> %d", initialSize, finalSize)
		}

		if !slicesEqual(initialSlice, finalSlice) {
			t.Errorf("Lista debería volver al estado inicial: %v -> %v", initialSlice, finalSlice)
		}
	}
}

// ========================================
// Helper functions
// ========================================

func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ========================================
// Benchmarks para Lista Ordenada
// ========================================

func BenchmarkSortedList_Insert(b *testing.B) {
	list := NewSortedList()
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Insert(rand.Intn(10000))
	}
}

func BenchmarkSortedList_Contains(b *testing.B) {
	list := NewSortedList()

	// Preparar lista con datos
	for i := 0; i < 1000; i++ {
		list.Insert(i)
	}

	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Contains(rand.Intn(1000))
	}
}

func BenchmarkSortedList_Remove(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Preparar lista para cada iteración
		list := NewSortedList()
		for j := 0; j < 100; j++ {
			list.Insert(j)
		}
		b.StartTimer()

		list.Remove(rand.Intn(100))
	}
}
