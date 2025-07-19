# 📦 Variables y Constantes: Los Bloques Fundamentales

> *"Variables are the cells of programming; constants are its DNA"* - Programming Wisdom

Variables y constantes son los bloques fundamentales que dan vida a tus programas. En esta lección dominarás cada aspecto de cómo Go maneja el estado y los valores inmutables.

## 🎯 Objetivos de Esta Lección

Al finalizar esta lección serás capaz de:
- ✅ **Dominar todas las formas** de declarar variables
- ✅ **Usar constantes** de forma idiomática y avanzada
- ✅ **Entender scope y shadowing** profundamente
- ✅ **Aplicar zero values** estratégicamente
- ✅ **Optimizar memoria** con decisiones inteligentes
- ✅ **Crear enumeraciones** potentes con iota

---

## 🧬 Variables: El ADN de tus Programas

### 🎨 Formas de Declarar Variables

Go ofrece múltiples formas de declarar variables, cada una con su propósito específico:

```go
package main

import "fmt"

// Variables globales
var globalCounter int                    // Zero value: 0
var applicationName string = "Go Deep"   // Inicializada
var isProduction = false                // Tipo inferido

func main() {
    // 1. Declaración explícita con var
    var userName string
    var userAge int
    var isActive bool
    
    // 2. Declaración con inicialización
    var city string = "Madrid"
    var population int = 3223334
    
    // 3. Declaración con tipo inferido
    var temperature = 25.5      // float64
    var message = "Hola"        // string
    
    // 4. Declaración múltiple
    var (
        firstName string = "Juan"
        lastName  string = "Pérez"
        age       int    = 30
    )
    
    // 5. Short variable declaration (:=)
    email := "juan@example.com"     // Solo dentro de funciones
    score := 95.5                   // Tipo inferido: float64
    
    // 6. Asignación múltiple
    name, surname := "Ana", "García"
    x, y, z := 1, 2, 3
    
    fmt.Printf("Usuario: %s %s\n", firstName, lastName)
    fmt.Printf("Email: %s, Score: %.1f\n", email, score)
}
```

### 🧠 Analogía: Variables como Cajas Etiquetadas

Imagina las variables como **cajas etiquetadas** en un almacén:

```
📦 [userName: ""]     ← Caja vacía (zero value)
📦 [age: 25]          ← Caja con contenido
📦 [isActive: true]   ← Caja con etiqueta específica
📦 [temp: nil]        ← Caja para punteros (puede estar vacía)
```

### 🎯 Cuándo Usar Cada Forma

```go
package main

import (
    "fmt"
    "os"
)

func variableStyleGuide() {
    // ✅ var para zero values intencionales
    var buffer []byte        // nil slice intencional
    var config *Config       // nil pointer hasta configurar
    
    // ✅ var para globals y package-level
    var (
        version    = "1.0.0"
        buildDate  = "2025-01-15"
        commitHash = "abc123"
    )
    
    // ✅ := para variables locales obvias
    user := getCurrentUser()
    count := len(users)
    
    // ✅ var para declarar antes de usar
    var result string
    if condition {
        result = "success"
    } else {
        result = "failure"
    }
    
    // ✅ := para valores de retorno múltiples
    data, err := os.ReadFile("config.json")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Data length: %d\n", len(data))
}
```

---

## 🔄 Assignment vs Declaration

### 📝 Diferencias Cruciales

```go
package main

import "fmt"

func assignmentVsDeclaration() {
    // DECLARACIÓN - Primera vez que aparece la variable
    var name string              // Declaración con zero value
    var age int = 25            // Declaración con inicialización
    email := "test@email.com"   // Declaración corta
    
    // ASIGNACIÓN - Variable ya existe
    name = "Juan"               // Asignación simple
    age = 30                    // Reasignación
    // email := "nuevo@email.com"  // ❌ Error! Ya existe
    email = "nuevo@email.com"   // ✅ Asignación correcta
    
    // MIXED - Algunas nuevas, algunas existentes
    city, country := "Madrid", "España"  // Ambas nuevas
    city, population := "Barcelona", 1620343  // city existe, population nueva
    
    fmt.Printf("User: %s, %d years old\n", name, age)
    fmt.Printf("From: %s, %s (pop: %d)\n", city, country, population)
}
```

