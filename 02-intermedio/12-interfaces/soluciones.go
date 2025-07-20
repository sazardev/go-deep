// Archivo: soluciones.go
// Lección 12: Interfaces - Soluciones Completas
// Este archivo contiene las soluciones para todos los ejercicios de interfaces

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ==============================================
// SOLUCIÓN EJERCICIO 1: Formas Geométricas
// ==============================================

type FormaSol interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

type RectanguloSol struct {
	Ancho, Alto float64
}

func (r RectanguloSol) Area() float64 {
	return r.Ancho * r.Alto
}

func (r RectanguloSol) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

func (r RectanguloSol) Nombre() string {
	return fmt.Sprintf("Rectángulo (%.1fx%.1f)", r.Ancho, r.Alto)
}

type CirculoSol struct {
	Radio float64
}

func (c CirculoSol) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c CirculoSol) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c CirculoSol) Nombre() string {
	return fmt.Sprintf("Círculo (r=%.1f)", c.Radio)
}

type TrianguloSol struct {
	Base, Altura, Lado1, Lado2, Lado3 float64
}

func (t TrianguloSol) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t TrianguloSol) Perimetro() float64 {
	return t.Lado1 + t.Lado2 + t.Lado3
}

func (t TrianguloSol) Nombre() string {
	return fmt.Sprintf("Triángulo (base=%.1f, altura=%.1f)", t.Base, t.Altura)
}

func solucion1() {
	fmt.Println("=== SOLUCIÓN 1: Formas Geométricas ===")

	rectangulo := RectanguloSol{Ancho: 5, Alto: 3}
	circulo := CirculoSol{Radio: 2.5}
	triangulo := TrianguloSol{Base: 4, Altura: 3, Lado1: 3, Lado2: 4, Lado3: 5}

	formas := []FormaSol{rectangulo, circulo, triangulo}

	for _, forma := range formas {
		mostrarInfoFormaSol(forma)
	}
}

func mostrarInfoFormaSol(f FormaSol) {
	fmt.Printf("📐 %s:\n", f.Nombre())
	fmt.Printf("   Área: %.2f unidades²\n", f.Area())
	fmt.Printf("   Perímetro: %.2f unidades\n", f.Perimetro())
	fmt.Println()
}

// ==============================================
// SOLUCIÓN EJERCICIO 2: Sistema de Transporte
// ==============================================

type VehiculoSol interface {
	Acelerar() string
	Frenar() string
	ObtenerVelocidad() int
	ObtenerTipo() string
}

type AutoSol struct {
	Marca     string
	Velocidad int
}

func (a *AutoSol) Acelerar() string {
	a.Velocidad += 20
	return fmt.Sprintf("🚗 %s acelerando...", a.Marca)
}

func (a *AutoSol) Frenar() string {
	if a.Velocidad > 0 {
		a.Velocidad = 0
	}
	return fmt.Sprintf("🛑 %s frenando...", a.Marca)
}

func (a *AutoSol) ObtenerVelocidad() int {
	return a.Velocidad
}

func (a *AutoSol) ObtenerTipo() string {
	return fmt.Sprintf("Auto %s", a.Marca)
}

type MotocicletaSol struct {
	Marca     string
	Velocidad int
}

func (m *MotocicletaSol) Acelerar() string {
	m.Velocidad += 30
	return fmt.Sprintf("🏍️ %s acelerando rápidamente...", m.Marca)
}

func (m *MotocicletaSol) Frenar() string {
	if m.Velocidad > 0 {
		m.Velocidad = 0
	}
	return fmt.Sprintf("🛑 %s frenando...", m.Marca)
}

func (m *MotocicletaSol) ObtenerVelocidad() int {
	return m.Velocidad
}

func (m *MotocicletaSol) ObtenerTipo() string {
	return fmt.Sprintf("Motocicleta %s", m.Marca)
}

type BicicletaSol struct {
	Tipo      string
	Velocidad int
}

func (b *BicicletaSol) Acelerar() string {
	b.Velocidad += 10
	return fmt.Sprintf("🚴 Bicicleta %s pedaleando...", b.Tipo)
}

func (b *BicicletaSol) Frenar() string {
	if b.Velocidad > 0 {
		b.Velocidad = 0
	}
	return "🛑 Bicicleta frenando..."
}

func (b *BicicletaSol) ObtenerVelocidad() int {
	return b.Velocidad
}

