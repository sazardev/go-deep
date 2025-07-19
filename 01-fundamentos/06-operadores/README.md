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