### 🚨 Errores Comunes

```go
// ❌ Error 1: Redeclaración en mismo scope
func badExample1() {
    name := "Juan"
    name := "Ana"  // Error: no new variables on left side
}

// ✅ Corrección: Reasignación
func goodExample1() {
    name := "Juan"
    name = "Ana"   // OK: asignación
}

// ❌ Error 2: := en scope global
var global := "value"  // Error: no se puede usar := fuera de funciones

// ✅ Corrección: usar var
var global = "value"   // OK

// ❌ Error 3: Variable no usada
func badExample2() {
    unused := "value"  // Error: unused variable
}

// ✅ Corrección: usar _ para ignorar
func goodExample2() {
    _ = "value"        // OK: blank identifier
    // o simplemente no declarar si no necesitas
}
```

---

## 🌍 Scope: El Reino de las Variables

### 🏰 Tipos de Scope

```go
package main

import "fmt"

// PACKAGE SCOPE - Visible en todo el package
var packageVariable = "I'm package-wide"

func scopeDemo() {
    // FUNCTION SCOPE - Visible en toda la función
    functionVariable := "I'm function-wide"
    
    if true {
        // BLOCK SCOPE - Visible solo en este bloque
        blockVariable := "I'm block-scoped"
        
        // Acceso a scopes externos
        fmt.Println(packageVariable)   // ✅ OK
        fmt.Println(functionVariable)  // ✅ OK
        fmt.Println(blockVariable)     // ✅ OK
    }
    
    // fmt.Println(blockVariable)     // ❌ Error: undefined
    fmt.Println(functionVariable)     // ✅ OK
}

// Ejemplo complejo de scopes anidados
func complexScopeExample() {
    name := "Outer"
    
    for i := 0; i < 3; i++ {
        name := "Loop"  // Shadowing - nueva variable
        
        if i == 1 {
            name := "Inner"  // Más shadowing
            fmt.Printf("Inner scope: %s\n", name)  // "Inner"
        }
        
        fmt.Printf("Loop scope: %s\n", name)   // "Loop"
    }
    
    fmt.Printf("Outer scope: %s\n", name)     // "Outer"
}
```

### 👤 Variable Shadowing

```go
package main

import "fmt"

var x = "global"

func shadowingDemo() {
    fmt.Println("1:", x)  // "global"
    
    x := "function"       // Shadow global x
    fmt.Println("2:", x)  // "function"
    
    {
        x := "block"      // Shadow function x
        fmt.Println("3:", x)  // "block"
        
        {
            x := "inner"  // Shadow block x
            fmt.Println("4:", x)  // "inner"
        }
        
        fmt.Println("5:", x)  // "block"
    }
    
    fmt.Println("6:", x)  // "function"
}

// ⚠️ Shadowing peligroso con err
func dangerousShadowing() error {
    data, err := readFile("file1.txt")
    if err != nil {
        return err
    }
    
    // ⚠️ PELIGRO: err es shadowed aquí
    if needsMoreData {
        data, err := readFile("file2.txt")  // Nueva variable err!
        if err != nil {
            return err  // ✅ Retorna el err local
        }
        // err local se sale de scope aquí
    }
    
    // El err original sigue siendo nil aunque file2 falle!
    return err  // Podría retornar nil inesperadamente
}

// ✅ Forma segura
func safeShadowing() error {
    data, err := readFile("file1.txt")
    if err != nil {
        return err
    }
    
    if needsMoreData {
        moreData, err := readFile("file2.txt")  // Reutiliza err
        if err != nil {
            return err
        }
        data = append(data, moreData...)
    }
    
    return nil
}
```

---

## 📏 Zero Values: El Poder de lo Vacío

### 🎯 Zero Values por Tipo

