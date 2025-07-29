# 🚀 Lección 01: Performance Optimization
## *El Arte de Hacer que Go Vuele*

> *"La optimización prematura es la raíz de todo mal, pero no optimizar cuando es necesario es peor"* - Donald Knuth (adaptado)

### 🎯 ¿Qué Aprenderás Hoy?

Al final de esta lección, serás capaz de:
- ⚡ **Optimizar código Go** hasta conseguir mejoras del 300-500%
- 🔍 **Identificar cuellos de botella** usando profiling avanzado
- 📊 **Medir performance** de manera científica y reproducible
- 🧠 **Aplicar técnicas** de optimización de nivel enterprise
- 🏗️ **Diseñar algoritmos** cache-friendly y memory-efficient
- ⚙️ **Aprovechar características** específicas del runtime de Go

### 🧠 Analogías para Entender Performance

#### 🏎️ **Tu Código es como un Auto de Carrera**

```mermaid
graph LR
    A[🏎️ Motor<br/>Algoritmo Core] --> B[⚙️ Transmisión<br/>Data Structures]
    B --> C[🛞 Neumáticos<br/>Memory Access]
    C --> D[🏁 Finish Line<br/>Usuario Final]
    
    style A fill:#ff6b6b,color:#fff
    style B fill:#4ecdc4,color:#fff
    style C fill:#45b7d1,color:#fff
    style D fill:#2ed573,color:#fff
```

- **🏎️ Motor (Algoritmo)**: La lógica central que mueve todo
- **⚙️ Transmisión (Estructuras)**: Cómo transferimos y organizamos datos
- **🛞 Neumáticos (Memoria)**: El contacto directo con el hardware
- **🏁 Meta**: La experiencia final del usuario

#### 🏭 **Performance como una Fábrica**

> *Imagina tu aplicación como una fábrica. Cada función es una estación de trabajo, cada variable es material, y el CPU es la energía que mueve todo.*

- **🔍 Profiling = Auditoría Industrial**: Encontrar dónde se pierde tiempo
- **⚡ Optimization = Lean Manufacturing**: Eliminar desperdicios
- **📊 Benchmarking = Control de Calidad**: Medir resultados consistentemente

## 📚 Teoría Fundamental

### 🎯 **Los 4 Pilares de Performance en Go**

#### 1. 🧮 **Computational Efficiency**
*"Hacer menos trabajo, más inteligentemente"*

```go
// 🐌 Naive approach - O(n²)
func FindDuplicatesNaive(slice []int) []int {
    var duplicates []int
    for i := 0; i < len(slice); i++ {
        for j := i + 1; j < len(slice); j++ {
            if slice[i] == slice[j] {
                duplicates = append(duplicates, slice[i])
                break
            }
        }
    }
    return duplicates
}

// 🚀 Optimized approach - O(n)
func FindDuplicatesOptimized(slice []int) []int {
    seen := make(map[int]bool)
    duplicates := make(map[int]bool)
    var result []int
    
    for _, val := range slice {
        if seen[val] {
            if !duplicates[val] {
                result = append(result, val)
                duplicates[val] = true
            }
        } else {
            seen[val] = true
        }
    }
    return result
}
```

#### 2. 🧠 **Memory Efficiency**
*"Cada byte cuenta"*

```go
// 🐌 Memory-heavy approach
type Person struct {
    Name        string    // 16 bytes (8+8)
    Age         int       // 8 bytes
    Email       string    // 16 bytes (8+8)
    IsActive    bool      // 1 byte
    // Padding: 7 bytes
    // Total: 48 bytes per person
}

// 🚀 Memory-optimized approach
type PersonOptimized struct {
    Name     string // 16 bytes
    Email    string // 16 bytes
    Age      int32  // 4 bytes
    IsActive bool   // 1 byte
    // Padding: 3 bytes
    // Total: 40 bytes per person (17% savings!)
}
```

#### 3. 🔄 **Concurrency Optimization**
*"El poder de hacer múltiples cosas bien"*

