// 🎯 EJERCICIOS PRÁCTICOS - OPERADORES EN GO
// ===========================================
//
// Instrucciones:
// 1. Completa cada función con la implementación correcta
// 2. Ejecuta: go run ejercicios.go para probar tus soluciones
// 3. Todas las funciones deben pasar los tests incluidos
// 4. Lee los comentarios para entender qué se espera

package main

import (
	"fmt"
	"math"
	"strings"
)

// 📅 EJERCICIO 1: Calculadora de Fechas
// ====================================

// esBisiesto determina si un año es bisiesto
// Reglas: Divisible por 4, pero no por 100, excepto si es divisible por 400
func esBisiesto(año int) bool {
	// TODO: Implementa la lógica usando operadores lógicos y aritméticos
	// Pista: Usa % y operadores lógicos (&&, ||)
	return false
}

// diasEnMes retorna el número de días en un mes específico
func diasEnMes(mes, año int) int {
	// TODO: Implementa usando switch/if y la función esBisiesto
	// Febrero tiene 28 días (29 si es bisiesto)
	// Abril, junio, septiembre, noviembre tienen 30 días
	// Los demás tienen 31 días
	return 0
}

// diasEnAño calcula el total de días desde el 1 de enero hasta una fecha
func diasEnAño(dia, mes, año int) int {
	// TODO: Suma los días de los meses anteriores + el día actual
	// Usa la función diasEnMes()
	return 0
}

// diasEntreFechas calcula los días entre dos fechas
func diasEntreFechas(dia1, mes1, año1, dia2, mes2, año2 int) int {
	// TODO: Implementa el cálculo considerando años bisiestos
	// Pista: Calcula días totales desde año 0 para cada fecha y resta
	return 0
}

// 🔐 EJERCICIO 2: Sistema de Permisos con Bitwise
// ===============================================

type Permiso uint8

const (
	Leer     Permiso = 1 << iota // 1
	Escribir                     // 2
	Ejecutar                     // 4
	Eliminar                     // 8
	Admin                        // 16
)

type Usuario struct {
	Nombre   string
	Permisos Permiso
}

// TienePermiso verifica si el usuario tiene un permiso específico
func (u *Usuario) TienePermiso(permiso Permiso) bool {
	// TODO: Usa operadores bitwise para verificar si el bit está activado
	return false
}

// AgregarPermiso añade un permiso al usuario
func (u *Usuario) AgregarPermiso(permiso Permiso) {
	// TODO: Usa OR bitwise para activar el bit
}

// QuitarPermiso remueve un permiso del usuario
func (u *Usuario) QuitarPermiso(permiso Permiso) {
	// TODO: Usa AND NOT bitwise para desactivar el bit
}

// ListarPermisos retorna una lista de permisos activos
func (u *Usuario) ListarPermisos() []string {
	var permisos []string
	// TODO: Verifica cada permiso y añade su nombre si está activo

	return permisos
}

// TienePermisosCompletos verifica si tiene al menos todos los permisos dados
func (u *Usuario) TienePermisosCompletos(permisos Permiso) bool {
	// TODO: Verifica que todos los bits requeridos estén activados
	return false
}

// 🧮 EJERCICIO 3: Evaluador de Expresiones Matemáticas
// ====================================================

// evaluarExpresion evalúa una expresión matemática simple
// Soporta: +, -, *, /, ( )
// Ejemplo: "2 + 3 * 4" = 14, "(2 + 3) * 4" = 20
func evaluarExpresion(expresion string) (float64, error) {
	// TODO: Implementa un parser básico que respete precedencia
	// Pista: Usa recursión o una pila para manejar paréntesis
	// Esta es una versión simplificada - puedes usar strings.Fields() para separar tokens

	// Eliminar espacios
	expresion = strings.ReplaceAll(expresion, " ", "")

	// TODO: Tu implementación aquí
	// Puedes empezar con casos simples sin paréntesis

	return 0, fmt.Errorf("no implementado")
}

// calcularOperacion realiza una operación básica
func calcularOperacion(a, b float64, operador string) (float64, error) {
	// TODO: Implementa usando switch y operadores aritméticos
	switch operador {
	case "+":
		// TODO
		return 0, nil
	case "-":
		// TODO
		return 0, nil
	case "*":
		// TODO
		return 0, nil
	case "/":
		// TODO: Verificar división por cero
		return 0, nil
	default:
		return 0, fmt.Errorf("operador desconocido: %s", operador)
	}
}

// 🔢 EJERCICIO 4: Manipulación de Bits
// ====================================

// contarBitsActivos cuenta cuántos bits están en 1
func contarBitsActivos(n uint32) int {
	// TODO: Implementa usando operadores bitwise
	// Pista: Usa n & 1 para verificar el bit menos significativo
	// y n >> 1 para desplazar
	return 0
}

// esPotenciaDe2 verifica si un número es potencia de 2
func esPotenciaDe2(n uint32) bool {
	// TODO: Una potencia de 2 tiene exactamente un bit activado
	// Pista: n & (n-1) == 0 para potencias de 2 (excepto 0)
	return false
}

