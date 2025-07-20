# 🎯 Lección 15: Context Package - Control y Cancelación

## 🎯 Objetivos de la Lección

Al finalizar esta lección, serás capaz de:
- Entender qué es el Context package y por qué es esencial
- Implementar cancelación elegante en aplicaciones Go
- Usar timeouts y deadlines efectivamente
- Propagar valores a través de call stacks
- Aplicar mejores prácticas con Context
- Debuggear problemas relacionados con Context
- Diseñar APIs que usen Context correctamente

---

## 🧠 Analogía: Context como el "Control Remoto Universal"

Imagina que tienes múltiples dispositivos funcionando en tu casa (televisión, aire acondicionado, luces, música) y necesitas un **control remoto universal** que pueda:

```
🏠 Casa (Aplicación Go)
├── 📺 TV (Goroutine 1)
├── ❄️ A/C (Goroutine 2) 
├── 💡 Luces (Goroutine 3)
├── 🎵 Música (Goroutine 4)
└── 🎮 Control Remoto (Context)
    ├── ⏹️ Apagar Todo (Cancel)
    ├── ⏰ Temporizador (Timeout)
    ├── 🔋 Estado Batería (Values)
    └── 📡 Señal (Propagation)
```

El **Context** es ese control remoto que:
- **Cancela** todas las operaciones coordinadamente
- **Controla timeouts** para operaciones que tardan mucho
- **Propaga información** como user ID, request ID, etc.
- **Se hereda** de padres a hijos automáticamente

---

## 📚 Fundamentos del Context Package

### 🔧 ¿Qué es Context?

**Context** es un package estándar de Go que proporciona una forma de transmitir:
- **Cancelación**: Señales para terminar operaciones
- **Timeouts**: Límites de tiempo para operaciones
- **Deadlines**: Fechas específicas de expiración
- **Valores**: Información que debe pasar por el call stack

### 🎭 Los Cuatro Tipos de Context

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func ejemplosContext() {
    // 1. Background Context - El padre de todos
    ctx := context.Background()
    fmt.Printf("Background: %T\n", ctx)
    
    // 2. TODO Context - Para desarrollo
    todoCtx := context.TODO()
    fmt.Printf("TODO: %T\n", todoCtx)
    
    // 3. WithCancel - Para cancelación manual
    cancelCtx, cancel := context.WithCancel(ctx)
    defer cancel()
    fmt.Printf("WithCancel: %T\n", cancelCtx)
    
    // 4. WithTimeout - Para límites de tiempo
    timeoutCtx, cancel2 := context.WithTimeout(ctx, 5*time.Second)
    defer cancel2()
    fmt.Printf("WithTimeout: %T\n", timeoutCtx)
    
    // 5. WithDeadline - Para fechas específicas
    deadline := time.Now().Add(10 * time.Second)
    deadlineCtx, cancel3 := context.WithDeadline(ctx, deadline)
    defer cancel3()
    fmt.Printf("WithDeadline: %T\n", deadlineCtx)
    
    // 6. WithValue - Para propagar valores
    valueCtx := context.WithValue(ctx, "userID", "12345")
    fmt.Printf("WithValue: %T\n", valueCtx)
}
```

---

## 🚫 Cancelación con Context

### 1. 🌟 Cancelación Básica

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("👷 Worker %d: recibió cancelación - %v\n", id, ctx.Err())
            return
        case <-time.After(500 * time.Millisecond):
            fmt.Printf("👷 Worker %d: trabajando...\n", id)
        }
    }
}

func ejemploCancelacionBasica() {
    fmt.Println("🚫 Ejemplo: Cancelación Básica")
    fmt.Println("==============================")
    
    // Crear context cancelable
    ctx, cancel := context.WithCancel(context.Background())
    
    // Lanzar workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }
    
    // Dejar trabajar por 3 segundos
    fmt.Println("⏱️ Trabajadores activos por 3 segundos...")
    time.Sleep(3 * time.Second)
    
    // Cancelar todos los workers
    fmt.Println("📤 Enviando señal de cancelación...")
    cancel()
    
    // Dar tiempo para que terminen
    time.Sleep(500 * time.Millisecond)
    fmt.Println("✅ Todos los workers terminados\n")
}
```

### 2. 🎯 Cancelación Jerárquica

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func taskManager(ctx context.Context, name string, tasks int) {
    // Crear sub-context para este manager
    subCtx, cancel := context.WithCancel(ctx)
    defer cancel()
    
    var wg sync.WaitGroup
    
    fmt.Printf("🎯 Task Manager '%s' iniciando con %d tasks\n", name, tasks)
    
    // Lanzar tasks
    for i := 1; i <= tasks; i++ {
        wg.Add(1)
        go func(taskID int) {
            defer wg.Done()
            executeTask(subCtx, fmt.Sprintf("%s-Task-%d", name, taskID))
        }(i)
    }
    
    // Esperar completion o cancelación
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        fmt.Printf("✅ Task Manager '%s' completado normalmente\n", name)
    case <-ctx.Done():
        fmt.Printf("🚫 Task Manager '%s' cancelado: %v\n", name, ctx.Err())
        cancel() // Cancelar sub-tasks
        wg.Wait() // Esperar que terminen
    }
}

