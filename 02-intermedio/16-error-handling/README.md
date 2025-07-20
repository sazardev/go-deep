# 🚨 Lección 16: Error Handling Avanzado - Robustez y Resilencia

## 🎯 Objetivos de la Lección

Al finalizar esta lección, serás capaz de:
- Entender la filosofía de error handling en Go
- Crear tipos de errores personalizados y descriptivos
- Implementar error wrapping y unwrapping
- Usar el paquete `errors` de Go 1.13+
- Aplicar patrones avanzados de manejo de errores
- Diseñar APIs resilientes con error handling
- Implementar recovery y fallback strategies
- Crear sistemas de logging y monitoring de errores

---

## 🧠 Analogía: Error Handling como un Sistema de Alertas Médicas

Imagina un **sistema médico de emergencias** donde cada problema debe ser:

```
🏥 Hospital (Aplicación)
├── 🚨 Detector de Síntomas (Error Detection)
├── 🔍 Diagnóstico (Error Classification)
├── 📋 Historial Médico (Error Context)
├── 💊 Tratamiento (Error Recovery)
├── 📞 Notificación (Error Reporting)
└── 📊 Seguimiento (Error Monitoring)
```

**En Go, el error handling es similar:**
- **Detectamos** problemas inmediatamente
- **Clasificamos** el tipo de error
- **Agregamos contexto** para facilitar el diagnóstico
- **Implementamos recovery** cuando es posible
- **Reportamos** errores para monitoreo
- **Prevenimos** que errores menores se vuelvan críticos

---

## 📚 Filosofía del Error Handling en Go

### 🔧 Principios Fundamentales

Go adopta una filosofía **explícita y pragmática** para errores:

```go
// ❌ Otros lenguajes: try-catch oculta el flujo
try {
    result = riskyOperation()
} catch (Exception e) {
    // Error manejado "invisiblemente"
}

// ✅ Go: Error handling explícito y visible
result, err := riskyOperation()
if err != nil {
    // Error manejado explícitamente
    return fmt.Errorf("failed to perform operation: %w", err)
}
```

### 🎭 Ventajas del Approach de Go

| **Aspecto** | **Go Approach** | **Beneficio** |
|-------------|-----------------|---------------|
| **Visibilidad** | Errores explícitos en el código | Flujo de control claro |
| **Performance** | Sin overhead de stack unwinding | Mejor rendimiento |
| **Simplicidad** | Un solo mecanismo para errores | Fácil de entender |
| **Composabilidad** | Errores como valores | Fácil testing y wrapping |

---

## 🔍 Error Interface y Tipos Básicos

### 1. **La Interface Error**

```go
package main

import (
    "fmt"
    "errors"
)

// La interface error built-in
type error interface {
    Error() string
}

// Ejemplo básico de error
func ejemploBasico() {
    // Crear error simple
    err1 := errors.New("algo salió mal")
    fmt.Printf("Error 1: %v\n", err1)
    
    // Crear error con formato
    err2 := fmt.Errorf("falló la operación con código: %d", 500)
    fmt.Printf("Error 2: %v\n", err2)
    
    // Verificar si hay error
    if err2 != nil {
        fmt.Println("Se detectó un error!")
    }
}

// Función que retorna error
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero no permitida")
    }
    return a / b, nil
}

func ejemploDivision() {
    resultado, err := dividir(10, 0)
    if err != nil {
        fmt.Printf("Error en división: %v\n", err)
        return
    }
    fmt.Printf("Resultado: %.2f\n", resultado)
}
```

### 2. **Errores Personalizados**

```go
package main

import (
    "fmt"
    "time"
)

// Error personalizado básico
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field '%s' with value '%v': %s", 
        e.Field, e.Value, e.Message)
}

// Error personalizado avanzado
type DatabaseError struct {
    Operation string
    Table     string
    Timestamp time.Time
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database %s operation failed on table '%s' at %v: %v",
        e.Operation, e.Table, e.Timestamp.Format("2006-01-02 15:04:05"), e.Err)
}

// Implementar Unwrap para error wrapping
func (e DatabaseError) Unwrap() error {
    return e.Err
}

// Error con múltiples niveles de información
type APIError struct {
    StatusCode int
    Method     string
    URL        string
    Message    string
    Cause      error
    RequestID  string
}

func (e APIError) Error() string {
    return fmt.Sprintf("API error [%d] %s %s: %s (request: %s)",
        e.StatusCode, e.Method, e.URL, e.Message, e.RequestID)
}

func (e APIError) Unwrap() error {
    return e.Cause
}

// Ejemplo de uso
func ejemploErroresPersonalizados() {
    // Error de validación
    valErr := ValidationError{
        Field:   "email",
        Value:   "not-an-email",
        Message: "debe ser un email válido",
    }
    fmt.Printf("Validation Error: %v\n", valErr)
    
    // Error de base de datos
    dbErr := DatabaseError{
        Operation: "INSERT",
        Table:     "users",
        Timestamp: time.Now(),
        Err:       errors.New("duplicate key violation"),
    }
    fmt.Printf("Database Error: %v\n", dbErr)
    
    // Error de API
    apiErr := APIError{
        StatusCode: 503,
        Method:     "POST",
        URL:        "/api/users",
        Message:    "service temporarily unavailable",
        RequestID:  "req-123-456",
        Cause:      dbErr,
    }
    fmt.Printf("API Error: %v\n", apiErr)
}
```

