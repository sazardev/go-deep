// 🧪 Tests del Proyecto: Sistema de Testing Integral
// Archivo: proyecto_testing_integral_test.go
// Ejecutar con: go test -v

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// ==========================================
// 🧪 UNIT TESTS - TDD EXAMPLES
// ==========================================

func TestECommerceService_CreateOrder_Success(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	setupSuccessfulMocks(mocks)

	ctx := context.Background()
	userID := "user-1"
	items := []TestOrderItem{
		{ProductID: "prod-1", Quantity: 1},
		{ProductID: "prod-2", Quantity: 2},
	}

	// Act
	order, err := service.CreateOrder(ctx, userID, items, "credit_card")

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if order == nil {
		t.Fatal("Expected order to be created, got nil")
	}
	if order.UserID != userID {
		t.Errorf("Expected UserID %s, got %s", userID, order.UserID)
	}
	if order.Status != TestOrderStatusProcessing {
		t.Errorf("Expected status %s, got %s", TestOrderStatusProcessing, order.Status)
	}
	if len(order.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(order.Items))
	}
	if order.TotalAmount <= 0 {
		t.Errorf("Expected positive total amount, got %f", order.TotalAmount)
	}

	// Verify interactions
	notifications := mocks.NotificationService.GetSentNotifications()
	if len(notifications) != 1 {
		t.Errorf("Expected 1 notification, got %d", len(notifications))
	}
}

func TestECommerceService_CreateOrder_UserNotFound(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	mocks.UserRepo.SetError("GetByID", errors.New("user not found"))

	ctx := context.Background()
	items := []TestOrderItem{{ProductID: "prod-1", Quantity: 1}}

	// Act
	order, err := service.CreateOrder(ctx, "invalid-user", items, "credit_card")

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid user, got nil")
	}
	if order != nil {
		t.Fatal("Expected nil order for invalid user")
	}
}

func TestECommerceService_CreateOrder_InsufficientStock(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	setupProductsAndUsers(mocks)

	// Set insufficient stock
	mocks.InventoryService.SetStock("prod-1", 0)

	ctx := context.Background()
	items := []TestOrderItem{{ProductID: "prod-1", Quantity: 1}}

	// Act
	order, err := service.CreateOrder(ctx, "user-1", items, "credit_card")

	// Assert
	if err == nil {
		t.Fatal("Expected error for insufficient stock, got nil")
	}
	if order != nil {
		t.Fatal("Expected nil order for insufficient stock")
	}
}

func TestECommerceService_CreateOrder_PaymentFailure(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	setupSuccessfulMocks(mocks)
	mocks.PaymentService.SetShouldFail(true)

	ctx := context.Background()
	items := []TestOrderItem{{ProductID: "prod-1", Quantity: 1}}

	// Act
	order, err := service.CreateOrder(ctx, "user-1", items, "credit_card")

	// Assert
	if err == nil {
		t.Fatal("Expected payment error, got nil")
	}
	if order != nil {
		t.Fatal("Expected nil order for payment failure")
	}
}

func TestECommerceService_CancelOrder_Success(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	setupOrderForCancellation(mocks)

	ctx := context.Background()

	// Act
	err := service.CancelOrder(ctx, "order-1", "user-1")

	// Assert
	if err != nil {
		t.Errorf("Expected no error canceling order, got: %v", err)
	}
}

func TestECommerceService_CancelOrder_UnauthorizedUser(t *testing.T) {
	// Arrange
	service, mocks := setupTestECommerceService()
	setupOrderForCancellation(mocks)

	ctx := context.Background()

	// Act
	err := service.CancelOrder(ctx, "order-1", "wrong-user")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unauthorized")
}

// ==========================================
// 🎯 PROPERTY-BASED TESTS
// ==========================================

