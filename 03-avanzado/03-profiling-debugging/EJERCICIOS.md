# 🧪 Ejercicios Prácticos: Profiling & Debugging

## 🎯 Ejercicio 1: Análisis de Performance con Benchmarks

### Objetivo
Comparar diferentes implementaciones y identificar la más eficiente.

### Código Base
```go
package main

import (
    "fmt"
    "strings"
    "testing"
)

// Implementación 1: Concatenación naive
func ConcatenateNaive(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}

// Implementación 2: strings.Builder
func ConcatenateBuilder(strs []string) string {
    var builder strings.Builder
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}

// Implementación 3: strings.Join
func ConcatenateJoin(strs []string) string {
    return strings.Join(strs, "")
}

// Implementación 4: Builder con capacidad pre-asignada
func ConcatenateBuilderPrealloc(strs []string) string {
    totalLen := 0
    for _, s := range strs {
        totalLen += len(s)
    }
    
    var builder strings.Builder
    builder.Grow(totalLen)
    
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}
```

### TODO: Implementa los benchmarks
```go
func BenchmarkConcatenateNaive(b *testing.B) {
    // Tu implementación aquí
}

func BenchmarkConcatenateBuilder(b *testing.B) {
    // Tu implementación aquí
}

func BenchmarkConcatenateJoin(b *testing.B) {
    // Tu implementación aquí
}

func BenchmarkConcatenateBuilderPrealloc(b *testing.B) {
    // Tu implementación aquí
}

// Benchmark con diferentes tamaños
func BenchmarkConcatenateComparison(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    
    for _, size := range sizes {
        // Implementa benchmarks para cada tamaño
    }
}
```

### Comandos para ejecutar
```bash
# Ejecutar benchmarks
go test -bench=. -benchmem

# Generar CPU profile
go test -bench=BenchmarkConcatenateNaive -cpuprofile=cpu.prof

# Generar memory profile
go test -bench=BenchmarkConcatenateNaive -memprofile=mem.prof

# Analizar profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

---

## 🎯 Ejercicio 2: Detección y Resolución de Deadlocks

### Objetivo
Identificar, diagnosticar y resolver deadlocks en código concurrente.

### Código con Deadlock
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type BankAccount struct {
    mu      sync.Mutex
    balance int
    id      string
}

func (ba *BankAccount) Transfer(to *BankAccount, amount int) error {
    // PROBLEMA: Orden de locks inconsistente puede causar deadlock
    ba.mu.Lock()
    defer ba.mu.Unlock()
    
    time.Sleep(100 * time.Millisecond) // Simula procesamiento
    
    to.mu.Lock()
    defer to.mu.Unlock()
    
    if ba.balance < amount {
        return fmt.Errorf("insufficient funds")
    }
    
    ba.balance -= amount
    to.balance += amount
    
    fmt.Printf("Transferred %d from %s to %s\n", amount, ba.id, to.id)
    return nil
}

func main() {
    account1 := &BankAccount{balance: 1000, id: "A"}
    account2 := &BankAccount{balance: 1000, id: "B"}
    
    // Estas dos goroutines pueden generar deadlock
    go func() {
        for i := 0; i < 10; i++ {
            account1.Transfer(account2, 100)
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := 0; i < 10; i++ {
            account2.Transfer(account1, 100)
            time.Sleep(50 * time.Millisecond)
        }
    }()
    
    time.Sleep(5 * time.Second)
    fmt.Println("Finished")
}
```

### TODO: Resolver el deadlock
1. **Análisis**: Identifica por qué ocurre el deadlock
2. **Solución 1**: Implementa orden consistente de locks
3. **Solución 2**: Usa un solo mutex global
4. **Solución 3**: Implementa timeout con context

### Herramientas para debugging
```bash
# Detectar deadlock con race detector
go run -race main.go

# Debugging con Delve
dlv debug main.go
# Comandos en dlv:
# b main.main
# c
# goroutines
# goroutine 1 bt
```

---

## 🎯 Ejercicio 3: Memory Leak Detection

### Objetivo
Identificar y resolver memory leaks en una aplicación de larga duración.

