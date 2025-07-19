// 🎯 EJERCICIOS PRÁCTICOS - FUNCIONES EN GO
// =========================================
//
// Instrucciones:
// 1. Implementa cada función según las especificaciones
// 2. Ejecuta: go run ejercicios.go para probar tus soluciones
// 3. Todas las funciones deben compilar sin errores
// 4. Lee los comentarios para entender qué se espera

package main

import (
	"fmt"
) // 📐 EJERCICIO 1: Funciones Básicas
// ==================================

// calcularArea calcula el área de un círculo dado su radio
func calcularArea(radio float64) float64 {
	// TODO: Implementa usando π * r²
	// Puedes usar 3.14159 como aproximación de π
	return 0 // Reemplaza con tu implementación
}

// esParImpar determina si un número es par
func esParImpar(numero int) bool {
	// TODO: Usa el operador módulo (%)
	return false // Reemplaza con tu implementación
}

// crearSaludo crea un saludo personalizado
func crearSaludo(nombre, apellido string, edad int) string {
	// TODO: Retorna un string como "Hola, Juan Pérez (25 años)"
	return "" // Reemplaza con tu implementación
}

// buscarMaximo encuentra el número mayor entre tres números
func buscarMaximo(a, b, c int) int {
	// TODO: Compara los tres números y retorna el mayor
	return 0 // Reemplaza con tu implementación
}

// 🔄 EJERCICIO 2: Múltiples Valores de Retorno
// =============================================

// dividirEnteros realiza división entera y retorna cociente y resto
func dividirEnteros(dividendo, divisor int) (int, int, error) {
	// TODO: Maneja la división por cero retornando un error
	// Retorna cociente, resto, error
	return 0, 0, nil // Reemplaza con tu implementación
}

// contarCaracteres analiza un texto y retorna estadísticas
func contarCaracteres(texto string) (palabras int, caracteres int, vocales int) {
	// TODO: Cuenta palabras, caracteres y vocales
	// Considera a, e, i, o, u como vocales (mayúsculas y minúsculas)
	return // Retorno naked - usa las variables nombradas
}

// cambiarTemperatura convierte entre Celsius y Fahrenheit
func cambiarTemperatura(grados float64, escala string) (float64, string, error) {
	// TODO: Si escala es "C", convierte a Fahrenheit
	// Si escala es "F", convierte a Celsius
	// Retorna: temperatura convertida, escala resultante, error si la escala es inválida
	return 0, "", nil // Reemplaza con tu implementación
}

// 📊 EJERCICIO 3: Funciones Variádicas
// ====================================

// sacarPromedio calcula el promedio de números dados
func sacarPromedio(numeros ...float64) float64 {
	// TODO: Si no hay números, retorna 0
	// Calcula la suma de todos los números y divide por la cantidad
	return 0 // Reemplaza con tu implementación
}

// unirTextos une strings con un separador personalizado
func unirTextos(separador string, textos ...string) string {
	// TODO: Une todos los textos usando el separador
	// Similar a strings.Join pero implementado por ti
	return "" // Reemplaza con tu implementación
}

// buscarEnRango busca números dentro de un rango específico
func buscarEnRango(min, max int, numeros ...int) []int {
	// TODO: Retorna slice con números que están entre min y max (inclusive)
	return nil // Reemplaza con tu implementación
}

// 🧮 EJERCICIO 4: Funciones de Primera Clase
// ===========================================

// Operacion define una función matemática
type Operacion func(float64, float64) float64

// hacerCalculadora retorna una función que aplica operaciones
func hacerCalculadora(operacion Operacion) func(float64, float64) string {
	// TODO: Retorna una función que:
	// 1. Aplica la operación a dos números
	// 2. Retorna el resultado formateado como string "a op b = resultado"
	return nil // Reemplaza con tu implementación
}

// procesarSlice aplica una función a todos los elementos de un slice
func procesarSlice(numeros []int, fn func(int) int) []int {
	// TODO: Crea un nuevo slice aplicando fn a cada elemento
	return nil // Reemplaza con tu implementación
}

// seleccionarElementos filtra elementos de un slice según un predicado
func seleccionarElementos(numeros []int, predicado func(int) bool) []int {
	// TODO: Retorna nuevo slice con elementos que cumplen el predicado
	return nil // Reemplaza con tu implementación
}

// 🔒 EJERCICIO 5: Closures
// ========================

// hacerContador crea un contador que incrementa en cada llamada
func hacerContador(inicial int) func() int {
	// TODO: Retorna función que incrementa y retorna un contador
	// Debe mantener estado entre llamadas
	return nil // Reemplaza con tu implementación
}

// hacerValidador crea un validador para rangos específicos
func hacerValidador(min, max int) func(int) bool {
	// TODO: Retorna función que valida si un número está en el rango
	return nil // Reemplaza con tu implementación
}

