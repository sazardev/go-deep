package main

import (
	"fmt"
	"time"
)

// Constantes para ejercicios
const (
	// Días de la semana
	Domingo = iota
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

const (
	// Tamaños de archivo
	Byte = 1
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
)

// Estados de tarea
const (
	Pendiente = iota
	EnProceso
	Completada
	Cancelada
)

// Permisos con bit flags
const (
	Lectura Permission = 1 << iota
	Escritura
	Ejecucion
	Eliminacion
)

type Permission int

// Constantes de tiempo
const (
	SegundosEnDia      = 60 * 60 * 24
	MilisegundosEnHora = 60 * 60 * 1000
)

// Tipos custom
type TemperaturaType float64
type EstadoType string

const (
	Frio   EstadoType = "frío"
	Tibio  EstadoType = "tibio"
	Calido EstadoType = "cálido"
)

func (t TemperaturaType) ToFahrenheit() TemperaturaType {
	return t*9/5 + 32
}

// Ejercicio 1: Declara variables de diferentes formas
func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Declaraciones de Variables ===")

	// Solución:
	var nombre string  // var sin inicializar
	var edad int = 25  // var con inicialización
	activo := true     // inferencia de tipos
	x, y, z := 1, 2, 3 // asignación múltiple

	fmt.Printf("Nombre: '%s', Edad: %d, Activo: %t\n", nombre, edad, activo)
	fmt.Printf("Coordenadas: x=%d, y=%d, z=%d\n", x, y, z)
}

// Ejercicio 2: Experimenta con scope y shadowing
func ejercicio2() {
	fmt.Println("\n=== Ejercicio 2: Scope y Shadowing ===")

	mensaje := "exterior"

	// Solución:
	fmt.Printf("Mensaje inicial: %s\n", mensaje)

	{
		mensaje := "interior" // Shadow variable
		fmt.Printf("Mensaje en bloque: %s\n", mensaje)
	}

	fmt.Printf("Mensaje después del bloque: %s\n", mensaje)
}

// Ejercicio 3: Trabaja con zero values
func ejercicio3() {
	fmt.Println("\n=== Ejercicio 3: Zero Values ===")

	// Solución:
	var b bool
	var i int
	var f float64
	var s string
	var slice []int
	var m map[string]int
	var ptr *int

	fmt.Printf("bool: %t\n", b)
	fmt.Printf("int: %d\n", i)
	fmt.Printf("float64: %.1f\n", f)
	fmt.Printf("string: '%s'\n", s)
	fmt.Printf("slice: %v (nil: %t)\n", slice, slice == nil)
	fmt.Printf("map: %v (nil: %t)\n", m, m == nil)
	fmt.Printf("pointer: %v (nil: %t)\n", ptr, ptr == nil)
}

// Ejercicio 4: Crea constantes con iota
func ejercicio4() {
	fmt.Println("\n=== Ejercicio 4: Constantes con iota ===")

	// Las constantes están definidas arriba
	fmt.Printf("Miércoles: %d\n", Miercoles)
	fmt.Printf("1 MB: %d bytes\n", MB)
	fmt.Printf("Estado Completada: %d\n", Completada)
}

// Función auxiliar para verificar permisos
func tienePermiso(permisos, permiso Permission) bool {
	return permisos&permiso != 0
}

// Ejercicio 5: Sistema de permisos con bit flags
func ejercicio5() {
	fmt.Println("\n=== Ejercicio 5: Sistema de Permisos ===")

	// Solución:
	usuario := Lectura | Escritura
	moderador := Lectura | Escritura | Eliminacion
	admin := Lectura | Escritura | Ejecucion | Eliminacion

	fmt.Printf("Usuario puede leer: %t\n", tienePermiso(usuario, Lectura))
	fmt.Printf("Usuario puede eliminar: %t\n", tienePermiso(usuario, Eliminacion))
	fmt.Printf("Moderador puede eliminar: %t\n", tienePermiso(moderador, Eliminacion))
	fmt.Printf("Admin puede ejecutar: %t\n", tienePermiso(admin, Ejecucion))
}

// Ejercicio 6: Struct que aprovecha zero values
type Contador struct {
	valor int      // Zero value: 0
	items []string // Zero value: nil slice
}

func (c *Contador) Incrementar() {
	c.valor++
}

func (c *Contador) Decrementar() {
	c.valor--
}

func (c *Contador) Valor() int {
	return c.valor
}

func (c *Contador) AgregarItem(item string) {
	c.items = append(c.items, item)
}

func (c *Contador) Items() []string {
	return c.items
}

func ejercicio6() {
	fmt.Println("\n=== Ejercicio 6: Zero Values en Structs ===")

	// Solución:
	var contador Contador // Zero value funciona perfectamente

	fmt.Printf("Valor inicial: %d\n", contador.Valor())

	contador.Incrementar()
	contador.Incrementar()
	contador.AgregarItem("primer item")
	contador.AgregarItem("segundo item")

	fmt.Printf("Valor después de incrementar: %d\n", contador.Valor())
	fmt.Printf("Items: %v\n", contador.Items())
}