### Código con Memory Leaks
```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type DataProcessor struct {
    mu        sync.RWMutex
    cache     map[string][]byte  // LEAK: nunca se limpia
    workers   []chan []byte      // LEAK: workers crecen infinitamente
    callbacks []func([]byte)     // LEAK: callbacks nunca se remueven
    stats     []ProcessingStat   // LEAK: estadísticas crecen sin límite
}

type ProcessingStat struct {
    Timestamp time.Time
    DataSize  int
    Duration  time.Duration
}

func NewDataProcessor() *DataProcessor {
    return &DataProcessor{
        cache:     make(map[string][]byte),
        workers:   make([]chan []byte, 0),
        callbacks: make([]func([]byte), 0),
        stats:     make([]ProcessingStat, 0),
    }
}

func (dp *DataProcessor) ProcessData(key string, data []byte) {
    start := time.Now()
    
    // LEAK: Cache crece infinitamente
    dp.mu.Lock()
    dp.cache[key] = make([]byte, len(data))
    copy(dp.cache[key], data)
    dp.mu.Unlock()
    
    // LEAK: Workers crecen cuando están ocupados
    var workerChan chan []byte
    for _, worker := range dp.workers {
        select {
        case worker <- data:
            workerChan = worker
            goto process
        default:
            continue
        }
    }
    
    // Crear nuevo worker si todos están ocupados
    newWorker := make(chan []byte, 100)
    dp.workers = append(dp.workers, newWorker)
    workerChan = newWorker
    
    go func(ch <-chan []byte) {
        for data := range ch {
            // Simular procesamiento
            time.Sleep(time.Duration(len(data)) * time.Microsecond)
            
            // Notificar callbacks
            dp.mu.RLock()
            for _, callback := range dp.callbacks {
                callback(data)
            }
            dp.mu.RUnlock()
        }
    }(newWorker)
    
process:
    workerChan <- data
    
    // LEAK: Estadísticas crecen sin límite
    stat := ProcessingStat{
        Timestamp: start,
        DataSize:  len(data),
        Duration:  time.Since(start),
    }
    
    dp.mu.Lock()
    dp.stats = append(dp.stats, stat)
    dp.mu.Unlock()
}

func (dp *DataProcessor) RegisterCallback(cb func([]byte)) {
    dp.mu.Lock()
    defer dp.mu.Unlock()
    
    // LEAK: Callbacks nunca se remueven
    dp.callbacks = append(dp.callbacks, cb)
}

func (dp *DataProcessor) GetStats() []ProcessingStat {
    dp.mu.RLock()
    defer dp.mu.RUnlock()
    
    // LEAK: Devuelve todo el slice, no hay límite
    result := make([]ProcessingStat, len(dp.stats))
    copy(result, dp.stats)
    return result
}

// Simulación de uso intensivo
func main() {
    processor := NewDataProcessor()
    
    // Registrar múltiples callbacks
    for i := 0; i < 100; i++ {
        processor.RegisterCallback(func(data []byte) {
            // Callback que hace algo pesado
            time.Sleep(1 * time.Millisecond)
        })
    }
    
    // Procesar datos continuamente
    go func() {
        for i := 0; ; i++ {
            key := fmt.Sprintf("data_%d", i%1000) // Rotación de keys
            data := make([]byte, 1024*1024) // 1MB por request
            
            processor.ProcessData(key, data)
            
            if i%100 == 0 {
                stats := processor.GetStats()
                fmt.Printf("Processed %d items, stats: %d\n", i, len(stats))
            }
            
            time.Sleep(10 * time.Millisecond)
        }
    }()
    
    // Simular aplicación de larga duración
    time.Sleep(1 * time.Hour) // En testing, usa menos tiempo
}
```

### TODO: Resolver los memory leaks
1. **Implementa cache con TTL y límite de tamaño**
2. **Limita el número de workers**
3. **Permite removal de callbacks**
4. **Implementa rotación de estadísticas**
5. **Añade métricas de memoria**

### Herramientas para detectar leaks
```bash
# Memory profiling
go build -o app main.go
./app &
APP_PID=$!

# Tomar snapshots periódicos
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof -base heap_base.prof heap_current.prof

# Monitoring continuo
watch -n 5 'ps aux | grep app'
```

---

## 🎯 Ejercicio 4: Goroutine Leak Detection

### Objetivo
Detectar y resolver leaks de goroutines en servicios de larga duración.

