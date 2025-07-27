// Lección 13: Interfaces Básicas en Go
// Soluciones completas de todos los ejercicios

/*
INSTRUCCIONES PARA USAR ESTE ARCHIVO:

Este archivo contiene las soluciones completas de todos los ejercicios.
Para ejecutar las soluciones:

1. Desde la terminal:
   go run soluciones.go

2. Las soluciones están organizadas por ejercicio y puedes ejecutar
   cada una individualmente modificando la función main.

3. Para comparar con los ejercicios:
   - Abre ejercicios.go en una ventana
   - Abre soluciones.go en otra ventana
   - Compara las implementaciones lado a lado

Cada ejercicio está completamente implementado y funcionando.
*/

package soluciones

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
type FormaSol interface {
	Area() float64
	Perimetro() float64
	Descripcion() string
}

// Solución: Struct Rectangulo con campos requeridos
type RectanguloSol struct {
	Ancho float64
	Alto  float64
}

// Solución: Implementación de métodos para Rectangulo
func (r RectanguloSol) Area() float64 {
	return r.Ancho * r.Alto
}

func (r RectanguloSol) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

func (r RectanguloSol) Descripcion() string {
	return fmt.Sprintf("Rectángulo de %.1f x %.1f", r.Ancho, r.Alto)
}

// Solución: Struct Circulo con campo requerido
type CirculoSol struct {
	Radio float64
}

// Solución: Implementación de métodos para Circulo
func (c CirculoSol) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c CirculoSol) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c CirculoSol) Descripcion() string {
	return fmt.Sprintf("Círculo con radio %.1f", c.Radio)
}

// Solución: Función que maneja cualquier forma
func MostrarInformacionFormaSol(f FormaSol) {
	fmt.Printf("🔶 %s\n", f.Descripcion())
	fmt.Printf("   Área: %.2f\n", f.Area())
	fmt.Printf("   Perímetro: %.2f\n", f.Perimetro())
	fmt.Println()
}

func ejercicio1Solucion() {
	fmt.Println("=== Ejercicio 1: Sistema de Formas Geométricas ===")

	rectangulo := RectanguloSol{Ancho: 5.0, Alto: 3.0}
	circulo := CirculoSol{Radio: 2.5}

	formas := []FormaSol{rectangulo, circulo}

	for _, forma := range formas {
		MostrarInformacionFormaSol(forma)
	}

	fmt.Println("✅ Ejercicio 1 completado\n")
}

// ========================================
// Ejercicio 2: Polimorfismo - Sistema de Animales
// ========================================

type AnimalSol interface {
	HacerSonido() string
	Moverse() string
	Comer(comida string) string
}

type PerroSol struct {
	Nombre string
	Raza   string
}

func (p PerroSol) HacerSonido() string {
	return fmt.Sprintf("%s hace: ¡Guau guau!", p.Nombre)
}

func (p PerroSol) Moverse() string {
	return fmt.Sprintf("%s corre alegremente", p.Nombre)
}

func (p PerroSol) Comer(comida string) string {
	return fmt.Sprintf("%s está comiendo %s con mucha energía", p.Nombre, comida)
}

type GatoSol struct {
	Nombre string
	Color  string
}

func (g GatoSol) HacerSonido() string {
	return fmt.Sprintf("%s hace: Miau miau", g.Nombre)
}

func (g GatoSol) Moverse() string {
	return fmt.Sprintf("%s camina sigilosamente", g.Nombre)
}

func (g GatoSol) Comer(comida string) string {
	return fmt.Sprintf("%s está comiendo %s elegantemente", g.Nombre, comida)
}

type PajaroSol struct {
	Nombre     string
	Especie    string
	PuedeVolar bool
}

func (p PajaroSol) HacerSonido() string {
	return fmt.Sprintf("%s hace: ¡Pío pío!", p.Nombre)
}

func (p PajaroSol) Moverse() string {
	if p.PuedeVolar {
		return fmt.Sprintf("%s vuela graciosamente", p.Nombre)
	}
	return fmt.Sprintf("%s camina saltando", p.Nombre)
}

func (p PajaroSol) Comer(comida string) string {
	return fmt.Sprintf("%s está picoteando %s", p.Nombre, comida)
}

func CuidarAnimalSol(a AnimalSol) {
	fmt.Printf("🐾 Cuidando a un animal:\n")
	fmt.Printf("   Sonido: %s\n", a.HacerSonido())
	fmt.Printf("   Movimiento: %s\n", a.Moverse())
	fmt.Printf("   Alimentación: %s\n", a.Comer("comida"))
	fmt.Println()
}

