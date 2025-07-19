# 🗺️ Lección 10: Maps

## 🎯 Objetivos de Aprendizaje

Al completar esta lección, serás capaz de:

- Entender qué son los maps y cuándo usarlos
- Crear e inicializar maps de diferentes maneras
- Realizar operaciones CRUD (Create, Read, Update, Delete) eficientemente
- Manejar la ausencia de claves con el patrón "comma ok"
- Iterar sobre maps con for-range
- Aplicar maps en casos de uso reales
- Optimizar el rendimiento de operaciones con maps
- Implementar estructuras de datos avanzadas usando maps

## 📖 Contenido

### 1. Introducción a los Maps

Los **maps** en Go son estructuras de datos que asocian **claves** con **valores**. Son equivalentes a hash tables, diccionarios o arrays asociativos en otros lenguajes.

#### Características de los Maps:
- **Clave-Valor**: Cada elemento es un par clave-valor
- **Claves Únicas**: No puede haber claves duplicadas
- **Tipos Específicos**: Claves y valores tienen tipos definidos
- **Referencia**: Los maps son tipos de referencia
- **Desordenados**: No mantienen orden de inserción
- **Zero Value**: El zero value de un map es `nil`

#### Sintaxis Básica

```go
// Declaración del tipo map
var mapVariable map[TipoClave]TipoValor

// Ejemplos de tipos de maps
var edades map[string]int           // string -> int
var precios map[int]float64         // int -> float64
var configuracion map[string]bool   // string -> bool
var usuarios map[int]Usuario        // int -> struct
```

### 2. Creación e Inicialización de Maps

#### Método 1: Usando make()

```go
func ejemploMake() {
    // Crear map vacío
    edades := make(map[string]int)
    
    // Agregar elementos
    edades["Ana"] = 25
    edades["Carlos"] = 30
    edades["María"] = 28
    
    fmt.Printf("Edades: %v\n", edades)
    
    // Map con capacidad inicial (optimización)
    scores := make(map[string]int, 100)
    scores["jugador1"] = 1500
    scores["jugador2"] = 1200
}
```

#### Método 2: Map Literals

```go
func ejemploLiterales() {
    // Inicialización directa
    colores := map[string]string{
        "rojo":   "#FF0000",
        "verde":  "#00FF00",
        "azul":   "#0000FF",
        "negro":  "#000000",
        "blanco": "#FFFFFF",
    }
    
    // Map de structs
    type Producto struct {
        Nombre string
        Precio float64
    }
    
    inventario := map[int]Producto{
        1001: {Nombre: "Laptop", Precio: 999.99},
        1002: {Nombre: "Mouse", Precio: 25.50},
        1003: {Nombre: "Teclado", Precio: 75.00},
    }
    
    fmt.Printf("Inventario: %v\n", inventario)
}
```

#### Método 3: Maps de Maps (Anidados)

```go
func ejemploMapsAnidados() {
    // Map bidimensional
    matriz := map[int]map[int]string{
        0: {0: "A", 1: "B", 2: "C"},
        1: {0: "D", 1: "E", 2: "F"},
        2: {0: "G", 1: "H", 2: "I"},
    }
    
    // Acceder a elementos anidados
    fmt.Printf("Elemento [1][1]: %s\n", matriz[1][1]) // "E"
    
    // Configuración jerárquica
    config := map[string]map[string]interface{}{
        "database": {
            "host":     "localhost",
            "port":     5432,
            "ssl":      true,
            "timeout":  30,
        },
        "redis": {
            "host":     "redis.ejemplo.com",
            "port":     6379,
            "password": "secreto",
        },
    }
    
    // Acceso seguro a configuración anidada
    if dbConfig, exists := config["database"]; exists {
        if host, ok := dbConfig["host"].(string); ok {
            fmt.Printf("Database host: %s\n", host)
        }
    }
}
```

### 3. Operaciones Básicas con Maps