```go
// 🐌 Sequential processing
func ProcessFilesSequential(files []string) []Result {
    var results []Result
    for _, file := range files {
        result := processFile(file) // Blocking call
        results = append(results, result)
    }
    return results
}

// 🚀 Concurrent processing
func ProcessFilesConcurrent(files []string) []Result {
    const workers = 4
    jobs := make(chan string, len(files))
    results := make(chan Result, len(files))
    
    // Start workers
    for w := 0; w < workers; w++ {
        go worker(jobs, results)
    }
    
    // Send jobs
    for _, file := range files {
        jobs <- file
    }
    close(jobs)
    
    // Collect results
    var finalResults []Result
    for i := 0; i < len(files); i++ {
        finalResults = append(finalResults, <-results)
    }
    
    return finalResults
}

func worker(jobs <-chan string, results chan<- Result) {
    for file := range jobs {
        results <- processFile(file)
    }
}
```

#### 4. 🎯 **Cache Optimization**
*"Datos más cercanos = acceso más rápido"*

```go
// 🐌 Cache-unfriendly: Array of Structs
type AoS []struct {
    X, Y, Z float64
    Active  bool
}

func (aos AoS) ProcessActivePoints() float64 {
    var sum float64
    for i := range aos {
        if aos[i].Active { // Memory jump every iteration
            sum += aos[i].X + aos[i].Y + aos[i].Z
        }
    }
    return sum
}

// 🚀 Cache-friendly: Struct of Arrays
type SoA struct {
    X, Y, Z []float64
    Active  []bool
}

func (soa SoA) ProcessActivePoints() float64 {
    var sum float64
    for i := range soa.Active {
        if soa.Active[i] { // Sequential memory access
            sum += soa.X[i] + soa.Y[i] + soa.Z[i]
        }
    }
    return sum
}
```

## 🔬 Herramientas de Profiling

### 📊 **1. Benchmarking Científico**

```go
package main

import (
    "fmt"
    "testing"
    "time"
)

// 🧪 Ejemplo: Comparando algoritmos de búsqueda
func BenchmarkLinearSearch(b *testing.B) {
    data := generateTestData(10000)
    target := data[5000]
    
    b.ResetTimer() // ¡Crucial! Excluye setup time
    for i := 0; i < b.N; i++ {
        _ = linearSearch(data, target)
    }
}

func BenchmarkBinarySearch(b *testing.B) {
    data := generateSortedTestData(10000)
    target := data[5000]
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = binarySearch(data, target)
    }
}

// 🎯 Benchmark con diferentes tamaños
func BenchmarkSearchAlgorithms(b *testing.B) {
    sizes := []int{100, 1000, 10000, 100000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("Linear-%d", size), func(b *testing.B) {
            data := generateTestData(size)
            target := data[size/2]
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                _ = linearSearch(data, target)
            }
        })
        
        b.Run(fmt.Sprintf("Binary-%d", size), func(b *testing.B) {
            data := generateSortedTestData(size)
            target := data[size/2]
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                _ = binarySearch(data, target)
            }
        })
    }
}
```

### 🔍 **2. CPU Profiling con pprof**

```go
package main

import (
    "os"
    "runtime/pprof"
)

func main() {
    // 🔍 Setup CPU profiling
    cpuFile, err := os.Create("cpu.prof")
    if err != nil {
        panic(err)
    }
    defer cpuFile.Close()
    
    if err := pprof.StartCPUProfile(cpuFile); err != nil {
        panic(err)
    }
    defer pprof.StopCPUProfile()
    
    // 🚀 Tu código a perfilar aquí
    heavyComputation()
}

func heavyComputation() {
    // Simulación de trabajo pesado
    data := make([]int, 1000000)
    for i := range data {
        data[i] = complexCalculation(i)
    }
}

func complexCalculation(n int) int {
    result := 0
    for i := 0; i < n%1000; i++ {
        result += i * i
    }
    return result
}
```