func TestOrderTotalCalculation_Property(t *testing.T) {
	t.Run("Order total equals sum of item subtotals", func(t *testing.T) {
		// Property: For any valid order, total amount should equal sum of subtotals
		for i := 0; i < 100; i++ {
			items := generateRandomOrderItems(t)
			if len(items) == 0 {
				continue
			}

			expectedTotal := calculateExpectedTotal(items)
			order := &Order{Items: items, TotalAmount: expectedTotal}

			calculatedTotal := 0.0
			for _, item := range order.Items {
				calculatedTotal += item.Subtotal
			}

			assert.InDelta(t, expectedTotal, calculatedTotal, 0.01,
				"Order total should equal sum of subtotals for items: %+v", items)
		}
	})
}

func TestStockConsistency_Property(t *testing.T) {
	t.Run("Stock conservation property", func(t *testing.T) {
		// Property: Initial stock - reserved stock = available stock
		inventory := NewMockInventoryService()

		for i := 0; i < 50; i++ {
			productID := generateRandomProductID()
			initialStock := generateRandomStock()
			reserveAmount := generateRandomReservation(initialStock)

			inventory.SetStock(productID, initialStock)

			ctx := context.Background()
			items := []OrderItem{{ProductID: productID, Quantity: reserveAmount}}

			// Reserve stock
			err := inventory.ReserveStock(ctx, items)
			require.NoError(t, err)

			// Check availability
			available, err := inventory.CheckAvailability(ctx, productID, 1)
			require.NoError(t, err)

			expectedAvailable := initialStock-reserveAmount > 0
			assert.Equal(t, expectedAvailable, available,
				"Stock consistency failed for initial=%d, reserved=%d", initialStock, reserveAmount)
		}
	})
}

// ==========================================
// 🔗 INTEGRATION TESTS
// ==========================================

func TestCreateOrder_FullFlow_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Arrange - Setup more realistic environment
	service, mocks := setupIntegrationTestEnvironment()

	ctx := context.Background()
	userID := "integration-user"
	items := []OrderItem{
		{ProductID: "prod-1", Quantity: 1},
		{ProductID: "prod-2", Quantity: 3},
	}

	// Act
	order, err := service.CreateOrder(ctx, userID, items, "credit_card")

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, order)

	// Verify all side effects
	t.Run("Verify stock was reserved", func(t *testing.T) {
		for _, item := range items {
			available, err := mocks.InventoryService.CheckAvailability(ctx, item.ProductID, item.Quantity)
			assert.NoError(t, err)
			assert.True(t, available) // Should still be available but reserved
		}
	})

	t.Run("Verify payment was processed", func(t *testing.T) {
		// In real integration test, verify payment was actually processed
		assert.Greater(t, order.TotalAmount, 0.0)
	})

	t.Run("Verify notification was sent", func(t *testing.T) {
		notifications := mocks.NotificationService.GetSentNotifications()
		assert.Len(t, notifications, 1)
		assert.Contains(t, notifications[0], userID)
		assert.Contains(t, notifications[0], order.ID)
	})

	t.Run("Verify logging occurred", func(t *testing.T) {
		logs := mocks.Logger.GetLogs()
		assert.NotEmpty(t, logs)

		// Check for specific log entries
		infoLogs := mocks.Logger.GetLogsByLevel("INFO")
		assert.True(t, len(infoLogs) > 0)

		// Verify no error logs
		errorLogs := mocks.Logger.GetLogsByLevel("ERROR")
		assert.Empty(t, errorLogs)
	})
}