### Código con Goroutine Leaks
```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "sync"
    "time"
)

type WebCrawler struct {
    client     *http.Client
    activeJobs sync.WaitGroup
    results    chan CrawlResult
}

type CrawlResult struct {
    URL    string
    Status int
    Error  error
}

func NewWebCrawler() *WebCrawler {
    return &WebCrawler{
        client:  &http.Client{Timeout: 10 * time.Second},
        results: make(chan CrawlResult, 1000),
    }
}

func (wc *WebCrawler) CrawlURL(url string) {
    // LEAK: Goroutine sin context para cancelación
    go func() {
        defer wc.activeJobs.Done()
        
        // LEAK: No hay timeout, puede colgarse indefinidamente
        resp, err := wc.client.Get(url)
        if err != nil {
            wc.results <- CrawlResult{URL: url, Error: err}
            return
        }
        defer resp.Body.Close()
        
        wc.results <- CrawlResult{
            URL:    url,
            Status: resp.StatusCode,
        }
    }()
    
    wc.activeJobs.Add(1)
}

func (wc *WebCrawler) StartWorkers(numWorkers int) {
    // LEAK: Workers sin mecanismo de shutdown
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            for {
                select {
                case result := <-wc.results:
                    fmt.Printf("Worker %d processed: %+v\n", workerID, result)
                    // LEAK: No hay case para shutdown
                }
            }
        }(i)
    }
}

func (wc *WebCrawler) WaitForCompletion() {
    wc.activeJobs.Wait()
}

func main() {
    crawler := NewWebCrawler()
    crawler.StartWorkers(10)
    
    urls := []string{
        "http://httpbin.org/delay/1",
        "http://httpbin.org/delay/2",
        "http://httpbin.org/delay/5",
        "http://httpbin.org/status/404",
        "http://httpbin.org/status/500",
    }
    
    // Crawl URLs indefinidamente
    for {
        for _, url := range urls {
            crawler.CrawlURL(url)
        }
        time.Sleep(1 * time.Second)
    }
}
```

### TODO: Resolver los goroutine leaks
1. **Implementa context para cancelación**
2. **Añade shutdown channel para workers**
3. **Implementa graceful shutdown**
4. **Añade monitoring de goroutines**

---

## 🎯 Ejercicio 5: Real-time Performance Dashboard

### Objetivo
Crear un dashboard en tiempo real que muestre métricas de la aplicación.

### TODO: Implementa un dashboard completo
```go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "runtime"
    "sync"
    "time"
)

// TODO: Implementa estas estructuras y funcionalidades

type PerformanceDashboard struct {
    // Añade campos necesarios
}

type Metrics struct {
    // CPU, memoria, goroutines, etc.
}

func (pd *PerformanceDashboard) CollectMetrics() {
    // Implementa recolección de métricas
}

func (pd *PerformanceDashboard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Implementa endpoint para métricas
}

func (pd *PerformanceDashboard) StartCollection(ctx context.Context) {
    // Implementa loop de recolección
}

func main() {
    // TODO: Implementa servidor con dashboard
}
```

### Requisitos del Dashboard
1. **Métricas en tiempo real**: CPU, memoria, goroutines
2. **Histórico**: Últimos 60 puntos de datos
3. **Alertas**: Detectar anomalías automáticamente
4. **Interface web**: HTML con JavaScript para visualización
5. **API REST**: Endpoints para consumir métricas

---

## 🔧 Herramientas y Comandos Útiles

### pprof Commands
```bash
# CPU profiling
go tool pprof cpu.prof
(pprof) top10
(pprof) list functionName
(pprof) web

# Memory profiling
go tool pprof mem.prof
(pprof) top10 -cum
(pprof) list functionName

# Goroutine profiling
go tool pprof goroutine.prof
(pprof) traces
```

### Delve Commands
```bash
# Basic debugging
dlv debug
(dlv) b main.main
(dlv) c
(dlv) n
(dlv) s
(dlv) p variable

# Advanced debugging
(dlv) goroutines
(dlv) goroutine 5
(dlv) bt
(dlv) locals
```

### Race Detection
```bash
go run -race main.go
go test -race ./...
go build -race
```

### Benchmarking
```bash
go test -bench=.
go test -bench=. -benchmem
go test -bench=. -cpuprofile=cpu.prof
go test -bench=. -memprofile=mem.prof
```

## 📊 Criterios de Evaluación

### Ejercicio 1: Benchmarking
- [ ] Benchmarks implementados correctamente
- [ ] Profiles generados y analizados
- [ ] Optimizaciones aplicadas
- [ ] Mejoras medidas cuantitativamente

### Ejercicio 2: Deadlock Resolution
- [ ] Deadlock identificado correctamente
- [ ] Múltiples soluciones implementadas
- [ ] Código sin race conditions
- [ ] Performance preservada

### Ejercicio 3: Memory Leak Detection
- [ ] Leaks identificados todos
- [ ] Soluciones implementadas
- [ ] Memory usage estabilizado
- [ ] Monitoring añadido

### Ejercicio 4: Goroutine Leak Resolution
- [ ] Leaks resueltos
- [ ] Graceful shutdown implementado
- [ ] Context usado correctamente
- [ ] Resource cleanup completo

### Ejercicio 5: Performance Dashboard
- [ ] Dashboard funcional
- [ ] Métricas precisas
- [ ] Interface intuitiva
- [ ] Alertas configurables