// obtenerBit retorna el valor del bit en la posición dada (0-indexado)
func obtenerBit(n uint32, posicion int) bool {
	// TODO: Usa desplazamiento y AND para obtener el bit
	return false
}

// establecerBit activa el bit en la posición dada
func establecerBit(n uint32, posicion int) uint32 {
	// TODO: Usa OR para activar el bit
	return 0
}

// limpiarBit desactiva el bit en la posición dada
func limpiarBit(n uint32, posicion int) uint32 {
	// TODO: Usa AND NOT para desactivar el bit
	return 0
}

// alternarBit cambia el estado del bit en la posición dada
func alternarBit(n uint32, posicion int) uint32 {
	// TODO: Usa XOR para alternar el bit
	return 0
}

// intercambiarBits intercambia los bits en dos posiciones
func intercambiarBits(n uint32, pos1, pos2 int) uint32 {
	// TODO: Obtén los bits, verifica si son diferentes, y intercámbia si es necesario
	return 0
}

// encontrarBitMasSignificativo encuentra la posición del bit más alto activado
func encontrarBitMasSignificativo(n uint32) int {
	// TODO: Implementa usando desplazamientos o logaritmos
	// Retorna -1 si n es 0
	return -1
}

// 💰 EJERCICIO 5: Calculadora de Interés Compuesto
// ================================================

// calcularInteresCompuesto calcula el monto final con interés compuesto
// Formula: A = P(1 + r/n)^(nt)
// P = principal, r = tasa anual, n = compounding frequency, t = tiempo en años
func calcularInteresCompuesto(principal, tasaAnual float64, compoundingFreq int, años float64) float64 {
	// TODO: Implementa la fórmula usando math.Pow()
	// Verifica que los valores sean positivos
	return 0
}

// compararInversiones compara dos opciones de inversión
func compararInversiones(principal1, tasa1 float64, freq1 int, años1 float64,
	principal2, tasa2 float64, freq2 int, años2 float64) (mejor int, diferencia float64) {
	// TODO: Calcula ambas inversiones y retorna cuál es mejor (1 o 2) y la diferencia
	// Usa la función calcularInteresCompuesto()
	return 0, 0
}

// 🎲 EJERCICIO 6: Simulador de Dados y Probabilidades
// ===================================================

// simularTiradaDados simula tirar n dados de 6 caras
func simularTiradaDados(numDados int, semilla int64) []int {
	// TODO: Implementa usando math/rand
	// Usa la semilla para resultados reproducibles
	// Cada dado debe dar un valor entre 1 y 6
	return nil
}

// calcularProbabilidadSuma calcula la probabilidad teórica de obtener una suma específica
func calcularProbabilidadSuma(numDados, sumaObjetivo int) float64 {
	// TODO: Calcula usando combinatoria
	// Esta es una función avanzada - puedes implementar una versión simple
	// que cuente todas las combinaciones posibles
	return 0.0
}

// 🧪 TESTS Y DEMOSTRACIÓN
// =======================

func mainEjercicios() {
	fmt.Println("🎯 EJERCICIOS DE OPERADORES - TESTS")
	fmt.Println("===================================")

	// Test Ejercicio 1: Fechas
	fmt.Println("\n📅 Test 1: Calculadora de Fechas")
	testearFechas()

	// Test Ejercicio 2: Permisos
	fmt.Println("\n🔐 Test 2: Sistema de Permisos")
	testearPermisos()

	// Test Ejercicio 3: Expresiones
	fmt.Println("\n🧮 Test 3: Evaluador de Expresiones")
	testearExpresiones()

	// Test Ejercicio 4: Bits
	fmt.Println("\n🔢 Test 4: Manipulación de Bits")
	testearBits()

	// Test Ejercicio 5: Interés
	fmt.Println("\n💰 Test 5: Interés Compuesto")
	testearInteres()

	// Test Ejercicio 6: Dados
	fmt.Println("\n🎲 Test 6: Simulador de Dados")
	testearDados()
}

func testearFechas() {
	// Tests para años bisiestos
	casos := []struct {
		año      int
		esperado bool
	}{
		{2000, true},  // Divisible por 400
		{1900, false}, // Divisible por 100 pero no por 400
		{2004, true},  // Divisible por 4 pero no por 100
		{2001, false}, // No divisible por 4
	}

	for _, caso := range casos {
		resultado := esBisiesto(caso.año)
		status := "❌"
		if resultado == caso.esperado {
			status = "✅"
		}
		fmt.Printf("%s Año %d bisiesto: %t (esperado: %t)\n",
			status, caso.año, resultado, caso.esperado)
	}

	// Test días en febrero
	diasFeb2020 := diasEnMes(2, 2020)
	diasFeb2021 := diasEnMes(2, 2021)
	fmt.Printf("Febrero 2020: %d días (esperado: 29)\n", diasFeb2020)
	fmt.Printf("Febrero 2021: %d días (esperado: 28)\n", diasFeb2021)
}