#### Insertar y Actualizar

```go
func operacionesBasicas() {
    usuarios := make(map[string]int)
    
    // Insertar nuevos elementos
    usuarios["alice"] = 100
    usuarios["bob"] = 200
    usuarios["charlie"] = 150
    
    // Actualizar elementos existentes
    usuarios["alice"] = 120  // Alice ahora tiene 120 puntos
    
    fmt.Printf("Usuarios: %v\n", usuarios)
}
```

#### Leer con el Patrón "Comma OK"

```go
func patternCommaOK() {
    puntuaciones := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    
    // Leer valor - forma básica
    puntajeAlice := puntuaciones["Alice"]
    fmt.Printf("Puntaje Alice: %d\n", puntajeAlice)
    
    // Problema: qué pasa si la clave no existe?
    puntajeDavid := puntuaciones["David"] // Retorna 0 (zero value)
    fmt.Printf("Puntaje David: %d\n", puntajeDavid) // 0
    
    // Solución: Patrón "comma ok"
    if puntaje, existe := puntuaciones["David"]; existe {
        fmt.Printf("David tiene %d puntos\n", puntaje)
    } else {
        fmt.Println("David no está en el mapa")
    }
    
    // Ejemplo práctico: contador de palabras
    contador := make(map[string]int)
    palabras := []string{"go", "es", "genial", "go", "es", "rápido", "go"}
    
    for _, palabra := range palabras {
        // Incrementar contador (0 si no existe)
        contador[palabra]++
    }
    
    fmt.Printf("Contador de palabras: %v\n", contador)
}
```

#### Eliminar Elementos

```go
func eliminarElementos() {
    inventario := map[string]int{
        "manzanas": 50,
        "peras":    30,
        "bananas":  25,
        "naranjas": 40,
    }
    
    fmt.Printf("Inventario inicial: %v\n", inventario)
    
    // Eliminar elemento específico
    delete(inventario, "peras")
    fmt.Printf("Después de eliminar peras: %v\n", inventario)
    
    // Verificar si existe antes de eliminar
    if _, existe := inventario["kiwis"]; existe {
        delete(inventario, "kiwis")
        fmt.Println("Kiwis eliminados")
    } else {
        fmt.Println("Kiwis no estaban en inventario")
    }
    
    // Limpiar todo el map
    for clave := range inventario {
        delete(inventario, clave)
    }
    fmt.Printf("Inventario limpio: %v\n", inventario)
}
```

### 4. Iteración sobre Maps

#### For-Range Básico

```go
func iteracionBasica() {
    notas := map[string]float64{
        "Matemáticas": 8.5,
        "Física":      9.0,
        "Química":     7.8,
        "Historia":    8.2,
        "Literatura":  9.5,
    }
    
    // Iterar sobre clave y valor
    fmt.Println("Notas del estudiante:")
    for materia, nota := range notas {
        fmt.Printf("%s: %.1f\n", materia, nota)
    }
    
    // Solo claves
    fmt.Println("\nMaterias:")
    for materia := range notas {
        fmt.Printf("- %s\n", materia)
    }
    
    // Solo valores (usando blank identifier)
    var suma float64
    for _, nota := range notas {
        suma += nota
    }
    promedio := suma / float64(len(notas))
    fmt.Printf("\nPromedio: %.2f\n", promedio)
}
```

#### Iteración Ordenada

```go
func iteracionOrdenada() {
    ventas := map[string]int{
        "enero":     15000,
        "febrero":   18000,
        "marzo":     22000,
        "abril":     19500,
        "mayo":      25000,
    }
    
    // Los maps no garantizan orden, así que ordenamos las claves
    var meses []string
    for mes := range ventas {
        meses = append(meses, mes)
    }
    
    // Ordenar claves
    sort.Strings(meses)
    
    // Iterar en orden
    fmt.Println("Ventas por mes (ordenado):")
    for _, mes := range meses {
        fmt.Printf("%s: $%d\n", mes, ventas[mes])
    }
}
```

