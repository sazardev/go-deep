# 🔄 Lección 7: Estructuras de Control

> **Nivel**: Fundamentos  
> **Duración estimada**: 2-3 horas  
> **Prerrequisitos**: Variables, constantes y operadores

## 📋 Objetivos de Aprendizaje

Al finalizar esta lección, podrás:

- ✅ Usar condicionales `if`, `else if`, `else` efectivamente
- ✅ Implementar bucles `for` en todas sus variantes
- ✅ Dominar el statement `switch` y type switch
- ✅ Controlar el flujo con `break`, `continue`, `goto` y `defer`
- ✅ Escribir código limpio y eficiente con estructuras de control
- ✅ Aplicar patrones comunes de control de flujo en Go

## 🎯 ¿Por Qué Son Importantes las Estructuras de Control?

Las estructuras de control son el **corazón** de la programación. Te permiten:

- **🔀 Tomar decisiones**: Ejecutar código basado en condiciones
- **🔁 Repetir tareas**: Automatizar procesos repetitivos  
- **🎛️ Controlar flujo**: Dirigir la ejecución del programa
- **⚡ Optimizar performance**: Escribir algoritmos eficientes

---

## 🔀 1. Condicionales (if, else if, else)

### Sintaxis Básica

```go
if condición {
    // código a ejecutar si la condición es verdadera
}
```

### Ejemplo Básico

```go
package main

import "fmt"

func main() {
    edad := 18
    
    if edad >= 18 {
        fmt.Println("Eres mayor de edad")
    }
}
```

### if-else

```go
func evaluarEdad(edad int) {
    if edad >= 18 {
        fmt.Println("Eres mayor de edad")
    } else {
        fmt.Println("Eres menor de edad")
    }
}
```

### if-else if-else

```go
func clasificarEdad(edad int) {
    if edad < 13 {
        fmt.Println("Niño")
    } else if edad < 20 {
        fmt.Println("Adolescente")
    } else if edad < 60 {
        fmt.Println("Adulto")
    } else {
        fmt.Println("Adulto mayor")
    }
}
```

### 🌟 Inicialización en if

Go permite inicializar variables dentro del `if`:

```go
func verificarUsuario(id int) {
    // Variable inicializada en el if
    if usuario, encontrado := buscarUsuario(id); encontrado {
        fmt.Printf("Usuario encontrado: %s\n", usuario.Nombre)
    } else {
        fmt.Println("Usuario no encontrado")
    }
    // 'usuario' no existe fuera del bloque if
}
```

### Ejemplo Práctico: Validador de Contraseña

```go
func validarContrasena(password string) bool {
    if len(password) < 8 {
        fmt.Println("❌ La contraseña debe tener al menos 8 caracteres")
        return false
    }
    
    if !tieneNumero(password) {
        fmt.Println("❌ La contraseña debe contener al menos un número")
        return false
    }
    
    if !tieneMayuscula(password) {
        fmt.Println("❌ La contraseña debe contener al menos una mayúscula")
        return false
    }
    
    fmt.Println("✅ Contraseña válida")
    return true
}
```

---

## 🔁 2. Bucles (for)

### Go Solo Tiene `for`

A diferencia de otros lenguajes, Go **solo** tiene el bucle `for`, pero es muy versátil.

### Sintaxis Básica

```go
for inicialización; condición; incremento {
    // código a repetir
}
```

### Ejemplo Básico

```go
func main() {
    // Bucle clásico tipo C
    for i := 0; i < 5; i++ {
        fmt.Printf("Iteración %d\n", i)
    }
}
```

### 🔄 Variantes del for

#### 1. While Loop (Solo Condición)

```go
func contarHasta(limite int) {
    i := 0
    for i < limite {
        fmt.Printf("Contador: %d\n", i)
        i++
    }
}
```

#### 2. Bucle Infinito

```go
func servidor() {
    for {
        // Bucle infinito
        procesarSolicitud()
        if debeParar {
            break
        }
    }
}
```

#### 3. Range Loop (Arrays, Slices, Maps)