func testearPermisos() {
	usuario := Usuario{Nombre: "Juan", Permisos: 0}

	// Agregar permisos
	usuario.AgregarPermiso(Leer)
	usuario.AgregarPermiso(Escribir)

	fmt.Printf("Usuario: %s\n", usuario.Nombre)
	fmt.Printf("Puede leer: %t\n", usuario.TienePermiso(Leer))
	fmt.Printf("Puede escribir: %t\n", usuario.TienePermiso(Escribir))
	fmt.Printf("Puede ejecutar: %t\n", usuario.TienePermiso(Ejecutar))

	// Listar permisos
	permisos := usuario.ListarPermisos()
	fmt.Printf("Permisos activos: %v\n", permisos)
}

func testearExpresiones() {
	expresiones := []struct {
		expr     string
		esperado float64
	}{
		{"2 + 3", 5},
		{"2 * 3 + 4", 10},
		{"2 + 3 * 4", 14},
		{"(2 + 3) * 4", 20},
		{"10 / 2", 5},
	}

	for _, expr := range expresiones {
		resultado, err := evaluarExpresion(expr.expr)
		status := "❌"
		if err == nil && math.Abs(resultado-expr.esperado) < 0.0001 {
			status = "✅"
		}
		fmt.Printf("%s %s = %.1f (esperado: %.1f)\n",
			status, expr.expr, resultado, expr.esperado)
	}
}

func testearBits() {
	// Test contar bits
	fmt.Printf("Bits activos en 7 (111): %d (esperado: 3)\n", contarBitsActivos(7))
	fmt.Printf("Bits activos en 15 (1111): %d (esperado: 4)\n", contarBitsActivos(15))

	// Test potencia de 2
	fmt.Printf("8 es potencia de 2: %t (esperado: true)\n", esPotenciaDe2(8))
	fmt.Printf("10 es potencia de 2: %t (esperado: false)\n", esPotenciaDe2(10))

	// Test manipulación de bits
	n := uint32(5) // 101 en binario
	fmt.Printf("Número original: %d (%08b)\n", n, n)

	n = establecerBit(n, 1) // Debería ser 111 (7)
	fmt.Printf("Después de activar bit 1: %d (%08b)\n", n, n)

	n = limpiarBit(n, 2) // Debería ser 011 (3)
	fmt.Printf("Después de limpiar bit 2: %d (%08b)\n", n, n)
}

func testearInteres() {
	// Inversión de $1000 al 5% anual, capitalización mensual, por 2 años
	resultado := calcularInteresCompuesto(1000, 0.05, 12, 2)
	fmt.Printf("$1000 al 5%% por 2 años: $%.2f\n", resultado)

	// Comparar dos inversiones
	mejor, diff := compararInversiones(1000, 0.05, 12, 2, 1000, 0.04, 4, 2)
	fmt.Printf("Mejor inversión: %d, diferencia: $%.2f\n", mejor, diff)
}

func testearDados() {
	dados := simularTiradaDados(2, 42) // Semilla fija para reproducibilidad
	fmt.Printf("Tirada de 2 dados: %v\n", dados)
	if len(dados) == 2 {
		suma := dados[0] + dados[1]
		fmt.Printf("Suma: %d\n", suma)
	}

	// Probabilidad teórica
	prob := calcularProbabilidadSuma(2, 7)
	fmt.Printf("Probabilidad de suma 7 con 2 dados: %.4f\n", prob)
}

// 💡 PISTAS Y AYUDAS
// ==================

/*
PISTAS PARA LOS EJERCICIOS:

📅 Ejercicio 1 (Fechas):
- Año bisiesto: (año % 4 == 0 && año % 100 != 0) || (año % 400 == 0)
- Usa arrays o maps para los días de cada mes
- Para calcular días entre fechas, convierte todo a días desde una fecha base

🔐 Ejercicio 2 (Permisos):
- TienePermiso: return (u.Permisos & permiso) != 0
- AgregarPermiso: u.Permisos |= permiso
- QuitarPermiso: u.Permisos &^= permiso

🧮 Ejercicio 3 (Expresiones):
- Empezar sin paréntesis, solo +, -, *, /
- Usar strings.Fields() o regex para parsear tokens
- Implementar precedencia: primero *, / luego +, -

🔢 Ejercicio 4 (Bits):
- ContarBits: usar n & 1 y n >>= 1 en un loop
- EsPotenciaDe2: return n > 0 && (n & (n-1)) == 0
- ObtenerBit: return (n & (1 << posicion)) != 0

💰 Ejercicio 5 (Interés):
- Usar math.Pow(1 + r/n, n*t)
- Verificar inputs válidos (positivos)

🎲 Ejercicio 6 (Dados):
- Usar rand.Seed() y rand.Intn(6) + 1
- Para probabilidades, contar combinaciones posibles
*/