```go
package main

import "fmt"

func zeroValuesDemo() {
    // Tipos básicos
    var b bool        // false
    var i int         // 0
    var f float64     // 0.0
    var s string      // ""
    
    // Tipos compuestos
    var arr [3]int    // [0, 0, 0]
    var slice []int   // nil
    var m map[string]int  // nil
    var ch chan int   // nil
    
    // Punteros e interfaces
    var ptr *int      // nil
    var iface interface{}  // nil
    
    // Structs
    type Person struct {
        Name string
        Age  int
    }
    var p Person      // {Name: "", Age: 0}
    
    fmt.Printf("bool: %t\n", b)
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float64: %.1f\n", f)
    fmt.Printf("string: '%s'\n", s)
    fmt.Printf("array: %v\n", arr)
    fmt.Printf("slice: %v (nil: %t)\n", slice, slice == nil)
    fmt.Printf("map: %v (nil: %t)\n", m, m == nil)
    fmt.Printf("channel: %v (nil: %t)\n", ch, ch == nil)
    fmt.Printf("pointer: %v (nil: %t)\n", ptr, ptr == nil)
    fmt.Printf("interface: %v (nil: %t)\n", iface, iface == nil)
    fmt.Printf("struct: %+v\n", p)
}
```

### 💡 Diseñando con Zero Values

```go
package main

import (
    "bytes"
    "fmt"
)

// ✅ Buffer aprovecha zero value de bytes.Buffer
type Logger struct {
    buffer bytes.Buffer  // Zero value es buffer vacío funcional
    prefix string
}

// No necesita constructor - zero value funciona!
func (l *Logger) Log(message string) {
    l.buffer.WriteString(l.prefix + message + "\n")
}

func (l *Logger) String() string {
    return l.buffer.String()
}

// ✅ Counter aprovecha zero value de int
type Counter struct {
    value int  // Zero value: 0
}

func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Value() int {
    return c.value
}

// ✅ StringSet aprovecha zero value de map
type StringSet struct {
    items map[string]bool  // nil map
}

func (s *StringSet) Add(item string) {
    if s.items == nil {
        s.items = make(map[string]bool)  // Lazy initialization
    }
    s.items[item] = true
}

func (s *StringSet) Contains(item string) bool {
    return s.items[item]  // Safe even if items is nil
}

func zeroValuePatternsDemo() {
    // Todos funcionan sin inicialización!
    var logger Logger
    logger.Log("First message")
    
    var counter Counter
    counter.Increment()
    counter.Increment()
    
    var set StringSet
    set.Add("hello")
    set.Add("world")
    
    fmt.Printf("Logger output:\n%s", logger.String())
    fmt.Printf("Counter value: %d\n", counter.Value())
    fmt.Printf("Set contains 'hello': %t\n", set.Contains("hello"))
}
```

---

## 🔒 Constantes: Los Inmutables

### 📝 Declaración de Constantes

```go
package main

import (
    "fmt"
    "math"
)

// Constantes globales
const (
    AppName    = "Go Deep"
    Version    = "1.0.0"
    MaxRetries = 3
    Pi         = 3.14159
)

func constantsDemo() {
    // Constantes locales
    const (
        maxUsers    = 1000
        timeout     = 30
        defaultPort = 8080
    )
    
    // Constante única
    const greeting = "¡Hola, Gopher!"
    
    // Constantes tipadas vs no tipadas
    const typedPi float64 = 3.14159      // Tipada
    const untypedPi = 3.14159            // No tipada (más flexible)
    
    // Expresiones constantes
    const (
        kilobyte = 1024
        megabyte = kilobyte * 1024
        gigabyte = megabyte * 1024
        terabyte = gigabyte * 1024
    )
    
    // Usando math constants
    const circleArea = math.Pi * 5 * 5  // Pi del package math
    
    fmt.Printf("App: %s v%s\n", AppName, Version)
    fmt.Printf("1 GB = %d bytes\n", gigabyte)
    fmt.Printf("Circle area: %.2f\n", circleArea)
}
```

### 🎭 Constantes Tipadas vs No Tipadas