func (b *BicicletaSol) ObtenerTipo() string {
	return fmt.Sprintf("Bicicleta de %s", b.Tipo)
}

func solucion2() {
	fmt.Println("=== SOLUCIÓN 2: Sistema de Transporte ===")

	auto := &AutoSol{Marca: "Toyota", Velocidad: 0}
	moto := &MotocicletaSol{Marca: "Honda", Velocidad: 0}
	bici := &BicicletaSol{Tipo: "Montaña", Velocidad: 0}

	vehiculos := []VehiculoSol{auto, moto, bici}

	for _, vehiculo := range vehiculos {
		conducirVehiculoSol(vehiculo)
	}
}

func conducirVehiculoSol(v VehiculoSol) {
	fmt.Printf("\n🚦 %s:\n", v.ObtenerTipo())
	fmt.Printf("   %s\n", v.Acelerar())
	fmt.Printf("   Velocidad actual: %d km/h\n", v.ObtenerVelocidad())
	fmt.Printf("   %s\n", v.Frenar())
	fmt.Printf("   Velocidad actual: %d km/h\n", v.ObtenerVelocidad())
}

// ==============================================
// SOLUCIÓN EJERCICIO 3: Sistema de Archivos
// ==============================================

type LectorSol interface {
	Read() ([]byte, error)
}

type EscritorSol interface {
	Write(data []byte) error
}

type CerradorSol interface {
	Close() error
}

type LectorEscritorSol interface {
	LectorSol
	EscritorSol
}

type ArchivoCompletoSol interface {
	LectorSol
	EscritorSol
	CerradorSol
}

type ArchivoSol struct {
	nombre   string
	datos    []byte
	posicion int
	cerrado  bool
}

func (a *ArchivoSol) Read() ([]byte, error) {
	if a.cerrado {
		return nil, fmt.Errorf("archivo %s está cerrado", a.nombre)
	}

	datos := make([]byte, len(a.datos))
	copy(datos, a.datos)
	return datos, nil
}

func (a *ArchivoSol) Write(data []byte) error {
	if a.cerrado {
		return fmt.Errorf("archivo %s está cerrado", a.nombre)
	}

	a.datos = append(a.datos, data...)
	return nil
}

func (a *ArchivoSol) Close() error {
	if a.cerrado {
		return fmt.Errorf("archivo %s ya está cerrado", a.nombre)
	}

	a.cerrado = true
	return nil
}

func solucion3() {
	fmt.Println("=== SOLUCIÓN 3: Sistema de Archivos ===")

	archivo := &ArchivoSol{
		nombre: "test.txt",
		datos:  []byte("Contenido inicial del archivo"),
	}

	usarComoLectorSol(archivo)
	usarComoEscritorSol(archivo)
	usarComoArchivoCompletoSol(archivo)
}

func usarComoLectorSol(r LectorSol) {
	fmt.Println("📖 Usando como Lector:")
	data, err := r.Read()
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Leído: %s\n", string(data))
	}
}

func usarComoEscritorSol(w EscritorSol) {
	fmt.Println("✏️  Usando como Escritor:")
	err := w.Write([]byte(" - Datos agregados"))
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   Escritura exitosa")
	}
}

func usarComoArchivoCompletoSol(ac ArchivoCompletoSol) {
	fmt.Println("📁 Usando como ArchivoCompleto:")
	data, _ := ac.Read()
	fmt.Printf("   Contenido final: %s\n", string(data))
	fmt.Println("   Cerrando archivo...")
	ac.Close()
}

// ==============================================
// SOLUCIÓN EJERCICIO 4: Procesador Universal
// ==============================================

func ProcesarDatosSol(datos interface{}) string {
	switch v := datos.(type) {
	case string:
		return fmt.Sprintf("String con %d caracteres: '%s'", len(v), v)
	case int:
		if v%2 == 0 {
			return fmt.Sprintf("Número par: %d", v)
		}
		return fmt.Sprintf("Número impar: %d", v)
	case float64:
		return fmt.Sprintf("Float redondeado: %.2f", v)
	case []int:
		suma := 0
		for _, num := range v {
			suma += num
		}
		promedio := float64(suma) / float64(len(v))
		return fmt.Sprintf("Slice con %d elementos, suma: %d, promedio: %.2f",
			len(v), suma, promedio)
	case map[string]int:
		return fmt.Sprintf("Map con %d claves: %v", len(v), getKeys(v))
	case bool:
		if v {
			return "Valor booleano: verdadero"
		}
		return "Valor booleano: falso"
	default:
		return fmt.Sprintf("Tipo no soportado: %T", v)
	}
}