func ejercicio2Solucion() {
	fmt.Println("=== Ejercicio 2: Sistema de Animales ===")

	perro := PerroSol{Nombre: "Max", Raza: "Golden Retriever"}
	gato := GatoSol{Nombre: "Luna", Color: "Negro"}
	pajaro := PajaroSol{Nombre: "Pipo", Especie: "Canario", PuedeVolar: true}

	animales := []AnimalSol{perro, gato, pajaro}

	for _, animal := range animales {
		CuidarAnimalSol(animal)
	}

	fmt.Println("✅ Ejercicio 2 completado\n")
}

// ========================================
// Ejercicio 3: Interfaces Estándar - fmt.Stringer y sort.Interface
// ========================================

type ProductoSol struct {
	Nombre    string
	Precio    float64
	Categoria string
}

func (p ProductoSol) String() string {
	return fmt.Sprintf("%s ($%.2f) - %s", p.Nombre, p.Precio, p.Categoria)
}

type ProductosPorPrecioSol []ProductoSol

func (p ProductosPorPrecioSol) Len() int {
	return len(p)
}

func (p ProductosPorPrecioSol) Less(i, j int) bool {
	return p[i].Precio < p[j].Precio
}

func (p ProductosPorPrecioSol) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ejercicio3Solucion() {
	fmt.Println("=== Ejercicio 3: Interfaces Estándar ===")

	productos := ProductosPorPrecioSol{
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

func ProcesarDatoSol(dato interface{}) {
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

func ejercicio4Solucion() {
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
		ProcesarDatoSol(dato)
	}

	fmt.Println("✅ Ejercicio 4 completado\n")
}

// ========================================
// Ejercicio 5: Empty Interface - Sistema de Logging
// ========================================

type LoggerSol interface {
	Log(level string, message string, data interface{})
}

type ConsoleLoggerSol struct{}

func (cl ConsoleLoggerSol) Log(level string, message string, data interface{}) {
	fmt.Printf("[%s] %s", strings.ToUpper(level), message)
	if data != nil {
		fmt.Printf(": %v", data)
	}
	fmt.Println()
}

type FileLoggerSol struct {
	archivo string
}

func (fl FileLoggerSol) Log(level string, message string, data interface{}) {
	fmt.Printf("[FILE:%s] [%s] %s", fl.archivo, strings.ToUpper(level), message)
	if data != nil {
		fmt.Printf(": %v", data)
	}
	fmt.Println()
}

func LogearEventoSol(logger LoggerSol, evento string, datos ...interface{}) {
	logger.Log("INFO", fmt.Sprintf("Evento: %s", evento), nil)

	for i, dato := range datos {
		logger.Log("DEBUG", fmt.Sprintf("Dato %d", i+1), dato)
	}
}

func ejercicio5Solucion() {
	fmt.Println("=== Ejercicio 5: Sistema de Logging ===")

	consoleLogger := ConsoleLoggerSol{}
	fileLogger := FileLoggerSol{archivo: "app.log"}

	fmt.Println("📝 Console Logger:")
	LogearEventoSol(consoleLogger, "usuario_login",
		map[string]interface{}{"user_id": 123, "ip": "192.168.1.1"},
		"timestamp: 2025-01-15T10:30:00Z",
		[]string{"session", "auth", "web"})

	fmt.Println("\n📁 File Logger:")
	LogearEventoSol(fileLogger, "compra_realizada",
		map[string]interface{}{"order_id": 456, "amount": 99.99},
		[]string{"laptop", "mouse"},
		true)

	fmt.Println("✅ Ejercicio 5 completado\n")
}

// ========================================
// Ejercicio 6: Strategy Pattern - Sistema de Descuentos
// ========================================

type EstrategiaDescuentoSol interface {
	AplicarDescuento(precio float64) float64
	DescribirDescuento() string
}

type SinDescuentoSol struct{}

func (sd SinDescuentoSol) AplicarDescuento(precio float64) float64 {
	return precio
}

func (sd SinDescuentoSol) DescribirDescuento() string {
	return "Sin descuento"
}

type DescuentoPorcentajeSol struct {
	Porcentaje float64
}

func (dp DescuentoPorcentajeSol) AplicarDescuento(precio float64) float64 {
	return precio * (1 - dp.Porcentaje/100)
}

func (dp DescuentoPorcentajeSol) DescribirDescuento() string {
	return fmt.Sprintf("%.0f%% de descuento", dp.Porcentaje)
}

type DescuentoFijoSol struct {
	Cantidad float64
}

func (df DescuentoFijoSol) AplicarDescuento(precio float64) float64 {
	resultado := precio - df.Cantidad
	if resultado < 0 {
		return 0
	}
	return resultado
}

func (df DescuentoFijoSol) DescribirDescuento() string {
	return fmt.Sprintf("$%.2f de descuento fijo", df.Cantidad)
}

type DescuentoPorCantidadSol struct {
	CantidadMinima int
	Descuento      float64
	cantidadItems  int
}

func (dpc *DescuentoPorCantidadSol) SetCantidadItems(cantidad int) {
	dpc.cantidadItems = cantidad
}

func (dpc DescuentoPorCantidadSol) AplicarDescuento(precio float64) float64 {
	if dpc.cantidadItems >= dpc.CantidadMinima {
		return precio * (1 - dpc.Descuento/100)
	}
	return precio
}

func (dpc DescuentoPorCantidadSol) DescribirDescuento() string {
	return fmt.Sprintf("%.0f%% descuento si compras %d+ items",
		dpc.Descuento, dpc.CantidadMinima)
}

type CalculadoraPrecioSol struct {
	estrategia EstrategiaDescuentoSol
}

func (cp *CalculadoraPrecioSol) SetEstrategia(estrategia EstrategiaDescuentoSol) {
	cp.estrategia = estrategia
}

func (cp CalculadoraPrecioSol) CalcularPrecioFinal(precio float64) (float64, string) {
	if cp.estrategia == nil {
		cp.estrategia = SinDescuentoSol{}
	}
	precioFinal := cp.estrategia.AplicarDescuento(precio)
	descripcion := cp.estrategia.DescribirDescuento()
	return precioFinal, descripcion
}

func ejercicio6Solucion() {
	fmt.Println("=== Ejercicio 6: Strategy Pattern - Descuentos ===")

	calculadora := CalculadoraPrecioSol{}
	precioOriginal := 100.0

	estrategias := []EstrategiaDescuentoSol{
		SinDescuentoSol{},
		DescuentoPorcentajeSol{Porcentaje: 15},
		DescuentoPorcentajeSol{Porcentaje: 25},
		DescuentoFijoSol{Cantidad: 20},
		DescuentoFijoSol{Cantidad: 150}, // Más que el precio original
	}

	// Estrategia especial que necesita configuración
	descuentoCantidad := &DescuentoPorCantidadSol{
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

type ObserverSol interface {
	Actualizar(evento string, datos interface{})
	ObtenerID() string
}

type SujetoSol interface {
	Suscribir(observer ObserverSol)
	Desuscribir(id string)
	Notificar(evento string, datos interface{})
}

type GestorEventosSol struct {
	observadores map[string]ObserverSol
}

func NewGestorEventosSol() *GestorEventosSol {
	return &GestorEventosSol{
		observadores: make(map[string]ObserverSol),
	}
}

func (ge *GestorEventosSol) Suscribir(observer ObserverSol) {
	ge.observadores[observer.ObtenerID()] = observer
	fmt.Printf("✅ %s suscrito a eventos\n", observer.ObtenerID())
}

func (ge *GestorEventosSol) Desuscribir(id string) {
	delete(ge.observadores, id)
	fmt.Printf("❌ %s desuscrito de eventos\n", id)
}

func (ge *GestorEventosSol) Notificar(evento string, datos interface{}) {
	fmt.Printf("📢 Notificando evento: %s\n", evento)
	for _, observer := range ge.observadores {
		observer.Actualizar(evento, datos)
	}
	fmt.Println()
}

type ObservadorEmailSol struct {
	email string
}

func (oe ObservadorEmailSol) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📧 [EMAIL a %s] Evento '%s': %v\n", oe.email, evento, datos)
}

func (oe ObservadorEmailSol) ObtenerID() string {
	return "email_" + oe.email
}

type ObservadorSMSSol struct {
	telefono string
}

func (os ObservadorSMSSol) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📱 [SMS a %s] Evento '%s': %v\n", os.telefono, evento, datos)
}

func (os ObservadorSMSSol) ObtenerID() string {
	return "sms_" + os.telefono
}

type ObservadorAnalyticsSol struct {
	servicio string
}

func (oa ObservadorAnalyticsSol) Actualizar(evento string, datos interface{}) {
	fmt.Printf("  📊 [ANALYTICS-%s] Tracking evento '%s': %v\n",
		oa.servicio, evento, datos)
}

func (oa ObservadorAnalyticsSol) ObtenerID() string {
	return "analytics_" + oa.servicio
}

func ejercicio7Solucion() {
	fmt.Println("=== Ejercicio 7: Observer Pattern ===")

	gestor := NewGestorEventosSol()

	// Crear observadores
	emailObs := ObservadorEmailSol{email: "admin@empresa.com"}
	smsObs := ObservadorSMSSol{telefono: "+1234567890"}
	analyticsObs := ObservadorAnalyticsSol{servicio: "GoogleAnalytics"}

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

type BaseDatosSol interface {
	Conectar() string
	EjecutarConsulta(sql string) string
	Cerrar() string
	ObtenerTipo() string
}

type MySQLSol struct {
	host   string
	puerto int
}

func (m MySQLSol) Conectar() string {
	return fmt.Sprintf("✅ Conectado a MySQL en %s:%d", m.host, m.puerto)
}

func (m MySQLSol) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 MySQL ejecutando: %s → Resultado: [rows: 3]", sql)
}

func (m MySQLSol) Cerrar() string {
	return "❌ Conexión MySQL cerrada"
}

func (m MySQLSol) ObtenerTipo() string {
	return "MySQL"
}

type PostgreSQLSol struct {
	host      string
	baseDatos string
}

func (p PostgreSQLSol) Conectar() string {
	return fmt.Sprintf("✅ Conectado a PostgreSQL: %s/%s", p.host, p.baseDatos)
}

func (p PostgreSQLSol) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 PostgreSQL ejecutando: %s → Resultado: {affected: 2}", sql)
}

func (p PostgreSQLSol) Cerrar() string {
	return "❌ Conexión PostgreSQL cerrada"
}

func (p PostgreSQLSol) ObtenerTipo() string {
	return "PostgreSQL"
}

type MongoDBSol struct {
	uri       string
	coleccion string
}

func (m MongoDBSol) Conectar() string {
	return fmt.Sprintf("✅ Conectado a MongoDB: %s (colección: %s)", m.uri, m.coleccion)
}

func (m MongoDBSol) EjecutarConsulta(sql string) string {
	return fmt.Sprintf("🔍 MongoDB ejecutando query: %s → Resultado: [{_id: 1}, {_id: 2}]", sql)
}

func (m MongoDBSol) Cerrar() string {
	return "❌ Conexión MongoDB cerrada"
}

func (m MongoDBSol) ObtenerTipo() string {
	return "MongoDB"
}

type FactoryBaseDatosSol interface {
	CrearBaseDatos(tipo string, config map[string]string) BaseDatosSol
}

type FactoryDBSol struct{}

func (f FactoryDBSol) CrearBaseDatos(tipo string, config map[string]string) BaseDatosSol {
	tipoLower := strings.ToLower(tipo)

	switch tipoLower {
	case "mysql":
		host := config["host"]
		if host == "" {
			host = "localhost"
		}
		return MySQLSol{host: host, puerto: 3306}

	case "postgresql", "postgres":
		host := config["host"]
		if host == "" {
			host = "localhost"
		}
		baseDatos := config["database"]
		if baseDatos == "" {
			baseDatos = "defaultdb"
		}
		return PostgreSQLSol{host: host, baseDatos: baseDatos}

	case "mongodb", "mongo":
		uri := config["uri"]
		if uri == "" {
			uri = "mongodb://localhost:27017"
		}
		coleccion := config["collection"]
		if coleccion == "" {
			coleccion = "documents"
		}
		return MongoDBSol{uri: uri, coleccion: coleccion}

	default:
		// Fallback a MySQL
		fmt.Printf("⚠️  Tipo '%s' no reconocido, usando MySQL por defecto\n", tipo)
		return MySQLSol{host: "localhost", puerto: 3306}
	}
}

func ProbarBaseDatosSol(factory FactoryBaseDatosSol, tipo string, config map[string]string) {
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

func ejercicio8Solucion() {
	fmt.Println("=== Ejercicio 8: Factory Pattern ===")

	factory := FactoryDBSol{}

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
		ProbarBaseDatosSol(factory, config.tipo, config.config)
	}

	fmt.Println("✅ Ejercicio 8 completado\n")
}

// ========================================
// Función principal con todas las soluciones
// ========================================

func mainSoluciones() {
	fmt.Println("🎪 Soluciones: Interfaces Básicas en Go")
	fmt.Println("======================================")
	fmt.Println()

	ejercicio1Solucion()
	ejercicio2Solucion()
	ejercicio3Solucion()
	ejercicio4Solucion()
	ejercicio5Solucion()
	ejercicio6Solucion()
	ejercicio7Solucion()
	ejercicio8Solucion()

	fmt.Println("🎉 ¡Todas las soluciones ejecutadas correctamente!")
	fmt.Println("\n💡 Estos ejemplos demuestran el poder de las interfaces en Go:")
	fmt.Println("   • Polimorfismo elegante")
	fmt.Println("   • Patrones de diseño robustos")
	fmt.Println("   • Código flexible y mantenible")
	fmt.Println("   • Testing fácil con mocks")
}

// Ejecutar todas las soluciones
func main() {
	mainSoluciones()
}
