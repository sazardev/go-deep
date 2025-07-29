# 🧪 Ejercicios Prácticos: Performance Optimization
## *Pon a prueba tus nuevas habilidades*

### 🎯 Instrucciones Generales

Para cada ejercicio:
1. **🔍 Analiza** el problema y identifica posibles optimizaciones
2. **💻 Implementa** al menos 2 versiones diferentes
3. **📊 Benchmarka** usando `go test -bench=.`
4. **📈 Compara** resultados y explica las diferencias
5. **📝 Documenta** tus hallazgos

---

## 🧪 Ejercicio 1: Optimización de Búsqueda

### 📋 Problema
Tienes una aplicación que necesita buscar elementos en grandes datasets. Implementa y compara diferentes algoritmos de búsqueda.

### 🎯 Tareas
```go
// TODO: Implementa estas funciones y crea benchmarks

// 🐌 Búsqueda lineal básica
func LinearSearch(data []int, target int) int {
    // Tu implementación aquí
}

// 🚀 Búsqueda binaria
func BinarySearch(data []int, target int) int {
    // Tu implementación aquí (asume que data está ordenado)
}

// 🧠 Búsqueda con hash map
func HashSearch(data []int, target int) int {
    // Tu implementación aquí (crea map una vez, búsqueda O(1))
}

// 📊 Benchmark function
func BenchmarkSearchMethods(b *testing.B) {
    // Crea benchmarks para diferentes tamaños: 1K, 10K, 100K, 1M
}
```

### 🏆 Objetivos de Performance
- **Linear Search**: Baseline
- **Binary Search**: 10-100x más rápido que linear
- **Hash Search**: Búsqueda constante O(1)

---

## 🧪 Ejercicio 2: String Processing Optimization

### 📋 Problema
Tu aplicación procesa logs de texto y necesita operaciones de string extremadamente rápidas.

### 🎯 Tareas
```go
// TODO: Implementa y optimiza estas funciones

// 🐌 Concatenación naive
func ConcatenateNaive(strings []string, separator string) string {
    // Tu implementación aquí
}

// 🚀 Concatenación con strings.Builder
func ConcatenateBuilder(strings []string, separator string) string {
    // Tu implementación aquí
}

// 🚀 Concatenación con pre-allocation
func ConcatenatePrealloc(strings []string, separator string) string {
    // Tu implementación aquí
}

// 🔄 Función de reversa de string optimizada
func ReverseStringOptimized(s string) string {
    // Implementa la versión más rápida posible
}

// 📊 Word count optimizado
func WordCountOptimized(text string) map[string]int {
    // Cuenta palabras de manera eficiente
}
```

### 🏆 Objetivos de Performance
- **Builder vs Naive**: 5-10x mejora en concatenación
- **Reverse**: Sub-microsegundo para strings de 1KB
- **Word Count**: Procesar 1MB de texto en <100ms

---

## 🧪 Ejercicio 3: Memory Pool Implementation

### 📋 Problema
Tu aplicación crea/destruye muchos objetos del mismo tipo, causando presión en el garbage collector.

### 🎯 Tareas
```go
// TODO: Implementa un sistema de object pooling completo

// 🏗️ Objeto que es caro de crear
type ExpensiveResource struct {
    Buffer []byte
    Connections map[string]interface{}
    // Agrega más campos según necesites
}

// 🔄 Pool interface
type ResourcePool interface {
    Get() *ExpensiveResource
    Put(*ExpensiveResource)
    Stats() (total, inUse, available int)
}

// 🚀 Implementación del pool
type OptimizedPool struct {
    // Tu implementación aquí
}

// 📊 Función para testear con y sin pool
func ProcessWorkloads(usePool bool, workloadSize int) {
    // Simula carga de trabajo pesada
}
```

### 🏆 Objetivos de Performance
- **Reducción de GC**: 50-80% menos allocations
- **Latencia**: 30-50% mejora en p99
- **Throughput**: 20-40% más operaciones/segundo

---

## 🧪 Ejercicio 4: Cache-Friendly Data Structures

### 📋 Problema
Necesitas procesar millones de estructuras de datos de manera eficiente, maximizando cache locality.

### 🎯 Tareas
```go
// TODO: Diseña estructuras de datos cache-friendly

// 🐌 Estructura cache-unfriendly
type PersonAoS struct {
    ID       int64
    Name     string
    Email    string
    Age      int32
    Salary   float64
    Active   bool
}

// 🚀 Estructura cache-friendly
type PersonSoA struct {
    // Tu diseño aquí
}

// 📊 Operaciones para comparar
func CalculateAverageSalaryAoS(people []PersonAoS) float64 {
    // Solo para personas activas
}

func CalculateAverageSalarySoA(people PersonSoA) float64 {
    // Solo para personas activas
}

// 🎯 Filtrado optimizado
func FilterActivePeopleOptimized(people PersonSoA) PersonSoA {
    // Retorna solo personas activas
}
```

### 🏆 Objetivos de Performance
- **SoA vs AoS**: 2-5x mejora en operaciones selectivas
- **Memory Usage**: 10-30% menos uso de memoria
- **Cache Misses**: 50-70% reducción en cache misses

---

## 🧪 Ejercicio 5: High-Performance JSON Processing

### 📋 Problema
Tu API procesa JSON de gran tamaño y necesita máximo throughput.

### 🎯 Tareas
```go
// TODO: Optimiza el procesamiento de JSON

// 🐌 Procesamiento básico
func ProcessJSONBasic(data []byte) ([]User, error) {
    // Implementación estándar
}

// 🚀 Procesamiento con streaming
func ProcessJSONStream(reader io.Reader) ([]User, error) {
    // Procesamiento streaming
}

// 🚀 Procesamiento con pooling
func ProcessJSONPooled(data []byte) ([]User, error) {
    // Usa pools para decoders y buffers
}

// 🎯 Serialización optimizada
func SerializeUsersOptimized(users []User) ([]byte, error) {
    // Serialización de alta performance
}
```

