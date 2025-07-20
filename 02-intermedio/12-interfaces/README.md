# 🔌 Lección 12: Interfaces en Go

## 🎯 Objetivos de Aprendizaje

Al completar esta lección, podrás:

- ✅ Entender qué son las interfaces y por qué son fundamentales en Go
- ✅ Declarar e implementar interfaces de diferentes tipos
- ✅ Aplicar polimorfismo usando interfaces
- ✅ Trabajar con interfaces embebidas (interface embedding)
- ✅ Usar la interface{} (empty interface) eficientemente
- ✅ Implementar type assertions y type switches
- ✅ Aplicar patrones de diseño con interfaces
- ✅ Desarrollar un sistema modular usando interfaces

---

## 📚 Tabla de Contenidos

1. [¿Qué son las Interfaces?](#-qué-son-las-interfaces)
2. [Declaración e Implementación](#-declaración-e-implementación)
3. [Polimorfismo con Interfaces](#-polimorfismo-con-interfaces)
4. [Interface Embedding](#-interface-embedding)
5. [Empty Interface (interface{})](#-empty-interface-interface)
6. [Type Assertions y Type Switches](#-type-assertions-y-type-switches)
7. [Interfaces Estándar de Go](#-interfaces-estándar-de-go)
8. [Patrones de Diseño con Interfaces](#-patrones-de-diseño-con-interfaces)
9. [Best Practices](#-best-practices)
10. [Errores Comunes](#-errores-comunes)

---

## 🔍 ¿Qué son las Interfaces?

### 🧠 Analogía: Las Interfaces como Contratos

Imagina que las **interfaces** son como **contratos laborales**:

```
📋 CONTRATO: "Conductor"
┌─────────────────────────────┐
│ REQUISITOS OBLIGATORIOS:    │
│ ✅ Saber acelerar()         │
│ ✅ Saber frenar()           │
│ ✅ Saber girar()            │
│ ✅ Saber estacionar()       │
└─────────────────────────────┘
        ⬇️
🚗 CANDIDATOS VÁLIDOS:
│ ├── ConductorAuto ✅
│ ├── ConductorMoto ✅  
│ ├── ConductorCamión ✅
│ └── ConductorBus ✅
```

### 💡 Concepto Clave

```go
// Una interface define QUÉ debe hacer algo
// NO CÓMO lo hace
type Conductor interface {
    Acelerar() error
    Frenar() error
    Girar(direccion string) error
    Estacionar() error
}
```

---

## 📝 Declaración e Implementación

### 🔧 Sintaxis Básica

```go
// Declarar una interface
type NombreInterface interface {
    Metodo1() TipoRetorno
    Metodo2(parametro Tipo) (Tipo, error)
    Metodo3(param1 Tipo, param2 Tipo) Tipo
}
```

### 📖 Ejemplo Completo: Sistema de Formas

```go
package main

import (
    "fmt"
    "math"
)

// 1. Declarar la interface
type Forma interface {
    Area() float64
    Perimetro() float64
    String() string
}

// 2. Implementar la interface con diferentes tipos
type Rectangulo struct {
    Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
    return 2 * (r.Ancho + r.Alto)
}

func (r Rectangulo) String() string {
    return fmt.Sprintf("Rectángulo(%.1fx%.1f)", r.Ancho, r.Alto)
}

type Circulo struct {
    Radio float64
}

func (c Circulo) Area() float64 {
    return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
    return 2 * math.Pi * c.Radio
}

func (c Circulo) String() string {
    return fmt.Sprintf("Círculo(r=%.1f)", c.Radio)
}

// 3. Usar polimorfismo
func MostrarInfo(f Forma) {
    fmt.Printf("%s:\n", f.String())
    fmt.Printf("  Área: %.2f\n", f.Area())
    fmt.Printf("  Perímetro: %.2f\n", f.Perimetro())
}

func main() {
    formas := []Forma{
        Rectangulo{Ancho: 5, Alto: 3},
        Circulo{Radio: 2.5},
        Rectangulo{Ancho: 10, Alto: 10},
    }
    
    for _, forma := range formas {
        MostrarInfo(forma)
        fmt.Println()
    }
}
```

### 🎯 Output:
```
Rectángulo(5.0x3.0):
  Área: 15.00
  Perímetro: 16.00

Círculo(r=2.5):
  Área: 19.63
  Perímetro: 15.71

Rectángulo(10.0x10.0):
  Área: 100.00
  Perímetro: 40.00
```

---

## 🔄 Polimorfismo con Interfaces

### 💡 Concepto: "Una Interface, Múltiples Implementaciones"

```go
// Interface común
type Reproductor interface {
    Play() string
    Pause() string
    Stop() string
    GetTitulo() string
}

// Implementación para música
type ReproductorMusica struct {
    Cancion string
    Artista string
}

func (rm ReproductorMusica) Play() string {
    return fmt.Sprintf("🎵 Reproduciendo: %s - %s", rm.Cancion, rm.Artista)
}

func (rm ReproductorMusica) Pause() string {
    return "⏸️ Música pausada"
}

func (rm ReproductorMusica) Stop() string {
    return "⏹️ Música detenida"
}

func (rm ReproductorMusica) GetTitulo() string {
    return rm.Cancion
}

// Implementación para video
type ReproductorVideo struct {
    Titulo   string
    Duracion string
}

func (rv ReproductorVideo) Play() string {
    return fmt.Sprintf("🎬 Reproduciendo video: %s (%s)", rv.Titulo, rv.Duracion)
}

func (rv ReproductorVideo) Pause() string {
    return "⏸️ Video pausado"
}

func (rv ReproductorVideo) Stop() string {
    return "⏹️ Video detenido"
}

func (rv ReproductorVideo) GetTitulo() string {
    return rv.Titulo
}

// Función que acepta cualquier reproductor
func ControlarReproduccion(r Reproductor) {
    fmt.Println(r.Play())
    fmt.Println(r.Pause())
    fmt.Println(r.Stop())
}
```

---

## 🔗 Interface Embedding

### 💡 Concepto: Combinar Interfaces

```go
// Interfaces básicas
type Lector interface {
    Read([]byte) (int, error)
}

type Escritor interface {
    Write([]byte) (int, error)
}

type Cerrador interface {
    Close() error
}

// Interface compuesta (embedding)
type LectorEscritor interface {
    Lector
    Escritor
}

type ReadWriteCloser interface {
    Lector
    Escritor
    Cerrador
}

// Implementación
type Archivo struct {
    nombre string
    datos  []byte
    pos    int
    abierto bool
}

func (a *Archivo) Read(p []byte) (int, error) {
    if !a.abierto {
        return 0, fmt.Errorf("archivo cerrado")
    }
    
    n := copy(p, a.datos[a.pos:])
    a.pos += n
    return n, nil
}

func (a *Archivo) Write(p []byte) (int, error) {
    if !a.abierto {
        return 0, fmt.Errorf("archivo cerrado")
    }
    
    a.datos = append(a.datos, p...)
    return len(p), nil
}

func (a *Archivo) Close() error {
    a.abierto = false
    return nil
}

func NuevoArchivo(nombre string) *Archivo {
    return &Archivo{
        nombre:  nombre,
        datos:   make([]byte, 0),
        abierto: true,
    }
}
```

---

## 🌐 Empty Interface (interface{})

### 💡 Concepto: La Interface Universal

```go
// interface{} puede contener CUALQUIER tipo
func MostrarTipo(valor interface{}) {
    fmt.Printf("Valor: %v, Tipo: %T\n", valor, valor)
}

func main() {
    valores := []interface{}{
        42,
        "Hola mundo",
        3.14,
        []int{1, 2, 3},
        map[string]int{"a": 1},
        true,
    }
    
    for _, valor := range valores {
        MostrarTipo(valor)
    }
}
```

### 🎯 Output:
```
Valor: 42, Tipo: int
Valor: Hola mundo, Tipo: string
Valor: 3.14, Tipo: float64
Valor: [1 2 3], Tipo: []int
Valor: map[a:1], Tipo: map[string]int
Valor: true, Tipo: bool
```

---

## 🔍 Type Assertions y Type Switches

### 🔧 Type Assertions

```go
func ProcesarValor(valor interface{}) {
    // Type assertion con verificación
    if str, ok := valor.(string); ok {
        fmt.Printf("Es string: %s (longitud: %d)\n", str, len(str))
        return
    }
    
    if num, ok := valor.(int); ok {
        fmt.Printf("Es int: %d (doble: %d)\n", num, num*2)
        return
    }
    
    fmt.Printf("Tipo no manejado: %T\n", valor)
}
```

### 🔀 Type Switches

```go
func AnalizarTipo(valor interface{}) {
    switch v := valor.(type) {
    case string:
        fmt.Printf("String: '%s' (longitud: %d)\n", v, len(v))
    case int:
        fmt.Printf("Int: %d (par: %t)\n", v, v%2 == 0)
    case float64:
        fmt.Printf("Float64: %.2f (entero: %.0f)\n", v, v)
    case []int:
        fmt.Printf("Slice de int: %v (suma: %d)\n", v, sumarSlice(v))
    case bool:
        fmt.Printf("Bool: %t\n", v)
    default:
        fmt.Printf("Tipo desconocido: %T\n", v)
    }
}

func sumarSlice(nums []int) int {
    suma := 0
    for _, num := range nums {
        suma += num
    }
    return suma
}
```

---

## 📦 Interfaces Estándar de Go

### 🔧 fmt.Stringer

```go
type Persona struct {
    Nombre string
    Edad   int
}

// Implementar fmt.Stringer
func (p Persona) String() string {
    return fmt.Sprintf("%s (%d años)", p.Nombre, p.Edad)
}

func main() {
    p := Persona{"Ana", 25}
    fmt.Println(p) // Usa automáticamente String()
}
```

### 📊 sort.Interface

```go
type Producto struct {
    Nombre string
    Precio float64
}

type PorPrecio []Producto

func (p PorPrecio) Len() int           { return len(p) }
func (p PorPrecio) Less(i, j int) bool { return p[i].Precio < p[j].Precio }
func (p PorPrecio) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
    productos := PorPrecio{
        {"Laptop", 1200.00},
        {"Mouse", 25.50},
        {"Teclado", 75.00},
    }
    
    sort.Sort(productos)
    fmt.Println("Ordenados por precio:", productos)
}
```

### 📝 io.Reader/Writer

```go
type Buffer struct {
    datos []byte
}

func (b *Buffer) Read(p []byte) (int, error) {
    if len(b.datos) == 0 {
        return 0, io.EOF
    }
    
    n := copy(p, b.datos)
    b.datos = b.datos[n:]
    return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
    b.datos = append(b.datos, p...)
    return len(p), nil
}
```

---

## 🏗️ Patrones de Diseño con Interfaces

### 🔧 Strategy Pattern

```go
// Interface para estrategias de descuento
type EstrategiaDescuento interface {
    AplicarDescuento(precio float64) float64
    Descripcion() string
}

// Implementaciones concretas
type DescuentoRegular struct{}

func (d DescuentoRegular) AplicarDescuento(precio float64) float64 {
    return precio * 0.95 // 5% descuento
}

func (d DescuentoRegular) Descripcion() string {
    return "Descuento regular (5%)"
}

type DescuentoPremium struct{}

func (d DescuentoPremium) AplicarDescuento(precio float64) float64 {
    return precio * 0.85 // 15% descuento
}

func (d DescuentoPremium) Descripcion() string {
    return "Descuento premium (15%)"
}

// Contexto que usa la estrategia
type CarritoCompras struct {
    items     []float64
    estrategia EstrategiaDescuento
}

func (c *CarritoCompras) SetEstrategia(e EstrategiaDescuento) {
    c.estrategia = e
}

func (c *CarritoCompras) CalcularTotal() float64 {
    total := 0.0
    for _, precio := range c.items {
        total += c.estrategia.AplicarDescuento(precio)
    }
    return total
}
```

### 🏭 Factory Pattern

```go
type BaseDatos interface {
    Conectar() error
    Ejecutar(query string) ([]map[string]interface{}, error)
    Cerrar() error
}

type FactoryDB interface {
    CrearDB(config map[string]string) BaseDatos
}

type MySQLFactory struct{}

func (f MySQLFactory) CrearDB(config map[string]string) BaseDatos {
    return &MySQL{
        host:     config["host"],
        puerto:   config["puerto"],
        database: config["database"],
    }
}

type PostgreSQLFactory struct{}

func (f PostgreSQLFactory) CrearDB(config map[string]string) BaseDatos {
    return &PostgreSQL{
        host:     config["host"],
        puerto:   config["puerto"],
        database: config["database"],
    }
}
```

---

## ✅ Best Practices

### 📏 1. Interfaces Pequeñas

```go
// ✅ BIEN: Interface pequeña y específica
type Leible interface {
    Read([]byte) (int, error)
}

// ❌ MAL: Interface demasiado grande
type ServicioCompleto interface {
    Leer() error
    Escribir() error
    Validar() error
    Procesar() error
    Guardar() error
    Enviar() error
    // ... muchos más métodos
}
```

### 🎯 2. Acepta Interfaces, Retorna Tipos Concretos

```go
// ✅ BIEN: Acepta interface (flexible)
func ProcesarDatos(r io.Reader) (*Resultado, error) {
    // Implementación
}

// ❌ MAL: Acepta tipo concreto (rígido)
func ProcesarArchivo(archivo *os.File) (*Resultado, error) {
    // Implementación
}
```

### 🔧 3. Definir Interfaces en el Consumidor

```go
// ✅ BIEN: Interface definida donde se usa
package procesador

type Almacenador interface {
    Guardar(datos []byte) error
}

func Procesar(datos []byte, storage Almacenador) error {
    // Procesar datos
    return storage.Guardar(datos)
}

// El paquete que implementa NO define la interface
package mysql

type DB struct {
    // campos
}

func (db *DB) Guardar(datos []byte) error {
    // Implementación específica de MySQL
}
```

---

## ❌ Errores Comunes

### 🚫 1. Interfaces Innecesarias

```go
// ❌ MAL: Interface con un solo implementador
type UsuarioRepositorio interface {
    CrearUsuario(u Usuario) error
    ObtenerUsuario(id int) (Usuario, error)
}

type MySQLUsuarioRepo struct{}

// Si solo tienes una implementación, no necesitas interface
```

### 🚫 2. Interface{} Excesivo

```go
// ❌ MAL: Uso innecesario de interface{}
func Procesar(datos interface{}) interface{} {
    // Pierdes type safety
}

// ✅ BIEN: Usa tipos específicos cuando sea posible
func ProcesarTexto(texto string) string {
    // Type safe
}
```

### 🚫 3. Type Assertions Sin Verificación

```go
// ❌ MAL: Puede causar panic
func MalManejo(valor interface{}) {
    str := valor.(string) // ¡Panic si no es string!
    fmt.Println(str)
}

// ✅ BIEN: Verificación segura
func BuenManejo(valor interface{}) {
    if str, ok := valor.(string); ok {
        fmt.Println(str)
    } else {
        fmt.Println("No es un string")
    }
}
```

---

## 🎯 Casos de Uso Reales

### 🔧 1. Sistema de Logging

```go
type Logger interface {
    Log(level string, message string)
    Error(message string)
    Info(message string)
    Debug(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{ archivo string }
type RemoteLogger struct{ endpoint string }
```

### 📊 2. Procesadores de Datos

```go
type DataProcessor interface {
    Process(data []byte) ([]byte, error)
    Validate(data []byte) error
}

type JSONProcessor struct{}
type XMLProcessor struct{}
type CSVProcessor struct{}
```

### 🔒 3. Sistemas de Autenticación

```go
type Authenticator interface {
    Authenticate(username, password string) (User, error)
    ValidateToken(token string) (User, error)
}

type JWTAuth struct{}
type OAuthAuth struct{}
type LDAPAuth struct{}
```

---

## 🚀 Resumen de la Lección

### ✅ Lo que has aprendido:

1. **Conceptos Fundamentales**: Qué son las interfaces y por qué son importantes
2. **Implementación Implícita**: Cómo Go maneja las interfaces automáticamente
3. **Polimorfismo**: Una interface, múltiples implementaciones
4. **Interface Embedding**: Combinar interfaces para crear abstracciones complejas
5. **Empty Interface**: Manejo de tipos dinámicos con interface{}
6. **Type Assertions/Switches**: Trabajar con tipos dinámicos de forma segura
7. **Interfaces Estándar**: fmt.Stringer, sort.Interface, io.Reader/Writer
8. **Patrones de Diseño**: Strategy, Factory y otros patrones con interfaces
9. **Best Practices**: Interfaces pequeñas, acepta interfaces/retorna concretos
10. **Errores Comunes**: Qué evitar al trabajar con interfaces

### 🎯 Próximos Pasos:

- Practicar con los ejercicios incluidos
- Implementar el proyecto de sistema de plugins
- Explorar interfaces más avanzadas en librerías estándar
- Estudiar cómo las interfaces facilitan el testing y mocking

---

**¡Las interfaces son el corazón del polimorfismo en Go!** 🚀
