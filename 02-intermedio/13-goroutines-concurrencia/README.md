# 🚀 Lección 13: Goroutines y Concurrencia

## 🎯 Objetivos de la Lección

Al finalizar esta lección, serás capaz de:
- Entender qué son las goroutines y por qué son revolucionarias
- Implementar concurrencia eficiente con goroutines
- Gestionar el ciclo de vida de goroutines
- Evitar race conditions y memory leaks
- Diseñar patrones de concurrencia escalables
- Optimizar performance con paralelización
- Debuggear problemas de concurrencia

---

## 🧠 Analogía: La Fábrica de Pizzas Concurrente

Imagina una pizzería tradicional vs una pizzería moderna:

### 🍕 Pizzería Tradicional (Secuencial)
```
Chef -> Masa -> Salsa -> Queso -> Hornear -> Entregar
 |       |       |       |        |         |
 5min   3min    2min    1min     15min     2min
                Total: 28 minutos por pizza
```

### 🏭 Pizzería Moderna (Concurrente)
```
Chef A: Masa -----> Masa -----> Masa
Chef B:     Salsa -----> Salsa -----> Salsa  
Chef C:         Queso -----> Queso -----> Queso
Horno:              Pizza1 -> Pizza2 -> Pizza3
Delivery:               Entrega1 -> Entrega2
                Resultado: 3 pizzas en 20 minutos
```

**Las goroutines son como tener múltiples chefs especializados trabajando en paralelo, coordinándose perfectamente sin pisarse los pies.**

---

## 📚 Fundamentos de Goroutines

### 🔧 ¿Qué es una Goroutine?

Una **goroutine** es una función que puede ejecutarse concurrentemente con otras funciones. Son:

- **Ligeras**: Ocupan solo ~2KB de memoria inicial
- **Eficientes**: Manejadas por el runtime de Go
- **Escalables**: Puedes tener millones sin problemas
- **Coordinadas**: Comunicación via channels

### 🎭 Diferencias Clave

| Concepto | Threads OS | Goroutines |
|----------|------------|------------|
| **Memoria** | ~2MB | ~2KB |
| **Creación** | Costosa | Muy barata |
| **Switching** | Kernel space | User space |
| **Scheduling** | OS Scheduler | Go Scheduler |
| **Máximo** | Miles | Millones |

---

## 🚀 Creando y Usando Goroutines

### 1. 🌟 Goroutine Básica

```go
package main

import (
    "fmt"
    "time"
)

// Función normal
func saludar(nombre string) {
    for i := 0; i < 5; i++ {
        fmt.Printf("👋 Hola %s #%d\n", nombre, i+1)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    fmt.Println("🎬 Inicio del programa")
    
    // Ejecución secuencial
    fmt.Println("\n📝 Ejecución Secuencial:")
    saludar("Ana")
    saludar("Bob")
    
    fmt.Println("\n🚀 Ejecución Concurrente:")
    // Lanzar goroutines
    go saludar("Carlos")  // ¡Magia! Solo agregar 'go'
    go saludar("Diana")
    
    // Esperar a que terminen (método básico)
    time.Sleep(600 * time.Millisecond)
    
    fmt.Println("🏁 Fin del programa")
}
```