// Ejercicio 7: Constantes complejas con expresiones
func ejercicio7() {
	fmt.Println("\n=== Ejercicio 7: Constantes Complejas ===")

	// Las constantes están definidas arriba
	fmt.Printf("Segundos en un día: %d\n", SegundosEnDia)
	fmt.Printf("Milisegundos en una hora: %d\n", MilisegundosEnHora)
	fmt.Printf("1 GB en bytes: %d\n", GB)

	// Calcular cuántos días de video caben en 1GB (estimando 1MB por minuto)
	minutosEnGB := GB / MB
	horasEnGB := minutosEnGB / 60
	fmt.Printf("Horas de video en 1GB (1MB/min): %d horas\n", horasEnGB)
}

// Ejercicio 8: Manejo de variables temporales
func ejercicio8() {
	fmt.Println("\n=== Ejercicio 8: Variables Temporales ===")

	a, b := 10, 20
	fmt.Printf("Antes del intercambio: a=%d, b=%d\n", a, b)

	// Método 1 - con variable temporal:
	temp := a
	a = b
	b = temp

	fmt.Printf("Después del intercambio 1: a=%d, b=%d\n", a, b)

	// Método 2 - asignación múltiple:
	a, b = b, a

	fmt.Printf("Después del intercambio 2: a=%d, b=%d\n", a, b)
}

// Ejercicio 9: Variables con tipos custom
func ejercicio9() {
	fmt.Println("\n=== Ejercicio 9: Tipos Custom ===")

	// Solución:
	var temp TemperaturaType = 25.0 // Celsius
	estado := Tibio

	fmt.Printf("Temperatura: %.1f°C (%.1f°F)\n", temp, temp.ToFahrenheit())
	fmt.Printf("Estado: %s\n", estado)

	// Diferentes temperaturas
	temperaturas := []TemperaturaType{0, 25, 37, 100}
	for _, t := range temperaturas {
		fmt.Printf("%.0f°C = %.1f°F\n", t, t.ToFahrenheit())
	}
}

// Ejercicio 10: Sistema completo de configuración
type Configuracion struct {
	Puerto   int             // 0 = puerto automático
	Debug    bool            // false = modo producción
	Timeout  time.Duration   // 0 = sin timeout
	Hosts    []string        // nil = todos los hosts
	Features map[string]bool // nil = sin features especiales
	LogLevel string          // "" = nivel por defecto
}

func (c *Configuracion) CargarDefaults() {
	if c.Puerto == 0 {
		c.Puerto = 8080
	}
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}
	if c.LogLevel == "" {
		c.LogLevel = "INFO"
	}
	if c.Hosts == nil {
		c.Hosts = []string{"localhost"}
	}
}

func (c *Configuracion) HabilitarFeature(feature string) {
	if c.Features == nil {
		c.Features = make(map[string]bool)
	}
	c.Features[feature] = true
}

func (c *Configuracion) String() string {
	var features []string
	for feature, enabled := range c.Features {
		if enabled {
			features = append(features, feature)
		}
	}

	return fmt.Sprintf("Config{Puerto: %d, Debug: %t, Timeout: %v, LogLevel: %s, Hosts: %v, Features: %v}",
		c.Puerto, c.Debug, c.Timeout, c.LogLevel, c.Hosts, features)
}

func ejercicio10() {
	fmt.Println("\n=== Ejercicio 10: Sistema de Configuración ===")

	// Solución:
	var config Configuracion // Zero values funcionan perfectamente

	fmt.Printf("Configuración inicial:\n%s\n", config.String())

	// Cargar defaults
	config.CargarDefaults()
	fmt.Printf("\nDespués de cargar defaults:\n%s\n", config.String())

	// Habilitar features
	config.HabilitarFeature("cache")
	config.HabilitarFeature("metrics")
	config.Debug = true

	fmt.Printf("\nConfiguración final:\n%s\n", config.String())
}

func main() {
	fmt.Println("🧪 === LABORATORIO: Variables y Constantes ===\n")

	// Ejecuta todos los ejercicios
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()
	ejercicio9()
	ejercicio10()

	fmt.Println("\n🎉 ¡Laboratorio completado! Todas las soluciones implementadas.")
	fmt.Println("\n💡 Conceptos demostrados:")
	fmt.Println("   ✅ Declaraciones de variables (var, :=, múltiples)")
	fmt.Println("   ✅ Scope y shadowing")
	fmt.Println("   ✅ Zero values efectivos")
	fmt.Println("   ✅ Constantes con iota")
	fmt.Println("   ✅ Bit flags para permisos")
	fmt.Println("   ✅ Structs con zero values inteligentes")
	fmt.Println("   ✅ Tipos custom con métodos")
	fmt.Println("   ✅ Sistema de configuración robusto")
}
