// Proyecto: Calculadora Avanzada con Operadores
package main

import (
	"fmt"
	"math"
)

// Calculadora que demuestra todos los operadores
type CalculadoraAvanzada struct {
	memoria   float64
	historial []string
}

func NewCalculadora() *CalculadoraAvanzada {
	return &CalculadoraAvanzada{
		memoria:   0,
		historial: make([]string, 0),
	}
}

// Operaciones aritméticas
func (c *CalculadoraAvanzada) Sumar(a, b float64) float64 {
	result := a + b
	c.agregarHistorial(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

func (c *CalculadoraAvanzada) Restar(a, b float64) float64 {
	result := a - b
	c.agregarHistorial(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

func (c *CalculadoraAvanzada) Multiplicar(a, b float64) float64 {
	result := a * b
	c.agregarHistorial(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
	return result
}

func (c *CalculadoraAvanzada) Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	result := a / b
	c.agregarHistorial(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
	return result, nil
}

// Operaciones bitwise para enteros
func (c *CalculadoraAvanzada) AND(a, b int) int {
	result := a & b
	c.agregarHistorial(fmt.Sprintf("%d & %d = %d", a, b, result))
	return result
}

func (c *CalculadoraAvanzada) OR(a, b int) int {
	result := a | b
	c.agregarHistorial(fmt.Sprintf("%d | %d = %d", a, b, result))
	return result
}

func (c *CalculadoraAvanzada) XOR(a, b int) int {
	result := a ^ b
	c.agregarHistorial(fmt.Sprintf("%d ^ %d = %d", a, b, result))
	return result
}

// Operaciones de comparación
func (c *CalculadoraAvanzada) Comparar(a, b float64) map[string]bool {
	result := map[string]bool{
		"igual":       c.sonIguales(a, b),
		"mayor":       a > b,
		"menor":       a < b,
		"mayor_igual": a >= b,
		"menor_igual": a <= b,
		"diferente":   !c.sonIguales(a, b),
	}
	c.agregarHistorial(fmt.Sprintf("Comparación de %.2f y %.2f", a, b))
	return result
}

func (c *CalculadoraAvanzada) sonIguales(a, b float64) bool {
	epsilon := 1e-9
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}

// Operaciones de memoria
func (c *CalculadoraAvanzada) GuardarMemoria(valor float64) {
	c.memoria = valor
	c.agregarHistorial(fmt.Sprintf("Memoria: %.2f", valor))
}

func (c *CalculadoraAvanzada) RecuperarMemoria() float64 {
	return c.memoria
}

// Funciones matemáticas avanzadas
func (c *CalculadoraAvanzada) Potencia(base, exponente float64) float64 {
	result := math.Pow(base, exponente)
	c.agregarHistorial(fmt.Sprintf("%.2f^%.2f = %.2f", base, exponente, result))
	return result
}

func (c *CalculadoraAvanzada) RaizCuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("no se puede calcular raíz cuadrada de número negativo")
	}
	result := math.Sqrt(n)
	c.agregarHistorial(fmt.Sprintf("√%.2f = %.2f", n, result))
	return result, nil
}

// Gestión de historial
func (c *CalculadoraAvanzada) agregarHistorial(operacion string) {
	c.historial = append(c.historial, operacion)
	if len(c.historial) > 5 { // Mantener solo las últimas 5
		c.historial = c.historial[1:]
	}
}

func (c *CalculadoraAvanzada) MostrarHistorial() {
	fmt.Println("\n=== Historial de Operaciones ===")
	for i, op := range c.historial {
		fmt.Printf("%d. %s\n", i+1, op)
	}
}

// Sistema de permisos con bitwise
type PermisoCalc uint8

const (
	Calcular  PermisoCalc = 1 << iota // 1
	Memoria                           // 2
	Avanzado                          // 4
	AdminCalc                         // 8
)

type UsuarioCalculadora struct {
	Nombre   string
	Permisos PermisoCalc
}

func (u *UsuarioCalculadora) TienePermiso(p PermisoCalc) bool {
	return (u.Permisos & p) != 0
}

func (u *UsuarioCalculadora) AgregarPermiso(p PermisoCalc) {
	u.Permisos |= p
}

// Demostración principal
func main() {
	fmt.Println("🎯 PROYECTO: CALCULADORA AVANZADA CON OPERADORES")
	fmt.Println("===============================================")

	// Crear calculadora
	calc := NewCalculadora()

	// Demostrar operaciones aritméticas
	fmt.Println("\n--- Operaciones Aritméticas ---")
	fmt.Printf("Suma: %.2f\n", calc.Sumar(15.5, 4.3))
	fmt.Printf("Resta: %.2f\n", calc.Restar(20.0, 5.5))
	fmt.Printf("Multiplicación: %.2f\n", calc.Multiplicar(3.5, 2.0))

	if div, err := calc.Dividir(10.0, 3.0); err == nil {
		fmt.Printf("División: %.2f\n", div)
	}

	// Demostrar operaciones bitwise
	fmt.Println("\n--- Operaciones Bitwise ---")
	fmt.Printf("12 & 10 = %d\n", calc.AND(12, 10))
	fmt.Printf("12 | 10 = %d\n", calc.OR(12, 10))
	fmt.Printf("12 ^ 10 = %d\n", calc.XOR(12, 10))

	// Demostrar comparaciones
	fmt.Println("\n--- Comparaciones ---")
	comp := calc.Comparar(7.5, 7.5)
	for operacion, resultado := range comp {
		fmt.Printf("%s: %t\n", operacion, resultado)
	}

	// Demostrar operaciones avanzadas
	fmt.Println("\n--- Operaciones Avanzadas ---")
	fmt.Printf("2^8 = %.0f\n", calc.Potencia(2, 8))
	if sqrt, err := calc.RaizCuadrada(16); err == nil {
		fmt.Printf("√16 = %.2f\n", sqrt)
	}

	// Demostrar memoria
	fmt.Println("\n--- Memoria ---")
	calc.GuardarMemoria(42.0)
	fmt.Printf("Valor en memoria: %.2f\n", calc.RecuperarMemoria())

	// Mostrar historial
	calc.MostrarHistorial()

	// Demostrar sistema de permisos
	fmt.Println("\n--- Sistema de Permisos ---")
	usuario := UsuarioCalculadora{Nombre: "Ana", Permisos: 0}
	usuario.AgregarPermiso(Calcular)
	usuario.AgregarPermiso(Memoria)

	fmt.Printf("Usuario: %s\n", usuario.Nombre)
	fmt.Printf("Puede calcular: %t\n", usuario.TienePermiso(Calcular))
	fmt.Printf("Puede usar memoria: %t\n", usuario.TienePermiso(Memoria))
	fmt.Printf("Tiene permisos avanzados: %t\n", usuario.TienePermiso(Avanzado))

	// Demostrar precedencia de operadores
	fmt.Println("\n--- Precedencia de Operadores ---")
	a, b, c := 2, 3, 4
	resultado1 := a + b*c     // 2 + (3*4) = 14
	resultado2 := (a + b) * c // (2+3) * 4 = 20

	fmt.Printf("a + b * c = %d + %d * %d = %d\n", a, b, c, resultado1)
	fmt.Printf("(a + b) * c = (%d + %d) * %d = %d\n", a, b, c, resultado2)

	// Demostrar operadores lógicos
	fmt.Println("\n--- Operadores Lógicos ---")
	edad := 25
	tienePermiso := true
	puedeAcceder := edad >= 18 && tienePermiso
	fmt.Printf("Edad: %d, Tiene permiso: %t\n", edad, tienePermiso)
	fmt.Printf("Puede acceder: %t\n", puedeAcceder)

	fmt.Println("\n¡Calculadora avanzada completada! 🎉")
}
