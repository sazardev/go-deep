# 🎭 Lección 12: Métodos en Go

> **Nivel**: Fundamentos  
> **Duración estimada**: 4-5 horas  
> **Prerrequisitos**: Structs, variables, funciones

## 🎯 Objetivos de Aprendizaje

Al finalizar esta lección, podrás:

- ✅ **Definir métodos** en tipos custom y structs
- ✅ **Entender la diferencia** entre value receivers y pointer receivers
- ✅ **Implementar method chaining** para APIs fluent
- ✅ **Usar métodos** para encapsulación y abstracción
- ✅ **Aplicar patrones de diseño** con métodos
- ✅ **Crear constructores** y factory methods
- ✅ **Validar datos** con métodos de validación
- ✅ **Serializar objetos** con métodos String() y JSON

---

## 📚 Tabla de Contenidos

1. [¿Qué son los Métodos?](#-1-qué-son-los-métodos)
2. [Sintaxis y Declaración](#-2-sintaxis-y-declaración)
3. [Value vs Pointer Receivers](#-3-value-vs-pointer-receivers)
4. [Métodos en Tipos Custom](#-4-métodos-en-tipos-custom)
5. [Method Chaining](#-5-method-chaining)
6. [Constructores y Factory Methods](#-6-constructores-y-factory-methods)
7. [Métodos de Validación](#-7-métodos-de-validación)
8. [Serialización y String()](#-8-serialización-y-string)
9. [Patrones y Best Practices](#-9-patrones-y-best-practices)
10. [Ejercicios Prácticos](#-10-ejercicios-prácticos)

---

## 🔍 1. ¿Qué son los Métodos?

### 🧠 Analogía: Métodos como Verbos

Imagina que los **structs** son **sustantivos** y los **métodos** son **verbos**:

- **Struct**: `Perro` (sustantivo)
- **Métodos**: `Ladrar()`, `Correr()`, `Comer()` (verbos)

```go
type Perro struct {
    Nombre string
    Edad   int
    Raza   string
}

// Métodos = acciones que puede hacer el perro
func (p Perro) Ladrar() {
    fmt.Printf("%s está ladrando: ¡Guau guau!\n", p.Nombre)
}

func (p Perro) Correr() {
    fmt.Printf("%s está corriendo muy rápido\n", p.Nombre)
}

func (p *Perro) Envejecer() {
    p.Edad++
    fmt.Printf("%s ahora tiene %d años\n", p.Nombre, p.Edad)
}
```

### 🔄 Métodos vs Funciones

| Aspecto | Función | Método |
|---------|---------|--------|
| **Sintaxis** | `func Nombre(params) {}` | `func (receiver) Nombre(params) {}` |
| **Llamada** | `Funcion(obj, params)` | `obj.Metodo(params)` |
| **Encapsulación** | Menos natural | Más intuitivo |
| **Polimorfismo** | No | Sí (via interfaces) |

```go
// Como función
func LadrarFuncion(p Perro) {
    fmt.Printf("%s ladra\n", p.Nombre)
}

// Como método  
func (p Perro) Ladrar() {
    fmt.Printf("%s ladra\n", p.Nombre)
}

// Uso
perro := Perro{Nombre: "Rex"}
LadrarFuncion(perro)  // Estilo funcional
perro.Ladrar()        // Estilo OOP - más natural
```

---

## 🛠️ 2. Sintaxis y Declaración

### 📝 Sintaxis Básica

```go
func (receiver TipoReceptor) NombreMetodo(parametros) tipoRetorno {
    // Cuerpo del método
}
```

**Componentes:**
- **receiver**: El tipo al que pertenece el método
- **TipoReceptor**: El tipo (struct, custom type, etc.)
- **NombreMetodo**: Nombre del método (PascalCase para público)
- **parametros**: Parámetros del método
- **tipoRetorno**: Tipo de retorno (opcional)

### 💻 Ejemplo Completo

```go
package main

import (
    "fmt"
    "math"
)

// Struct para representar un círculo
type Circulo struct {
    Radio float64
    Color string
}

// Método para calcular área (value receiver)
func (c Circulo) Area() float64 {
    return math.Pi * c.Radio * c.Radio
}

// Método para calcular perímetro (value receiver)
func (c Circulo) Perimetro() float64 {
    return 2 * math.Pi * c.Radio
}

// Método para cambiar color (pointer receiver)
func (c *Circulo) CambiarColor(nuevoColor string) {
    c.Color = nuevoColor
}

// Método para escalar el círculo (pointer receiver)
func (c *Circulo) Escalar(factor float64) {
    c.Radio *= factor
}

// Método String() para representación como string
func (c Circulo) String() string {
    return fmt.Sprintf("Círculo(radio=%.2f, color=%s)", c.Radio, c.Color)
}

func main() {
    // Crear círculo
    circulo := Circulo{Radio: 5.0, Color: "rojo"}
    
    // Usar métodos
    fmt.Println("Círculo:", circulo)
    fmt.Printf("Área: %.2f\n", circulo.Area())
    fmt.Printf("Perímetro: %.2f\n", circulo.Perimetro())
    
    // Modificar círculo
    circulo.CambiarColor("azul")
    circulo.Escalar(2.0)
    
    fmt.Println("Después de modificar:", circulo)
    fmt.Printf("Nueva área: %.2f\n", circulo.Area())
}
```

---

## ⚖️ 3. Value vs Pointer Receivers

### 🔍 Diferencias Fundamentales

| Aspecto | Value Receiver | Pointer Receiver |
|---------|----------------|------------------|
| **Sintaxis** | `(v Tipo)` | `(v *Tipo)` |
| **Modifica original** | ❌ No | ✅ Sí |
| **Performance** | Copia el valor | Referencia directa |
| **Memory usage** | Más memoria | Menos memoria |
| **Nil safety** | ✅ Seguro | ⚠️ Puede ser nil |

### 💡 Cuándo Usar Cada Uno

```go
type Persona struct {
    Nombre   string
    Edad     int
    Saldo    float64
    Activa   bool
}

// ✅ VALUE RECEIVER - Cuando:
// 1. No necesitas modificar el struct
// 2. El struct es pequeño
// 3. Operaciones de consulta/lectura

func (p Persona) EsMayorDeEdad() bool {
    return p.Edad >= 18
}

func (p Persona) NombreCompleto() string {
    return fmt.Sprintf("Sr/a. %s", p.Nombre)
}

func (p Persona) PuedePagar(monto float64) bool {
    return p.Saldo >= monto && p.Activa
}

// ✅ POINTER RECEIVER - Cuando:
// 1. Necesitas modificar el struct
// 2. El struct es grande (evitar copias)
// 3. Consistencia en el tipo

func (p *Persona) Cumpleanos() {
    p.Edad++
    fmt.Printf("¡Feliz cumpleaños! %s ahora tiene %d años\n", p.Nombre, p.Edad)
}

func (p *Persona) Depositar(monto float64) error {
    if monto <= 0 {
        return fmt.Errorf("el monto debe ser positivo")
    }
    p.Saldo += monto
    return nil
}

func (p *Persona) Retirar(monto float64) error {
    if !p.PuedePagar(monto) {
        return fmt.Errorf("fondos insuficientes o cuenta inactiva")
    }
    p.Saldo -= monto
    return nil
}

func (p *Persona) Desactivar() {
    p.Activa = false
    fmt.Printf("Cuenta de %s desactivada\n", p.Nombre)
}
```

### ⚠️ Importante: Consistencia

```go
// ❌ MAL - Mezclar value y pointer receivers
type Contador struct {
    valor int
}

func (c Contador) Obtener() int { return c.valor }      // value
func (c *Contador) Incrementar() { c.valor++ }          // pointer

// ✅ BIEN - Ser consistente
type ContadorMejorado struct {
    valor int
}

func (c *ContadorMejorado) Obtener() int { return c.valor }
func (c *ContadorMejorado) Incrementar() { c.valor++ }
func (c *ContadorMejorado) Decrementar() { c.valor-- }
func (c *ContadorMejorado) Reset() { c.valor = 0 }
```

### 🧪 Ejemplo Interactivo

```go
package main

import "fmt"

type Temperatura struct {
    celsius float64
}

// Value receiver - no modifica original
func (t Temperatura) Fahrenheit() float64 {
    return t.celsius*9/5 + 32
}

func (t Temperatura) Kelvin() float64 {
    return t.celsius + 273.15
}

// Pointer receiver - modifica original
func (t *Temperatura) SetCelsius(c float64) {
    t.celsius = c
}

func (t *Temperatura) SetFahrenheit(f float64) {
    t.celsius = (f - 32) * 5 / 9
}

func main() {
    temp := Temperatura{celsius: 25.0}
    
    fmt.Printf("Temperatura: %.1f°C\n", temp.celsius)
    fmt.Printf("En Fahrenheit: %.1f°F\n", temp.Fahrenheit())
    fmt.Printf("En Kelvin: %.1fK\n", temp.Kelvin())
    
    // Modificar temperatura
    temp.SetFahrenheit(86.0)
    fmt.Printf("Después de setear 86°F: %.1f°C\n", temp.celsius)
}
```

---

## 🎯 4. Métodos en Tipos Custom

### 📝 Tipos Custom Básicos

```go
// Definir tipos custom
type ID int
type Email string
type Dinero float64
type Estado string

// Métodos en tipos básicos
func (id ID) EsValido() bool {
    return id > 0
}

func (e Email) EsValido() bool {
    return strings.Contains(string(e), "@") && strings.Contains(string(e), ".")
}

func (d Dinero) String() string {
    return fmt.Sprintf("$%.2f", float64(d))
}

func (d Dinero) SumarImpuesto(porcentaje float64) Dinero {
    return d * Dinero(1+porcentaje/100)
}

func (est Estado) EsActivo() bool {
    return est == "activo" || est == "pendiente"
}
```

### 🏗️ Tipos Slice Custom

```go
// Tipo slice custom
type Numeros []int

// Métodos para slice
func (nums Numeros) Suma() int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func (nums Numeros) Promedio() float64 {
    if len(nums) == 0 {
        return 0
    }
    return float64(nums.Suma()) / float64(len(nums))
}

func (nums Numeros) Max() int {
    if len(nums) == 0 {
        return 0
    }
    max := nums[0]
    for _, num := range nums[1:] {
        if num > max {
            max = num
        }
    }
    return max
}

func (nums *Numeros) Agregar(n int) {
    *nums = append(*nums, n)
}

func (nums *Numeros) Filtrar(fn func(int) bool) {
    resultado := make(Numeros, 0)
    for _, num := range *nums {
        if fn(num) {
            resultado = append(resultado, num)
        }
    }
    *nums = resultado
}

// Uso
func main() {
    nums := Numeros{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Printf("Números: %v\n", nums)
    fmt.Printf("Suma: %d\n", nums.Suma())
    fmt.Printf("Promedio: %.2f\n", nums.Promedio())
    fmt.Printf("Máximo: %d\n", nums.Max())
    
    // Filtrar números pares
    nums.Filtrar(func(n int) bool { return n%2 == 0 })
    fmt.Printf("Solo pares: %v\n", nums)
}
```

### 🗺️ Tipos Map Custom

```go
// Map custom para cache
type Cache map[string]interface{}

func (c Cache) Set(key string, value interface{}) {
    c[key] = value
}

func (c Cache) Get(key string) (interface{}, bool) {
    value, exists := c[key]
    return value, exists
}

func (c Cache) Delete(key string) {
    delete(c, key)
}

func (c Cache) Keys() []string {
    keys := make([]string, 0, len(c))
    for key := range c {
        keys = append(keys, key)
    }
    return keys
}

func (c Cache) Clear() {
    for key := range c {
        delete(c, key)
    }
}

func (c Cache) Size() int {
    return len(c)
}

func (c Cache) Has(key string) bool {
    _, exists := c[key]
    return exists
}

// Uso
func main() {
    cache := make(Cache)
    
    cache.Set("user:1", "Alice")
    cache.Set("user:2", "Bob")
    cache.Set("config:debug", true)
    
    fmt.Printf("Tamaño del cache: %d\n", cache.Size())
    fmt.Printf("Claves: %v\n", cache.Keys())
    
    if value, exists := cache.Get("user:1"); exists {
        fmt.Printf("Encontrado: %v\n", value)
    }
}
```

---

## 🔗 5. Method Chaining

### 🌊 Patrón Fluent Interface

Method chaining permite escribir código más legible y expresivo:

```go
type StringBuilder struct {
    buffer []string
}

// Cada método retorna *StringBuilder para permitir chaining
func (sb *StringBuilder) Add(s string) *StringBuilder {
    sb.buffer = append(sb.buffer, s)
    return sb
}

func (sb *StringBuilder) AddLine(s string) *StringBuilder {
    sb.buffer = append(sb.buffer, s, "\n")
    return sb
}

func (sb *StringBuilder) AddSpace() *StringBuilder {
    sb.buffer = append(sb.buffer, " ")
    return sb
}

func (sb *StringBuilder) AddTab() *StringBuilder {
    sb.buffer = append(sb.buffer, "\t")
    return sb
}

func (sb *StringBuilder) Clear() *StringBuilder {
    sb.buffer = sb.buffer[:0]
    return sb
}

func (sb *StringBuilder) String() string {
    return strings.Join(sb.buffer, "")
}

func (sb *StringBuilder) Length() int {
    return len(sb.String())
}

// Uso con method chaining
func main() {
    sb := &StringBuilder{}
    
    result := sb.
        Add("Hola").
        AddSpace().
        Add("mundo").
        AddLine("!").
        Add("Este es Go").
        AddSpace().
        Add("con method chaining").
        String()
    
    fmt.Println(result)
    
    // También se puede usar sin chaining
    sb.Clear()
    sb.Add("Sin")
    sb.AddSpace()
    sb.Add("chaining")
    fmt.Println(sb.String())
}
```

### 🏗️ Builder Pattern Avanzado

```go
type HTTPRequest struct {
    url         string
    method      string
    headers     map[string]string
    body        []byte
    timeout     time.Duration
    retries     int
}

type HTTPRequestBuilder struct {
    request *HTTPRequest
}

func NewHTTPRequest() *HTTPRequestBuilder {
    return &HTTPRequestBuilder{
        request: &HTTPRequest{
            method:  "GET",
            headers: make(map[string]string),
            timeout: 30 * time.Second,
            retries: 3,
        },
    }
}

func (b *HTTPRequestBuilder) URL(url string) *HTTPRequestBuilder {
    b.request.url = url
    return b
}

func (b *HTTPRequestBuilder) Method(method string) *HTTPRequestBuilder {
    b.request.method = method
    return b
}

func (b *HTTPRequestBuilder) Header(key, value string) *HTTPRequestBuilder {
    b.request.headers[key] = value
    return b
}

func (b *HTTPRequestBuilder) Body(body []byte) *HTTPRequestBuilder {
    b.request.body = body
    return b
}

func (b *HTTPRequestBuilder) JSONBody(data interface{}) *HTTPRequestBuilder {
    jsonData, _ := json.Marshal(data)
    b.request.body = jsonData
    b.Header("Content-Type", "application/json")
    return b
}

func (b *HTTPRequestBuilder) Timeout(timeout time.Duration) *HTTPRequestBuilder {
    b.request.timeout = timeout
    return b
}

func (b *HTTPRequestBuilder) Retries(retries int) *HTTPRequestBuilder {
    b.request.retries = retries
    return b
}

func (b *HTTPRequestBuilder) Build() *HTTPRequest {
    return b.request
}

// Uso
func main() {
    request := NewHTTPRequest().
        URL("https://api.example.com/users").
        Method("POST").
        Header("Authorization", "Bearer token123").
        Header("User-Agent", "MyApp/1.0").
        JSONBody(map[string]string{
            "name":  "Alice",
            "email": "alice@example.com",
        }).
        Timeout(10 * time.Second).
        Retries(5).
        Build()
    
    fmt.Printf("Request: %+v\n", request)
}
```

---

## 🏭 6. Constructores y Factory Methods

### 🔨 Constructor Básico

```go
type Usuario struct {
    ID       int
    Nombre   string
    Email    string
    Activo   bool
    FechaAlta time.Time
}

// Constructor básico
func NewUsuario(nombre, email string) *Usuario {
    return &Usuario{
        ID:        generateID(), // función auxiliar
        Nombre:    nombre,
        Email:     email,
        Activo:    true,
        FechaAlta: time.Now(),
    }
}

// Constructor con validación
func NewUsuarioConValidacion(nombre, email string) (*Usuario, error) {
    if nombre == "" {
        return nil, fmt.Errorf("nombre no puede estar vacío")
    }
    
    if !isValidEmail(email) {
        return nil, fmt.Errorf("email inválido: %s", email)
    }
    
    return &Usuario{
        ID:        generateID(),
        Nombre:    nombre,
        Email:     email,
        Activo:    true,
        FechaAlta: time.Now(),
    }, nil
}

// Constructor con configuración
type UsuarioConfig struct {
    Nombre    string
    Email     string
    Activo    bool
    EsAdmin   bool
    Permisos  []string
}

func NewUsuarioConConfig(config UsuarioConfig) (*Usuario, error) {
    if err := config.Validar(); err != nil {
        return nil, err
    }
    
    usuario := &Usuario{
        ID:        generateID(),
        Nombre:    config.Nombre,
        Email:     config.Email,
        Activo:    config.Activo,
        FechaAlta: time.Now(),
    }
    
    // Lógica adicional basada en configuración
    if config.EsAdmin {
        usuario.AsignarPermisos([]string{"read", "write", "admin"})
    }
    
    return usuario, nil
}
```

### 🏭 Factory Methods

```go
// Factory para diferentes tipos de usuarios
type TipoUsuario string

const (
    UsuarioRegular TipoUsuario = "regular"
    UsuarioAdmin   TipoUsuario = "admin"
    UsuarioGuest   TipoUsuario = "guest"
)

// Factory method principal
func NewUsuarioPorTipo(tipo TipoUsuario, nombre, email string) (*Usuario, error) {
    switch tipo {
    case UsuarioRegular:
        return newUsuarioRegular(nombre, email)
    case UsuarioAdmin:
        return newUsuarioAdmin(nombre, email)
    case UsuarioGuest:
        return newUsuarioGuest(nombre, email)
    default:
        return nil, fmt.Errorf("tipo de usuario desconocido: %s", tipo)
    }
}

// Factory methods específicos
func newUsuarioRegular(nombre, email string) (*Usuario, error) {
    usuario, err := NewUsuarioConValidacion(nombre, email)
    if err != nil {
        return nil, err
    }
    
    usuario.AsignarPermisos([]string{"read", "write"})
    return usuario, nil
}

func newUsuarioAdmin(nombre, email string) (*Usuario, error) {
    usuario, err := NewUsuarioConValidacion(nombre, email)
    if err != nil {
        return nil, err
    }
    
    usuario.AsignarPermisos([]string{"read", "write", "admin", "delete"})
    usuario.MarcarComoAdmin()
    return usuario, nil
}

func newUsuarioGuest(nombre, email string) (*Usuario, error) {
    // Los guests no necesitan email válido
    if nombre == "" {
        nombre = "Guest"
    }
    
    return &Usuario{
        ID:        generateID(),
        Nombre:    nombre,
        Email:     email,
        Activo:    true,
        FechaAlta: time.Now(),
        Permisos:  []string{"read"},
    }, nil
}

// Uso
func main() {
    // Constructor básico
    usuario1 := NewUsuario("Alice", "alice@example.com")
    
    // Constructor con validación
    usuario2, err := NewUsuarioConValidacion("Bob", "bob@example.com")
    if err != nil {
        log.Fatal(err)
    }
    
    // Factory methods
    admin, err := NewUsuarioPorTipo(UsuarioAdmin, "Admin", "admin@example.com")
    if err != nil {
        log.Fatal(err)
    }
    
    guest, _ := NewUsuarioPorTipo(UsuarioGuest, "", "")
    
    fmt.Printf("Usuario regular: %+v\n", usuario1)
    fmt.Printf("Usuario validado: %+v\n", usuario2)
    fmt.Printf("Admin: %+v\n", admin)
    fmt.Printf("Guest: %+v\n", guest)
}
```

---

## ✅ 7. Métodos de Validación

### 🔍 Validaciones Básicas

```go
type Producto struct {
    ID          int
    Nombre      string
    Precio      float64
    Categoria   string
    Stock       int
    Activo      bool
    FechaCreado time.Time
}

// Validaciones individuales
func (p Producto) ValidarNombre() error {
    if strings.TrimSpace(p.Nombre) == "" {
        return fmt.Errorf("nombre del producto es requerido")
    }
    if len(p.Nombre) < 3 {
        return fmt.Errorf("nombre debe tener al menos 3 caracteres")
    }
    if len(p.Nombre) > 100 {
        return fmt.Errorf("nombre no puede exceder 100 caracteres")
    }
    return nil
}

func (p Producto) ValidarPrecio() error {
    if p.Precio <= 0 {
        return fmt.Errorf("precio debe ser mayor a 0")
    }
    if p.Precio > 1000000 {
        return fmt.Errorf("precio no puede exceder $1,000,000")
    }
    return nil
}

func (p Producto) ValidarStock() error {
    if p.Stock < 0 {
        return fmt.Errorf("stock no puede ser negativo")
    }
    return nil
}

func (p Producto) ValidarCategoria() error {
    categoriasValidas := []string{
        "electronica", "ropa", "hogar", "deportes", "libros",
    }
    
    for _, cat := range categoriasValidas {
        if p.Categoria == cat {
            return nil
        }
    }
    
    return fmt.Errorf("categoría '%s' no es válida. Categorías válidas: %v", 
        p.Categoria, categoriasValidas)
}

// Validación completa
func (p Producto) Validar() error {
    validaciones := []func() error{
        p.ValidarNombre,
        p.ValidarPrecio,
        p.ValidarStock,
        p.ValidarCategoria,
    }
    
    for _, validacion := range validaciones {
        if err := validacion(); err != nil {
            return err
        }
    }
    
    return nil
}

// Validación con múltiples errores
func (p Producto) ValidarCompleto() []error {
    var errores []error
    
    if err := p.ValidarNombre(); err != nil {
        errores = append(errores, err)
    }
    
    if err := p.ValidarPrecio(); err != nil {
        errores = append(errores, err)
    }
    
    if err := p.ValidarStock(); err != nil {
        errores = append(errores, err)
    }
    
    if err := p.ValidarCategoria(); err != nil {
        errores = append(errores, err)
    }
    
    return errores
}
```

### 🛡️ Validaciones Avanzadas

```go
// Interfaz para validables
type Validable interface {
    Validar() error
}

// Tipo de error de validación custom
type ErrorValidacion struct {
    Campo   string
    Valor   interface{}
    Mensaje string
}

func (e ErrorValidacion) Error() string {
    return fmt.Sprintf("validación fallida en '%s': %s (valor: %v)", 
        e.Campo, e.Mensaje, e.Valor)
}

// Validador de reglas
type ValidadorReglas struct {
    reglas map[string][]func(interface{}) error
}

func NewValidadorReglas() *ValidadorReglas {
    return &ValidadorReglas{
        reglas: make(map[string][]func(interface{}) error),
    }
}

func (v *ValidadorReglas) AgregarRegla(campo string, regla func(interface{}) error) {
    v.reglas[campo] = append(v.reglas[campo], regla)
}

func (v *ValidadorReglas) Validar(campo string, valor interface{}) error {
    for _, regla := range v.reglas[campo] {
        if err := regla(valor); err != nil {
            return &ErrorValidacion{
                Campo:   campo,
                Valor:   valor,
                Mensaje: err.Error(),
            }
        }
    }
    return nil
}

// Reglas de validación comunes
var (
    Requerido = func(valor interface{}) error {
        if valor == nil {
            return fmt.Errorf("es requerido")
        }
        if str, ok := valor.(string); ok && strings.TrimSpace(str) == "" {
            return fmt.Errorf("no puede estar vacío")
        }
        return nil
    }
    
    LongitudMinima = func(min int) func(interface{}) error {
        return func(valor interface{}) error {
            if str, ok := valor.(string); ok {
                if len(str) < min {
                    return fmt.Errorf("debe tener al menos %d caracteres", min)
                }
            }
            return nil
        }
    }
    
    RangoNumerico = func(min, max float64) func(interface{}) error {
        return func(valor interface{}) error {
            var num float64
            switch v := valor.(type) {
            case int:
                num = float64(v)
            case float64:
                num = v
            default:
                return fmt.Errorf("debe ser un número")
            }
            
            if num < min || num > max {
                return fmt.Errorf("debe estar entre %.2f y %.2f", min, max)
            }
            return nil
        }
    }
)

// Producto con validación avanzada
func (p Producto) ValidarConReglas() error {
    validador := NewValidadorReglas()
    
    // Configurar reglas
    validador.AgregarRegla("nombre", Requerido)
    validador.AgregarRegla("nombre", LongitudMinima(3))
    validador.AgregarRegla("precio", Requerido)
    validador.AgregarRegla("precio", RangoNumerico(0.01, 1000000))
    
    // Validar campos
    campos := map[string]interface{}{
        "nombre": p.Nombre,
        "precio": p.Precio,
    }
    
    for campo, valor := range campos {
        if err := validador.Validar(campo, valor); err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## 📄 8. Serialización y String()

### 🎨 Método String()

```go
type Persona struct {
    Nombre    string
    Edad      int
    Email     string
    Telefono  string
    Direccion string
}

// Implementación básica de String()
func (p Persona) String() string {
    return fmt.Sprintf("Persona{Nombre: %s, Edad: %d, Email: %s}", 
        p.Nombre, p.Edad, p.Email)
}

// String() con formato más legible
func (p Persona) StringDetallado() string {
    var builder strings.Builder
    
    builder.WriteString("=== PERSONA ===\n")
    builder.WriteString(fmt.Sprintf("Nombre: %s\n", p.Nombre))
    builder.WriteString(fmt.Sprintf("Edad: %d años\n", p.Edad))
    builder.WriteString(fmt.Sprintf("Email: %s\n", p.Email))
    
    if p.Telefono != "" {
        builder.WriteString(fmt.Sprintf("Teléfono: %s\n", p.Telefono))
    }
    
    if p.Direccion != "" {
        builder.WriteString(fmt.Sprintf("Dirección: %s\n", p.Direccion))
    }
    
    return builder.String()
}

// String() condicional
func (p Persona) StringPrivado() string {
    // No mostrar información sensible
    return fmt.Sprintf("Persona{Nombre: %s, Edad: %d}", p.Nombre, p.Edad)
}
```

### 📊 Serialización JSON

```go
import (
    "encoding/json"
    "time"
)

type Usuario struct {
    ID           int       `json:"id"`
    Nombre       string    `json:"nombre"`
    Email        string    `json:"email"`
    Password     string    `json:"-"` // No serializar
    FechaNac     time.Time `json:"fecha_nacimiento"`
    Activo       bool      `json:"activo"`
    UltimoLogin  time.Time `json:"ultimo_login,omitempty"`
}

// ToJSON convierte a JSON
func (u Usuario) ToJSON() ([]byte, error) {
    return json.MarshalIndent(u, "", "  ")
}

// ToJSONString convierte a JSON string
func (u Usuario) ToJSONString() string {
    data, err := u.ToJSON()
    if err != nil {
        return fmt.Sprintf(`{"error": "%s"}`, err.Error())
    }
    return string(data)
}

// FromJSON crea Usuario desde JSON
func (u *Usuario) FromJSON(data []byte) error {
    return json.Unmarshal(data, u)
}

// ToMap convierte a map
func (u Usuario) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":                u.ID,
        "nombre":            u.Nombre,
        "email":             u.Email,
        "fecha_nacimiento":  u.FechaNac,
        "activo":            u.Activo,
        "ultimo_login":      u.UltimoLogin,
    }
}

// Clone crea una copia
func (u Usuario) Clone() Usuario {
    return Usuario{
        ID:          u.ID,
        Nombre:      u.Nombre,
        Email:       u.Email,
        Password:    u.Password,
        FechaNac:    u.FechaNac,
        Activo:      u.Activo,
        UltimoLogin: u.UltimoLogin,
    }
}

// Equals compara dos usuarios
func (u Usuario) Equals(otro Usuario) bool {
    return u.ID == otro.ID &&
           u.Nombre == otro.Nombre &&
           u.Email == otro.Email &&
           u.FechaNac.Equal(otro.FechaNac) &&
           u.Activo == otro.Activo
}
```

### 🔄 Serialización Custom

```go
// MarshalJSON custom
func (u Usuario) MarshalJSON() ([]byte, error) {
    // Alias para evitar recursión infinita
    type Alias Usuario
    
    return json.Marshal(&struct {
        *Alias
        EdadCalculada int    `json:"edad_calculada"`
        Tipo          string `json:"tipo"`
    }{
        Alias:         (*Alias)(&u),
        EdadCalculada: u.CalcularEdad(),
        Tipo:          u.ObtenerTipo(),
    })
}

// UnmarshalJSON custom
func (u *Usuario) UnmarshalJSON(data []byte) error {
    type Alias Usuario
    aux := &struct {
        *Alias
        FechaNacStr string `json:"fecha_nacimiento"`
    }{
        Alias: (*Alias)(u),
    }
    
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    
    // Procesar fecha custom
    if aux.FechaNacStr != "" {
        fecha, err := time.Parse("2006-01-02", aux.FechaNacStr)
        if err != nil {
            return err
        }
        u.FechaNac = fecha
    }
    
    return nil
}
```

---

## 📋 9. Patrones y Best Practices

### 🎯 Builder Pattern Completo

```go
type Configuracion struct {
    host     string
    puerto   int
    database string
    usuario  string
    password string
    ssl      bool
    timeout  time.Duration
    retries  int
    debug    bool
}

type ConfigBuilder struct {
    config *Configuracion
}

func NewConfigBuilder() *ConfigBuilder {
    return &ConfigBuilder{
        config: &Configuracion{
            puerto:  5432,
            ssl:     false,
            timeout: 30 * time.Second,
            retries: 3,
            debug:   false,
        },
    }
}

func (b *ConfigBuilder) Host(host string) *ConfigBuilder {
    b.config.host = host
    return b
}

func (b *ConfigBuilder) Puerto(puerto int) *ConfigBuilder {
    b.config.puerto = puerto
    return b
}

func (b *ConfigBuilder) Database(db string) *ConfigBuilder {
    b.config.database = db
    return b
}

func (b *ConfigBuilder) Credenciales(usuario, password string) *ConfigBuilder {
    b.config.usuario = usuario
    b.config.password = password
    return b
}

func (b *ConfigBuilder) ConSSL() *ConfigBuilder {
    b.config.ssl = true
    return b
}

func (b *ConfigBuilder) Timeout(timeout time.Duration) *ConfigBuilder {
    b.config.timeout = timeout
    return b
}

func (b *ConfigBuilder) Retries(retries int) *ConfigBuilder {
    b.config.retries = retries
    return b
}

func (b *ConfigBuilder) Debug(debug bool) *ConfigBuilder {
    b.config.debug = debug
    return b
}

func (b *ConfigBuilder) Build() (*Configuracion, error) {
    if err := b.config.validar(); err != nil {
        return nil, err
    }
    return b.config, nil
}

func (c *Configuracion) validar() error {
    if c.host == "" {
        return fmt.Errorf("host es requerido")
    }
    if c.database == "" {
        return fmt.Errorf("database es requerido")
    }
    if c.puerto <= 0 || c.puerto > 65535 {
        return fmt.Errorf("puerto debe estar entre 1 y 65535")
    }
    return nil
}

// Uso
func main() {
    config, err := NewConfigBuilder().
        Host("localhost").
        Puerto(5432).
        Database("miapp").
        Credenciales("admin", "password123").
        ConSSL().
        Timeout(10 * time.Second).
        Debug(true).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Configuración: %+v\n", config)
}
```

### 🔄 State Pattern

```go
type EstadoPedido interface {
    Procesar(p *Pedido) error
    Cancelar(p *Pedido) error
    String() string
}

type Pedido struct {
    ID       int
    Items    []string
    Total    float64
    estado   EstadoPedido
}

// Estados concretos
type EstadoPendiente struct{}
type EstadoProcesando struct{}
type EstadoEnviado struct{}
type EstadoEntregado struct{}
type EstadoCancelado struct{}

// Implementaciones del estado Pendiente
func (e EstadoPendiente) Procesar(p *Pedido) error {
    p.estado = EstadoProcesando{}
    fmt.Printf("Pedido %d ahora está en procesamiento\n", p.ID)
    return nil
}

func (e EstadoPendiente) Cancelar(p *Pedido) error {
    p.estado = EstadoCancelado{}
    fmt.Printf("Pedido %d cancelado\n", p.ID)
    return nil
}

func (e EstadoPendiente) String() string {
    return "Pendiente"
}

// Implementaciones del estado Procesando
func (e EstadoProcesando) Procesar(p *Pedido) error {
    p.estado = EstadoEnviado{}
    fmt.Printf("Pedido %d enviado\n", p.ID)
    return nil
}

func (e EstadoProcesando) Cancelar(p *Pedido) error {
    return fmt.Errorf("no se puede cancelar un pedido en procesamiento")
}

func (e EstadoProcesando) String() string {
    return "Procesando"
}

// Métodos del pedido
func (p *Pedido) Procesar() error {
    return p.estado.Procesar(p)
}

func (p *Pedido) Cancelar() error {
    return p.estado.Cancelar(p)
}

func (p *Pedido) Estado() string {
    return p.estado.String()
}

func NewPedido(id int, items []string, total float64) *Pedido {
    return &Pedido{
        ID:     id,
        Items:  items,
        Total:  total,
        estado: EstadoPendiente{},
    }
}
```

### 🎭 Observer Pattern

```go
type Observer interface {
    Notificar(evento string, data interface{})
}

type Observable struct {
    observers []Observer
}

func (o *Observable) AgregarObserver(observer Observer) {
    o.observers = append(o.observers, observer)
}

func (o *Observable) RemoverObserver(observer Observer) {
    for i, obs := range o.observers {
        if obs == observer {
            o.observers = append(o.observers[:i], o.observers[i+1:]...)
            break
        }
    }
}

func (o *Observable) NotificarObservers(evento string, data interface{}) {
    for _, observer := range o.observers {
        observer.Notificar(evento, data)
    }
}

// Cuenta bancaria observable
type CuentaBancaria struct {
    Observable
    numero string
    saldo  float64
}

func NewCuentaBancaria(numero string, saldoInicial float64) *CuentaBancaria {
    return &CuentaBancaria{
        numero: numero,
        saldo:  saldoInicial,
    }
}

func (c *CuentaBancaria) Depositar(monto float64) {
    c.saldo += monto
    c.NotificarObservers("deposito", map[string]interface{}{
        "cuenta": c.numero,
        "monto":  monto,
        "saldo":  c.saldo,
    })
}

func (c *CuentaBancaria) Retirar(monto float64) error {
    if monto > c.saldo {
        return fmt.Errorf("fondos insuficientes")
    }
    
    c.saldo -= monto
    c.NotificarObservers("retiro", map[string]interface{}{
        "cuenta": c.numero,
        "monto":  monto,
        "saldo":  c.saldo,
    })
    
    return nil
}

// Observadores concretos
type LoggerObserver struct{}

func (l LoggerObserver) Notificar(evento string, data interface{}) {
    fmt.Printf("[LOG] Evento: %s, Data: %+v\n", evento, data)
}

type EmailObserver struct{}

func (e EmailObserver) Notificar(evento string, data interface{}) {
    if evento == "retiro" {
        fmt.Printf("[EMAIL] Alerta: Retiro realizado\n")
    }
}
```

---

## 🎯 10. Ejercicios Prácticos

### 📝 Lista de Ejercicios

1. **Calculadora Científica** - Métodos para operaciones matemáticas
2. **Sistema de Biblioteca** - Gestión de libros con validaciones
3. **API Client Builder** - Builder pattern para HTTP requests
4. **State Machine** - Máquina de estados para semáforo
5. **Observable Store** - Store con observers para estado
6. **Validador de Formularios** - Sistema de validación flexible
7. **Cache LRU** - Cache con política Least Recently Used
8. **Árbol Binario** - Estructura de datos con métodos de búsqueda

### 🧮 Ejercicio 1: Calculadora Científica

```go
type CalculadoraCientifica struct {
    historial []string
    precision int
}

// TODO: Implementar estos métodos
func NewCalculadora() *CalculadoraCientifica {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Sumar(a, b float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Restar(a, b float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Multiplicar(a, b float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Dividir(a, b float64) (float64, error) {
    // Tu código aquí - manejar división por cero
}

func (c *CalculadoraCientifica) Potencia(base, exponente float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) RaizCuadrada(x float64) (float64, error) {
    // Tu código aquí - manejar números negativos
}

func (c *CalculadoraCientifica) Sin(x float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Cos(x float64) float64 {
    // Tu código aquí
}

func (c *CalculadoraCientifica) Logaritmo(x float64) (float64, error) {
    // Tu código aquí
}

func (c *CalculadoraCientifica) ObtenerHistorial() []string {
    // Tu código aquí
}

func (c *CalculadoraCientifica) LimpiarHistorial() {
    // Tu código aquí
}

func (c *CalculadoraCientifica) String() string {
    // Tu código aquí
}
```

---

## ✅ Checklist de Dominio

Antes de continuar a la siguiente lección, asegúrate de poder:

- [ ] Definir métodos con value y pointer receivers
- [ ] Explicar cuándo usar cada tipo de receiver
- [ ] Implementar method chaining
- [ ] Crear constructores y factory methods
- [ ] Escribir métodos de validación efectivos
- [ ] Implementar el método String() 
- [ ] Serializar structs a JSON con métodos custom
- [ ] Aplicar patrones de diseño con métodos
- [ ] Crear APIs fluent y expresivas
- [ ] Manejar errores en métodos apropiadamente

---

## 🔗 Navegación

⬅️ **Anterior**: [Lección 11: Structs](../11-structs/README.md)  
➡️ **Siguiente**: [Lección 13: Interfaces Básicas](../13-interfaces-basicas/README.md)  
🏠 **Inicio**: [Fundamentos de Go](../README.md)  
📚 **Curso**: [Go Deep - Domina Go](../../README.md)

---

## 📞 Soporte

¿Tienes dudas sobre métodos? 

- 💬 **Discusión**: [GitHub Discussions](../../discussions)
- 🐛 **Problemas**: [GitHub Issues](../../issues)
- 📧 **Email**: [contacto@go-deep.dev](mailto:contacto@go-deep.dev)

---

¡Los métodos son la puerta de entrada a la programación orientada a objetos en Go! 🎭 Practica mucho con diferentes patterns y verás cómo tu código se vuelve más expresivo y mantenible.

**¡Continuemos construyendo tus habilidades en Go!** 🚀
