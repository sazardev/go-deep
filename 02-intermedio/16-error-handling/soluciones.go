// 🚨 Soluciones: Error Handling Avanzado
// Lección 16: Robustez y Resilencia en Go
// Ejecutar con: go run soluciones.go

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ========================================
// Ejercicio 1: Error Básico Personalizado
// ========================================

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with value '%v': %s",
		e.Field, e.Value, e.Message)
}

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "must be between 0 and 150",
		}
	}
	return nil
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Error Personalizado ===")

	// Probar edad válida
	err := validateAge(25)
	if err == nil {
		fmt.Println("✅ Edad 25: válida")
	}

	// Probar edad inválida
	err = validateAge(-5)
	if err != nil {
		fmt.Printf("❌ Edad -5: %v\n", err)
	}

	err = validateAge(200)
	if err != nil {
		fmt.Printf("❌ Edad 200: %v\n", err)
	}
}

// ==========================================
// Ejercicio 2: Error Wrapping con Contexto
// ==========================================

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrConnectionFailed = errors.New("database connection failed")
)

func queryDatabase(userID string) (string, error) {
	if userID == "" {
		return "", ValidationError{
			Field:   "userID",
			Value:   userID,
			Message: "cannot be empty",
		}
	}

	if userID == "999" {
		return "", fmt.Errorf("user lookup failed: %w", ErrUserNotFound)
	}

	if userID == "error" {
		return "", fmt.Errorf("database query failed: %w", ErrConnectionFailed)
	}

	return fmt.Sprintf("User data for ID: %s", userID), nil
}

func getUserService(userID string) (string, error) {
	data, err := queryDatabase(userID)
	if err != nil {
		return "", fmt.Errorf("user service failed for ID %s: %w", userID, err)
	}
	return data, nil
}

func ejercicio2() {
	fmt.Println("\n=== Ejercicio 2: Error Wrapping ===")

	testCases := []string{"123", "", "999", "error"}

	for _, userID := range testCases {
		data, err := getUserService(userID)
		if err != nil {
			fmt.Printf("❌ UserID '%s': %v\n", userID, err)

			// Mostrar error original
			if originalErr := errors.Unwrap(err); originalErr != nil {
				fmt.Printf("   📍 Original error: %v\n", originalErr)
			}

			// Verificar tipos específicos
			if errors.Is(err, ErrUserNotFound) {
				fmt.Println("   🔍 Detected: User not found error")
			}
			if errors.Is(err, ErrConnectionFailed) {
				fmt.Println("   🔍 Detected: Connection error")
			}
		} else {
			fmt.Printf("✅ UserID '%s': %s\n", userID, data)
		}
	}
}

// ========================================
// Ejercicio 3: Error Accumulator Pattern
// ========================================

type ErrorAccumulator struct {
	errorList []error
}

func (ea *ErrorAccumulator) Add(err error) {
	if err != nil {
		ea.errorList = append(ea.errorList, err)
	}
}

func (ea *ErrorAccumulator) HasErrors() bool {
	return len(ea.errorList) > 0
}

func (ea *ErrorAccumulator) Error() string {
	if len(ea.errorList) == 0 {
		return ""
	}

	var messages []string
	for _, err := range ea.errorList {
		messages = append(messages, err.Error())
	}
	return fmt.Sprintf("multiple validation errors: [%s]", strings.Join(messages, "; "))
}

func (ea *ErrorAccumulator) Errors() []error {
	return ea.errorList
}

func validateUser(name, email string, age int) error {
	var errAccumulator ErrorAccumulator

	// Validar nombre
	if len(name) < 2 {
		errAccumulator.Add(ValidationError{
			Field:   "name",
			Value:   name,
			Message: "must have at least 2 characters",
		})
	}

	// Validar email
	if !strings.Contains(email, "@") {
		errAccumulator.Add(ValidationError{
			Field:   "email",
			Value:   email,
			Message: "must contain @ symbol",
		})
	}

	// Validar edad
	if ageErr := validateAge(age); ageErr != nil {
		errAccumulator.Add(ageErr)
	}

	if errAccumulator.HasErrors() {
		return &errAccumulator
	}
	return nil
}

