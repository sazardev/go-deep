// 🚀 Proyecto Final: Sistema de Gestión de APIs con Context
// =========================================================
// Este proyecto demuestra el uso profesional del Context package
// en un sistema de gestión de APIs real

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ==========================================
// 🏗️ TIPOS Y ESTRUCTURAS
// ==========================================

// Context keys tipadas
type contextKey string

const (
	UserIDKey    contextKey = "userID"
	RequestIDKey contextKey = "requestID"
	TenantIDKey  contextKey = "tenantID"
	TraceIDKey   contextKey = "traceID"
	SessionIDKey contextKey = "sessionID"
)

// Request representa una petición API
type Request struct {
	ID        string                 `json:"id"`
	Method    string                 `json:"method"`
	Endpoint  string                 `json:"endpoint"`
	UserID    string                 `json:"user_id"`
	TenantID  string                 `json:"tenant_id"`
	Headers   map[string]string      `json:"headers"`
	Body      map[string]interface{} `json:"body"`
	Timestamp time.Time              `json:"timestamp"`
}

// Response representa una respuesta API
type Response struct {
	RequestID  string                 `json:"request_id"`
	StatusCode int                    `json:"status_code"`
	Body       map[string]interface{} `json:"body"`
	Headers    map[string]string      `json:"headers"`
	Duration   time.Duration          `json:"duration"`
	Error      string                 `json:"error,omitempty"`
	Timestamp  time.Time              `json:"timestamp"`
}

// APIMetrics almacena métricas de rendimiento
type APIMetrics struct {
	TotalRequests     int           `json:"total_requests"`
	SuccessfulReqs    int           `json:"successful_requests"`
	FailedReqs        int           `json:"failed_requests"`
	TimeoutReqs       int           `json:"timeout_requests"`
	AvgResponseTime   time.Duration `json:"avg_response_time"`
	MaxResponseTime   time.Duration `json:"max_response_time"`
	ActiveConnections int           `json:"active_connections"`
	mutex             sync.RWMutex
}

// ==========================================
// 🎯 SERVICIO DE GESTIÓN DE APIS
// ==========================================

