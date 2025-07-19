package main

import (
	"fmt"
	"math"
	"strings"
)

// ========== EJERCICIO 1: OPERADORES ARITMÉTICOS ==========

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Operadores Aritméticos ===")

	// Operaciones básicas
	a, b := 25, 4
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Suma: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Resta: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplicación: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("División entera: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("División flotante: %.2f / %.2f = %.2f\n", float64(a), float64(b), float64(a)/float64(b))
	fmt.Printf("Módulo: %d %% %d = %d\n", a, b, a%b)

	// Operadores de asignación
	x := 10
	fmt.Printf("\nOperadores de asignación (x inicial = %d):\n", x)
	x += 5
	fmt.Printf("x += 5: %d\n", x)
	x *= 2
	fmt.Printf("x *= 2: %d\n", x)
	x /= 3
	fmt.Printf("x /= 3: %d\n", x)
	x++
	fmt.Printf("x++: %d\n", x)
	x--
	fmt.Printf("x--: %d\n", x)
}

// ========== EJERCICIO 2: OPERADORES DE COMPARACIÓN ==========

func ejercicio2() {
	fmt.Println("\n=== Ejercicio 2: Operadores de Comparación ===")

	a, b := 15, 10
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a == b: %t\n", a == b)
	fmt.Printf("a != b: %t\n", a != b)
	fmt.Printf("a > b: %t\n", a > b)
	fmt.Printf("a < b: %t\n", a < b)
	fmt.Printf("a >= b: %t\n", a >= b)
	fmt.Printf("a <= b: %t\n", a <= b)

	// Comparación de strings
	str1, str2 := "apple", "banana"
	fmt.Printf("\nComparación de strings:\n")
	fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)

	// Comparación segura de flotantes
	fmt.Println("\nComparación segura de flotantes:")
	f1, f2 := 0.1+0.2, 0.3
	fmt.Printf("0.1 + 0.2 == 0.3: %t (¡Cuidado!)\n", f1 == f2)
	fmt.Printf("Comparación segura: %t\n", floatEquals(f1, f2, 1e-9))
}

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// ========== EJERCICIO 3: OPERADORES LÓGICOS ==========

func ejercicio3() {
	fmt.Println("\n=== Ejercicio 3: Operadores Lógicos ===")

	p, q := true, false
	fmt.Printf("p = %t, q = %t\n", p, q)
	fmt.Printf("p && q: %t\n", p && q)
	fmt.Printf("p || q: %t\n", p || q)
	fmt.Printf("!p: %t\n", !p)
	fmt.Printf("!q: %t\n", !q)

	// Expresiones complejas
	fmt.Println("\nExpresiones complejas:")
	a, b, c := true, false, true
	fmt.Printf("a && b || c: %t\n", a && b || c)
	fmt.Printf("a && (b || c): %t\n", a && (b || c))
	fmt.Printf("(a && b) || c: %t\n", (a && b) || c)

	// Cortocircuito
	fmt.Println("\nEvaluación de cortocircuito:")
	result1 := expensiveTrue() && expensiveFalse()
	fmt.Printf("true && false = %t\n", result1)

	result2 := expensiveFalse() && expensiveTrue()
	fmt.Printf("false && true = %t (segunda función no se ejecuta)\n", result2)
}

func expensiveTrue() bool {
	fmt.Print("  -> expensiveTrue() ")
	return true
}

func expensiveFalse() bool {
	fmt.Print("  -> expensiveFalse() ")
	return false
}

// ========== EJERCICIO 4: OPERADORES BITWISE ==========

