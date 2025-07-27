// 🧪 Ejemplo de Implementación: Tests de Concurrencia
// ==================================================
// Este archivo muestra cómo testear código concurrente y detectar race conditions
// Demuestra el uso de go test -race y testing paralelo

package main

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

// =============================================================================
// 📝 TESTS BÁSICOS DEL SHOPPING CART
// =============================================================================

func TestShoppingCart_AddItem(t *testing.T) {
	tests := []struct {
		name         string
		productID    int
		productName  string
		price        float64
		quantity     int
		wantError    bool
		errorMessage string
	}{
		{
			name:        "valid item",
			productID:   1,
			productName: "Test Product",
			price:       10.99,
			quantity:    2,
			wantError:   false,
		},
		{
			name:         "zero quantity",
			productID:    1,
			productName:  "Test Product",
			price:        10.99,
			quantity:     0,
			wantError:    true,
			errorMessage: "quantity must be positive",
		},
		{
			name:         "negative quantity",
			productID:    1,
			productName:  "Test Product",
			price:        10.99,
			quantity:     -1,
			wantError:    true,
			errorMessage: "quantity must be positive",
		},
		{
			name:         "negative price",
			productID:    1,
			productName:  "Test Product",
			price:        -10.99,
			quantity:     1,
			wantError:    true,
			errorMessage: "price cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			cart := NewShoppingCart()

			// Act
			err := cart.AddItem(tt.productID, tt.productName, tt.price, tt.quantity)

			// Assert
			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				if err != nil && err.Error() != tt.errorMessage {
					t.Errorf("Expected error message '%s', got '%s'",
						tt.errorMessage, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Verificar que el item fue añadido
				items := cart.GetItems()
				item, exists := items[tt.productID]
				if !exists {
					t.Error("Expected item to be added to cart")
				} else {
					if item.ProductID != tt.productID {
						t.Errorf("Expected ProductID %d, got %d",
							tt.productID, item.ProductID)
					}
					if item.Name != tt.productName {
						t.Errorf("Expected Name '%s', got '%s'",
							tt.productName, item.Name)
					}
					if item.Price != tt.price {
						t.Errorf("Expected Price %.2f, got %.2f",
							tt.price, item.Price)
					}
					if item.Quantity != tt.quantity {
						t.Errorf("Expected Quantity %d, got %d",
							tt.quantity, item.Quantity)
					}
				}

				// Verificar total
				expectedTotal := tt.price * float64(tt.quantity)
				if cart.GetTotal() != expectedTotal {
					t.Errorf("Expected total %.2f, got %.2f",
						expectedTotal, cart.GetTotal())
				}
			}
		})
	}
}

func TestShoppingCart_AddItemTwice(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	productID := 1
	productName := "Test Product"
	price := 10.99
	quantity1 := 2
	quantity2 := 3

	// Act
	err1 := cart.AddItem(productID, productName, price, quantity1)
	err2 := cart.AddItem(productID, productName, price, quantity2)

	// Assert
	if err1 != nil {
		t.Errorf("First AddItem failed: %v", err1)
	}
	if err2 != nil {
		t.Errorf("Second AddItem failed: %v", err2)
	}

	items := cart.GetItems()
	item, exists := items[productID]
	if !exists {
		t.Fatal("Expected item to exist in cart")
	}

	expectedQuantity := quantity1 + quantity2
	if item.Quantity != expectedQuantity {
		t.Errorf("Expected quantity %d, got %d", expectedQuantity, item.Quantity)
	}

	expectedTotal := price * float64(expectedQuantity)
	if cart.GetTotal() != expectedTotal {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, cart.GetTotal())
	}
}