### 2. 🔄 Goroutines con WaitGroup

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func trabajador(id int, wg *sync.WaitGroup) {
    defer wg.Done() // ¡Crucial! Avisar que terminó
    
    fmt.Printf("👷 Trabajador %d iniciando\n", id)
    
    // Simular trabajo
    tiempo := time.Duration(id*100) * time.Millisecond
    time.Sleep(tiempo)
    
    fmt.Printf("✅ Trabajador %d terminó\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    fmt.Println("🏗️ Iniciando trabajo en paralelo")
    
    // Lanzar 5 trabajadores
    for i := 1; i <= 5; i++ {
        wg.Add(1) // Incrementar contador
        go trabajador(i, &wg)
    }
    
    fmt.Println("⏳ Esperando que terminen todos...")
    wg.Wait() // Bloquear hasta que todos terminen
    
    fmt.Println("🎉 ¡Todos los trabajadores terminaron!")
}
```

### 3. 🎯 Goroutines con Valores de Retorno

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Resultado struct {
    ID    int
    Valor int
    Error error
}

func calcularNumeroAleatorio(id int, resultados chan<- Resultado, wg *sync.WaitGroup) {
    defer wg.Done()
    
    // Simular procesamiento
    time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
    
    // Generar resultado
    valor := rand.Intn(100)
    
    resultado := Resultado{
        ID:    id,
        Valor: valor,
        Error: nil,
    }
    
    resultados <- resultado
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    var wg sync.WaitGroup
    resultados := make(chan Resultado, 10) // Buffer para resultados
    
    // Lanzar goroutines
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go calcularNumeroAleatorio(i, resultados, &wg)
    }
    
    // Cerrar canal cuando terminen todas
    go func() {
        wg.Wait()
        close(resultados)
    }()
    
    // Recoger resultados
    fmt.Println("📊 Resultados:")
    for resultado := range resultados {
        fmt.Printf("  ID: %d, Valor: %d\n", resultado.ID, resultado.Valor)
    }
}
```

---

## 🔒 Sincronización y Race Conditions

### ⚠️ Problema: Race Condition

```go
package main

import (
    "fmt"
    "sync"
)

// ❌ CÓDIGO PELIGROSO - Race Condition
var contador int

func incrementar(wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        contador++ // ¡PELIGRO! Race condition
    }
}

func ejemploRaceCondition() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go incrementar(&wg)
    }
    
    wg.Wait()
    fmt.Printf("❌ Contador final (con race condition): %d\n", contador)
    fmt.Println("   Esperábamos: 10000")
}
```

### ✅ Solución 1: Mutex

```go
package main

import (
    "fmt"
    "sync"
)

var (
    contadorSeguro int
    mutex         sync.Mutex
)

func incrementarSeguro(wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        mutex.Lock()   // 🔒 Bloquear acceso
        contadorSeguro++
        mutex.Unlock() // 🔓 Liberar acceso
    }
}

func ejemploMutex() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go incrementarSeguro(&wg)
    }
    
    wg.Wait()
    fmt.Printf("✅ Contador final (con mutex): %d\n", contadorSeguro)
}
```

### ✅ Solución 2: Atomic Operations

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

var contadorAtomico int64

func incrementarAtomico(wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000; i++ {
        atomic.AddInt64(&contadorAtomico, 1) // ⚡ Operación atómica
    }
}

func ejemploAtomic() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go incrementarAtomico(&wg)
    }
    
    wg.Wait()
    
    valor := atomic.LoadInt64(&contadorAtomico)
    fmt.Printf("✅ Contador final (atomic): %d\n", valor)
}
```

---

## 🏃‍♂️ Patrones de Concurrencia

### 1. 🏗️ Worker Pool Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Trabajo struct {
    ID   int
    Data string
}

type Resultado struct {
    Trabajo   Trabajo
    Resultado string
    Error     error
}

func worker(id int, trabajos <-chan Trabajo, resultados chan<- Resultado, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for trabajo := range trabajos {
        fmt.Printf("👷 Worker %d procesando trabajo %d\n", id, trabajo.ID)
        
        // Simular procesamiento
        time.Sleep(100 * time.Millisecond)
        
        resultado := Resultado{
            Trabajo:   trabajo,
            Resultado: fmt.Sprintf("Procesado por worker %d", id),
            Error:     nil,
        }
        
        resultados <- resultado
    }
    
    fmt.Printf("👷 Worker %d terminando\n", id)
}

func ejemploWorkerPool() {
    const numWorkers = 3
    const numTrabajos = 10
    
    trabajos := make(chan Trabajo, numTrabajos)
    resultados := make(chan Resultado, numTrabajos)
    
    var wg sync.WaitGroup
    
    // Crear workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, trabajos, resultados, &wg)
    }
    
    // Enviar trabajos
    for i := 1; i <= numTrabajos; i++ {
        trabajo := Trabajo{
            ID:   i,
            Data: fmt.Sprintf("Datos del trabajo %d", i),
        }
        trabajos <- trabajo
    }
    close(trabajos) // ¡Importante! Cerrar canal
    
    // Cerrar resultados cuando todos terminen
    go func() {
        wg.Wait()
        close(resultados)
    }()
    
    // Recoger resultados
    fmt.Println("\n📊 Resultados:")
    for resultado := range resultados {
        fmt.Printf("  Trabajo %d: %s\n", 
            resultado.Trabajo.ID, resultado.Resultado)
    }
}
```

### 2. 🔄 Pipeline Pattern

```go
package main

import (
    "fmt"
    "sync"
)

func generarNumeros(nums chan<- int) {
    defer close(nums)
    
    for i := 1; i <= 10; i++ {
        nums <- i
    }
}

func elevarCuadrado(nums <-chan int, cuadrados chan<- int) {
    defer close(cuadrados)
    
    for num := range nums {
        cuadrados <- num * num
    }
}

func filtrarPares(cuadrados <-chan int, pares chan<- int) {
    defer close(pares)
    
    for cuadrado := range cuadrados {
        if cuadrado%2 == 0 {
            pares <- cuadrado
        }
    }
}

func sumarTodos(pares <-chan int, resultado chan<- int) {
    defer close(resultado)
    
    suma := 0
    for par := range pares {
        suma += par
    }
    
    resultado <- suma
}

func ejemploPipeline() {
    // Crear canales
    nums := make(chan int)
    cuadrados := make(chan int)
    pares := make(chan int)
    resultado := make(chan int)
    
    // Pipeline: generar -> elevar -> filtrar -> sumar
    go generarNumeros(nums)
    go elevarCuadrado(nums, cuadrados)
    go filtrarPares(cuadrados, pares)
    go sumarTodos(pares, resultado)
    
    // Obtener resultado final
    suma := <-resultado
    fmt.Printf("🧮 Suma de cuadrados pares: %d\n", suma)
}
```