func ejercicio4() {
	fmt.Println("\n=== Ejercicio 4: Operadores Bitwise ===")

	a, b := uint8(12), uint8(10) // 1100 y 1010 en binario
	fmt.Printf("a = %d (%08b), b = %d (%08b)\n", a, a, b, b)
	fmt.Printf("a & b  (AND): %d (%08b)\n", a&b, a&b)
	fmt.Printf("a | b  (OR):  %d (%08b)\n", a|b, a|b)
	fmt.Printf("a ^ b  (XOR): %d (%08b)\n", a^b, a^b)
	fmt.Printf("^a     (NOT): %d (%08b)\n", ^a, ^a)

	// Desplazamientos
	x := uint8(5) // 00000101
	fmt.Printf("\nDesplazamientos (x = %d, %08b):\n", x, x)
	fmt.Printf("x << 2: %d (%08b)\n", x<<2, x<<2)
	fmt.Printf("x >> 1: %d (%08b)\n", x>>1, x>>1)

	// Aplicaciones prácticas
	fmt.Println("\nAplicaciones prácticas:")
	numbers := []int{1, 2, 3, 4, 8, 15, 16, 32}
	for _, n := range numbers {
		isPowerOf2 := n > 0 && (n&(n-1)) == 0
		fmt.Printf("%d es potencia de 2: %t\n", n, isPowerOf2)
	}
}

// ========== EJERCICIO 5: PRECEDENCIA DE OPERADORES ==========

func ejercicio5() {
	fmt.Println("\n=== Ejercicio 5: Precedencia de Operadores ===")

	// Aritmética vs comparación
	result1 := 2 + 3*4     // = 2 + 12 = 14
	result2 := (2 + 3) * 4 // = 5 * 4 = 20
	fmt.Printf("2 + 3 * 4 = %d\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d\n", result2)

	// Comparación vs lógicos
	a, b, c := 5, 10, 15
	result3 := a < b && b < c // true && true = true
	fmt.Printf("%d < %d && %d < %d = %t\n", a, b, b, c, result3)

	// Bitwise vs aritmética
	x := 8 | 4 + 2   // = 8 | 6 = 14
	y := (8 | 4) + 2 // = 12 + 2 = 14
	fmt.Printf("8 | 4 + 2 = %d\n", x)
	fmt.Printf("(8 | 4) + 2 = %d\n", y)

	// Expresión compleja
	result4 := 1+2*3 == 7 && 4 < 5 || false
	fmt.Printf("1 + 2 * 3 == 7 && 4 < 5 || false = %t\n", result4)

	// Trampas comunes
	fmt.Println("\nTrampas de precedencia:")
	val1 := 1 + 2<<3     // = 1 + (2 << 3) = 1 + 16 = 17
	val2 := (1 + 2) << 3 // = 3 << 3 = 24
	fmt.Printf("1 + 2 << 3 = %d (¡cuidado!), (1 + 2) << 3 = %d\n", val1, val2)
}

// ========== EJERCICIO 6: SISTEMA DE PERMISOS ==========

type Permission uint8

const (
	Read    Permission = 1 << iota // 00000001
	Write                          // 00000010
	Execute                        // 00000100
	Delete                         // 00001000
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
	if len(perms) == 0 {
		return "None"
	}
	return strings.Join(perms, "|")
}

func ejercicio6() {
	fmt.Println("\n=== Ejercicio 6: Sistema de Permisos ===")

	// Definir roles
	guest := Permission(0)                   // Sin permisos
	user := Read | Write                     // Lectura y escritura
	moderator := Read | Write | Delete       // + Eliminar
	admin := Read | Write | Execute | Delete // Todos los permisos

	fmt.Printf("Guest: %s (%08b)\n", guest, guest)
	fmt.Printf("User: %s (%08b)\n", user, user)
	fmt.Printf("Moderator: %s (%08b)\n", moderator, moderator)
	fmt.Printf("Admin: %s (%08b)\n", admin, admin)

	// Verificar permisos
	fmt.Println("\nVerificación de permisos:")
	fmt.Printf("User puede leer: %t\n", hasPermission(user, Read))
	fmt.Printf("User puede eliminar: %t\n", hasPermission(user, Delete))
	fmt.Printf("Moderator puede eliminar: %t\n", hasPermission(moderator, Delete))
	fmt.Printf("Admin puede ejecutar: %t\n", hasPermission(admin, Execute))

	// Modificar permisos
	fmt.Println("\nModificar permisos:")
	userUpdated := addPermission(user, Execute)
	fmt.Printf("User + Execute: %s\n", userUpdated)

	userDowngraded := removePermission(userUpdated, Write)
	fmt.Printf("User - Write: %s\n", userDowngraded)

	userToggled := togglePermission(user, Delete)
	fmt.Printf("User toggle Delete: %s\n", userToggled)
}

