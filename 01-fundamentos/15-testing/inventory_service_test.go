// 🧪 Ejemplo de Implementación: Tests con Mocks y Dependency Injection
// ====================================================================
// Este archivo muestra cómo implementar los tests del Ejercicio 2
// Demuestra mocking manual y testing de servicios complejos

package main

import (
	"errors"
	"testing"
)

// =============================================================================
// 📝 MOCK IMPLEMENTATIONS
// =============================================================================

// MockProductRepository implementación mock de ProductRepository
type MockProductRepository struct {
	products map[int]*Product
	nextID   int

	// Configuración de errores para testing
	saveError   error
	findError   error
	updateError error
	deleteError error

	// Tracking de llamadas para verificación
	saveCalled        bool
	findByIDCalled    map[int]bool
	updateStockCalled map[int]int
	deleteCalled      map[int]bool
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products:          make(map[int]*Product),
		nextID:            1,
		findByIDCalled:    make(map[int]bool),
		updateStockCalled: make(map[int]int),
		deleteCalled:      make(map[int]bool),
	}
}

func (m *MockProductRepository) Save(product *Product) error {
	m.saveCalled = true

	if m.saveError != nil {
		return m.saveError
	}

	if product.ID == 0 {
		product.ID = m.nextID
		m.nextID++
	}

	// Crear copia para evitar modificaciones externas
	m.products[product.ID] = &Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}

	return nil
}

func (m *MockProductRepository) FindByID(id int) (*Product, error) {
	m.findByIDCalled[id] = true

	if m.findError != nil {
		return nil, m.findError
	}

	product, exists := m.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}

	// Retornar copia para evitar modificaciones
	return &Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (m *MockProductRepository) FindByCategory(category string) ([]*Product, error) {
	if m.findError != nil {
		return nil, m.findError
	}

	var results []*Product
	for _, product := range m.products {
		if product.Category == category {
			results = append(results, &Product{
				ID:       product.ID,
				Name:     product.Name,
				Price:    product.Price,
				Stock:    product.Stock,
				Category: product.Category,
			})
		}
	}

	return results, nil
}

func (m *MockProductRepository) UpdateStock(id int, quantity int) error {
	m.updateStockCalled[id] = quantity

	if m.updateError != nil {
		return m.updateError
	}

	product, exists := m.products[id]
	if !exists {
		return errors.New("product not found")
	}

	product.Stock += quantity
	return nil
}

func (m *MockProductRepository) Delete(id int) error {
	m.deleteCalled[id] = true

	if m.deleteError != nil {
		return m.deleteError
	}

	delete(m.products, id)
	return nil
}

// Helper methods para verificación en tests
func (m *MockProductRepository) WasSaveCalled() bool {
	return m.saveCalled
}

func (m *MockProductRepository) WasFindByIDCalled(id int) bool {
	return m.findByIDCalled[id]
}

func (m *MockProductRepository) GetUpdateStockCall(id int) (int, bool) {
	quantity, called := m.updateStockCalled[id]
	return quantity, called
}

func (m *MockProductRepository) SetSaveError(err error) {
	m.saveError = err
}

func (m *MockProductRepository) SetFindError(err error) {
	m.findError = err
}

func (m *MockProductRepository) SetUpdateError(err error) {
	m.updateError = err
}

// MockNotificationService implementación mock de NotificationService
type MockNotificationService struct {
	lowStockAlerts   []*Product
	outOfStockAlerts []*Product
	lowStockError    error
	outOfStockError  error
}

func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{
		lowStockAlerts:   make([]*Product, 0),
		outOfStockAlerts: make([]*Product, 0),
	}
}

func (m *MockNotificationService) SendLowStockAlert(product *Product) error {
	if m.lowStockError != nil {
		return m.lowStockError
	}

	m.lowStockAlerts = append(m.lowStockAlerts, product)
	return nil
}

func (m *MockNotificationService) SendOutOfStockAlert(product *Product) error {
	if m.outOfStockError != nil {
		return m.outOfStockError
	}

	m.outOfStockAlerts = append(m.outOfStockAlerts, product)
	return nil
}

