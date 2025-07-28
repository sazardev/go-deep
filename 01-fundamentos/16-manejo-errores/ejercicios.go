// 🚨 Lección 16: Ejercicios de Manejo de Errores
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚨 === LECCIÓN 16: MANEJO DE ERRORES ===")
	fmt.Println()

	// Ejecutar todos los ejercicios
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
}

// ====================================
// 💡 EJERCICIO 1: Errores Básicos
// ====================================

func ejercicio1() {
	fmt.Println("💡 Ejercicio 1: Creando Errores Básicos")
	fmt.Println("=====================================")

	// Función con diferentes tipos de errores
	testDivision := func(a, b float64) {
		result, err := divide(a, b)
		if err != nil {
			fmt.Printf("Error dividiendo %.2f / %.2f: %v\n", a, b, err)
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
		}
	}

	testDivision(10, 2)   // OK
	testDivision(10, 0)   // Error: división por cero
	testDivision(-5, 2.5) // OK
	testDivision(0, 0)    // Error: división por cero

	fmt.Println()

	// Función con validación múltiple
	testValidateUser := func(name, email string, age int) {
		err := validateUser(name, email, age)
		if err != nil {
			fmt.Printf("Validación falló para %s: %v\n", name, err)
		} else {
			fmt.Printf("Usuario %s válido ✅\n", name)
		}
	}

	testValidateUser("Juan", "juan@email.com", 25)    // OK
	testValidateUser("", "juan@email.com", 25)        // Error: nombre vacío
	testValidateUser("Ana", "email-inválido", 30)     // Error: email inválido
	testValidateUser("Pedro", "pedro@email.com", -5)  // Error: edad negativa
	testValidateUser("María", "maria@email.com", 200) // Error: edad muy alta

	fmt.Println()
}

// divide números con manejo de error básico
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// validateUser valida datos de usuario
func validateUser(name, email string, age int) error {
	if name == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	if email == "" {
		return errors.New("el email no puede estar vacío")
	}

	if !strings.Contains(email, "@") {
		return fmt.Errorf("formato de email inválido: %s", email)
	}

	if age < 0 {
		return fmt.Errorf("la edad no puede ser negativa: %d", age)
	}

	if age > 150 {
		return fmt.Errorf("edad demasiado alta: %d (máximo: 150)", age)
	}

	return nil
}

// ====================================
// 🎭 EJERCICIO 2: Errores Personalizados
// ====================================

func ejercicio2() {
	fmt.Println("🎭 Ejercicio 2: Errores Personalizados")
	fmt.Println("=====================================")

	// Test de ValidationError personalizado
	user1 := User{Name: "", Email: "juan@email.com", Age: 25}
	user2 := User{Name: "Ana", Email: "email-sin-arroba", Age: 30}
	user3 := User{Name: "Pedro", Email: "pedro@email.com", Age: -10}
	user4 := User{Name: "María", Email: "maria@email.com", Age: 28}

	users := []User{user1, user2, user3, user4}

	for i, user := range users {
		fmt.Printf("Validando usuario %d: %+v\n", i+1, user)
		err := validateUserAdvanced(user)
		if err != nil {
			// Verificar si es ValidationError
			var valErr *ValidationError
			if errors.As(err, &valErr) {
				fmt.Printf("  ❌ Error de validación en campo '%s': %s\n", valErr.Field, valErr.Message)
				fmt.Printf("     Valor inválido: %v\n", valErr.Value)
			} else {
				fmt.Printf("  ❌ Error general: %v\n", err)
			}
		} else {
			fmt.Printf("  ✅ Usuario válido\n")
		}
		fmt.Println()
	}

	// Test de APIError
	fmt.Println("Testing API errors:")
	apiErrors := []APIError{
		{StatusCode: 404, Message: "User not found", Retryable: false},
		{StatusCode: 500, Message: "Internal server error", Retryable: true},
		{StatusCode: 429, Message: "Rate limit exceeded", Retryable: true},
		{StatusCode: 400, Message: "Bad request", Retryable: false},
	}

	for _, apiErr := range apiErrors {
		fmt.Printf("API Error: %v (Retryable: %t)\n", apiErr, apiErr.IsRetryable())
	}

	fmt.Println()
}