### 5. Maps Avanzados

#### Maps de Slices

```go
func mapsDeSlices() {
    // Agrupar estudiantes por grado
    estudiantesPorGrado := map[int][]string{
        9:  {"Ana", "Carlos", "Diana"},
        10: {"Eduardo", "Fernanda", "Gabriel"},
        11: {"Helena", "Ignacio", "Julia"},
        12: {"Kevin", "Laura", "Miguel"},
    }
    
    // Agregar estudiante a un grado
    estudiantesPorGrado[10] = append(estudiantesPorGrado[10], "Nuevo Estudiante")
    
    // Mostrar todos los estudiantes
    for grado, estudiantes := range estudiantesPorGrado {
        fmt.Printf("Grado %d: %v\n", grado, estudiantes)
    }
    
    // Categorizar tareas por prioridad
    tareasPorPrioridad := make(map[string][]string)
    
    agregarTarea := func(prioridad, tarea string) {
        tareasPorPrioridad[prioridad] = append(tareasPorPrioridad[prioridad], tarea)
    }
    
    agregarTarea("alta", "Revisar código crítico")
    agregarTarea("alta", "Solucionar bug de producción")
    agregarTarea("media", "Escribir documentación")
    agregarTarea("baja", "Actualizar README")
    
    fmt.Println("\nTareas por prioridad:")
    for prioridad, tareas := range tareasPorPrioridad {
        fmt.Printf("%s: %v\n", prioridad, tareas)
    }
}
```

#### Maps de Maps (Matrices Dispersas)

```go
func matrizDispersa() {
    // Representar una matriz dispersa usando map de maps
    type Matriz map[int]map[int]float64
    
    crearMatriz := func() Matriz {
        return make(Matriz)
    }
    
    establecer := func(m Matriz, fila, col int, valor float64) {
        if m[fila] == nil {
            m[fila] = make(map[int]float64)
        }
        m[fila][col] = valor
    }
    
    obtener := func(m Matriz, fila, col int) float64 {
        if filaMap, existe := m[fila]; existe {
            if valor, existe := filaMap[col]; existe {
                return valor
            }
        }
        return 0.0 // Valor por defecto para posiciones vacías
    }
    
    // Crear y llenar matriz
    matriz := crearMatriz()
    establecer(matriz, 0, 0, 1.5)
    establecer(matriz, 0, 2, 3.7)
    establecer(matriz, 1, 1, 2.8)
    establecer(matriz, 2, 0, 4.2)
    establecer(matriz, 2, 2, 5.9)
    
    // Mostrar matriz
    fmt.Println("Matriz dispersa:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            valor := obtener(matriz, i, j)
            fmt.Printf("%.1f ", valor)
        }
        fmt.Println()
    }
}
```

#### Maps con Structs Complejos

```go
type Usuario struct {
    ID       int
    Nombre   string
    Email    string
    Activo   bool
    Metadata map[string]interface{}
}

func mapsConStructs() {
    usuarios := map[string]Usuario{
        "user001": {
            ID:     1,
            Nombre: "Ana García",
            Email:  "ana@ejemplo.com",
            Activo: true,
            Metadata: map[string]interface{}{
                "last_login":    "2024-01-15",
                "login_count":   42,
                "preferences":   map[string]string{"theme": "dark", "lang": "es"},
                "notifications": true,
            },
        },
        "user002": {
            ID:     2,
            Nombre: "Carlos López",
            Email:  "carlos@ejemplo.com",
            Activo: false,
            Metadata: map[string]interface{}{
                "last_login":  "2023-12-20",
                "login_count": 15,
                "suspended":   true,
                "reason":      "Violation of terms",
            },
        },
    }
    
    // Buscar usuario activo
    for userID, usuario := range usuarios {
        if usuario.Activo {
            fmt.Printf("Usuario activo: %s (%s)\n", usuario.Nombre, userID)
            
            // Acceder a metadata
            if loginCount, ok := usuario.Metadata["login_count"].(int); ok {
                fmt.Printf("  Logins: %d\n", loginCount)
            }
        }
    }
    
    // Actualizar usuario
    if usuario, existe := usuarios["user002"]; existe {
        usuario.Activo = true
        delete(usuario.Metadata, "suspended")
        usuario.Metadata["reactivated_date"] = "2024-01-16"
        usuarios["user002"] = usuario // Importante: reasignar al map
        fmt.Println("Usuario user002 reactivado")
    }
}
```