func ejercicio3() {
	fmt.Println("\n=== Ejercicio 3: Error Accumulator ===")

	// Caso con múltiples errores
	err := validateUser("A", "not-email", -5)
	if err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)

		if errList, ok := err.(*ErrorAccumulator); ok {
			fmt.Println("📋 Individual errors:")
			for i, e := range errList.Errors() {
				fmt.Printf("   %d. %v\n", i+1, e)
			}
		}
	}

	// Caso exitoso
	err = validateUser("John Doe", "john@example.com", 30)
	if err == nil {
		fmt.Println("✅ User validation successful")
	}
}

// ===================================
// Ejercicio 4: Result Type Pattern
// ===================================

type Result[T any] struct {
	Value T
	Error error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value, Error: nil}
}

func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{Value: zero, Error: err}
}

func (r Result[T]) IsOk() bool {
	return r.Error == nil
}

func (r Result[T]) IsErr() bool {
	return r.Error != nil
}

func (r Result[T]) Unwrap() (T, error) {
	return r.Value, r.Error
}

func parseNumber(s string) Result[int] {
	value, err := strconv.Atoi(s)
	if err != nil {
		return Err[int](fmt.Errorf("failed to parse '%s' as number: %w", s, err))
	}
	return Ok(value)
}

func ejercicio4() {
	fmt.Println("\n=== Ejercicio 4: Result Type Pattern ===")

	testInputs := []string{"123", "abc", "456", "not-a-number"}

	for _, input := range testInputs {
		result := parseNumber(input)

		if result.IsOk() {
			fmt.Printf("✅ '%s' → %d\n", input, result.Value)
		} else {
			fmt.Printf("❌ '%s' → %v\n", input, result.Error)
		}

		// También podemos usar Unwrap
		value, err := result.Unwrap()
		if err == nil {
			fmt.Printf("   📦 Unwrapped value: %d\n", value)
		}
	}
}

// =====================================
// Ejercicio 5: Circuit Breaker Simple
// =====================================

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

func (s CircuitState) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

type CircuitBreaker struct {
	mu           sync.Mutex
	state        CircuitState
	failureCount int
	threshold    int
	successCount int
}

func NewCircuitBreaker(threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		state:     StateClosed,
		threshold: threshold,
	}
}

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open")

func (cb *CircuitBreaker) Execute(operation func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	// Si está abierto, no ejecutar
	if cb.state == StateOpen {
		return ErrCircuitBreakerOpen
	}

	// Ejecutar operación
	err := operation()

	if err != nil {
		cb.onFailure()
	} else {
		cb.onSuccess()
	}

	return err
}

func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	if cb.failureCount >= cb.threshold {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) onSuccess() {
	cb.failureCount = 0
	if cb.state == StateHalfOpen {
		cb.successCount++
		if cb.successCount >= 2 {
			cb.state = StateClosed
			cb.successCount = 0
		}
	}
}

func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}

// Simular operación que falla frecuentemente
var operationCounter int

func unreliableOperation() error {
	operationCounter++
	// Fallar 70% del tiempo
	if operationCounter%10 < 7 {
		return errors.New("operation failed")
	}
	return nil
}

func ejercicio5() {
	fmt.Println("\n=== Ejercicio 5: Circuit Breaker ===")

	cb := NewCircuitBreaker(3)

	// Probar múltiples operaciones
	for i := 0; i < 10; i++ {
		err := cb.Execute(unreliableOperation)
		state := cb.GetState()

		if err != nil {
			if errors.Is(err, ErrCircuitBreakerOpen) {
				fmt.Printf("Attempt %d: ⛔ Circuit Breaker OPEN\n", i+1)
			} else {
				fmt.Printf("Attempt %d: ❌ Operation failed [%s]\n", i+1, state)
			}
		} else {
			fmt.Printf("Attempt %d: ✅ Operation succeeded [%s]\n", i+1, state)
		}
	}
}

// ===================================
// Ejercicio 6: Error Middleware HTTP
// ===================================

type HTTPError struct {
	Code       string
	Message    string
	StatusCode int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewBadRequest(message string) HTTPError {
	return HTTPError{
		Code:       "BAD_REQUEST",
		Message:    message,
		StatusCode: 400,
	}
}

func NewNotFound(resource string) HTTPError {
	return HTTPError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s not found", resource),
		StatusCode: 404,
	}
}

func NewInternalError(err error) HTTPError {
	return HTTPError{
		Code:       "INTERNAL_ERROR",
		Message:    "Internal server error occurred",
		StatusCode: 500,
	}
}

