# 📡 Lección 14: Channels - Comunicación entre Goroutines

## 🎯 Objetivos de la Lección

Al finalizar esta lección, serás capaz de:
- Entender qué son los channels y por qué son fundamentales
- Crear y usar diferentes tipos de channels
- Implementar comunicación segura entre goroutines
- Dominar el statement `select` para multiplexación
- Diseñar patrones avanzados con channels
- Evitar deadlocks y channel leaks
- Optimizar performance con buffering

---

## 🧠 Analogía: El Sistema de Tuberías de la Ciudad

Imagina las **goroutines** como **casas** en una ciudad, y los **channels** como el **sistema de tuberías** que conecta estas casas:

```
🏠 Casa A ═══════════════╗
                         ║  📡 Channel (tubería)
🏠 Casa B ═══════════════╬═══════════════ 🏠 Casa C
                         ║
🏠 Casa D ═══════════════╝
```

- **🏠 Casas (Goroutines)**: Realizan trabajos independientes
- **📡 Tuberías (Channels)**: Transportan información entre casas
- **🚰 Agua (Datos)**: Fluye de una casa a otra a través de las tuberías
- **🔧 Válvulas (Select)**: Controlan qué tubería usar en cada momento

---

## 📚 Fundamentos de Channels

### 🔧 ¿Qué es un Channel?

Un **channel** es un conducto de comunicación que permite a las goroutines intercambiar datos de forma segura. Son:

- **Type-safe**: Solo transportan un tipo específico de datos
- **Thread-safe**: Sincronización automática entre goroutines
- **First-class**: Pueden pasarse como parámetros y valores de retorno
- **Blocking**: Operaciones síncronas por defecto

### 🎭 Filosofía de Go

> **"No comuniques compartiendo memoria; comparte memoria comunicando"**
> 
> *- Rob Pike, co-creador de Go*

```go
// ❌ MAL: Compartir memoria
var shared int
var mutex sync.Mutex

func increment() {
    mutex.Lock()
    shared++
    mutex.Unlock()
}

// ✅ BIEN: Compartir comunicando
ch := make(chan int)

go func() {
    ch <- 42  // Enviar
}()

value := <-ch  // Recibir
```

---

## 📡 Creando y Usando Channels

### 1. 🌟 Channel Básico

```go
package main

import (
    "fmt"
    "time"
)

func productor(ch chan<- string) {  // Send-only channel
    mensajes := []string{"Hola", "Mundo", "desde", "Go!"}
    
    for _, mensaje := range mensajes {
        fmt.Printf("📤 Enviando: %s\n", mensaje)
        ch <- mensaje                    // Enviar al channel
        time.Sleep(500 * time.Millisecond)
    }
    
    close(ch)  // ¡Importante! Cerrar cuando termine
}

func consumidor(ch <-chan string) {  // Receive-only channel
    for mensaje := range ch {        // Recibir hasta que se cierre
        fmt.Printf("📥 Recibido: %s\n", mensaje)
    }
    fmt.Println("✅ Channel cerrado, consumidor terminando")
}

func main() {
    fmt.Println("📡 Demo: Channel básico")
    
    // Crear channel unbuffered
    ch := make(chan string)
    
    // Lanzar productor y consumidor
    go productor(ch)
    go consumidor(ch)
    
    // Esperar a que terminen
    time.Sleep(3 * time.Second)
}
```

### 2. 🎯 Channel con Sincronización

```go
package main

import (
    "fmt"
    "time"
)

func trabajador(id int, trabajos <-chan int, resultados chan<- int) {
    for trabajo := range trabajos {
        fmt.Printf("👷 Worker %d iniciando trabajo %d\n", id, trabajo)
        
        // Simular procesamiento
        tiempo := time.Duration(trabajo) * 100 * time.Millisecond
        time.Sleep(tiempo)
        
        resultado := trabajo * trabajo
        resultados <- resultado
        
        fmt.Printf("✅ Worker %d completó trabajo %d\n", id, trabajo)
    }
}

func ejemploSincronizacion() {
    trabajos := make(chan int, 10)      // Buffered para trabajos
    resultados := make(chan int, 10)    // Buffered para resultados
    
    // Lanzar 3 workers
    for i := 1; i <= 3; i++ {
        go trabajador(i, trabajos, resultados)
    }
    
    // Enviar trabajos
    fmt.Println("📋 Enviando trabajos...")
    for j := 1; j <= 9; j++ {
        trabajos <- j
    }
    close(trabajos)
    
    // Recoger resultados
    fmt.Println("\n📊 Recogiendo resultados:")
    for r := 1; r <= 9; r++ {
        resultado := <-resultados
        fmt.Printf("📈 Resultado %d: %d\n", r, resultado)
    }
}
```