```bash
# 🔬 Analizar el profile
go tool pprof cpu.prof

# 📊 Comandos útiles en pprof:
# top10    - Top 10 funciones por CPU time
# web      - Visualización gráfica (requiere graphviz)
# list functionName - Ver código de función específica
# peek functionName - Ver llamadas desde/hacia función
```

### 🧠 **3. Memory Profiling**

```go
package main

import (
    "os"
    "runtime"
    "runtime/pprof"
)

func main() {
    // 🧠 Memory profiling setup
    defer func() {
        memFile, err := os.Create("mem.prof")
        if err != nil {
            panic(err)
        }
        defer memFile.Close()
        
        runtime.GC() // Forzar garbage collection
        if err := pprof.WriteHeapProfile(memFile); err != nil {
            panic(err)
        }
    }()
    
    // 🚀 Código que usa memoria
    memoryIntensiveOperation()
}

func memoryIntensiveOperation() {
    // Simulación de uso intensivo de memoria
    var data [][]int
    for i := 0; i < 1000; i++ {
        slice := make([]int, 10000)
        for j := range slice {
            slice[j] = j * i
        }
        data = append(data, slice)
    }
    
    // Simular que hacemos algo con los datos
    processData(data)
}

func processData(data [][]int) {
    sum := 0
    for _, slice := range data {
        for _, val := range slice {
            sum += val
        }
    }
}
```

## 🏗️ Técnicas de Optimización Avanzadas

### 1. 🧮 **String Optimization**

```go
package main

import (
    "strings"
    "fmt"
)

// 🐌 Inefficient string concatenation
func ConcatenateNaive(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s // Creates new string each time!
    }
    return result
}

// 🚀 Efficient with strings.Builder
func ConcatenateOptimized(strs []string) string {
    var builder strings.Builder
    
    // Pre-allocate capacity if we know approximate size
    totalLen := 0
    for _, s := range strs {
        totalLen += len(s)
    }
    builder.Grow(totalLen)
    
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}

// 🚀 Even more efficient with strings.Join
func ConcatenateJoin(strs []string) string {
    return strings.Join(strs, "")
}
```

### 2. 🔄 **Slice Optimization**

```go
// 🐌 Inefficient slice growth
func AppendNaive(data []int) []int {
    var result []int
    for _, val := range data {
        result = append(result, val*2) // Potential reallocations
    }
    return result
}

// 🚀 Pre-allocate capacity
func AppendOptimized(data []int) []int {
    result := make([]int, 0, len(data)) // Pre-allocate capacity
    for _, val := range data {
        result = append(result, val*2)
    }
    return result
}

// 🚀 Even better: use indexing instead of append
func TransformInPlace(data []int) []int {
    result := make([]int, len(data))
    for i, val := range data {
        result[i] = val * 2
    }
    return result
}
```

### 3. 🏊 **Object Pooling**

```go
package main

import (
    "sync"
)

// 🚀 Object pool for expensive objects
type ExpensiveObject struct {
    Data []byte
    // Other expensive fields...
}

var expensiveObjectPool = sync.Pool{
    New: func() interface{} {
        return &ExpensiveObject{
            Data: make([]byte, 1024*1024), // 1MB buffer
        }
    },
}

// 🔄 Get object from pool
func GetExpensiveObject() *ExpensiveObject {
    return expensiveObjectPool.Get().(*ExpensiveObject)
}

// 🔄 Return object to pool
func PutExpensiveObject(obj *ExpensiveObject) {
    // Reset object state
    for i := range obj.Data {
        obj.Data[i] = 0
    }
    expensiveObjectPool.Put(obj)
}

// 💡 Usage example
func ProcessData(input []byte) []byte {
    obj := GetExpensiveObject()
    defer PutExpensiveObject(obj)
    
    // Use obj.Data for processing
    copy(obj.Data, input)
    // ... processing logic ...
    
    return obj.Data[:len(input)]
}
```

### 4. ⚡ **Escape Analysis Optimization**