### 🏆 Objetivos de Performance
- **Streaming vs Basic**: 3-5x mejora en memoria
- **Pooled vs Basic**: 30-50% mejora en throughput
- **Zero-Copy**: Minimizar allocations

---

## 🧪 Ejercicio 6: Concurrent Processing Optimization

### 📋 Problema
Procesas archivos grandes usando concurrencia, pero necesitas optimizar el balance entre paralelismo y overhead.

### 🎯 Tareas
```go
// TODO: Optimiza el procesamiento concurrente

// 🐌 Procesamiento secuencial
func ProcessFilesSequential(files []string) []Result {
    // Baseline
}

// 🚀 Procesamiento concurrente básico
func ProcessFilesConcurrent(files []string, workers int) []Result {
    // Worker pool simple
}

// 🚀 Procesamiento con work stealing
func ProcessFilesWorkStealing(files []string) []Result {
    // Algoritmo avanzado de distribución de trabajo
}

// 📊 Función para encontrar número óptimo de workers
func FindOptimalWorkerCount(files []string) int {
    // Experimenta y encuentra el sweet spot
}
```

### 🏆 Objetivos de Performance
- **Concurrent vs Sequential**: 2-8x mejora (depende de CPU cores)
- **Work Stealing**: 20-40% mejor que worker pool básico
- **CPU Utilization**: >80% en todos los cores

---

## 📊 Criterios de Evaluación

### 🎯 **Performance Metrics**
```go
type BenchmarkResult struct {
    Algorithm    string
    InputSize    int
    NsPerOp      int64
    AllocsPerOp  int64
    BytesPerOp   int64
    MBPerSec     float64
}
```

### 🏆 **Scoring**
- **🥇 Gold (90-100%)**: Supera todos los objetivos + innovación
- **🥈 Silver (80-89%)**: Cumple todos los objetivos
- **🥉 Bronze (70-79%)**: Cumple la mayoría de objetivos
- **📈 Improvement**: Cualquier mejora significativa es valiosa

### 📈 **Bonus Points**
- **🧠 Memory Efficiency**: Reduce allocations significativamente
- **⚡ CPU Optimization**: Usa todas las características del CPU
- **🔍 Profiling**: Incluye análisis con pprof
- **📊 Visualization**: Gráficas de performance
- **📝 Documentation**: Explica las optimizaciones detalladamente

---

## 🛠️ Herramientas de Testing

### 📊 **Comandos de Benchmark**
```bash
# 🚀 Benchmark básico
go test -bench=.

# 📊 Con información de memoria
go test -bench=. -benchmem

# 🔄 Múltiples runs para estabilidad
go test -bench=. -count=5

# 🎯 Benchmarks específicos
go test -bench=BenchmarkMyFunction

# 📈 Output a archivo para análisis
go test -bench=. -benchmem > results.txt

# 🔍 Profiling durante benchmark
go test -bench=BenchmarkMyFunction -cpuprofile=cpu.prof
go test -bench=BenchmarkMyFunction -memprofile=mem.prof
```

### 🔍 **Análisis con pprof**
```bash
# 📊 CPU profiling
go tool pprof cpu.prof

# 🧠 Memory profiling  
go tool pprof mem.prof

# 🌐 Web interface
go tool pprof -http=:8080 cpu.prof
```

---

## 📚 Recursos Adicionales

### 📖 **Lecturas Recomendadas**
- [Go Performance Tips](https://github.com/dgryski/go-perfbook)
- [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
- [Profiling Go Programs](https://blog.golang.org/pprof)

### 🎯 **Benchmarks de Referencia**
- **String Operations**: 1M ops/sec como mínimo
- **Slice Operations**: <1ns per element
- **Map Operations**: <100ns per lookup
- **JSON Processing**: >100MB/sec throughput

### 🏆 **Casos de Estudio**
- **Netflix**: Optimizaciones de microservicios
- **Uber**: High-throughput systems
- **Google**: Compiler optimizations
- **Docker**: Container performance

---

## 🎯 Entregables

Para cada ejercicio, entrega:

1. **💻 Código**: Implementaciones completas con comentarios
2. **📊 Benchmarks**: Resultados detallados de performance
3. **📈 Análisis**: Comparación de métodos y explicación
4. **🔍 Profiling**: Resultados de pprof cuando sea relevante
5. **📝 Reporte**: Documento con hallazgos y recomendaciones

### 📋 **Template de Reporte**
```markdown
# Performance Optimization Report

## Ejercicio: [Nombre]

### Implementaciones
- Método 1: [Descripción]
- Método 2: [Descripción]
- Método 3: [Descripción]

### Benchmark Results
| Método  | Input Size | ns/op | allocs/op | bytes/op | Speedup |
| ------- | ---------- | ----- | --------- | -------- | ------- |
| Método1 | 1K         | ...   | ...       | ...      | 1.0x    |
| Método2 | 1K         | ...   | ...       | ...      | 2.3x    |

### Análisis
[Explicación detallada de resultados]

### Conclusiones
[Recomendaciones y aprendizajes]
```

---

## 🚀 ¡A Optimizar!

```bash
# 🎯 Setup para empezar
mkdir performance-exercises
cd performance-exercises
go mod init performance-exercises

# 🧪 Crear primer ejercicio
touch exercise1_search.go
touch exercise1_test.go

# 🔥 ¡Que comience la optimización!
echo "Ready to make Go fly! 🚀"
```

**Recuerda**: *"La optimización es un arte que se perfecciona con la práctica. Cada microsegundo cuenta!"* 💪