---

## 🎁 Error Wrapping y Unwrapping (Go 1.13+)

### 1. **Error Wrapping Básico**

```go
package main

import (
    "errors"
    "fmt"
)

// Ejemplo de wrapping con fmt.Errorf
func operacionCompleja() error {
    err := operacionBasica()
    if err != nil {
        // Wrapping con %w verb
        return fmt.Errorf("operación compleja falló: %w", err)
    }
    return nil
}

func operacionBasica() error {
    return errors.New("error de conexión de red")
}

// Ejemplo de unwrapping
func ejemploWrapping() {
    err := operacionCompleja()
    if err != nil {
        fmt.Printf("Error principal: %v\n", err)
        
        // Unwrap para obtener el error original
        originalErr := errors.Unwrap(err)
        if originalErr != nil {
            fmt.Printf("Error original: %v\n", originalErr)
        }
        
        // Verificar si contiene un error específico
        if errors.Is(err, errors.New("error de conexión de red")) {
            fmt.Println("Es un error de red!")
        }
    }
}

// Error wrapping manual
type ServiceError struct {
    Service string
    Err     error
}

func (e ServiceError) Error() string {
    return fmt.Sprintf("service '%s' error: %v", e.Service, e.Err)
}

func (e ServiceError) Unwrap() error {
    return e.Err
}

func ejemploWrappingManual() {
    baseErr := errors.New("connection timeout")
    serviceErr := ServiceError{
        Service: "user-service",
        Err:     baseErr,
    }
    
    fmt.Printf("Service Error: %v\n", serviceErr)
    fmt.Printf("Unwrapped: %v\n", errors.Unwrap(serviceErr))
}
```

### 2. **Funciones errors.Is y errors.As**

```go
package main

import (
    "errors"
    "fmt"
    "net"
    "os"
)

// Errores sentinela (predefinidos)
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrUnauthorized = errors.New("unauthorized access")
)

// Error personalizado para demostrar errors.As
type TemporaryError struct {
    Message string
    Retry   bool
}

func (e TemporaryError) Error() string {
    return e.Message
}

func (e TemporaryError) Temporary() bool {
    return e.Retry
}

func ejemploIsAs() {
    fmt.Println("=== Ejemplo errors.Is ===")
    
    // Crear error wrapeado
    err := fmt.Errorf("database operation failed: %w", ErrUserNotFound)
    
    // errors.Is encuentra errores en la cadena
    if errors.Is(err, ErrUserNotFound) {
        fmt.Println("✅ Es un error de usuario no encontrado")
    }
    
    // Comparación directa no funciona con wrapping
    if err == ErrUserNotFound {
        fmt.Println("❌ No se ejecuta (comparación directa)")
    } else {
        fmt.Println("✅ Comparación directa falló (esperado con wrapping)")
    }
    
    fmt.Println("\n=== Ejemplo errors.As ===")
    
    // Crear error con tipo específico
    tempErr := TemporaryError{
        Message: "temporary network failure",
        Retry:   true,
    }
    wrappedTempErr := fmt.Errorf("operation failed: %w", tempErr)
    
    // errors.As extrae tipos específicos
    var temp TemporaryError
    if errors.As(wrappedTempErr, &temp) {
        fmt.Printf("✅ Error temporal encontrado: %s, Retry: %t\n", 
            temp.Message, temp.Temporary())
    }
    
    // Ejemplo con errores del sistema
    fmt.Println("\n=== Ejemplo con errores del sistema ===")
    _, err = os.Open("archivo_inexistente.txt")
    if err != nil {
        var pathErr *os.PathError
        if errors.As(err, &pathErr) {
            fmt.Printf("✅ Error de path: %s\n", pathErr.Path)
        }
    }
    
    // Ejemplo con errores de red
    _, err = net.Dial("tcp", "direccion-invalida:80")
    if err != nil {
        fmt.Printf("Error de red: %v\n", err)
        
        var netErr net.Error
        if errors.As(err, &netErr) {
            fmt.Printf("✅ Es error de red, Timeout: %t, Temporary: %t\n",
                netErr.Timeout(), netErr.Temporary())
        }
    }
}
```

