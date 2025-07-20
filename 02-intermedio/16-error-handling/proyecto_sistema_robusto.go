// 🏗️ Proyecto: Sistema de Procesamiento de Pedidos Robusto
// Lección 16: Error Handling Avanzado
// Demostración de patrones avanzados de manejo de errores

package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// ============================================
// 🏗️ ARQUITECTURA DEL SISTEMA DE ERRORES
// ============================================

// Errores sentinela del dominio
var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrInsufficientStock  = errors.New("insufficient stock")
	ErrPaymentFailed      = errors.New("payment processing failed")
	ErrInvalidOrderData   = errors.New("invalid order data")
	ErrServiceUnavailable = errors.New("service temporarily unavailable")
)

// Error personalizado para validación
type ValidationError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s (value: %v)",
		e.Field, e.Message, e.Value)
}

// Error de dominio de negocio
type BusinessError struct {
	Operation string    `json:"operation"`
	Reason    string    `json:"reason"`
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Err       error     `json:"-"`
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("business error in %s: %s [%s]",
		e.Operation, e.Reason, e.Code)
}

func (e BusinessError) Unwrap() error {
	return e.Err
}

// Error de infraestructura
type InfrastructureError struct {
	Service   string    `json:"service"`
	Operation string    `json:"operation"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Retryable bool      `json:"retryable"`
	Err       error     `json:"-"`
}

func (e InfrastructureError) Error() string {
	retryMsg := ""
	if e.Retryable {
		retryMsg = " (retryable)"
	}
	return fmt.Sprintf("infrastructure error in %s.%s: %s%s",
		e.Service, e.Operation, e.Message, retryMsg)
}

func (e InfrastructureError) Unwrap() error {
	return e.Err
}

func (e InfrastructureError) IsRetryable() bool {
	return e.Retryable
}

// ============================================
// 🎯 RESULT TYPE PATTERN
// ============================================

type Result[T any] struct {
	value T
	err   error
}

func Success[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

func Failure[T any](err error) Result[T] {
	var zero T
	return Result[T]{value: zero, err: err}
}

func (r Result[T]) IsSuccess() bool {
	return r.err == nil
}

func (r Result[T]) IsFailure() bool {
	return r.err != nil
}

func (r Result[T]) Value() T {
	return r.value
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

// Map transformation para Result
func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	if r.IsFailure() {
		return Failure[U](r.err)
	}
	return Success(fn(r.value))
}

// FlatMap para operaciones que pueden fallar
func FlatMap[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	if r.IsFailure() {
		return Failure[U](r.err)
	}
	return fn(r.value)
}

// ============================================
// 🔄 CIRCUIT BREAKER AVANZADO
// ============================================

type CircuitState int

const (
	CircuitClosed CircuitState = iota
	CircuitOpen
	CircuitHalfOpen
)

func (s CircuitState) String() string {
	switch s {
	case CircuitClosed:
		return "CLOSED"
	case CircuitOpen:
		return "OPEN"
	case CircuitHalfOpen:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

type CircuitBreakerConfig struct {
	FailureThreshold int
	RecoveryTimeout  time.Duration
	SuccessThreshold int
	MaxRequests      int
}

type CircuitBreaker struct {
	mu              sync.RWMutex
	state           CircuitState
	failureCount    int
	successCount    int
	requestCount    int
	lastFailureTime time.Time
	config          CircuitBreakerConfig
}

func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker {
	return &CircuitBreaker{
		state:  CircuitClosed,
		config: config,
	}
}

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open")

func (cb *CircuitBreaker) Execute(operation func() error) error {
	// Verificar si podemos ejecutar
	if !cb.canExecute() {
		return ErrCircuitBreakerOpen
	}

	// Ejecutar operación
	err := operation()

	// Actualizar estado basado en resultado
	cb.recordResult(err)

	return err
}

func (cb *CircuitBreaker) canExecute() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	switch cb.state {
	case CircuitClosed:
		return true
	case CircuitOpen:
		// Verificar si es tiempo de intentar recovery
		if time.Since(cb.lastFailureTime) > cb.config.RecoveryTimeout {
			cb.mu.RUnlock()
			cb.mu.Lock()
			cb.state = CircuitHalfOpen
			cb.requestCount = 0
			cb.mu.Unlock()
			cb.mu.RLock()
			return true
		}
		return false
	case CircuitHalfOpen:
		return cb.requestCount < cb.config.MaxRequests
	default:
		return false
	}
}

func (cb *CircuitBreaker) recordResult(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == CircuitHalfOpen {
		cb.requestCount++
	}

	if err != nil {
		cb.onFailure()
	} else {
		cb.onSuccess()
	}
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()

	if cb.state == CircuitHalfOpen || cb.failureCount >= cb.config.FailureThreshold {
		cb.state = CircuitOpen
		cb.successCount = 0
	}
}

func (cb *CircuitBreaker) onSuccess() {
	cb.failureCount = 0

	if cb.state == CircuitHalfOpen {
		cb.successCount++
		if cb.successCount >= cb.config.SuccessThreshold {
			cb.state = CircuitClosed
			cb.successCount = 0
		}
	}
}

func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

func (cb *CircuitBreaker) GetStats() map[string]interface{} {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	return map[string]interface{}{
		"state":        cb.state.String(),
		"failures":     cb.failureCount,
		"successes":    cb.successCount,
		"requests":     cb.requestCount,
		"last_failure": cb.lastFailureTime,
	}
}

// ============================================
// 📊 SISTEMA DE MÉTRICAS Y OBSERVABILIDAD
// ============================================

type ErrorMetrics struct {
	mu              sync.RWMutex
	totalErrors     int64
	errorsByType    map[string]int64
	errorsByService map[string]int64
	recentErrors    []ErrorEvent
	maxRecent       int
}

type ErrorEvent struct {
	Timestamp time.Time   `json:"timestamp"`
	ErrorType string      `json:"error_type"`
	Service   string      `json:"service"`
	Operation string      `json:"operation"`
	Message   string      `json:"message"`
	Severity  string      `json:"severity"`
	Metadata  interface{} `json:"metadata,omitempty"`
}

func NewErrorMetrics(maxRecent int) *ErrorMetrics {
	return &ErrorMetrics{
		errorsByType:    make(map[string]int64),
		errorsByService: make(map[string]int64),
		recentErrors:    make([]ErrorEvent, 0, maxRecent),
		maxRecent:       maxRecent,
	}
}

func (em *ErrorMetrics) RecordError(event ErrorEvent) {
	em.mu.Lock()
	defer em.mu.Unlock()

	em.totalErrors++
	em.errorsByType[event.ErrorType]++
	em.errorsByService[event.Service]++

	// Agregar a errores recientes
	if len(em.recentErrors) >= em.maxRecent {
		em.recentErrors = em.recentErrors[1:]
	}
	em.recentErrors = append(em.recentErrors, event)
}

func (em *ErrorMetrics) GetStats() map[string]interface{} {
	em.mu.RLock()
	defer em.mu.RUnlock()

	return map[string]interface{}{
		"total_errors":      em.totalErrors,
		"errors_by_type":    em.copyStringInt64Map(em.errorsByType),
		"errors_by_service": em.copyStringInt64Map(em.errorsByService),
		"recent_count":      len(em.recentErrors),
	}
}

func (em *ErrorMetrics) GetRecentErrors(limit int) []ErrorEvent {
	em.mu.RLock()
	defer em.mu.RUnlock()

	start := len(em.recentErrors) - limit
	if start < 0 {
		start = 0
	}

	recent := make([]ErrorEvent, len(em.recentErrors[start:]))
	copy(recent, em.recentErrors[start:])
	return recent
}

func (em *ErrorMetrics) copyStringInt64Map(m map[string]int64) map[string]int64 {
	copy := make(map[string]int64)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

// ============================================
// 🔧 RETRY MECHANISM CON BACKOFF INTELIGENTE
// ============================================

type RetryConfig struct {
	MaxAttempts     int
	InitialDelay    time.Duration
	MaxDelay        time.Duration
	BackoffFactor   float64
	RetryableErrors []error
}

type RetryableError interface {
	IsRetryable() bool
}

func RetryWithBackoff(operation func() error, config RetryConfig) error {
	var lastErr error
	delay := config.InitialDelay

	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		err := operation()
		if err == nil {
			return nil // Éxito
		}

		lastErr = err

		// Verificar si el error es reintentable
		if !isRetryableError(err, config.RetryableErrors) {
			return fmt.Errorf("non-retryable error: %w", err)
		}

		if attempt == config.MaxAttempts-1 {
			break // Último intento
		}

		// Esperar antes del siguiente intento
		log.Printf("Attempt %d failed: %v. Retrying in %v...",
			attempt+1, err, delay)
		time.Sleep(delay)

		// Calcular siguiente delay con backoff exponencial
		delay = time.Duration(float64(delay) * config.BackoffFactor)
		if delay > config.MaxDelay {
			delay = config.MaxDelay
		}
	}

	return fmt.Errorf("operation failed after %d attempts: %w",
		config.MaxAttempts, lastErr)
}

func isRetryableError(err error, retryableErrors []error) bool {
	// Verificar si implementa RetryableError
	var retryable RetryableError
	if errors.As(err, &retryable) {
		return retryable.IsRetryable()
	}

	// Verificar contra lista de errores reintentables
	for _, retryableErr := range retryableErrors {
		if errors.Is(err, retryableErr) {
			return true
		}
	}

	return false
}

// ============================================
// 🏢 MODELOS DE DOMINIO
// ============================================

type Order struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type PaymentRequest struct {
	OrderID    string  `json:"order_id"`
	Amount     float64 `json:"amount"`
	Method     string  `json:"method"`
	CustomerID string  `json:"customer_id"`
}

type PaymentResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
	ProcessedAt   time.Time `json:"processed_at"`
}

// ============================================
// 🏗️ SERVICIOS CON ERROR HANDLING ROBUSTO
// ============================================

// Servicio de validación
type ValidationService struct {
	metrics *ErrorMetrics
}

func NewValidationService(metrics *ErrorMetrics) *ValidationService {
	return &ValidationService{metrics: metrics}
}

func (vs *ValidationService) ValidateOrder(order Order) Result[Order] {
	var validationErrors []error

	// Validar ID
	if order.ID == "" {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "id",
			Value:   order.ID,
			Message: "cannot be empty",
			Code:    "REQUIRED_FIELD",
		})
	}

	// Validar CustomerID
	if order.CustomerID == "" {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "customer_id",
			Value:   order.CustomerID,
			Message: "cannot be empty",
			Code:    "REQUIRED_FIELD",
		})
	}

	// Validar Items
	if len(order.Items) == 0 {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "items",
			Value:   len(order.Items),
			Message: "order must contain at least one item",
			Code:    "MIN_ITEMS",
		})
	}

	// Validar TotalAmount
	if order.TotalAmount <= 0 {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "total_amount",
			Value:   order.TotalAmount,
			Message: "must be greater than 0",
			Code:    "INVALID_AMOUNT",
		})
	}

	if len(validationErrors) > 0 {
		err := fmt.Errorf("order validation failed: %v", validationErrors)

		vs.metrics.RecordError(ErrorEvent{
			Timestamp: time.Now(),
			ErrorType: "validation_error",
			Service:   "validation",
			Operation: "validate_order",
			Message:   err.Error(),
			Severity:  "warning",
			Metadata:  map[string]interface{}{"order_id": order.ID},
		})

		return Failure[Order](err)
	}

	return Success(order)
}

// Servicio de inventario
type InventoryService struct {
	circuitBreaker *CircuitBreaker
	metrics        *ErrorMetrics
	stockData      map[string]int // Simulación de stock
	mu             sync.RWMutex
}

func NewInventoryService(metrics *ErrorMetrics) *InventoryService {
	config := CircuitBreakerConfig{
		FailureThreshold: 3,
		RecoveryTimeout:  5 * time.Second,
		SuccessThreshold: 2,
		MaxRequests:      5,
	}

	// Stock inicial simulado
	stockData := map[string]int{
		"product-1": 100,
		"product-2": 50,
		"product-3": 0, // Sin stock
		"product-4": 25,
	}

	return &InventoryService{
		circuitBreaker: NewCircuitBreaker(config),
		metrics:        metrics,
		stockData:      stockData,
	}
}

func (is *InventoryService) CheckStock(items []Item) Result[bool] {
	operation := func() error {
		return is.checkStockInternal(items)
	}

	err := is.circuitBreaker.Execute(operation)
	if err != nil {
		if errors.Is(err, ErrCircuitBreakerOpen) {
			infraErr := InfrastructureError{
				Service:   "inventory",
				Operation: "check_stock",
				Message:   "service circuit breaker is open",
				Timestamp: time.Now(),
				Retryable: true,
				Err:       err,
			}

			is.metrics.RecordError(ErrorEvent{
				Timestamp: time.Now(),
				ErrorType: "circuit_breaker",
				Service:   "inventory",
				Operation: "check_stock",
				Message:   infraErr.Error(),
				Severity:  "error",
				Metadata:  map[string]interface{}{"circuit_state": "open"},
			})

			return Failure[bool](infraErr)
		}

		return Failure[bool](err)
	}

	return Success(true)
}

func (is *InventoryService) checkStockInternal(items []Item) error {
	is.mu.RLock()
	defer is.mu.RUnlock()

	// Simular falla ocasional del servicio
	if time.Now().UnixNano()%10 < 3 { // 30% de probabilidad de fallo
		err := InfrastructureError{
			Service:   "inventory",
			Operation: "check_stock",
			Message:   "database connection timeout",
			Timestamp: time.Now(),
			Retryable: true,
			Err:       ErrServiceUnavailable,
		}

		is.metrics.RecordError(ErrorEvent{
			Timestamp: time.Now(),
			ErrorType: "infrastructure_error",
			Service:   "inventory",
			Operation: "check_stock",
			Message:   err.Error(),
			Severity:  "error",
		})

		return err
	}

	// Verificar stock para cada item
	for _, item := range items {
		available, exists := is.stockData[item.ProductID]
		if !exists {
			err := BusinessError{
				Operation: "check_stock",
				Reason:    fmt.Sprintf("product %s not found", item.ProductID),
				Code:      "PRODUCT_NOT_FOUND",
				Timestamp: time.Now(),
				Err:       ErrOrderNotFound,
			}

			is.metrics.RecordError(ErrorEvent{
				Timestamp: time.Now(),
				ErrorType: "business_error",
				Service:   "inventory",
				Operation: "check_stock",
				Message:   err.Error(),
				Severity:  "warning",
				Metadata:  map[string]interface{}{"product_id": item.ProductID},
			})

			return err
		}

		if available < item.Quantity {
			err := BusinessError{
				Operation: "check_stock",
				Reason: fmt.Sprintf("insufficient stock for product %s: available %d, requested %d",
					item.ProductID, available, item.Quantity),
				Code:      "INSUFFICIENT_STOCK",
				Timestamp: time.Now(),
				Err:       ErrInsufficientStock,
			}

			is.metrics.RecordError(ErrorEvent{
				Timestamp: time.Now(),
				ErrorType: "business_error",
				Service:   "inventory",
				Operation: "check_stock",
				Message:   err.Error(),
				Severity:  "warning",
				Metadata: map[string]interface{}{
					"product_id": item.ProductID,
					"available":  available,
					"requested":  item.Quantity,
				},
			})

			return err
		}
	}

	return nil
}

// Servicio de pagos
type PaymentService struct {
	circuitBreaker *CircuitBreaker
	metrics        *ErrorMetrics
}

func NewPaymentService(metrics *ErrorMetrics) *PaymentService {
	config := CircuitBreakerConfig{
		FailureThreshold: 5,
		RecoveryTimeout:  10 * time.Second,
		SuccessThreshold: 3,
		MaxRequests:      10,
	}

	return &PaymentService{
		circuitBreaker: NewCircuitBreaker(config),
		metrics:        metrics,
	}
}

func (ps *PaymentService) ProcessPayment(request PaymentRequest) Result[PaymentResponse] {
	retryConfig := RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Second,
		MaxDelay:      5 * time.Second,
		BackoffFactor: 2.0,
		RetryableErrors: []error{
			ErrServiceUnavailable,
		},
	}

	var response PaymentResponse
	var finalErr error

	operation := func() error {
		return ps.circuitBreaker.Execute(func() error {
			var err error
			response, err = ps.processPaymentInternal(request)
			return err
		})
	}

	finalErr = RetryWithBackoff(operation, retryConfig)

	if finalErr != nil {
		return Failure[PaymentResponse](finalErr)
	}

	return Success(response)
}

func (ps *PaymentService) processPaymentInternal(request PaymentRequest) (PaymentResponse, error) {
	// Simular diferentes tipos de errores
	switch request.Method {
	case "invalid_card":
		err := BusinessError{
			Operation: "process_payment",
			Reason:    "invalid credit card information",
			Code:      "INVALID_PAYMENT_METHOD",
			Timestamp: time.Now(),
			Err:       ErrPaymentFailed,
		}

		ps.metrics.RecordError(ErrorEvent{
			Timestamp: time.Now(),
			ErrorType: "business_error",
			Service:   "payment",
			Operation: "process_payment",
			Message:   err.Error(),
			Severity:  "warning",
			Metadata:  map[string]interface{}{"order_id": request.OrderID},
		})

		return PaymentResponse{}, err

	case "network_error":
		err := InfrastructureError{
			Service:   "payment",
			Operation: "process_payment",
			Message:   "payment gateway connection failed",
			Timestamp: time.Now(),
			Retryable: true,
			Err:       ErrServiceUnavailable,
		}

		ps.metrics.RecordError(ErrorEvent{
			Timestamp: time.Now(),
			ErrorType: "infrastructure_error",
			Service:   "payment",
			Operation: "process_payment",
			Message:   err.Error(),
			Severity:  "error",
			Metadata:  map[string]interface{}{"order_id": request.OrderID},
		})

		return PaymentResponse{}, err
	}

	// Pago exitoso
	response := PaymentResponse{
		TransactionID: fmt.Sprintf("txn_%d", time.Now().Unix()),
		Status:        "completed",
		ProcessedAt:   time.Now(),
	}

	return response, nil
}

// ============================================
// 🎯 ORQUESTADOR PRINCIPAL
// ============================================

type OrderProcessor struct {
	validationService *ValidationService
	inventoryService  *InventoryService
	paymentService    *PaymentService
	metrics           *ErrorMetrics
}

func NewOrderProcessor() *OrderProcessor {
	metrics := NewErrorMetrics(100)

	return &OrderProcessor{
		validationService: NewValidationService(metrics),
		inventoryService:  NewInventoryService(metrics),
		paymentService:    NewPaymentService(metrics),
		metrics:           metrics,
	}
}

func (op *OrderProcessor) ProcessOrder(order Order) Result[string] {
	// 1. Validación
	validationResult := op.validationService.ValidateOrder(order)
	if validationResult.IsFailure() {
		return Failure[string](fmt.Errorf("order validation failed: %w", validationResult.Error()))
	}

	// 2. Verificar stock usando FlatMap
	stockResult := FlatMap(validationResult, func(validOrder Order) Result[bool] {
		return op.inventoryService.CheckStock(validOrder.Items)
	})

	if stockResult.IsFailure() {
		return Failure[string](fmt.Errorf("stock check failed: %w", stockResult.Error()))
	}

	// 3. Procesar pago
	paymentRequest := PaymentRequest{
		OrderID:    order.ID,
		Amount:     order.TotalAmount,
		Method:     "credit_card", // Simulado
		CustomerID: order.CustomerID,
	}

	paymentResult := op.paymentService.ProcessPayment(paymentRequest)
	if paymentResult.IsFailure() {
		return Failure[string](fmt.Errorf("payment processing failed: %w", paymentResult.Error()))
	}

	// Éxito completo
	successMessage := fmt.Sprintf("Order %s processed successfully. Transaction: %s",
		order.ID, paymentResult.Value().TransactionID)

	return Success(successMessage)
}

func (op *OrderProcessor) GetMetrics() map[string]interface{} {
	return op.metrics.GetStats()
}

func (op *OrderProcessor) GetRecentErrors(limit int) []ErrorEvent {
	return op.metrics.GetRecentErrors(limit)
}

// ============================================
// 🧪 DEMOSTRACIÓN DEL SISTEMA
// ============================================

func main() {
	fmt.Println("🏗️ SISTEMA DE PROCESAMIENTO DE PEDIDOS ROBUSTO")
	fmt.Println("============================================")
	fmt.Println("Demostrando patrones avanzados de error handling")

	processor := NewOrderProcessor()

	// Casos de prueba con diferentes tipos de errores
	testCases := []struct {
		name  string
		order Order
	}{
		{
			name: "Pedido Válido",
			order: Order{
				ID:         "order-001",
				CustomerID: "customer-123",
				Items: []Item{
					{ProductID: "product-1", Quantity: 2, Price: 25.99},
					{ProductID: "product-2", Quantity: 1, Price: 15.50},
				},
				TotalAmount: 67.48,
				Status:      "pending",
				CreatedAt:   time.Now(),
			},
		},
		{
			name: "Pedido con Datos Inválidos",
			order: Order{
				ID:          "", // Campo requerido vacío
				CustomerID:  "customer-456",
				Items:       []Item{}, // Sin items
				TotalAmount: -10.0,    // Monto inválido
				CreatedAt:   time.Now(),
			},
		},
		{
			name: "Pedido con Stock Insuficiente",
			order: Order{
				ID:         "order-003",
				CustomerID: "customer-789",
				Items: []Item{
					{ProductID: "product-3", Quantity: 5, Price: 30.00}, // Sin stock
				},
				TotalAmount: 150.00,
				CreatedAt:   time.Now(),
			},
		},
		{
			name: "Pedido con Error de Pago",
			order: Order{
				ID:         "order-004",
				CustomerID: "customer-invalid",
				Items: []Item{
					{ProductID: "product-1", Quantity: 1, Price: 25.99},
				},
				TotalAmount: 25.99,
				CreatedAt:   time.Now(),
			},
		},
	}

	fmt.Println("\n🔄 Procesando pedidos de prueba...")

	// Procesar cada caso de prueba
	for i, testCase := range testCases {
		fmt.Printf("\n--- Caso %d: %s ---\n", i+1, testCase.name)

		result := processor.ProcessOrder(testCase.order)

		if result.IsSuccess() {
			fmt.Printf("✅ %s\n", result.Value())
		} else {
			fmt.Printf("❌ Error: %v\n", result.Error())

			// Analizar tipo de error
			err := result.Error()

			var validationErr ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("   🔍 Tipo: Error de validación (Campo: %s)\n", validationErr.Field)
			}

			var businessErr BusinessError
			if errors.As(err, &businessErr) {
				fmt.Printf("   🔍 Tipo: Error de negocio (Código: %s)\n", businessErr.Code)
			}

			var infraErr InfrastructureError
			if errors.As(err, &infraErr) {
				fmt.Printf("   🔍 Tipo: Error de infraestructura (Reintentable: %t)\n", infraErr.Retryable)
			}
		}
	}

	// Mostrar métricas del sistema
	fmt.Println("\n📊 MÉTRICAS DEL SISTEMA")
	fmt.Println("========================")

	metrics := processor.GetMetrics()

	if totalErrors, ok := metrics["total_errors"].(int64); ok {
		fmt.Printf("Total de errores: %d\n", totalErrors)
	}

	if errorsByType, ok := metrics["errors_by_type"].(map[string]int64); ok {
		fmt.Println("\nErrores por tipo:")
		for errorType, count := range errorsByType {
			fmt.Printf("  %s: %d\n", errorType, count)
		}
	}

	if errorsByService, ok := metrics["errors_by_service"].(map[string]int64); ok {
		fmt.Println("\nErrores por servicio:")
		for service, count := range errorsByService {
			fmt.Printf("  %s: %d\n", service, count)
		}
	}

	// Mostrar errores recientes
	fmt.Println("\n🕒 ERRORES RECIENTES")
	fmt.Println("==================")

	recentErrors := processor.GetRecentErrors(5)
	for i, errorEvent := range recentErrors {
		fmt.Printf("%d. [%s] %s.%s: %s (%s)\n",
			i+1,
			errorEvent.Timestamp.Format("15:04:05"),
			errorEvent.Service,
			errorEvent.Operation,
			errorEvent.Message,
			errorEvent.Severity)
	}

	fmt.Println("\n🎉 DEMOSTRACIÓN COMPLETADA")
	fmt.Println("==========================")
	fmt.Println("✅ Patrones implementados:")
	fmt.Println("   - Errores personalizados con contexto")
	fmt.Println("   - Error wrapping y unwrapping")
	fmt.Println("   - Result type pattern")
	fmt.Println("   - Circuit breaker avanzado")
	fmt.Println("   - Retry con backoff exponencial")
	fmt.Println("   - Métricas y observabilidad")
	fmt.Println("   - Error handling estructurado")
	fmt.Println("\n💡 Este sistema demuestra como crear aplicaciones robustas")
	fmt.Println("   que manejan errores de manera elegante y proporcionan")
	fmt.Println("   visibilidad completa sobre el comportamiento del sistema.")
}
