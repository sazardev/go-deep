// Archivo: ejercicios.go
// Lección 12: Interfaces - Ejercicios Prácticos
// Cubre: declaración, polimorfismo, embedding, type assertions, patrones

package main

import (
	"fmt"
	"sort"
)

// ==============================================
// EJERCICIO 1: Interface Básica - Formas Geométricas
// ==============================================

// TODO: Define una interface Forma con métodos:
// - Area() float64
// - Perimetro() float64
// - Nombre() string

type Forma interface {
	// Implementa aquí los métodos requeridos
}

// TODO: Implementa la interface con estas estructuras:

type Rectangulo struct {
	Ancho, Alto float64
}

type Circulo struct {
	Radio float64
}

type Triangulo struct {
	Base, Altura, Lado1, Lado2, Lado3 float64
}

// TODO: Implementa los métodos para cada forma

func ejercicio1() {
	fmt.Println("=== EJERCICIO 1: Formas Geométricas ===")

	// Crear instancias de cada forma
	rectangulo := Rectangulo{Ancho: 5, Alto: 3}
	circulo := Circulo{Radio: 2.5}
	triangulo := Triangulo{Base: 4, Altura: 3, Lado1: 3, Lado2: 4, Lado3: 5}

	// Slice de formas usando polimorfismo
	formas := []Forma{rectangulo, circulo, triangulo}

	// Mostrar información de cada forma
	for _, forma := range formas {
		mostrarInfoForma(forma)
	}
}

func mostrarInfoForma(f Forma) {
	fmt.Printf("%s:\n", f.Nombre())
	fmt.Printf("  Área: %.2f\n", f.Area())
	fmt.Printf("  Perímetro: %.2f\n", f.Perimetro())
	fmt.Println()
}

// ==============================================
// EJERCICIO 2: Polimorfismo - Sistema de Transporte
// ==============================================

// TODO: Define una interface Vehiculo con métodos:
// - Acelerar() string
// - Frenar() string
// - ObtenerVelocidad() int
// - ObtenerTipo() string

type Vehiculo interface {
	// Implementa aquí
}

// TODO: Implementa diferentes tipos de vehículos

type Auto struct {
	Marca     string
	Velocidad int
}

type Motocicleta struct {
	Marca     string
	Velocidad int
}

type Bicicleta struct {
	Tipo      string
	Velocidad int
}

// TODO: Implementa los métodos para cada vehículo

func ejercicio2() {
	fmt.Println("=== EJERCICIO 2: Sistema de Transporte ===")

	// Crear vehículos
	auto := &Auto{Marca: "Toyota", Velocidad: 0}
	moto := &Motocicleta{Marca: "Honda", Velocidad: 0}
	bici := &Bicicleta{Tipo: "Montaña", Velocidad: 0}

	vehiculos := []Vehiculo{auto, moto, bici}

	// Simular conducción
	for _, vehiculo := range vehiculos {
		conducirVehiculo(vehiculo)
	}
}

func conducirVehiculo(v Vehiculo) {
	fmt.Printf("\n%s:\n", v.ObtenerTipo())
	fmt.Printf("- %s\n", v.Acelerar())
	fmt.Printf("- Velocidad actual: %d km/h\n", v.ObtenerVelocidad())
	fmt.Printf("- %s\n", v.Frenar())
	fmt.Printf("- Velocidad actual: %d km/h\n", v.ObtenerVelocidad())
}

// ==============================================
// EJERCICIO 3: Interface Embedding - Sistema de Archivos
// ==============================================

// TODO: Define interfaces básicas

type Lector interface {
	Read() ([]byte, error)
}

type Escritor interface {
	Write(data []byte) error
}

type Cerrador interface {
	Close() error
}

// TODO: Define interfaces compuestas usando embedding

type LectorEscritor interface {
	// Combina Lector y Escritor
}

type ArchivoCompleto interface {
	// Combina Lector, Escritor y Cerrador
}

// TODO: Implementa una estructura que satisfaga ArchivoCompleto

type Archivo struct {
	nombre   string
	datos    []byte
	posicion int
	cerrado  bool
}

// TODO: Implementa los métodos necesarios

func ejercicio3() {
	fmt.Println("=== EJERCICIO 3: Sistema de Archivos ===")

	// Crear archivo
	archivo := &Archivo{
		nombre: "test.txt",
		datos:  []byte("Contenido inicial del archivo"),
	}

	// Usar como diferentes interfaces
	usarComoLector(archivo)
	usarComoEscritor(archivo)
	usarComoArchivoCompleto(archivo)
}