func TestShoppingCart_RemoveItem(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	productID := 1
	cart.AddItem(productID, "Test Product", 10.99, 2)

	// Act
	err := cart.RemoveItem(productID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	items := cart.GetItems()
	if _, exists := items[productID]; exists {
		t.Error("Expected item to be removed from cart")
	}

	if cart.GetTotal() != 0 {
		t.Errorf("Expected total 0, got %.2f", cart.GetTotal())
	}

	if cart.GetItemCount() != 0 {
		t.Errorf("Expected item count 0, got %d", cart.GetItemCount())
	}
}

func TestShoppingCart_RemoveNonExistentItem(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()

	// Act
	err := cart.RemoveItem(999)

	// Assert
	if err == nil {
		t.Error("Expected error when removing non-existent item")
	}
	if err.Error() != "item not found in cart" {
		t.Errorf("Expected 'item not found in cart', got '%s'", err.Error())
	}
}

func TestShoppingCart_UpdateQuantity(t *testing.T) {
	tests := []struct {
		name         string
		initialQty   int
		newQty       int
		expectRemove bool
		wantError    bool
	}{
		{
			name:       "increase quantity",
			initialQty: 2,
			newQty:     5,
		},
		{
			name:       "decrease quantity",
			initialQty: 5,
			newQty:     2,
		},
		{
			name:         "zero quantity removes item",
			initialQty:   2,
			newQty:       0,
			expectRemove: true,
		},
		{
			name:         "negative quantity removes item",
			initialQty:   2,
			newQty:       -1,
			expectRemove: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			cart := NewShoppingCart()
			productID := 1
			price := 10.99
			cart.AddItem(productID, "Test Product", price, tt.initialQty)

			// Act
			err := cart.UpdateQuantity(productID, tt.newQty)

			// Assert
			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				items := cart.GetItems()
				if tt.expectRemove {
					if _, exists := items[productID]; exists {
						t.Error("Expected item to be removed")
					}
					if cart.GetTotal() != 0 {
						t.Errorf("Expected total 0, got %.2f", cart.GetTotal())
					}
				} else {
					item, exists := items[productID]
					if !exists {
						t.Fatal("Expected item to exist")
					}
					if item.Quantity != tt.newQty {
						t.Errorf("Expected quantity %d, got %d",
							tt.newQty, item.Quantity)
					}
					expectedTotal := price * float64(tt.newQty)
					if cart.GetTotal() != expectedTotal {
						t.Errorf("Expected total %.2f, got %.2f",
							expectedTotal, cart.GetTotal())
					}
				}
			}
		})
	}
}

func TestShoppingCart_GetItemCount(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()

	// Initially empty
	if cart.GetItemCount() != 0 {
		t.Errorf("Expected item count 0, got %d", cart.GetItemCount())
	}

	// Add items
	cart.AddItem(1, "Product 1", 10.99, 2)
	cart.AddItem(2, "Product 2", 5.99, 3)
	cart.AddItem(3, "Product 3", 15.99, 1)

	expectedCount := 2 + 3 + 1 // Total quantity across all items
	if cart.GetItemCount() != expectedCount {
		t.Errorf("Expected item count %d, got %d", expectedCount, cart.GetItemCount())
	}
}

func TestShoppingCart_Clear(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	cart.AddItem(1, "Product 1", 10.99, 2)
	cart.AddItem(2, "Product 2", 5.99, 3)

	// Act
	cart.Clear()

	// Assert
	if len(cart.GetItems()) != 0 {
		t.Error("Expected cart to be empty after clear")
	}
	if cart.GetTotal() != 0 {
		t.Errorf("Expected total 0 after clear, got %.2f", cart.GetTotal())
	}
	if cart.GetItemCount() != 0 {
		t.Errorf("Expected item count 0 after clear, got %d", cart.GetItemCount())
	}
}

// =============================================================================
// 📝 TESTS DE CONCURRENCIA
// =============================================================================

func TestShoppingCart_ConcurrentAddItems(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	numGoroutines := 100
	itemsPerGoroutine := 10
	var wg sync.WaitGroup

	// Act - múltiples goroutines añadiendo items
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < itemsPerGoroutine; j++ {
				productID := goroutineID*itemsPerGoroutine + j
				err := cart.AddItem(productID, "Test Product", 1.0, 1)
				if err != nil {
					t.Errorf("Goroutine %d: AddItem failed: %v", goroutineID, err)
				}
			}
		}(i)
	}

	wg.Wait()

	// Assert
	expectedItems := numGoroutines * itemsPerGoroutine
	items := cart.GetItems()
	if len(items) != expectedItems {
		t.Errorf("Expected %d items, got %d", expectedItems, len(items))
	}

	expectedTotal := float64(expectedItems) * 1.0
	if cart.GetTotal() != expectedTotal {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, cart.GetTotal())
	}
}

func TestShoppingCart_ConcurrentAddSameItem(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	numGoroutines := 50
	quantityPerGoroutine := 2
	productID := 1
	price := 10.99
	var wg sync.WaitGroup

	// Act - múltiples goroutines añadiendo el mismo item
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := cart.AddItem(productID, "Test Product", price, quantityPerGoroutine)
			if err != nil {
				t.Errorf("AddItem failed: %v", err)
			}
		}()
	}

	wg.Wait()

	// Assert
	items := cart.GetItems()
	if len(items) != 1 {
		t.Errorf("Expected 1 unique item, got %d", len(items))
	}

	item, exists := items[productID]
	if !exists {
		t.Fatal("Expected item to exist")
	}

	expectedQuantity := numGoroutines * quantityPerGoroutine
	if item.Quantity != expectedQuantity {
		t.Errorf("Expected quantity %d, got %d", expectedQuantity, item.Quantity)
	}

	expectedTotal := price * float64(expectedQuantity)
	if cart.GetTotal() != expectedTotal {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, cart.GetTotal())
	}
}