// User representa un usuario
type User struct {
	Name  string
	Email string
	Age   int
}

// ValidationError error personalizado para validación
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validación falló en campo '%s' con valor '%v': %s",
		e.Field, e.Value, e.Message)
}

// APIError error personalizado para APIs
type APIError struct {
	StatusCode int
	Message    string
	Retryable  bool
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

func (e APIError) IsRetryable() bool {
	return e.Retryable
}

// validateUserAdvanced validación avanzada con errores personalizados
func validateUserAdvanced(user User) error {
	if user.Name == "" {
		return &ValidationError{
			Field:   "name",
			Value:   user.Name,
			Message: "no puede estar vacío",
		}
	}

	if user.Email == "" {
		return &ValidationError{
			Field:   "email",
			Value:   user.Email,
			Message: "no puede estar vacío",
		}
	}

	if !strings.Contains(user.Email, "@") {
		return &ValidationError{
			Field:   "email",
			Value:   user.Email,
			Message: "formato inválido (debe contener @)",
		}
	}

	if user.Age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "no puede ser negativo",
		}
	}

	if user.Age > 150 {
		return &ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "demasiado alto (máximo 150)",
		}
	}

	return nil
}

// ====================================
// 🔗 EJERCICIO 3: Error Wrapping
// ====================================

func ejercicio3() {
	fmt.Println("🔗 Ejercicio 3: Error Wrapping y Unwrapping")
	fmt.Println("==========================================")

	// Simular operaciones en capas
	userIDs := []string{"user123", "missing", "invalid", "user456"}

	for _, userID := range userIDs {
		fmt.Printf("Procesando usuario: %s\n", userID)
		err := controllerLayer(userID)
		if err != nil {
			fmt.Printf("Error final: %v\n", err)

			// Unwrapping para mostrar la cadena
			fmt.Println("Cadena de errores:")
			current := err
			level := 1
			for current != nil {
				fmt.Printf("  %d. %v\n", level, current)
				current = errors.Unwrap(current)
				level++
			}

			// Verificar errores específicos
			if errors.Is(err, ErrUserNotFound) {
				fmt.Println("  → Acción: Crear nuevo usuario")
			} else if errors.Is(err, ErrInvalidUserID) {
				fmt.Println("  → Acción: Solicitar ID válido")
			}
		} else {
			fmt.Printf("Procesamiento exitoso ✅\n")
		}
		fmt.Println()
	}
}

// Errores predefinidos
var (
	ErrUserNotFound  = errors.New("usuario no encontrado")
	ErrInvalidUserID = errors.New("ID de usuario inválido")
	ErrDBConnection  = errors.New("error de conexión a base de datos")
)

// Simular capas de la aplicación
func databaseLayer(userID string) error {
	switch userID {
	case "missing":
		return ErrUserNotFound
	case "invalid":
		return ErrInvalidUserID
	case "db_error":
		return ErrDBConnection
	default:
		return nil
	}
}

func serviceLayer(userID string) error {
	err := databaseLayer(userID)
	if err != nil {
		return fmt.Errorf("service layer: falló la operación para usuario %s: %w", userID, err)
	}
	return nil
}

func controllerLayer(userID string) error {
	err := serviceLayer(userID)
	if err != nil {
		return fmt.Errorf("controller layer: error procesando solicitud: %w", err)
	}
	return nil
}

// ====================================
// 🔄 EJERCICIO 4: Retry Pattern
// ====================================

