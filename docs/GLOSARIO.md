# 📚 Glosario - Go Deep

*Terminología completa para tu journey en Go*

Este glosario contiene todos los términos técnicos, conceptos y jerga que encontrarás en el mundo Go. Úsalo como referencia rápida durante tu aprendizaje.

---

## A

### **API (Application Programming Interface)**
Conjunto de reglas y protocolos que permite a diferentes aplicaciones comunicarse entre sí.
```go
// Ejemplo: API REST endpoint
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Lógica del endpoint
}
```

### **Array**
Colección de elementos del mismo tipo con tamaño fijo.
```go
var numbers [5]int // Array de 5 enteros
```

### **Assignment (Asignación)**
Proceso de dar un valor a una variable.
```go
x := 42        // Short variable declaration
var y int = 42 // Explicit assignment
```

---

## B

### **Backing Array**
Array subyacente que almacena los datos de un slice.
```go
arr := [10]int{1,2,3,4,5,6,7,8,9,10}
slice := arr[2:5] // slice usa arr como backing array
```

### **Binary**
Archivo ejecutable compilado desde código Go.
```bash
go build main.go  # Genera binario ejecutable
```

### **Boolean**
Tipo de dato que puede ser `true` o `false`.
```go
var isActive bool = true
```

### **Buffered Channel**
Channel con capacidad para almacenar múltiples valores antes de bloquearse.
```go
ch := make(chan int, 5) // Buffer de 5 elementos
```

---

## C

### **Channel**
Mecanismo de comunicación entre goroutines.
```go
ch := make(chan string)
go func() { ch <- "mensaje" }()
msg := <-ch
```

### **Closure**
Función anónima que captura variables de su scope exterior.
```go
func outer() func() int {
    x := 0
    return func() int { // closure
        x++
        return x
    }
}
```

### **Compilation**
Proceso de convertir código Go a código máquina ejecutable.
```bash
go build    # Compila el package actual
go install  # Compila e instala
```

### **Concurrency**
Capacidad de manejar múltiples tareas aparentemente al mismo tiempo.
```go
go doWork() // Ejecuta concurrentemente
```

### **Constant**
Valor que no puede cambiar durante la ejecución.
```go
const Pi = 3.14159
const MaxUsers = 100
```

---

## D

### **Deadlock**
Situación donde goroutines se bloquean mutuamente esperando recursos.
```go
// DEADLOCK: goroutines esperándose mutuamente
ch1 := make(chan int)
ch2 := make(chan int)
go func() { ch1 <- <-ch2 }()
go func() { ch2 <- <-ch1 }()
```

### **Defer**
Palabra clave que postpone la ejecución de una función hasta que la función actual retorna.
```go
func example() {
    defer fmt.Println("Esto se ejecuta al final")
    fmt.Println("Esto se ejecuta primero")
}
```

---

## E

### **Embedding**
Técnica para incluir un tipo dentro de otro tipo.
```go
type Person struct {
    Name string
}

type Employee struct {
    Person  // Embedding
    ID      int
}
```

### **Error Interface**
Interface built-in para manejo de errores.
```go
type error interface {
    Error() string
}
```

### **Exported**
Identificador que comienza con mayúscula y es visible fuera del package.
```go
func PublicFunction() {} // Exported
func privateFunction() {} // Unexported
```

---

## F

### **Function**
Bloque de código reutilizable que realiza una tarea específica.
```go
func add(a, b int) int {
    return a + b
}
```

### **Function Literal**
Función anónima definida inline.
```go
func(x int) int { return x * 2 }
```

---

## G

### **Garbage Collection (GC)**
Proceso automático de liberación de memoria no utilizada.

### **gofmt**
Herramienta oficial para formatear código Go automáticamente.
```bash
gofmt -w main.go  # Formatea y sobrescribe
```

### **Goroutine**
Función que se ejecuta concurrentemente con otras goroutines.
```go
go func() {
    fmt.Println("Ejecutándose en goroutine")
}()
```

### **GOPATH**
Variable de entorno que define el workspace de Go (legacy).

### **Go Modules**
Sistema moderno de gestión de dependencias en Go.
```bash
go mod init myproject
go mod tidy
```

---

## H

### **Handler**
Función que maneja solicitudes HTTP.
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
}
```

---

## I

### **Interface**
Tipo que define un conjunto de signatures de métodos.
```go
type Writer interface {
    Write([]byte) (int, error)
}
```

### **iota**
Identificador que genera constantes incrementales.
```go
const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
)
```

---

## J

### **JSON**
Formato de intercambio de datos. Go tiene excelente soporte built-in.
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

---

## L

### **Literal**
Valor escrito directamente en el código.
```go
42          // Integer literal
"hello"     // String literal
true        // Boolean literal
```

---

## M

### **Map**
Estructura de datos key-value.
```go
m := make(map[string]int)
m["key"] = 42
```

### **Method**
Función con un receiver que le permite ser llamada en un tipo específico.
```go
func (p Person) String() string {
    return p.Name
}
```

### **Mutex**
Mecanismo de sincronización para prevenir race conditions.
```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