```go
package main

import "fmt"

func typedVsUntypedConstants() {
    // Constantes no tipadas - Flexibles
    const untypedInt = 42
    const untypedFloat = 3.14
    
    // Pueden usarse con diferentes tipos
    var i8 int8 = untypedInt      // ✅ OK
    var i16 int16 = untypedInt    // ✅ OK  
    var i32 int32 = untypedInt    // ✅ OK
    var f32 float32 = untypedFloat // ✅ OK
    var f64 float64 = untypedFloat // ✅ OK
    
    // Constantes tipadas - Estrictas
    const typedInt int = 42
    const typedFloat float64 = 3.14
    
    var i int = typedInt          // ✅ OK - mismo tipo
    // var i16 int16 = typedInt   // ❌ Error - tipos diferentes
    var i16 int16 = int16(typedInt) // ✅ OK - conversión explícita
    
    // Demostración de flexibilidad
    fmt.Printf("Untyped used as different types:\n")
    fmt.Printf("int8: %d, int16: %d, int32: %d\n", i8, i16, i32)
    fmt.Printf("float32: %.2f, float64: %.2f\n", f32, f64)
}
```

### 🔢 iota: El Generador de Constantes

```go
package main

import "fmt"

// Ejemplo básico de iota
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

// iota con valores personalizedos
const (
    _  = iota             // 0 - ignorado con blank identifier
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                    // 1 << 20 = 1048576
    GB                    // 1 << 30 = 1073741824
    TB                    // 1 << 40 = 1099511627776
)

// iota con expresiones complejas
const (
    _      = iota                    // 0
    Red    = 1 << (iota - 1)        // 1 << 0 = 1
    Green                           // 1 << 1 = 2
    Blue                            // 1 << 2 = 4
    Yellow = Red | Green            // 1 | 2 = 3
    Cyan   = Green | Blue           // 2 | 4 = 6
    Magenta = Red | Blue            // 1 | 4 = 5
)

// Estados con iota
type ConnectionState int

const (
    Disconnected ConnectionState = iota
    Connecting
    Connected
    Reconnecting
    Failed
)

func (cs ConnectionState) String() string {
    states := []string{
        "Disconnected",
        "Connecting", 
        "Connected",
        "Reconnecting",
        "Failed",
    }
    
    if int(cs) < len(states) {
        return states[cs]
    }
    return "Unknown"
}

// HTTP Status codes con iota
const (
    StatusContinue           = 100 + iota  // 100
    StatusSwitchingProtocols               // 101
    StatusProcessing                       // 102
    StatusEarlyHints                       // 103
)

const (
    StatusOK                   = 200 + iota  // 200
    StatusCreated                           // 201
    StatusAccepted                          // 202
    StatusNonAuthoritativeInfo              // 203
)

func iotaDemo() {
    fmt.Printf("Days of week:\n")
    fmt.Printf("Sunday: %d, Monday: %d, Saturday: %d\n", Sunday, Monday, Saturday)
    
    fmt.Printf("\nFile sizes:\n")
    fmt.Printf("1 KB = %d bytes\n", KB)
    fmt.Printf("1 MB = %d bytes\n", MB)
    fmt.Printf("1 GB = %d bytes\n", GB)
    
    fmt.Printf("\nColors (bit flags):\n")
    fmt.Printf("Red: %d, Green: %d, Blue: %d\n", Red, Green, Blue)
    fmt.Printf("Yellow: %d, Cyan: %d, Magenta: %d\n", Yellow, Cyan, Magenta)
    
    fmt.Printf("\nConnection states:\n")
    var state ConnectionState = Connected
    fmt.Printf("Current state: %s (%d)\n", state, state)
    
    fmt.Printf("\nHTTP Status codes:\n")
    fmt.Printf("Continue: %d, OK: %d, Created: %d\n", 
        StatusContinue, StatusOK, StatusCreated)
}
```

### 🎨 Patrones Avanzados con iota