func ejercicio4() {
	fmt.Println("🔄 Ejercicio 4: Retry Pattern con Backoff")
	fmt.Println("=======================================")

	// Configuración de retry
	config := RetryConfig{
		MaxAttempts:   3,
		InitialDelay:  100 * time.Millisecond,
		MaxDelay:      2 * time.Second,
		BackoffFactor: 2.0,
		Jitter:        true,
	}

	ctx := context.Background()

	// Test con operación que falla y luego tiene éxito
	fmt.Println("Test 1: Operación que falla 2 veces y luego tiene éxito")
	counter := 0
	operation1 := func() error {
		counter++
		if counter < 3 {
			return &APIError{StatusCode: 500, Message: "Server error", Retryable: true}
		}
		return nil
	}

	err := RetryWithBackoff(ctx, config, operation1)
	if err != nil {
		fmt.Printf("❌ Falló después de retry: %v\n", err)
	} else {
		fmt.Printf("✅ Éxito después de %d intentos\n", counter)
	}
	fmt.Println()

	// Test con error no retryable
	fmt.Println("Test 2: Error no retryable")
	operation2 := func() error {
		return &APIError{StatusCode: 400, Message: "Bad request", Retryable: false}
	}

	err = RetryWithBackoff(ctx, config, operation2)
	if err != nil {
		fmt.Printf("❌ Error no retryable: %v\n", err)
	}
	fmt.Println()

	// Test con timeout del contexto
	fmt.Println("Test 3: Context timeout durante retry")
	timeoutCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	operation3 := func() error {
		return &APIError{StatusCode: 503, Message: "Service unavailable", Retryable: true}
	}

	err = RetryWithBackoff(timeoutCtx, config, operation3)
	if err != nil {
		fmt.Printf("❌ Timeout durante retry: %v\n", err)
	}
	fmt.Println()
}

// RetryConfig configuración para reintentos
type RetryConfig struct {
	MaxAttempts   int
	InitialDelay  time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
	Jitter        bool
}

// RetryableError interface para errores retryables
type RetryableError interface {
	error
	IsRetryable() bool
}

// RetryWithBackoff ejecuta operación con retry y backoff exponencial
func RetryWithBackoff(ctx context.Context, config RetryConfig, operation func() error) error {
	var lastErr error
	delay := config.InitialDelay

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		// Ejecutar operación
		err := operation()
		if err == nil {
			return nil // ¡Éxito!
		}

		lastErr = err

		// Verificar si el error es retryable
		if retryableErr, ok := err.(RetryableError); ok && !retryableErr.IsRetryable() {
			return fmt.Errorf("error no retryable en intento %d: %w", attempt, err)
		}

		// Si es el último intento, no esperar
		if attempt == config.MaxAttempts {
			break
		}

		// Calcular delay con backoff exponencial
		actualDelay := delay
		if config.Jitter {
			// Añadir jitter aleatorio ±25%
			jitter := time.Duration(rand.Float64() * 0.5 * float64(delay))
			if rand.Float64() < 0.5 {
				actualDelay += jitter
			} else {
				actualDelay -= jitter
			}
		}

		fmt.Printf("Intento %d falló, reintentando en %v: %v\n", attempt, actualDelay, err)

		// Esperar con posibilidad de cancelación
		select {
		case <-ctx.Done():
			return fmt.Errorf("retry cancelado por contexto después de %d intentos: %w", attempt, ctx.Err())
		case <-time.After(actualDelay):
			// Continuar con siguiente intento
		}

		// Incrementar delay para siguiente intento
		delay = time.Duration(float64(delay) * config.BackoffFactor)
		if delay > config.MaxDelay {
			delay = config.MaxDelay
		}
	}

	return fmt.Errorf("operación falló después de %d intentos: %w", config.MaxAttempts, lastErr)
}

// ====================================
// 🛡️ EJERCICIO 5: Circuit Breaker
// ====================================

func ejercicio5() {
	fmt.Println("🛡️ Ejercicio 5: Circuit Breaker Pattern")
	fmt.Println("======================================")

	// Crear circuit breaker
	cb := NewCircuitBreaker(3, 2*time.Second)

	// Simular servicio poco confiable
	failureCount := 0
	unreliableService := func() error {
		failureCount++
		// Fallar primeras 5 llamadas, luego tener éxito
		if failureCount <= 5 {
			return fmt.Errorf("servicio falló (llamada %d)", failureCount)
		}
		return nil
	}

	// Realizar múltiples llamadas
	for i := 0; i < 15; i++ {
		fmt.Printf("Llamada %d: ", i+1)
		err := cb.Execute(unreliableService)
		if err != nil {
			fmt.Printf("❌ %v (Estado: %s)\n", err, cb.GetState())
		} else {
			fmt.Printf("✅ Éxito (Estado: %s)\n", cb.GetState())
		}

		// Esperar un poco entre llamadas
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println()
}

// CircuitState estado del circuit breaker
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
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

// CircuitBreaker implementación
type CircuitBreaker struct {
	mu               sync.RWMutex
	state            CircuitState
	failureCount     int
	successCount     int
	lastFailureTime  time.Time
	maxFailures      int
	resetTimeout     time.Duration
	halfOpenMaxCalls int
}

// NewCircuitBreaker constructor
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		maxFailures:      maxFailures,
		resetTimeout:     resetTimeout,
		halfOpenMaxCalls: 3,
	}
}