### 3. 📦 Channels Buffered vs Unbuffered

```go
package main

import (
    "fmt"
    "time"
)

func demoUnbuffered() {
    fmt.Println("🔄 Channel Unbuffered (Síncrono)")
    ch := make(chan string)  // Sin buffer
    
    go func() {
        fmt.Println("📤 Enviando 'Hola'...")
        ch <- "Hola"             // Bloquea hasta que alguien reciba
        fmt.Println("📤 'Hola' enviado!")
        
        fmt.Println("📤 Enviando 'Mundo'...")
        ch <- "Mundo"            // Bloquea hasta que alguien reciba
        fmt.Println("📤 'Mundo' enviado!")
        close(ch)
    }()
    
    time.Sleep(1 * time.Second)  // Simular delay en receptor
    
    for msg := range ch {
        fmt.Printf("📥 Recibido: %s\n", msg)
        time.Sleep(1 * time.Second)  // Simular procesamiento
    }
}

func demoBuffered() {
    fmt.Println("\n📦 Channel Buffered (Asíncrono)")
    ch := make(chan string, 3)  // Buffer de 3 elementos
    
    go func() {
        mensajes := []string{"Uno", "Dos", "Tres"}
        for _, msg := range mensajes {
            fmt.Printf("📤 Enviando: %s\n", msg)
            ch <- msg  // No bloquea mientras haya espacio en buffer
            fmt.Printf("✅ %s enviado inmediatamente!\n", msg)
        }
        close(ch)
    }()
    
    time.Sleep(2 * time.Second)  // Delay antes de leer
    
    for msg := range ch {
        fmt.Printf("📥 Recibido: %s\n", msg)
    }
}

func main() {
    demoUnbuffered()
    demoBuffered()
}
```

---

## 🎛️ El Statement SELECT

### 1. 🚀 Select Básico

```go
package main

import (
    "fmt"
    "time"
)

func servidor(puerto string, ch chan<- string) {
    for i := 1; i <= 3; i++ {
        tiempo := time.Duration(i*200) * time.Millisecond
        time.Sleep(tiempo)
        
        mensaje := fmt.Sprintf("Respuesta del %s #%d", puerto, i)
        ch <- mensaje
    }
    close(ch)
}

func ejemploSelectBasico() {
    fmt.Println("🎛️ Select Statement - Múltiples Channels")
    
    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)
    
    // Lanzar servidores concurrentes
    go servidor("Puerto 8080", ch1)
    go servidor("Puerto 8081", ch2)
    go servidor("Puerto 8082", ch3)
    
    // Recibir de cualquier channel que esté listo
    for {
        select {
        case msg1, ok := <-ch1:
            if ok {
                fmt.Printf("🔵 [8080]: %s\n", msg1)
            } else {
                ch1 = nil  // Desactivar channel cerrado
            }
            
        case msg2, ok := <-ch2:
            if ok {
                fmt.Printf("🟢 [8081]: %s\n", msg2)
            } else {
                ch2 = nil  // Desactivar channel cerrado
            }
            
        case msg3, ok := <-ch3:
            if ok {
                fmt.Printf("🟡 [8082]: %s\n", msg3)
            } else {
                ch3 = nil  // Desactivar channel cerrado
            }
        }
        
        // Salir cuando todos los channels estén cerrados
        if ch1 == nil && ch2 == nil && ch3 == nil {
            break
        }
    }
    
    fmt.Println("🏁 Todos los servidores han terminado")
}
```

### 2. ⏰ Select con Timeout