func executeTask(ctx context.Context, taskName string) {
    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()
    
    iterations := 0
    maxIterations := 10
    
    for iterations < maxIterations {
        select {
        case <-ctx.Done():
            fmt.Printf("  🔴 %s cancelado después de %d iteraciones\n", taskName, iterations)
            return
        case <-ticker.C:
            iterations++
            fmt.Printf("  ⚡ %s iteración %d/%d\n", taskName, iterations, maxIterations)
        }
    }
    
    fmt.Printf("  ✅ %s completado exitosamente\n", taskName)
}

func ejemploCancelacionJerarquica() {
    fmt.Println("🎯 Ejemplo: Cancelación Jerárquica")
    fmt.Println("=================================")
    
    // Context principal
    rootCtx, rootCancel := context.WithCancel(context.Background())
    
    // Lanzar múltiples task managers
    go taskManager(rootCtx, "Frontend", 2)
    go taskManager(rootCtx, "Backend", 3)
    go taskManager(rootCtx, "Database", 2)
    
    // Dejar ejecutar por 3 segundos
    time.Sleep(3 * time.Second)
    
    // Cancelar todo desde la raíz
    fmt.Println("\n📤 Cancelando desde la raíz...")
    rootCancel()
    
    // Dar tiempo para cleanup
    time.Sleep(1 * time.Second)
    fmt.Println("🏁 Cancelación jerárquica completada\n")
}
```

---

## ⏰ Timeouts y Deadlines

### 1. 🕐 Timeout Básico

```go
package main

import (
    "context"
    "fmt"
    "math/rand"
    "time"
)

func operacionLenta(ctx context.Context, nombre string, duracion time.Duration) error {
    fmt.Printf("🚀 Iniciando %s (duración: %v)\n", nombre, duracion)
    
    select {
    case <-time.After(duracion):
        fmt.Printf("✅ %s completada exitosamente\n", nombre)
        return nil
    case <-ctx.Done():
        fmt.Printf("⏰ %s cancelada por timeout: %v\n", nombre, ctx.Err())
        return ctx.Err()
    }
}

func ejemploTimeoutBasico() {
    fmt.Println("⏰ Ejemplo: Timeout Básico")
    fmt.Println("=========================")
    
    operaciones := []struct {
        nombre   string
        duracion time.Duration
        timeout  time.Duration
    }{
        {"Operación Rápida", 1 * time.Second, 3 * time.Second},
        {"Operación Lenta", 5 * time.Second, 3 * time.Second},
        {"Operación Media", 2 * time.Second, 3 * time.Second},
    }
    
    for _, op := range operaciones {
        // Crear context con timeout
        ctx, cancel := context.WithTimeout(context.Background(), op.timeout)
        
        fmt.Printf("\n🎯 Ejecutando %s con timeout de %v\n", op.nombre, op.timeout)
        err := operacionLenta(ctx, op.nombre, op.duracion)
        
        if err != nil {
            fmt.Printf("❌ Error: %v\n", err)
        }
        
        cancel() // Siempre cancelar para liberar recursos
    }
    
    fmt.Println()
}
```

### 2. 📅 Deadline Específico

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func procesamientoPorLotes(ctx context.Context, batchSize int) {
    fmt.Printf("📦 Iniciando procesamiento de %d elementos\n", batchSize)
    
    for i := 1; i <= batchSize; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("⏰ Procesamiento interrumpido en elemento %d: %v\n", i, ctx.Err())
            return
        default:
            // Simular procesamiento de elemento
            time.Sleep(200 * time.Millisecond)
            fmt.Printf("  ✅ Elemento %d/%d procesado\n", i, batchSize)
        }
    }
    
    fmt.Printf("🎉 Procesamiento por lotes completado\n")
}

func ejemploDeadlineEspecifico() {
    fmt.Println("📅 Ejemplo: Deadline Específico")
    fmt.Println("==============================")
    
    // Deadline: procesar hasta las 5 segundos desde ahora
    deadline := time.Now().Add(5 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    
    fmt.Printf("⏰ Deadline establecido: %v\n", deadline.Format("15:04:05"))
    fmt.Printf("🕐 Tiempo actual: %v\n", time.Now().Format("15:04:05"))
    
    // Intentar procesar 30 elementos (más de lo que debería dar tiempo)
    procesamientoPorLotes(ctx, 30)
    
    // Mostrar información del deadline
    if deadline, ok := ctx.Deadline(); ok {
        timeLeft := time.Until(deadline)
        fmt.Printf("⏱️ Tiempo restante hasta deadline: %v\n", timeLeft)
    }
    
    fmt.Println()
}
```

### 3. 🔄 Timeouts Escalonados

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func servicioWeb(ctx context.Context, url string, timeout time.Duration) error {
    // Crear sub-context con timeout específico
    reqCtx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    
    fmt.Printf("🌐 Consultando %s (timeout: %v)\n", url, timeout)
    
    // Simular diferentes duraciones de respuesta
    var duracion time.Duration
    switch url {
    case "api-rapida.com":
        duracion = 500 * time.Millisecond
    case "api-lenta.com":
        duracion = 3 * time.Second
    case "api-muy-lenta.com":
        duracion = 8 * time.Second
    }
    
    select {
    case <-time.After(duracion):
        fmt.Printf("  ✅ %s respondió en %v\n", url, duracion)
        return nil
    case <-reqCtx.Done():
        fmt.Printf("  ⏰ %s timeout después de %v: %v\n", url, timeout, reqCtx.Err())
        return reqCtx.Err()
    }
}

