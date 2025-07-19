# ⚡ Operadores: El Poder de las Operaciones

> *"Operators are the verbs of programming; they make data dance"* - Programming Wisdom

Los operadores son las herramientas que transforman, comparan y manipulan datos. En esta lección dominarás todos los operadores de Go y aprenderás a usarlos de forma idiomática y eficiente.

## 🎯 Objetivos de Esta Lección

Al finalizar esta lección serás capaz de:
- ✅ **Dominar todos los operadores** aritméticos, lógicos y bitwise
- ✅ **Usar operadores de comparación** de forma idiomática
- ✅ **Aplicar precedencia** correctamente sin confusiones
- ✅ **Evitar errores comunes** con operadores
- ✅ **Optimizar código** usando operadores apropiados
- ✅ **Implementar algoritmos** eficientes con operaciones bitwise

---

## 🧮 Operadores Aritméticos: Las Matemáticas de Go

### 📝 Operadores Básicos

```go
package main

import "fmt"

func operadoresBasicos() {
    fmt.Println("=== Operadores Aritméticos Básicos ===")
    
    a, b := 15, 4
    
    // Operadores básicos
    fmt.Printf("a = %d, b = %d\n", a, b)
    fmt.Printf("Suma:        a + b = %d\n", a+b)      // 19
    fmt.Printf("Resta:       a - b = %d\n", a-b)      // 11
    fmt.Printf("Multiplicación: a * b = %d\n", a*b)   // 60
    fmt.Printf("División:    a / b = %d\n", a/b)      // 3 (división entera)
    fmt.Printf("Módulo:      a %% b = %d\n", a%b)     // 3 (resto)
    
    // División con flotantes
    fmt.Printf("\nDivisión con flotantes:\n")
    fmt.Printf("float64(a) / float64(b) = %.2f\n", float64(a)/float64(b)) // 3.75
    
    // Operadores unarios
    x := 10
    fmt.Printf("\nOperadores unarios:\n")
    fmt.Printf("x = %d\n", x)
    fmt.Printf("+x = %d\n", +x)    // 10 (positivo explícito)
    fmt.Printf("-x = %d\n", -x)    // -10 (negativo)
}

// Demostración de precedencia aritmética
func precedenciaAritmetica() {
    fmt.Println("\n=== Precedencia Aritmética ===")
    
    // Sin paréntesis
    result1 := 2 + 3 * 4      // 14 (no 20)
    result2 := 10 - 6 / 2     // 7 (no 2)
    result3 := 5 * 2 + 3      // 13
    
    fmt.Printf("2 + 3 * 4 = %d (multiplicación primero)\n", result1)
    fmt.Printf("10 - 6 / 2 = %d (división primero)\n", result2)
    fmt.Printf("5 * 2 + 3 = %d (multiplicación primero)\n", result3)
    
    // Con paréntesis para claridad
    result4 := (2 + 3) * 4    // 20
    result5 := (10 - 6) / 2   // 2
    result6 := 5 * (2 + 3)    // 25
    
    fmt.Printf("(2 + 3) * 4 = %d\n", result4)
    fmt.Printf("(10 - 6) / 2 = %d\n", result5)
    fmt.Printf("5 * (2 + 3) = %d\n", result6)
}

// Demostración de overflow y underflow
func overflowDemo() {
    fmt.Println("\n=== Overflow y Underflow ===")
    
    // Overflow con int8
    var small int8 = 127  // Valor máximo para int8
    fmt.Printf("int8 máximo: %d\n", small)
    
    // Esto causaría overflow en compilación
    // small = 128  // Error: constant 128 overflows int8
    
    // Overflow en runtime (cuidado)
    small += 1  // Esto compila pero puede dar comportamiento inesperado
    fmt.Printf("Después de +1: %d (¡overflow!)\n", small)
    
    // Underflow con uint
    var unsigned uint = 0
    fmt.Printf("uint mínimo: %d\n", unsigned)
    
    // Esto causaría underflow
    // unsigned -= 1  // Behavior depends on implementation
    
    // ✅ Verificación segura
    if unsigned > 0 {
        unsigned -= 1
    } else {
        fmt.Println("⚠️ Underflow evitado")
    }
}
```
}
```

### 🧠 Analogía: Operadores como Herramientas

Imagina los operadores como **herramientas en un taller**:

```
🔧 + (suma)         → Soldadora (une piezas)
✂️ - (resta)        → Sierra (corta/separa)
🔨 * (multiplicación) → Martillo (amplifica fuerza)
📏 / (división)     → Regla (mide/divide)
🗑️ % (módulo)       → Filtro (separa resto)
```

### ⚠️ Trampas Comunes con División

```go
package main

import "fmt"

func trampasDivision() {
    fmt.Println("=== Trampas Comunes con División ===")
    
    // ❌ División entera inesperada
    resultado1 := 5 / 2
    fmt.Printf("5 / 2 = %d (¡Sorpresa! División entera)\n", resultado1) // 2
    
    // ✅ División con flotantes
    resultado2 := 5.0 / 2.0
    fmt.Printf("5.0 / 2.0 = %.1f\n", resultado2) // 2.5
    
    // ✅ Conversión explícita
    a, b := 5, 2
    resultado3 := float64(a) / float64(b)
    fmt.Printf("float64(%d) / float64(%d) = %.1f\n", a, b, resultado3) // 2.5
    
    // ⚠️ División por cero en runtime
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("¡Pánico capturado!: %v\n", r)
        }
    }()
    
    divisor := 0
    // fmt.Printf("10 / 0 = %d\n", 10/divisor) // ¡Pánico!
    
    // ✅ Verificación segura
    if divisor != 0 {
        fmt.Printf("10 / %d = %d\n", divisor, 10/divisor)
    } else {
        fmt.Println("⚠️ División por cero evitada")
    }
}
```

### 🎯 Operadores de Asignación

```go
package main

import "fmt"

func operadoresAsignacion() {
    fmt.Println("=== Operadores de Asignación ===")
    
    // Asignación simple
    x := 10
    fmt.Printf("Inicial: x = %d\n", x)
    
    // Operadores compuestos
    x += 5    // x = x + 5
    fmt.Printf("x += 5:  x = %d\n", x)  // 15
    
    x -= 3    // x = x - 3
    fmt.Printf("x -= 3:  x = %d\n", x)  // 12
    
    x *= 2    // x = x * 2
    fmt.Printf("x *= 2:  x = %d\n", x)  // 24
    
    x /= 4    // x = x / 4
    fmt.Printf("x /= 4:  x = %d\n", x)  // 6
    
    x %= 4    // x = x % 4
    fmt.Printf("x %%= 4: x = %d\n", x)  // 2
    
    // Incremento y decremento
    fmt.Println("\nIncremento y decremento:")
    y := 5
    fmt.Printf("y inicial = %d\n", y)
    
    y++  // Post-incremento (solo esta forma en Go)
    fmt.Printf("y++ = %d\n", y)  // 6
    
    y--  // Post-decremento
    fmt.Printf("y-- = %d\n", y)  // 5
    
    // ❌ Go NO tiene pre-incremento
    // ++y  // Error de compilación
    // --y  // Error de compilación
}
```

---

## 🔍 Operadores de Comparación: Evaluando la Verdad

### 📊 Operadores de Igualdad y Relacionales

```go
package main

import "fmt"

func operadoresComparacion() {
    fmt.Println("=== Operadores de Comparación ===")
    
    a, b := 10, 20
    fmt.Printf("a = %d, b = %d\n", a, b)
    
    // Operadores de igualdad
    fmt.Printf("a == b: %t\n", a == b)  // false
    fmt.Printf("a != b: %t\n", a != b)  // true
    
    // Operadores relacionales
    fmt.Printf("a < b:  %t\n", a < b)   // true
    fmt.Printf("a <= b: %t\n", a <= b)  // true
    fmt.Printf("a > b:  %t\n", a > b)   // false
    fmt.Printf("a >= b: %t\n", a >= b)  // false
    
    // Comparación de strings
    fmt.Println("\nComparación de strings:")
    str1, str2 := "apple", "banana"
    fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)  // true (orden lexicográfico)
    
    // Comparación de tipos compatibles
    var x int = 42
    var y int64 = 42
    // fmt.Printf("x == y: %t\n", x == y)  // ❌ Error: tipos diferentes
    fmt.Printf("x == int(y): %t\n", x == int(y))  // ✅ true
    fmt.Printf("int64(x) == y: %t\n", int64(x) == y)  // ✅ true
}