```go
package main

import (
    "fmt"
    "time"
)

func operacionLenta(ch chan<- string) {
    // Simular operación que puede tardar mucho
    time.Sleep(3 * time.Second)
    ch <- "Operación completada"
}

func ejemploSelectTimeout() {
    fmt.Println("⏰ Select con Timeout")
    
    ch := make(chan string)
    go operacionLenta(ch)
    
    select {
    case resultado := <-ch:
        fmt.Printf("✅ %s\n", resultado)
        
    case <-time.After(2 * time.Second):
        fmt.Println("⏰ Timeout: La operación tardó demasiado")
        
    case <-time.After(1 * time.Second):
        fmt.Println("⚡ Timeout rápido: Cancelando operación")
    }
}

func ejemploSelectNonBlocking() {
    fmt.Println("\n🚫 Select Non-blocking (Default Case)")
    
    ch := make(chan string, 1)
    
    // Intentar enviar sin bloquear
    select {
    case ch <- "Mensaje":
        fmt.Println("📤 Mensaje enviado")
    default:
        fmt.Println("❌ Channel lleno, no se pudo enviar")
    }
    
    // Intentar recibir sin bloquear
    select {
    case msg := <-ch:
        fmt.Printf("📥 Recibido: %s\n", msg)
    default:
        fmt.Println("❌ No hay mensajes disponibles")
    }
    
    // Segundo intento de recibir
    select {
    case msg := <-ch:
        fmt.Printf("📥 Recibido: %s\n", msg)
    default:
        fmt.Println("❌ Channel vacío")
    }
}

func main() {
    ejemploSelectBasico()
    ejemploSelectTimeout()
    ejemploSelectNonBlocking()
}
```

---

## 🏗️ Patrones Avanzados con Channels

### 1. 🔄 Pipeline Pattern

```go
package main

import (
    "fmt"
    "math"
)

// Etapa 1: Generar números
func generarNumeros() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 1; i <= 10; i++ {
            ch <- i
        }
    }()
    return ch
}

// Etapa 2: Elevar al cuadrado
func elevarCuadrado(input <-chan int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for num := range input {
            ch <- num * num
        }
    }()
    return ch
}

// Etapa 3: Filtrar pares
func filtrarPares(input <-chan int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for num := range input {
            if num%2 == 0 {
                ch <- num
            }
        }
    }()
    return ch
}

// Etapa 4: Calcular raíz cuadrada
func calcularRaiz(input <-chan int) <-chan float64 {
    ch := make(chan float64)
    go func() {
        defer close(ch)
        for num := range input {
            ch <- math.Sqrt(float64(num))
        }
    }()
    return ch
}

func ejemploPipeline() {
    fmt.Println("🔄 Pipeline Pattern")
    fmt.Println("Números → Cuadrado → Filtrar Pares → Raíz Cuadrada")
    
    // Crear pipeline encadenado
    numeros := generarNumeros()
    cuadrados := elevarCuadrado(numeros)
    pares := filtrarPares(cuadrados)
    raices := calcularRaiz(pares)
    
    // Procesar resultados
    fmt.Println("\n📊 Resultados del pipeline:")
    for resultado := range raices {
        fmt.Printf("%.2f ", resultado)
    }
    fmt.Println()
}
```

### 2. 📊 Fan-Out / Fan-In Pattern

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Trabajo struct {
    ID   int
    Data string
}

type Resultado struct {
    TrabajoID int
    Resultado string
    WorkerID  int
}

func generarTrabajos(numTrabajos int) <-chan Trabajo {
    trabajos := make(chan Trabajo)
    
    go func() {
        defer close(trabajos)
        for i := 1; i <= numTrabajos; i++ {
            trabajo := Trabajo{
                ID:   i,
                Data: fmt.Sprintf("Datos del trabajo %d", i),
            }
            trabajos <- trabajo
        }
    }()
    
    return trabajos
}

func worker(id int, trabajos <-chan Trabajo) <-chan Resultado {
    resultados := make(chan Resultado)
    
    go func() {
        defer close(resultados)
        for trabajo := range trabajos {
            // Simular procesamiento variable
            tiempo := time.Duration(rand.Intn(200)) * time.Millisecond
            time.Sleep(tiempo)
            
            resultado := Resultado{
                TrabajoID: trabajo.ID,
                Resultado: fmt.Sprintf("Procesado: %s", trabajo.Data),
                WorkerID:  id,
            }
            
            resultados <- resultado
        }
    }()
    
    return resultados
}