func ejemploTimeoutsEscalonados() {
    fmt.Println("🔄 Ejemplo: Timeouts Escalonados")
    fmt.Println("===============================")
    
    // Context principal con timeout general
    mainCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    servicios := []struct {
        url     string
        timeout time.Duration
    }{
        {"api-rapida.com", 1 * time.Second},
        {"api-lenta.com", 2 * time.Second},
        {"api-muy-lenta.com", 4 * time.Second},
    }
    
    fmt.Println("🎯 Consultando servicios con timeouts escalonados:")
    
    for _, servicio := range servicios {
        err := servicioWeb(mainCtx, servicio.url, servicio.timeout)
        if err != nil {
            fmt.Printf("❌ Error consultando %s: %v\n", servicio.url, err)
        }
        
        // Verificar si el context principal ya expiró
        select {
        case <-mainCtx.Done():
            fmt.Printf("🚫 Context principal expirado: %v\n", mainCtx.Err())
            return
        default:
            // Continuar con el siguiente servicio
        }
    }
    
    fmt.Println("🎉 Todas las consultas completadas\n")
}
```

---

## 📦 Propagación de Valores

### 1. 🏷️ Context Values Básico

```go
package main

import (
    "context"
    "fmt"
)

// Tipos para keys de context - buena práctica
type contextKey string

const (
    UserIDKey     contextKey = "userID"
    RequestIDKey  contextKey = "requestID"
    CorrelationIDKey contextKey = "correlationID"
)

func authenticate(ctx context.Context, token string) context.Context {
    // Simular autenticación
    userID := "user-" + token[len(token)-4:] // Últimos 4 caracteres
    
    // Agregar userID al context
    return context.WithValue(ctx, UserIDKey, userID)
}

func authorize(ctx context.Context, resource string) bool {
    userID, ok := ctx.Value(UserIDKey).(string)
    if !ok {
        fmt.Printf("❌ No se encontró userID en el context\n")
        return false
    }
    
    // Simular autorización
    fmt.Printf("🔐 Autorizando %s para acceder a %s\n", userID, resource)
    return userID != "user-deny" // Denegar usuarios específicos
}

func logRequest(ctx context.Context, action string) {
    userID := ctx.Value(UserIDKey)
    requestID := ctx.Value(RequestIDKey)
    correlationID := ctx.Value(CorrelationIDKey)
    
    fmt.Printf("📝 [%v] [%v] [%v] %s\n", requestID, correlationID, userID, action)
}

func handleRequest(ctx context.Context, resource string) {
    logRequest(ctx, fmt.Sprintf("Iniciando acceso a %s", resource))
    
    // Verificar autorización
    if !authorize(ctx, resource) {
        logRequest(ctx, "Acceso denegado")
        return
    }
    
    // Simular procesamiento
    logRequest(ctx, "Procesando request")
    
    // Completar
    logRequest(ctx, "Request completado exitosamente")
}

func ejemploContextValues() {
    fmt.Println("📦 Ejemplo: Propagación de Valores")
    fmt.Println("=================================")
    
    // Context base
    ctx := context.Background()
    
    // Agregar request ID y correlation ID
    ctx = context.WithValue(ctx, RequestIDKey, "req-12345")
    ctx = context.WithValue(ctx, CorrelationIDKey, "corr-abcde")
    
    // Simular requests con diferentes tokens
    tokens := []string{"valid-user-1234", "valid-admin-5678", "valid-deny-9999"}
    resources := []string{"documents", "admin-panel", "user-profile"}
    
    for i, token := range tokens {
        fmt.Printf("\n🔑 Procesando request con token: %s\n", token)
        
        // Autenticar y obtener context con userID
        authCtx := authenticate(ctx, token)
        
        // Procesar request
        handleRequest(authCtx, resources[i])
    }
    
    fmt.Println()
}
```

### 2. 🏗️ Context Builder Pattern

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// RequestContext es un builder para crear contexts de request
type RequestContext struct {
    ctx context.Context
}

func NewRequestContext() *RequestContext {
    return &RequestContext{
        ctx: context.Background(),
    }
}

func (r *RequestContext) WithUserID(userID string) *RequestContext {
    r.ctx = context.WithValue(r.ctx, UserIDKey, userID)
    return r
}

func (r *RequestContext) WithRequestID(requestID string) *RequestContext {
    r.ctx = context.WithValue(r.ctx, RequestIDKey, requestID)
    return r
}

func (r *RequestContext) WithTimeout(timeout time.Duration) *RequestContext {
    var cancel context.CancelFunc
    r.ctx, cancel = context.WithTimeout(r.ctx, timeout)
    
    // Programar cancelación automática
    go func() {
        <-r.ctx.Done()
        cancel()
    }()
    
    return r
}

func (r *RequestContext) WithCorrelationID(correlationID string) *RequestContext {
    r.ctx = context.WithValue(r.ctx, CorrelationIDKey, correlationID)
    return r
}

func (r *RequestContext) Build() context.Context {
    return r.ctx
}

// Helper para extraer valores del context
func GetUserID(ctx context.Context) (string, bool) {
    userID, ok := ctx.Value(UserIDKey).(string)
    return userID, ok
}

func GetRequestID(ctx context.Context) (string, bool) {
    requestID, ok := ctx.Value(RequestIDKey).(string)
    return requestID, ok
}

func processBusinessLogic(ctx context.Context, operation string) error {
    userID, _ := GetUserID(ctx)
    requestID, _ := GetRequestID(ctx)
    
    fmt.Printf("🔄 [%s] Usuario %s ejecutando: %s\n", requestID, userID, operation)
    
    // Simular trabajo que puede ser cancelado
    select {
    case <-time.After(1 * time.Second):
        fmt.Printf("✅ [%s] Operación %s completada\n", requestID, operation)
        return nil
    case <-ctx.Done():
        fmt.Printf("❌ [%s] Operación %s cancelada: %v\n", requestID, operation, ctx.Err())
        return ctx.Err()
    }
}

func ejemploContextBuilder() {
    fmt.Println("🏗️ Ejemplo: Context Builder Pattern")
    fmt.Println("==================================")
    
    // Crear diferentes contexts usando el builder
    contexts := []struct {
        name string
        ctx  context.Context
    }{
        {
            name: "Request Rápido",
            ctx: NewRequestContext().
                WithUserID("user-123").
                WithRequestID("req-fast").
                WithCorrelationID("corr-001").
                WithTimeout(3 * time.Second).
                Build(),
        },
        {
            name: "Request Lento",
            ctx: NewRequestContext().
                WithUserID("user-456").
                WithRequestID("req-slow").
                WithCorrelationID("corr-002").
                WithTimeout(500 * time.Millisecond). // Timeout corto
                Build(),
        },
    }
    
    for _, contextInfo := range contexts {
        fmt.Printf("\n🎯 Procesando %s\n", contextInfo.name)
        err := processBusinessLogic(contextInfo.ctx, "calculate-report")
        if err != nil {
            fmt.Printf("⚠️ Error en %s: %v\n", contextInfo.name, err)
        }
    }
    
    fmt.Println()
}
```

