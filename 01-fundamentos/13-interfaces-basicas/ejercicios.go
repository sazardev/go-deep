// Lección 13: Interfaces Básicas en Go
// Ejercicios prácticos para dominar las interfaces

package main

import "fmt"

// ========================================
// Ejercicio 1: Interface Básica - Sistema de Formas Geométricas
// ========================================

// TODO: Define una interface llamada "Forma" que tenga los métodos:
// - Area() float64
// - Perimetro() float64
// - Descripcion() string

type Forma interface {
	// TODO: Implementar métodos de la interface
}

// TODO: Crea un struct "Rectangulo" con campos Ancho y Alto
type Rectangulo struct {
	// TODO: Agregar campos
}

// TODO: Implementa los métodos de la interface Forma para Rectangulo

// TODO: Crea un struct "Circulo" con campo Radio
type Circulo struct {
	// TODO: Agregar campos
}

// TODO: Implementa los métodos de la interface Forma para Circulo

// TODO: Crea una función que reciba una Forma y muestre toda su información
func MostrarInformacionForma(f Forma) {
	// TODO: Implementar
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Sistema de Formas Geométricas ===")

	// TODO: Crear instancias de Rectangulo y Circulo
	// TODO: Usar la función MostrarInformacionForma con ambas

	fmt.Println("✅ Ejercicio 1 completado\n")
}

// ========================================
// Ejercicio 2: Polimorfismo - Sistema de Animales
// ========================================

// TODO: Define una interface "Animal" con métodos:
// - HacerSonido() string
// - Moverse() string
// - Comer(comida string) string

type Animal interface {
	// TODO: Implementar métodos
}

// TODO: Crea structs para diferentes animales: Perro, Gato, Pajaro
// Cada uno debe implementar la interface Animal de manera diferente

type Perro struct {
	Nombre string
	Raza   string
}

type Gato struct {
	Nombre string
	Color  string
}

type Pajaro struct {
	Nombre     string
	Especie    string
	PuedeVolar bool
}

// TODO: Implementar métodos Animal para cada struct

// TODO: Crea una función "CuidarAnimal" que reciba un Animal
// y simule darle de comer y pedirle que haga sonido
func CuidarAnimal(a Animal) {
	// TODO: Implementar
}

func ejercicio2() {
	fmt.Println("=== Ejercicio 2: Sistema de Animales ===")

	// TODO: Crear diferentes animales y usar CuidarAnimal con cada uno

	fmt.Println("✅ Ejercicio 2 completado\n")
}

// ========================================
// Ejercicio 3: Interfaces Estándar - fmt.Stringer y sort.Interface
// ========================================

// TODO: Crea un struct "Producto" con campos: Nombre, Precio, Categoria
type Producto struct {
	// TODO: Agregar campos
}

// TODO: Implementa fmt.Stringer para Producto
// El String() debe retornar algo como: "Laptop ($1200.00) - Electrónicos"

// TODO: Crea un tipo personalizado "ProductosPorPrecio" que sea un slice de Producto
type ProductosPorPrecio []Producto

// TODO: Implementa sort.Interface para ProductosPorPrecio
// (Len, Less, Swap) para ordenar por precio de menor a mayor

func ejercicio3() {
	fmt.Println("=== Ejercicio 3: Interfaces Estándar ===")

	// TODO: Crear varios productos
	// TODO: Mostrar productos usando fmt.Println (usará String())
	// TODO: Ordenar productos por precio usando sort.Sort
	// TODO: Mostrar productos ordenados

	fmt.Println("✅ Ejercicio 3 completado\n")
}

// ========================================
// Ejercicio 4: Type Assertions - Procesador de Datos
// ========================================

// TODO: Crea una función "ProcesarDato" que reciba interface{} y:
// - Si es string: muestre la longitud y el texto en mayúsculas
// - Si es int: muestre si es par/impar y su cuadrado
// - Si es float64: muestre la raíz cuadrada y si es mayor a 100
// - Si es []int: muestre la suma y el promedio
// - Si es map[string]int: muestre las claves y la suma de valores
// - Para cualquier otro tipo: muestre el tipo y valor

func ProcesarDato(dato interface{}) {
	// TODO: Implementar usando type switch
}

func ejercicio4() {
	fmt.Println("=== Ejercicio 4: Type Assertions ===")

	// TODO: Crear slice con diferentes tipos de datos
	// TODO: Procesar cada dato con ProcesarDato

	datos := []interface{}{
		"Hola Go",
		42,
		3.14159,
		[]int{1, 2, 3, 4, 5},
		map[string]int{"a": 10, "b": 20, "c": 30},
		true,
	}

	for i, dato := range datos {
		fmt.Printf("Dato %d: ", i+1)
		ProcesarDato(dato)
	}

	fmt.Println("✅ Ejercicio 4 completado\n")
}

// ========================================
// Ejercicio 5: Empty Interface - Sistema de Logging
// ========================================

// TODO: Define una interface "Logger" con método:
// - Log(level string, message string, data interface{})

type Logger interface {
	// TODO: Implementar
}

// TODO: Crea struct "ConsoleLogger" que implemente Logger
// Debe imprimir a consola con formato: [NIVEL] mensaje: datos
type ConsoleLogger struct{}

// TODO: Crea struct "FileLogger" que simule escribir a archivo
type FileLogger struct {
	archivo string
}