// Execute ejecuta operación a través del circuit breaker
func (cb *CircuitBreaker) Execute(operation func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	// Verificar si debemos cambiar de Open a Half-Open
	if cb.state == StateOpen {
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
			cb.successCount = 0
			fmt.Printf("Circuit breaker: %s -> %s ", StateOpen, StateHalfOpen)
		} else {
			return errors.New("circuit breaker está OPEN")
		}
	}

	// En Half-Open, limitar número de llamadas
	if cb.state == StateHalfOpen && cb.successCount >= cb.halfOpenMaxCalls {
		return errors.New("circuit breaker HALF-OPEN: máximo de llamadas excedido")
	}

	// Ejecutar operación
	err := operation()

	if err != nil {
		cb.onFailure()
		return fmt.Errorf("operación falló (circuit: %s): %w", cb.state, err)
	}

	cb.onSuccess()
	return nil
}

// onFailure maneja fallos
func (cb *CircuitBreaker) onFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()

	switch cb.state {
	case StateClosed:
		if cb.failureCount >= cb.maxFailures {
			cb.state = StateOpen
			fmt.Printf("Circuit breaker: CLOSED -> OPEN (fallos: %d) ", cb.failureCount)
		}
	case StateHalfOpen:
		cb.state = StateOpen
		cb.failureCount = 1 // Reset counter
		fmt.Printf("Circuit breaker: HALF-OPEN -> OPEN ")
	}
}

// onSuccess maneja éxitos
func (cb *CircuitBreaker) onSuccess() {
	switch cb.state {
	case StateClosed:
		cb.failureCount = 0 // Reset failure count
	case StateHalfOpen:
		cb.successCount++
		if cb.successCount >= cb.halfOpenMaxCalls {
			cb.state = StateClosed
			cb.failureCount = 0
			fmt.Printf("Circuit breaker: HALF-OPEN -> CLOSED ")
		}
	}
}

// GetState retorna estado actual
func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// ====================================
// 📊 EJERCICIO 6: Error Aggregation
// ====================================

func ejercicio6() {
	fmt.Println("📊 Ejercicio 6: Error Aggregation")
	fmt.Println("================================")

	// Test de validación múltiple
	fmt.Println("Test 1: Validación múltiple con MultiError")
	user := User{Name: "", Email: "email-inválido", Age: -5}
	err := validateUserMultiple(user)
	if err != nil {
		fmt.Printf("Errores de validación:\n%v\n", err)
	}

	// Test de procesamiento paralelo
	fmt.Println("Test 2: Procesamiento paralelo con errores")
	items := []string{"item1", "bad_item", "item3", "another_bad", "item5"}
	err = processItemsConcurrently(items)
	if err != nil {
		fmt.Printf("Errores de procesamiento:\n%v\n", err)
	}

	fmt.Println()
}

// MultiError para manejar múltiples errores
type MultiError struct {
	errors []error
	mu     sync.Mutex
}

// NewMultiError constructor
func NewMultiError() *MultiError {
	return &MultiError{
		errors: make([]error, 0),
	}
}

// Add añade error al agregador
func (me *MultiError) Add(err error) {
	if err == nil {
		return
	}

	me.mu.Lock()
	defer me.mu.Unlock()
	me.errors = append(me.errors, err)
}

// Error implementa la interface error
func (me *MultiError) Error() string {
	me.mu.Lock()
	defer me.mu.Unlock()

	if len(me.errors) == 0 {
		return ""
	}

	if len(me.errors) == 1 {
		return me.errors[0].Error()
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("múltiples errores (%d):\n", len(me.errors)))

	for i, err := range me.errors {
		builder.WriteString(fmt.Sprintf("  %d. %v\n", i+1, err))
	}

	return builder.String()
}

// HasErrors verifica si hay errores
func (me *MultiError) HasErrors() bool {
	me.mu.Lock()
	defer me.mu.Unlock()
	return len(me.errors) > 0
}