---

## 🎯 Patrones Avanzados de Error Handling

### 1. **Result Type Pattern**

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Result type genérico
type Result[T any] struct {
    Value T
    Error error
}

// Constructor para éxito
func Ok[T any](value T) Result[T] {
    return Result[T]{Value: value, Error: nil}
}

// Constructor para error
func Err[T any](err error) Result[T] {
    var zero T
    return Result[T]{Value: zero, Error: err}
}

// Métodos útiles
func (r Result[T]) IsOk() bool {
    return r.Error == nil
}

func (r Result[T]) IsErr() bool {
    return r.Error != nil
}

func (r Result[T]) Unwrap() (T, error) {
    return r.Value, r.Error
}

// Ejemplo de uso
func parseJSON(data string) Result[map[string]interface{}] {
    var result map[string]interface{}
    err := json.Unmarshal([]byte(data), &result)
    if err != nil {
        return Err[map[string]interface{}](err)
    }
    return Ok(result)
}

func ejemploResultType() {
    fmt.Println("=== Result Type Pattern ===")
    
    // Caso exitoso
    validJSON := `{"name": "Go", "version": "1.24"}`
    result := parseJSON(validJSON)
    
    if result.IsOk() {
        fmt.Printf("✅ JSON parseado: %+v\n", result.Value)
    }
    
    // Caso de error
    invalidJSON := `{"name": "Go", "version":}`
    result = parseJSON(invalidJSON)
    
    if result.IsErr() {
        fmt.Printf("❌ Error parsing JSON: %v\n", result.Error)
    }
}
```

### 2. **Error Accumulator Pattern**

```go
package main

import (
    "fmt"
    "strings"
)

// Acumulador de errores
type ErrorList struct {
    errors []error
}

func (el *ErrorList) Add(err error) {
    if err != nil {
        el.errors = append(el.errors, err)
    }
}

func (el *ErrorList) HasErrors() bool {
    return len(el.errors) > 0
}

func (el *ErrorList) Error() string {
    if len(el.errors) == 0 {
        return ""
    }
    
    var messages []string
    for _, err := range el.errors {
        messages = append(messages, err.Error())
    }
    return fmt.Sprintf("multiple errors: [%s]", strings.Join(messages, "; "))
}

func (el *ErrorList) Errors() []error {
    return el.errors
}

// Ejemplo de validación que acumula errores
func validateUser(name, email string, age int) error {
    var errList ErrorList
    
    // Validar nombre
    if len(name) < 2 {
        errList.Add(errors.New("nombre debe tener al menos 2 caracteres"))
    }
    
    // Validar email
    if !strings.Contains(email, "@") {
        errList.Add(errors.New("email debe contener @"))
    }
    
    // Validar edad
    if age < 0 || age > 150 {
        errList.Add(errors.New("edad debe estar entre 0 y 150"))
    }
    
    if errList.HasErrors() {
        return &errList
    }
    return nil
}

func ejemploErrorAccumulator() {
    fmt.Println("=== Error Accumulator Pattern ===")
    
    // Validación con múltiples errores
    err := validateUser("A", "not-email", -5)
    if err != nil {
        fmt.Printf("❌ Errores de validación: %v\n", err)
        
        // Acceder a errores individuales
        if errList, ok := err.(*ErrorList); ok {
            fmt.Println("Errores individuales:")
            for i, e := range errList.Errors() {
                fmt.Printf("  %d. %v\n", i+1, e)
            }
        }
    }
    
    // Validación exitosa
    err = validateUser("John Doe", "john@example.com", 30)
    if err == nil {
        fmt.Println("✅ Usuario válido")
    }
}
```

### 3. **Circuit Breaker Pattern**

```go
package main

import (
    "fmt"
    "time"
    "sync"
    "errors"
)

// Estados del circuit breaker
type State int

const (
    StateClosed State = iota
    StateOpen
    StateHalfOpen
)

func (s State) String() string {
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

// Circuit breaker
type CircuitBreaker struct {
    mu              sync.RWMutex
    state           State
    failureCount    int
    successCount    int
    threshold       int
    timeout         time.Duration
    lastFailureTime time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:     StateClosed,
        threshold: threshold,
        timeout:   timeout,
    }
}

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open")