### 6. Patrones Comunes con Maps

#### Caché y Memoización

```go
func patronCache() {
    // Cache simple para funciones costosas
    type Cache map[int]int
    
    var fibonacciCache = make(Cache)
    
    fibonacci := func(n int) int {
        // Verificar cache primero
        if resultado, cached := fibonacciCache[n]; cached {
            return resultado
        }
        
        var resultado int
        if n <= 1 {
            resultado = n
        } else {
            resultado = fibonacci(n-1) + fibonacci(n-2)
        }
        
        // Guardar en cache
        fibonacciCache[n] = resultado
        return resultado
    }
    
    // Usar función con cache
    fmt.Println("Fibonacci con caché:")
    for i := 0; i <= 10; i++ {
        fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
    }
    
    fmt.Printf("Cache: %v\n", fibonacciCache)
}
```

#### Índices e Inversiones

```go
func patronIndices() {
    // Lista de empleados
    empleados := []struct {
        ID         int
        Nombre     string
        Departamento string
        Sueldo     float64
    }{
        {1, "Ana García", "IT", 75000},
        {2, "Carlos López", "Ventas", 60000},
        {3, "María Rodríguez", "IT", 80000},
        {4, "Juan Pérez", "Marketing", 55000},
        {5, "Laura Martín", "Ventas", 62000},
    }
    
    // Crear índices
    porID := make(map[int]string)
    porDepartamento := make(map[string][]string)
    porRangoSueldo := make(map[string][]string)
    
    for _, emp := range empleados {
        // Índice por ID
        porID[emp.ID] = emp.Nombre
        
        // Índice por departamento
        porDepartamento[emp.Departamento] = append(
            porDepartamento[emp.Departamento], emp.Nombre,
        )
        
        // Índice por rango de sueldo
        var rangoSueldo string
        switch {
        case emp.Sueldo < 60000:
            rangoSueldo = "bajo"
        case emp.Sueldo < 75000:
            rangoSueldo = "medio"
        default:
            rangoSueldo = "alto"
        }
        porRangoSueldo[rangoSueldo] = append(
            porRangoSueldo[rangoSueldo], emp.Nombre,
        )
    }
    
    // Usar índices
    fmt.Printf("Empleado con ID 3: %s\n", porID[3])
    fmt.Printf("Empleados de IT: %v\n", porDepartamento["IT"])
    fmt.Printf("Empleados con sueldo alto: %v\n", porRangoSueldo["alto"])
}
```

#### Set (Conjunto) usando Maps