type APIManager struct {
	activeRequests sync.Map
	metrics        *APIMetrics
	rateLimiter    *RateLimiter
	authenticator  *Authenticator
	requestLogger  *RequestLogger
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewAPIManager() *APIManager {
	ctx, cancel := context.WithCancel(context.Background())

	return &APIManager{
		metrics:       &APIMetrics{},
		rateLimiter:   NewRateLimiter(100), // 100 requests por minuto
		authenticator: NewAuthenticator(),
		requestLogger: NewRequestLogger(),
		ctx:           ctx,
		cancel:        cancel,
	}
}

func (am *APIManager) ProcessRequest(req *Request) *Response {
	start := time.Now()

	// Crear context para este request con timeout
	reqCtx, cancel := context.WithTimeout(am.ctx, 30*time.Second)
	defer cancel()

	// Enriquecer context con metadata del request
	reqCtx = context.WithValue(reqCtx, RequestIDKey, req.ID)
	reqCtx = context.WithValue(reqCtx, UserIDKey, req.UserID)
	reqCtx = context.WithValue(reqCtx, TenantIDKey, req.TenantID)
	reqCtx = context.WithValue(reqCtx, TraceIDKey, generateTraceID())

	// Registrar request activo
	am.activeRequests.Store(req.ID, req)
	defer am.activeRequests.Delete(req.ID)

	// Pipeline de procesamiento con middleware
	response := am.processWithMiddleware(reqCtx, req)

	// Calcular duración y actualizar métricas
	response.Duration = time.Since(start)
	am.updateMetrics(response)

	return response
}

func (am *APIManager) processWithMiddleware(ctx context.Context, req *Request) *Response {
	// Middleware 1: Logging
	ctx = am.loggingMiddleware(ctx, req)

	// Middleware 2: Authentication
	if !am.authenticationMiddleware(ctx, req) {
		return &Response{
			RequestID:  req.ID,
			StatusCode: 401,
			Body:       map[string]interface{}{"error": "Unauthorized"},
			Error:      "Authentication failed",
			Timestamp:  time.Now(),
		}
	}

	// Middleware 3: Rate Limiting
	if !am.rateLimitingMiddleware(ctx, req) {
		return &Response{
			RequestID:  req.ID,
			StatusCode: 429,
			Body:       map[string]interface{}{"error": "Rate limit exceeded"},
			Error:      "Rate limit exceeded",
			Timestamp:  time.Now(),
		}
	}

	// Procesar request principal
	return am.handleBusinessLogic(ctx, req)
}

func (am *APIManager) loggingMiddleware(ctx context.Context, req *Request) context.Context {
	traceID := ctx.Value(TraceIDKey).(string)

	am.requestLogger.Log(fmt.Sprintf("[%s] %s %s - User: %s, Tenant: %s",
		traceID, req.Method, req.Endpoint, req.UserID, req.TenantID))

	return ctx
}

func (am *APIManager) authenticationMiddleware(ctx context.Context, req *Request) bool {
	traceID := ctx.Value(TraceIDKey).(string)

	// Simular autenticación asíncrona
	authChan := make(chan bool, 1)

	go func() {
		select {
		case <-time.After(50 * time.Millisecond): // Simular auth query
			authChan <- am.authenticator.Authenticate(req.UserID, req.Headers["Authorization"])
		case <-ctx.Done():
			authChan <- false
		}
	}()

	select {
	case authenticated := <-authChan:
		if authenticated {
			am.requestLogger.Log(fmt.Sprintf("[%s] Authentication successful", traceID))
		} else {
			am.requestLogger.Log(fmt.Sprintf("[%s] Authentication failed", traceID))
		}
		return authenticated
	case <-ctx.Done():
		am.requestLogger.Log(fmt.Sprintf("[%s] Authentication timeout", traceID))
		return false
	}
}

func (am *APIManager) rateLimitingMiddleware(ctx context.Context, req *Request) bool {
	traceID := ctx.Value(TraceIDKey).(string)

	allowed := am.rateLimiter.Allow(req.UserID)
	if !allowed {
		am.requestLogger.Log(fmt.Sprintf("[%s] Rate limit exceeded for user %s", traceID, req.UserID))
	}

	return allowed
}

func (am *APIManager) handleBusinessLogic(ctx context.Context, req *Request) *Response {
	traceID := ctx.Value(TraceIDKey).(string)

	am.requestLogger.Log(fmt.Sprintf("[%s] Processing business logic for %s", traceID, req.Endpoint))

	// Simular diferentes tipos de endpoints con diferentes latencias
	var processingTime time.Duration
	var shouldFail bool

	switch req.Endpoint {
	case "/api/users":
		processingTime = 100 * time.Millisecond
		shouldFail = rand.Float32() < 0.1 // 10% failure rate
	case "/api/orders":
		processingTime = 300 * time.Millisecond
		shouldFail = rand.Float32() < 0.05 // 5% failure rate
	case "/api/analytics":
		processingTime = 1 * time.Second
		shouldFail = rand.Float32() < 0.15 // 15% failure rate
	case "/api/reports":
		processingTime = 2 * time.Second
		shouldFail = rand.Float32() < 0.2 // 20% failure rate
	default:
		processingTime = 500 * time.Millisecond
		shouldFail = rand.Float32() < 0.1
	}

	// Procesar con posibilidad de cancelación
	select {
	case <-time.After(processingTime):
		if shouldFail {
			am.requestLogger.Log(fmt.Sprintf("[%s] Business logic failed", traceID))
			return &Response{
				RequestID:  req.ID,
				StatusCode: 500,
				Body:       map[string]interface{}{"error": "Internal server error"},
				Error:      "Processing failed",
				Timestamp:  time.Now(),
			}
		}

		am.requestLogger.Log(fmt.Sprintf("[%s] Business logic completed successfully", traceID))
		return &Response{
			RequestID:  req.ID,
			StatusCode: 200,
			Body: map[string]interface{}{
				"result":   "success",
				"data":     fmt.Sprintf("Processed %s for user %s", req.Endpoint, req.UserID),
				"trace_id": traceID,
			},
			Headers: map[string]string{
				"Content-Type": "application/json",
				"X-Trace-ID":   traceID,
			},
			Timestamp: time.Now(),
		}

	case <-ctx.Done():
		am.requestLogger.Log(fmt.Sprintf("[%s] Business logic cancelled: %v", traceID, ctx.Err()))
		return &Response{
			RequestID:  req.ID,
			StatusCode: 408,
			Body:       map[string]interface{}{"error": "Request timeout"},
			Error:      ctx.Err().Error(),
			Timestamp:  time.Now(),
		}
	}
}

func (am *APIManager) updateMetrics(response *Response) {
	am.metrics.mutex.Lock()
	defer am.metrics.mutex.Unlock()

	am.metrics.TotalRequests++

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		am.metrics.SuccessfulReqs++
	} else if response.StatusCode == 408 {
		am.metrics.TimeoutReqs++
	} else {
		am.metrics.FailedReqs++
	}

	// Actualizar tiempos de respuesta
	if response.Duration > am.metrics.MaxResponseTime {
		am.metrics.MaxResponseTime = response.Duration
	}

	// Calcular tiempo promedio (simplificado)
	am.metrics.AvgResponseTime = time.Duration(
		(int64(am.metrics.AvgResponseTime)*int64(am.metrics.TotalRequests-1) + int64(response.Duration)) /
			int64(am.metrics.TotalRequests))
}