```go
func ejemplosRange() {
    // Array/Slice
    numeros := []int{1, 2, 3, 4, 5}
    
    // Índice y valor
    for i, valor := range numeros {
        fmt.Printf("Índice %d: %d\n", i, valor)
    }
    
    // Solo valor
    for _, valor := range numeros {
        fmt.Printf("Valor: %d\n", valor)
    }
    
    // Solo índice
    for i := range numeros {
        fmt.Printf("Índice: %d\n", i)
    }
}
```

#### 4. Range con Maps

```go
func iterarMap() {
    estudiantes := map[string]int{
        "Ana":   20,
        "Carlos": 22,
        "Elena":  19,
    }
    
    for nombre, edad := range estudiantes {
        fmt.Printf("%s tiene %d años\n", nombre, edad)
    }
}
```

#### 5. Range con Strings

```go
func analizarTexto(texto string) {
    for i, caracter := range texto {
        fmt.Printf("Posición %d: %c (Unicode: %d)\n", i, caracter, caracter)
    }
}
```

### Ejemplo Práctico: Calculadora de Estadísticas

```go
func calcularEstadisticas(numeros []float64) {
    if len(numeros) == 0 {
        fmt.Println("No hay datos para procesar")
        return
    }
    
    var suma, min, max float64
    min = numeros[0]
    max = numeros[0]
    
    for i, num := range numeros {
        suma += num
        
        if num < min {
            min = num
        }
        
        if num > max {
            max = num
        }
        
        // Mostrar progreso cada 1000 elementos
        if i%1000 == 0 && i > 0 {
            fmt.Printf("Procesando... %d/%d elementos\n", i, len(numeros))
        }
    }
    
    promedio := suma / float64(len(numeros))
    
    fmt.Printf("📊 Estadísticas:\n")
    fmt.Printf("   Total elementos: %d\n", len(numeros))
    fmt.Printf("   Suma: %.2f\n", suma)
    fmt.Printf("   Promedio: %.2f\n", promedio)
    fmt.Printf("   Mínimo: %.2f\n", min)
    fmt.Printf("   Máximo: %.2f\n", max)
}
```

---

## 🎛️ 3. Switch Statement

### Sintaxis Básica

```go
switch variable {
case valor1:
    // código
case valor2:
    // código
default:
    // código por defecto
}
```

### Ejemplo Básico

```go
func obtenerDiaSemana(dia int) string {
    switch dia {
    case 1:
        return "Lunes"
    case 2:
        return "Martes"
    case 3:
        return "Miércoles"
    case 4:
        return "Jueves"
    case 5:
        return "Viernes"
    case 6:
        return "Sábado"
    case 7:
        return "Domingo"
    default:
        return "Día inválido"
    }
}
```

### 🌟 Características Únicas de Go

#### 1. No Necesita `break`

```go
func ejemplo() {
    valor := 2
    
    switch valor {
    case 1:
        fmt.Println("Uno")
        // No necesita break, automáticamente sale
    case 2:
        fmt.Println("Dos")
        // No continúa al siguiente case
    case 3:
        fmt.Println("Tres")
    }
}
```

#### 2. Múltiples Valores en un Case

```go
func tipoCaracter(c rune) string {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
        return "vocal"
    case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
        return "consonante"
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return "dígito"
    default:
        return "otro"
    }
}
```

#### 3. Switch con Expresiones

```go
func clasificarNota(nota int) string {
    switch {
    case nota >= 90:
        return "Excelente"
    case nota >= 80:
        return "Muy bueno"
    case nota >= 70:
        return "Bueno"
    case nota >= 60:
        return "Suficiente"
    default:
        return "Insuficiente"
    }
}
```

#### 4. Inicialización en Switch

```go
func procesarArchivo(nombre string) {
    switch ext := obtenerExtension(nombre); ext {
    case ".jpg", ".png", ".gif":
        procesarImagen(nombre)
    case ".mp4", ".avi", ".mov":
        procesarVideo(nombre)
    case ".txt", ".md":
        procesarTexto(nombre)
    default:
        fmt.Printf("Tipo de archivo no soportado: %s\n", ext)
    }
}
```

#### 5. Type Switch