func merge(resultados ...<-chan Resultado) <-chan Resultado {
    var wg sync.WaitGroup
    merged := make(chan Resultado)
    
    // Función para copiar de un channel al merged
    copiar := func(c <-chan Resultado) {
        defer wg.Done()
        for resultado := range c {
            merged <- resultado
        }
    }
    
    // Lanzar una goroutine por cada channel de entrada
    wg.Add(len(resultados))
    for _, c := range resultados {
        go copiar(c)
    }
    
    // Cerrar merged cuando todas las goroutines terminen
    go func() {
        wg.Wait()
        close(merged)
    }()
    
    return merged
}

func ejemploFanOutFanIn() {
    fmt.Println("📊 Fan-Out / Fan-In Pattern")
    
    rand.Seed(time.Now().UnixNano())
    
    // Generar trabajos (Fan-Out: distribuir)
    trabajos := generarTrabajos(15)
    
    // Crear múltiples workers (Fan-Out)
    numWorkers := 5
    resultados := make([]<-chan Resultado, numWorkers)
    
    fmt.Printf("🏭 Lanzando %d workers...\n", numWorkers)
    for i := 0; i < numWorkers; i++ {
        resultados[i] = worker(i+1, trabajos)
    }
    
    // Combinar todos los resultados (Fan-In)
    merged := merge(resultados...)
    
    // Procesar resultados combinados
    fmt.Println("\n📈 Resultados combinados:")
    contador := 0
    for resultado := range merged {
        contador++
        fmt.Printf("  [Worker %d] Trabajo %d: %s\n", 
            resultado.WorkerID, resultado.TrabajoID, resultado.Resultado)
    }
    
    fmt.Printf("\n✅ Procesados %d trabajos en total\n", contador)
}
```

### 3. 🎯 Quit Channel Pattern

```go
package main

import (
    "fmt"
    "time"
)

func servicioConQuit(quit <-chan bool) {
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()
    
    contador := 0
    
    for {
        select {
        case <-ticker.C:
            contador++
            fmt.Printf("⚡ Servicio activo... tick #%d\n", contador)
            
        case <-quit:
            fmt.Println("🛑 Señal de quit recibida")
            fmt.Println("🧹 Limpiando recursos...")
            time.Sleep(200 * time.Millisecond)  // Simular cleanup
            fmt.Println("✅ Servicio terminado correctamente")
            return
        }
    }
}

func ejemploQuitChannel() {
    fmt.Println("🎯 Quit Channel Pattern")
    
    quit := make(chan bool)
    
    // Lanzar servicio
    go servicioConQuit(quit)
    
    // Dejar que corra por un tiempo
    time.Sleep(3 * time.Second)
    
    // Enviar señal de quit
    fmt.Println("📤 Enviando señal de quit...")
    quit <- true
    
    // Dar tiempo para que termine limpiamente
    time.Sleep(500 * time.Millisecond)
}
```

---

## ⚠️ Problemas Comunes y Soluciones

### 1. 🚫 Deadlock

```go
package main

import "fmt"

// ❌ MAL: Deadlock
func deadlockMalo() {
    ch := make(chan int)
    
    // Esto causará deadlock porque nadie está leyendo
    ch <- 42  // Bloquea para siempre
    
    fmt.Println("Nunca llegará aquí")
}

// ✅ BIEN: Sin deadlock
func deadlockBueno() {
    ch := make(chan int, 1)  // Channel buffered
    
    ch <- 42               // No bloquea porque hay buffer
    value := <-ch          // Leer el valor
    
    fmt.Printf("Valor recibido: %d\n", value)
}

// ✅ MEJOR: Con goroutine
func deadlockMejor() {
    ch := make(chan int)
    
    go func() {
        ch <- 42  // Enviar en goroutine
    }()
    
    value := <-ch  // Recibir en main
    fmt.Printf("Valor recibido: %d\n", value)
}
```

### 2. 💧 Channel Leak

```go
package main