```go
package main

// 🐌 This escapes to heap
func CreateOnHeap() *int {
    x := 42
    return &x // Pointer escapes, x goes to heap
}

// 🚀 This stays on stack
func CreateOnStack() int {
    x := 42
    return x // Value copy, stays on stack
}

// 🔍 Check escape analysis:
// go build -gcflags="-m" main.go
```

## 🧪 Ejemplos Prácticos

### 🎯 **Ejemplo 1: Optimizando JSON Processing**

```go
package main

import (
    "encoding/json"
    "io"
    "sync"
)

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    IsActive bool   `json:"is_active"`
}

// 🐌 Naive JSON processing
func ProcessJSONNaive(data []byte) ([]User, error) {
    var users []User
    err := json.Unmarshal(data, &users)
    return users, err
}

// 🚀 Optimized with decoder and pooling
var decoderPool = sync.Pool{
    New: func() interface{} {
        return json.NewDecoder(nil)
    },
}

func ProcessJSONOptimized(reader io.Reader) ([]User, error) {
    decoder := decoderPool.Get().(*json.Decoder)
    defer decoderPool.Put(decoder)
    
    decoder.Reset(reader)
    
    var users []User
    err := decoder.Decode(&users)
    return users, err
}

// 🚀 Streaming JSON for large datasets
func ProcessJSONStream(reader io.Reader, callback func(User)) error {
    decoder := json.NewDecoder(reader)
    
    // Expect array start
    _, err := decoder.Token()
    if err != nil {
        return err
    }
    
    for decoder.More() {
        var user User
        if err := decoder.Decode(&user); err != nil {
            return err
        }
        callback(user)
    }
    
    return nil
}
```

### 🎯 **Ejemplo 2: Cache-Optimized Matrix Operations**

```go
package main

// 🐌 Cache-unfriendly matrix multiplication
func MultiplyNaive(a, b [][]float64) [][]float64 {
    n := len(a)
    result := make([][]float64, n)
    for i := range result {
        result[i] = make([]float64, n)
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                result[i][j] += a[i][k] * b[k][j] // Poor cache locality
            }
        }
    }
    return result
}

// 🚀 Cache-friendly with loop reordering
func MultiplyCacheFriendly(a, b [][]float64) [][]float64 {
    n := len(a)
    result := make([][]float64, n)
    for i := range result {
        result[i] = make([]float64, n)
    }
    
    for i := 0; i < n; i++ {
        for k := 0; k < n; k++ {
            for j := 0; j < n; j++ {
                result[i][j] += a[i][k] * b[k][j] // Better cache locality
            }
        }
    }
    return result
}

// 🚀 Blocked matrix multiplication for large matrices
func MultiplyBlocked(a, b [][]float64, blockSize int) [][]float64 {
    n := len(a)
    result := make([][]float64, n)
    for i := range result {
        result[i] = make([]float64, n)
    }
    
    for ii := 0; ii < n; ii += blockSize {
        for jj := 0; jj < n; jj += blockSize {
            for kk := 0; kk < n; kk += blockSize {
                // Process block
                for i := ii; i < min(ii+blockSize, n); i++ {
                    for j := jj; j < min(jj+blockSize, n); j++ {
                        for k := kk; k < min(kk+blockSize, n); k++ {
                            result[i][j] += a[i][k] * b[k][j]
                        }
                    }
                }
            }
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 🔥 Técnicas Avanzadas

### 🧬 **1. Assembly Optimization (Cuando es Crítico)**

```go
package main

import "unsafe"

// 🚀 High-performance byte operations using unsafe
func ByteSum(data []byte) int {
    if len(data) == 0 {
        return 0
    }
    
    sum := 0
    i := 0
    
    // Process 8 bytes at a time using unsafe
    for i <= len(data)-8 {
        val := *(*uint64)(unsafe.Pointer(&data[i]))
        sum += int(val&0xFF) + int((val>>8)&0xFF) + int((val>>16)&0xFF) + int((val>>24)&0xFF) +
               int((val>>32)&0xFF) + int((val>>40)&0xFF) + int((val>>48)&0xFF) + int((val>>56)&0xFF)
        i += 8
    }
    
    // Process remaining bytes
    for i < len(data) {
        sum += int(data[i])
        i++
    }
    
    return sum
}