// ErrorOrNil retorna error o nil si no hay errores
func (me *MultiError) ErrorOrNil() error {
	if me.HasErrors() {
		return me
	}
	return nil
}

// validateUserMultiple validación con múltiples errores
func validateUserMultiple(user User) error {
	multiErr := NewMultiError()

	// Validar nombre
	if user.Name == "" {
		multiErr.Add(fmt.Errorf("nombre es requerido"))
	} else if len(user.Name) < 2 {
		multiErr.Add(fmt.Errorf("nombre debe tener al menos 2 caracteres"))
	}

	// Validar email
	if user.Email == "" {
		multiErr.Add(fmt.Errorf("email es requerido"))
	} else if !strings.Contains(user.Email, "@") {
		multiErr.Add(fmt.Errorf("formato de email inválido"))
	}

	// Validar edad
	if user.Age < 0 {
		multiErr.Add(fmt.Errorf("edad no puede ser negativa"))
	} else if user.Age > 150 {
		multiErr.Add(fmt.Errorf("edad no puede ser mayor a 150"))
	}

	return multiErr.ErrorOrNil()
}

// processItemsConcurrently procesa items en paralelo
func processItemsConcurrently(items []string) error {
	multiErr := NewMultiError()
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		go func(index int, item string) {
			defer wg.Done()

			if err := processItem(item); err != nil {
				multiErr.Add(fmt.Errorf("item %d (%s): %w", index, item, err))
			}
		}(i, item)
	}

	wg.Wait()
	return multiErr.ErrorOrNil()
}

func processItem(item string) error {
	if strings.Contains(item, "bad") {
		return fmt.Errorf("item es inválido")
	}
	return nil
}

// ====================================
// 📝 EJERCICIO 7: Structured Logging
// ====================================

func ejercicio7() {
	fmt.Println("📝 Ejercicio 7: Structured Logging de Errores")
	fmt.Println("=============================================")

	logger := NewErrorLogger()
	ctx := context.Background()

	// Simular diferentes errores con contexto
	errors := []struct {
		err     error
		context ErrorContext
	}{
		{
			err: &ValidationError{Field: "email", Value: "invalid", Message: "formato inválido"},
			context: ErrorContext{
				UserID:    "user123",
				RequestID: "req-001",
				Operation: "validate_user",
				Component: "validation_service",
				Duration:  50 * time.Millisecond,
			},
		},
		{
			err: &APIError{StatusCode: 500, Message: "Database connection failed", Retryable: true},
			context: ErrorContext{
				UserID:    "user456",
				RequestID: "req-002",
				Operation: "fetch_user",
				Component: "database_service",
				Duration:  2 * time.Second,
				Metadata: map[string]interface{}{
					"table":   "users",
					"timeout": "30s",
				},
			},
		},
		{
			err: fmt.Errorf("nested error: %w", ErrUserNotFound),
			context: ErrorContext{
				UserID:    "user789",
				RequestID: "req-003",
				Operation: "update_profile",
				Component: "user_service",
			},
		},
	}

	for _, e := range errors {
		logger.LogError(ctx, e.err, e.context)
	}

	fmt.Println("✅ Errores loggeados con contexto estructurado")
	fmt.Println()
}