func simulateHTTPHandler(endpoint string) error {
	switch endpoint {
	case "/invalid":
		return NewBadRequest("invalid request parameters")
	case "/missing":
		return NewNotFound("Resource")
	case "/error":
		return NewInternalError(errors.New("database connection failed"))
	default:
		return nil // Éxito
	}
}

func ejercicio6() {
	fmt.Println("\n=== Ejercicio 6: HTTP Error Handling ===")

	endpoints := []string{"/valid", "/invalid", "/missing", "/error"}

	for _, endpoint := range endpoints {
		err := simulateHTTPHandler(endpoint)

		if err != nil {
			var httpErr HTTPError
			if errors.As(err, &httpErr) {
				fmt.Printf("❌ %s → [%d] %s\n", endpoint, httpErr.StatusCode, httpErr.Error())
			} else {
				fmt.Printf("❌ %s → Unknown error: %v\n", endpoint, err)
			}
		} else {
			fmt.Printf("✅ %s → Success\n", endpoint)
		}
	}
}

// =====================================
// Ejercicio 7: Error Metrics y Logging
// =====================================

type ErrorMetrics struct {
	mu          sync.RWMutex
	errorCounts map[string]int
}

func NewErrorMetrics() *ErrorMetrics {
	return &ErrorMetrics{
		errorCounts: make(map[string]int),
	}
}

func (em *ErrorMetrics) RecordError(errorType string) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.errorCounts[errorType]++
}

func (em *ErrorMetrics) GetStats() map[string]int {
	em.mu.RLock()
	defer em.mu.RUnlock()

	// Crear copia para evitar race conditions
	stats := make(map[string]int)
	for k, v := range em.errorCounts {
		stats[k] = v
	}
	return stats
}

type ErrorLogger struct {
	metrics *ErrorMetrics
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		metrics: NewErrorMetrics(),
	}
}

func (el *ErrorLogger) LogError(err error, service string) {
	errorType := "unknown"

	// Determinar tipo de error
	var httpErr HTTPError
	if errors.As(err, &httpErr) {
		errorType = httpErr.Code
	} else if errors.Is(err, ErrUserNotFound) {
		errorType = "user_not_found"
	} else if errors.Is(err, ErrConnectionFailed) {
		errorType = "connection_failed"
	}

	// Registrar en métricas
	el.metrics.RecordError(errorType)

	// Log estructurado
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] ERROR %s/%s: %v\n", timestamp, service, errorType, err)
}

func (el *ErrorLogger) PrintStats() {
	stats := el.metrics.GetStats()
	fmt.Println("\n📊 Error Statistics:")
	for errorType, count := range stats {
		fmt.Printf("  %s: %d\n", errorType, count)
	}
}

func ejercicio7() {
	fmt.Println("\n=== Ejercicio 7: Error Metrics ===")

	logger := NewErrorLogger()

	// Simular varios errores
	logger.LogError(NewBadRequest("invalid input"), "user-service")
	logger.LogError(NewNotFound("User"), "user-service")
	logger.LogError(NewInternalError(errors.New("db error")), "order-service")
	logger.LogError(ErrUserNotFound, "auth-service")
	logger.LogError(NewBadRequest("missing field"), "user-service")

	logger.PrintStats()
}

// =======================================
// Ejercicio 8: Retry con Backoff Exponencial
// =======================================

func retryWithBackoff(operation func() error, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := operation()
		if err == nil {
			return nil // Éxito
		}

		lastErr = err

		if attempt == maxRetries-1 {
			break // No más reintentos
		}

		// Backoff exponencial: 1s, 2s, 4s, 8s...
		delay := time.Duration(1<<attempt) * time.Second
		fmt.Printf("   ⏳ Attempt %d failed, retrying in %v...\n", attempt+1, delay)
		time.Sleep(delay)
	}

	return fmt.Errorf("operation failed after %d attempts: %w", maxRetries, lastErr)
}

var attemptCount int

func flakeyOperation() error {
	attemptCount++
	fmt.Printf("   🔄 Attempt %d...\n", attemptCount)

	if attemptCount < 3 {
		return errors.New("temporary failure")
	}
	return nil // Éxito después del 3er intento
}

func ejercicio8() {
	fmt.Println("\n=== Ejercicio 8: Retry con Backoff ===")

	attemptCount = 0 // Reset contador

	err := retryWithBackoff(flakeyOperation, 5)
	if err != nil {
		fmt.Printf("❌ Final result: %v\n", err)
	} else {
		fmt.Printf("✅ Operation succeeded after %d attempts\n", attemptCount)
	}
}