```go
func patronSet() {
    // Implementar Set usando map[T]bool
    type StringSet map[string]bool
    
    nuevoSet := func() StringSet {
        return make(StringSet)
    }
    
    agregar := func(s StringSet, elemento string) {
        s[elemento] = true
    }
    
    contiene := func(s StringSet, elemento string) bool {
        return s[elemento]
    }
    
    eliminar := func(s StringSet, elemento string) {
        delete(s, elemento)
    }
    
    elementos := func(s StringSet) []string {
        var resultado []string
        for elemento := range s {
            resultado = append(resultado, elemento)
        }
        return resultado
    }
    
    union := func(s1, s2 StringSet) StringSet {
        resultado := nuevoSet()
        for elemento := range s1 {
            agregar(resultado, elemento)
        }
        for elemento := range s2 {
            agregar(resultado, elemento)
        }
        return resultado
    }
    
    interseccion := func(s1, s2 StringSet) StringSet {
        resultado := nuevoSet()
        for elemento := range s1 {
            if contiene(s2, elemento) {
                agregar(resultado, elemento)
            }
        }
        return resultado
    }
    
    // Usar Set
    set1 := nuevoSet()
    agregar(set1, "go")
    agregar(set1, "python")
    agregar(set1, "java")
    
    set2 := nuevoSet()
    agregar(set2, "go")
    agregar(set2, "rust")
    agregar(set2, "c++")
    
    fmt.Printf("Set1: %v\n", elementos(set1))
    fmt.Printf("Set2: %v\n", elementos(set2))
    fmt.Printf("Unión: %v\n", elementos(union(set1, set2)))
    fmt.Printf("Intersección: %v\n", elementos(interseccion(set1, set2)))
}
```

### 7. Optimización y Rendimiento

#### Preasignación de Capacidad

```go
func optimizacionCapacidad() {
    // Ineficiente: muchas reasignaciones
    mapSinCapacidad := make(map[int]string)
    
    // Eficiente: capacidad preasignada
    mapConCapacidad := make(map[int]string, 10000)
    
    // Benchmark básico
    numElementos := 10000
    
    inicio := time.Now()
    for i := 0; i < numElementos; i++ {
        mapSinCapacidad[i] = fmt.Sprintf("valor_%d", i)
    }
    tiempoSinCapacidad := time.Since(inicio)
    
    inicio = time.Now()
    for i := 0; i < numElementos; i++ {
        mapConCapacidad[i] = fmt.Sprintf("valor_%d", i)
    }
    tiempoConCapacidad := time.Since(inicio)
    
    fmt.Printf("Sin capacidad: %v\n", tiempoSinCapacidad)
    fmt.Printf("Con capacidad: %v\n", tiempoConCapacidad)
}
```

#### Evitar Allocaciones Innecesarias

```go
func optimizacionAllocaciones() {
    datos := []string{"apple", "banana", "cherry", "date", "elderberry"}
    
    // Ineficiente: crear string en cada comparación
    contarLargosIneficiente := func(items []string, longitud int) int {
        count := 0
        cache := make(map[string]int) // Recreado cada vez
        
        for _, item := range items {
            if len(item) == longitud {
                cache[item] = len(item) // Redundante
                count++
            }
        }
        return count
    }
    
    // Eficiente: reutilizar estructuras
    contarLargosEficiente := func(items []string, longitud int) int {
        count := 0
        for _, item := range items {
            if len(item) == longitud {
                count++
            }
        }
        return count
    }
    
    resultado1 := contarLargosIneficiente(datos, 5)
    resultado2 := contarLargosEficiente(datos, 5)
    
    fmt.Printf("Resultado ineficiente: %d\n", resultado1)
    fmt.Printf("Resultado eficiente: %d\n", resultado2)
}
```

### 8. Maps Thread-Safe

#### Usando Mutex

```go
import (
    "sync"
)

type MapThreadSafe struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func NuevoMapThreadSafe() *MapThreadSafe {
    return &MapThreadSafe{
        data: make(map[string]interface{}),
    }
}

func (m *MapThreadSafe) Set(clave string, valor interface{}) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.data[clave] = valor
}

func (m *MapThreadSafe) Get(clave string) (interface{}, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    valor, existe := m.data[clave]
    return valor, existe
}

func (m *MapThreadSafe) Delete(clave string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    delete(m.data, clave)
}

func (m *MapThreadSafe) Len() int {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return len(m.data)
}

func ejemploThreadSafe() {
    mapSeguro := NuevoMapThreadSafe()
    
    // Simular acceso concurrente
    var wg sync.WaitGroup
    
    // Escritores
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            mapSeguro.Set(fmt.Sprintf("key_%d", id), id*10)
        }(i)
    }
    
    // Lectores
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            if valor, existe := mapSeguro.Get(fmt.Sprintf("key_%d", id)); existe {
                fmt.Printf("Leído: key_%d = %v\n", id, valor)
            }
        }(i)
    }
    
    wg.Wait()
    fmt.Printf("Elementos finales: %d\n", mapSeguro.Len())
}
```