func hasPermission(userPerms, perm Permission) bool {
	return userPerms&perm != 0
}

func addPermission(userPerms, perm Permission) Permission {
	return userPerms | perm
}

func removePermission(userPerms, perm Permission) Permission {
	return userPerms &^ perm // AND NOT
}

func togglePermission(userPerms, perm Permission) Permission {
	return userPerms ^ perm
}

// ========== EJERCICIO 7: CALCULADORA SIMPLE ==========

type Calculator struct {
	memory float64
}

func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

func (c *Calculator) Power(base, exp float64) float64 {
	return math.Pow(base, exp)
}

func (c *Calculator) StoreMemory(value float64) {
	c.memory = value
}

func (c *Calculator) RecallMemory() float64 {
	return c.memory
}

func (c *Calculator) ClearMemory() {
	c.memory = 0
}

func ejercicio7() {
	fmt.Println("\n=== Ejercicio 7: Calculadora Simple ===")

	calc := &Calculator{}

	// Operaciones básicas
	fmt.Printf("5 + 3 = %.2f\n", calc.Add(5, 3))
	fmt.Printf("10 - 4 = %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("6 * 7 = %.2f\n", calc.Multiply(6, 7))

	if result, err := calc.Divide(15, 3); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("15 / 3 = %.2f\n", result)
	}

	fmt.Printf("2^8 = %.0f\n", calc.Power(2, 8))

	// Operaciones de memoria
	fmt.Println("\nOperaciones de memoria:")
	calc.StoreMemory(42.5)
	fmt.Printf("Memoria almacenada: %.2f\n", calc.RecallMemory())
	calc.ClearMemory()
	fmt.Printf("Memoria después de limpiar: %.2f\n", calc.RecallMemory())
}

// ========== EJERCICIO 8: ALGORITMOS BITWISE ==========

func ejercicio8() {
	fmt.Println("\n=== Ejercicio 8: Algoritmos Bitwise ===")

	// Contar bits establecidos
	fmt.Println("Contar bits establecidos:")
	numbers := []uint32{7, 15, 255, 1023}
	for _, n := range numbers {
		count := popCount(n)
		fmt.Printf("%d (%b) tiene %d bits establecidos\n", n, n, count)
	}

	// Intercambio sin variable temporal
	fmt.Println("\nIntercambio XOR:")
	a, b := 42, 73
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("Después: a=%d, b=%d\n", a, b)

	// Bit menos significativo
	fmt.Println("\nBit menos significativo:")
	for _, n := range numbers {
		if n > 0 {
			lsb := n & (-n)
			fmt.Printf("%d (%b) LSB: %d (%b)\n", n, n, lsb, lsb)
		}
	}

	// Verificar potencias de 2
	fmt.Println("\nVerificar potencias de 2:")
	testNumbers := []int{1, 2, 3, 4, 5, 8, 16, 17, 32}
	for _, n := range testNumbers {
		isPowerOf2 := n > 0 && (n&(n-1)) == 0
		fmt.Printf("%d es potencia de 2: %t\n", n, isPowerOf2)
	}
}

func popCount(x uint32) int {
	count := 0
	for x != 0 {
		x &= x - 1 // Elimina el bit menos significativo
		count++
	}
	return count
}

// ========== EJERCICIO 9: VALIDACIÓN DE DATOS ==========

func ejercicio9() {
	fmt.Println("\n=== Ejercicio 9: Validación de Datos ===")

	// Validar rangos de edad
	ages := []int{-5, 0, 18, 25, 65, 120, 150}
	for _, age := range ages {
		isValid := age >= 0 && age <= 120
		category := getAgeCategory(age)
		fmt.Printf("Edad %d: válida=%t, categoría=%s\n", age, isValid, category)
	}

	// Validar emails (simplificado)
	emails := []string{"test@example.com", "invalid-email", "user@domain", "a@b.co"}
	for _, email := range emails {
		isValid := isValidEmail(email)
		fmt.Printf("Email '%s': válido=%t\n", email, isValid)
	}

	// Validar rangos numéricos
	fmt.Println("\nValidación de rangos:")
	numbers := []float64{-10, 0, 50, 100, 150}
	for _, num := range numbers {
		inRange := num >= 0 && num <= 100
		fmt.Printf("%.1f está en rango [0-100]: %t\n", num, inRange)
	}
}