// ErrorContext estructura para contexto de error
type ErrorContext struct {
	UserID    string                 `json:"user_id,omitempty"`
	RequestID string                 `json:"request_id,omitempty"`
	Operation string                 `json:"operation,omitempty"`
	Component string                 `json:"component,omitempty"`
	Duration  time.Duration          `json:"duration,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ErrorLogger wrapper con contexto de error
type ErrorLogger struct {
	// En un caso real, usarías slog o logrus
	// Aquí simplificamos con log estándar
}

// NewErrorLogger constructor
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{}
}

// LogError logea error con contexto completo
func (el *ErrorLogger) LogError(ctx context.Context, err error, errorCtx ErrorContext) {
	if err == nil {
		return
	}

	// Construir estructura de log
	logEntry := map[string]interface{}{
		"timestamp":  time.Now().UTC(),
		"level":      "ERROR",
		"error":      err.Error(),
		"error_type": fmt.Sprintf("%T", err),
	}

	// Añadir contexto
	if errorCtx.UserID != "" {
		logEntry["user_id"] = errorCtx.UserID
	}
	if errorCtx.RequestID != "" {
		logEntry["request_id"] = errorCtx.RequestID
	}
	if errorCtx.Operation != "" {
		logEntry["operation"] = errorCtx.Operation
	}
	if errorCtx.Component != "" {
		logEntry["component"] = errorCtx.Component
	}
	if errorCtx.Duration > 0 {
		logEntry["duration_ms"] = errorCtx.Duration.Milliseconds()
	}

	// Añadir metadata
	if errorCtx.Metadata != nil {
		for key, value := range errorCtx.Metadata {
			logEntry[key] = value
		}
	}

	// Error chain si existe
	if unwrapped := errors.Unwrap(err); unwrapped != nil {
		chain := []string{}
		current := err
		for current != nil {
			chain = append(chain, current.Error())
			current = errors.Unwrap(current)
		}
		logEntry["error_chain"] = chain
	}

	// Serializar a JSON
	if jsonData, err := json.Marshal(logEntry); err == nil {
		log.Printf("STRUCTURED_ERROR: %s", string(jsonData))
	} else {
		log.Printf("ERROR: %v (failed to serialize: %v)", err, err)
	}
}

// ====================================
// 📈 EJERCICIO 8: Error Metrics
// ====================================

func ejercicio8() {
	fmt.Println("📈 Ejercicio 8: Error Metrics y Monitoring")
	fmt.Println("=========================================")

	metrics := NewErrorMetrics()

	// Simular errores de diferentes tipos
	errorTypes := []string{
		"ValidationError",
		"APIError",
		"DatabaseError",
		"ValidationError",
		"APIError",
		"NetworkError",
		"ValidationError",
	}

	components := []string{
		"auth_service",
		"user_service",
		"db_service",
		"api_gateway",
		"payment_service",
	}

	// Generar errores con diferentes frecuencias
	for i, errorType := range errorTypes {
		component := components[i%len(components)]
		operation := fmt.Sprintf("operation_%d", i%3)

		metrics.RecordError(errorType, component, operation)

		// Simular tiempo entre errores
		time.Sleep(10 * time.Millisecond)
	}

	// Mostrar métricas
	summary := metrics.GetMetricsSummary()
	fmt.Println("📊 Resumen de métricas:")

	if jsonData, err := json.MarshalIndent(summary, "", "  "); err == nil {
		fmt.Println(string(jsonData))
	}

	// Mostrar rates específicos
	fmt.Println("\n📊 Error rates por tipo:")
	for errorType := range metrics.errorCounts {
		rate := metrics.GetErrorRate(errorType)
		fmt.Printf("  %s: %.2f errores/min\n", errorType, rate)
	}

	fmt.Println()
}

// ErrorMetrics recolector de métricas de errores
type ErrorMetrics struct {
	mu                 sync.RWMutex
	errorCounts        map[string]int64
	errorRates         map[string][]time.Time
	componentErrors    map[string]int64
	operationErrors    map[string]int64
	lastErrorTimestamp time.Time
	totalErrors        int64
}

// NewErrorMetrics constructor
func NewErrorMetrics() *ErrorMetrics {
	return &ErrorMetrics{
		errorCounts:     make(map[string]int64),
		errorRates:      make(map[string][]time.Time),
		componentErrors: make(map[string]int64),
		operationErrors: make(map[string]int64),
	}
}

// RecordError registra un error en las métricas
func (em *ErrorMetrics) RecordError(errorType, component, operation string) {
	em.mu.Lock()
	defer em.mu.Unlock()

	now := time.Now()

	// Incrementar contadores
	em.errorCounts[errorType]++
	em.componentErrors[component]++
	em.operationErrors[operation]++
	em.totalErrors++
	em.lastErrorTimestamp = now

	// Mantener timestamps para rate calculation
	if _, exists := em.errorRates[errorType]; !exists {
		em.errorRates[errorType] = make([]time.Time, 0)
	}

	em.errorRates[errorType] = append(em.errorRates[errorType], now)

	// Limpiar timestamps antiguos (> 5 minutos)
	cutoff := now.Add(-5 * time.Minute)
	filtered := make([]time.Time, 0)
	for _, timestamp := range em.errorRates[errorType] {
		if timestamp.After(cutoff) {
			filtered = append(filtered, timestamp)
		}
	}
	em.errorRates[errorType] = filtered
}

// GetErrorRate retorna rate de errores por minuto
func (em *ErrorMetrics) GetErrorRate(errorType string) float64 {
	em.mu.RLock()
	defer em.mu.RUnlock()

	timestamps, exists := em.errorRates[errorType]
	if !exists || len(timestamps) == 0 {
		return 0
	}

	// Contar errores en últimos 5 minutos
	cutoff := time.Now().Add(-5 * time.Minute)
	count := 0
	for _, timestamp := range timestamps {
		if timestamp.After(cutoff) {
			count++
		}
	}

	return float64(count) / 5.0 // Errores por minuto
}

// GetMetricsSummary retorna resumen de métricas
func (em *ErrorMetrics) GetMetricsSummary() map[string]interface{} {
	em.mu.RLock()
	defer em.mu.RUnlock()

	summary := map[string]interface{}{
		"total_errors":         em.totalErrors,
		"last_error_timestamp": em.lastErrorTimestamp,
		"error_counts":         copyMapInt64(em.errorCounts),
		"component_errors":     copyMapInt64(em.componentErrors),
		"operation_errors":     copyMapInt64(em.operationErrors),
	}

	return summary
}

func copyMapInt64(original map[string]int64) map[string]int64 {
	copy := make(map[string]int64)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

// ====================================
// 🎯 EJERCICIO 9: Error Builder Pattern
// ====================================

func ejercicio9() {
	fmt.Println("🎯 Ejercicio 9: Error Builder Pattern")
	fmt.Println("====================================")

	// Crear errores usando builder pattern
	err1 := NewApplicationError(ErrorTypeValidation).
		Code("EMAIL_INVALID").
		Message("Email format is invalid").
		Detail("email", "user@invalid").
		Detail("expected_format", "user@domain.com").
		Context("user_id", "12345").
		Context("operation", "user_registration").
		Build()

	err2 := NewApplicationError(ErrorTypeDatabase).
		Code("CONNECTION_TIMEOUT").
		Message("Database connection timed out").
		Cause(fmt.Errorf("tcp connection failed")).
		Detail("timeout_seconds", 30).
		Detail("retry_count", 3).
		Context("service", "user_service").
		Context("database", "postgres_primary").
		Build()

	err3 := NewApplicationError(ErrorTypeBusiness).
		Code("INSUFFICIENT_BALANCE").
		Message("Account balance is insufficient").
		Detail("current_balance", 50.25).
		Detail("required_amount", 100.00).
		Detail("currency", "USD").
		Context("account_id", "acc_123").
		Context("transaction_id", "tx_456").
		Build()

	errors := []*ApplicationError{err1, err2, err3}

	for i, err := range errors {
		fmt.Printf("Error %d:\n", i+1)
		fmt.Printf("  Type: %s\n", err.Type)
		fmt.Printf("  Code: %s\n", err.Code)
		fmt.Printf("  Message: %s\n", err.Message)
		fmt.Printf("  Error: %v\n", err)
		if len(err.Details) > 0 {
			fmt.Printf("  Details: %+v\n", err.Details)
		}
		if len(err.Context) > 0 {
			fmt.Printf("  Context: %+v\n", err.Context)
		}
		if err.Cause != nil {
			fmt.Printf("  Cause: %v\n", err.Cause)
		}
		fmt.Printf("  Timestamp: %v\n", err.Timestamp.Format(time.RFC3339))
		fmt.Println()
	}
}

// ErrorType para categorizar errores
type ErrorType string

const (
	ErrorTypeValidation ErrorType = "validation"
	ErrorTypeNetwork    ErrorType = "network"
	ErrorTypeDatabase   ErrorType = "database"
	ErrorTypeBusiness   ErrorType = "business"
	ErrorTypeSystem     ErrorType = "system"
)

// ApplicationError error principal de la aplicación
type ApplicationError struct {
	Type      ErrorType              `json:"type"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Details   map[string]interface{} `json:"details,omitempty"`
	Cause     error                  `json:"-"`
	Timestamp time.Time              `json:"timestamp"`
	Context   map[string]string      `json:"context,omitempty"`
}