```go
func procesarInterfaz(v interface{}) {
    switch valor := v.(type) {
    case int:
        fmt.Printf("Es un entero: %d\n", valor)
    case string:
        fmt.Printf("Es un string: %s\n", valor)
    case bool:
        fmt.Printf("Es un booleano: %t\n", valor)
    case []int:
        fmt.Printf("Es un slice de enteros con %d elementos\n", len(valor))
    default:
        fmt.Printf("Tipo desconocido: %T\n", valor)
    }
}
```

### Ejemplo Práctico: Sistema de Menús

```go
func mostrarMenu() {
    for {
        fmt.Println("\n=== SISTEMA DE GESTIÓN ===")
        fmt.Println("1. Crear usuario")
        fmt.Println("2. Listar usuarios")
        fmt.Println("3. Buscar usuario")
        fmt.Println("4. Eliminar usuario")
        fmt.Println("5. Salir")
        fmt.Print("Selecciona una opción: ")
        
        var opcion int
        fmt.Scanln(&opcion)
        
        switch opcion {
        case 1:
            crearUsuario()
        case 2:
            listarUsuarios()
        case 3:
            buscarUsuario()
        case 4:
            eliminarUsuario()
        case 5:
            fmt.Println("¡Hasta luego!")
            return
        default:
            fmt.Println("❌ Opción inválida. Intenta de nuevo.")
        }
    }
}
```

---

## ⚡ 4. Control de Flujo

### break y continue

#### `break` - Salir del Bucle

```go
func buscarNumero(numeros []int, objetivo int) int {
    for i, num := range numeros {
        if num == objetivo {
            fmt.Printf("¡Encontrado en la posición %d!\n", i)
            break // Sale inmediatamente del bucle
        }
    }
    return -1
}
```

#### `continue` - Saltar Iteración

```go
func procesarNumerosPares(numeros []int) {
    for _, num := range numeros {
        if num%2 != 0 {
            continue // Salta números impares
        }
        
        fmt.Printf("Procesando número par: %d\n", num)
        // Lógica para números pares
    }
}
```

#### Etiquetas con break y continue

```go
func buscarEnMatriz(matriz [][]int, objetivo int) (int, int) {
    exterior:
    for i, fila := range matriz {
        for j, valor := range fila {
            if valor == objetivo {
                fmt.Printf("Encontrado en [%d][%d]\n", i, j)
                break exterior // Sale de ambos bucles
            }
        }
    }
    return -1, -1
}
```

### defer - Diferir Ejecución

```go
func trabajarConArchivo(nombre string) error {
    archivo, err := os.Open(nombre)
    if err != nil {
        return err
    }
    defer archivo.Close() // Se ejecuta al final, sin importar cómo salga la función
    
    // Trabajar con el archivo
    datos, err := io.ReadAll(archivo)
    if err != nil {
        return err // archivo.Close() se ejecuta automáticamente
    }
    
    fmt.Printf("Leídos %d bytes\n", len(datos))
    return nil // archivo.Close() se ejecuta automáticamente
}
```

#### Múltiples defer (LIFO - Last In, First Out)

```go
func ejemploDefer() {
    fmt.Println("Inicio")
    
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")
    defer fmt.Println("Defer 3")
    
    fmt.Println("Medio")
    fmt.Println("Final")
}

// Salida:
// Inicio
// Medio
// Final
// Defer 3
// Defer 2
// Defer 1
```

### goto (Uso Muy Limitado)

```go
func ejemploGoto() {
    i := 0
    
    inicio:
    if i < 3 {
        fmt.Printf("Iteración %d\n", i)
        i++
        goto inicio
    }
    
    fmt.Println("Terminado")
}
```

> **⚠️ Advertencia**: `goto` debe usarse con **extrema moderación**. En la mayoría de casos, hay alternativas más limpias.

---

## 🛠️ 5. Patrones Comunes y Mejores Prácticas

### Patrón: Validación Múltiple