func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func solucion4() {
	fmt.Println("=== SOLUCIÓN 4: Procesador Universal ===")

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
		resultado := ProcesarDatosSol(dato)
		fmt.Printf("📊 Dato %d: %v\n   -> %s\n\n", i+1, dato, resultado)
	}
}

// ==============================================
// SOLUCIÓN EJERCICIO 5: Type Assertions
// ==============================================

type EmpleadoSol interface {
	ObtenerSalario() float64
	ObtenerNombre() string
}

type EmpleadoBaseSol struct {
	Nombre  string
	Salario float64
}

func (e EmpleadoBaseSol) ObtenerSalario() float64 {
	return e.Salario
}

func (e EmpleadoBaseSol) ObtenerNombre() string {
	return e.Nombre
}

type DesarrolladorSol struct {
	EmpleadoBaseSol
	Lenguajes []string
	Nivel     string
}

func (d DesarrolladorSol) ProgramarEn(lenguaje string) string {
	for _, l := range d.Lenguajes {
		if l == lenguaje {
			return fmt.Sprintf("💻 %s está programando en %s", d.Nombre, lenguaje)
		}
	}
	return fmt.Sprintf("❌ %s no conoce %s", d.Nombre, lenguaje)
}

type GerenteSol struct {
	EmpleadoBaseSol
	Departamento string
	Equipo       int
}

func (g GerenteSol) Dirigir() string {
	return fmt.Sprintf("👔 %s está dirigiendo el departamento de %s con %d personas",
		g.Nombre, g.Departamento, g.Equipo)
}

func AnalizarEmpleadoSol(e EmpleadoSol) {
	fmt.Printf("👤 Empleado: %s (Salario: $%.0f)\n",
		e.ObtenerNombre(), e.ObtenerSalario())

	// Type assertion con verificación
	if dev, ok := e.(DesarrolladorSol); ok {
		fmt.Printf("   🔧 Rol: Desarrollador %s\n", dev.Nivel)
		fmt.Printf("   🔤 Lenguajes: %v\n", dev.Lenguajes)
		fmt.Printf("   %s\n", dev.ProgramarEn("Go"))
	} else if ger, ok := e.(GerenteSol); ok {
		fmt.Printf("   📋 Rol: Gerente\n")
		fmt.Printf("   🏢 Departamento: %s\n", ger.Departamento)
		fmt.Printf("   👥 Equipo: %d personas\n", ger.Equipo)
		fmt.Printf("   %s\n", ger.Dirigir())
	} else {
		fmt.Printf("   ❓ Rol: Empleado base\n")
	}
}

func solucion5() {
	fmt.Println("=== SOLUCIÓN 5: Type Assertions ===")

	empleados := []EmpleadoSol{
		DesarrolladorSol{
			EmpleadoBaseSol: EmpleadoBaseSol{Nombre: "Ana García", Salario: 75000},
			Lenguajes:       []string{"Go", "Python", "JavaScript"},
			Nivel:           "Senior",
		},
		GerenteSol{
			EmpleadoBaseSol: EmpleadoBaseSol{Nombre: "Carlos López", Salario: 95000},
			Departamento:    "Tecnología",
			Equipo:          12,
		},
		DesarrolladorSol{
			EmpleadoBaseSol: EmpleadoBaseSol{Nombre: "María Rodríguez", Salario: 68000},
			Lenguajes:       []string{"Java", "C++"},
			Nivel:           "Mid",
		},
	}

	for _, empleado := range empleados {
		AnalizarEmpleadoSol(empleado)
		fmt.Println()
	}
}

// ==============================================
// SOLUCIÓN EJERCICIO 6: Interfaces Estándar
// ==============================================

type ProductoSol struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

// Implementar fmt.Stringer
func (p ProductoSol) String() string {
	return fmt.Sprintf("📦 %s (#%d) - $%.2f (Stock: %d)",
		p.Nombre, p.ID, p.Precio, p.Stock)
}

// Ordenamiento por precio
type PorPrecioSol []ProductoSol