---

## 🏗️ Patrones Avanzados con Context

### 1. 🔄 Context Middleware

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// Middleware type
type Middleware func(context.Context, func(context.Context)) context.Context

// Logging middleware
func LoggingMiddleware(ctx context.Context, next func(context.Context)) context.Context {
    start := time.Now()
    requestID, _ := GetRequestID(ctx)
    
    fmt.Printf("🔍 [%s] Request iniciado\n", requestID)
    
    // Ejecutar siguiente función
    next(ctx)
    
    duration := time.Since(start)
    fmt.Printf("📊 [%s] Request completado en %v\n", requestID, duration)
    
    return ctx
}

// Authentication middleware
func AuthenticationMiddleware(ctx context.Context, next func(context.Context)) context.Context {
    requestID, _ := GetRequestID(ctx)
    
    // Simular verificación de autenticación
    if userID, ok := GetUserID(ctx); ok {
        fmt.Printf("🔐 [%s] Usuario autenticado: %s\n", requestID, userID)
        next(ctx)
    } else {
        fmt.Printf("❌ [%s] Usuario no autenticado\n", requestID)
    }
    
    return ctx
}

// Rate limiting middleware
func RateLimitingMiddleware(ctx context.Context, next func(context.Context)) context.Context {
    requestID, _ := GetRequestID(ctx)
    userID, _ := GetUserID(ctx)
    
    // Simular rate limiting
    if userID == "user-blocked" {
        fmt.Printf("🚫 [%s] Rate limit excedido para %s\n", requestID, userID)
        return ctx
    }
    
    fmt.Printf("✅ [%s] Rate limit OK para %s\n", requestID, userID)
    next(ctx)
    
    return ctx
}

// Request handler final
func handleBusinessRequest(ctx context.Context) {
    requestID, _ := GetRequestID(ctx)
    userID, _ := GetUserID(ctx)
    
    fmt.Printf("💼 [%s] Ejecutando lógica de negocio para %s\n", requestID, userID)
    
    // Simular trabajo
    select {
    case <-time.After(500 * time.Millisecond):
        fmt.Printf("🎉 [%s] Lógica de negocio completada\n", requestID)
    case <-ctx.Done():
        fmt.Printf("⏰ [%s] Lógica de negocio cancelada: %v\n", requestID, ctx.Err())
    }
}

// Pipeline de middlewares
func ProcessRequest(ctx context.Context, middlewares []Middleware, handler func(context.Context)) {
    // Crear una función que aplique todos los middlewares
    var processedHandler func(context.Context) = handler
    
    // Aplicar middlewares en orden inverso (último primero)
    for i := len(middlewares) - 1; i >= 0; i-- {
        middleware := middlewares[i]
        currentHandler := processedHandler
        
        processedHandler = func(ctx context.Context) {
            middleware(ctx, currentHandler)
        }
    }
    
    // Ejecutar el pipeline completo
    processedHandler(ctx)
}

