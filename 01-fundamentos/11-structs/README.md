# 📦 Lección 11: Structs en Go

## 🎯 Objetivos de Aprendizaje

Al completar esta lección, podrás:

- ✅ Entender qué son los structs y por qué son fundamentales
- ✅ Declarar y utilizar structs de diferentes formas
- ✅ Trabajar con campos, tags y metadata
- ✅ Implementar composition y embedding
- ✅ Usar constructores e inicialización
- ✅ Manejar structs anónimos y anidados
- ✅ Aplicar técnicas avanzadas con structs
- ✅ Desarrollar un sistema completo usando structs

---

## 📚 Tabla de Contenidos

1. [¿Qué son los Structs?](#-qué-son-los-structs)
2. [Declaración y Sintaxis](#-declaración-y-sintaxis)
3. [Inicialización de Structs](#-inicialización-de-structs)
4. [Campos y Métodos](#-campos-y-métodos)
5. [Embedding y Composition](#-embedding-y-composition)
6. [Tags y Metadata](#-tags-y-metadata)
7. [Structs Anónimos](#-structs-anónimos)
8. [Técnicas Avanzadas](#-técnicas-avanzadas)
9. [Patrones y Best Practices](#-patrones-y-best-practices)
10. [Errores Comunes](#-errores-comunes)

---

## 🔍 ¿Qué son los Structs?

### 🧠 Analogía: Los Structs como Planos Arquitectónicos

Imagina que los **structs** son como **planos arquitectónicos** para construir casas:

```
🏗️ PLANO (Struct Definition)
┌─────────────────────────────┐
│ 🏠 Casa                     │
│ ├── 🚪 Puertas: 3          │
│ ├── 🪟 Ventanas: 8         │
│ ├── 🛏️ Habitaciones: 4     │
│ ├── 🚿 Baños: 2            │
│ └── 📐 Metros²: 120        │
└─────────────────────────────┘
        ⬇️
🏘️ CONSTRUCCIONES (Instances)
🏠casa1  🏠casa2  🏠casa3
```

### 📖 Definición Técnica

Un **struct** es un tipo de dato personalizado que agrupa diferentes campos (variables) bajo un mismo nombre. Es la forma principal de crear tipos complejos en Go.

```go
// Definición de un struct
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}

// Creación de una instancia
persona := Persona{
    Nombre: "Ana García",
    Edad:   25,
    Email:  "ana@ejemplo.com",
}
```

---

## 🛠️ Declaración y Sintaxis

### 1️⃣ Sintaxis Básica

```go
// Formato general
type NombreStruct struct {
    campo1 tipo1
    campo2 tipo2
    campo3 tipo3
}
```

### 2️⃣ Ejemplo Completo

```go
package main

import (
    "fmt"
    "time"
)

// Struct básico para representar un libro
type Libro struct {
    ID          int
    Titulo      string
    Autor       string
    Paginas     int
    Precio      float64
    Disponible  bool
    FechaPublic time.Time
}

func main() {
    // Crear instancia usando campos nombrados
    libro1 := Libro{
        ID:          1,
        Titulo:      "El Principito",
        Autor:       "Antoine de Saint-Exupéry",
        Paginas:     96,
        Precio:      15.99,
        Disponible:  true,
        FechaPublic: time.Date(1943, 4, 6, 0, 0, 0, 0, time.UTC),
    }

    fmt.Printf("📖 %s por %s\n", libro1.Titulo, libro1.Autor)
    fmt.Printf("💰 Precio: $%.2f\n", libro1.Precio)
}
```

### 3️⃣ Campos con Múltiples Tipos

```go
type Configuracion struct {
    // Tipos básicos
    AppName     string
    Port        int
    Debug       bool
    Version     float32
    
    // Slices y arrays
    Servidores  []string
    Puertos     [3]int
    
    // Maps
    Variables   map[string]string
    
    // Otros structs
    Database    DatabaseConfig
    
    // Punteros
    Logger      *LoggerConfig
    
    // Interfaces
    Handler     interface{}
}

type DatabaseConfig struct {
    Host     string
    Port     int
    Database string
    SSL      bool
}

type LoggerConfig struct {
    Level  string
    Output string
}
```

---

## 🚀 Inicialización de Structs

### 1️⃣ Formas de Inicialización

```go
type Usuario struct {
    ID       int
    Nombre   string
    Email    string
    Activo   bool
}

func ejemplosInicializacion() {
    // 1. Zero value (valores por defecto)
    var u1 Usuario
    fmt.Printf("Zero value: %+v\n", u1)
    // Output: {ID:0 Nombre: Email: Activo:false}
    
    // 2. Literal con campos nombrados (RECOMENDADO)
    u2 := Usuario{
        ID:     1,
        Nombre: "Ana",
        Email:  "ana@test.com",
        Activo: true,
    }
    
    // 3. Literal sin nombres (NO RECOMENDADO)
    u3 := Usuario{2, "Carlos", "carlos@test.com", true}
    
    // 4. Inicialización parcial
    u4 := Usuario{
        ID:     3,
        Nombre: "María",
        // Email y Activo tendrán zero values
    }
    
    // 5. Usando new()
    u5 := new(Usuario)
    u5.ID = 4
    u5.Nombre = "Pedro"
    
    // 6. Con función constructora
    u6 := NewUsuario("Luis", "luis@test.com")
    
    fmt.Printf("u2: %+v\n", u2)
    fmt.Printf("u6: %+v\n", u6)
}

// Función constructora (patrón común)
func NewUsuario(nombre, email string) Usuario {
    return Usuario{
        ID:     generarID(),
        Nombre: nombre,
        Email:  email,
        Activo: true,
    }
}

func generarID() int {
    // Lógica para generar ID único
    return int(time.Now().Unix())
}
```

### 2️⃣ Constructores Avanzados

```go
type Servidor struct {
    Host            string
    Puerto          int
    SSL             bool
    TimeoutSegundos int
    MaxConexiones   int
    Configuracion   map[string]interface{}
}

// Constructor básico
func NewServidor(host string, puerto int) *Servidor {
    return &Servidor{
        Host:            host,
        Puerto:          puerto,
        SSL:             false,
        TimeoutSegundos: 30,
        MaxConexiones:   100,
        Configuracion:   make(map[string]interface{}),
    }
}

// Constructor con opciones
func NewServidorConOpciones(host string, puerto int, opciones ...func(*Servidor)) *Servidor {
    servidor := NewServidor(host, puerto)
    
    // Aplicar opciones
    for _, opcion := range opciones {
        opcion(servidor)
    }
    
    return servidor
}

// Funciones de opción
func ConSSL() func(*Servidor) {
    return func(s *Servidor) {
        s.SSL = true
    }
}

func ConTimeout(segundos int) func(*Servidor) {
    return func(s *Servidor) {
        s.TimeoutSegundos = segundos
    }
}

func ConMaxConexiones(max int) func(*Servidor) {
    return func(s *Servidor) {
        s.MaxConexiones = max
    }
}

// Uso del constructor con opciones
func ejemploConstructorOpciones() {
    servidor := NewServidorConOpciones(
        "localhost", 8080,
        ConSSL(),
        ConTimeout(60),
        ConMaxConexiones(500),
    )
    
    fmt.Printf("Servidor: %+v\n", servidor)
}
```

---

## 🔧 Campos y Métodos

### 1️⃣ Acceso a Campos

```go
type Producto struct {
    ID          int
    Nombre      string
    Precio      float64
    Categoria   string
    EnStock     bool
    Descuento   float64
}

func ejemploAccesoCampos() {
    producto := Producto{
        ID:        101,
        Nombre:    "Laptop Gaming",
        Precio:    1299.99,
        Categoria: "Electrónicos",
        EnStock:   true,
        Descuento: 0.10,
    }
    
    // Leer campos
    fmt.Printf("Producto: %s\n", producto.Nombre)
    fmt.Printf("Precio original: $%.2f\n", producto.Precio)
    
    // Modificar campos
    producto.Precio = 1199.99
    producto.Descuento = 0.15
    
    // Calcular precio final
    precioFinal := producto.Precio * (1 - producto.Descuento)
    fmt.Printf("Precio final: $%.2f\n", precioFinal)
    
    // Verificar disponibilidad
    if producto.EnStock {
        fmt.Println("✅ Disponible en stock")
    } else {
        fmt.Println("❌ Agotado")
    }
}
```

### 2️⃣ Métodos en Structs

```go
type Rectangulo struct {
    Ancho  float64
    Alto   float64
}

// Método con receptor por valor
func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Alto
}

// Método con receptor por puntero
func (r *Rectangulo) Escalar(factor float64) {
    r.Ancho *= factor
    r.Alto *= factor
}

// Método que devuelve múltiples valores
func (r Rectangulo) Dimensiones() (float64, float64) {
    return r.Ancho, r.Alto
}

// Método que modifica y devuelve nuevo struct
func (r Rectangulo) ConMargen(margen float64) Rectangulo {
    return Rectangulo{
        Ancho: r.Ancho + 2*margen,
        Alto:  r.Alto + 2*margen,
    }
}

func ejemploMetodos() {
    rect := Rectangulo{Ancho: 10, Alto: 5}
    
    fmt.Printf("Área original: %.2f\n", rect.Area())
    
    // Escalar (modifica el original)
    rect.Escalar(2)
    fmt.Printf("Después de escalar x2: %.2f x %.2f\n", rect.Dimensiones())
    
    // Crear versión con margen (no modifica original)
    rectConMargen := rect.ConMargen(1)
    fmt.Printf("Con margen: %.2f x %.2f\n", rectConMargen.Dimensiones())
    fmt.Printf("Original sigue igual: %.2f x %.2f\n", rect.Dimensiones())
}
```

---

## 🏗️ Embedding y Composition

### 1️⃣ Embedding Básico

```go
// Struct base
type Direccion struct {
    Calle     string
    Ciudad    string
    CodigoP   string
    Pais      string
}

func (d Direccion) String() string {
    return fmt.Sprintf("%s, %s %s, %s", d.Calle, d.Ciudad, d.CodigoP, d.Pais)
}

// Struct que embebe Direccion
type Persona struct {
    Nombre   string
    Edad     int
    Email    string
    Direccion // Embedding anónimo
}

// Struct con embedding nombrado
type Empresa struct {
    Nombre     string
    Telefono   string
    Direccion  Direccion // Embedding nombrado
}

func ejemploEmbedding() {
    persona := Persona{
        Nombre: "Ana García",
        Edad:   30,
        Email:  "ana@ejemplo.com",
        Direccion: Direccion{
            Calle:   "Av. Principal 123",
            Ciudad:  "Madrid",
            CodigoP: "28001",
            Pais:    "España",
        },
    }
    
    // Acceso directo a campos embebidos (promoción de campos)
    fmt.Printf("Vive en: %s\n", persona.Ciudad)
    fmt.Printf("Dirección completa: %s\n", persona.Direccion.String())
    
    empresa := Empresa{
        Nombre:   "Tech Corp",
        Telefono: "+34-91-123-4567",
        Direccion: Direccion{
            Calle:   "Gran Vía 45",
            Ciudad:  "Madrid",
            CodigoP: "28013",
            Pais:    "España",
        },
    }
    
    // Con embedding nombrado necesitas el nombre del campo
    fmt.Printf("Empresa en: %s\n", empresa.Direccion.Ciudad)
}
```

### 2️⃣ Múltiple Embedding

```go
type Identificable struct {
    ID string
}

func (i Identificable) GetID() string {
    return i.ID
}

type Timestampable struct {
    CreadoEn      time.Time
    ActualizadoEn time.Time
}

func (t *Timestampable) ActualizarTimestamp() {
    t.ActualizadoEn = time.Now()
}

type Auditable struct {
    CreadoPor      string
    ActualizadoPor string
}

// Struct que combina múltiples embeddings
type Documento struct {
    Identificable  // ID y GetID()
    Timestampable  // Campos de tiempo y ActualizarTimestamp()
    Auditable      // Campos de auditoría
    
    Titulo    string
    Contenido string
    Version   int
}

func (d *Documento) Actualizar(contenido, usuario string) {
    d.Contenido = contenido
    d.Version++
    d.ActualizadoPor = usuario
    d.ActualizarTimestamp()
}

func ejemploMultipleEmbedding() {
    doc := Documento{
        Identificable: Identificable{ID: "DOC-001"},
        Timestampable: Timestampable{
            CreadoEn:      time.Now(),
            ActualizadoEn: time.Now(),
        },
        Auditable: Auditable{
            CreadoPor:      "admin",
            ActualizadoPor: "admin",
        },
        Titulo:    "Manual de Usuario",
        Contenido: "Contenido inicial...",
        Version:   1,
    }
    
    fmt.Printf("Documento ID: %s\n", doc.GetID())
    
    // Actualizar documento
    doc.Actualizar("Contenido actualizado...", "editor")
    
    fmt.Printf("Versión: %d\n", doc.Version)
    fmt.Printf("Actualizado por: %s\n", doc.ActualizadoPor)
    fmt.Printf("Última actualización: %s\n", doc.ActualizadoEn.Format("2006-01-02 15:04:05"))
}
```

---

## 🏷️ Tags y Metadata

### 1️⃣ Struct Tags Básicos

```go
import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "reflect"
)

type Usuario struct {
    ID       int    `json:"id" xml:"id,attr" db:"user_id" validate:"required"`
    Nombre   string `json:"name" xml:"name" db:"full_name" validate:"required,min=2"`
    Email    string `json:"email" xml:"email" db:"email_address" validate:"required,email"`
    Edad     int    `json:"age,omitempty" xml:"age,omitempty" db:"age" validate:"min=0,max=120"`
    Password string `json:"-" xml:"-" db:"password_hash" validate:"required,min=8"`
    Activo   bool   `json:"active" xml:"active" db:"is_active"`
}

func ejemploTags() {
    usuario := Usuario{
        ID:       1,
        Nombre:   "Ana García",
        Email:    "ana@ejemplo.com",
        Edad:     25,
        Password: "secreto123",
        Activo:   true,
    }
    
    // JSON marshaling
    jsonData, _ := json.MarshalIndent(usuario, "", "  ")
    fmt.Println("JSON:")
    fmt.Println(string(jsonData))
    
    // XML marshaling
    xmlData, _ := xml.MarshalIndent(usuario, "", "  ")
    fmt.Println("\nXML:")
    fmt.Println(string(xmlData))
    
    // Reflection para leer tags
    fmt.Println("\nStruct Tags:")
    inspeccionar(usuario)
}

func inspeccionar(v interface{}) {
    t := reflect.TypeOf(v)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        
        fmt.Printf("Campo: %s\n", field.Name)
        fmt.Printf("  JSON tag: %s\n", field.Tag.Get("json"))
        fmt.Printf("  DB tag: %s\n", field.Tag.Get("db"))
        fmt.Printf("  Validate tag: %s\n", field.Tag.Get("validate"))
        fmt.Println()
    }
}
```

### 2️⃣ Tags Personalizados

```go
type ConfiguracionApp struct {
    Puerto      int    `env:"PORT" default:"8080" description:"Puerto del servidor"`
    Host        string `env:"HOST" default:"localhost" description:"Host del servidor"`
    Debug       bool   `env:"DEBUG" default:"false" description:"Modo debug"`
    DatabaseURL string `env:"DATABASE_URL" required:"true" description:"URL de la base de datos"`
    SecretKey   string `env:"SECRET_KEY" required:"true" sensitive:"true" description:"Clave secreta"`
}

// Función que lee configuración usando tags
func cargarConfiguracion() ConfiguracionApp {
    config := ConfiguracionApp{}
    t := reflect.TypeOf(config)
    v := reflect.ValueOf(&config).Elem()
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)
        
        // Leer variable de entorno
        envName := field.Tag.Get("env")
        envValue := os.Getenv(envName)
        
        if envValue == "" {
            // Usar valor por defecto si no hay variable de entorno
            defaultValue := field.Tag.Get("default")
            envValue = defaultValue
        }
        
        // Verificar si es requerido
        if field.Tag.Get("required") == "true" && envValue == "" {
            fmt.Printf("Error: %s es requerido\n", envName)
            continue
        }
        
        // Asignar valor según el tipo
        switch fieldValue.Kind() {
        case reflect.String:
            fieldValue.SetString(envValue)
        case reflect.Int:
            if intVal, err := strconv.Atoi(envValue); err == nil {
                fieldValue.SetInt(int64(intVal))
            }
        case reflect.Bool:
            if boolVal, err := strconv.ParseBool(envValue); err == nil {
                fieldValue.SetBool(boolVal)
            }
        }
        
        // Mostrar información (sin valores sensibles)
        description := field.Tag.Get("description")
        sensitive := field.Tag.Get("sensitive") == "true"
        
        displayValue := envValue
        if sensitive && envValue != "" {
            displayValue = "*****"
        }
        
        fmt.Printf("%s (%s): %s - %s\n", envName, field.Name, displayValue, description)
    }
    
    return config
}
```

---

## 👤 Structs Anónimos

### 1️⃣ Declaración y Uso

```go
func ejemploStructsAnonimos() {
    // Struct anónimo simple
    usuario := struct {
        Nombre string
        Edad   int
    }{
        Nombre: "Carlos",
        Edad:   30,
    }
    
    fmt.Printf("Usuario anónimo: %+v\n", usuario)
    
    // Slice de structs anónimos
    empleados := []struct {
        ID       int
        Nombre   string
        Salario  float64
        Activo   bool
    }{
        {1, "Ana", 50000, true},
        {2, "Luis", 55000, true},
        {3, "María", 48000, false},
    }
    
    fmt.Println("\nEmpleados:")
    for _, emp := range empleados {
        status := "Inactivo"
        if emp.Activo {
            status = "Activo"
        }
        fmt.Printf("ID: %d, %s - $%.2f (%s)\n", 
            emp.ID, emp.Nombre, emp.Salario, status)
    }
    
    // Map con structs anónimos
    configuraciones := map[string]struct {
        Host    string
        Puerto  int
        SSL     bool
    }{
        "desarrollo": {"localhost", 8080, false},
        "pruebas":    {"test.ejemplo.com", 8443, true},
        "produccion": {"prod.ejemplo.com", 443, true},
    }
    
    fmt.Println("\nConfiguraciones:")
    for env, config := range configuraciones {
        protocol := "http"
        if config.SSL {
            protocol = "https"
        }
        fmt.Printf("%s: %s://%s:%d\n", 
            env, protocol, config.Host, config.Puerto)
    }
}
```

### 2️⃣ Casos de Uso Avanzados

```go
// Función que retorna struct anónimo
func obtenerEstadisticas() struct {
    TotalUsuarios    int
    UsuariosActivos  int
    PromedioEdad     float64
    DistribucionPais map[string]int
} {
    // Simulación de cálculos
    return struct {
        TotalUsuarios    int
        UsuariosActivos  int
        PromedioEdad     float64
        DistribucionPais map[string]int
    }{
        TotalUsuarios:   1000,
        UsuariosActivos: 750,
        PromedioEdad:    28.5,
        DistribucionPais: map[string]int{
            "España": 400,
            "México": 300,
            "Argentina": 200,
            "Colombia": 100,
        },
    }
}

// API response con struct anónimo
func procesarRespuestaAPI() {
    respuesta := struct {
        Status int    `json:"status"`
        Mensaje string `json:"message"`
        Datos   struct {
            Usuarios []struct {
                ID     int    `json:"id"`
                Nombre string `json:"name"`
                Email  string `json:"email"`
            } `json:"users"`
            Total int `json:"total"`
        } `json:"data"`
    }{}
    
    // Simular parsing de JSON
    jsonStr := `{
        "status": 200,
        "message": "success",
        "data": {
            "users": [
                {"id": 1, "name": "Ana", "email": "ana@test.com"},
                {"id": 2, "name": "Luis", "email": "luis@test.com"}
            ],
            "total": 2
        }
    }`
    
    json.Unmarshal([]byte(jsonStr), &respuesta)
    
    fmt.Printf("Status: %d - %s\n", respuesta.Status, respuesta.Mensaje)
    fmt.Printf("Total usuarios: %d\n", respuesta.Datos.Total)
    
    for _, usuario := range respuesta.Datos.Usuarios {
        fmt.Printf("- %s (%s)\n", usuario.Nombre, usuario.Email)
    }
}
```

---

## 🚀 Técnicas Avanzadas

### 1️⃣ Struct Factory Pattern

```go
type ConexionDB struct {
    Host         string
    Puerto       int
    Database     string
    Usuario      string
    Password     string
    MaxConex     int
    Timeout      time.Duration
    SSL          bool
    conexion     interface{} // conexión real
}

type DBFactory struct {
    configuraciones map[string]ConexionDB
}

func NewDBFactory() *DBFactory {
    return &DBFactory{
        configuraciones: map[string]ConexionDB{
            "mysql": {
                Host:     "localhost",
                Puerto:   3306,
                MaxConex: 10,
                Timeout:  30 * time.Second,
                SSL:      false,
            },
            "postgres": {
                Host:     "localhost",
                Puerto:   5432,
                MaxConex: 20,
                Timeout:  45 * time.Second,
                SSL:      true,
            },
            "redis": {
                Host:     "localhost",
                Puerto:   6379,
                MaxConex: 50,
                Timeout:  10 * time.Second,
                SSL:      false,
            },
        },
    }
}

func (f *DBFactory) CrearConexion(tipo, database, usuario, password string) (*ConexionDB, error) {
    config, existe := f.configuraciones[tipo]
    if !existe {
        return nil, fmt.Errorf("tipo de base de datos no soportado: %s", tipo)
    }
    
    // Personalizar configuración
    config.Database = database
    config.Usuario = usuario
    config.Password = password
    
    // Simular conexión
    fmt.Printf("🔌 Conectando a %s://%s:%d/%s\n", 
        tipo, config.Host, config.Puerto, config.Database)
    
    return &config, nil
}
```

### 2️⃣ Struct Validation

```go
import (
    "errors"
    "regexp"
    "strings"
)

type Validable interface {
    Validar() error
}

type UsuarioCompleto struct {
    ID          int    `validate:"required,min=1"`
    Nombre      string `validate:"required,min=2,max=50"`
    Apellido    string `validate:"required,min=2,max=50"`
    Email       string `validate:"required,email"`
    Telefono    string `validate:"phone"`
    Edad        int    `validate:"required,min=18,max=100"`
    Website     string `validate:"url,optional"`
    Contraseña  string `validate:"required,min=8,complexity"`
}

func (u UsuarioCompleto) Validar() error {
    var errores []string
    
    // Validar ID
    if u.ID <= 0 {
        errores = append(errores, "ID debe ser mayor que 0")
    }
    
    // Validar nombre y apellido
    if len(strings.TrimSpace(u.Nombre)) < 2 {
        errores = append(errores, "Nombre debe tener al menos 2 caracteres")
    }
    if len(strings.TrimSpace(u.Apellido)) < 2 {
        errores = append(errores, "Apellido debe tener al menos 2 caracteres")
    }
    
    // Validar email
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(u.Email) {
        errores = append(errores, "Email no tiene formato válido")
    }
    
    // Validar teléfono
    phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
    if u.Telefono != "" && !phoneRegex.MatchString(u.Telefono) {
        errores = append(errores, "Teléfono no tiene formato válido")
    }
    
    // Validar edad
    if u.Edad < 18 || u.Edad > 100 {
        errores = append(errores, "Edad debe estar entre 18 y 100 años")
    }
    
    // Validar website
    if u.Website != "" {
        urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
        if !urlRegex.MatchString(u.Website) {
            errores = append(errores, "Website debe ser una URL válida")
        }
    }
    
    // Validar complejidad de contraseña
    if len(u.Contraseña) < 8 {
        errores = append(errores, "Contraseña debe tener al menos 8 caracteres")
    }
    if !regexp.MustCompile(`[A-Z]`).MatchString(u.Contraseña) {
        errores = append(errores, "Contraseña debe contener al menos una mayúscula")
    }
    if !regexp.MustCompile(`[a-z]`).MatchString(u.Contraseña) {
        errores = append(errores, "Contraseña debe contener al menos una minúscula")
    }
    if !regexp.MustCompile(`[0-9]`).MatchString(u.Contraseña) {
        errores = append(errores, "Contraseña debe contener al menos un número")
    }
    
    if len(errores) > 0 {
        return errors.New(strings.Join(errores, "; "))
    }
    
    return nil
}

func ejemploValidacion() {
    usuario := UsuarioCompleto{
        ID:         1,
        Nombre:     "Ana",
        Apellido:   "García",
        Email:      "ana@ejemplo.com",
        Telefono:   "+34612345678",
        Edad:       25,
        Website:    "https://ana-garcia.com",
        Contraseña: "MiPassword123",
    }
    
    if err := usuario.Validar(); err != nil {
        fmt.Printf("❌ Errores de validación: %v\n", err)
    } else {
        fmt.Println("✅ Usuario válido")
    }
}
```

---

## 📋 Patrones y Best Practices

### 1️⃣ Builder Pattern

```go
type ServidorHTTP struct {
    host            string
    puerto          int
    ssl             bool
    certificadoSSL  string
    claveSSL        string
    timeout         time.Duration
    maxConexiones   int
    middleware      []func(http.HandlerFunc) http.HandlerFunc
    rutas           map[string]http.HandlerFunc
    logLevel        string
    cors            bool
}

type ServidorBuilder struct {
    servidor *ServidorHTTP
}

func NewServidorBuilder() *ServidorBuilder {
    return &ServidorBuilder{
        servidor: &ServidorHTTP{
            host:          "localhost",
            puerto:        8080,
            ssl:           false,
            timeout:       30 * time.Second,
            maxConexiones: 100,
            middleware:    []func(http.HandlerFunc) http.HandlerFunc{},
            rutas:         make(map[string]http.HandlerFunc),
            logLevel:      "info",
            cors:          false,
        },
    }
}

func (b *ServidorBuilder) Host(host string) *ServidorBuilder {
    b.servidor.host = host
    return b
}

func (b *ServidorBuilder) Puerto(puerto int) *ServidorBuilder {
    b.servidor.puerto = puerto
    return b
}

func (b *ServidorBuilder) ConSSL(certificado, clave string) *ServidorBuilder {
    b.servidor.ssl = true
    b.servidor.certificadoSSL = certificado
    b.servidor.claveSSL = clave
    return b
}

func (b *ServidorBuilder) Timeout(timeout time.Duration) *ServidorBuilder {
    b.servidor.timeout = timeout
    return b
}

func (b *ServidorBuilder) MaxConexiones(max int) *ServidorBuilder {
    b.servidor.maxConexiones = max
    return b
}

func (b *ServidorBuilder) AgregarMiddleware(mw func(http.HandlerFunc) http.HandlerFunc) *ServidorBuilder {
    b.servidor.middleware = append(b.servidor.middleware, mw)
    return b
}

func (b *ServidorBuilder) AgregarRuta(patron string, handler http.HandlerFunc) *ServidorBuilder {
    b.servidor.rutas[patron] = handler
    return b
}

func (b *ServidorBuilder) LogLevel(level string) *ServidorBuilder {
    b.servidor.logLevel = level
    return b
}

func (b *ServidorBuilder) ConCORS() *ServidorBuilder {
    b.servidor.cors = true
    return b
}

func (b *ServidorBuilder) Build() *ServidorHTTP {
    return b.servidor
}

func (s *ServidorHTTP) Iniciar() error {
    fmt.Printf("🚀 Iniciando servidor en %s:%d\n", s.host, s.puerto)
    fmt.Printf("📊 SSL: %v, Timeout: %v, Max Conexiones: %d\n", 
        s.ssl, s.timeout, s.maxConexiones)
    fmt.Printf("🛣️ Rutas registradas: %d\n", len(s.rutas))
    fmt.Printf("🔧 Middleware: %d\n", len(s.middleware))
    
    // Simular inicio del servidor
    return nil
}

func ejemploBuilder() {
    servidor := NewServidorBuilder().
        Host("0.0.0.0").
        Puerto(8443).
        ConSSL("/path/to/cert.pem", "/path/to/key.pem").
        Timeout(60 * time.Second).
        MaxConexiones(500).
        LogLevel("debug").
        ConCORS().
        AgregarRuta("/api/users", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintln(w, "Lista de usuarios")
        }).
        AgregarRuta("/api/health", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintln(w, "OK")
        }).
        Build()
    
    servidor.Iniciar()
}
```

### 2️⃣ Prototype Pattern

```go
type Clonable interface {
    Clonar() interface{}
}

type ConfiguracionServicio struct {
    Nombre        string
    Version       string
    Puerto        int
    Database      DatabaseConfig
    Cache         CacheConfig
    Logs          LogConfig
    Caracteristicas map[string]bool
}

type DatabaseConfig struct {
    Host     string
    Puerto   int
    Database string
    Pool     int
}

type CacheConfig struct {
    Host    string
    Puerto  int
    TTL     time.Duration
    MaxKeys int
}

type LogConfig struct {
    Level  string
    Output string
    Format string
}

func (c ConfiguracionServicio) Clonar() interface{} {
    // Crear copia profunda
    clone := ConfiguracionServicio{
        Nombre:  c.Nombre,
        Version: c.Version,
        Puerto:  c.Puerto,
        Database: DatabaseConfig{
            Host:     c.Database.Host,
            Puerto:   c.Database.Puerto,
            Database: c.Database.Database,
            Pool:     c.Database.Pool,
        },
        Cache: CacheConfig{
            Host:    c.Cache.Host,
            Puerto:  c.Cache.Puerto,
            TTL:     c.Cache.TTL,
            MaxKeys: c.Cache.MaxKeys,
        },
        Logs: LogConfig{
            Level:  c.Logs.Level,
            Output: c.Logs.Output,
            Format: c.Logs.Format,
        },
        Caracteristicas: make(map[string]bool),
    }
    
    // Copiar map
    for k, v := range c.Caracteristicas {
        clone.Caracteristicas[k] = v
    }
    
    return clone
}

func ejemploPrototype() {
    // Configuración base
    configBase := ConfiguracionServicio{
        Nombre:  "MiServicio",
        Version: "1.0.0",
        Puerto:  8080,
        Database: DatabaseConfig{
            Host:     "localhost",
            Puerto:   5432,
            Database: "myapp",
            Pool:     10,
        },
        Cache: CacheConfig{
            Host:    "localhost",
            Puerto:  6379,
            TTL:     5 * time.Minute,
            MaxKeys: 1000,
        },
        Logs: LogConfig{
            Level:  "info",
            Output: "stdout",
            Format: "json",
        },
        Caracteristicas: map[string]bool{
            "metrics":     true,
            "healthcheck": true,
            "auth":        false,
        },
    }
    
    // Crear configuración para desarrollo
    configDev := configBase.Clonar().(ConfiguracionServicio)
    configDev.Nombre = "MiServicio-Dev"
    configDev.Puerto = 8081
    configDev.Database.Database = "myapp_dev"
    configDev.Logs.Level = "debug"
    configDev.Caracteristicas["debug"] = true
    
    // Crear configuración para testing
    configTest := configBase.Clonar().(ConfiguracionServicio)
    configTest.Nombre = "MiServicio-Test"
    configTest.Puerto = 8082
    configTest.Database.Database = "myapp_test"
    configTest.Cache.TTL = 1 * time.Minute
    configTest.Caracteristicas["mock"] = true
    
    fmt.Printf("Base: %s en puerto %d\n", configBase.Nombre, configBase.Puerto)
    fmt.Printf("Dev: %s en puerto %d\n", configDev.Nombre, configDev.Puerto)
    fmt.Printf("Test: %s en puerto %d\n", configTest.Nombre, configTest.Puerto)
}
```

---

## ⚠️ Errores Comunes

### 1️⃣ Comparación de Structs

```go
type Punto struct {
    X, Y float64
}

type Complejo struct {
    Real, Imaginario float64
    metadata         map[string]interface{} // ❌ No comparable
}

func erroresComparacion() {
    // ✅ OK - Structs comparables
    p1 := Punto{1, 2}
    p2 := Punto{1, 2}
    p3 := Punto{2, 3}
    
    fmt.Printf("p1 == p2: %v\n", p1 == p2) // true
    fmt.Printf("p1 == p3: %v\n", p1 == p3) // false
    
    // ❌ ERROR - Structs con campos no comparables
    /*
    c1 := Complejo{Real: 1, Imaginario: 2, metadata: make(map[string]interface{})}
    c2 := Complejo{Real: 1, Imaginario: 2, metadata: make(map[string]interface{})}
    fmt.Printf("c1 == c2: %v\n", c1 == c2) // ERROR: invalid operation
    */
    
    // ✅ Solución - Implementar método de comparación
    c1 := Complejo{Real: 1, Imaginario: 2, metadata: make(map[string]interface{})}
    c2 := Complejo{Real: 1, Imaginario: 2, metadata: make(map[string]interface{})}
    
    fmt.Printf("c1 equals c2: %v\n", c1.Equals(c2))
}

func (c Complejo) Equals(other Complejo) bool {
    return c.Real == other.Real && c.Imaginario == other.Imaginario
}
```

### 2️⃣ Receivers por Valor vs Puntero

```go
type Contador struct {
    valor int
}

// ❌ INCORRECTO - Receptor por valor no modifica original
func (c Contador) IncrementarMal() {
    c.valor++
}

// ✅ CORRECTO - Receptor por puntero modifica original
func (c *Contador) Incrementar() {
    c.valor++
}

// ✅ OK - Receptor por valor para métodos que no modifican
func (c Contador) Valor() int {
    return c.valor
}

func erroresReceivers() {
    contador := Contador{valor: 0}
    
    fmt.Printf("Inicial: %d\n", contador.Valor())
    
    // Esto NO funciona
    contador.IncrementarMal()
    fmt.Printf("Después de IncrementarMal: %d\n", contador.Valor()) // Sigue siendo 0
    
    // Esto SÍ funciona
    contador.Incrementar()
    fmt.Printf("Después de Incrementar: %d\n", contador.Valor()) // Ahora es 1
}
```

### 3️⃣ Zero Values y Inicialización

```go
type ConfiguracionErronea struct {
    Database map[string]string // ❌ Zero value es nil
    Logger   *log.Logger       // ❌ Zero value es nil
    Timeout  time.Duration     // ✅ Zero value es 0 (válido)
    Enabled  bool              // ✅ Zero value es false (válido)
}

type ConfiguracionCorrecta struct {
    Database map[string]string
    Logger   *log.Logger
    Timeout  time.Duration
    Enabled  bool
}

func NewConfiguracionCorrecta() *ConfiguracionCorrecta {
    return &ConfiguracionCorrecta{
        Database: make(map[string]string), // ✅ Inicializar map
        Logger:   log.New(os.Stdout, "", log.LstdFlags), // ✅ Inicializar logger
        Timeout:  30 * time.Second,       // ✅ Valor sensato
        Enabled:  true,                   // ✅ Valor por defecto
    }
}

func erroresZeroValues() {
    // ❌ Uso peligroso sin inicializar
    var configMala ConfiguracionErronea
    
    // Esto causará panic
    // configMala.Database["key"] = "value" // panic: assignment to entry in nil map
    
    // ✅ Uso correcto con constructor
    configBuena := NewConfiguracionCorrecta()
    configBuena.Database["host"] = "localhost"
    configBuena.Logger.Println("Configuración lista")
    
    fmt.Printf("Timeout: %v\n", configBuena.Timeout)
}
```

---

## 📝 Resumen de Mejores Prácticas

### ✅ DO (Hacer)

1. **Usar constructores** para inicialización compleja
2. **Nombrar structs con PascalCase** (públicos) o camelCase (privados)
3. **Agrupar campos relacionados** lógicamente
4. **Usar embedding** para composition
5. **Implementar interfaces** cuando sea apropiado
6. **Validar datos** en constructores o métodos
7. **Usar punteros para receivers** que modifican el struct
8. **Documentar structs públicos** con comentarios

### ❌ DON'T (No hacer)

1. **No usar structs con muchos campos** (considerar separar)
2. **No ignorar zero values** problemáticos
3. **No mezclar concerns** en un solo struct
4. **No usar interfaces innecesarias** para structs simples
5. **No olvidar thread-safety** en operaciones concurrentes
6. **No usar reflection** sin necesidad real
7. **No crear dependencias circulares** entre structs
8. **No ignorar el impacto en memoria** de structs grandes

---

## 🎯 Ejercicios Prácticos

¡Ahora es tu turno! Ve al archivo `ejercicios.go` para practicar con ejercicios progresivos que cubren todos los conceptos de esta lección.

## 🚀 Proyecto Final

En el archivo `proyecto_ecommerce.go` encontrarás un proyecto completo que implementa un sistema de e-commerce usando structs de manera avanzada.

---

**¡Felicidades! 🎉 Has completado la lección sobre Structs en Go. Los structs son la base para crear aplicaciones complejas y bien estructuradas en Go.**