// 🎯 SIMD-style operations for numeric slices
func SumFloat64SIMD(data []float64) float64 {
    if len(data) == 0 {
        return 0
    }
    
    sum := 0.0
    i := 0
    
    // Unroll loop for better performance
    for i <= len(data)-4 {
        sum += data[i] + data[i+1] + data[i+2] + data[i+3]
        i += 4
    }
    
    // Handle remainder
    for i < len(data) {
        sum += data[i]
        i++
    }
    
    return sum
}
```

### 🔄 **2. Lock-Free Programming**

```go
package main

import (
    "sync/atomic"
    "unsafe"
)

// 🚀 Lock-free stack using atomic operations
type LockFreeStack struct {
    head unsafe.Pointer
}

type node struct {
    value interface{}
    next  unsafe.Pointer
}

func (s *LockFreeStack) Push(value interface{}) {
    newNode := &node{value: value}
    for {
        head := atomic.LoadPointer(&s.head)
        newNode.next = head
        if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
            break
        }
        // Retry on contention
    }
}

func (s *LockFreeStack) Pop() interface{} {
    for {
        head := atomic.LoadPointer(&s.head)
        if head == nil {
            return nil
        }
        headNode := (*node)(head)
        next := atomic.LoadPointer(&headNode.next)
        if atomic.CompareAndSwapPointer(&s.head, head, next) {
            return headNode.value
        }
        // Retry on contention
    }
}
```

## 📊 Measuring and Monitoring

### 🎯 **Production Performance Monitoring**

```go
package main

import (
    "context"
    "time"
    "sync/atomic"
)

// 🔍 Performance metrics collector
type PerformanceMetrics struct {
    requestCount     int64
    totalLatency     int64
    errorCount       int64
    maxLatency       int64
}

func (m *PerformanceMetrics) RecordRequest(latency time.Duration, isError bool) {
    atomic.AddInt64(&m.requestCount, 1)
    atomic.AddInt64(&m.totalLatency, int64(latency))
    
    if isError {
        atomic.AddInt64(&m.errorCount, 1)
    }
    
    // Update max latency
    for {
        current := atomic.LoadInt64(&m.maxLatency)
        if int64(latency) <= current {
            break
        }
        if atomic.CompareAndSwapInt64(&m.maxLatency, current, int64(latency)) {
            break
        }
    }
}

func (m *PerformanceMetrics) GetStats() (count, avgLatency, maxLatency, errorRate float64) {
    reqCount := atomic.LoadInt64(&m.requestCount)
    if reqCount == 0 {
        return 0, 0, 0, 0
    }
    
    totalLat := atomic.LoadInt64(&m.totalLatency)
    maxLat := atomic.LoadInt64(&m.maxLatency)
    errCount := atomic.LoadInt64(&m.errorCount)
    
    return float64(reqCount),
           float64(totalLat) / float64(reqCount),
           float64(maxLat),
           float64(errCount) / float64(reqCount)
}

// 🎯 Function timing decorator
func TimeOperation(operation func() error) (time.Duration, error) {
    start := time.Now()
    err := operation()
    return time.Since(start), err
}

// 🚀 Context-aware performance tracking
func TrackPerformance(ctx context.Context, operationName string, fn func() error) error {
    start := time.Now()
    defer func() {
        duration := time.Since(start)
        // Log or send metrics
        metrics.RecordOperation(operationName, duration)
    }()
    
    return fn()
}
```

## 🏆 Proyecto Práctico: High-Performance Web Server

```go
package main

import (
    "bufio"
    "context"
    "fmt"
    "net"
    "net/http"
    "sync"
    "time"
)

// 🚀 High-performance HTTP server with optimizations
type OptimizedServer struct {
    server     *http.Server
    connPool   sync.Pool
    bufferPool sync.Pool
}