```go
package main

import "fmt"

// Enum con métodos
type Priority int

const (
    Low Priority = iota
    Medium
    High
    Critical
)

func (p Priority) String() string {
    return [...]string{"Low", "Medium", "High", "Critical"}[p]
}

func (p Priority) Color() string {
    return [...]string{"green", "yellow", "orange", "red"}[p]
}

// Bit flags para permisos
type Permission int

const (
    Read Permission = 1 << iota  // 1
    Write                        // 2
    Execute                      // 4
    Delete                       // 8
)

func (p Permission) String() string {
    var perms []string
    if p&Read != 0 {
        perms = append(perms, "Read")
    }
    if p&Write != 0 {
        perms = append(perms, "Write")
    }
    if p&Execute != 0 {
        perms = append(perms, "Execute")
    }
    if p&Delete != 0 {
        perms = append(perms, "Delete")
    }
    return strings.Join(perms, "|")
}

// Configuración con tamaños
type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
    FATAL
)

var logLevelNames = [...]string{
    DEBUG: "DEBUG",
    INFO:  "INFO", 
    WARN:  "WARN",
    ERROR: "ERROR",
    FATAL: "FATAL",
}

func (ll LogLevel) String() string {
    if ll < 0 || int(ll) >= len(logLevelNames) {
        return "UNKNOWN"
    }
    return logLevelNames[ll]
}

func advancedIotaDemo() {
    // Priority demo
    task := High
    fmt.Printf("Task priority: %s (color: %s)\n", task, task.Color())
    
    // Permission demo
    userPerms := Read | Write | Execute  // Combinar permisos
    fmt.Printf("User permissions: %s\n", userPerms)
    
    adminPerms := Read | Write | Execute | Delete
    fmt.Printf("Admin permissions: %s\n", adminPerms)
    
    // Check specific permission
    if userPerms&Write != 0 {
        fmt.Println("User can write")
    }
    
    // Log level demo
    currentLevel := ERROR
    fmt.Printf("Current log level: %s\n", currentLevel)
}
```

---

## 🧪 Laboratorio: Sistema de Gestión de Estado

### 🎯 Proyecto: User Management System

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

// User roles con iota
type Role int

const (
    Guest Role = iota
    User
    Moderator
    Admin
    SuperAdmin
)

func (r Role) String() string {
    roles := [...]string{
        "Guest", "User", "Moderator", "Admin", "SuperAdmin",
    }
    if int(r) < len(roles) {
        return roles[r]
    }
    return "Unknown"
}

// Permissions con bit flags
type Permission int

const (
    CanRead Permission = 1 << iota
    CanWrite
    CanDelete
    CanBan
    CanPromote
    CanManageSystem
)

func (p Permission) String() string {
    var perms []string
    permMap := map[Permission]string{
        CanRead:         "Read",
        CanWrite:        "Write", 
        CanDelete:       "Delete",
        CanBan:          "Ban",
        CanPromote:      "Promote",
        CanManageSystem: "ManageSystem",
    }
    
    for perm, name := range permMap {
        if p&perm != 0 {
            perms = append(perms, name)
        }
    }
    
    if len(perms) == 0 {
        return "None"
    }
    return strings.Join(perms, "|")
}

// Role permissions mapping
var rolePermissions = map[Role]Permission{
    Guest:      0, // No permissions
    User:       CanRead | CanWrite,
    Moderator:  CanRead | CanWrite | CanDelete | CanBan,
    Admin:      CanRead | CanWrite | CanDelete | CanBan | CanPromote,
    SuperAdmin: CanRead | CanWrite | CanDelete | CanBan | CanPromote | CanManageSystem,
}

// User states
type UserState int

const (
    Inactive UserState = iota
    Active
    Suspended
    Banned
    PendingVerification
)

func (us UserState) String() string {
    states := [...]string{
        "Inactive", "Active", "Suspended", "Banned", "PendingVerification",
    }
    if int(us) < len(states) {
        return states[us]
    }
    return "Unknown"
}

// Configuration constants
const (
    MaxLoginAttempts     = 3
    SessionTimeoutMinutes = 30
    PasswordMinLength    = 8
    UsernameMinLength    = 3
    MaxUsersPerPage      = 50
)

// Size constants
const (
    _          = iota             // ignore first value
    KB         = 1 << (10 * iota) // 1024
    MB                            // 1048576
    GB                            // 1073741824
    
    MaxAvatarSize = 2 * MB        // 2MB
    MaxPostSize   = 10 * KB       // 10KB
)