import (
    "fmt"
    "time"
)

// ❌ MAL: Channel leak
func channelLeakMalo() {
    ch := make(chan int)
    
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        // ¡Olvidamos cerrar el channel!
    }()
    
    // Solo leemos algunos valores
    for i := 0; i < 5; i++ {
        fmt.Printf("Valor: %d\n", <-ch)
    }
    
    // La goroutine queda bloqueada para siempre
    // ¡MEMORY LEAK!
}

// ✅ BIEN: Sin leak
func channelLeakBueno() {
    ch := make(chan int)
    quit := make(chan bool)
    
    go func() {
        defer close(ch)
        for i := 0; i < 10; i++ {
            select {
            case ch <- i:
                // Valor enviado exitosamente
            case <-quit:
                fmt.Println("Goroutine terminando por quit")
                return
            }
        }
    }()
    
    // Leer algunos valores
    for i := 0; i < 5; i++ {
        fmt.Printf("Valor: %d\n", <-ch)
    }
    
    // Señalar quit para terminar la goroutine
    close(quit)
    
    // Dar tiempo para que termine
    time.Sleep(100 * time.Millisecond)
}
```

### 3. 🔒 Channel Closed Panic

```go
package main

import "fmt"

// ❌ MAL: Panic por escribir a channel cerrado
func channelClosedMalo() {
    ch := make(chan int, 1)
    
    close(ch)
    
    // ¡PANIC! No se puede escribir a channel cerrado
    ch <- 42
}

// ✅ BIEN: Verificar si está cerrado
func channelClosedBueno() {
    ch := make(chan int, 1)
    
    ch <- 42
    close(ch)
    
    // Leer de forma segura
    if value, ok := <-ch; ok {
        fmt.Printf("Valor recibido: %d\n", value)
    } else {
        fmt.Println("Channel está cerrado")
    }
    
    // Intentar leer de nuevo
    if value, ok := <-ch; ok {
        fmt.Printf("Valor recibido: %d\n", value)
    } else {
        fmt.Println("Channel está cerrado (segunda lectura)")
    }
}
```

---

## ⚡ Optimización de Performance

### 1. 📦 Buffer Sizing

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func benchmarkBufferSize(bufferSize int, numMensajes int) time.Duration {
    ch := make(chan int, bufferSize)
    var wg sync.WaitGroup
    
    start := time.Now()
    
    // Productor
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(ch)
        for i := 0; i < numMensajes; i++ {
            ch <- i
        }
    }()
    
    // Consumidor
    wg.Add(1)
    go func() {
        defer wg.Done()
        for range ch {
            // Simular procesamiento mínimo
            runtime.Gosched()
        }
    }()
    
    wg.Wait()
    return time.Since(start)
}

func ejemploOptimizacionBuffer() {
    fmt.Println("📦 Optimización de Buffer Size")
    
    numMensajes := 10000
    bufferSizes := []int{0, 1, 10, 100, 1000}
    
    for _, size := range bufferSizes {
        duracion := benchmarkBufferSize(size, numMensajes)
        fmt.Printf("Buffer %4d: %v\n", size, duracion)
    }
}
```

### 2. 🎯 Channel Direction Optimization

```go
package main

import "fmt"

// Función que solo envía
func soloEnvia(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}

// Función que solo recibe
func soloRecibe(ch <-chan int) {
    for valor := range ch {
        fmt.Printf("📥 Recibido: %d\n", valor)
    }
}

// Función que transforma (recibe y envía)
func transformador(input <-chan int, output chan<- int) {
    defer close(output)
    for valor := range input {
        output <- valor * 2  // Duplicar
    }
}

func ejemploChannelDirection() {
    fmt.Println("🎯 Channel Direction Optimization")
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go soloEnvia(ch1)
    go transformador(ch1, ch2)
    soloRecibe(ch2)
}
```

---

## 🧪 Testing de Channels