#### Usando sync.Map (Go 1.9+)

```go
func ejemploSyncMap() {
    var mapConcurrente sync.Map
    
    // Almacenar valores
    mapConcurrente.Store("clave1", "valor1")
    mapConcurrente.Store("clave2", 42)
    mapConcurrente.Store("clave3", []string{"a", "b", "c"})
    
    // Leer valores
    if valor, ok := mapConcurrente.Load("clave1"); ok {
        fmt.Printf("clave1: %v\n", valor)
    }
    
    // Leer o almacenar (atomic)
    actual, loaded := mapConcurrente.LoadOrStore("clave4", "nuevo valor")
    fmt.Printf("clave4: %v, existía: %t\n", actual, loaded)
    
    // Iterar sobre todos los elementos
    mapConcurrente.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true // continuar iteración
    })
    
    // Eliminar
    mapConcurrente.Delete("clave2")
}
```

### 9. Casos de Uso Avanzados

#### Sistema de Configuración

```go
type ConfigManager struct {
    configs map[string]map[string]interface{}
    defaults map[string]interface{}
}

func NuevoConfigManager() *ConfigManager {
    return &ConfigManager{
        configs: make(map[string]map[string]interface{}),
        defaults: map[string]interface{}{
            "timeout": 30,
            "retries": 3,
            "debug":   false,
        },
    }
}

func (cm *ConfigManager) SetConfig(modulo, clave string, valor interface{}) {
    if cm.configs[modulo] == nil {
        cm.configs[modulo] = make(map[string]interface{})
    }
    cm.configs[modulo][clave] = valor
}

func (cm *ConfigManager) GetConfig(modulo, clave string) interface{} {
    if moduloConfig, existe := cm.configs[modulo]; existe {
        if valor, existe := moduloConfig[clave]; existe {
            return valor
        }
    }
    
    // Fallback a defaults
    if defaultVal, existe := cm.defaults[clave]; existe {
        return defaultVal
    }
    
    return nil
}

func ejemploConfigManager() {
    config := NuevoConfigManager()
    
    // Configurar módulos
    config.SetConfig("database", "host", "localhost")
    config.SetConfig("database", "port", 5432)
    config.SetConfig("database", "timeout", 60)
    
    config.SetConfig("redis", "host", "redis.ejemplo.com")
    config.SetConfig("redis", "port", 6379)
    
    // Leer configuraciones
    fmt.Printf("DB Host: %v\n", config.GetConfig("database", "host"))
    fmt.Printf("DB Timeout: %v\n", config.GetConfig("database", "timeout"))
    fmt.Printf("Redis Debug: %v\n", config.GetConfig("redis", "debug")) // Usa default
}
```

#### Sistema de Enrutamiento

```go
type Router struct {
    routes map[string]map[string]func(string) string
}

func NuevoRouter() *Router {
    return &Router{
        routes: make(map[string]map[string]func(string) string),
    }
}

func (r *Router) AddRoute(metodo, path string, handler func(string) string) {
    if r.routes[metodo] == nil {
        r.routes[metodo] = make(map[string]func(string) string)
    }
    r.routes[metodo][path] = handler
}

func (r *Router) HandleRequest(metodo, path, data string) string {
    if metodosMap, existe := r.routes[metodo]; existe {
        if handler, existe := metodosMap[path]; existe {
            return handler(data)
        }
    }
    return "404 Not Found"
}

func ejemploRouter() {
    router := NuevoRouter()
    
    // Definir rutas
    router.AddRoute("GET", "/users", func(data string) string {
        return "Lista de usuarios"
    })
    
    router.AddRoute("POST", "/users", func(data string) string {
        return fmt.Sprintf("Usuario creado: %s", data)
    })
    
    router.AddRoute("GET", "/products", func(data string) string {
        return "Lista de productos"
    })
    
    // Manejar requests
    fmt.Println(router.HandleRequest("GET", "/users", ""))
    fmt.Println(router.HandleRequest("POST", "/users", "{'name':'Juan'}"))
    fmt.Println(router.HandleRequest("DELETE", "/users", ""))
}
```