// User struct using zero values effectively
type User struct {
    ID          int64       // Zero value: 0 (will be set by DB)
    Username    string      // Zero value: "" (must be set)
    Email       string      // Zero value: "" (must be set)
    Role        Role        // Zero value: Guest (default role)
    State       UserState   // Zero value: Inactive (safe default)
    Permissions Permission  // Zero value: 0 (no permissions)
    CreatedAt   time.Time   // Zero value: time zero (will be set)
    LastLoginAt *time.Time  // Zero value: nil (not logged in yet)
    LoginAttempts int       // Zero value: 0 (no failed attempts)
}

// UserManager manages users with smart zero values
type UserManager struct {
    users    map[int64]*User  // nil map (lazy init)
    nextID   int64           // 0 (will start from 1)
    settings *Settings       // nil pointer (optional)
}

type Settings struct {
    AllowGuestAccess bool
    RequireEmailVerification bool
    MaxConcurrentSessions int
}

// CreateUser creates a new user with smart defaults
func (um *UserManager) CreateUser(username, email string) (*User, error) {
    // Lazy initialization of map
    if um.users == nil {
        um.users = make(map[int64]*User)
    }
    
    // Validation
    if len(username) < UsernameMinLength {
        return nil, fmt.Errorf("username too short (min %d chars)", UsernameMinLength)
    }
    
    // Generate ID
    um.nextID++
    
    // Create user with appropriate zero values and defaults
    user := &User{
        ID:          um.nextID,
        Username:    username,
        Email:       email,
        Role:        User,         // Default role (not Guest)
        State:       Active,       // Default to active (not Inactive)
        Permissions: rolePermissions[User], // Set permissions based on role
        CreatedAt:   time.Now(),
        // LastLoginAt stays nil (never logged in)
        // LoginAttempts stays 0
    }
    
    um.users[user.ID] = user
    return user, nil
}

// PromoteUser promotes a user to a new role
func (um *UserManager) PromoteUser(userID int64, newRole Role) error {
    user, exists := um.users[userID]
    if !exists {
        return fmt.Errorf("user not found")
    }
    
    user.Role = newRole
    user.Permissions = rolePermissions[newRole]
    return nil
}

// HasPermission checks if user has specific permission
func (u *User) HasPermission(perm Permission) bool {
    return u.Permissions&perm != 0
}

// CanPerformAction checks if user can perform an action
func (u *User) CanPerformAction(action string) bool {
    if u.State != Active {
        return false
    }
    
    switch action {
    case "read":
        return u.HasPermission(CanRead)
    case "write":
        return u.HasPermission(CanWrite)
    case "delete":
        return u.HasPermission(CanDelete)
    case "ban":
        return u.HasPermission(CanBan)
    case "promote":
        return u.HasPermission(CanPromote)
    case "system":
        return u.HasPermission(CanManageSystem)
    default:
        return false
    }
}

// RecordLogin records a successful login
func (u *User) RecordLogin() {
    now := time.Now()
    u.LastLoginAt = &now
    u.LoginAttempts = 0  // Reset failed attempts
}

// RecordFailedLogin records a failed login attempt
func (u *User) RecordFailedLogin() {
    u.LoginAttempts++
}

// IsLocked checks if user is locked due to failed attempts
func (u *User) IsLocked() bool {
    return u.LoginAttempts >= MaxLoginAttempts
}