// hacerSumador crea una función que acumula valores
func hacerSumador() (func(int), func() int) {
	// TODO: Retorna dos funciones:
	// 1. agregar(n) - suma n al acumulador
	// 2. obtener() - retorna el valor actual
	return nil, nil // Reemplaza con tu implementación
}

func main() {
	fmt.Println("🎯 EJERCICIOS DE FUNCIONES - IMPLEMENTACIÓN")
	fmt.Println("==========================================")

	// Ejercicio 1: Funciones básicas
	fmt.Println("\n📐 Ejercicio 1: Funciones Básicas")
	pruebaEjercicio1()

	// Ejercicio 2: Múltiples retornos
	fmt.Println("\n🔄 Ejercicio 2: Múltiples Valores de Retorno")
	pruebaEjercicio2()

	// Ejercicio 3: Funciones variádicas
	fmt.Println("\n📊 Ejercicio 3: Funciones Variádicas")
	pruebaEjercicio3()

	// Ejercicio 4: Funciones de primera clase
	fmt.Println("\n🧮 Ejercicio 4: Funciones de Primera Clase")
	pruebaEjercicio4()

	// Ejercicio 5: Closures
	fmt.Println("\n🔒 Ejercicio 5: Closures")
	pruebaEjercicio5()
}

// ============================================================================
// FUNCIONES DE PRUEBA
// ============================================================================

func pruebaEjercicio1() {
	// Probar área del círculo
	area := calcularArea(5.0)
	fmt.Printf("Área círculo radio 5: %.2f (esperado: ~78.54)\n", area)

	// Probar es par
	fmt.Printf("4 es par: %t (esperado: true)\n", esParImpar(4))
	fmt.Printf("7 es par: %t (esperado: false)\n", esParImpar(7))

	// Probar saludo
	saludo := crearSaludo("Ana", "García", 25)
	fmt.Printf("Saludo: %s\n", saludo)

	// Probar máximo
	max := buscarMaximo(3, 7, 5)
	fmt.Printf("Máximo de 3, 7, 5: %d (esperado: 7)\n", max)
}

func pruebaEjercicio2() {
	// Probar división
	if cociente, resto, err := dividirEnteros(17, 5); err == nil {
		fmt.Printf("17 ÷ 5 = %d resto %d\n", cociente, resto)
	}

	// Probar análisis de texto
	palabras, chars, vocales := contarCaracteres("Hola mundo")
	fmt.Printf("'Hola mundo': %d palabras, %d caracteres, %d vocales\n", palabras, chars, vocales)

	// Probar conversión de temperatura
	if temp, escala, err := cambiarTemperatura(100, "C"); err == nil {
		fmt.Printf("100°C = %.1f°%s\n", temp, escala)
	}
}

func pruebaEjercicio3() {
	// Probar promedio
	prom := sacarPromedio(1, 2, 3, 4, 5)
	fmt.Printf("Promedio de 1,2,3,4,5: %.1f (esperado: 3.0)\n", prom)

	// Probar concatenación
	resultado := unirTextos(" - ", "A", "B", "C")
	fmt.Printf("Concatenar con ' - ': %s\n", resultado)

	// Probar filtro por rango
	numeros := buscarEnRango(3, 7, 1, 4, 8, 5, 9, 6, 2)
	fmt.Printf("Números entre 3 y 7: %v\n", numeros)
}

func pruebaEjercicio4() {
	// Crear operaciones
	suma := func(a, b float64) float64 { return a + b }

	// Probar calculadora
	calcSuma := hacerCalculadora(suma)
	if calcSuma != nil {
		fmt.Printf("Calculadora suma: %s\n", calcSuma(5, 3))
	}

	// Probar aplicar a slice
	numeros := []int{1, 2, 3, 4, 5}
	duplicar := func(x int) int { return x * 2 }
	duplicados := procesarSlice(numeros, duplicar)
	fmt.Printf("Duplicados: %v\n", duplicados)

	// Probar filtro
	esPar := func(x int) bool { return x%2 == 0 }
	pares := seleccionarElementos(numeros, esPar)
	fmt.Printf("Pares: %v\n", pares)
}

func pruebaEjercicio5() {
	// Probar contador
	contador := hacerContador(0)
	if contador != nil {
		fmt.Printf("Contador: %d, %d, %d\n", contador(), contador(), contador())
	}

	// Probar validador
	validarEdad := hacerValidador(18, 65)
	if validarEdad != nil {
		fmt.Printf("Edad 25 válida: %t\n", validarEdad(25))
		fmt.Printf("Edad 10 válida: %t\n", validarEdad(10))
	}

	// Probar acumulador
	agregar, obtener := hacerSumador()
	if agregar != nil && obtener != nil {
		agregar(5)
		agregar(3)
		fmt.Printf("Acumulador: %d\n", obtener())
	}
}