func (cb *CircuitBreaker) Execute(operation func() error) error {
    cb.mu.RLock()
    state := cb.state
    cb.mu.RUnlock()
    
    // Si está abierto, verificar si puede pasar a half-open
    if state == StateOpen {
        cb.mu.Lock()
        if time.Since(cb.lastFailureTime) > cb.timeout {
            cb.state = StateHalfOpen
            cb.successCount = 0
        } else {
            cb.mu.Unlock()
            return ErrCircuitBreakerOpen
        }
        cb.mu.Unlock()
    }
    
    // Ejecutar operación
    err := operation()
    
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if err != nil {
        cb.onFailure()
    } else {
        cb.onSuccess()
    }
    
    return err
}

func (cb *CircuitBreaker) onFailure() {
    cb.failureCount++
    cb.lastFailureTime = time.Now()
    
    if cb.state == StateHalfOpen || cb.failureCount >= cb.threshold {
        cb.state = StateOpen
    }
}

func (cb *CircuitBreaker) onSuccess() {
    cb.failureCount = 0
    
    if cb.state == StateHalfOpen {
        cb.successCount++
        if cb.successCount >= 3 { // Requiere 3 éxitos para cerrar
            cb.state = StateClosed
        }
    }
}

func (cb *CircuitBreaker) GetState() State {
    cb.mu.RLock()
    defer cb.mu.RUnlock()
    return cb.state
}

// Simulación de servicio externo
var failureRate float64 = 0.7 // 70% de fallos

func unreliableService() error {
    if time.Now().UnixNano()%100 < int64(failureRate*100) {
        return errors.New("service failure")
    }
    return nil
}

func ejemploCircuitBreaker() {
    fmt.Println("=== Circuit Breaker Pattern ===")
    
    cb := NewCircuitBreaker(3, 2*time.Second)
    
    // Simular múltiples llamadas
    for i := 0; i < 10; i++ {
        err := cb.Execute(unreliableService)
        state := cb.GetState()
        
        if err != nil {
            if errors.Is(err, ErrCircuitBreakerOpen) {
                fmt.Printf("Llamada %d: ⛔ Circuit Breaker OPEN\n", i+1)
            } else {
                fmt.Printf("Llamada %d: ❌ Service Error [%s]\n", i+1, state)
            }
        } else {
            fmt.Printf("Llamada %d: ✅ Success [%s]\n", i+1, state)
        }
        
        time.Sleep(500 * time.Millisecond)
    }
    
    // Esperar y probar recovery
    fmt.Println("\n⏳ Esperando timeout del circuit breaker...")
    time.Sleep(3 * time.Second)
    
    // Reducir tasa de fallos para permitir recovery
    failureRate = 0.2
    
    fmt.Println("📈 Intentando recovery...")
    for i := 0; i < 5; i++ {
        err := cb.Execute(unreliableService)
        state := cb.GetState()
        
        if err != nil {
            fmt.Printf("Recovery %d: ❌ Error [%s]\n", i+1, state)
        } else {
            fmt.Printf("Recovery %d: ✅ Success [%s]\n", i+1, state)
        }
        
        time.Sleep(300 * time.Millisecond)
    }
}
```

---

## 🎛️ Error Handling Middleware y Decorators

### 1. **HTTP Error Middleware**

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

// Estructura de respuesta de error estandarizada
type ErrorResponse struct {
    Error struct {
        Code      string    `json:"code"`
        Message   string    `json:"message"`
        Details   string    `json:"details,omitempty"`
        Timestamp time.Time `json:"timestamp"`
        RequestID string    `json:"request_id,omitempty"`
    } `json:"error"`
}

// Tipos de errores HTTP personalizados
type HTTPError struct {
    Code       string
    Message    string
    Details    string
    StatusCode int
    Err        error
}

func (e HTTPError) Error() string {
    return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
}

func (e HTTPError) Unwrap() error {
    return e.Err
}

// Constructores de errores comunes
func NewBadRequestError(message string) HTTPError {
    return HTTPError{
        Code:       "BAD_REQUEST",
        Message:    message,
        StatusCode: http.StatusBadRequest,
    }
}

func NewNotFoundError(resource string) HTTPError {
    return HTTPError{
        Code:       "NOT_FOUND",
        Message:    fmt.Sprintf("%s not found", resource),
        StatusCode: http.StatusNotFound,
    }
}

func NewInternalError(err error) HTTPError {
    return HTTPError{
        Code:       "INTERNAL_ERROR",
        Message:    "Internal server error",
        Details:    "An unexpected error occurred",
        StatusCode: http.StatusInternalServerError,
        Err:        err,
    }
}

// Middleware de manejo de errores
func ErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic recovered: %v", err)
                
                httpErr := NewInternalError(fmt.Errorf("panic: %v", err))
                writeErrorResponse(w, httpErr, "")
            }
        }()
        
        // Wrapper para capturar errores
        errorHandler := &ErrorAwareWriter{
            ResponseWriter: w,
            Request:        r,
        }
        
        next(errorHandler, r)
    }
}

// Writer que puede manejar errores
type ErrorAwareWriter struct {
    http.ResponseWriter
    Request *http.Request
}

func writeErrorResponse(w http.ResponseWriter, err HTTPError, requestID string) {
    response := ErrorResponse{}
    response.Error.Code = err.Code
    response.Error.Message = err.Message
    response.Error.Details = err.Details
    response.Error.Timestamp = time.Now().UTC()
    response.Error.RequestID = requestID
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(err.StatusCode)
    
    if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
        log.Printf("Error encoding error response: %v", encodeErr)
    }
    
    // Log para monitoring
    log.Printf("HTTP Error [%d]: %s - %s (Request: %s)", 
        err.StatusCode, err.Code, err.Message, requestID)
}

// Ejemplo de handlers con error handling
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    if userID == "" {
        err := NewBadRequestError("user ID is required")
        writeErrorResponse(w, err, "req-123")
        return
    }
    
    // Simular búsqueda de usuario
    if userID == "999" {
        err := NewNotFoundError("User")
        writeErrorResponse(w, err, "req-123")
        return
    }
    
    // Simular error interno
    if userID == "error" {
        err := NewInternalError(errors.New("database connection failed"))
        writeErrorResponse(w, err, "req-123")
        return
    }
    
    // Respuesta exitosa
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id":   userID,
        "name": "John Doe",
        "email": "john@example.com",
    })
}

func ejemploHTTPErrorHandling() {
    fmt.Println("=== HTTP Error Handling ===")
    fmt.Println("Iniciando servidor HTTP en :8080")
    fmt.Println("Prueba estos endpoints:")
    fmt.Println("  - http://localhost:8080/user?id=123 (éxito)")
    fmt.Println("  - http://localhost:8080/user (bad request)")
    fmt.Println("  - http://localhost:8080/user?id=999 (not found)")
    fmt.Println("  - http://localhost:8080/user?id=error (internal error)")
    
    http.HandleFunc("/user", ErrorMiddleware(getUserHandler))
    
    // Solo mostrar el setup, no iniciar servidor real
    fmt.Println("✅ Servidor configurado (no iniciado para demo)")
}
```