func ejemploContextMiddleware() {
    fmt.Println("🔄 Ejemplo: Context Middleware")
    fmt.Println("=============================")
    
    // Definir middlewares
    middlewares := []Middleware{
        LoggingMiddleware,
        AuthenticationMiddleware,
        RateLimitingMiddleware,
    }
    
    // Procesar diferentes requests
    requests := []struct {
        name   string
        userID string
        reqID  string
    }{
        {"Request válido", "user-123", "req-001"},
        {"Request sin auth", "", "req-002"},
        {"Request rate limited", "user-blocked", "req-003"},
    }
    
    for _, req := range requests {
        fmt.Printf("\n🎯 Procesando: %s\n", req.name)
        
        ctx := NewRequestContext().
            WithUserID(req.userID).
            WithRequestID(req.reqID).
            WithTimeout(2 * time.Second).
            Build()
        
        ProcessRequest(ctx, middlewares, handleBusinessRequest)
    }
    
    fmt.Println()
}
```

### 2. 🔀 Context Composition

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Composite Context que combina múltiples contexts
type CompositeContext struct {
    contexts []context.Context
    done     chan struct{}
    err      error
    once     sync.Once
}

func NewCompositeContext(contexts ...context.Context) *CompositeContext {
    cc := &CompositeContext{
        contexts: contexts,
        done:     make(chan struct{}),
    }
    
    // Monitorear todos los contexts
    for _, ctx := range contexts {
        go func(c context.Context) {
            <-c.Done()
            cc.cancel(c.Err())
        }(ctx)
    }
    
    return cc
}

func (cc *CompositeContext) cancel(err error) {
    cc.once.Do(func() {
        cc.err = err
        close(cc.done)
    })
}

func (cc *CompositeContext) Done() <-chan struct{} {
    return cc.done
}

func (cc *CompositeContext) Err() error {
    select {
    case <-cc.done:
        return cc.err
    default:
        return nil
    }
}

func (cc *CompositeContext) Deadline() (deadline time.Time, ok bool) {
    // Retornar el deadline más cercano
    var earliest time.Time
    var hasDeadline bool
    
    for _, ctx := range cc.contexts {
        if d, ok := ctx.Deadline(); ok {
            if !hasDeadline || d.Before(earliest) {
                earliest = d
                hasDeadline = true
            }
        }
    }
    
    return earliest, hasDeadline
}

func (cc *CompositeContext) Value(key interface{}) interface{} {
    // Buscar el valor en todos los contexts
    for _, ctx := range cc.contexts {
        if value := ctx.Value(key); value != nil {
            return value
        }
    }
    return nil
}

func operacionCompleja(ctx context.Context, nombre string) {
    ticker := time.NewTicker(300 * time.Millisecond)
    defer ticker.Stop()
    
    iteracion := 0
    
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("🔴 %s cancelada en iteración %d: %v\n", nombre, iteracion, ctx.Err())
            return
        case <-ticker.C:
            iteracion++
            fmt.Printf("⚡ %s - iteración %d\n", nombre, iteracion)
            
            if iteracion >= 10 {
                fmt.Printf("✅ %s completada exitosamente\n", nombre)
                return
            }
        }
    }
}

func ejemploContextComposition() {
    fmt.Println("🔀 Ejemplo: Context Composition")
    fmt.Println("===============================")
    
    // Crear diferentes contexts con diferentes propósitos
    userCtx := context.WithValue(context.Background(), UserIDKey, "user-123")
    timeoutCtx, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel1()
    
    cancelableCtx, cancel2 := context.WithCancel(context.Background())
    defer cancel2()
    
    // Crear composite context
    compositeCtx := NewCompositeContext(userCtx, timeoutCtx, cancelableCtx)
    
    fmt.Println("🎯 Iniciando operación con context compuesto")
    fmt.Printf("👤 User ID: %v\n", compositeCtx.Value(UserIDKey))
    
    if deadline, ok := compositeCtx.Deadline(); ok {
        fmt.Printf("⏰ Deadline: %v\n", deadline.Format("15:04:05"))
    }
    
    // Lanzar operación
    go operacionCompleja(compositeCtx, "Operación Compuesta")
    
    // Simular cancelación manual después de 1.5 segundos
    go func() {
        time.Sleep(1500 * time.Millisecond)
        fmt.Println("📤 Enviando cancelación manual...")
        cancel2()
    }()
    
    // Esperar a que termine
    <-compositeCtx.Done()
    fmt.Printf("🏁 Context compuesto terminado: %v\n", compositeCtx.Err())
    
    time.Sleep(500 * time.Millisecond)
    fmt.Println()
}
```

---

## 🛡️ Mejores Prácticas con Context

### 1. 📋 Reglas de Oro

```go
// ✅ CORRECTO: Context como primer parámetro
func ProcessData(ctx context.Context, data []byte) error {
    return nil
}

// ❌ INCORRECTO: Context no es el primer parámetro
func ProcessData(data []byte, ctx context.Context) error {
    return nil
}

// ✅ CORRECTO: No almacenar context en structs
type Service struct {
    client HTTPClient
}

func (s *Service) FetchData(ctx context.Context) error {
    return s.client.Get(ctx, "/api/data")
}

// ❌ INCORRECTO: Context almacenado en struct
type BadService struct {
    ctx    context.Context // ¡No hagas esto!
    client HTTPClient
}

// ✅ CORRECTO: Usar context.TODO() durante desarrollo
func WorkInProgress(ctx context.Context) {
    if ctx == nil {
        ctx = context.TODO() // Temporal durante desarrollo
    }
    // ... código
}

// ✅ CORRECTO: Keys tipadas para valores
type key string
const userKey key = "user"

func SetUser(ctx context.Context, user User) context.Context {
    return context.WithValue(ctx, userKey, user)
}

// ❌ INCORRECTO: String keys directos
func BadSetUser(ctx context.Context, user User) context.Context {
    return context.WithValue(ctx, "user", user) // Propenso a colisiones
}
```

