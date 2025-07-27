# 📝 Resumen: Lección 18 - Patrones Avanzados de Concurrencia

## 🎯 Patrones Implementados

### 1. 🏭 **Worker Pool Pattern**
- **Propósito**: Distribuir trabajo entre un número fijo de workers
- **Ventajas**: Control de recursos, evita sobrecarga del sistema
- **Casos de uso**: Procesamiento de archivos, requests HTTP, tareas CPU-intensivas

```go
// Ejemplo de uso
jobs := []int{1, 2, 3, 4, 5}
results := WorkerPool(jobs, 3) // 3 workers
// Output: [2, 4, 6, 8, 10] (duplicados)
```

### 2. 🔄 **Pipeline Pattern**
- **Propósito**: Procesar datos en etapas secuenciales conectadas
- **Ventajas**: Separación de responsabilidades, composición modular
- **Casos de uso**: Transformación de datos, filtrado, validación

```go
// Pipeline: input -> multiplicar por 2 -> sumar 1
input := []int{1, 2, 3}
results := Pipeline(input)
// Output: [3, 5, 7] ((1*2)+1, (2*2)+1, (3*2)+1)
```

### 3. 📊 **Fan-Out / Fan-In Pattern**
- **Propósito**: Distribuir trabajo (fan-out) y recolectar resultados (fan-in)
- **Ventajas**: Máximo paralelismo, escalabilidad horizontal
- **Casos de uso**: Búsquedas paralelas, agregación de datos

```go
jobs := []int{1, 2, 3, 4}
results := FanOutFanIn(jobs, 2) // 2 workers
// Output: [2, 4, 6, 8] (orden puede variar)
```

### 4. ⏳ **Throttling / Rate Limiting**
- **Propósito**: Limitar el número de operaciones concurrentes
- **Ventajas**: Protege recursos externos, evita sobrecarga
- **Casos de uso**: APIs con límites, bases de datos, servicios externos

```go
jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
results := ThrottledProcessing(jobs, 2) // Máximo 2 concurrentes
// Se ejecuta más lento pero controlado
```

### 5. 🛑 **Graceful Shutdown**
- **Propósito**: Parar sistemas concurrentes de forma elegante
- **Ventajas**: Evita pérdida de datos, limpieza adecuada
- **Casos de uso**: Servidores web, workers de fondo, servicios

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
results := GracefulShutdown(ctx, jobs)
// Se detiene cuando se cancela el contexto
```

---

## 🧪 Tests y Benchmarks

### Tests Implementados
- ✅ **Unit tests**: Verifican funcionalidad básica
- ✅ **Edge cases**: Casos límite (slices vacíos, un worker)
- ✅ **Concurrency tests**: Verifican paralelismo real
- ✅ **Timeout tests**: Graceful shutdown con diferentes timeouts

### Benchmarks
```
BenchmarkWorkerPool-2      10807    107319 ns/op
BenchmarkPipeline-2         1983    534376 ns/op
BenchmarkFanOutFanIn-2      2695    418764 ns/op
```

---

## 🎯 Conceptos Clave Aprendidos

### 1. **Channel Patterns**
- Canales buffered vs unbuffered
- Cerrar canales correctamente
- Select con timeout

### 2. **Synchronization**
- sync.WaitGroup para coordinación
- Semáforos con canales buffered
- Context para cancelación

### 3. **Error Handling**
- Propagación de errores en concurrencia
- Timeout y cancelación
- Cleanup de recursos

### 4. **Performance**
- Overhead de goroutines
- Balanceo de carga
- Throttling para proteger recursos

---

## 🛠️ Herramientas Utilizadas

- **Goroutines**: Unidades de concurrencia ligeras
- **Channels**: Comunicación entre goroutines
- **sync.WaitGroup**: Esperar que múltiples goroutines terminen
- **context.Context**: Cancelación y timeouts
- **time.After**: Timeouts con select
- **Buffered channels**: Control de flujo y semáforos

---

## 📊 Casos de Uso Reales

### Worker Pool
- **Servidores web**: Pool de workers para requests
- **Procesamiento de archivos**: Leer/procesar múltiples archivos
- **ETL**: Extract, Transform, Load de datos

### Pipeline
- **Stream processing**: Kafka, Apache Storm
- **Image processing**: Filtros secuenciales
- **Compilation**: Lexer -> Parser -> Codegen

### Fan-Out/Fan-In
- **Microservicios**: Consultas paralelas a múltiples servicios
- **Map-Reduce**: Distribución y agregación
- **Search engines**: Búsqueda en múltiples índices

### Throttling
- **API clients**: Respetar rate limits
- **Database connections**: Pool de conexiones
- **Resource management**: CPU, memoria, I/O

### Graceful Shutdown
- **HTTP servers**: Terminar requests activos
- **Background workers**: Completar tareas en curso
- **Microservices**: Shutdown ordenado

---

## 🚀 Próximos Pasos

1. **Practica con proyectos reales**
2. **Aprende sobre context.Context avanzado**
3. **Estudia sync.Pool para reutilización de objetos**
4. **Explora errgroup para manejo de errores**
5. **Investiga monitoring y observabilidad**

---

## 💡 Best Practices

- 🎯 **Usa el patrón correcto** para el problema específico
- 🛡️ **Evita goroutine leaks** cerrando canales y usando context
- ⚡ **Mide el rendimiento** con benchmarks antes de optimizar
- 🧪 **Escribe tests** que verifiquen concurrencia real
- 📊 **Monitorea** el número de goroutines en producción

¡Has completado exitosamente los patrones avanzados de concurrencia en Go! 🎉