---

## 📊 Monitoring y Logging de Errores

### 1. **Sistema de Métricas de Errores**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Métricas de errores
type ErrorMetrics struct {
    mu                sync.RWMutex
    totalErrors       int64
    errorsByType      map[string]int64
    errorsByService   map[string]int64
    recentErrors      []ErrorEvent
    maxRecentErrors   int
}

type ErrorEvent struct {
    Timestamp time.Time
    Type      string
    Service   string
    Message   string
    Severity  string
}

func NewErrorMetrics(maxRecent int) *ErrorMetrics {
    return &ErrorMetrics{
        errorsByType:    make(map[string]int64),
        errorsByService: make(map[string]int64),
        recentErrors:    make([]ErrorEvent, 0, maxRecent),
        maxRecentErrors: maxRecent,
    }
}

func (em *ErrorMetrics) RecordError(errorType, service, message, severity string) {
    em.mu.Lock()
    defer em.mu.Unlock()
    
    em.totalErrors++
    em.errorsByType[errorType]++
    em.errorsByService[service]++
    
    // Agregar a errores recientes
    event := ErrorEvent{
        Timestamp: time.Now(),
        Type:      errorType,
        Service:   service,
        Message:   message,
        Severity:  severity,
    }
    
    if len(em.recentErrors) >= em.maxRecentErrors {
        // Remover el más antiguo
        em.recentErrors = em.recentErrors[1:]
    }
    em.recentErrors = append(em.recentErrors, event)
}

func (em *ErrorMetrics) GetStats() map[string]interface{} {
    em.mu.RLock()
    defer em.mu.RUnlock()
    
    // Copiar maps para evitar race conditions
    errorsByType := make(map[string]int64)
    for k, v := range em.errorsByType {
        errorsByType[k] = v
    }
    
    errorsByService := make(map[string]int64)
    for k, v := range em.errorsByService {
        errorsByService[k] = v
    }
    
    return map[string]interface{}{
        "total_errors":      em.totalErrors,
        "errors_by_type":    errorsByType,
        "errors_by_service": errorsByService,
        "recent_count":      len(em.recentErrors),
    }
}

func (em *ErrorMetrics) GetRecentErrors() []ErrorEvent {
    em.mu.RLock()
    defer em.mu.RUnlock()
    
    // Copiar slice
    recent := make([]ErrorEvent, len(em.recentErrors))
    copy(recent, em.recentErrors)
    return recent
}