## 🎯 Mejores Prácticas

### 1. **Inicialización**
```go
// ✅ Bueno: usar make() para maps vacíos
m := make(map[string]int)

// ✅ Bueno: usar map literals para inicialización
m := map[string]int{
    "key1": 1,
    "key2": 2,
}

// ❌ Malo: no inicializar
var m map[string]int
m["key"] = 1 // panic: assignment to entry in nil map
```

### 2. **Verificación de Existencia**
```go
// ✅ Bueno: usar comma ok idiom
if valor, existe := mapa["clave"]; existe {
    // Usar valor
}

// ❌ Malo: asumir que existe
valor := mapa["clave"] // Podría ser zero value
```

### 3. **Iteración Segura**
```go
// ✅ Bueno: copiar claves si vas a modificar durante iteración
var claves []string
for clave := range mapa {
    claves = append(claves, clave)
}
for _, clave := range claves {
    delete(mapa, clave)
}

// ❌ Malo: modificar durante iteración directa
for clave := range mapa {
    delete(mapa, clave) // Comportamiento impredecible
}
```

### 4. **Capacidad**
```go
// ✅ Bueno: especificar capacidad si se conoce el tamaño
m := make(map[string]int, 1000)

// ✅ Aceptable: para maps pequeños
m := make(map[string]int)
```

## ⚠️ Errores Comunes

### 1. **Map no Inicializado**
```go
// ❌ Error común
var m map[string]int
m["key"] = 1 // panic!

// ✅ Corrección
m := make(map[string]int)
m["key"] = 1
```

### 2. **Modificación Durante Iteración**
```go
// ❌ Problemático
for k, v := range m {
    if v == 0 {
        delete(m, k) // Puede causar problemas
    }
}

// ✅ Seguro
var keysToDelete []string
for k, v := range m {
    if v == 0 {
        keysToDelete = append(keysToDelete, k)
    }
}
for _, k := range keysToDelete {
    delete(m, k)
}
```

### 3. **Dependencia del Orden**
```go
// ❌ Malo: asumir orden en maps
for k, v := range m {
    // No asumas que aparecerán en orden específico
}

// ✅ Bueno: ordenar claves si necesitas orden
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    // Procesar en orden
}
```

## 📚 Ejercicios Prácticos

Completa los ejercicios en el archivo `ejercicios.go` para practicar:

1. **Operaciones básicas** con maps
2. **Contador de frecuencias** de palabras
3. **Sistema de caché** simple
4. **Índice invertido** para búsquedas
5. **Agrupación de datos** por criterios
6. **Map thread-safe** personalizado
7. **Sistema de configuración** jerárquico
8. **Router HTTP** básico

## 🔗 Próxima Lección

En la **Lección 11: Structs**, aprenderemos sobre:
- Definición y uso de structs
- Struct literals y inicialización
- Embedding y composición
- Anonymous structs
- Struct tags

## 📚 Recursos Adicionales

- [Go Maps in Action](https://go.dev/blog/maps)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
- [Go Data Structures](https://research.swtch.com/godata)

---
**Recuerda**: Los maps son fundamentales en Go para asociar datos. Dominar su uso eficiente te permitirá escribir código más expresivo y performante.