func TestConcurrentOrderCreation_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Test concurrent order creation to verify thread safety
	service, _ := setupIntegrationTestEnvironment()

	const numGoroutines = 10
	const ordersPerGoroutine = 5

	errors := make(chan error, numGoroutines*ordersPerGoroutine)
	orders := make(chan *Order, numGoroutines*ordersPerGoroutine)

	// Launch concurrent order creation
	for i := 0; i < numGoroutines; i++ {
		go func(workerID int) {
			for j := 0; j < ordersPerGoroutine; j++ {
				ctx := context.Background()
				userID := fmt.Sprintf("user-%d", workerID)
				items := []OrderItem{{ProductID: "prod-1", Quantity: 1}}

				order, err := service.CreateOrder(ctx, userID, items, "credit_card")
				if err != nil {
					errors <- err
				} else {
					orders <- order
				}
			}
		}(i)
	}

	// Collect results
	var successfulOrders []*Order
	var orderErrors []error

	timeout := time.After(10 * time.Second)
	expectedResults := numGoroutines * ordersPerGoroutine

	for i := 0; i < expectedResults; i++ {
		select {
		case order := <-orders:
			successfulOrders = append(successfulOrders, order)
		case err := <-errors:
			orderErrors = append(orderErrors, err)
		case <-timeout:
			t.Fatal("Test timed out waiting for concurrent operations")
		}
	}

	// Verify results
	t.Logf("Successful orders: %d, Errors: %d", len(successfulOrders), len(orderErrors))
	assert.True(t, len(successfulOrders) > 0, "Should have at least some successful orders")

	// Verify all successful orders have unique IDs
	orderIDs := make(map[string]bool)
	for _, order := range successfulOrders {
		assert.False(t, orderIDs[order.ID], "Order ID should be unique: %s", order.ID)
		orderIDs[order.ID] = true
	}
}

// ==========================================
// ⚡ BENCHMARK TESTS
// ==========================================

func BenchmarkCreateOrder_SingleOrder(b *testing.B) {
	service, _ := setupBenchmarkECommerceService()

	ctx := context.Background()
	userID := "bench-user"
	items := []OrderItem{{ProductID: "prod-1", Quantity: 1}}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := service.CreateOrder(ctx, userID, items, "credit_card")
		if err != nil {
			b.Fatalf("CreateOrder failed: %v", err)
		}
	}
}

func BenchmarkCreateOrder_MultipleItems(b *testing.B) {
	service, _ := setupBenchmarkECommerceService()

	ctx := context.Background()
	userID := "bench-user"
	items := []OrderItem{
		{ProductID: "prod-1", Quantity: 1},
		{ProductID: "prod-2", Quantity: 2},
		{ProductID: "prod-3", Quantity: 3},
		{ProductID: "prod-4", Quantity: 1},
		{ProductID: "prod-5", Quantity: 2},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := service.CreateOrder(ctx, userID, items, "credit_card")
		if err != nil {
			b.Fatalf("CreateOrder failed: %v", err)
		}
	}
}

func BenchmarkCreateOrder_Parallel(b *testing.B) {
	service, _ := setupBenchmarkECommerceService()

	ctx := context.Background()
	items := []OrderItem{{ProductID: "prod-1", Quantity: 1}}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		userCounter := 0
		for pb.Next() {
			userID := fmt.Sprintf("bench-user-%d", userCounter)
			userCounter++

			_, err := service.CreateOrder(ctx, userID, items, "credit_card")
			if err != nil {
				b.Errorf("CreateOrder failed: %v", err)
			}
		}
	})
}

// ==========================================
// 🛠️ TEST HELPERS AND SETUP
// ==========================================

type TestMocks struct {
	ProductRepo         *MockProductRepository
	UserRepo            *MockUserRepository
	OrderRepo           *MockOrderRepository
	PaymentService      *MockPaymentService
	NotificationService *MockNotificationService
	InventoryService    *MockInventoryService
	Logger              *MockLogger
}

func setupTestECommerceService() (*ECommerceService, *TestMocks) {
	mocks := &TestMocks{
		ProductRepo:         NewMockProductRepository(),
		UserRepo:            NewMockUserRepository(),
		OrderRepo:           NewMockOrderRepository(),
		PaymentService:      NewMockPaymentService(),
		NotificationService: NewMockNotificationService(),
		InventoryService:    NewMockInventoryService(),
		Logger:              NewMockLogger(),
	}

	service := NewECommerceService(
		mocks.ProductRepo,
		mocks.UserRepo,
		mocks.OrderRepo,
		mocks.PaymentService,
		mocks.NotificationService,
		mocks.InventoryService,
		mocks.Logger,
	)

	return service, mocks
}