// Logger de errores estructurado
type StructuredLogger struct {
    metrics *ErrorMetrics
}

func NewStructuredLogger() *StructuredLogger {
    return &StructuredLogger{
        metrics: NewErrorMetrics(100),
    }
}

func (sl *StructuredLogger) LogError(err error, service string, extra map[string]interface{}) {
    errorType := "unknown"
    severity := "error"
    
    // Determinar tipo de error
    var httpErr HTTPError
    if errors.As(err, &httpErr) {
        errorType = httpErr.Code
        if httpErr.StatusCode >= 500 {
            severity = "critical"
        } else if httpErr.StatusCode >= 400 {
            severity = "warning"
        }
    }
    
    // Registrar en métricas
    sl.metrics.RecordError(errorType, service, err.Error(), severity)
    
    // Log estructurado
    logEntry := map[string]interface{}{
        "timestamp": time.Now().UTC().Format(time.RFC3339),
        "level":     severity,
        "service":   service,
        "error":     err.Error(),
        "type":      errorType,
    }
    
    // Agregar campos extra
    for k, v := range extra {
        logEntry[k] = v
    }
    
    // En una aplicación real, esto iría a un sistema de logging
    fmt.Printf("LOG: %+v\n", logEntry)
}

func (sl *StructuredLogger) PrintStats() {
    stats := sl.metrics.GetStats()
    recent := sl.metrics.GetRecentErrors()
    
    fmt.Println("\n=== Error Statistics ===")
    fmt.Printf("Total Errors: %v\n", stats["total_errors"])
    fmt.Printf("Errors by Type: %+v\n", stats["errors_by_type"])
    fmt.Printf("Errors by Service: %+v\n", stats["errors_by_service"])
    
    fmt.Println("\n=== Recent Errors ===")
    for i, event := range recent {
        if i >= 5 { // Solo mostrar los últimos 5
            break
        }
        fmt.Printf("  [%s] %s/%s: %s (%s)\n",
            event.Timestamp.Format("15:04:05"),
            event.Service,
            event.Type,
            event.Message,
            event.Severity)
    }
}

func ejemploErrorMonitoring() {
    fmt.Println("=== Error Monitoring y Logging ===")
    
    logger := NewStructuredLogger()
    
    // Simular varios errores
    errors := []struct {
        err     error
        service string
        extra   map[string]interface{}
    }{
        {
            err:     NewBadRequestError("invalid email format"),
            service: "user-service",
            extra:   map[string]interface{}{"user_id": "123", "action": "create"},
        },
        {
            err:     NewNotFoundError("Product"),
            service: "catalog-service",
            extra:   map[string]interface{}{"product_id": "999", "category": "electronics"},
        },
        {
            err:     NewInternalError(errors.New("database timeout")),
            service: "order-service",
            extra:   map[string]interface{}{"order_id": "ord-456", "db_host": "db-1"},
        },
        {
            err:     NewBadRequestError("missing required field"),
            service: "user-service",
            extra:   map[string]interface{}{"field": "password", "action": "login"},
        },
        {
            err:     NewInternalError(errors.New("redis connection failed")),
            service: "session-service",
            extra:   map[string]interface{}{"cache_key": "sess-789"},
        },
    }
    
    // Log errores
    for _, errInfo := range errors {
        logger.LogError(errInfo.err, errInfo.service, errInfo.extra)
        time.Sleep(100 * time.Millisecond)
    }
    
    // Mostrar estadísticas
    logger.PrintStats()
}
```

---

## 🧪 Testing de Error Handling

### 1. **Test de Errores con Table-Driven Tests**

```go
package main

import (
    "errors"
    "fmt"
    "testing"
)

// Función para testear
func processData(input string) (string, error) {
    if input == "" {
        return "", NewBadRequestError("input cannot be empty")
    }
    if input == "invalid" {
        return "", NewInternalError(errors.New("processing failed"))
    }
    if len(input) > 100 {
        return "", NewBadRequestError("input too long")
    }
    return fmt.Sprintf("processed: %s", input), nil
}