func (e *ApplicationError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap permite error wrapping chains
func (e *ApplicationError) Unwrap() error {
	return e.Cause
}

// Builder pattern para crear errores
type ErrorBuilder struct {
	err *ApplicationError
}

func NewApplicationError(errType ErrorType) *ErrorBuilder {
	return &ErrorBuilder{
		err: &ApplicationError{
			Type:      errType,
			Timestamp: time.Now(),
			Details:   make(map[string]interface{}),
			Context:   make(map[string]string),
		},
	}
}

func (b *ErrorBuilder) Code(code string) *ErrorBuilder {
	b.err.Code = code
	return b
}

func (b *ErrorBuilder) Message(message string) *ErrorBuilder {
	b.err.Message = message
	return b
}

func (b *ErrorBuilder) Cause(err error) *ErrorBuilder {
	b.err.Cause = err
	return b
}

func (b *ErrorBuilder) Detail(key string, value interface{}) *ErrorBuilder {
	b.err.Details[key] = value
	return b
}

func (b *ErrorBuilder) Context(key, value string) *ErrorBuilder {
	b.err.Context[key] = value
	return b
}

func (b *ErrorBuilder) Build() *ApplicationError {
	return b.err
}

// ====================================
// 🔄 EJERCICIO 10: Recovery y Panic
// ====================================

func ejercicio10() {
	fmt.Println("🔄 Ejercicio 10: Recovery y Panic Handling")
	fmt.Println("=========================================")

	// Test de funciones que pueden hacer panic
	functions := []struct {
		name string
		fn   func()
	}{
		{
			name: "División por cero",
			fn: func() {
				_ = 42 / 0
			},
		},
		{
			name: "Acceso a índice inválido",
			fn: func() {
				slice := []int{1, 2, 3}
				_ = slice[10]
			},
		},
		{
			name: "Panic explícito",
			fn: func() {
				panic("algo salió muy mal")
			},
		},
		{
			name: "Función normal (sin panic)",
			fn: func() {
				fmt.Println("  Operación normal completada")
			},
		},
	}

	logger := NewErrorLogger()

	for _, test := range functions {
		fmt.Printf("Testing: %s\n", test.name)
		err := SafeExecute(test.fn, func(recovered interface{}) {
			// Callback para manejar recovery
			logger.LogErrorWithRecover(context.Background(), recovered, ErrorContext{
				Operation: test.name,
				Component: "safe_executor",
			})
		})

		if err != nil {
			fmt.Printf("  ❌ Error capturado: %v\n", err)
		} else {
			fmt.Printf("  ✅ Ejecutado sin errores\n")
		}
		fmt.Println()
	}
}

// SafeExecute ejecuta función con recovery de panic
func SafeExecute(fn func(), onRecover func(interface{})) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			// Log del panic
			if onRecover != nil {
				onRecover(recovered)
			}

			// Convertir panic a error
			switch v := recovered.(type) {
			case error:
				err = fmt.Errorf("panic recovered: %w", v)
			case string:
				err = fmt.Errorf("panic recovered: %s", v)
			default:
				err = fmt.Errorf("panic recovered: %+v", v)
			}
		}
	}()

	// Ejecutar función
	fn()
	return nil
}

// LogErrorWithRecover logea panic recovery
func (el *ErrorLogger) LogErrorWithRecover(ctx context.Context, recovered interface{}, errorCtx ErrorContext) {
	logEntry := map[string]interface{}{
		"timestamp":    time.Now().UTC(),
		"level":        "ERROR",
		"panic_value":  recovered,
		"error_type":   "panic",
		"panic_type":   fmt.Sprintf("%T", recovered),
		"stack_trace":  "stack trace would be here", // En un caso real, usarías runtime.Stack()
	}

	// Añadir contexto
	if errorCtx.Operation != "" {
		logEntry["operation"] = errorCtx.Operation
	}
	if errorCtx.Component != "" {
		logEntry["component"] = errorCtx.Component
	}

	// Log simplificado para el ejemplo
	if jsonData, err := json.Marshal(logEntry); err == nil {
		log.Printf("PANIC_RECOVERY: %s", string(jsonData))
	}
}