### 2. 🧪 Testing con Context

```go
package main

import (
    "context"
    "fmt"
    "testing"
    "time"
)

// Función que queremos testear
func ProcessWithTimeout(ctx context.Context, data string) (string, error) {
    // Simular procesamiento que puede tardar
    select {
    case <-time.After(100 * time.Millisecond):
        return "processed: " + data, nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

// Test de caso exitoso
func TestProcessWithTimeout_Success() {
    ctx := context.Background()
    result, err := ProcessWithTimeout(ctx, "test-data")
    
    if err != nil {
        fmt.Printf("❌ Test fallido: %v\n", err)
        return
    }
    
    expected := "processed: test-data"
    if result != expected {
        fmt.Printf("❌ Resultado incorrecto. Esperado: %s, Obtenido: %s\n", expected, result)
        return
    }
    
    fmt.Println("✅ TestProcessWithTimeout_Success pasó")
}

// Test de timeout
func TestProcessWithTimeout_Timeout() {
    // Context con timeout muy corto
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()
    
    result, err := ProcessWithTimeout(ctx, "test-data")
    
    if err == nil {
        fmt.Printf("❌ Test fallido: esperaba error de timeout\n")
        return
    }
    
    if err != context.DeadlineExceeded {
        fmt.Printf("❌ Test fallido: error incorrecto. Esperado: %v, Obtenido: %v\n", 
                   context.DeadlineExceeded, err)
        return
    }
    
    if result != "" {
        fmt.Printf("❌ Test fallido: resultado debería estar vacío\n")
        return
    }
    
    fmt.Println("✅ TestProcessWithTimeout_Timeout pasó")
}

// Test de cancelación
func TestProcessWithTimeout_Cancellation() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // Cancelar inmediatamente
    cancel()
    
    result, err := ProcessWithTimeout(ctx, "test-data")
    
    if err == nil {
        fmt.Printf("❌ Test fallido: esperaba error de cancelación\n")
        return
    }
    
    if err != context.Canceled {
        fmt.Printf("❌ Test fallido: error incorrecto. Esperado: %v, Obtenido: %v\n", 
                   context.Canceled, err)
        return
    }
    
    fmt.Println("✅ TestProcessWithTimeout_Cancellation pasó")
}

// Helper para crear context de testing
func TestContext(timeout time.Duration) (context.Context, context.CancelFunc) {
    if timeout <= 0 {
        return context.WithCancel(context.Background())
    }
    return context.WithTimeout(context.Background(), timeout)
}

func ejemploTestingContext() {
    fmt.Println("🧪 Ejemplo: Testing con Context")
    fmt.Println("==============================")
    
    TestProcessWithTimeout_Success()
    TestProcessWithTimeout_Timeout()
    TestProcessWithTimeout_Cancellation()
    
    fmt.Println()
}
```

---

## ⚠️ Errores Comunes y Cómo Evitarlos

### 1. 🚫 Context Leaks

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// ❌ MAL: No cancelar context
func badExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // ¡Olvidamos cancel()! - LEAK
    
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("Timeout alcanzado")
        }
    }()
}

// ✅ BIEN: Siempre cancelar
func goodExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // ¡Siempre cancelar!
    
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("Timeout alcanzado")
        }
    }()
}

// ❌ MAL: No propagar cancelación
func badChaining() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Crear nuevo context sin heredar cancelación
    newCtx := context.Background() // ¡Perdimos la cancelación!
    
    go func() {
        select {
        case <-newCtx.Done(): // Nunca se cancelará
            fmt.Println("Nunca se ejecuta")
        }
    }()
}

// ✅ BIEN: Heredar cancelación
func goodChaining() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Crear context que hereda cancelación
    newCtx, newCancel := context.WithCancel(ctx)
    defer newCancel()
    
    go func() {
        select {
        case <-newCtx.Done(): // Se cancelará cuando ctx se cancele
            fmt.Println("Cancelado correctamente")
        }
    }()
}
```

### 2. 💾 Context Values Anti-patterns

```go
package main

import (
    "context"
    "fmt"
)

// ❌ MAL: Usar context para pasar datos de configuración
func badConfiguration(ctx context.Context) {
    // ¡NO hagas esto!
    dbURL := ctx.Value("database_url").(string)
    apiKey := ctx.Value("api_key").(string)
    
    fmt.Printf("DB: %s, API: %s\n", dbURL, apiKey)
}

// ✅ BIEN: Configuración como parámetros o struct
type Config struct {
    DatabaseURL string
    APIKey      string
}

func goodConfiguration(ctx context.Context, config Config) {
    fmt.Printf("DB: %s, API: %s\n", config.DatabaseURL, config.APIKey)
}

// ❌ MAL: Valores grandes en context
func badLargeValues(ctx context.Context) {
    // ¡No pongas objetos grandes en context!
    bigData := make([]byte, 1024*1024) // 1MB
    ctx = context.WithValue(ctx, "big_data", bigData)
}