```go
func validarDatos(email, password string, edad int) error {
    if email == "" {
        return fmt.Errorf("el email es requerido")
    }
    
    if !strings.Contains(email, "@") {
        return fmt.Errorf("el email no es válido")
    }
    
    if len(password) < 8 {
        return fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
    }
    
    if edad < 18 {
        return fmt.Errorf("debes ser mayor de edad")
    }
    
    return nil
}
```

### Patrón: Procesamiento por Lotes

```go
func procesarLotes(datos []int, tamanoLote int) {
    for i := 0; i < len(datos); i += tamanoLote {
        fin := i + tamanoLote
        if fin > len(datos) {
            fin = len(datos)
        }
        
        lote := datos[i:fin]
        fmt.Printf("Procesando lote %d: %v\n", i/tamanoLote+1, lote)
        
        // Procesar el lote
        for _, item := range lote {
            procesarItem(item)
        }
    }
}
```

### Patrón: Retry con Backoff

```go
func intentarConexion(maxIntentos int) error {
    for intento := 1; intento <= maxIntentos; intento++ {
        err := conectar()
        if err == nil {
            fmt.Println("✅ Conexión exitosa")
            return nil
        }
        
        if intento == maxIntentos {
            return fmt.Errorf("falló después de %d intentos: %v", maxIntentos, err)
        }
        
        // Backoff exponencial
        tiempo := time.Duration(intento*intento) * time.Second
        fmt.Printf("⏳ Intento %d falló, reintentando en %v...\n", intento, tiempo)
        time.Sleep(tiempo)
    }
    
    return nil
}
```

### Patrón: Estado Máquina Simple

```go
type Estado int

const (
    Inicial Estado = iota
    Procesando
    Completado
    Error
)

func maquinaEstados(eventos []string) {
    estado := Inicial
    
    for _, evento := range eventos {
        switch estado {
        case Inicial:
            switch evento {
            case "iniciar":
                estado = Procesando
                fmt.Println("🚀 Iniciando procesamiento...")
            default:
                fmt.Printf("❌ Evento '%s' no válido en estado inicial\n", evento)
            }
            
        case Procesando:
            switch evento {
            case "completar":
                estado = Completado
                fmt.Println("✅ Procesamiento completado")
            case "error":
                estado = Error
                fmt.Println("💥 Error en procesamiento")
            default:
                fmt.Printf("❌ Evento '%s' no válido durante procesamiento\n", evento)
            }
            
        case Completado:
            switch evento {
            case "reiniciar":
                estado = Inicial
                fmt.Println("🔄 Reiniciando...")
            default:
                fmt.Printf("❌ Evento '%s' no válido cuando completado\n", evento)
            }
            
        case Error:
            switch evento {
            case "reiniciar":
                estado = Inicial
                fmt.Println("🔄 Reiniciando después del error...")
            default:
                fmt.Printf("❌ Evento '%s' no válido en estado de error\n", evento)
            }
        }
    }
}
```

---

## 🎯 6. Ejemplos Prácticos Integrados

### Sistema de Autenticación

```go
type Usuario struct {
    ID       int
    Nombre   string
    Email    string
    Password string
    Activo   bool
}

func autenticarUsuario(email, password string, usuarios []Usuario) *Usuario {
    for _, usuario := range usuarios {
        if !usuario.Activo {
            continue // Saltar usuarios inactivos
        }
        
        if usuario.Email == email {
            if verificarPassword(usuario.Password, password) {
                fmt.Printf("✅ Bienvenido, %s!\n", usuario.Nombre)
                return &usuario
            } else {
                fmt.Println("❌ Contraseña incorrecta")
                return nil
            }
        }
    }
    
    fmt.Println("❌ Usuario no encontrado")
    return nil
}

func gestionarSesion(usuario *Usuario) {
    if usuario == nil {
        fmt.Println("No hay usuario autenticado")
        return
    }
    
    for {
        fmt.Println("\n=== PANEL DE USUARIO ===")
        fmt.Println("1. Ver perfil")
        fmt.Println("2. Cambiar contraseña")
        fmt.Println("3. Cerrar sesión")
        
        var opcion int
        fmt.Print("Selecciona una opción: ")
        fmt.Scanln(&opcion)
        
        switch opcion {
        case 1:
            mostrarPerfil(usuario)
        case 2:
            cambiarPassword(usuario)
        case 3:
            fmt.Println("👋 Sesión cerrada")
            return
        default:
            fmt.Println("Opción inválida")
        }
    }
}
```