// Comparación de flotantes - ¡Cuidado!
func comparacionFlotantes() {
    fmt.Println("\n=== Comparación de Flotantes (¡Peligroso!) ===")
    
    // ❌ Comparación directa de flotantes puede fallar
    a := 0.1 + 0.2
    b := 0.3
    fmt.Printf("0.1 + 0.2 = %.17f\n", a)
    fmt.Printf("0.3 = %.17f\n", b)
    fmt.Printf("¿Son iguales? %t\n", a == b)  // ¡Puede ser false!
    
    // ✅ Comparación con epsilon
    epsilon := 1e-9
    diff := a - b
    if diff < 0 {
        diff = -diff  // Valor absoluto
    }
    isEqual := diff < epsilon
    fmt.Printf("¿Son iguales con epsilon? %t\n", isEqual)
    
    // ✅ Función helper para comparar flotantes
    floatEqual := func(a, b, epsilon float64) bool {
        diff := a - b
        if diff < 0 {
            diff = -diff
        }
        return diff < epsilon
    }
    
    fmt.Printf("Usando función helper: %t\n", floatEqual(a, b, 1e-9))
}
```

---

## 🔗 Operadores Lógicos: La Lógica Booleana

### 🧠 AND, OR, NOT - Los Fundamentos

```go
package main

import "fmt"

func operadoresLogicos() {
    fmt.Println("=== Operadores Lógicos ===")
    
    a, b := true, false
    fmt.Printf("a = %t, b = %t\n", a, b)
    
    // Operador AND (&&)
    fmt.Printf("a && b = %t\n", a && b)  // false
    fmt.Printf("a && true = %t\n", a && true)  // true
    
    // Operador OR (||)
    fmt.Printf("a || b = %t\n", a || b)  // true
    fmt.Printf("false || b = %t\n", false || b)  // false
    
    // Operador NOT (!)
    fmt.Printf("!a = %t\n", !a)  // false
    fmt.Printf("!b = %t\n", !b)  // true
    
    // Combinaciones complejas
    fmt.Printf("!(a && b) = %t\n", !(a && b))  // true
    fmt.Printf("!a || !b = %t\n", !a || !b)    // true (De Morgan)
}

// Demostración de cortocircuito
func cortocircuito() {
    fmt.Println("\n=== Evaluación de Cortocircuito ===")
    
    // Función que imprime y retorna
    check := func(name string, value bool) bool {
        fmt.Printf("  Evaluando %s: %t\n", name, value)
        return value
    }
    
    fmt.Println("AND con cortocircuito:")
    result1 := check("false", false) && check("true", true)
    fmt.Printf("Resultado: %t\n", result1)
    // Solo evalúa el primer operando (false)
    
    fmt.Println("\nOR con cortocircuito:")
    result2 := check("true", true) || check("false", false)
    fmt.Printf("Resultado: %t\n", result2)
    // Solo evalúa el primer operando (true)
    
    fmt.Println("\nSin cortocircuito:")
    result3 := check("false", false) && check("true", true)
    fmt.Printf("Resultado: %t\n", result3)
}

// Patrones útiles con operadores lógicos
func patronesLogicos() {
    fmt.Println("\n=== Patrones Útiles ===")
    
    // Validación de rangos
    age := 25
    isValidAge := age >= 18 && age <= 65
    fmt.Printf("Edad %d es válida: %t\n", age, isValidAge)
    
    // Valores por defecto
    name := ""
    displayName := name
    if name == "" {
        displayName = "Usuario Anónimo"
    }
    // Forma más idiomática:
    // displayName := name != "" && name || "Usuario Anónimo"  // No funciona así en Go
    
    fmt.Printf("Nombre a mostrar: %s\n", displayName)
    
    // Verificación de múltiples condiciones
    hour := 14
    isWorkingHours := hour >= 9 && hour <= 17
    isWeekend := false  // Simplificado
    isAvailable := isWorkingHours && !isWeekend
    
    fmt.Printf("Disponible: %t\n", isAvailable)
}
```

---

## 🔢 Operadores Bitwise: Manipulando Bits

### ⚡ Operaciones a Nivel de Bit

```go
package main

import "fmt"

func operadoresBitwise() {
    fmt.Println("=== Operadores Bitwise ===")
    
    a, b := 12, 10  // 1100 y 1010 en binario
    fmt.Printf("a = %d (binario: %08b)\n", a, a)
    fmt.Printf("b = %d (binario: %08b)\n", b, b)
    
    // AND bitwise (&)
    and := a & b
    fmt.Printf("a & b  = %d (binario: %08b)\n", and, and)  // 8 (1000)
    
    // OR bitwise (|)
    or := a | b
    fmt.Printf("a | b  = %d (binario: %08b)\n", or, or)    // 14 (1110)
    
    // XOR bitwise (^)
    xor := a ^ b
    fmt.Printf("a ^ b  = %d (binario: %08b)\n", xor, xor)  // 6 (0110)
    
    // NOT bitwise (^) - unario
    not := ^a
    fmt.Printf("^a     = %d (binario: %032b)\n", not, not)  // Complemento
    
    // Desplazamientos
    fmt.Println("\nDesplazamientos:")
    left := a << 2   // Izquierda: multiplica por 2^n
    right := a >> 1  // Derecha: divide por 2^n
    
    fmt.Printf("a << 2 = %d (binario: %08b) // %d * 4\n", left, left, a)
    fmt.Printf("a >> 1 = %d (binario: %08b) // %d / 2\n", right, right, a)
}

// Casos de uso prácticos con bitwise
func casosPracticosBitwise() {
    fmt.Println("\n=== Casos Prácticos Bitwise ===")
    
    // 1. Verificar si un número es par
    num := 42
    isPar := (num & 1) == 0
    fmt.Printf("%d es par: %t\n", num, isPar)
    
    // 2. Potencias de 2 rápidas
    fmt.Println("\nPotencias de 2:")
    for i := 0; i < 5; i++ {
        power := 1 << i  // 2^i
        fmt.Printf("2^%d = %d\n", i, power)
    }
    
    // 3. Intercambio sin variable temporal (XOR swap)
    x, y := 25, 30
    fmt.Printf("Antes: x=%d, y=%d\n", x, y)
    x = x ^ y
    y = x ^ y
    x = x ^ y
    fmt.Printf("Después: x=%d, y=%d\n", x, y)
    
    // 4. Contar bits activados (población)
    value := 23  // 10111 en binario
    count := 0
    temp := value
    for temp != 0 {
        count += temp & 1
        temp >>= 1
    }
    fmt.Printf("%d (%08b) tiene %d bits activados\n", value, value, count)
}

// Sistema de flags/permisos con bitwise
type Permission uint8

const (
    Read    Permission = 1 << iota  // 1 (00000001)
    Write                           // 2 (00000010)
    Execute                         // 4 (00000100)
    Delete                          // 8 (00001000)
)

func sistemaPermisos() {
    fmt.Println("\n=== Sistema de Permisos con Bitwise ===")
    
    // Crear permisos
    userPerms := Read | Write                    // 3 (00000011)
    adminPerms := Read | Write | Execute | Delete // 15 (00001111)
    
    fmt.Printf("Usuario: %08b (%d)\n", userPerms, userPerms)
    fmt.Printf("Admin:   %08b (%d)\n", adminPerms, adminPerms)
    
    // Verificar permisos
    hasRead := userPerms&Read != 0
    hasDelete := userPerms&Delete != 0
    
    fmt.Printf("Usuario puede leer: %t\n", hasRead)
    fmt.Printf("Usuario puede eliminar: %t\n", hasDelete)
    
    // Agregar permiso
    userPerms |= Execute
    fmt.Printf("Usuario con Execute: %08b (%d)\n", userPerms, userPerms)
    
    // Quitar permiso
    userPerms &^= Write  // AND NOT
    fmt.Printf("Usuario sin Write: %08b (%d)\n", userPerms, userPerms)
    
    // Toggle permiso
    userPerms ^= Read
    fmt.Printf("Usuario toggle Read: %08b (%d)\n", userPerms, userPerms)
}
```

---

## 📐 Precedencia de Operadores: El Orden Importa

### 🎯 Tabla de Precedencia

```go
package main