// ================================
// Ejercicio 9: Error Testing Helper
// ================================

func assertErrorType[T error](t interface{}, err error, expectedType T) bool {
	if err == nil {
		fmt.Printf("❌ Expected error of type %T, but got nil\n", expectedType)
		return false
	}

	var target T
	if errors.As(err, &target) {
		fmt.Printf("✅ Error has expected type %T\n", expectedType)
		return true
	}

	fmt.Printf("❌ Expected error type %T, got %T\n", expectedType, err)
	return false
}

func assertErrorIs(t interface{}, err error, target error) bool {
	if errors.Is(err, target) {
		fmt.Printf("✅ Error matches target error: %v\n", target)
		return true
	}

	fmt.Printf("❌ Error does not match target. Got: %v, Expected: %v\n", err, target)
	return false
}

func processInput(input string) error {
	if input == "" {
		return NewBadRequest("input cannot be empty")
	}
	if input == "not-found" {
		return NewNotFound("resource")
	}
	return nil
}

func ejercicio9() {
	fmt.Println("\n=== Ejercicio 9: Error Testing ===")

	// Test caso 1: input vacío
	err1 := processInput("")
	assertErrorType(nil, err1, HTTPError{})

	// Test caso 2: not-found
	err2 := processInput("not-found")
	assertErrorType(nil, err2, HTTPError{})

	// Test caso 3: input válido
	err3 := processInput("valid")
	if err3 == nil {
		fmt.Println("✅ No error for valid input")
	}

	// Test verificación de error específico
	if httpErr, ok := err1.(HTTPError); ok {
		if httpErr.Code == "BAD_REQUEST" {
			fmt.Println("✅ Correct error code for empty input")
		}
	}
}

// ===================================
// Ejercicio 10: Sistema de Error Completo
// ===================================

type RobustService struct {
	circuitBreaker *CircuitBreaker
	logger         *ErrorLogger
	retryCount     int
}

func NewRobustService() *RobustService {
	return &RobustService{
		circuitBreaker: NewCircuitBreaker(3),
		logger:         NewErrorLogger(),
		retryCount:     3,
	}
}

var requestCounter int

func simulateBusinessLogic(requestID string) error {
	requestCounter++

	// Simular diferentes comportamientos según el requestID
	switch requestID {
	case "fail":
		return errors.New("business logic failure")
	case "timeout":
		return errors.New("operation timeout")
	default:
		// 30% de probabilidad de éxito para otros IDs
		if requestCounter%10 < 3 {
			return nil
		}
		return errors.New("random failure")
	}
}

func (rs *RobustService) ProcessRequest(requestID string) error {
	fmt.Printf("🔄 Processing request: %s\n", requestID)

	// Usar retry con circuit breaker
	err := retryWithBackoff(func() error {
		return rs.circuitBreaker.Execute(func() error {
			return simulateBusinessLogic(requestID)
		})
	}, rs.retryCount)

	if err != nil {
		// Loggear error con métricas
		rs.logger.LogError(err, "robust-service")
		return fmt.Errorf("request %s failed: %w", requestID, err)
	}

	fmt.Printf("✅ Request %s processed successfully\n", requestID)
	return nil
}

func ejercicio10() {
	fmt.Println("\n=== Ejercicio 10: Sistema Robusto Completo ===")

	service := NewRobustService()
	requests := []string{"success", "fail", "timeout", "random1", "random2", "success"}

	for i, requestID := range requests {
		fmt.Printf("\n--- Request %d: %s ---\n", i+1, requestID)
		err := service.ProcessRequest(requestID)

		state := service.circuitBreaker.GetState()
		fmt.Printf("Circuit Breaker State: %s\n", state)

		if err != nil {
			fmt.Printf("❌ Result: %v\n", err)
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Mostrar estadísticas finales
	fmt.Println("\n=== Final Statistics ===")
	service.logger.PrintStats()
}

// ===============================
// Función principal de soluciones
// ===============================

func main() {
	fmt.Println("🚨 SOLUCIONES: ERROR HANDLING AVANZADO")
	fmt.Println("=====================================")

	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()
	ejercicio9()
	ejercicio10()

	fmt.Println("\n🎉 ¡Todos los ejercicios completados!")
	fmt.Println("💡 Has dominado el error handling avanzado en Go")
	fmt.Println("🔥 Ahora puedes crear sistemas robustos y resilientes")
}