func (m *MockNotificationService) GetLowStockAlerts() []*Product {
	return m.lowStockAlerts
}

func (m *MockNotificationService) GetOutOfStockAlerts() []*Product {
	return m.outOfStockAlerts
}

func (m *MockNotificationService) SetLowStockError(err error) {
	m.lowStockError = err
}

func (m *MockNotificationService) SetOutOfStockError(err error) {
	m.outOfStockError = err
}

// MockPriceCalculator implementación mock de PriceCalculator
type MockPriceCalculator struct {
	discountRate  float64
	taxRate       float64
	discountError error
	taxError      error
}

func NewMockPriceCalculator() *MockPriceCalculator {
	return &MockPriceCalculator{
		discountRate: 0.1, // 10% descuento por defecto
		taxRate:      0.2, // 20% tax por defecto
	}
}

func (m *MockPriceCalculator) CalculateDiscount(price float64, category string) float64 {
	// Simulamos descuentos diferentes por categoría
	switch category {
	case "Electronics":
		return price * (1 - 0.15) // 15% descuento
	case "Books":
		return price * (1 - 0.05) // 5% descuento
	default:
		return price * (1 - m.discountRate)
	}
}

func (m *MockPriceCalculator) CalculateTax(price float64) float64 {
	return price * (1 + m.taxRate)
}

func (m *MockPriceCalculator) SetDiscountRate(rate float64) {
	m.discountRate = rate
}

func (m *MockPriceCalculator) SetTaxRate(rate float64) {
	m.taxRate = rate
}

// =============================================================================
// 📝 TESTS DE INVENTORYSERVICE
// =============================================================================