func getAgeCategory(age int) string {
	if age < 0 {
		return "inválida"
	} else if age < 13 {
		return "niño"
	} else if age < 18 {
		return "adolescente"
	} else if age < 65 {
		return "adulto"
	} else if age <= 120 {
		return "adulto mayor"
	} else {
		return "inválida"
	}
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// ========== EJERCICIO 10: OPTIMIZACIONES ==========

func ejercicio10() {
	fmt.Println("\n=== Ejercicio 10: Optimizaciones ===")

	// Multiplicación y división por potencias de 2
	fmt.Println("Multiplicación/división por potencias de 2:")
	x := 25
	fmt.Printf("x = %d\n", x)
	fmt.Printf("x * 4 = %d (usando x << 2 = %d)\n", x*4, x<<2)
	fmt.Printf("x * 8 = %d (usando x << 3 = %d)\n", x*8, x<<3)
	fmt.Printf("x / 4 = %d (usando x >> 2 = %d)\n", x/4, x>>2)

	// Verificar paridad
	fmt.Println("\nVerificar paridad:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, n := range numbers {
		isEven := (n & 1) == 0
		fmt.Printf("%d es par: %t\n", n, isEven)
	}

	// Intercambio de variables usando XOR (solo para enteros)
	fmt.Println("\nIntercambio optimizado:")
	a, b := 15, 27
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	// Método tradicional vs XOR
	// temp := a; a = b; b = temp  // Método tradicional
	a ^= b
	b ^= a
	a ^= b // Método XOR (sin variable temporal)
	fmt.Printf("Después: a=%d, b=%d\n", a, b)

	// Encontrar el máximo sin condicionales
	fmt.Println("\nMáximo sin condicionales:")
	x1, y1 := 42, 37
	max1 := maxWithoutBranching(x1, y1)
	fmt.Printf("max(%d, %d) = %d\n", x1, y1, max1)

	x2, y2 := 15, 28
	max2 := maxWithoutBranching(x2, y2)
	fmt.Printf("max(%d, %d) = %d\n", x2, y2, max2)
}

// Máximo sin usar if/else (truco bitwise)
func maxWithoutBranching(a, b int) int {
	// Esta es una técnica avanzada, normalmente se prefiere usar math.Max o if
	diff := a - b
	sign := (diff >> 63) & 1 // En Go int es de 64 bits en arquitecturas de 64 bits
	return a - diff*sign
}

// Versión más práctica y legible
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("⚡ === LABORATORIO: OPERADORES EN GO ===")

	ejercicio1()  // Operadores aritméticos
	ejercicio2()  // Operadores de comparación
	ejercicio3()  // Operadores lógicos
	ejercicio4()  // Operadores bitwise
	ejercicio5()  // Precedencia
	ejercicio6()  // Sistema de permisos
	ejercicio7()  // Calculadora
	ejercicio8()  // Algoritmos bitwise
	ejercicio9()  // Validación de datos
	ejercicio10() // Optimizaciones

	fmt.Println("\n🎉 ¡Laboratorio completado!")
	fmt.Println("\n💡 Conceptos demostrados:")
	fmt.Println("   ✅ Operadores aritméticos y de asignación")
	fmt.Println("   ✅ Operadores de comparación y lógicos")
	fmt.Println("   ✅ Operadores bitwise y aplicaciones")
	fmt.Println("   ✅ Precedencia y asociatividad")
	fmt.Println("   ✅ Sistema de permisos con bit flags")
	fmt.Println("   ✅ Algoritmos eficientes con bitwise")
	fmt.Println("   ✅ Validación y optimizaciones")
}