### Analizador de Logs

```go
func analizarLogs(logs []string) {
    estadisticas := map[string]int{
        "INFO":    0,
        "WARNING": 0,
        "ERROR":   0,
        "DEBUG":   0,
    }
    
    var errores []string
    
    for numeroLinea, linea := range logs {
        linea = strings.TrimSpace(linea)
        
        if linea == "" {
            continue // Saltar líneas vacías
        }
        
        // Determinar tipo de log
        switch {
        case strings.Contains(linea, "[INFO]"):
            estadisticas["INFO"]++
        case strings.Contains(linea, "[WARNING]"):
            estadisticas["WARNING"]++
        case strings.Contains(linea, "[ERROR]"):
            estadisticas["ERROR"]++
            errores = append(errores, fmt.Sprintf("Línea %d: %s", numeroLinea+1, linea))
        case strings.Contains(linea, "[DEBUG]"):
            estadisticas["DEBUG"]++
        default:
            fmt.Printf("⚠️ Línea %d: Formato desconocido\n", numeroLinea+1)
        }
    }
    
    // Mostrar estadísticas
    fmt.Println("\n📊 ESTADÍSTICAS DE LOGS:")
    for tipo, cantidad := range estadisticas {
        fmt.Printf("   %s: %d\n", tipo, cantidad)
    }
    
    // Mostrar errores si los hay
    if len(errores) > 0 {
        fmt.Println("\n🚨 ERRORES ENCONTRADOS:")
        for _, error := range errores {
            fmt.Printf("   %s\n", error)
        }
    }
}
```

### Juego de Adivinanza

```go
func juegoAdivinanza() {
    rand.Seed(time.Now().UnixNano())
    numeroSecreto := rand.Intn(100) + 1
    intentos := 0
    maxIntentos := 7
    
    fmt.Println("🎲 ¡Adivina el número entre 1 y 100!")
    fmt.Printf("Tienes %d intentos.\n", maxIntentos)
    
    for intentos < maxIntentos {
        intentos++
        fmt.Printf("\nIntento %d/%d: ", intentos, maxIntentos)
        
        var numero int
        _, err := fmt.Scanln(&numero)
        if err != nil {
            fmt.Println("❌ Por favor ingresa un número válido")
            intentos-- // No contar este intento
            continue
        }
        
        switch {
        case numero < 1 || numero > 100:
            fmt.Println("❌ El número debe estar entre 1 y 100")
            intentos-- // No contar este intento
        case numero == numeroSecreto:
            fmt.Printf("🎉 ¡FELICIDADES! Adivinaste en %d intentos\n", intentos)
            return
        case numero < numeroSecreto:
            fmt.Println("📈 El número es mayor")
        case numero > numeroSecreto:
            fmt.Println("📉 El número es menor")
        }
        
        // Dar pistas adicionales
        if intentos == maxIntentos-1 {
            fmt.Printf("💡 Pista: El número %s par\n", 
                map[bool]string{true: "es", false: "no es"}[numeroSecreto%2 == 0])
        }
    }
    
    fmt.Printf("😞 Se acabaron los intentos. El número era %d\n", numeroSecreto)
}
```

---

## 💡 7. Consejos y Mejores Prácticas

### ✅ Buenas Prácticas

1. **Usa `range` para iterar colecciones**
```go
// ✅ Bueno
for i, valor := range slice {
    // ...
}

// ❌ Evitar
for i := 0; i < len(slice); i++ {
    valor := slice[i]
    // ...
}
```

2. **Inicializa variables en if/switch cuando sea apropiado**
```go
// ✅ Bueno
if user, found := users[id]; found {
    return user.Name
}

// ❌ Menos limpio
user, found := users[id]
if found {
    return user.Name
}
```