func TestInventoryService_AddProduct(t *testing.T) {
	tests := []struct {
		name      string
		product   *Product
		wantError bool
		errorMsg  string
		setupMock func(*MockProductRepository)
	}{
		{
			name: "valid product",
			product: &Product{
				Name:     "Test Product",
				Price:    10.99,
				Stock:    100,
				Category: "Test",
			},
			wantError: false,
		},
		{
			name: "empty name",
			product: &Product{
				Name:     "",
				Price:    10.99,
				Stock:    100,
				Category: "Test",
			},
			wantError: true,
			errorMsg:  "product name cannot be empty",
		},
		{
			name: "negative price",
			product: &Product{
				Name:     "Test Product",
				Price:    -10.99,
				Stock:    100,
				Category: "Test",
			},
			wantError: true,
			errorMsg:  "product price must be positive",
		},
		{
			name: "negative stock",
			product: &Product{
				Name:     "Test Product",
				Price:    10.99,
				Stock:    -1,
				Category: "Test",
			},
			wantError: true,
			errorMsg:  "product stock cannot be negative",
		},
		{
			name: "repository error",
			product: &Product{
				Name:     "Test Product",
				Price:    10.99,
				Stock:    100,
				Category: "Test",
			},
			wantError: true,
			setupMock: func(repo *MockProductRepository) {
				repo.SetSaveError(errors.New("database connection failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := NewMockProductRepository()
			mockNotification := NewMockNotificationService()
			mockCalculator := NewMockPriceCalculator()

			if tt.setupMock != nil {
				tt.setupMock(mockRepo)
			}

			service := NewInventoryService(mockRepo, mockNotification, mockCalculator)

			// Act
			err := service.AddProduct(tt.product)

			// Assert
			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				} else if tt.errorMsg != "" && err.Error() != tt.errorMsg {
					t.Errorf("Expected error message '%s', got '%s'",
						tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Verificar que se llamó al repository
				if !mockRepo.WasSaveCalled() {
					t.Error("Expected Save to be called on repository")
				}
			}
		})
	}
}

func TestInventoryService_UpdateStock(t *testing.T) {
	tests := []struct {
		name                  string
		productID             int
		quantity              int
		initialStock          int
		expectedFinalStock    int
		wantError             bool
		expectLowStockAlert   bool
		expectOutOfStockAlert bool
		setupMock             func(*MockProductRepository, *MockNotificationService)
	}{
		{
			name:               "successful stock increase",
			productID:          1,
			quantity:           10,
			initialStock:       20,
			expectedFinalStock: 30,
			wantError:          false,
		},
		{
			name:               "successful stock decrease",
			productID:          1,
			quantity:           -5,
			initialStock:       20,
			expectedFinalStock: 15,
			wantError:          false,
		},
		{
			name:                "stock decrease triggering low stock alert",
			productID:           1,
			quantity:            -15,
			initialStock:        20,
			expectedFinalStock:  5,
			wantError:           false,
			expectLowStockAlert: true,
		},
		{
			name:                  "stock decrease to zero triggering out of stock alert",
			productID:             1,
			quantity:              -20,
			initialStock:          20,
			expectedFinalStock:    0,
			wantError:             false,
			expectOutOfStockAlert: true,
		},
		{
			name:         "insufficient stock",
			productID:    1,
			quantity:     -30,
			initialStock: 20,
			wantError:    true,
		},
		{
			name:      "product not found",
			productID: 999,
			quantity:  10,
			wantError: true,
			setupMock: func(repo *MockProductRepository, notif *MockNotificationService) {
				repo.SetFindError(errors.New("product not found"))
			},
		},
		{
			name:         "repository update error",
			productID:    1,
			quantity:     10,
			initialStock: 20,
			wantError:    true,
			setupMock: func(repo *MockProductRepository, notif *MockNotificationService) {
				repo.SetUpdateError(errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := NewMockProductRepository()
			mockNotification := NewMockNotificationService()
			mockCalculator := NewMockPriceCalculator()

			// Setup product inicial si es necesario
			if tt.initialStock > 0 {
				product := &Product{
					ID:       tt.productID,
					Name:     "Test Product",
					Price:    10.99,
					Stock:    tt.initialStock,
					Category: "Test",
				}
				mockRepo.Save(product)
			}

			if tt.setupMock != nil {
				tt.setupMock(mockRepo, mockNotification)
			}

			service := NewInventoryService(mockRepo, mockNotification, mockCalculator)

			// Act
			err := service.UpdateStock(tt.productID, tt.quantity)

			// Assert
			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Verificar que se llamó UpdateStock en el repository
				quantity, called := mockRepo.GetUpdateStockCall(tt.productID)
				if !called {
					t.Error("Expected UpdateStock to be called on repository")
				}
				if quantity != tt.quantity {
					t.Errorf("Expected UpdateStock called with quantity %d, got %d",
						tt.quantity, quantity)
				}

				// Verificar alertas
				if tt.expectLowStockAlert {
					alerts := mockNotification.GetLowStockAlerts()
					if len(alerts) != 1 {
						t.Errorf("Expected 1 low stock alert, got %d", len(alerts))
					}
				}

				if tt.expectOutOfStockAlert {
					alerts := mockNotification.GetOutOfStockAlerts()
					if len(alerts) != 1 {
						t.Errorf("Expected 1 out of stock alert, got %d", len(alerts))
					}
				}
			}
		})
	}
}

func TestInventoryService_GetFinalPrice(t *testing.T) {
	tests := []struct {
		name          string
		productID     int
		basePrice     float64
		category      string
		expectedPrice float64
		wantError     bool
		setupMock     func(*MockProductRepository, *MockPriceCalculator)
	}{
		{
			name:          "electronics product with discount",
			productID:     1,
			basePrice:     100.0,
			category:      "Electronics",
			expectedPrice: 102.0, // 100 * 0.85 * 1.2 = 102
			wantError:     false,
		},
		{
			name:          "books product with discount",
			productID:     2,
			basePrice:     20.0,
			category:      "Books",
			expectedPrice: 22.8, // 20 * 0.95 * 1.2 = 22.8
			wantError:     false,
		},
		{
			name:          "default category product",
			productID:     3,
			basePrice:     50.0,
			category:      "Other",
			expectedPrice: 54.0, // 50 * 0.9 * 1.2 = 54
			wantError:     false,
		},
		{
			name:      "product not found",
			productID: 999,
			wantError: true,
			setupMock: func(repo *MockProductRepository, calc *MockPriceCalculator) {
				repo.SetFindError(errors.New("product not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := NewMockProductRepository()
			mockNotification := NewMockNotificationService()
			mockCalculator := NewMockPriceCalculator()

			// Setup product si es necesario
			if tt.basePrice > 0 {
				product := &Product{
					ID:       tt.productID,
					Name:     "Test Product",
					Price:    tt.basePrice,
					Stock:    10,
					Category: tt.category,
				}
				mockRepo.Save(product)
			}

			if tt.setupMock != nil {
				tt.setupMock(mockRepo, mockCalculator)
			}

			service := NewInventoryService(mockRepo, mockNotification, mockCalculator)

			// Act
			finalPrice, err := service.GetFinalPrice(tt.productID)

			// Assert
			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Verificar precio final (con tolerancia para floating point)
				tolerance := 0.01
				if abs(finalPrice-tt.expectedPrice) > tolerance {
					t.Errorf("Expected final price %.2f, got %.2f",
						tt.expectedPrice, finalPrice)
				}

				// Verificar que se consultó el producto
				if !mockRepo.WasFindByIDCalled(tt.productID) {
					t.Error("Expected FindByID to be called on repository")
				}
			}
		})
	}
}

// Helper function para valor absoluto
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// =============================================================================
// 📝 TESTS DE COMPORTAMIENTO DE MOCKS
// =============================================================================

func TestMockVerification(t *testing.T) {
	t.Run("mock repository tracks calls correctly", func(t *testing.T) {
		// Arrange
		mockRepo := NewMockProductRepository()
		product := &Product{Name: "Test", Price: 10.0, Stock: 5, Category: "Test"}

		// Act
		err := mockRepo.Save(product)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		_, err = mockRepo.FindByID(product.ID)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = mockRepo.UpdateStock(product.ID, 10)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Assert
		if !mockRepo.WasSaveCalled() {
			t.Error("Expected Save to be tracked as called")
		}

		if !mockRepo.WasFindByIDCalled(product.ID) {
			t.Error("Expected FindByID to be tracked as called")
		}

		quantity, called := mockRepo.GetUpdateStockCall(product.ID)
		if !called {
			t.Error("Expected UpdateStock to be tracked as called")
		}
		if quantity != 10 {
			t.Errorf("Expected UpdateStock called with 10, got %d", quantity)
		}
	})

	t.Run("mock notification service tracks alerts", func(t *testing.T) {
		// Arrange
		mockNotif := NewMockNotificationService()
		product := &Product{ID: 1, Name: "Test Product"}

		// Act
		err := mockNotif.SendLowStockAlert(product)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		err = mockNotif.SendOutOfStockAlert(product)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Assert
		lowStockAlerts := mockNotif.GetLowStockAlerts()
		if len(lowStockAlerts) != 1 {
			t.Errorf("Expected 1 low stock alert, got %d", len(lowStockAlerts))
		}

		outOfStockAlerts := mockNotif.GetOutOfStockAlerts()
		if len(outOfStockAlerts) != 1 {
			t.Errorf("Expected 1 out of stock alert, got %d", len(outOfStockAlerts))
		}
	})
}

// =============================================================================
// 📝 BENCHMARKS
// =============================================================================

func BenchmarkInventoryService_AddProduct(b *testing.B) {
	mockRepo := NewMockProductRepository()
	mockNotification := NewMockNotificationService()
	mockCalculator := NewMockPriceCalculator()
	service := NewInventoryService(mockRepo, mockNotification, mockCalculator)

	product := &Product{
		Name:     "Benchmark Product",
		Price:    99.99,
		Stock:    100,
		Category: "Test",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset product ID for each iteration
		testProduct := &Product{
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
		}
		service.AddProduct(testProduct)
	}
}

func BenchmarkInventoryService_GetFinalPrice(b *testing.B) {
	mockRepo := NewMockProductRepository()
	mockNotification := NewMockNotificationService()
	mockCalculator := NewMockPriceCalculator()
	service := NewInventoryService(mockRepo, mockNotification, mockCalculator)

	// Setup product
	product := &Product{
		ID:       1,
		Name:     "Benchmark Product",
		Price:    99.99,
		Stock:    100,
		Category: "Electronics",
	}
	mockRepo.Save(product)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.GetFinalPrice(1)
	}
}