// TODO: Crea una función "LogearEvento" que reciba un Logger y varios datos
func LogearEvento(logger Logger, evento string, datos ...interface{}) {
	// TODO: Usar el logger para registrar el evento con todos los datos
}

func ejercicio5() {
	fmt.Println("=== Ejercicio 5: Sistema de Logging ===")

	// TODO: Crear instancias de ConsoleLogger y FileLogger
	// TODO: Usar LogearEvento con diferentes tipos de datos

	fmt.Println("✅ Ejercicio 5 completado\n")
}

// ========================================
// Ejercicio 6: Strategy Pattern - Sistema de Descuentos
// ========================================

// TODO: Define interface "EstrategiaDescuento" con método:
// - AplicarDescuento(precio float64) float64
// - DescribirDescuento() string

type EstrategiaDescuento interface {
	// TODO: Implementar
}

// TODO: Crea diferentes estrategias de descuento:
// - SinDescuento: no aplica descuento
// - DescuentoPorcentaje: descuento por porcentaje
// - DescuentoFijo: descuento de cantidad fija
// - DescuentoPorCantidad: descuento basado en cantidad de items

type SinDescuento struct{}

type DescuentoPorcentaje struct {
	Porcentaje float64
}

type DescuentoFijo struct {
	Cantidad float64
}

type DescuentoPorCantidad struct {
	CantidadMinima int
	Descuento      float64
}

// TODO: Implementar las estrategias

// TODO: Crea struct "CalculadoraPrecio" que use una estrategia
type CalculadoraPrecio struct {
	estrategia EstrategiaDescuento
}

// TODO: Implementar métodos para cambiar estrategia y calcular precio final

func ejercicio6() {
	fmt.Println("=== Ejercicio 6: Strategy Pattern - Descuentos ===")

	// TODO: Crear calculadora y probar diferentes estrategias
	// TODO: Mostrar precio original y final con descripción del descuento

	fmt.Println("✅ Ejercicio 6 completado\n")
}

// ========================================
// Ejercicio 7: Observer Pattern - Sistema de Notificaciones
// ========================================

// TODO: Define interface "Observer" con métodos:
// - Actualizar(evento string, datos interface{})
// - ObtenerID() string

type Observer interface {
	// TODO: Implementar
}

// TODO: Define interface "Sujeto" para objetos observables con métodos:
// - Suscribir(observer Observer)
// - Desuscribir(id string)
// - Notificar(evento string, datos interface{})

type Sujeto interface {
	// TODO: Implementar
}

// TODO: Crea struct "GestorEventos" que implemente Sujeto
type GestorEventos struct {
	observadores map[string]Observer
}

// TODO: Implementar métodos de Sujeto

// TODO: Crea diferentes tipos de observadores:
// - ObservadorEmail: simula envío de emails
// - ObservadorSMS: simula envío de SMS
// - ObservadorAnalytics: simula registro de analytics

type ObservadorEmail struct {
	email string
}

type ObservadorSMS struct {
	telefono string
}

type ObservadorAnalytics struct {
	servicio string
}

// TODO: Implementar Observer para cada tipo

func ejercicio7() {
	fmt.Println("=== Ejercicio 7: Observer Pattern ===")

	// TODO: Crear gestor de eventos y diferentes observadores
	// TODO: Suscribir observadores
	// TODO: Generar diferentes eventos
	// TODO: Desuscribir un observador y generar otro evento

	fmt.Println("✅ Ejercicio 7 completado\n")
}

// ========================================
// Ejercicio 8: Factory Pattern - Conexiones de Base de Datos
// ========================================

// TODO: Define interface "BaseDatos" con métodos:
// - Conectar() string
// - EjecutarConsulta(sql string) string
// - Cerrar() string
// - ObtenerTipo() string

type BaseDatos interface {
	// TODO: Implementar
}

// TODO: Crea implementaciones para diferentes bases de datos:
// - MySQL, PostgreSQL, MongoDB
// Cada una debe comportarse de manera diferente

type MySQL struct {
	host   string
	puerto int
}

type PostgreSQL struct {
	host      string
	baseDatos string
}

type MongoDB struct {
	uri       string
	coleccion string
}

// TODO: Implementar BaseDatos para cada tipo

// TODO: Define interface "FactoryBaseDatos" con método:
// - CrearBaseDatos(tipo string, config map[string]string) BaseDatos

type FactoryBaseDatos interface {
	// TODO: Implementar
}

// TODO: Crea struct "FactoryDB" que implemente FactoryBaseDatos
type FactoryDB struct{}

// TODO: Implementar el factory

// TODO: Crea función "ProbarBaseDatos" que use el factory y pruebe la BD
func ProbarBaseDatos(factory FactoryBaseDatos, tipo string, config map[string]string) {
	// TODO: Crear BD usando factory, conectar, ejecutar consulta, cerrar
}

func ejercicio8() {
	fmt.Println("=== Ejercicio 8: Factory Pattern ===")

	// TODO: Crear factory y probar diferentes tipos de bases de datos

	fmt.Println("✅ Ejercicio 8 completado\n")
}

// ========================================
// Función principal
// ========================================

func main() {
	fmt.Println("🎪 Ejercicios de Interfaces Básicas en Go")
	fmt.Println("=========================================")
	fmt.Println()

	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()

	fmt.Println("🎉 ¡Todos los ejercicios completados!")
	fmt.Println("\n💡 Para ver las soluciones completas, revisa el archivo 'soluciones.go'")
}