func TestShoppingCart_ConcurrentOperations(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	numOperations := 1000
	var wg sync.WaitGroup

	// Act - operaciones concurrentes mixtas
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			switch id % 4 {
			case 0: // Add item
				cart.AddItem(id, "Product", 1.0, 1)
			case 1: // Update quantity
				cart.UpdateQuantity(id/2, 2)
			case 2: // Get total
				cart.GetTotal()
			case 3: // Get items
				cart.GetItems()
			}
		}(i)
	}

	wg.Wait()

	// Assert - verificar que no hay inconsistencias
	items := cart.GetItems()
	calculatedTotal := 0.0
	for _, item := range items {
		calculatedTotal += item.Price * float64(item.Quantity)
	}

	if cart.GetTotal() != calculatedTotal {
		t.Errorf("Total inconsistency: stored %.2f, calculated %.2f",
			cart.GetTotal(), calculatedTotal)
	}
}

func TestShoppingCart_ConcurrentReadWrite(t *testing.T) {
	// Arrange
	cart := NewShoppingCart()
	numReaders := 10
	numWriters := 5
	duration := 100 * time.Millisecond
	var wg sync.WaitGroup

	// Populate cart initially
	for i := 0; i < 10; i++ {
		cart.AddItem(i, "Product", 10.0, 1)
	}

	// Start readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			start := time.Now()
			readCount := 0

			for time.Since(start) < duration {
				_ = cart.GetTotal()
				_ = cart.GetItems()
				_ = cart.GetItemCount()
				readCount++
				runtime.Gosched() // Yield to other goroutines
			}

			t.Logf("Reader %d performed %d reads", readerID, readCount)
		}(i)
	}

	// Start writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			start := time.Now()
			writeCount := 0

			for time.Since(start) < duration {
				productID := writerID*1000 + writeCount
				cart.AddItem(productID, "Product", 1.0, 1)
				writeCount++
				runtime.Gosched() // Yield to other goroutines
			}

			t.Logf("Writer %d performed %d writes", writerID, writeCount)
		}(i)
	}

	wg.Wait()

	// Verify consistency
	items := cart.GetItems()
	calculatedTotal := 0.0
	for _, item := range items {
		calculatedTotal += item.Price * float64(item.Quantity)
	}

	if cart.GetTotal() != calculatedTotal {
		t.Errorf("Final total inconsistency: stored %.2f, calculated %.2f",
			cart.GetTotal(), calculatedTotal)
	}
}

// =============================================================================
// 📝 RACE CONDITION DETECTION TESTS
// =============================================================================

// Este test está diseñado para detectar race conditions
// Ejecutar con: go test -race -run TestShoppingCart_RaceDetection
func TestShoppingCart_RaceDetection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping race detection test in short mode")
	}

	// Arrange
	cart := NewShoppingCart()
	iterations := 1000
	var wg sync.WaitGroup

	// Act - operaciones que podrían causar race conditions
	for i := 0; i < iterations; i++ {
		wg.Add(3) // Tres operaciones concurrentes por iteración

		// Goroutine 1: Add item
		go func(id int) {
			defer wg.Done()
			cart.AddItem(id, "Product", 1.0, 1)
		}(i)

		// Goroutine 2: Read operations
		go func() {
			defer wg.Done()
			cart.GetTotal()
			cart.GetItems()
		}()

		// Goroutine 3: Update operations
		go func(id int) {
			defer wg.Done()
			if id > 0 {
				cart.UpdateQuantity(id-1, 2)
			}
		}(i)
	}

	wg.Wait()

	// Assert - verificar consistencia final
	items := cart.GetItems()
	calculatedTotal := 0.0
	for _, item := range items {
		calculatedTotal += item.Price * float64(item.Quantity)
	}

	tolerance := 0.01 // Pequeña tolerancia para floating point
	if abs(cart.GetTotal()-calculatedTotal) > tolerance {
		t.Errorf("Race condition detected: total mismatch %.2f vs %.2f",
			cart.GetTotal(), calculatedTotal)
	}
}

// =============================================================================
// 📝 BENCHMARKS DE CONCURRENCIA
// =============================================================================

func BenchmarkShoppingCart_AddItem_Sequential(b *testing.B) {
	cart := NewShoppingCart()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cart.AddItem(i, "Product", 1.0, 1)
	}
}

func BenchmarkShoppingCart_AddItem_Parallel(b *testing.B) {
	cart := NewShoppingCart()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cart.AddItem(i, "Product", 1.0, 1)
			i++
		}
	})
}