import "fmt"

func precedenciaCompleta() {
    fmt.Println("=== Precedencia de Operadores ===")
    
    // Tabla de precedencia (de mayor a menor):
    // 1. * / % << >> & &^
    // 2. + - | ^
    // 3. == != < <= > >=
    // 4. &&
    // 5. ||
    
    // Ejemplos de precedencia
    fmt.Println("Sin paréntesis:")
    result1 := 2 + 3*4 == 14    // true: (2 + (3*4)) == 14
    result2 := 10 > 5 && 3 < 7  // true: (10 > 5) && (3 < 7)
    result3 := 1 << 2 + 1       // 8: 1 << (2 + 1) = 1 << 3
    
    fmt.Printf("2 + 3*4 == 14: %t\n", result1)
    fmt.Printf("10 > 5 && 3 < 7: %t\n", result2)
    fmt.Printf("1 << 2 + 1: %d\n", result3)
    
    // Con paréntesis para claridad
    fmt.Println("\nCon paréntesis explícitos:")
    result4 := (2 + 3) * 4 == 14  // false: 20 != 14
    result5 := 1 << (2 + 1)       // 8: mismo resultado pero más claro
    
    fmt.Printf("(2 + 3) * 4 == 14: %t\n", result4)
    fmt.Printf("1 << (2 + 1): %d\n", result5)
}

// Casos complejos de precedencia
func casosComplejos() {
    fmt.Println("\n=== Casos Complejos ===")
    
    // Mezcla de operadores
    a, b, c := 2, 3, 4
    
    // Sin paréntesis - siguiendo precedencia
    result1 := a + b*c > 10 && a < 5
    // Evaluación: a + (b*c) > 10 && a < 5
    //            2 + (3*4) > 10 && 2 < 5
    //            2 + 12 > 10 && 2 < 5
    //            14 > 10 && 2 < 5
    //            true && true
    //            true
    
    fmt.Printf("a + b*c > 10 && a < 5: %t\n", result1)
    
    // Con paréntesis para cambiar orden
    result2 := (a + b) * c > 10 && a < 5
    // Evaluación: (2 + 3) * 4 > 10 && 2 < 5
    //            5 * 4 > 10 && 2 < 5
    //            20 > 10 && 2 < 5
    //            true && true
    //            true
    
    fmt.Printf("(a + b) * c > 10 && a < 5: %t\n", result2)
    
    // Bitwise con aritmética
    x := 5
    result3 := x << 1 + 1  // x << (1 + 1) = 5 << 2 = 20
    result4 := (x << 1) + 1  // (5 << 1) + 1 = 10 + 1 = 11
    
    fmt.Printf("x << 1 + 1: %d\n", result3)
    fmt.Printf("(x << 1) + 1: %d\n", result4)
}
```

---

## 🚀 Proyecto: Calculadora Avanzada

### 🎯 Sistema Completo con Todos los Operadores

```go
package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

// Calculadora avanzada que demuestra todos los operadores
type Calculadora struct {
    memoria     float64
    historial   []string
    precision   int
}

func NewCalculadora() *Calculadora {
    return &Calculadora{
        memoria:   0,
        historial: make([]string, 0),
        precision: 2,
    }
}