// ✅ BIEN: Solo metadata pequeña
func goodSmallValues(ctx context.Context) {
    // Solo IDs, tokens, flags pequeños
    ctx = context.WithValue(ctx, UserIDKey, "user-123")
    ctx = context.WithValue(ctx, "trace_id", "abc123")
}

// ❌ MAL: Context como miembro de struct
type BadService struct {
    ctx context.Context // ¡No hagas esto!
}

// ✅ BIEN: Context como parámetro
type GoodService struct {
    client HTTPClient
}

func (s *GoodService) DoWork(ctx context.Context) error {
    return nil
}
```

---

## 🚀 Context en el Mundo Real

### 1. 🌐 HTTP Server con Context

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type HTTPClient interface {
    Get(ctx context.Context, url string) (*http.Response, error)
}

// Simulación de servicio web
func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
    // El context viene del request HTTP
    ctx := r.Context()
    
    // Agregar timeout específico para esta operación
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    // Agregar request ID para tracing
    requestID := r.Header.Get("X-Request-ID")
    if requestID == "" {
        requestID = generateRequestID()
    }
    ctx = context.WithValue(ctx, RequestIDKey, requestID)
    
    // Procesar request
    result, err := processAPIRequest(ctx, r.URL.Query().Get("data"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Responder
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "result":     result,
        "request_id": requestID,
    })
}

func processAPIRequest(ctx context.Context, data string) (string, error) {
    requestID, _ := ctx.Value(RequestIDKey).(string)
    
    fmt.Printf("🔄 [%s] Procesando: %s\n", requestID, data)
    
    // Simular llamada a base de datos
    dbResult, err := queryDatabase(ctx, data)
    if err != nil {
        return "", err
    }
    
    // Simular llamada a servicio externo
    apiResult, err := callExternalAPI(ctx, dbResult)
    if err != nil {
        return "", err
    }
    
    fmt.Printf("✅ [%s] Procesamiento completado\n", requestID)
    return apiResult, nil
}

func queryDatabase(ctx context.Context, query string) (string, error) {
    requestID, _ := ctx.Value(RequestIDKey).(string)
    
    select {
    case <-time.After(100 * time.Millisecond): // Simular query
        fmt.Printf("💾 [%s] Database query completada\n", requestID)
        return "db_result_" + query, nil
    case <-ctx.Done():
        fmt.Printf("❌ [%s] Database query cancelada: %v\n", requestID, ctx.Err())
        return "", ctx.Err()
    }
}

func callExternalAPI(ctx context.Context, data string) (string, error) {
    requestID, _ := ctx.Value(RequestIDKey).(string)
    
    select {
    case <-time.After(200 * time.Millisecond): // Simular API call
        fmt.Printf("🌐 [%s] External API call completada\n", requestID)
        return "api_result_" + data, nil
    case <-ctx.Done():
        fmt.Printf("❌ [%s] External API call cancelada: %v\n", requestID, ctx.Err())
        return "", ctx.Err()
    }
}

func generateRequestID() string {
    return fmt.Sprintf("req-%d", time.Now().UnixNano()%10000)
}

func ejemploHTTPServer() {
    fmt.Println("🌐 Ejemplo: HTTP Server con Context")
    fmt.Println("==================================")
    
    // Simular algunos requests
    requests := []string{"user_data", "order_info", "product_catalog"}
    
    for _, req := range requests {
        fmt.Printf("\n📥 Simulando request: %s\n", req)
        
        // Crear request simulado
        httpReq, _ := http.NewRequest("GET", "/api/data?data="+req, nil)
        httpReq.Header.Set("X-Request-ID", generateRequestID())
        
        // Procesar con timeout del cliente
        ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
        httpReq = httpReq.WithContext(ctx)
        
        // Simular procesamiento
        result, err := processAPIRequest(httpReq.Context(), req)
        if err != nil {
            fmt.Printf("❌ Error: %v\n", err)
        } else {
            fmt.Printf("📤 Response: %s\n", result)
        }
        
        cancel()
    }
    
    fmt.Println()
}
```

### 2. 🔄 Worker Pool con Context

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID   int
    Data string
}

type Result struct {
    JobID  int
    Result string
    Error  error
}

type WorkerPool struct {
    workerCount int
    jobs        chan Job
    results     chan Result
    ctx         context.Context
    cancel      context.CancelFunc
    wg          sync.WaitGroup
}

func NewWorkerPool(workerCount, bufferSize int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    
    return &WorkerPool{
        workerCount: workerCount,
        jobs:        make(chan Job, bufferSize),
        results:     make(chan Result, bufferSize),
        ctx:         ctx,
        cancel:      cancel,
    }
}

func (wp *WorkerPool) Start() {
    fmt.Printf("🏭 Iniciando worker pool con %d workers\n", wp.workerCount)
    
    for i := 1; i <= wp.workerCount; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    fmt.Printf("👷 Worker %d iniciado\n", id)
    
    for {
        select {
        case job, ok := <-wp.jobs:
            if !ok {
                fmt.Printf("👷 Worker %d: canal cerrado, terminando\n", id)
                return
            }
            
            result := wp.processJob(id, job)
            
            select {
            case wp.results <- result:
            case <-wp.ctx.Done():
                fmt.Printf("👷 Worker %d: cancelado durante envío de resultado\n", id)
                return
            }
            
        case <-wp.ctx.Done():
            fmt.Printf("👷 Worker %d: cancelado\n", id)
            return
        }
    }
}