func usarComoLector(r Lector) {
	fmt.Println("Usando como Lector:")
	data, err := r.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Leído: %s\n", string(data))
	}
}

func usarComoEscritor(w Escritor) {
	fmt.Println("Usando como Escritor:")
	err := w.Write([]byte(" - Datos agregados"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Escritura exitosa")
	}
}

func usarComoArchivoCompleto(ac ArchivoCompleto) {
	fmt.Println("Usando como ArchivoCompleto:")
	data, _ := ac.Read()
	fmt.Printf("Contenido final: %s\n", string(data))
	ac.Close()
}

// ==============================================
// EJERCICIO 4: Empty Interface - Procesador Universal
// ==============================================

// TODO: Implementa una función que procese diferentes tipos de datos

func ProcesarDatos(datos interface{}) string {
	// TODO: Usa type switch para manejar diferentes tipos:
	// - string: contar caracteres
	// - int: verificar si es par o impar
	// - float64: redondear a 2 decimales
	// - []int: calcular suma y promedio
	// - map[string]int: contar claves
	// - bool: convertir a texto

	return "Procesado"
}

func ejercicio4() {
	fmt.Println("=== EJERCICIO 4: Procesador Universal ===")

	// Datos de diferentes tipos
	datos := []interface{}{
		"Hola Mundo",
		42,
		3.14159,
		[]int{1, 2, 3, 4, 5},
		map[string]int{"a": 1, "b": 2, "c": 3},
		true,
		false,
	}

	for i, dato := range datos {
		resultado := ProcesarDatos(dato)
		fmt.Printf("Dato %d: %v -> %s\n", i+1, dato, resultado)
	}
}

// ==============================================
// EJERCICIO 5: Type Assertions - Sistema de Empleados
// ==============================================

type Empleado interface {
	ObtenerSalario() float64
	ObtenerNombre() string
}

type EmpleadoBase struct {
	Nombre  string
	Salario float64
}

func (e EmpleadoBase) ObtenerSalario() float64 {
	return e.Salario
}

func (e EmpleadoBase) ObtenerNombre() string {
	return e.Nombre
}

type Desarrollador struct {
	EmpleadoBase
	Lenguajes []string
	Nivel     string
}

func (d Desarrollador) ProgramarEn(lenguaje string) string {
	for _, l := range d.Lenguajes {
		if l == lenguaje {
			return fmt.Sprintf("%s está programando en %s", d.Nombre, lenguaje)
		}
	}
	return fmt.Sprintf("%s no conoce %s", d.Nombre, lenguaje)
}

type Gerente struct {
	EmpleadoBase
	Departamento string
	Equipo       int
}

func (g Gerente) Dirigir() string {
	return fmt.Sprintf("%s está dirigiendo el departamento de %s con %d personas",
		g.Nombre, g.Departamento, g.Equipo)
}

// TODO: Implementa funciones que usen type assertions

func AnalizarEmpleado(e Empleado) {
	// TODO: Usa type assertions para determinar el tipo específico
	// Si es Desarrollador, mostrar sus lenguajes
	// Si es Gerente, mostrar su departamento

	fmt.Printf("Empleado: %s\n", e.ObtenerNombre())
}

func ejercicio5() {
	fmt.Println("=== EJERCICIO 5: Type Assertions ===")

	empleados := []Empleado{
		Desarrollador{
			EmpleadoBase: EmpleadoBase{Nombre: "Ana", Salario: 75000},
			Lenguajes:    []string{"Go", "Python", "JavaScript"},
			Nivel:        "Senior",
		},
		Gerente{
			EmpleadoBase: EmpleadoBase{Nombre: "Carlos", Salario: 95000},
			Departamento: "Tecnología",
			Equipo:       12,
		},
		Desarrollador{
			EmpleadoBase: EmpleadoBase{Nombre: "María", Salario: 68000},
			Lenguajes:    []string{"Java", "C++"},
			Nivel:        "Mid",
		},
	}

	for _, empleado := range empleados {
		AnalizarEmpleado(empleado)
		fmt.Println()
	}
}

// ==============================================
// EJERCICIO 6: Interfaces Estándar - Ordenamiento
// ==============================================

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

// TODO: Implementa sort.Interface para ordenar productos por precio

type PorPrecio []Producto

// TODO: Implementa sort.Interface para ordenar productos por stock

type PorStock []Producto

// TODO: Implementa fmt.Stringer para Producto

func ejercicio6() {
	fmt.Println("=== EJERCICIO 6: Interfaces Estándar ===")

	productos := []Producto{
		{1, "Laptop", 1200.00, 5},
		{2, "Mouse", 25.50, 50},
		{3, "Teclado", 75.00, 20},
		{4, "Monitor", 300.00, 8},
	}

	fmt.Println("Productos originales:")
	for _, p := range productos {
		fmt.Println(p)
	}

	// Ordenar por precio
	porPrecio := PorPrecio(productos)
	sort.Sort(porPrecio)
	fmt.Println("\nOrdenados por precio:")
	for _, p := range porPrecio {
		fmt.Println(p)
	}

	// Ordenar por stock
	porStock := PorStock(productos)
	sort.Sort(porStock)
	fmt.Println("\nOrdenados por stock:")
	for _, p := range porStock {
		fmt.Println(p)
	}
}

// ==============================================
// EJERCICIO 7: Strategy Pattern - Sistema de Descuentos
// ==============================================

// TODO: Define interface para estrategias de descuento

type EstrategiaDescuento interface {
	// AplicarDescuento(precio float64) float64
	// Descripcion() string
}

// TODO: Implementa diferentes estrategias

type SinDescuento struct{}

type DescuentoPorcentaje struct {
	Porcentaje float64
}

type DescuentoCantidadFija struct {
	Cantidad float64
}

type DescuentoPorCategoria struct {
	Categoria string
	Descuento float64
}

// TODO: Implementa los métodos para cada estrategia

// Estructura para usar las estrategias
type CalculadoraPrecios struct {
	estrategia EstrategiaDescuento
}

func (c *CalculadoraPrecios) SetEstrategia(e EstrategiaDescuento) {
	c.estrategia = e
}

func (c *CalculadoraPrecios) CalcularPrecio(precio float64) float64 {
	if c.estrategia == nil {
		return precio
	}
	return c.estrategia.AplicarDescuento(precio)
}

func ejercicio7() {
	fmt.Println("=== EJERCICIO 7: Strategy Pattern ===")

	calculadora := &CalculadoraPrecios{}
	precioOriginal := 100.0

	estrategias := []EstrategiaDescuento{
		SinDescuento{},
		DescuentoPorcentaje{Porcentaje: 10},
		DescuentoCantidadFija{Cantidad: 15},
		DescuentoPorCategoria{Categoria: "Electrónicos", Descuento: 20},
	}

	for _, estrategia := range estrategias {
		calculadora.SetEstrategia(estrategia)
		precioFinal := calculadora.CalcularPrecio(precioOriginal)
		fmt.Printf("%s: $%.2f -> $%.2f\n",
			estrategia.Descripcion(), precioOriginal, precioFinal)
	}
}

// ==============================================
// EJERCICIO 8: Factory Pattern - Procesadores de Datos
// ==============================================

// TODO: Define interface para procesadores

type ProcesadorDatos interface {
	// Procesar(datos string) (string, error)
	// TipoSoportado() string
}

// TODO: Define interface para factory

type FactoryProcesador interface {
	// CrearProcesador() ProcesadorDatos
}

// TODO: Implementa diferentes procesadores

type ProcesadorJSON struct{}

type ProcesadorXML struct{}

type ProcesadorCSV struct{}

// TODO: Implementa factories

type JSONFactory struct{}

type XMLFactory struct{}

type CSVFactory struct{}

// TODO: Implementa función para obtener factory por tipo

func ObtenerFactory(tipo string) FactoryProcesador {
	// TODO: Retorna la factory apropiada según el tipo
	return nil
}

func ejercicio8() {
	fmt.Println("=== EJERCICIO 8: Factory Pattern ===")

	datos := map[string]string{
		"json": `{"nombre": "Juan", "edad": 30}`,
		"xml":  `<persona><nombre>Juan</nombre><edad>30</edad></persona>`,
		"csv":  `nombre,edad\nJuan,30`,
	}

	for tipo, contenido := range datos {
		factory := ObtenerFactory(tipo)
		if factory != nil {
			procesador := factory.CrearProcesador()
			resultado, err := procesador.Procesar(contenido)
			if err != nil {
				fmt.Printf("%s: Error - %v\n", tipo, err)
			} else {
				fmt.Printf("%s: %s\n", tipo, resultado)
			}
		} else {
			fmt.Printf("%s: Factory no encontrada\n", tipo)
		}
	}
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func main() {
	fmt.Println("🔌 EJERCICIOS DE INTERFACES EN GO")
	fmt.Println("=================================")

	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()

	fmt.Println("\n✅ ¡Todos los ejercicios completados!")
	fmt.Println("Verifica tus soluciones con 'go run soluciones.go'")
}