// Test cases para errores
func TestProcessDataErrors(t *testing.T) {
    tests := []struct {
        name        string
        input       string
        wantResult  string
        wantErr     bool
        errType     string
        errCode     string
    }{
        {
            name:       "successful processing",
            input:      "valid data",
            wantResult: "processed: valid data",
            wantErr:    false,
        },
        {
            name:    "empty input error",
            input:   "",
            wantErr: true,
            errType: "HTTPError",
            errCode: "BAD_REQUEST",
        },
        {
            name:    "invalid input error",
            input:   "invalid",
            wantErr: true,
            errType: "HTTPError",
            errCode: "INTERNAL_ERROR",
        },
        {
            name:    "input too long error",
            input:   string(make([]byte, 101)), // 101 caracteres
            wantErr: true,
            errType: "HTTPError",
            errCode: "BAD_REQUEST",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := processData(tt.input)
            
            // Verificar si debe haber error
            if tt.wantErr && err == nil {
                t.Errorf("Expected error but got none")
                return
            }
            if !tt.wantErr && err != nil {
                t.Errorf("Unexpected error: %v", err)
                return
            }
            
            // Verificar resultado en caso exitoso
            if !tt.wantErr && result != tt.wantResult {
                t.Errorf("Expected result %q, got %q", tt.wantResult, result)
            }
            
            // Verificar tipo de error
            if tt.wantErr && err != nil {
                var httpErr HTTPError
                if !errors.As(err, &httpErr) {
                    t.Errorf("Expected HTTPError, got %T", err)
                    return
                }
                
                if httpErr.Code != tt.errCode {
                    t.Errorf("Expected error code %q, got %q", tt.errCode, httpErr.Code)
                }
            }
        })
    }
}

// Helper para testing
func assertError(t *testing.T, err error, expectedType interface{}) {
    t.Helper()
    
    if err == nil {
        t.Fatal("Expected error but got none")
    }
    
    switch expected := expectedType.(type) {
    case string:
        // Verificar mensaje de error
        if err.Error() != expected {
            t.Errorf("Expected error message %q, got %q", expected, err.Error())
        }
    case error:
        // Verificar que sea el mismo error
        if !errors.Is(err, expected) {
            t.Errorf("Expected error %v, got %v", expected, err)
        }
    default:
        // Verificar tipo con errors.As
        if !errors.As(err, expectedType) {
            t.Errorf("Expected error of type %T, got %T", expectedType, err)
        }
    }
}

func ejemploTesting() {
    fmt.Println("=== Testing de Error Handling ===")
    fmt.Println("Ejecutar: go test -v para ver los tests")
    fmt.Println("Tests implementados:")
    fmt.Println("  ✅ Casos exitosos")
    fmt.Println("  ✅ Diferentes tipos de errores")
    fmt.Println("  ✅ Verificación de códigos de error")
    fmt.Println("  ✅ Error assertions con errors.As")
    fmt.Println("  ✅ Table-driven tests para cobertura completa")
}
```

---

## 💡 Best Practices y Anti-patterns

### 1. **Do's and Don'ts**

```go
package main

import (
    "fmt"
    "log"
)

// ✅ GOOD PRACTICES

// 1. Error sentinela para comparaciones
var ErrConfigNotFound = errors.New("configuration not found")

// 2. Tipos de error descriptivos
type ConfigError struct {
    Key     string
    Reason  string
    Err     error
}

func (e ConfigError) Error() string {
    return fmt.Sprintf("config error for key '%s': %s", e.Key, e.Reason)
}

func (e ConfigError) Unwrap() error {
    return e.Err
}

// 3. Error wrapping con contexto
func loadConfig(path string) (*Config, error) {
    data, err := readFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
    }
    
    config, err := parseConfig(data)
    if err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    return config, nil
}

// 4. Validación de entrada con errores específicos
func validateEmail(email string) error {
    if email == "" {
        return ConfigError{
            Key:    "email",
            Reason: "cannot be empty",
        }
    }
    
    if !strings.Contains(email, "@") {
        return ConfigError{
            Key:    "email", 
            Reason: "must contain @ symbol",
        }
    }
    
    return nil
}

// 5. Error handling en defer
func processFile(filename string) (err error) {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            // Solo loggear si no hay error principal
            if err == nil {
                err = fmt.Errorf("failed to close file: %w", closeErr)
            } else {
                log.Printf("Warning: failed to close file: %v", closeErr)
            }
        }
    }()
    
    // Procesar archivo...
    return nil
}

// ❌ ANTI-PATTERNS

// 1. ❌ Ignorar errores (nunca hacer esto)
func badExample1() {
    data, _ := readFile("config.json") // ¡MAL! Ignora el error
    // usar data sin verificar si hay error
}

// 2. ❌ Panic por errores recuperables
func badExample2(id string) {
    user, err := findUser(id)
    if err != nil {
        panic(err) // ¡MAL! No hagas panic por errores normales
    }
    // usar user
}

// 3. ❌ Loggear y retornar error (double handling)
func badExample3() error {
    err := doSomething()
    if err != nil {
        log.Printf("Error: %v", err) // ¡MAL! Loggear...
        return err                   // ...y retornar duplica el handling
    }
    return nil
}