---

## N

### **nil**
Valor zero para pointers, slices, maps, channels, functions e interfaces.
```go
var p *int = nil
var s []int = nil
```

---

## P

### **Package**
Unidad de organización y reutilización de código.
```go
package main

import "fmt"
```

### **Panic**
Función built-in que causa que el programa termine abruptamente.
```go
if err != nil {
    panic(err)
}
```

### **Pointer**
Variable que almacena la dirección de memoria de otra variable.
```go
x := 42
p := &x  // p es pointer a x
fmt.Println(*p) // 42, dereference pointer
```

---

## R

### **Race Condition**
Error que ocurre cuando múltiples goroutines acceden concurrentemente a datos compartidos.

### **Receiver**
Parámetro especial en la definición de métodos que especifica el tipo.
```go
func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

### **Rune**
Alias para int32, representa un code point Unicode.
```go
var r rune = 'A'  // Unicode code point
```

---

## S

### **Slice**
Array dinámico con longitud variable.
```go
s := []int{1, 2, 3, 4, 5}
s = append(s, 6) // Agrega elemento
```

### **String**
Secuencia inmutable de bytes (UTF-8).
```go
s := "Hello, 世界"
```

### **Struct**
Tipo compuesto que agrupa campos relacionados.
```go
type Person struct {
    Name string
    Age  int
}
```

---

## T

### **Type Assertion**
Operación que extrae el valor subyacente de una interface.
```go
var i interface{} = "hello"
s := i.(string) // Type assertion
```

### **Type Switch**
Switch statement que opera sobre tipos.
```go
switch v := x.(type) {
case string:
    // v es string
case int:
    // v es int
}
```

---

## V

### **Variable**
Ubicación de memoria con un nombre que almacena un valor.
```go
var name string = "Go"
age := 25
```

### **Variadic Function**
Función que acepta un número variable de argumentos.
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

---

## W

### **WaitGroup**
Primitiva de sincronización que espera a que un conjunto de goroutines termine.
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // trabajo
}()
wg.Wait()
```

---

## Z

### **Zero Value**
Valor por defecto de cualquier tipo cuando se declara sin inicializar.
- `int`: 0
- `string`: ""
- `bool`: false
- `pointer`: nil
- `slice`: nil

```go
var x int    // x = 0 (zero value)
var s string // s = "" (zero value)
```

---

## 🎯 Términos Específicos de Go

### **Go Generate**
Herramienta para generar código automáticamente.
```go
//go:generate stringer -type=Day
```

### **Build Tags**
Comentarios especiales que controlan cuándo se compila un archivo.
```go
// +build linux,386
```

### **CGO**
Característica que permite llamar código C desde Go.
```go
/*
#include <stdio.h>
*/
import "C"
```

---

## 🔧 Herramientas y Comandos

### **go build**
Compila packages y dependencias.
```bash
go build main.go
go build -o myapp
```

### **go run**
Compila y ejecuta un programa Go.
```bash
go run main.go
```

### **go test**
Ejecuta tests y benchmarks.
```bash
go test ./...
go test -v -cover
```

### **go get**
Descarga e instala packages.
```bash
go get github.com/gin-gonic/gin
go get -u ./...  # actualiza dependencias
```

### **go mod**
Gestiona módulos Go.
```bash
go mod init myproject
go mod tidy
go mod vendor
```

---

## 💡 Patrones y Conceptos Avanzados

### **Builder Pattern**
Patrón para construir objetos complejos paso a paso.

### **Dependency Injection**
Técnica para proveer dependencias a un objeto en lugar de que las cree.

### **Middleware**
Función que procesa requests antes de que lleguen al handler final.

### **Pipeline Pattern**
Patrón que procesa datos a través de una serie de stages.

### **Fan-out/Fan-in**
Patrón de concurrencia para distribuir trabajo y reunir resultados.

---

## 📚 Recursos para Profundizar

- **[Go Language Specification](https://golang.org/ref/spec)** - Especificación oficial completa
- **[Effective Go](https://golang.org/doc/effective_go.html)** - Guía de mejores prácticas
- **[Go Memory Model](https://golang.org/ref/mem)** - Modelo de memoria y concurrencia

---

*¿Falta algún término? [Contribuye al glosario](../CONTRIBUTING.md) 📝*
