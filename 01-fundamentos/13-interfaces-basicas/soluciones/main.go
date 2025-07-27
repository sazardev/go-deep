// Lección 13: Interfaces Básicas en Go
// Soluciones completas de todos los ejercicios

/*
INSTRUCCIONES PARA USAR ESTE ARCHIVO:

Este archivo contiene las soluciones completas de todos los ejercicios.

Para ejecutar las soluciones:
1. Desde la terminal, navega al directorio de soluciones:
   cd soluciones

2. Ejecuta el archivo:
   go run main.go

3. Para ejecutar ejercicios específicos, modifica la función main
   para comentar/descomentar los ejercicios deseados.

Para comparar con los ejercicios:
1. Abre ../ejercicios.go en una ventana
2. Abre main.go en otra ventana
3. Compara las implementaciones lado a lado

Cada ejercicio está completamente implementado y funcionando.
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ========================================
// Ejercicio 1: Interface Básica - Sistema de Formas Geométricas
// ========================================

// Solución: Interface Forma con métodos necesarios
type Forma interface {
	Area() float64
	Perimetro() float64
	Descripcion() string
}

// Solución: Struct Rectangulo con campos requeridos
type Rectangulo struct {
	Ancho float64
	Alto  float64
}

// Solución: Implementación de métodos para Rectangulo
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

func (r Rectangulo) Descripcion() string {
	return fmt.Sprintf("Rectángulo de %.1f x %.1f", r.Ancho, r.Alto)
}

// Solución: Struct Circulo con campo requerido
type Circulo struct {
	Radio float64
}

// Solución: Implementación de métodos para Circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circulo) Descripcion() string {
	return fmt.Sprintf("Círculo con radio %.1f", c.Radio)
}

// Solución: Función que maneja cualquier forma
func MostrarInformacionForma(f Forma) {
	fmt.Printf("🔶 %s\n", f.Descripcion())
	fmt.Printf("   Área: %.2f\n", f.Area())
	fmt.Printf("   Perímetro: %.2f\n", f.Perimetro())
	fmt.Println()
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Sistema de Formas Geométricas ===")

	rectangulo := Rectangulo{Ancho: 5.0, Alto: 3.0}
	circulo := Circulo{Radio: 2.5}

	formas := []Forma{rectangulo, circulo}

	for _, forma := range formas {
		MostrarInformacionForma(forma)
	}

	fmt.Println("✅ Ejercicio 1 completado\n")
}

// ========================================
// Ejercicio 2: Polimorfismo - Sistema de Animales
// ========================================

type Animal interface {
	HacerSonido() string
	Moverse() string
	Comer(comida string) string
}

type Perro struct {
	Nombre string
	Raza   string
}

func (p Perro) HacerSonido() string {
	return fmt.Sprintf("%s hace: ¡Guau guau!", p.Nombre)
}

func (p Perro) Moverse() string {
	return fmt.Sprintf("%s corre alegremente", p.Nombre)
}

func (p Perro) Comer(comida string) string {
	return fmt.Sprintf("%s está comiendo %s con mucha energía", p.Nombre, comida)
}

type Gato struct {
	Nombre string
	Color  string
}

func (g Gato) HacerSonido() string {
	return fmt.Sprintf("%s hace: Miau miau", g.Nombre)
}

func (g Gato) Moverse() string {
	return fmt.Sprintf("%s camina sigilosamente", g.Nombre)
}

func (g Gato) Comer(comida string) string {
	return fmt.Sprintf("%s está comiendo %s elegantemente", g.Nombre, comida)
}

type Pajaro struct {
	Nombre     string
	Especie    string
	PuedeVolar bool
}

func (p Pajaro) HacerSonido() string {
	return fmt.Sprintf("%s hace: ¡Pío pío!", p.Nombre)
}

func (p Pajaro) Moverse() string {
	if p.PuedeVolar {
		return fmt.Sprintf("%s vuela graciosamente", p.Nombre)
	}
	return fmt.Sprintf("%s camina saltando", p.Nombre)
}

func (p Pajaro) Comer(comida string) string {
	return fmt.Sprintf("%s está picoteando %s", p.Nombre, comida)
}

func CuidarAnimal(a Animal) {
	fmt.Printf("🐾 Cuidando a un animal:\n")
	fmt.Printf("   Sonido: %s\n", a.HacerSonido())
	fmt.Printf("   Movimiento: %s\n", a.Moverse())
	fmt.Printf("   Alimentación: %s\n", a.Comer("comida"))
	fmt.Println()
}

func ejercicio2() {
	fmt.Println("=== Ejercicio 2: Sistema de Animales ===")

	perro := Perro{Nombre: "Max", Raza: "Golden Retriever"}
	gato := Gato{Nombre: "Luna", Color: "Negro"}
	pajaro := Pajaro{Nombre: "Pipo", Especie: "Canario", PuedeVolar: true}

	animales := []Animal{perro, gato, pajaro}

	for _, animal := range animales {
		CuidarAnimal(animal)
	}

	fmt.Println("✅ Ejercicio 2 completado\n")
}

// ========================================
// Ejercicio 3: Interfaces Estándar - fmt.Stringer y sort.Interface
// ========================================

type Producto struct {
	Nombre    string
	Precio    float64
	Categoria string
}

func (p Producto) String() string {
	return fmt.Sprintf("%s ($%.2f) - %s", p.Nombre, p.Precio, p.Categoria)
}

type ProductosPorPrecio []Producto

func (p ProductosPorPrecio) Len() int {
	return len(p)
}

func (p ProductosPorPrecio) Less(i, j int) bool {
	return p[i].Precio < p[j].Precio
}

func (p ProductosPorPrecio) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ejercicio3() {
	fmt.Println("=== Ejercicio 3: Interfaces Estándar ===")

	productos := ProductosPorPrecio{
		{"Laptop", 1200.00, "Electrónicos"},
		{"Mouse", 25.99, "Accesorios"},
		{"Teclado", 75.50, "Accesorios"},
		{"Monitor", 350.00, "Electrónicos"},
		{"Auriculares", 89.99, "Audio"},
	}

	fmt.Println("📦 Productos originales:")
	for i, producto := range productos {
		fmt.Printf("  %d. %s\n", i+1, producto)
	}

	sort.Sort(productos)

	fmt.Println("\n📦 Productos ordenados por precio:")
	for i, producto := range productos {
		fmt.Printf("  %d. %s\n", i+1, producto)
	}

	fmt.Println("✅ Ejercicio 3 completado\n")
}

// ========================================
// Ejercicio 4: Type Assertions - Procesador de Datos
// ========================================

func ProcesarDato(dato interface{}) {
	switch valor := dato.(type) {
	case string:
		fmt.Printf("String: '%s' (longitud: %d, mayúsculas: '%s')\n",
			valor, len(valor), strings.ToUpper(valor))
	case int:
		paridad := "impar"
		if valor%2 == 0 {
			paridad = "par"
		}
		fmt.Printf("Integer: %d (es %s, cuadrado: %d)\n",
			valor, paridad, valor*valor)
	case float64:
		mayorA100 := "no"
		if valor > 100 {
			mayorA100 = "sí"
		}
		fmt.Printf("Float64: %.3f (raíz cuadrada: %.3f, mayor a 100: %s)\n",
			valor, math.Sqrt(valor), mayorA100)
	case []int:
		suma := 0
		for _, num := range valor {
			suma += num
		}
		promedio := float64(suma) / float64(len(valor))
		fmt.Printf("Slice de int: %v (suma: %d, promedio: %.2f)\n",
			valor, suma, promedio)
	case map[string]int:
		claves := make([]string, 0, len(valor))
		suma := 0
		for k, v := range valor {
			claves = append(claves, k)
			suma += v
		}
		fmt.Printf("Map[string]int: claves=%v, suma de valores=%d\n",
			claves, suma)
	default:
		fmt.Printf("Tipo desconocido: %T con valor: %v\n", valor, valor)
	}
}

func ejercicio4() {
	fmt.Println("=== Ejercicio 4: Type Assertions ===")

	datos := []interface{}{
		"Hola Go",
		42,
		3.14159,
		[]int{1, 2, 3, 4, 5},
		map[string]int{"a": 10, "b": 20, "c": 30},
		true,
		struct{ X int }{X: 100},
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

type Logger interface {
	Log(level string, message string, data interface{})
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(level string, message string, data interface{}) {
	fmt.Printf("[%s] %s", strings.ToUpper(level), message)
	if data != nil {
		fmt.Printf(": %v", data)
	}
	fmt.Println()
}

type FileLogger struct {
	archivo string
}

func (fl FileLogger) Log(level string, message string, data interface{}) {
	fmt.Printf("[FILE:%s] [%s] %s", fl.archivo, strings.ToUpper(level), message)
	if data != nil {
		fmt.Printf(": %v", data)
	}
	fmt.Println()
}

func LogearEvento(logger Logger, evento string, datos ...interface{}) {
	logger.Log("INFO", fmt.Sprintf("Evento: %s", evento), nil)

	for i, dato := range datos {
		logger.Log("DEBUG", fmt.Sprintf("Dato %d", i+1), dato)
	}
}

func ejercicio5() {
	fmt.Println("=== Ejercicio 5: Sistema de Logging ===")

	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{archivo: "app.log"}

	fmt.Println("📝 Console Logger:")
	LogearEvento(consoleLogger, "usuario_login",
		map[string]interface{}{"user_id": 123, "ip": "192.168.1.1"},
		"timestamp: 2025-01-15T10:30:00Z",
		[]string{"session", "auth", "web"})

	fmt.Println("\n📁 File Logger:")
	LogearEvento(fileLogger, "compra_realizada",
		map[string]interface{}{"order_id": 456, "amount": 99.99},
		[]string{"laptop", "mouse"},
		true)

	fmt.Println("✅ Ejercicio 5 completado\n")
}

// ========================================
// Ejercicio 6: Strategy Pattern - Sistema de Descuentos
// ========================================

type EstrategiaDescuento interface {
	AplicarDescuento(precio float64) float64
	DescribirDescuento() string
}

type SinDescuento struct{}

func (sd SinDescuento) AplicarDescuento(precio float64) float64 {
	return precio
}

func (sd SinDescuento) DescribirDescuento() string {
	return "Sin descuento"
}

type DescuentoPorcentaje struct {
	Porcentaje float64
}

func (dp DescuentoPorcentaje) AplicarDescuento(precio float64) float64 {
	return precio * (1 - dp.Porcentaje/100)
}

func (dp DescuentoPorcentaje) DescribirDescuento() string {
	return fmt.Sprintf("%.0f%% de descuento", dp.Porcentaje)
}

type DescuentoFijo struct {
	Cantidad float64
}

func (df DescuentoFijo) AplicarDescuento(precio float64) float64 {
	resultado := precio - df.Cantidad
	if resultado < 0 {
		return 0
	}
	return resultado
}

func (df DescuentoFijo) DescribirDescuento() string {
	return fmt.Sprintf("$%.2f de descuento fijo", df.Cantidad)
}

type DescuentoPorCantidad struct {
	CantidadMinima int
	Descuento      float64
	cantidadItems  int
}

func (dpc *DescuentoPorCantidad) SetCantidadItems(cantidad int) {
	dpc.cantidadItems = cantidad
}

func (dpc DescuentoPorCantidad) AplicarDescuento(precio float64) float64 {
	if dpc.cantidadItems >= dpc.CantidadMinima {
		return precio * (1 - dpc.Descuento/100)
	}
	return precio
}

func (dpc DescuentoPorCantidad) DescribirDescuento() string {
	return fmt.Sprintf("%.0f%% descuento si compras %d+ items",
		dpc.Descuento, dpc.CantidadMinima)
}

type CalculadoraPrecio struct {
	estrategia EstrategiaDescuento
}

func (cp *CalculadoraPrecio) SetEstrategia(estrategia EstrategiaDescuento) {
	cp.estrategia = estrategia
}

func (cp CalculadoraPrecio) CalcularPrecioFinal(precio float64) (float64, string) {
	if cp.estrategia == nil {
		cp.estrategia = SinDescuento{}
	}
	precioFinal := cp.estrategia.AplicarDescuento(precio)
	descripcion := cp.estrategia.DescribirDescuento()
	return precioFinal, descripcion
}

func ejercicio6() {
	fmt.Println("=== Ejercicio 6: Strategy Pattern - Descuentos ===")

	calculadora := CalculadoraPrecio{}
	precioOriginal := 100.0

	estrategias := []EstrategiaDescuento{
		SinDescuento{},
		DescuentoPorcentaje{Porcentaje: 15},
		DescuentoPorcentaje{Porcentaje: 25},
		DescuentoFijo{Cantidad: 20},
		DescuentoFijo{Cantidad: 150}, // Más que el precio original
	}

	// Estrategia especial que necesita configuración
	descuentoCantidad := &DescuentoPorCantidad{
		CantidadMinima: 3,
		Descuento:      20,
	}
	descuentoCantidad.SetCantidadItems(5) // 5 items en el carrito
	estrategias = append(estrategias, descuentoCantidad)

	fmt.Printf("💰 Precio original: $%.2f\n\n", precioOriginal)

	for i, estrategia := range estrategias {
		calculadora.SetEstrategia(estrategia)
		precioFinal, descripcion := calculadora.CalcularPrecioFinal(precioOriginal)
		ahorro := precioOriginal - precioFinal

		fmt.Printf("%d. %s\n", i+1, descripcion)
		fmt.Printf("   Precio final: $%.2f\n", precioFinal)
		fmt.Printf("   Ahorro: $%.2f\n\n", ahorro)
	}

	fmt.Println("✅ Ejercicio 6 completado\n")
}

// ========================================
// Ejercicio 7: Observer Pattern - Sistema de Notificaciones
// ========================================

type Observer interface {
	Actualizar(evento string, datos interface{})
	ObtenerID() string
}

type Sujeto interface {
	Suscribir(observer Observer)
	Desuscribir(id string)
	Notificar(evento string, datos interface{})
}

type GestorEventos struct {
	observadores map[string]Observer
}

func NewGestorEventos() *GestorEventos {
	return &GestorEventos{
		observadores: make(map[string]Observer),
	}
}

func (ge *GestorEventos) Suscribir(observer Observer) {
	ge.observadores[observer.ObtenerID()] = observer
	fmt.Printf("✅ %s suscrito a eventos\n", observer.ObtenerID())
}

func (ge *GestorEventos) Desuscribir(id string) {
	delete(ge.observadores, id)
	fmt.Printf("❌ %s desuscrito de eventos\n", id)
}

func (ge *GestorEventos) Notificar(evento string, datos interface{}) {
	fmt.Printf("📢 Notificando evento: %s\n", evento)
	for _, observer := range ge.observadores {
		observer.Actualizar(evento, datos)
	}
	fmt.Println()
}

type ObservadorEmail struct {
	email string
}

func (oe ObservadorEmail) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📧 [EMAIL a %s] Evento '%s': %v\n", oe.email, evento, datos)
}

func (oe ObservadorEmail) ObtenerID() string {
	return "email_" + oe.email
}

type ObservadorSMS struct {
	telefono string
}

func (os ObservadorSMS) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📱 [SMS a %s] Evento '%s': %v\n", os.telefono, evento, datos)
}

func (os ObservadorSMS) ObtenerID() string {
	return "sms_" + os.telefono
}

type ObservadorAnalytics struct {
	servicio string
}

func (oa ObservadorAnalytics) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📊 [ANALYTICS-%s] Tracking evento '%s': %v\n",
		oa.servicio, evento, datos)
}

func (oa ObservadorAnalytics) ObtenerID() string {
	return "analytics_" + oa.servicio
}

func ejercicio7() {
	fmt.Println("=== Ejercicio 7: Observer Pattern ===")

	gestor := NewGestorEventos()

	// Crear observadores
	emailObs := ObservadorEmail{email: "admin@empresa.com"}
	smsObs := ObservadorSMS{telefono: "+1234567890"}
	analyticsObs := ObservadorAnalytics{servicio: "GoogleAnalytics"}

	// Suscribir observadores
	gestor.Suscribir(emailObs)
	gestor.Suscribir(smsObs)
	gestor.Suscribir(analyticsObs)
	fmt.Println()

	// Generar eventos
	gestor.Notificar("usuario_registro", map[string]interface{}{
		"user_id": 123,
		"email":   "nuevo@usuario.com",
		"plan":    "premium",
	})

	gestor.Notificar("compra_completada", map[string]interface{}{
		"order_id": 456,
		"amount":   199.99,
		"items":    []string{"producto1", "producto2"},
	})

	// Desuscribir un observador
	gestor.Desuscribir("sms_+1234567890")
	fmt.Println()

	gestor.Notificar("usuario_logout", map[string]interface{}{
		"user_id":          123,
		"session_duration": "45 minutos",
	})

	fmt.Println("✅ Ejercicio 7 completado\n")
}

// ========================================
// Ejercicio 8: Factory Pattern - Conexiones de Base de Datos
// ========================================

type BaseDatos interface {
	Conectar() string
	EjecutarConsulta(sql string) string
	Cerrar() string
	ObtenerTipo() string
}

type MySQL struct {
	host   string
	puerto int
}

func (m MySQL) Conectar() string {
	return fmt.Sprintf("✅ Conectado a MySQL en %s:%d", m.host, m.puerto)
}

func (m MySQL) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 MySQL ejecutando: %s → Resultado: [rows: 3]", sql)
}

func (m MySQL) Cerrar() string {
	return "❌ Conexión MySQL cerrada"
}

func (m MySQL) ObtenerTipo() string {
	return "MySQL"
}

type PostgreSQL struct {
	host      string
	baseDatos string
}

func (p PostgreSQL) Conectar() string {
	return fmt.Sprintf("✅ Conectado a PostgreSQL: %s/%s", p.host, p.baseDatos)
}

func (p PostgreSQL) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 PostgreSQL ejecutando: %s → Resultado: {affected: 2}", sql)
}

func (p PostgreSQL) Cerrar() string {
	return "❌ Conexión PostgreSQL cerrada"
}

func (p PostgreSQL) ObtenerTipo() string {
	return "PostgreSQL"
}

type MongoDB struct {
	uri       string
	coleccion string
}

func (m MongoDB) Conectar() string {
	return fmt.Sprintf("✅ Conectado a MongoDB: %s (colección: %s)", m.uri, m.coleccion)
}

func (m MongoDB) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 MongoDB ejecutando query: %s → Resultado: [{_id: 1}, {_id: 2}]", sql)
}

func (m MongoDB) Cerrar() string {
	return "❌ Conexión MongoDB cerrada"
}

func (m MongoDB) ObtenerTipo() string {
	return "MongoDB"
}

type FactoryBaseDatos interface {
	CrearBaseDatos(tipo string, config map[string]string) BaseDatos
}

type FactoryDB struct{}

func (f FactoryDB) CrearBaseDatos(tipo string, config map[string]string) BaseDatos {
	tipoLower := strings.ToLower(tipo)

	switch tipoLower {
	case "mysql":
		host := config["host"]
		if host == "" {
			host = "localhost"
		}
		return MySQL{host: host, puerto: 3306}

	case "postgresql", "postgres":
		host := config["host"]
		if host == "" {
			host = "localhost"
		}
		baseDatos := config["database"]
		if baseDatos == "" {
			baseDatos = "defaultdb"
		}
		return PostgreSQL{host: host, baseDatos: baseDatos}

	case "mongodb", "mongo":
		uri := config["uri"]
		if uri == "" {
			uri = "mongodb://localhost:27017"
		}
		coleccion := config["collection"]
		if coleccion == "" {
			coleccion = "documents"
		}
		return MongoDB{uri: uri, coleccion: coleccion}

	default:
		// Fallback a MySQL
		fmt.Printf("⚠️  Tipo '%s' no reconocido, usando MySQL por defecto\n", tipo)
		return MySQL{host: "localhost", puerto: 3306}
	}
}

func ProbarBaseDatos(factory FactoryBaseDatos, tipo string, config map[string]string) {
	fmt.Printf("=== Probando %s ===\n", strings.ToUpper(tipo))
	fmt.Printf("Configuración: %v\n", config)

	// Crear base de datos usando factory
	db := factory.CrearBaseDatos(tipo, config)

	// Probar operaciones
	fmt.Printf("Tipo: %s\n", db.ObtenerTipo())
	fmt.Println(db.Conectar())

	consultas := []string{
		"SELECT * FROM usuarios",
		"INSERT INTO productos (nombre, precio) VALUES ('Laptop', 1200)",
	}

	for _, consulta := range consultas {
		fmt.Println(db.EjecutarConsulta(consulta))
	}

	fmt.Println(db.Cerrar())
	fmt.Println()
}

func ejercicio8() {
	fmt.Println("=== Ejercicio 8: Factory Pattern ===")

	factory := FactoryDB{}

	configuraciones := []struct {
		tipo   string
		config map[string]string
	}{
		{
			tipo: "mysql",
			config: map[string]string{
				"host": "mysql-prod.empresa.com",
			},
		},
		{
			tipo: "postgresql",
			config: map[string]string{
				"host":     "postgres-cluster.empresa.com",
				"database": "ecommerce",
			},
		},
		{
			tipo: "mongodb",
			config: map[string]string{
				"uri":        "mongodb://mongo-cluster.empresa.com:27017",
				"collection": "users",
			},
		},
		{
			tipo: "redis", // Tipo no soportado
			config: map[string]string{
				"host": "redis-cache.empresa.com",
			},
		},
	}

	for _, config := range configuraciones {
		ProbarBaseDatos(factory, config.tipo, config.config)
	}

	fmt.Println("✅ Ejercicio 8 completado\n")
}

// ========================================
// Función principal con todas las soluciones
// ========================================

func main() {
	fmt.Println("🎪 Soluciones: Interfaces Básicas en Go")
	fmt.Println("======================================")
	fmt.Println()

	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()

	fmt.Println("🎉 ¡Todas las soluciones ejecutadas correctamente!")
	fmt.Println("\n💡 Estos ejemplos demuestran el poder de las interfaces en Go:")
	fmt.Println("   • Polimorfismo elegante")
	fmt.Println("   • Patrones de diseño robustos")
	fmt.Println("   • Código flexible y mantenible")
	fmt.Println("   • Testing fácil con mocks")
}