3. **Usa `switch` sin expresión para condiciones complejas**
```go
// ✅ Bueno
switch {
case x > 10 && y < 5:
    // ...
case x < 0:
    // ...
}

// ❌ Menos legible
if x > 10 && y < 5 {
    // ...
} else if x < 0 {
    // ...
}
```

4. **Usa `defer` para cleanup**
```go
// ✅ Bueno
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Siempre se ejecuta
    
    // proceso del archivo...
    return nil
}
```

### ❌ Anti-patrones a Evitar

1. **No usar `goto` excepto en casos muy específicos**
2. **No anidar demasiado profundo (máximo 3-4 niveles)**
3. **No usar `break` y `continue` excesivamente**
4. **No ignorar errores en bucles**

```go
// ❌ Malo
for _, item := range items {
    process(item) // ¿Qué pasa si falla?
}

// ✅ Mejor
for _, item := range items {
    if err := process(item); err != nil {
        log.Printf("Error procesando %v: %v", item, err)
        continue
    }
}
```

---

## 🧪 8. Laboratorio: Sistema de Gestión de Tareas

Vamos a construir un sistema completo que integre todas las estructuras de control:

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

type Tarea struct {
    ID          int
    Titulo      string
    Descripcion string
    Completada  bool
    Prioridad   int // 1=Baja, 2=Media, 3=Alta
    FechaLimite time.Time
}

type SistemaTareas struct {
    tareas   []Tarea
    siguienteID int
}

func (st *SistemaTareas) menu() {
    for {
        fmt.Println("\n📋 === GESTOR DE TAREAS ===")
        fmt.Println("1. Agregar tarea")
        fmt.Println("2. Listar tareas")
        fmt.Println("3. Marcar como completada")
        fmt.Println("4. Buscar tareas")
        fmt.Println("5. Estadísticas")
        fmt.Println("6. Limpiar completadas")
        fmt.Println("7. Salir")
        fmt.Print("Selecciona una opción: ")
        
        var opcion int
        if _, err := fmt.Scanln(&opcion); err != nil {
            fmt.Println("❌ Entrada inválida")
            continue
        }
        
        switch opcion {
        case 1:
            st.agregarTarea()
        case 2:
            st.listarTareas()
        case 3:
            st.completarTarea()
        case 4:
            st.buscarTareas()
        case 5:
            st.mostrarEstadisticas()
        case 6:
            st.limpiarCompletadas()
        case 7:
            fmt.Println("👋 ¡Hasta luego!")
            return
        default:
            fmt.Println("❌ Opción inválida")
        }
    }
}

func (st *SistemaTareas) agregarTarea() {
    fmt.Print("Título de la tarea: ")
    var titulo string
    fmt.Scanln(&titulo)
    
    if strings.TrimSpace(titulo) == "" {
        fmt.Println("❌ El título no puede estar vacío")
        return
    }
    
    fmt.Print("Descripción: ")
    var descripcion string
    fmt.Scanln(&descripcion)
    
    fmt.Print("Prioridad (1=Baja, 2=Media, 3=Alta): ")
    var prioridad int
    fmt.Scanln(&prioridad)
    
    if prioridad < 1 || prioridad > 3 {
        prioridad = 2 // Prioridad media por defecto
    }
    
    st.siguienteID++
    tarea := Tarea{
        ID:          st.siguienteID,
        Titulo:      titulo,
        Descripcion: descripcion,
        Completada:  false,
        Prioridad:   prioridad,
        FechaLimite: time.Now().AddDate(0, 0, 7), // 7 días desde hoy
    }
    
    st.tareas = append(st.tareas, tarea)
    fmt.Printf("✅ Tarea '%s' agregada con ID %d\n", titulo, st.siguienteID)
}

func (st *SistemaTareas) listarTareas() {
    if len(st.tareas) == 0 {
        fmt.Println("📝 No hay tareas registradas")
        return
    }
    
    fmt.Println("\n📋 LISTA DE TAREAS:")
    for i, tarea := range st.tareas {
        estado := "❌"
        if tarea.Completada {
            estado = "✅"
        }
        
        prioridad := ""
        switch tarea.Prioridad {
        case 1:
            prioridad = "🟢 Baja"
        case 2:
            prioridad = "🟡 Media"
        case 3:
            prioridad = "🔴 Alta"
        }
        
        fmt.Printf("%d. %s [ID:%d] %s - %s\n", 
            i+1, estado, tarea.ID, tarea.Titulo, prioridad)
        
        if tarea.Descripcion != "" {
            fmt.Printf("   📝 %s\n", tarea.Descripcion)
        }
    }
}