### 3. 📊 Fan-Out/Fan-In Pattern

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func generarTrabajos(trabajos chan<- int) {
    defer close(trabajos)
    
    for i := 1; i <= 20; i++ {
        trabajos <- i
    }
}

func procesarTrabajos(id int, trabajos <-chan int, resultados chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for trabajo := range trabajos {
        // Simular tiempo de procesamiento variable
        tiempo := time.Duration(rand.Intn(100)) * time.Millisecond
        time.Sleep(tiempo)
        
        resultado := fmt.Sprintf("Worker %d procesó trabajo %d", id, trabajo)
        resultados <- resultado
    }
}

func ejemploFanOutFanIn() {
    rand.Seed(time.Now().UnixNano())
    
    trabajos := make(chan int)
    resultados := make(chan string, 20)
    
    var wg sync.WaitGroup
    
    // Fan-Out: Distribuir trabajo a múltiples workers
    numWorkers := 5
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go procesarTrabajos(i, trabajos, resultados, &wg)
    }
    
    // Generar trabajos
    go generarTrabajos(trabajos)
    
    // Fan-In: Recoger todos los resultados
    go func() {
        wg.Wait()
        close(resultados)
    }()
    
    fmt.Println("📋 Resultados del procesamiento:")
    for resultado := range resultados {
        fmt.Printf("  %s\n", resultado)
    }
}
```

---

## ⚡ Optimización de Performance

### 1. 🎯 Control de Goroutines

```go
package main

import (
    "context"
    "fmt"
    "runtime"
    "sync"
    "time"
)

func workerConContext(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("👷 Worker %d recibió señal de cancelación\n", id)
            return
        default:
            // Simular trabajo
            fmt.Printf("👷 Worker %d trabajando...\n", id)
            time.Sleep(200 * time.Millisecond)
        }
    }
}

func ejemploControlGoroutines() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    var wg sync.WaitGroup
    
    // Lanzar workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go workerConContext(ctx, i, &wg)
    }
    
    // Esperar a que terminen
    wg.Wait()
    fmt.Println("✅ Todos los workers terminaron")
}
```

### 2. 📊 Monitoreo de Goroutines

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func monitorearGoroutines() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    
    for i := 0; i < 10; i++ {
        <-ticker.C
        fmt.Printf("📊 Goroutines activas: %d\n", runtime.NumGoroutine())
    }
}

func crearGoroutines(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    var localWg sync.WaitGroup
    
    for i := 0; i < n; i++ {
        localWg.Add(1)
        go func(id int) {
            defer localWg.Done()
            time.Sleep(200 * time.Millisecond)
        }(i)
    }
    
    localWg.Wait()
}

func ejemploMonitoreo() {
    fmt.Println("📊 Monitoreando goroutines...")
    
    go monitorearGoroutines()
    
    var wg sync.WaitGroup
    
    // Crear oleadas de goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go crearGoroutines(5, &wg)
        time.Sleep(50 * time.Millisecond)
    }
    
    wg.Wait()
    time.Sleep(500 * time.Millisecond) // Tiempo para el monitor
}
```

---

## ⚠️ Errores Comunes y Cómo Evitarlos

### 1. 🚫 Goroutine Leak

```go
// ❌ MAL: Goroutine que nunca termina
func goroutineLeakMalo() {
    ch := make(chan int)
    
    go func() {
        // Esta goroutine quedará bloqueada para siempre
        valor := <-ch
        fmt.Println(valor)
    }()
    
    // Función termina, pero la goroutine sigue viva
    // ¡MEMORY LEAK!
}

// ✅ BIEN: Goroutine con timeout
func goroutineLeakBueno() {
    ch := make(chan int)
    
    go func() {
        select {
        case valor := <-ch:
            fmt.Println(valor)
        case <-time.After(1 * time.Second):
            fmt.Println("Timeout: goroutine terminando")
            return
        }
    }()
    
    // La goroutine se auto-destruirá después de 1 segundo
}
```

### 2. 🔄 Compartir Memoria por Comunicación