func BenchmarkShoppingCart_GetTotal_Sequential(b *testing.B) {
	cart := NewShoppingCart()

	// Setup - añadir items
	for i := 0; i < 1000; i++ {
		cart.AddItem(i, "Product", 1.0, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cart.GetTotal()
	}
}

func BenchmarkShoppingCart_GetTotal_Parallel(b *testing.B) {
	cart := NewShoppingCart()

	// Setup - añadir items
	for i := 0; i < 1000; i++ {
		cart.AddItem(i, "Product", 1.0, 1)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cart.GetTotal()
		}
	})
}

func BenchmarkShoppingCart_MixedOperations(b *testing.B) {
	cart := NewShoppingCart()

	// Setup inicial
	for i := 0; i < 100; i++ {
		cart.AddItem(i, "Product", 1.0, 1)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			switch i % 3 {
			case 0:
				cart.AddItem(1000+i, "Product", 1.0, 1)
			case 1:
				cart.GetTotal()
			case 2:
				cart.GetItems()
			}
			i++
		}
	})
}

// =============================================================================
// 📝 STRESS TESTS
// =============================================================================

func TestShoppingCart_StressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	// Arrange
	cart := NewShoppingCart()
	numGoroutines := 100
	operationsPerGoroutine := 1000
	var wg sync.WaitGroup

	// Act
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < operationsPerGoroutine; j++ {
				productID := goroutineID*operationsPerGoroutine + j

				// 70% adds, 20% reads, 10% updates
				switch j % 10 {
				case 0, 1, 2, 3, 4, 5, 6: // Add operations
					cart.AddItem(productID, "Product", 1.0, 1)
				case 7, 8: // Read operations
					cart.GetTotal()
					cart.GetItems()
				case 9: // Update operations
					if productID > 0 {
						cart.UpdateQuantity(productID-1, 2)
					}
				}
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	// Assert
	totalOperations := numGoroutines * operationsPerGoroutine
	operationsPerSecond := float64(totalOperations) / duration.Seconds()

	t.Logf("Stress test completed:")
	t.Logf("- Goroutines: %d", numGoroutines)
	t.Logf("- Operations per goroutine: %d", operationsPerGoroutine)
	t.Logf("- Total operations: %d", totalOperations)
	t.Logf("- Duration: %v", duration)
	t.Logf("- Operations per second: %.2f", operationsPerSecond)
	t.Logf("- Final cart items: %d", len(cart.GetItems()))
	t.Logf("- Final cart total: %.2f", cart.GetTotal())

	// Verify consistency
	items := cart.GetItems()
	calculatedTotal := 0.0
	for _, item := range items {
		calculatedTotal += item.Price * float64(item.Quantity)
	}

	if abs(cart.GetTotal()-calculatedTotal) > 0.01 {
		t.Errorf("Stress test revealed inconsistency: stored %.2f, calculated %.2f",
			cart.GetTotal(), calculatedTotal)
	}
}

// =============================================================================
// 📝 HELPER FUNCTIONS PARA TESTS DE CONCURRENCIA
// =============================================================================

// Helper para crear múltiples items de prueba
func addMultipleItems(t *testing.T, cart *ShoppingCart, count int) {
	t.Helper()

	for i := 0; i < count; i++ {
		err := cart.AddItem(i, "Product", 10.0, 1)
		if err != nil {
			t.Fatalf("Failed to add item %d: %v", i, err)
		}
	}
}

// Helper para verificar consistencia del carrito
func verifyCartConsistency(t *testing.T, cart *ShoppingCart) {
	t.Helper()

	items := cart.GetItems()
	calculatedTotal := 0.0
	calculatedCount := 0

	for _, item := range items {
		calculatedTotal += item.Price * float64(item.Quantity)
		calculatedCount += item.Quantity
	}

	if abs(cart.GetTotal()-calculatedTotal) > 0.01 {
		t.Errorf("Total inconsistency: stored %.2f, calculated %.2f",
			cart.GetTotal(), calculatedTotal)
	}

	if cart.GetItemCount() != calculatedCount {
		t.Errorf("Count inconsistency: stored %d, calculated %d",
			cart.GetItemCount(), calculatedCount)
	}
}

// Helper para ejecutar operaciones concurrentes con timeout
func runConcurrentOperations(t *testing.T, operations []func(), timeout time.Duration) {
	t.Helper()

	var wg sync.WaitGroup
	done := make(chan bool, 1)

	// Start operations
	for _, op := range operations {
		wg.Add(1)
		go func(operation func()) {
			defer wg.Done()
			operation()
		}(op)
	}

	// Wait for completion or timeout
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Operations completed successfully
	case <-time.After(timeout):
		t.Error("Operations timed out")
	}
}