func (st *SistemaTareas) completarTarea() {
    fmt.Print("ID de la tarea a completar: ")
    var id int
    fmt.Scanln(&id)
    
    for i := range st.tareas {
        if st.tareas[i].ID == id {
            if st.tareas[i].Completada {
                fmt.Println("ℹ️ La tarea ya estaba completada")
            } else {
                st.tareas[i].Completada = true
                fmt.Printf("✅ Tarea '%s' marcada como completada\n", st.tareas[i].Titulo)
            }
            return
        }
    }
    
    fmt.Println("❌ No se encontró tarea con ese ID")
}

func (st *SistemaTareas) buscarTareas() {
    fmt.Print("Término de búsqueda: ")
    var termino string
    fmt.Scanln(&termino)
    
    termino = strings.ToLower(strings.TrimSpace(termino))
    if termino == "" {
        fmt.Println("❌ Término de búsqueda vacío")
        return
    }
    
    var encontradas []Tarea
    
    for _, tarea := range st.tareas {
        titulo := strings.ToLower(tarea.Titulo)
        descripcion := strings.ToLower(tarea.Descripcion)
        
        if strings.Contains(titulo, termino) || strings.Contains(descripcion, termino) {
            encontradas = append(encontradas, tarea)
        }
    }
    
    if len(encontradas) == 0 {
        fmt.Printf("🔍 No se encontraron tareas con el término '%s'\n", termino)
        return
    }
    
    fmt.Printf("🔍 Encontradas %d tarea(s):\n", len(encontradas))
    for _, tarea := range encontradas {
        estado := "❌"
        if tarea.Completada {
            estado = "✅"
        }
        fmt.Printf("   %s [ID:%d] %s\n", estado, tarea.ID, tarea.Titulo)
    }
}

func (st *SistemaTareas) mostrarEstadisticas() {
    if len(st.tareas) == 0 {
        fmt.Println("📊 No hay tareas para analizar")
        return
    }
    
    var completadas, pendientes int
    prioridades := map[int]int{1: 0, 2: 0, 3: 0}
    
    for _, tarea := range st.tareas {
        if tarea.Completada {
            completadas++
        } else {
            pendientes++
        }
        prioridades[tarea.Prioridad]++
    }
    
    porcentaje := float64(completadas) / float64(len(st.tareas)) * 100
    
    fmt.Println("\n📊 ESTADÍSTICAS:")
    fmt.Printf("   Total: %d tareas\n", len(st.tareas))
    fmt.Printf("   Completadas: %d (%.1f%%)\n", completadas, porcentaje)
    fmt.Printf("   Pendientes: %d\n", pendientes)
    fmt.Println("\n📈 Por prioridad:")
    fmt.Printf("   🟢 Baja: %d\n", prioridades[1])
    fmt.Printf("   🟡 Media: %d\n", prioridades[2])
    fmt.Printf("   🔴 Alta: %d\n", prioridades[3])
}

func (st *SistemaTareas) limpiarCompletadas() {
    var tareasActivas []Tarea
    eliminadas := 0
    
    for _, tarea := range st.tareas {
        if !tarea.Completada {
            tareasActivas = append(tareasActivas, tarea)
        } else {
            eliminadas++
        }
    }
    
    st.tareas = tareasActivas
    
    if eliminadas > 0 {
        fmt.Printf("🗑️ Se eliminaron %d tarea(s) completada(s)\n", eliminadas)
    } else {
        fmt.Println("ℹ️ No hay tareas completadas para eliminar")
    }
}