func (wp *WorkerPool) processJob(workerID int, job Job) Result {
    // Crear context con timeout para este job específico
    jobCtx, cancel := context.WithTimeout(wp.ctx, 2*time.Second)
    defer cancel()
    
    fmt.Printf("⚡ Worker %d procesando job %d: %s\n", workerID, job.ID, job.Data)
    
    // Simular procesamiento
    select {
    case <-time.After(time.Duration(job.ID*100) * time.Millisecond):
        result := fmt.Sprintf("processed_%s_by_worker_%d", job.Data, workerID)
        fmt.Printf("✅ Worker %d completó job %d\n", workerID, job.ID)
        return Result{JobID: job.ID, Result: result, Error: nil}
        
    case <-jobCtx.Done():
        fmt.Printf("⏰ Worker %d: job %d timeout\n", workerID, job.ID)
        return Result{JobID: job.ID, Result: "", Error: jobCtx.Err()}
    }
}

func (wp *WorkerPool) Submit(job Job) {
    select {
    case wp.jobs <- job:
    case <-wp.ctx.Done():
        fmt.Printf("❌ No se pudo enviar job %d: worker pool cancelado\n", job.ID)
    }
}

func (wp *WorkerPool) Close() {
    fmt.Println("🛑 Cerrando worker pool...")
    
    close(wp.jobs)
    wp.wg.Wait()
    close(wp.results)
    
    fmt.Println("✅ Worker pool cerrado")
}

func (wp *WorkerPool) Shutdown(timeout time.Duration) {
    fmt.Printf("🛑 Iniciando shutdown con timeout de %v\n", timeout)
    
    // Cerrar jobs channel
    close(wp.jobs)
    
    // Esperar con timeout
    done := make(chan struct{})
    go func() {
        wp.wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        fmt.Println("✅ Shutdown completado correctamente")
    case <-time.After(timeout):
        fmt.Println("⏰ Shutdown timeout, forzando cancelación")
        wp.cancel()
        wp.wg.Wait()
    }
    
    close(wp.results)
}

func ejemploWorkerPoolContext() {
    fmt.Println("🔄 Ejemplo: Worker Pool con Context")
    fmt.Println("==================================")
    
    // Crear worker pool
    pool := NewWorkerPool(3, 10)
    pool.Start()
    
    // Enviar jobs
    jobs := []Job{
        {ID: 1, Data: "task_fast"},
        {ID: 2, Data: "task_medium"},
        {ID: 3, Data: "task_slow"},
        {ID: 4, Data: "task_very_slow"},
        {ID: 5, Data: "task_quick"},
    }
    
    // Goroutine para recoger resultados
    go func() {
        for result := range pool.results {
            if result.Error != nil {
                fmt.Printf("❌ Job %d falló: %v\n", result.JobID, result.Error)
            } else {
                fmt.Printf("📊 Job %d resultado: %s\n", result.JobID, result.Result)
            }
        }
    }()
    
    // Enviar jobs
    fmt.Println("📤 Enviando jobs...")
    for _, job := range jobs {
        pool.Submit(job)
    }
    
    // Esperar un poco y hacer shutdown
    time.Sleep(3 * time.Second)
    pool.Shutdown(1 * time.Second)
    
    fmt.Println()
}
```

---

## 🎯 Resumen de la Lección

### ✅ Conceptos Clave Aprendidos

1. **🎯 Context Fundamentals**: Qué es y por qué es esencial
2. **🚫 Cancelación**: Coordinación elegante de operaciones
3. **⏰ Timeouts y Deadlines**: Control temporal de operaciones
4. **📦 Value Propagation**: Paso de metadata a través del call stack
5. **🏗️ Patrones Avanzados**: Middleware, composition, builder patterns
6. **🛡️ Best Practices**: Reglas de oro y errores comunes
7. **🧪 Testing**: Cómo testear código que usa Context
8. **🚀 Real World**: Aplicaciones prácticas en HTTP servers y worker pools

### 🏆 Logros Desbloqueados

- [ ] 🥇 **Context Novice**: Primer Context exitoso
- [ ] 🥈 **Timeout Master**: Dominio de timeouts y deadlines
- [ ] 🥉 **Cancellation Expert**: Cancelación coordinada
- [ ] 🏅 **Value Propagator**: Paso eficiente de metadata
- [ ] 🎖️ **Pattern Architect**: Patrones avanzados implementados
- [ ] 🏆 **Context Wizard**: Aplicaciones del mundo real

### 📚 Próximos Pasos

En la **Lección 16: Error Handling**, aprenderemos:
- Filosofía de manejo de errores en Go
- Patrones avanzados de error handling
- Error wrapping y unwrapping
- Custom errors y error types
- Logging y observabilidad

---

**🎉 ¡Felicitaciones! Has dominado el Context package de Go. Ahora puedes crear aplicaciones que manejan cancelación, timeouts y propagación de valores de forma elegante y eficiente.**

*Recuerda: "Context es el hilo conductor que conecta todas las operaciones de tu aplicación" - Úsalo sabiamente para crear sistemas robustos y observables.*