func (p PorPrecioSol) Len() int           { return len(p) }
func (p PorPrecioSol) Less(i, j int) bool { return p[i].Precio < p[j].Precio }
func (p PorPrecioSol) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Ordenamiento por stock
type PorStockSol []ProductoSol

func (p PorStockSol) Len() int           { return len(p) }
func (p PorStockSol) Less(i, j int) bool { return p[i].Stock < p[j].Stock }
func (p PorStockSol) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func solucion6() {
	fmt.Println("=== SOLUCIÓN 6: Interfaces Estándar ===")

	productos := []ProductoSol{
		{1, "Laptop Gaming", 1200.00, 5},
		{2, "Mouse Inalámbrico", 25.50, 50},
		{3, "Teclado Mecánico", 75.00, 20},
		{4, "Monitor 4K", 300.00, 8},
	}

	fmt.Println("📋 Productos originales:")
	for _, p := range productos {
		fmt.Printf("   %s\n", p)
	}

	// Ordenar por precio
	porPrecio := PorPrecioSol(productos)
	sort.Sort(porPrecio)
	fmt.Println("\n💰 Ordenados por precio (menor a mayor):")
	for _, p := range porPrecio {
		fmt.Printf("   %s\n", p)
	}

	// Ordenar por stock
	porStock := PorStockSol(productos)
	sort.Sort(porStock)
	fmt.Println("\n📦 Ordenados por stock (menor a mayor):")
	for _, p := range porStock {
		fmt.Printf("   %s\n", p)
	}
}

// ==============================================
// SOLUCIÓN EJERCICIO 7: Strategy Pattern
// ==============================================

type EstrategiaDescuentoSol interface {
	AplicarDescuento(precio float64) float64
	Descripcion() string
}

type SinDescuentoSol struct{}

func (s SinDescuentoSol) AplicarDescuento(precio float64) float64 {
	return precio
}

func (s SinDescuentoSol) Descripcion() string {
	return "Sin descuento"
}

type DescuentoPorcentajeSol struct {
	Porcentaje float64
}

func (d DescuentoPorcentajeSol) AplicarDescuento(precio float64) float64 {
	return precio * (1 - d.Porcentaje/100)
}

func (d DescuentoPorcentajeSol) Descripcion() string {
	return fmt.Sprintf("Descuento %.0f%%", d.Porcentaje)
}

type DescuentoCantidadFijaSol struct {
	Cantidad float64
}

func (d DescuentoCantidadFijaSol) AplicarDescuento(precio float64) float64 {
	resultado := precio - d.Cantidad
	if resultado < 0 {
		return 0
	}
	return resultado
}

func (d DescuentoCantidadFijaSol) Descripcion() string {
	return fmt.Sprintf("Descuento fijo $%.2f", d.Cantidad)
}

type DescuentoPorCategoriaSol struct {
	Categoria string
	Descuento float64
}

func (d DescuentoPorCategoriaSol) AplicarDescuento(precio float64) float64 {
	return precio * (1 - d.Descuento/100)
}

func (d DescuentoPorCategoriaSol) Descripcion() string {
	return fmt.Sprintf("Descuento %s (%.0f%%)", d.Categoria, d.Descuento)
}

type CalculadoraPreciosSol struct {
	estrategia EstrategiaDescuentoSol
}

func (c *CalculadoraPreciosSol) SetEstrategia(e EstrategiaDescuentoSol) {
	c.estrategia = e
}

func (c *CalculadoraPreciosSol) CalcularPrecio(precio float64) float64 {
	if c.estrategia == nil {
		return precio
	}
	return c.estrategia.AplicarDescuento(precio)
}

func solucion7() {
	fmt.Println("=== SOLUCIÓN 7: Strategy Pattern ===")

	calculadora := &CalculadoraPreciosSol{}
	precioOriginal := 100.0

	estrategias := []EstrategiaDescuentoSol{
		SinDescuentoSol{},
		DescuentoPorcentajeSol{Porcentaje: 10},
		DescuentoCantidadFijaSol{Cantidad: 15},
		DescuentoPorCategoriaSol{Categoria: "Electrónicos", Descuento: 20},
	}

	fmt.Printf("💰 Precio original: $%.2f\n\n", precioOriginal)

	for _, estrategia := range estrategias {
		calculadora.SetEstrategia(estrategia)
		precioFinal := calculadora.CalcularPrecio(precioOriginal)
		ahorro := precioOriginal - precioFinal
		fmt.Printf("🏷️  %s:\n", estrategia.Descripcion())
		fmt.Printf("   Precio final: $%.2f (Ahorro: $%.2f)\n\n", precioFinal, ahorro)
	}
}