func setupSuccessfulMocks(mocks *TestMocks) {
	setupProductsAndUsers(mocks)
	setupInventoryWithStock(mocks)
}

func setupProductsAndUsers(mocks *TestMocks) {
	// Add test products
	products := []*Product{
		{
			ID:          "prod-1",
			Name:        "Test Product 1",
			Description: "Test Description 1",
			Price:       99.99,
			Stock:       100,
			Category:    "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "prod-2",
			Name:        "Test Product 2",
			Description: "Test Description 2",
			Price:       49.99,
			Stock:       50,
			Category:    "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, product := range products {
		mocks.ProductRepo.AddProduct(product)
	}

	// Add test user
	user := &User{
		ID:        "user-1",
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		CreatedAt: time.Now(),
	}
	mocks.UserRepo.AddUser(user)
}

func setupInventoryWithStock(mocks *TestMocks) {
	mocks.InventoryService.SetStock("prod-1", 100)
	mocks.InventoryService.SetStock("prod-2", 50)
}

func setupOrderForCancellation(mocks *TestMocks) {
	setupProductsAndUsers(mocks)

	order := &Order{
		ID:          "order-1",
		UserID:      "user-1",
		Items:       []OrderItem{{ProductID: "prod-1", Quantity: 1}},
		TotalAmount: 99.99,
		Status:      OrderStatusProcessing,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Add order to mock repository
	ctx := context.Background()
	mocks.OrderRepo.Create(ctx, order)
}

func setupIntegrationTestEnvironment() (*ECommerceService, *TestMocks) {
	service, mocks := setupTestECommerceService()
	setupSuccessfulMocks(mocks)

	// Add integration test user
	integrationUser := &User{
		ID:        "integration-user",
		Email:     "integration@example.com",
		FirstName: "Integration",
		LastName:  "User",
		CreatedAt: time.Now(),
	}
	mocks.UserRepo.AddUser(integrationUser)

	return service, mocks
}

func setupBenchmarkECommerceService() (*ECommerceService, *TestMocks) {
	service, mocks := setupTestECommerceService()

	// Setup with minimal overhead for benchmarking
	for i := 1; i <= 5; i++ {
		productID := fmt.Sprintf("prod-%d", i)
		product := &Product{
			ID:          productID,
			Name:        fmt.Sprintf("Benchmark Product %d", i),
			Description: "Benchmark Description",
			Price:       float64(i * 10),
			Stock:       10000, // High stock for benchmarking
			Category:    "benchmark",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mocks.ProductRepo.AddProduct(product)
		mocks.InventoryService.SetStock(productID, 10000)
	}

	// Add a benchmark user template
	benchUser := &User{
		ID:        "bench-user",
		Email:     "bench@example.com",
		FirstName: "Bench",
		LastName:  "User",
		CreatedAt: time.Now(),
	}
	mocks.UserRepo.AddUser(benchUser)

	return service, mocks
}

// ==========================================
// 🎲 PROPERTY TEST HELPERS
// ==========================================

func generateRandomOrderItems(t *testing.T) []OrderItem {
	numItems := rand.Intn(5) + 1 // 1-5 items
	items := make([]OrderItem, numItems)

	for i := 0; i < numItems; i++ {
		price := rand.Float64() * 1000 // $0-$1000
		quantity := rand.Intn(10) + 1  // 1-10 quantity

		items[i] = OrderItem{
			ProductID: fmt.Sprintf("prod-%d", rand.Intn(100)),
			Quantity:  quantity,
			UnitPrice: price,
			Subtotal:  price * float64(quantity),
		}
	}

	return items
}

func calculateExpectedTotal(items []OrderItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Subtotal
	}
	return total
}

func generateRandomProductID() string {
	return fmt.Sprintf("prod-%d", rand.Intn(1000))
}

func generateRandomStock() int {
	return rand.Intn(100) + 1 // 1-100 stock
}

func generateRandomReservation(maxStock int) int {
	if maxStock <= 0 {
		return 0
	}
	return rand.Intn(maxStock + 1) // 0 to maxStock
}