func (am *APIManager) GetMetrics() *APIMetrics {
	am.metrics.mutex.RLock()
	defer am.metrics.mutex.RUnlock()

	// Contar conexiones activas
	activeCount := 0
	am.activeRequests.Range(func(key, value interface{}) bool {
		activeCount++
		return true
	})
	am.metrics.ActiveConnections = activeCount

	// Retornar copia de las métricas
	return &APIMetrics{
		TotalRequests:     am.metrics.TotalRequests,
		SuccessfulReqs:    am.metrics.SuccessfulReqs,
		FailedReqs:        am.metrics.FailedReqs,
		TimeoutReqs:       am.metrics.TimeoutReqs,
		AvgResponseTime:   am.metrics.AvgResponseTime,
		MaxResponseTime:   am.metrics.MaxResponseTime,
		ActiveConnections: activeCount,
	}
}

func (am *APIManager) Shutdown(timeout time.Duration) {
	fmt.Printf("🛑 Iniciando shutdown del API Manager (timeout: %v)\n", timeout)

	// Context para shutdown con timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Cancelar context principal
	am.cancel()

	// Esperar que terminen las requests activas
	done := make(chan struct{})
	go func() {
		for {
			activeCount := 0
			am.activeRequests.Range(func(key, value interface{}) bool {
				activeCount++
				return true
			})

			if activeCount == 0 {
				close(done)
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-done:
		fmt.Println("✅ Shutdown completado: todas las requests terminaron")
	case <-shutdownCtx.Done():
		fmt.Println("⏰ Shutdown timeout: forzando terminación")
		// Aquí podrías implementar cleanup forzado
	}
}

// ==========================================
// 🔐 SERVICIO DE AUTENTICACIÓN
// ==========================================

type Authenticator struct {
	validUsers map[string]string // userID -> token
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{
		validUsers: map[string]string{
			"user-1": "token-abc123",
			"user-2": "token-def456",
			"user-3": "token-ghi789",
			"admin":  "token-admin999",
		},
	}
}

func (a *Authenticator) Authenticate(userID, authHeader string) bool {
	if expectedToken, exists := a.validUsers[userID]; exists {
		return authHeader == "Bearer "+expectedToken
	}
	return false
}

// ==========================================
// 🚦 RATE LIMITER
// ==========================================

type RateLimiter struct {
	requests map[string][]time.Time
	limit    int
	window   time.Duration
	mutex    sync.RWMutex
}

func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    requestsPerMinute,
		window:   time.Minute,
	}

	// Goroutine para cleanup periódico
	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) Allow(userID string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// Obtener requests del usuario
	userRequests := rl.requests[userID]

	// Filtrar requests dentro de la ventana
	var validRequests []time.Time
	for _, reqTime := range userRequests {
		if reqTime.After(windowStart) {
			validRequests = append(validRequests, reqTime)
		}
	}

	// Verificar límite
	if len(validRequests) >= rl.limit {
		rl.requests[userID] = validRequests
		return false
	}

	// Agregar nueva request
	validRequests = append(validRequests, now)
	rl.requests[userID] = validRequests

	return true
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mutex.Lock()

		now := time.Now()
		windowStart := now.Add(-rl.window)

		for userID, requests := range rl.requests {
			var validRequests []time.Time
			for _, reqTime := range requests {
				if reqTime.After(windowStart) {
					validRequests = append(validRequests, reqTime)
				}
			}

			if len(validRequests) == 0 {
				delete(rl.requests, userID)
			} else {
				rl.requests[userID] = validRequests
			}
		}

		rl.mutex.Unlock()
	}
}

// ==========================================
// 📝 REQUEST LOGGER
// ==========================================

type RequestLogger struct {
	logs    []string
	mutex   sync.RWMutex
	maxLogs int
}

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{
		logs:    make([]string, 0),
		maxLogs: 1000,
	}
}

func (rl *RequestLogger) Log(message string) {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	timestamp := time.Now().Format("15:04:05.000")
	logEntry := fmt.Sprintf("[%s] %s", timestamp, message)

	rl.logs = append(rl.logs, logEntry)

	// Mantener solo los últimos maxLogs
	if len(rl.logs) > rl.maxLogs {
		rl.logs = rl.logs[len(rl.logs)-rl.maxLogs:]
	}

	// También imprimir a consola
	fmt.Println(logEntry)
}

func (rl *RequestLogger) GetRecentLogs(count int) []string {
	rl.mutex.RLock()
	defer rl.mutex.RUnlock()

	if count > len(rl.logs) {
		count = len(rl.logs)
	}

	start := len(rl.logs) - count
	if start < 0 {
		start = 0
	}

	return rl.logs[start:]
}