```go
// ❌ MAL: Compartir estado con mutex
type ContadorMutex struct {
    valor int
    mu    sync.Mutex
}

func (c *ContadorMutex) Incrementar() {
    c.mu.Lock()
    c.valor++
    c.mu.Unlock()
}

// ✅ BIEN: Comunicar estado con channels
type ContadorChannel struct {
    ch chan int
}

func NewContadorChannel() *ContadorChannel {
    c := &ContadorChannel{
        ch: make(chan int),
    }
    
    go func() {
        valor := 0
        for incremento := range c.ch {
            valor += incremento
            fmt.Printf("Contador: %d\n", valor)
        }
    }()
    
    return c
}

func (c *ContadorChannel) Incrementar() {
    c.ch <- 1
}
```

---

## 💡 Tips de Experto

### 🎯 Cuándo Usar Goroutines

```go
// ✅ BUENO: I/O operations
func procesarArchivos(archivos []string) {
    var wg sync.WaitGroup
    
    for _, archivo := range archivos {
        wg.Add(1)
        go func(nombre string) {
            defer wg.Done()
            // Leer archivo (I/O bound)
            procesarArchivo(nombre)
        }(archivo)
    }
    
    wg.Wait()
}

// ✅ BUENO: CPU intensive con múltiples cores
func calcularParalelo(datos []int) []int {
    numCPU := runtime.NumCPU()
    resultados := make([]int, len(datos))
    
    chunkSize := len(datos) / numCPU
    var wg sync.WaitGroup
    
    for i := 0; i < numCPU; i++ {
        start := i * chunkSize
        end := start + chunkSize
        if i == numCPU-1 {
            end = len(datos) // Último chunk toma el resto
        }
        
        wg.Add(1)
        go func(inicio, fin int) {
            defer wg.Done()
            for j := inicio; j < fin; j++ {
                resultados[j] = calcularComplejo(datos[j])
            }
        }(start, end)
    }
    
    wg.Wait()
    return resultados
}
```

### 🔧 Configuración Óptima

```go
func configurarOptima() {
    // Configurar número de threads del OS
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    // Para debugging: habilitar race detector
    // go run -race main.go
    
    // Configurar garbage collector para baja latencia
    // GOGC=100 go run main.go
}
```

---

## 🧪 Testing de Código Concurrente

```go
package main

import (
    "sync"
    "testing"
    "time"
)

func TestConcurrencia(t *testing.T) {
    const numGoroutines = 100
    const incrementosPorGoroutine = 1000
    
    var contador int64
    var wg sync.WaitGroup
    
    inicio := time.Now()
    
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < incrementosPorGoroutine; j++ {
                atomic.AddInt64(&contador, 1)
            }
        }()
    }
    
    wg.Wait()
    duracion := time.Since(inicio)
    
    esperado := int64(numGoroutines * incrementosPorGoroutine)
    if contador != esperado {
        t.Errorf("Esperado %d, obtenido %d", esperado, contador)
    }
    
    t.Logf("Operaciones: %d en %v", contador, duracion)
}
```

---

## 🎯 Resumen de la Lección

### ✅ Conceptos Clave Aprendidos

1. **🚀 Goroutines**: Funciones concurrentes ligeras
2. **🔄 WaitGroup**: Sincronización de múltiples goroutines  
3. **🔒 Mutex**: Protección de recursos compartidos
4. **⚡ Atomic**: Operaciones thread-safe rápidas
5. **🏗️ Worker Pool**: Patrón para procesar trabajos
6. **🔄 Pipeline**: Cadena de procesamiento
7. **📊 Fan-Out/Fan-In**: Distribución y agregación
8. **🎯 Context**: Control y cancelación
9. **⚠️ Race Conditions**: Problemas y soluciones
10. **💡 Best Practices**: Cuándo y cómo usar

### 🏆 Logros Desbloqueados

- [ ] 🥇 **Concurrent Novice**: Primera goroutine exitosa
- [ ] 🥈 **Race Detector**: Identificar y corregir race condition
- [ ] 🥉 **Worker Master**: Implementar worker pool
- [ ] 🏅 **Pipeline Engineer**: Crear pipeline de procesamiento
- [ ] 🎖️ **Performance Optimizer**: Optimizar código concurrente

### 📚 Próximos Pasos

En la **Lección 14: Channels**, aprenderemos:
- Comunicación entre goroutines
- Tipos de channels
- Patrones avanzados con channels
- Select statement
- Channel buffering

---

**🎉 ¡Felicitaciones! Has dominado los fundamentos de la concurrencia en Go. Las goroutines son una de las características más poderosas del lenguaje y ahora tienes las herramientas para crear aplicaciones escalables y eficientes.**

*Recuerda: "No comuniques compartiendo memoria; comparte memoria comunicando" - Rob Pike*