// Operaciones aritméticas básicas
func (c *Calculadora) Sumar(a, b float64) float64 {
    result := a + b
    c.agregarHistorial(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculadora) Restar(a, b float64) float64 {
    result := a - b
    c.agregarHistorial(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculadora) Multiplicar(a, b float64) float64 {
    result := a * b
    c.agregarHistorial(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculadora) Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    result := a / b
    c.agregarHistorial(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculadora) Modulo(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("módulo por cero")
    }
    result := a % b
    c.agregarHistorial(fmt.Sprintf("%d %% %d = %d", a, b, result))
    return result, nil
}

// Operaciones de potencia
func (c *Calculadora) Potencia(base, exponente float64) float64 {
    result := math.Pow(base, exponente)
    c.agregarHistorial(fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponente, result))
    return result
}

// Operaciones bitwise (solo para enteros)
func (c *Calculadora) AND(a, b int) int {
    result := a & b
    c.agregarHistorial(fmt.Sprintf("%d & %d = %d (%08b & %08b = %08b)", 
        a, b, result, a, b, result))
    return result
}

func (c *Calculadora) OR(a, b int) int {
    result := a | b
    c.agregarHistorial(fmt.Sprintf("%d | %d = %d (%08b | %08b = %08b)", 
        a, b, result, a, b, result))
    return result
}

func (c *Calculadora) XOR(a, b int) int {
    result := a ^ b
    c.agregarHistorial(fmt.Sprintf("%d ^ %d = %d (%08b ^ %08b = %08b)", 
        a, b, result, a, b, result))
    return result
}

func (c *Calculadora) DesplazarIzquierda(a, n int) int {
    result := a << n
    c.agregarHistorial(fmt.Sprintf("%d << %d = %d (%08b << %d = %08b)", 
        a, n, result, a, n, result))
    return result
}

func (c *Calculadora) DesplazarDerecha(a, n int) int {
    result := a >> n
    c.agregarHistorial(fmt.Sprintf("%d >> %d = %d (%08b >> %d = %08b)", 
        a, n, result, a, n, result))
    return result
}

// Operaciones de comparación
func (c *Calculadora) Comparar(a, b float64) map[string]bool {
    result := map[string]bool{
        "igual":     c.sonIguales(a, b),
        "mayor":     a > b,
        "menor":     a < b,
        "mayor_igual": a >= b,
        "menor_igual": a <= b,
        "diferente": !c.sonIguales(a, b),
    }
    
    c.agregarHistorial(fmt.Sprintf("Comparación de %.2f y %.2f", a, b))
    return result
}

func (c *Calculadora) sonIguales(a, b float64) bool {
    epsilon := 1e-9
    diff := a - b
    if diff < 0 {
        diff = -diff
    }
    return diff < epsilon
}

// Funciones de memoria
func (c *Calculadora) GuardarMemoria(valor float64) {
    c.memoria = valor
    c.agregarHistorial(fmt.Sprintf("Memoria guardada: %.2f", valor))
}

func (c *Calculadora) RecuperarMemoria() float64 {
    c.agregarHistorial(fmt.Sprintf("Memoria recuperada: %.2f", c.memoria))
    return c.memoria
}

func (c *Calculadora) LimpiarMemoria() {
    c.memoria = 0
    c.agregarHistorial("Memoria limpiada")
}

// Gestión de historial
func (c *Calculadora) agregarHistorial(operacion string) {
    c.historial = append(c.historial, operacion)
    if len(c.historial) > 10 {  // Mantener solo las últimas 10
        c.historial = c.historial[1:]
    }
}

func (c *Calculadora) MostrarHistorial() {
    fmt.Println("\n=== Historial de Operaciones ===")
    if len(c.historial) == 0 {
        fmt.Println("No hay operaciones en el historial")
        return
    }
    
    for i, op := range c.historial {
        fmt.Printf("%d. %s\n", i+1, op)
    }
}

func (c *Calculadora) LimpiarHistorial() {
    c.historial = c.historial[:0]
    fmt.Println("Historial limpiado")
}

// Demostración de la calculadora
func demoCalculadora() {
    fmt.Println("=== Demo Calculadora Avanzada ===")
    
    calc := NewCalculadora()
    
    // Operaciones aritméticas
    fmt.Println("\n--- Operaciones Aritméticas ---")
    fmt.Printf("Suma: %.2f\n", calc.Sumar(15.5, 4.3))
    fmt.Printf("Resta: %.2f\n", calc.Restar(20.0, 5.5))
    fmt.Printf("Multiplicación: %.2f\n", calc.Multiplicar(3.5, 2.0))
    
    if div, err := calc.Dividir(10.0, 3.0); err == nil {
        fmt.Printf("División: %.2f\n", div)
    }
    
    if mod, err := calc.Modulo(17, 5); err == nil {
        fmt.Printf("Módulo: %d\n", mod)
    }
    
    // Operaciones bitwise
    fmt.Println("\n--- Operaciones Bitwise ---")
    fmt.Printf("AND: %d\n", calc.AND(12, 10))
    fmt.Printf("OR: %d\n", calc.OR(12, 10))
    fmt.Printf("XOR: %d\n", calc.XOR(12, 10))
    fmt.Printf("Desplazar izquierda: %d\n", calc.DesplazarIzquierda(5, 2))
    fmt.Printf("Desplazar derecha: %d\n", calc.DesplazarDerecha(20, 2))
    
    // Comparaciones
    fmt.Println("\n--- Comparaciones ---")
    comp := calc.Comparar(7.5, 7.5)
    for op, result := range comp {
        fmt.Printf("%s: %t\n", op, result)
    }
    
    // Memoria
    fmt.Println("\n--- Funciones de Memoria ---")
    calc.GuardarMemoria(42.0)
    fmt.Printf("Valor en memoria: %.2f\n", calc.RecuperarMemoria())
    
    // Mostrar historial
    calc.MostrarHistorial()
}

func main() {
    operadoresBasicos()
    precedenciaAritmetica()
    overflowDemo()
    operadoresAsignacion()
    operadoresComparacion()
    comparacionFlotantes()
    operadoresLogicos()
    cortocircuito()
    patronesLogicos()
    operadoresBitwise()
    casosPracticosBitwise()
    sistemaPermisos()
    precedenciaCompleta()
    casosComplejos()
    demoCalculadora()
}
```
```

### 🎭 Comparaciones Especiales

```go
package main

import (
    "fmt"
    "math"
)

func comparacionesEspeciales() {
    fmt.Println("=== Comparaciones Especiales ===")
    
    // Floating point comparisons
    fmt.Println("Comparaciones de punto flotante:")
    a := 0.1 + 0.2
    b := 0.3
    fmt.Printf("0.1 + 0.2 = %.17f\n", a)
    fmt.Printf("0.3 = %.17f\n", b)
    fmt.Printf("(0.1 + 0.2) == 0.3: %t\n", a == b)  // ¡Puede ser false!
    
    // ✅ Comparación segura de flotantes
    epsilon := 1e-9
    fmt.Printf("Diferencia absoluta: %.2e\n", math.Abs(a-b))
    fmt.Printf("¿Son aproximadamente iguales?: %t\n", math.Abs(a-b) < epsilon)
    
    // Comparación de slices (no se puede con ==)
    fmt.Println("\nComparación de slices:")
    slice1 := []int{1, 2, 3}
    slice2 := []int{1, 2, 3}
    // fmt.Printf("slice1 == slice2: %t\n", slice1 == slice2)  // ❌ Error
    
    // ✅ Comparación manual
    fmt.Printf("¿Slices iguales?: %t\n", slicesIguales(slice1, slice2))
    
    // Comparación de punteros
    fmt.Println("\nComparación de punteros:")
    x, y := 42, 42
    px1, px2 := &x, &x
    py := &y
    
    fmt.Printf("px1 == px2: %t (mismo objeto)\n", px1 == px2)  // true
    fmt.Printf("px1 == py:  %t (objetos diferentes)\n", px1 == py)   // false
    fmt.Printf("*px1 == *py: %t (mismo valor)\n", *px1 == *py) // true
}

func slicesIguales(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
```

---

## 🧠 Operadores Lógicos: La Lógica del Código

### ⚡ AND, OR, NOT

```go
package main

import "fmt"

func operadoresLogicos() {
    fmt.Println("=== Operadores Lógicos ===")
    
    // Variables de prueba
    a, b := true, false
    fmt.Printf("a = %t, b = %t\n", a, b)
    
    // Operadores lógicos básicos
    fmt.Printf("a && b (AND): %t\n", a && b)  // false
    fmt.Printf("a || b (OR):  %t\n", a || b)  // true
    fmt.Printf("!a (NOT):     %t\n", !a)      // false
    fmt.Printf("!b (NOT):     %t\n", !b)      // true
    
    // Combinaciones complejas
    fmt.Println("\nCombinaciones complejas:")
    x, y, z := true, false, true
    fmt.Printf("x && y || z:     %t\n", x && y || z)      // true
    fmt.Printf("x && (y || z):   %t\n", x && (y || z))    // true
    fmt.Printf("(x && y) || z:   %t\n", (x && y) || z)    // true
    fmt.Printf("x && y && z:     %t\n", x && y && z)      // false
    fmt.Printf("x || y || z:     %t\n", x || y || z)      // true
    fmt.Printf("!(x && y):       %t\n", !(x && y))        // true (De Morgan)
    fmt.Printf("!x || !y:        %t\n", !x || !y)         // true (De Morgan)
}
```

### ⚡ Evaluación de Cortocircuito

```go
package main

import "fmt"

func evaluacionCortocircuito() {
    fmt.Println("=== Evaluación de Cortocircuito ===")
    
    // AND (&&) - se detiene en el primer false
    fmt.Println("Prueba AND (&&):")
    resultado1 := falseFunction() && trueFunction()
    fmt.Printf("Resultado: %t\n", resultado1)
    
    fmt.Println("\nPrueba AND (&&) - orden inverso:")
    resultado2 := trueFunction() && falseFunction()
    fmt.Printf("Resultado: %t\n", resultado2)
    
    // OR (||) - se detiene en el primer true
    fmt.Println("\nPrueba OR (||):")
    resultado3 := trueFunction() || falseFunction()
    fmt.Printf("Resultado: %t\n", resultado3)
    
    fmt.Println("\nPrueba OR (||) - orden inverso:")
    resultado4 := falseFunction() || trueFunction()
    fmt.Printf("Resultado: %t\n", resultado4)
    
    // Uso práctico: verificación segura
    fmt.Println("\nVerificación segura:")
    var ptr *int
    // if ptr != nil && *ptr > 0 {  // Seguro - no desreferencia si ptr es nil
    //     fmt.Println("Valor positivo")
    // }
    
    x := 42
    ptr = &x
    if ptr != nil && *ptr > 0 {
        fmt.Printf("Valor positivo: %d\n", *ptr)
    }
}

func trueFunction() bool {
    fmt.Print("  -> trueFunction() ejecutada ")
    return true
}

func falseFunction() bool {
    fmt.Print("  -> falseFunction() ejecutada ")
    return false
}
```

### 🎯 Patrones Idiomáticos

```go
package main

import "fmt"

func patronesLogicos() {
    fmt.Println("=== Patrones Lógicos Idiomáticos ===")
    
    // Patrón: Verificación de múltiples condiciones
    age := 25
    hasLicense := true
    hasInsurance := true
    
    canDrive := age >= 18 && hasLicense && hasInsurance
    fmt.Printf("Puede conducir: %t\n", canDrive)
    
    // Patrón: Valores por defecto con OR
    config := getConfig()
    timeout := config.Timeout
    if timeout <= 0 {
        timeout = 30  // Valor por defecto
    }
    // Más idiomático (aunque Go no tiene operador ternario):
    // timeout = config.Timeout > 0 ? config.Timeout : 30  // ❌ No existe en Go
    
    fmt.Printf("Timeout: %d segundos\n", timeout)
    
    // Patrón: Validación en cadena
    user := User{Name: "Juan", Email: "juan@test.com", Age: 25}
    if isValidUser(user) {
        fmt.Println("Usuario válido")
    }
    
    // Patrón: Guard clauses (cláusulas de guarda)
    if err := processUser(user); err != nil {
        fmt.Printf("Error procesando usuario: %v\n", err)
        return
    }
    fmt.Println("Usuario procesado exitosamente")
}

type Config struct {
    Timeout int
}

func getConfig() Config {
    return Config{Timeout: 0}  // Simula configuración sin timeout
}

type User struct {
    Name  string
    Email string
    Age   int
}

func isValidUser(u User) bool {
    return u.Name != "" && 
           u.Email != "" && 
           u.Age >= 0 && 
           u.Age <= 150
}

func processUser(u User) error {
    // Guard clauses
    if u.Name == "" {
        return fmt.Errorf("nombre requerido")
    }
    if u.Email == "" {
        return fmt.Errorf("email requerido")
    }
    if u.Age < 0 {
        return fmt.Errorf("edad inválida")
    }
    
    // Lógica principal
    fmt.Printf("Procesando usuario: %s\n", u.Name)
    return nil
}
```

---

## 🔧 Operadores Bitwise: El Poder de los Bits

### 🎛️ Operaciones Básicas de Bits

```go
package main

import "fmt"

func operadoresBitwise() {
    fmt.Println("=== Operadores Bitwise ===")
    
    a, b := 12, 10  // 1100 y 1010 en binario
    fmt.Printf("a = %d (%04b), b = %d (%04b)\n", a, a, b, b)
    
    // Operadores bitwise
    fmt.Printf("a & b  (AND): %d (%04b)\n", a&b, a&b)    // 8 (1000)
    fmt.Printf("a | b  (OR):  %d (%04b)\n", a|b, a|b)    // 14 (1110)
    fmt.Printf("a ^ b  (XOR): %d (%04b)\n", a^b, a^b)    // 6 (0110)
    fmt.Printf("^a     (NOT): %d (%b)\n", ^a, ^a)        // -13 (complemento)
    
    // Desplazamientos
    fmt.Println("\nDesplazamientos:")
    x := 5  // 101 en binario
    fmt.Printf("x = %d (%03b)\n", x, x)
    fmt.Printf("x << 2: %d (%05b) (multiplicar por 4)\n", x<<2, x<<2)  // 20 (10100)
    fmt.Printf("x >> 1: %d (%02b) (dividir por 2)\n", x>>1, x>>1)     // 2 (10)
    
    // Operadores de asignación bitwise
    fmt.Println("\nAsignación bitwise:")
    y := 15  // 1111
    fmt.Printf("y inicial: %d (%04b)\n", y, y)
    
    y &= 12  // y = y & 12
    fmt.Printf("y &= 12:   %d (%04b)\n", y, y)  // 12 (1100)
    
    y |= 3   // y = y | 3
    fmt.Printf("y |= 3:    %d (%04b)\n", y, y)  // 15 (1111)
    
    y ^= 5   // y = y ^ 5
    fmt.Printf("y ^= 5:    %d (%04b)\n", y, y)  // 10 (1010)
}
```

### 🎯 Aplicaciones Prácticas de Bitwise

```go
package main

import "fmt"

// Sistema de permisos con bit flags
type Permission uint8

const (
    Read    Permission = 1 << iota  // 00000001
    Write                           // 00000010
    Execute                         // 00000100
    Delete                          // 00001000
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
    return fmt.Sprintf("[%s]", strings.Join(perms, ", "))
}

func aplicacionesBitwise() {
    fmt.Println("=== Aplicaciones Prácticas de Bitwise ===")
    
    // 1. Sistema de permisos
    fmt.Println("1. Sistema de permisos:")
    userPerms := Read | Write           // Combinar permisos
    adminPerms := Read | Write | Execute | Delete
    
    fmt.Printf("Usuario: %s\n", userPerms)
    fmt.Printf("Admin: %s\n", adminPerms)
    
    // Verificar permisos
    if userPerms&Write != 0 {
        fmt.Println("  ✓ Usuario puede escribir")
    }
    if userPerms&Delete == 0 {
        fmt.Println("  ✗ Usuario NO puede eliminar")
    }
    
    // 2. Algoritmos eficientes
    fmt.Println("\n2. Algoritmos eficientes:")
    
    // Verificar si un número es potencia de 2
    numbers := []int{1, 2, 3, 4, 5, 8, 16, 17}
    for _, n := range numbers {
        isPowerOf2 := n > 0 && (n&(n-1)) == 0
        fmt.Printf("%d es potencia de 2: %t\n", n, isPowerOf2)
    }
    
    // 3. Manipulación de bits individuales
    fmt.Println("\n3. Manipulación de bits:")
    var flags uint8 = 0b10101010  // 170 en decimal
    fmt.Printf("Flags inicial: %08b (%d)\n", flags, flags)
    
    // Establecer bit (OR)
    pos := 1
    flags |= (1 << pos)
    fmt.Printf("Establecer bit %d: %08b (%d)\n", pos, flags, flags)
    
    // Limpiar bit (AND con NOT)
    pos = 3
    flags &= ^(1 << pos)
    fmt.Printf("Limpiar bit %d: %08b (%d)\n", pos, flags, flags)
    
    // Alternar bit (XOR)
    pos = 7
    flags ^= (1 << pos)
    fmt.Printf("Alternar bit %d: %08b (%d)\n", pos, flags, flags)
    
    // Verificar bit
    pos = 1
    isSet := (flags>>pos)&1 == 1
    fmt.Printf("Bit %d está establecido: %t\n", pos, isSet)
    
    // 4. Máscara de bits
    fmt.Println("\n4. Máscara de bits:")
    rgb := uint32(0xFF5733)  // Color en hexadecimal
    fmt.Printf("Color RGB: 0x%06X\n", rgb)
    
    // Extraer componentes
    red := (rgb >> 16) & 0xFF
    green := (rgb >> 8) & 0xFF
    blue := rgb & 0xFF
    
    fmt.Printf("Rojo: %d, Verde: %d, Azul: %d\n", red, green, blue)
    
    // Combinar componentes
    newColor := (red << 16) | (green << 8) | blue
    fmt.Printf("Color reconstruido: 0x%06X\n", newColor)
}
```

### 🚀 Algoritmos Avanzados con Bitwise

```go
package main

import "fmt"

func algoritmosAvanzadosBitwise() {
    fmt.Println("=== Algoritmos Avanzados con Bitwise ===")
    
    // 1. Contar bits establecidos (Population Count)
    fmt.Println("1. Contar bits establecidos:")
    numbers := []uint32{7, 15, 255, 1023}
    for _, n := range numbers {
        count := popCount(n)
        fmt.Printf("%d (%b) tiene %d bits establecidos\n", n, n, count)
    }
    
    // 2. Encontrar el bit menos significativo
    fmt.Println("\n2. Bit menos significativo:")
    for _, n := range numbers {
        if n > 0 {
            lsb := n & (-n)  // Truco bitwise para LSB
            fmt.Printf("%d (%b) LSB: %d (%b)\n", n, n, lsb, lsb)
        }
    }
    
    // 3. Intercambiar dos números sin variable temporal
    fmt.Println("\n3. Intercambio sin variable temporal:")
    a, b := 42, 73
    fmt.Printf("Antes: a=%d, b=%d\n", a, b)
    
    a ^= b  // a = a XOR b
    b ^= a  // b = b XOR (a XOR b) = original_a
    a ^= b  // a = (a XOR b) XOR original_a = original_b
    
    fmt.Printf("Después: a=%d, b=%d\n", a, b)
    
    // 4. Verificar si dos números tienen signos opuestos
    fmt.Println("\n4. Signos opuestos:")
    pairs := [][2]int{{5, -3}, {-7, 8}, {4, 9}, {-2, -6}}
    for _, pair := range pairs {
        x, y := pair[0], pair[1]
        opposite := (x ^ y) < 0
        fmt.Printf("%d y %d tienen signos opuestos: %t\n", x, y, opposite)
    }
    
    // 5. Rotación de bits
    fmt.Println("\n5. Rotación de bits:")
    value := uint8(0b10110010)  // 178
    fmt.Printf("Valor original: %08b (%d)\n", value, value)
    
    rotated := rotateLeft8(value, 3)
    fmt.Printf("Rotado 3 izq:   %08b (%d)\n", rotated, rotated)
    
    rotated = rotateRight8(value, 2)
    fmt.Printf("Rotado 2 der:   %08b (%d)\n", rotated, rotated)
}

// Contar bits establecidos (método Brian Kernighan)
func popCount(x uint32) int {
    count := 0
    for x != 0 {
        x &= x - 1  // Elimina el bit menos significativo
        count++
    }
    return count
}

// Rotación a la izquierda para uint8
func rotateLeft8(value uint8, shift int) uint8 {
    shift %= 8  // Manejar rotaciones > 8
    return (value << shift) | (value >> (8 - shift))
}

// Rotación a la derecha para uint8
func rotateRight8(value uint8, shift int) uint8 {
    shift %= 8  // Manejar rotaciones > 8
    return (value >> shift) | (value << (8 - shift))
}
```

---

## 📐 Precedencia de Operadores: El Orden Importa

### 🎯 Tabla de Precedencia

```go
package main

import "fmt"

func precedenciaOperadores() {
    fmt.Println("=== Precedencia de Operadores ===")
    
    // Tabla de precedencia (mayor a menor):
    // 1. * / % << >> & &^        (multiplicativos y bitwise)
    // 2. + - | ^                 (aditivos y bitwise)
    // 3. == != < <= > >=         (comparación)
    // 4. &&                      (AND lógico)
    // 5. ||                      (OR lógico)
    
    fmt.Println("Ejemplos de precedencia:")
    
    // Aritmética vs comparación
    result1 := 2 + 3 * 4    // = 2 + 12 = 14
    result2 := (2 + 3) * 4  // = 5 * 4 = 20
    fmt.Printf("2 + 3 * 4 = %d\n", result1)
    fmt.Printf("(2 + 3) * 4 = %d\n", result2)
    
    // Comparación vs lógicos
    a, b, c := 5, 10, 15
    result3 := a < b && b < c     // (a < b) && (b < c) = true && true = true
    result4 := a < b && c > b     // (a < b) && (c > b) = true && true = true
    fmt.Printf("%d < %d && %d < %d = %t\n", a, b, b, c, result3)
    fmt.Printf("%d < %d && %d > %d = %t\n", a, b, c, b, result4)
    
    // Bitwise vs aritmética
    x := 8 | 4 + 2    // = 8 | (4 + 2) = 8 | 6 = 14
    y := (8 | 4) + 2  // = 12 + 2 = 14
    fmt.Printf("8 | 4 + 2 = %d\n", x)
    fmt.Printf("(8 | 4) + 2 = %d\n", y)
    
    // Casos complejos
    fmt.Println("\nCasos complejos:")
    result5 := 1 + 2 * 3 == 7 && 4 < 5 || false
    // = (1 + (2 * 3)) == 7 && 4 < 5 || false
    // = 7 == 7 && true || false
    // = true && true || false
    // = true || false
    // = true
    fmt.Printf("1 + 2 * 3 == 7 && 4 < 5 || false = %t\n", result5)
}
```

### ⚠️ Trampas de Precedencia

```go
package main

import "fmt"

func trampasPrecedencia() {
    fmt.Println("=== Trampas Comunes de Precedencia ===")
    
    // ❌ Trampa 1: Desplazamiento vs aritmética
    x := 1 + 2 << 3    // = 1 + (2 << 3) = 1 + 16 = 17
    y := (1 + 2) << 3  // = 3 << 3 = 24
    fmt.Printf("1 + 2 << 3 = %d (no %d)\n", x, y)
    
    // ❌ Trampa 2: Bitwise AND vs comparación
    a, b := 6, 4
    result1 := a & b == 4    // = a & (b == 4) = 6 & true = 6 & 1 = 0
    result2 := (a & b) == 4  // = 4 == 4 = true
    fmt.Printf("6 & 4 == 4: resultado incorrecto = %d, correcto = %t\n", result1, result2)
    
    // ❌ Trampa 3: Asignación vs comparación
    flag := true
    // if flag = false {  // ❌ Error: asignación en if no permitida en Go
    //     fmt.Println("Esto no compila")
    // }
    
    if flag == false {  // ✅ Correcto
        fmt.Println("Flag es false")
    } else {
        fmt.Println("Flag es true")
    }
    
    // ✅ Mejor práctica: usar paréntesis para claridad
    fmt.Println("\nMejores prácticas:")
    result3 := (1 + 2) * (3 + 4)    // Claro y explícito
    result4 := (a & b) == 4         // Precedencia clara
    result5 := (x > 0) && (y < 10)  // Lógica clara
    
    fmt.Printf("Expresiones claras: %d, %t, %t\n", result3, result4, result5)
}
```

---

## 🧪 Laboratorio: Sistema de Cálculos Avanzados

### 🎯 Proyecto: Calculadora Científica

```go
package main

import (
    "fmt"
    "math"
    "strings"
)

// BitField para representar conjunto de características
type Features uint32

const (
    BasicMath Features = 1 << iota  // 0001
    Scientific                      // 0010
    Programming                     // 0100
    Statistics                      // 1000
)

func (f Features) String() string {
    var features []string
    if f&BasicMath != 0 {
        features = append(features, "BasicMath")
    }
    if f&Scientific != 0 {
        features = append(features, "Scientific")
    }
    if f&Programming != 0 {
        features = append(features, "Programming")
    }
    if f&Statistics != 0 {
        features = append(features, "Statistics")
    }
    return strings.Join(features, "|")
}

// Calculadora con diferentes modos
type Calculator struct {
    features Features
    precision int
    memory   float64
}

// Constructor que aprovecha zero values
func NewCalculator() *Calculator {
    return &Calculator{
        features:  BasicMath,  // Características básicas por defecto
        precision: 2,          // 2 decimales por defecto
        // memory queda en 0 (zero value)
    }
}

// Operaciones básicas
func (c *Calculator) Add(a, b float64) float64 {
    return c.round(a + b)
}

func (c *Calculator) Subtract(a, b float64) float64 {
    return c.round(a - b)
}

func (c *Calculator) Multiply(a, b float64) float64 {
    return c.round(a * b)
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return c.round(a / b), nil
}

func (c *Calculator) Modulo(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("módulo por cero")
    }
    return a % b, nil
}

// Operaciones científicas (requiere feature Scientific)
func (c *Calculator) Power(base, exp float64) (float64, error) {
    if !c.hasFeature(Scientific) {
        return 0, fmt.Errorf("función científica no disponible")
    }
    return c.round(math.Pow(base, exp)), nil
}

func (c *Calculator) SquareRoot(x float64) (float64, error) {
    if !c.hasFeature(Scientific) {
        return 0, fmt.Errorf("función científica no disponible")
    }
    if x < 0 {
        return 0, fmt.Errorf("raíz cuadrada de número negativo")
    }
    return c.round(math.Sqrt(x)), nil
}

// Operaciones de programación (requiere feature Programming)
func (c *Calculator) BitwiseAND(a, b uint64) (uint64, error) {
    if !c.hasFeature(Programming) {
        return 0, fmt.Errorf("función de programación no disponible")
    }
    return a & b, nil
}

func (c *Calculator) BitwiseOR(a, b uint64) (uint64, error) {
    if !c.hasFeature(Programming) {
        return 0, fmt.Errorf("función de programación no disponible")
    }
    return a | b, nil
}

func (c *Calculator) BitwiseXOR(a, b uint64) (uint64, error) {
    if !c.hasFeature(Programming) {
        return 0, fmt.Errorf("función de programación no disponible")
    }
    return a ^ b, nil
}

func (c *Calculator) LeftShift(value uint64, positions int) (uint64, error) {
    if !c.hasFeature(Programming) {
        return 0, fmt.Errorf("función de programación no disponible")
    }
    if positions < 0 {
        return 0, fmt.Errorf("posiciones negativas no permitidas")
    }
    return value << positions, nil
}

func (c *Calculator) RightShift(value uint64, positions int) (uint64, error) {
    if !c.hasFeature(Programming) {
        return 0, fmt.Errorf("función de programación no disponible")
    }
    if positions < 0 {
        return 0, fmt.Errorf("posiciones negativas no permitidas")
    }
    return value >> positions, nil
}

// Operaciones estadísticas (requiere feature Statistics)
func (c *Calculator) Mean(values []float64) (float64, error) {
    if !c.hasFeature(Statistics) {
        return 0, fmt.Errorf("función estadística no disponible")
    }
    if len(values) == 0 {
        return 0, fmt.Errorf("no se puede calcular media de slice vacío")
    }
    
    sum := 0.0
    for _, v := range values {
        sum += v
    }
    return c.round(sum / float64(len(values))), nil
}

func (c *Calculator) Max(values []float64) (float64, error) {
    if !c.hasFeature(Statistics) {
        return 0, fmt.Errorf("función estadística no disponible")
    }
    if len(values) == 0 {
        return 0, fmt.Errorf("no se puede calcular máximo de slice vacío")
    }
    
    max := values[0]
    for _, v := range values[1:] {
        if v > max {
            max = v
        }
    }
    return max, nil
}

// Gestión de memoria
func (c *Calculator) StoreMemory(value float64) {
    c.memory = value
}

func (c *Calculator) RecallMemory() float64 {
    return c.memory
}

func (c *Calculator) AddToMemory(value float64) {
    c.memory += value
}

func (c *Calculator) ClearMemory() {
    c.memory = 0
}

// Gestión de características
func (c *Calculator) EnableFeature(feature Features) {
    c.features |= feature  // Usar OR para activar bit
}

func (c *Calculator) DisableFeature(feature Features) {
    c.features &= ^feature  // Usar AND con NOT para desactivar bit
}

func (c *Calculator) hasFeature(feature Features) bool {
    return c.features&feature != 0
}

func (c *Calculator) GetFeatures() Features {
    return c.features
}

// Configuración
func (c *Calculator) SetPrecision(precision int) error {
    if precision < 0 || precision > 10 {
        return fmt.Errorf("precisión debe estar entre 0 y 10")
    }
    c.precision = precision
    return nil
}

// Función auxiliar para redondear
func (c *Calculator) round(value float64) float64 {
    multiplier := math.Pow(10, float64(c.precision))
    return math.Round(value*multiplier) / multiplier
}

// Función auxiliar para verificar si un número es potencia de 2
func (c *Calculator) IsPowerOfTwo(n uint64) bool {
    return n > 0 && (n&(n-1)) == 0
}

// Función auxiliar para contar bits establecidos
func (c *Calculator) CountBits(n uint64) int {
    count := 0
    for n != 0 {
        n &= n - 1  // Elimina el bit menos significativo
        count++
    }
    return count
}

func main() {
    fmt.Println("🧮 === CALCULADORA CIENTÍFICA AVANZADA ===\n")
    
    // Crear calculadora
    calc := NewCalculator()
    fmt.Printf("Calculadora creada con características: %s\n", calc.GetFeatures())
    fmt.Printf("Precisión: %d decimales\n\n", calc.precision)
    
    // Operaciones básicas
    fmt.Println("=== OPERACIONES BÁSICAS ===")
    fmt.Printf("15 + 7 = %.2f\n", calc.Add(15, 7))
    fmt.Printf("15 - 7 = %.2f\n", calc.Subtract(15, 7))
    fmt.Printf("15 * 7 = %.2f\n", calc.Multiply(15, 7))
    
    if result, err := calc.Divide(15, 7); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("15 / 7 = %.2f\n", result)
    }
    
    if result, err := calc.Modulo(15, 7); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("15 %% 7 = %d\n", result)
    }
    
    // Probar funciones científicas sin habilitarlas
    fmt.Println("\n=== INTENTAR FUNCIONES CIENTÍFICAS ===")
    if _, err := calc.Power(2, 3); err != nil {
        fmt.Printf("Error esperado: %v\n", err)
    }
    
    // Habilitar funciones científicas
    calc.EnableFeature(Scientific)
    fmt.Printf("\nCaracterísticas habilitadas: %s\n", calc.GetFeatures())
    
    if result, err := calc.Power(2, 3); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("2^3 = %.2f\n", result)
    }
    
    if result, err := calc.SquareRoot(16); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("√16 = %.2f\n", result)
    }
    
    // Habilitar funciones de programación
    calc.EnableFeature(Programming)
    fmt.Printf("\nCaracterísticas habilitadas: %s\n", calc.GetFeatures())
    
    fmt.Println("\n=== OPERACIONES BITWISE ===")
    a, b := uint64(12), uint64(10)  // 1100 y 1010 en binario
    fmt.Printf("a = %d (%04b), b = %d (%04b)\n", a, a, b, b)
    
    if result, err := calc.BitwiseAND(a, b); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("a & b = %d (%04b)\n", result, result)
    }
    
    if result, err := calc.BitwiseOR(a, b); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("a | b = %d (%04b)\n", result, result)
    }
    
    if result, err := calc.BitwiseXOR(a, b); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("a ^ b = %d (%04b)\n", result, result)
    }
    
    if result, err := calc.LeftShift(5, 2); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("5 << 2 = %d\n", result)
    }
    
    // Funciones auxiliares
    fmt.Println("\n=== FUNCIONES AUXILIARES ===")
    numbers := []uint64{1, 2, 3, 4, 8, 15, 16}
    for _, n := range numbers {
        isPower := calc.IsPowerOfTwo(n)
        bitCount := calc.CountBits(n)
        fmt.Printf("%d: potencia de 2? %t, bits: %d\n", n, isPower, bitCount)
    }
    
    // Habilitar estadísticas
    calc.EnableFeature(Statistics)
    fmt.Printf("\nCaracterísticas habilitadas: %s\n", calc.GetFeatures())
    
    fmt.Println("\n=== OPERACIONES ESTADÍSTICAS ===")
    values := []float64{1.5, 2.3, 3.7, 4.1, 5.9}
    
    if mean, err := calc.Mean(values); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Media de %v = %.2f\n", values, mean)
    }
    
    if max, err := calc.Max(values); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Máximo de %v = %.2f\n", values, max)
    }
    
    // Operaciones de memoria
    fmt.Println("\n=== OPERACIONES DE MEMORIA ===")
    calc.StoreMemory(42.5)
    fmt.Printf("Memoria almacenada: %.2f\n", calc.RecallMemory())
    
    calc.AddToMemory(7.5)
    fmt.Printf("Memoria después de sumar 7.5: %.2f\n", calc.RecallMemory())
    
    calc.ClearMemory()
    fmt.Printf("Memoria después de limpiar: %.2f\n", calc.RecallMemory())
    
    // Cambiar precisión
    fmt.Println("\n=== CAMBIO DE PRECISIÓN ===")
    calc.SetPrecision(4)
    result := calc.Divide(22, 7)  // π aproximado
    fmt.Printf("22/7 con 4 decimales = %.4f\n", result)
    
    fmt.Println("\n🎉 ¡Calculadora funcionando perfectamente!")
    fmt.Printf("Características finales: %s\n", calc.GetFeatures())
}
```

---

## 🎯 Best Practices

### ✅ Operadores Aritméticos

1. **Cuidado con división entera** - Usa conversión explícita para flotantes
2. **Verifica división por cero** antes de dividir
3. **Usa operadores compuestos** (+=, -=, etc.) para claridad
4. **Prefiere incremento/decremento** (++, --) cuando sea apropiado

### ✅ Operadores de Comparación

1. **Usa epsilon para comparar flotantes** - Nunca uses == directamente
2. **Convierte tipos explícitamente** para comparaciones
3. **Aprovecha cortocircuito** en expresiones lógicas
4. **Usa paréntesis** para clarificar precedencia

### ✅ Operadores Bitwise

1. **Documenta operaciones bitwise** complejas
2. **Usa constantes con nombres** para bit flags
3. **Prefiere métodos con nombres** para operaciones complejas
4. **Considera rendimiento** - bitwise es muy eficiente

### ✅ Patrones Recomendados

```go
// ✅ División segura
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

// ✅ Comparación de flotantes
func floatEquals(a, b, epsilon float64) bool {
    return math.Abs(a-b) < epsilon
}

// ✅ Bit flags bien documentados
type FileMode uint32

const (
    ReadOnly  FileMode = 1 << iota  // 001
    WriteOnly                       // 010
    ReadWrite = ReadOnly | WriteOnly // 011
)

// ✅ Guard clauses
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("datos vacíos")
    }
    if len(data) > maxSize {
        return errors.New("datos demasiado grandes")
    }
    
    // Lógica principal aquí
    return nil
}
```

---

## 🎉 ¡Felicitaciones!

¡Has dominado todos los operadores de Go! Ahora puedes:

- ✅ **Usar operadores aritméticos** de forma segura y eficiente
- ✅ **Comparar valores** correctamente incluyendo flotantes
- ✅ **Aplicar lógica compleja** con operadores lógicos
- ✅ **Manipular bits** para algoritmos eficientes
- ✅ **Entender precedencia** y usar paréntesis apropiadamente
- ✅ **Implementar sistemas complejos** como la calculadora

### 🔥 Conceptos Dominados:

1. **Operadores aritméticos** - Matemáticas seguras y eficientes
2. **Operadores de comparación** - Evaluación correcta de condiciones
3. **Operadores lógicos** - Lógica booleana y cortocircuito
4. **Operadores bitwise** - Manipulación de bits y algoritmos eficientes
5. **Precedencia** - Orden correcto de evaluación
6. **Patrones idiomáticos** - Código limpio y mantenible

---

## 📚 Mejores Prácticas y Estilo

### ✅ Recomendaciones

```go
// ✅ BUENAS PRÁCTICAS

// 1. Usa paréntesis para clarificar precedencia compleja
result := (a + b) * (c - d)  // ✅ Claro
// vs
result := a + b * c - d      // ❌ Confuso

// 2. Prefiere operadores de asignación cuando sea apropiado
counter += 1    // ✅ Idiomático
counter++       // ✅ Aún mejor para incremento
counter = counter + 1  // ❌ Verboso

// 3. Para flotantes, siempre considera la precisión
func isEqual(a, b, epsilon float64) bool {
    diff := a - b
    if diff < 0 {
        diff = -diff
    }
    return diff < epsilon
}

// 4. Usa operadores bitwise para flags y permisos
type Status uint8
const (
    Active   Status = 1 << iota
    Verified 
    Premium  
)

// 5. Evita operadores complejos en condicionales
// ✅ Claro
isValid := age >= 18 && age <= 65
hasPermission := user.Role == "admin"
if isValid && hasPermission {
    // ...
}

// ❌ Confuso
if age >= 18 && age <= 65 && user.Role == "admin" && user.Active && !user.Suspended {
    // ...
}
```

### ❌ Errores Comunes

```go
// ❌ ERRORES FRECUENTES

// 1. Comparación directa de flotantes
if 0.1 + 0.2 == 0.3 {  // ❌ Puede fallar
    // ...
}

// 2. Confundir precedencia
result := 2 + 3 * 4  // Es 14, no 20

// 3. No manejar overflow
var x int8 = 127
x++  // ❌ Overflow a -128

// 4. Usar bitwise en lugar de lógico (o viceversa)
if flag & mask {     // ❌ En Go debe ser != 0
    // ...
}

if flag && mask {    // ❌ Si son números
    // ...
}

// 5. Asignación en lugar de comparación
if x = 5 {  // ❌ Error de compilación en Go (¡bien!)
    // ...
}
```

---

## 🎯 Ejercicios Prácticos

### Ejercicio 1: Calculadora de Días
```go
// Implementa una función que calcule:
// 1. Días entre dos fechas
// 2. Si un año es bisiesto
// 3. Días restantes hasta fin de año

func diasEntreFechas(dia1, mes1, año1, dia2, mes2, año2 int) int {
    // Tu implementación aquí
    return 0
}

func esBisiesto(año int) bool {
    // Tu implementación aquí
    return false
}
```

### Ejercicio 2: Sistema de Permisos
```go
// Implementa un sistema de permisos con bitwise
// Permisos: READ(1), WRITE(2), EXECUTE(4), DELETE(8), ADMIN(16)

type Usuario struct {
    Nombre    string
    Permisos  uint8
}

func (u *Usuario) TienePermiso(permiso uint8) bool {
    // Tu implementación
    return false
}

func (u *Usuario) AgregarPermiso(permiso uint8) {
    // Tu implementación
}

func (u *Usuario) QuitarPermiso(permiso uint8) {
    // Tu implementación
}
```

### Ejercicio 3: Evaluador de Expresiones
```go
// Implementa un evaluador simple de expresiones matemáticas
// Debe manejar +, -, *, / y paréntesis
// Ejemplo: "2 + 3 * 4" = 14, "(2 + 3) * 4" = 20

func evaluarExpresion(expresion string) (float64, error) {
    // Tu implementación aquí
    return 0, nil
}
```

### Ejercicio 4: Manipulación de Bits
```go
// Implementa las siguientes funciones:
// 1. Contar bits activos en un número
// 2. Verificar si es potencia de 2
// 3. Encontrar el bit más significativo
// 4. Intercambiar dos bits específicos

func contarBits(n uint32) int {
    // Tu implementación
    return 0
}

func esPotenciaDe2(n uint32) bool {
    // Tu implementación
    return false
}

func bitMasSignificativo(n uint32) int {
    // Tu implementación
    return 0
}

func intercambiarBits(n uint32, pos1, pos2 int) uint32 {
    // Tu implementación
    return 0
}
```

---

## 🏆 Desafíos Avanzados

### Desafío 1: Calculadora de Números Complejos
```go
type Complejo struct {
    Real, Imag float64
}

// Implementa todas las operaciones básicas (+, -, *, /)
// y funciones como módulo, argumento, conjugado
```

### Desafío 2: Parser de Expresiones con Precedencia
```go
// Implementa un parser que maneje:
// - Operadores aritméticos con precedencia correcta
// - Funciones matemáticas (sin, cos, sqrt, etc.)
// - Variables y constantes
// - Paréntesis anidados
```

### Desafío 3: Sistema de Flags Avanzado
```go
// Implementa un sistema de configuración usando bitwise
// que permita:
// - Múltiples categorías de flags
// - Serialización/deserialización
// - Validación de combinaciones válidas
// - Herencia de permisos
```

---

## 📖 Conceptos Clave para Recordar

1. **Precedencia de Operadores**: Memoriza el orden básico y usa paréntesis cuando dudes
2. **Cortocircuito**: && y || no evalúan el segundo operando si no es necesario
3. **Comparación de Flotantes**: Nunca uses == directamente, siempre con epsilon
4. **Overflow**: Siempre considera los límites de tus tipos de datos
5. **Operadores Bitwise**: Poderosos para flags, optimizaciones y manipulación de bits
6. **Asignación Compuesta**: Más eficiente y legible que la asignación completa

---

## 🎓 Resumen y Siguientes Pasos

### 📝 Lo que Aprendiste

En esta lección has dominado:

- ✅ **Operadores Aritméticos**: Suma, resta, multiplicación, división, módulo
- ✅ **Operadores de Asignación**: =, +=, -=, *=, /=, %=, ++, --
- ✅ **Operadores de Comparación**: ==, !=, <, <=, >, >=
- ✅ **Operadores Lógicos**: &&, ||, ! y evaluación de cortocircuito
- ✅ **Operadores Bitwise**: &, |, ^, <<, >>, &^ y sus aplicaciones
- ✅ **Precedencia y Asociatividad**: Orden de evaluación y buenas prácticas
- ✅ **Casos Especiales**: Overflow, comparación de flotantes, flags con bits

### 🚀 Próximos Pasos

1. **Practica los Ejercicios**: Completa todos los ejercicios propuestos
2. **Experimenta**: Crea tus propias combinaciones de operadores
3. **Siguiente Lección**: Control de Flujo (if, switch, loops)
4. **Proyecto**: Implementa una calculadora científica completa

### 💡 Consejos para el Éxito

- **Práctica Diaria**: Usa operadores en pequeños programas cada día
- **Lectura de Código**: Analiza código Go real para ver operadores en contexto
- **Debugging**: Aprende a debuggear expresiones complejas paso a paso
- **Performance**: Entiende cuándo los operadores bitwise pueden optimizar tu código

---

> 💬 **Reflexión**: Los operadores son las herramientas básicas con las que construyes lógica en Go. Dominarlos no solo te hace más eficiente, sino que te permite escribir código más elegante y expresivo.

**¡Excelente trabajo!** 🎉 Ahora tienes una base sólida en operadores de Go. En la siguiente lección exploraremos cómo usar estos operadores para controlar el flujo de ejecución de tus programas.

---

📁 **Archivos de esta lección:**
- `README.md` - Teoría completa y ejemplos
- `ejercicios.go` - Ejercicios prácticos para resolver
- `soluciones.go` - Soluciones detalladas y explicadas
- `proyecto_calculadora.go` - Proyecto práctico completo

---

### 🚀 Próximo Nivel

¡Es hora de controlar el flujo de tus programas!

**[→ Ir a la Lección 7: Estructuras de Control](../07-estructuras-control/)**

---

## 📞 ¿Preguntas?

- 💬 **Discord**: [Go Deep Community](#)
- 📧 **Email**: support@go-deep.dev
- 🐛 **Issues**: [GitHub Issues](../../../issues)

---

*¡Tus operaciones están bajo control! Hora de estructurar el flujo ⚡*