// ==============================================
// SOLUCIÓN EJERCICIO 8: Factory Pattern
// ==============================================

type ProcesadorDatosSol interface {
	Procesar(datos string) (string, error)
	TipoSoportado() string
}

type FactoryProcesadorSol interface {
	CrearProcesador() ProcesadorDatosSol
}

// Procesadores
type ProcesadorJSONSol struct{}

func (p ProcesadorJSONSol) Procesar(datos string) (string, error) {
	if !strings.Contains(datos, "{") || !strings.Contains(datos, "}") {
		return "", fmt.Errorf("formato JSON inválido")
	}
	return fmt.Sprintf("✅ JSON procesado: %d caracteres", len(datos)), nil
}

func (p ProcesadorJSONSol) TipoSoportado() string {
	return "JSON"
}

type ProcesadorXMLSol struct{}

func (p ProcesadorXMLSol) Procesar(datos string) (string, error) {
	if !strings.Contains(datos, "<") || !strings.Contains(datos, ">") {
		return "", fmt.Errorf("formato XML inválido")
	}
	return fmt.Sprintf("✅ XML procesado: %d caracteres", len(datos)), nil
}

func (p ProcesadorXMLSol) TipoSoportado() string {
	return "XML"
}

type ProcesadorCSVSol struct{}

func (p ProcesadorCSVSol) Procesar(datos string) (string, error) {
	lineas := strings.Split(datos, "\n")
	if len(lineas) < 2 {
		return "", fmt.Errorf("CSV debe tener al menos 2 líneas")
	}
	return fmt.Sprintf("✅ CSV procesado: %d líneas", len(lineas)), nil
}

func (p ProcesadorCSVSol) TipoSoportado() string {
	return "CSV"
}

// Factories
type JSONFactorySol struct{}

func (f JSONFactorySol) CrearProcesador() ProcesadorDatosSol {
	return ProcesadorJSONSol{}
}

type XMLFactorySol struct{}

func (f XMLFactorySol) CrearProcesador() ProcesadorDatosSol {
	return ProcesadorXMLSol{}
}

type CSVFactorySol struct{}

func (f CSVFactorySol) CrearProcesador() ProcesadorDatosSol {
	return ProcesadorCSVSol{}
}

func ObtenerFactorySol(tipo string) FactoryProcesadorSol {
	switch strings.ToLower(tipo) {
	case "json":
		return JSONFactorySol{}
	case "xml":
		return XMLFactorySol{}
	case "csv":
		return CSVFactorySol{}
	default:
		return nil
	}
}

func solucion8() {
	fmt.Println("=== SOLUCIÓN 8: Factory Pattern ===")

	datos := map[string]string{
		"json": `{"nombre": "Juan", "edad": 30, "activo": true}`,
		"xml":  `<persona><nombre>Juan</nombre><edad>30</edad></persona>`,
		"csv":  `nombre,edad,activo\nJuan,30,true\nMaria,25,false`,
	}

	fmt.Println("🏭 Procesando datos con diferentes factories:\n")

	for tipo, contenido := range datos {
		factory := ObtenerFactorySol(tipo)
		if factory != nil {
			procesador := factory.CrearProcesador()
			resultado, err := procesador.Procesar(contenido)

			fmt.Printf("📄 Tipo: %s\n", strings.ToUpper(tipo))
			fmt.Printf("   Datos: %s\n", contenido)

			if err != nil {
				fmt.Printf("   ❌ Error: %v\n", err)
			} else {
				fmt.Printf("   %s\n", resultado)
			}
			fmt.Println()
		} else {
			fmt.Printf("❌ %s: Factory no encontrada\n\n", tipo)
		}
	}
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func main() {
	fmt.Println("🔌 SOLUCIONES DE EJERCICIOS - INTERFACES")
	fmt.Println("========================================")

	solucion1()
	solucion2()
	solucion3()
	solucion4()
	solucion5()
	solucion6()
	solucion7()
	solucion8()

	fmt.Println("✅ ¡Todas las soluciones ejecutadas exitosamente!")
	fmt.Println("🎓 Has dominado las interfaces en Go!")
}