// ==========================================
// 🎲 UTILIDADES
// ==========================================

func generateRequestID() string {
	return fmt.Sprintf("req-%d", time.Now().UnixNano()%1000000)
}

func generateTraceID() string {
	return fmt.Sprintf("trace-%d", time.Now().UnixNano()%1000000)
}

func createSampleRequest(userID, endpoint string) *Request {
	return &Request{
		ID:       generateRequestID(),
		Method:   "GET",
		Endpoint: endpoint,
		UserID:   userID,
		TenantID: "tenant-" + userID[len(userID)-1:],
		Headers: map[string]string{
			"Authorization": "Bearer token-abc123",
			"Content-Type":  "application/json",
		},
		Body:      map[string]interface{}{"query": "sample"},
		Timestamp: time.Now(),
	}
}

// ==========================================
// 🎯 SIMULADOR DE CARGA
// ==========================================

type LoadSimulator struct {
	apiManager  *APIManager
	requestRate int // requests por segundo
	duration    time.Duration
}

func NewLoadSimulator(apiManager *APIManager, requestRate int, duration time.Duration) *LoadSimulator {
	return &LoadSimulator{
		apiManager:  apiManager,
		requestRate: requestRate,
		duration:    duration,
	}
}

func (ls *LoadSimulator) Run(ctx context.Context) {
	fmt.Printf("🚀 Iniciando simulación de carga: %d req/s por %v\n", ls.requestRate, ls.duration)

	// Context con timeout para la simulación
	simCtx, cancel := context.WithTimeout(ctx, ls.duration)
	defer cancel()

	// Ticker para controlar la tasa de requests
	interval := time.Second / time.Duration(ls.requestRate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	endpoints := []string{"/api/users", "/api/orders", "/api/analytics", "/api/reports"}
	users := []string{"user-1", "user-2", "user-3", "admin"}

	requestCount := 0

	for {
		select {
		case <-simCtx.Done():
			fmt.Printf("✅ Simulación completada: %d requests enviadas\n", requestCount)
			return

		case <-ticker.C:
			// Seleccionar endpoint y usuario aleatoriamente
			endpoint := endpoints[rand.Intn(len(endpoints))]
			user := users[rand.Intn(len(users))]

			req := createSampleRequest(user, endpoint)

			// Procesar request en goroutine separada
			go func() {
				response := ls.apiManager.ProcessRequest(req)
				_ = response // Podríamos procesar la respuesta aquí
			}()

			requestCount++

			// Cada 50 requests, mostrar métricas
			if requestCount%50 == 0 {
				ls.showMetrics()
			}
		}
	}
}

func (ls *LoadSimulator) showMetrics() {
	metrics := ls.apiManager.GetMetrics()

	fmt.Printf("📊 Métricas actuales:\n")
	fmt.Printf("   Total: %d | ✅ %d | ❌ %d | ⏰ %d | 🔗 %d activas\n",
		metrics.TotalRequests,
		metrics.SuccessfulReqs,
		metrics.FailedReqs,
		metrics.TimeoutReqs,
		metrics.ActiveConnections)
	fmt.Printf("   Tiempo promedio: %v | Máximo: %v\n",
		metrics.AvgResponseTime,
		metrics.MaxResponseTime)
}

// ==========================================
// 🏃‍♂️ FUNCIÓN PRINCIPAL
// ==========================================

func main() {
	fmt.Println("🚀 Sistema de Gestión de APIs con Context")
	fmt.Println("=========================================")

	// Inicializar el gestor de APIs
	apiManager := NewAPIManager()

	// Crear simulador de carga
	simulator := NewLoadSimulator(apiManager, 10, 30*time.Second) // 10 req/s por 30s

	// Context principal de la aplicación
	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Goroutine para capturar señales de interrupción (simulado)
	go func() {
		time.Sleep(35 * time.Second) // Simular Ctrl+C después de 35s
		fmt.Println("\n🛑 Señal de interrupción recibida")
		cancel()
	}()

	// Ejecutar simulación
	simulator.Run(mainCtx)

	// Mostrar métricas finales
	fmt.Println("\n📊 Métricas Finales:")
	fmt.Println("===================")
	finalMetrics := apiManager.GetMetrics()

	metricsJSON, _ := json.MarshalIndent(finalMetrics, "", "  ")
	fmt.Println(string(metricsJSON))

	// Mostrar logs recientes
	fmt.Println("\n📝 Últimos 10 Logs:")
	fmt.Println("===================")
	recentLogs := apiManager.requestLogger.GetRecentLogs(10)
	for _, log := range recentLogs {
		fmt.Println(log)
	}

	// Shutdown elegante
	fmt.Println()
	apiManager.Shutdown(5 * time.Second)

	fmt.Println("\n🎉 Sistema finalizado correctamente")
}