func main() {
    sistema := &SistemaTareas{
        tareas:      make([]Tarea, 0),
        siguienteID: 0,
    }
    
    // Datos de ejemplo
    sistema.tareas = []Tarea{
        {1, "Estudiar Go", "Completar lección de estructuras de control", false, 3, time.Now().AddDate(0, 0, 2)},
        {2, "Ejercicio", "Salir a correr 30 minutos", true, 2, time.Now().AddDate(0, 0, 1)},
        {3, "Proyecto", "Terminar aplicación web", false, 3, time.Now().AddDate(0, 0, 10)},
    }
    sistema.siguienteID = 3
    
    fmt.Println("🚀 Bienvenido al Sistema de Gestión de Tareas")
    sistema.menu()
}
```

---

## 📚 9. Resumen de Conceptos

### Estructuras de Control en Go

| Estructura | Sintaxis | Uso |
|------------|----------|-----|
| `if` | `if condición { ... }` | Ejecución condicional |
| `for` | `for init; cond; post { ... }` | Bucles |
| `for range` | `for k, v := range col { ... }` | Iterar colecciones |
| `switch` | `switch val { case x: ... }` | Múltiples condiciones |
| `type switch` | `switch v.(type) { ... }` | Verificar tipos |
| `break` | `break` | Salir de bucle |
| `continue` | `continue` | Siguiente iteración |
| `defer` | `defer func()` | Ejecución diferida |

### 🎯 Puntos Clave

1. **Go solo tiene `for`** - pero es muy versátil
2. **`switch` no necesita `break`** - no hay fall-through por defecto
3. **`defer` es poderoso** - siempre se ejecuta
4. **Inicialización en if/switch** - mantiene scope limitado
5. **`range` es idiomático** - úsalo para colecciones

---

## 🎯 10. Ejercicios Prácticos

### Ejercicio 1: Validador de Números
Crea un programa que:
- Lea números del usuario hasta que ingrese 0
- Valide que sean números positivos
- Calcule estadísticas (suma, promedio, máximo, mínimo)

### Ejercicio 2: Juego de Piedra, Papel, Tijera
Implementa el juego clásico con:
- Menú de opciones
- Contador de victorias
- Opción de jugar múltiples rondas

### Ejercicio 3: Analizador de Texto
Programa que analice un texto y muestre:
- Número de caracteres, palabras, líneas
- Frecuencia de cada letra
- Palabras más comunes

### Ejercicio 4: Sistema de Calificaciones
Gestor que permita:
- Agregar estudiantes y calificaciones
- Calcular promedios
- Mostrar estadísticas por materia
- Generar reportes

### Ejercicio 5: Conversor de Unidades
Aplicación que convierta entre diferentes unidades:
- Longitud (metros, pies, pulgadas)
- Peso (kg, libras, onzas)
- Temperatura (Celsius, Fahrenheit, Kelvin)

---

## ✅ Checklist de Dominio

Antes de continuar a la siguiente lección, asegúrate de poder:

- [ ] Escribir condicionales `if-else` complejas
- [ ] Usar todas las variantes del bucle `for`
- [ ] Implementar `switch` statements efectivamente
- [ ] Usar `break`, `continue` y `defer` apropiadamente
- [ ] Escribir código limpio con estructuras de control
- [ ] Debuggear problemas de lógica de control
- [ ] Optimizar bucles para performance
- [ ] Usar patrones idiomáticos de Go

---

## 🔗 Navegación

⬅️ **Anterior**: [Lección 6: Operadores](../06-operadores/README.md)  
➡️ **Siguiente**: [Lección 8: Funciones](../08-funciones/README.md)  
🏠 **Inicio**: [Fundamentos de Go](../README.md)  
📚 **Curso**: [Go Deep - Domina Go](../../README.md)

---

## 📞 Soporte

¿Tienes dudas sobre estructuras de control? 

- 💬 **Discusión**: [GitHub Discussions](../../discussions)
- 🐛 **Problemas**: [GitHub Issues](../../issues)
- 📧 **Email**: [contacto@go-deep.dev](mailto:contacto@go-deep.dev)

---

¡Las estructuras de control son la base de la lógica de programación! 🎯 Practica mucho y experimenta con diferentes patrones. 

**¡Continuemos construyendo tus habilidades en Go!** 🚀