func main() {
    fmt.Println("=== User Management System Demo ===\n")
    
    // UserManager starts with zero values (no initialization needed!)
    var manager UserManager
    
    // Create users
    user1, err := manager.CreateUser("alice", "alice@example.com")
    if err != nil {
        fmt.Printf("Error creating user: %v\n", err)
        return
    }
    
    user2, err := manager.CreateUser("bob", "bob@example.com") 
    if err != nil {
        fmt.Printf("Error creating user: %v\n", err)
        return
    }
    
    // Display initial state
    fmt.Printf("Created users:\n")
    fmt.Printf("User 1: %s (Role: %s, State: %s, Permissions: %s)\n",
        user1.Username, user1.Role, user1.State, user1.Permissions)
    fmt.Printf("User 2: %s (Role: %s, State: %s, Permissions: %s)\n",
        user2.Username, user2.Role, user2.State, user2.Permissions)
    
    // Test permissions
    fmt.Printf("\nPermission tests:\n")
    fmt.Printf("Alice can read: %t\n", user1.CanPerformAction("read"))
    fmt.Printf("Alice can ban: %t\n", user1.CanPerformAction("ban"))
    fmt.Printf("Bob can write: %t\n", user2.CanPerformAction("write"))
    
    // Promote user
    fmt.Printf("\nPromoting Alice to Admin...\n")
    manager.PromoteUser(user1.ID, Admin)
    fmt.Printf("Alice new permissions: %s\n", user1.Permissions)
    fmt.Printf("Alice can ban now: %t\n", user1.CanPerformAction("ban"))
    
    // Login simulation
    fmt.Printf("\nLogin simulation:\n")
    fmt.Printf("Bob failed login attempts: %d\n", user2.LoginAttempts)
    user2.RecordFailedLogin()
    user2.RecordFailedLogin()
    fmt.Printf("After 2 failed attempts: %d (locked: %t)\n", 
        user2.LoginAttempts, user2.IsLocked())
    
    user2.RecordFailedLogin()
    fmt.Printf("After 3 failed attempts: %d (locked: %t)\n", 
        user2.LoginAttempts, user2.IsLocked())
    
    // Successful login
    user1.RecordLogin()
    fmt.Printf("Alice last login: %v\n", user1.LastLoginAt)
    
    // Constants demo
    fmt.Printf("\nSystem constants:\n")
    fmt.Printf("Max avatar size: %d bytes (%.1f MB)\n", MaxAvatarSize, float64(MaxAvatarSize)/float64(MB))
    fmt.Printf("Session timeout: %d minutes\n", SessionTimeoutMinutes)
    fmt.Printf("Max login attempts: %d\n", MaxLoginAttempts)
}
```

---

## 🎯 Best Practices

### ✅ Variables

1. **Usa nombres descriptivos** pero concisos
2. **Prefiere := para locals** obvias
3. **Usa var para zero values** intencionales
4. **Evita shadowing** en código crítico
5. **Declara cerca del uso** cuando sea posible

### ✅ Constantes

1. **Usa constantes para valores mágicos**
2. **Prefiere iota para enumeraciones**
3. **Agrupa constantes relacionadas**
4. **Usa tipos custom** para type safety
5. **Documenta constantes complejas**

### ✅ Patrones Recomendados

```go
// ✅ Constantes bien organizadas
const (
    // HTTP Status codes
    StatusOK                 = 200
    StatusBadRequest         = 400
    StatusInternalServerError = 500
    
    // Limits
    MaxRetries    = 3
    TimeoutSecs   = 30
    MaxUploadSize = 10 * MB
)

// ✅ Enums con métodos
type Status int

const (
    Pending Status = iota
    Processing
    Completed
    Failed
)

func (s Status) String() string { /* ... */ }

// ✅ Zero value friendly structs
type Config struct {
    Port     int           // 0 = auto-assign
    Debug    bool          // false = production mode
    LogLevel string        // "" = default level
    Features map[string]bool // nil = no features
}
```

---

## 🎉 ¡Felicitaciones!

¡Has dominado variables y constantes en Go! Ahora puedes:

- ✅ **Declarar variables** de todas las formas posibles
- ✅ **Manejar scope** y evitar shadowing peligroso
- ✅ **Usar zero values** estratégicamente
- ✅ **Crear constantes** poderosas con iota
- ✅ **Diseñar APIs** que aprovechan valores por defecto

### 🔥 Conceptos Dominados:

1. **Declaración de variables** - var, :=, múltiples formas
2. **Scope management** - Package, function, block scope
3. **Zero values** - Diseño inteligente con valores por defecto
4. **Constantes** - Tipadas, no tipadas, expresiones
5. **iota** - Generación automática de constantes
6. **Enumeraciones** - Patrones idiomáticos con tipos custom

### 🚀 Próximo Nivel

¡Es hora de dominar los operadores que dan poder a tus variables!

**[→ Ir a la Lección 6: Operadores](../06-operadores/)**

---

## 📞 ¿Preguntas?

- 💬 **Discord**: [Go Deep Community](#)
- 📧 **Email**: support@go-deep.dev
- 🐛 **Issues**: [GitHub Issues](../../../issues)

---

*¡Tus variables están bajo control! Hora de operarlas como un pro ⚡*