// 4. ❌ Error strings con mayúscula
func badExample4() error {
    return errors.New("Error occurred") // ¡MAL! No empezar con mayúscula
}

// 5. ❌ No preservar error original
func badExample5() error {
    err := doSomething()
    if err != nil {
        return errors.New("operation failed") // ¡MAL! Pierde contexto original
    }
    return nil
}

// ✅ CORRECCIONES

// 1. ✅ Siempre verificar errores
func goodExample1() error {
    data, err := readFile("config.json")
    if err != nil {
        return fmt.Errorf("failed to read config: %w", err)
    }
    // usar data seguramente
    return nil
}

// 2. ✅ Retornar error, no panic
func goodExample2(id string) (*User, error) {
    user, err := findUser(id)
    if err != nil {
        return nil, fmt.Errorf("failed to find user %s: %w", id, err)
    }
    return user, nil
}

// 3. ✅ Loggear en el nivel más alto o retornar
func goodExample3() error {
    err := doSomething()
    if err != nil {
        // Solo retornar, dejar que el caller loggee
        return fmt.Errorf("operation failed: %w", err)
    }
    return nil
}

// 4. ✅ Error strings en minúscula, sin puntuación
func goodExample4() error {
    return errors.New("operation failed") // ✅ Correcto
}

// 5. ✅ Preservar error original con wrapping
func goodExample5() error {
    err := doSomething()
    if err != nil {
        return fmt.Errorf("operation failed: %w", err) // ✅ Preserva original
    }
    return nil
}

func ejemploBestPractices() {
    fmt.Println("=== Best Practices de Error Handling ===")
    fmt.Println("✅ DO:")
    fmt.Println("  - Siempre verificar errores")
    fmt.Println("  - Usar error wrapping con contexto")
    fmt.Println("  - Crear tipos de error descriptivos")
    fmt.Println("  - Error strings en minúscula")
    fmt.Println("  - Loggear en el nivel apropiado")
    fmt.Println("  - Usar defer para cleanup con error handling")
    
    fmt.Println("\n❌ DON'T:")
    fmt.Println("  - Ignorar errores con _")
    fmt.Println("  - Hacer panic por errores recuperables")
    fmt.Println("  - Double handling (log + return)")
    fmt.Println("  - Perder contexto del error original")
    fmt.Println("  - Error strings con mayúscula o puntuación")
    fmt.Println("  - Errores genéricos sin contexto")
}

// Stubs para ejemplos (no implementados)
type Config struct{}
type User struct{}

func readFile(path string) ([]byte, error) { return nil, nil }
func parseConfig(data []byte) (*Config, error) { return nil, nil }
func findUser(id string) (*User, error) { return nil, nil }
func doSomething() error { return nil }
```

---

## 🎯 Resumen de la Lección

### ✅ Conceptos Clave Aprendidos

1. **🔍 Error Interface**: La base del sistema de errores en Go
2. **🎁 Error Wrapping**: Preservar contexto con fmt.Errorf y %w
3. **🔧 Tipos Personalizados**: Crear errores descriptivos y tipados
4. **🎯 errors.Is/As**: Verificar y extraer tipos de errores
5. **🏗️ Patrones Avanzados**: Result type, accumulator, circuit breaker
6. **🎛️ Middleware**: Manejo centralizado de errores HTTP
7. **📊 Monitoring**: Métricas y logging estructurado
8. **🧪 Testing**: Verificación robusta de error handling
9. **💡 Best Practices**: Patrones correctos y anti-patterns

### 🏆 Logros Desbloqueados

- [ ] 🥇 **Error Novice**: Manejo básico de errores
- [ ] 🥈 **Wrapper Master**: Dominio de error wrapping
- [ ] 🥉 **Type Expert**: Tipos personalizados avanzados
- [ ] 🏅 **Pattern Architect**: Implementación de patrones complejos
- [ ] 🎖️ **Resilience Engineer**: Sistemas resilientes con error handling
- [ ] 🏆 **Error Wizard**: Monitoreo y debugging de errores

### 📚 Próximos Pasos

En la **Lección 17: Testing Avanzado**, aprenderemos:
- Test-driven development (TDD)
- Mocking y test doubles
- Property-based testing
- Integration testing

---

**🎉 ¡Felicitaciones! Ahora dominas el error handling avanzado en Go. Tus aplicaciones serán robustas, resilientes y fáciles de debuggear. El error handling sólido es la diferencia entre código amateur y código de producción.**

*Recuerda: "Un error bien manejado es una oportunidad de crear software más resiliente" - Software Wisdom*