func NewOptimizedServer(addr string) *OptimizedServer {
    s := &OptimizedServer{
        server: &http.Server{
            Addr:           addr,
            ReadTimeout:    5 * time.Second,
            WriteTimeout:   10 * time.Second,
            IdleTimeout:    120 * time.Second,
            MaxHeaderBytes: 1 << 16, // 64KB
        },
    }
    
    // Initialize pools
    s.connPool.New = func() interface{} {
        return make([]byte, 32*1024) // 32KB buffer
    }
    
    s.bufferPool.New = func() interface{} {
        return bufio.NewWriterSize(nil, 4096)
    }
    
    return s
}

// 🎯 Optimized handler with pooling
func (s *OptimizedServer) OptimizedHandler(w http.ResponseWriter, r *http.Request) {
    // Get buffer from pool
    writer := s.bufferPool.Get().(*bufio.Writer)
    defer s.bufferPool.Put(writer)
    
    writer.Reset(w)
    defer writer.Flush()
    
    // Fast path for common operations
    switch r.URL.Path {
    case "/health":
        writer.WriteString("OK")
    case "/metrics":
        s.writeMetrics(writer)
    default:
        s.handleRequest(writer, r)
    }
}

func (s *OptimizedServer) writeMetrics(w *bufio.Writer) {
    // Pre-allocated string building for metrics
    w.WriteString(`{"status":"ok","uptime":`)
    w.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
    w.WriteString("}")
}

func (s *OptimizedServer) handleRequest(w *bufio.Writer, r *http.Request) {
    // Handle other requests efficiently
    w.WriteString("Hello, World!")
}
```

## 🎯 Ejercicios Prácticos

### 🧪 **Ejercicio 1: Benchmark Comparison**

Crea benchmarks para comparar diferentes implementaciones de estas funciones:

```go
// TODO: Implementa y benchmarka estas funciones
func ReverseStringMethod1(s string) string {
    // Implementación 1: usando []rune
}

func ReverseStringMethod2(s string) string {
    // Implementación 2: usando strings.Builder
}

func ReverseStringMethod3(s string) string {
    // Implementación 3: usando byte manipulation
}
```

### 🧪 **Ejercicio 2: Memory Pool Implementation**

```go
// TODO: Implementa un memory pool genérico
type MemoryPool[T any] struct {
    // Tu implementación aquí
}

func NewMemoryPool[T any](newFunc func() T) *MemoryPool[T] {
    // Tu implementación aquí
}

func (p *MemoryPool[T]) Get() T {
    // Tu implementación aquí
}

func (p *MemoryPool[T]) Put(item T) {
    // Tu implementación aquí
}
```

### 🧪 **Ejercicio 3: Cache-Friendly Data Structure**

```go
// TODO: Diseña una estructura de datos cache-friendly para almacenar
// millones de puntos 3D y realizar operaciones eficientes sobre ellos

type Point3D struct {
    X, Y, Z float64
}

type Point3DCollection interface {
    Add(p Point3D)
    FindNearby(center Point3D, radius float64) []Point3D
    GetBounds() (min, max Point3D)
}

// Implementa al menos dos versiones y compara su performance
```

## 📊 Métricas de Éxito

Al completar esta lección deberías poder:

- ✅ **Acelerar** código existente en 200-500%
- ✅ **Reducir** uso de memoria en 30-50%
- ✅ **Identificar** cuellos de botella usando profiling
- ✅ **Implementar** object pooling efectivamente
- ✅ **Optimizar** para cache locality
- ✅ **Medir** performance de manera científica

## 🚀 Próximos Pasos

1. **📚 Estudia**: Lee sobre algoritmos cache-oblivious
2. **🔬 Experimenta**: Perfila una aplicación real
3. **🏗️ Construye**: Implementa un servidor HTTP optimizado
4. **📊 Mide**: Compara con benchmarks de la industria

---

**[⬅️ Volver a Avanzado](../README.md) | [🏠 Inicio](../../README.md) | [➡️ Siguiente: Memory Management](../02-memory-management/)**