```go
package main

import (
    "testing"
    "time"
)

func productor(ch chan<- int, n int) {
    defer close(ch)
    for i := 0; i < n; i++ {
        ch <- i
    }
}

func consumidor(ch <-chan int) int {
    suma := 0
    for valor := range ch {
        suma += valor
    }
    return suma
}

func TestChannel(t *testing.T) {
    ch := make(chan int)
    
    go productor(ch, 10)
    suma := consumidor(ch)
    
    esperado := 45  // 0+1+2+...+9
    if suma != esperado {
        t.Errorf("Esperado %d, obtenido %d", esperado, suma)
    }
}

func TestChannelTimeout(t *testing.T) {
    ch := make(chan int)
    
    select {
    case <-ch:
        t.Error("No debería recibir nada")
    case <-time.After(100 * time.Millisecond):
        // Correcto, timeout esperado
    }
}

func BenchmarkChannel(b *testing.B) {
    ch := make(chan int, 100)
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        go func() {
            ch <- i
        }()
        <-ch
    }
}
```

---

## 💡 Tips de Experto

### 🎯 Cuándo Usar Cada Tipo de Channel

```go
// ✅ Unbuffered: Sincronización estricta
handshake := make(chan bool)

// ✅ Buffered pequeño: Suavizar picos
notificaciones := make(chan string, 10)

// ✅ Buffered grande: Desacoplar productor/consumidor
cola := make(chan Trabajo, 1000)

// ✅ Channel de channels: Multiplexación dinámica
canales := make(chan chan int)
```

### 🔧 Patrones de Cancelación

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func trabajadorConContext(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("👷 Worker %d: cancelado\n", id)
            return
        case <-time.After(200 * time.Millisecond):
            fmt.Printf("👷 Worker %d: trabajando...\n", id)
        }
    }
}

func ejemploCancelacion() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    for i := 1; i <= 3; i++ {
        go trabajadorConContext(ctx, i)
    }
    
    time.Sleep(3 * time.Second)
}
```

---

## ⚠️ Errores Comunes y Cómo Evitarlos

### 1. 🚫 Channel Nil

```go
// ❌ MAL: Usar channel nil
var ch chan int
ch <- 42  // Bloquea para siempre

// ✅ BIEN: Inicializar channel
ch = make(chan int)
ch <- 42
```

### 2. 🔄 Range en Channel No Cerrado

```go
// ❌ MAL: Range sin close
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    // Olvida cerrar
}()

for v := range ch {  // Deadlock eventual
    fmt.Println(v)
}

// ✅ BIEN: Siempre cerrar
ch := make(chan int)
go func() {
    defer close(ch)  // ¡Importante!
    ch <- 1
    ch <- 2
}()

for v := range ch {
    fmt.Println(v)
}
```

---

## 🎯 Resumen de la Lección

### ✅ Conceptos Clave Aprendidos

1. **📡 Channels**: Tuberías de comunicación entre goroutines
2. **🔄 Buffered vs Unbuffered**: Síncrono vs asíncrono
3. **🎛️ Select Statement**: Multiplexación de channels
4. **🏗️ Patrones**: Pipeline, Fan-Out/Fan-In, Quit
5. **⚠️ Problemas**: Deadlocks, leaks, panics
6. **⚡ Optimización**: Buffer sizing, direcciones
7. **🧪 Testing**: Verificación de comportamiento concurrente
8. **💡 Best Practices**: Cuándo y cómo usar channels

### 🏆 Logros Desbloqueados

- [ ] 🥇 **Channel Novice**: Primer channel exitoso
- [ ] 🥈 **Select Master**: Dominar multiplexación
- [ ] 🥉 **Pipeline Engineer**: Crear pipeline complejo
- [ ] 🏅 **Fan-Out Expert**: Implementar distribución de trabajo
- [ ] 🎖️ **Deadlock Detector**: Identificar y evitar bloqueos

### 📚 Próximos Pasos

En la **Lección 15: Context Package**, aprenderemos:
- Context para cancelación
- Propagación de deadlines
- Valores en contexto
- Best practices con context

---

**🎉 ¡Felicitaciones! Has dominado los channels de Go. Ahora puedes crear sistemas de comunicación robustos y eficientes entre goroutines, aplicando los patrones fundamentales de la programación concurrente.**

*Recuerda: "Los channels son las tuberías que conectan las goroutines concurrentes" - Aprende a diseñar estas conexiones y crearás arquitecturas escalables.*
