// Archivo: ejercicios.go
// Lección 11: Structs - Ejercicios Prácticos
// Cubre: declaración, inicialización, métodos, embedding, tags, validación

package main

import (
	"fmt"
	"time"
)

// ==============================================
// EJERCICIO 1: Struct Básico - Sistema de Libros
// ==============================================

// TODO: Define un struct Libro con los campos:
// - ID (int)
// - Titulo (string)
// - Autor (string)
// - Paginas (int)
// - Precio (float64)
// - Disponible (bool)
// - FechaPublicacion (time.Time)

type Libro struct {
	// Tu código aquí
}

// TODO: Implementa un constructor NewLibro que reciba titulo, autor, paginas, precio
// y establezca ID automáticamente, Disponible como true y FechaPublicacion como ahora
func NewLibro(titulo, autor string, paginas int, precio float64) Libro {
	// Tu código aquí
	return Libro{}
}

// TODO: Implementa un método String() que devuelva una representación legible del libro
func (l Libro) String() string {
	// Tu código aquí
	return ""
}

// TODO: Implementa un método Descuento() que aplique un descuento al precio
func (l *Libro) Descuento(porcentaje float64) {
	// Tu código aquí
}

// TODO: Implementa un método EsCaroq() que retorne true si el libro cuesta más de 25.00
func (l Libro) EsCaro() bool {
	// Tu código aquí
	return false
}

func ejercicio1() {
	fmt.Println("=== EJERCICIO 1: Struct Básico ===")

	// Crear algunos libros
	libro1 := NewLibro("El Quijote", "Cervantes", 863, 29.99)
	libro2 := NewLibro("Go Programming", "Kernighan", 380, 45.50)

	fmt.Println("Libros creados:")
	fmt.Println(libro1)
	fmt.Println(libro2)

	// Aplicar descuentos
	libro1.Descuento(0.10) // 10% descuento
	libro2.Descuento(0.15) // 15% descuento

	fmt.Println("\nDespués de descuentos:")
	fmt.Println(libro1)
	fmt.Println(libro2)

	// Verificar si son caros
	fmt.Printf("\n¿El Quijote es caro? %v\n", libro1.EsCaro())
	fmt.Printf("¿Go Programming es caro? %v\n", libro2.EsCaro())
}

// ==============================================
// EJERCICIO 2: Embedding - Sistema de Empleados
// ==============================================

// TODO: Define un struct Persona con campos básicos
type Persona struct {
	// Tu código aquí - agregar campos: Nombre, Apellido, Edad, Email
	Nombre   string
	Apellido string
	Edad     int
	Email    string
}

// TODO: Define un struct Direccion
type Direccion struct {
	// Tu código aquí - agregar campos: Calle, Ciudad, Estado, CodigoPostal, Pais
	Calle        string
	Ciudad       string
	Estado       string
	CodigoPostal string
	Pais         string
}

// TODO: Define un struct Empleado que embeba Persona y contenga Direccion
type Empleado struct {
	// Tu código aquí - usar embedding para Persona
	Persona           // Embedding
	Direccion         // Campo struct
	ID                int
	Departamento      string
	SalarioMensual    float64
	FechaContratacion time.Time
}

// TODO: Implementa un método NombreCompleto() para Persona
func (p Persona) NombreCompleto() string {
	// Tu código aquí
	return ""
}

// TODO: Implementa un método DireccionCompleta() para Direccion
func (d Direccion) DireccionCompleta() string {
	// Tu código aquí
	return ""
}

// TODO: Implementa un método CalcularSalarioAnual() para Empleado
func (e Empleado) CalcularSalarioAnual() float64 {
	// Tu código aquí
	return 0
}

// TODO: Implementa un método AumentarSalario() para Empleado
func (e *Empleado) AumentarSalario(porcentaje float64) {
	// Tu código aquí
}

func ejercicio2() {
	fmt.Println("\n=== EJERCICIO 2: Embedding ===")

	// Crear empleado
	empleado := Empleado{
		// Tu código aquí - inicializar todos los campos
	}

	fmt.Printf("Empleado: %s\n", empleado.NombreCompleto())
	fmt.Printf("Dirección: %s\n", empleado.Direccion.DireccionCompleta())
	fmt.Printf("Salario anual: $%.2f\n", empleado.CalcularSalarioAnual())

	// Aumentar salario
	empleado.AumentarSalario(0.10) // 10% de aumento
	fmt.Printf("Nuevo salario anual: $%.2f\n", empleado.CalcularSalarioAnual())
}

// ==============================================
// EJERCICIO 3: Struct Tags - Configuración JSON
// ==============================================

// TODO: Define un struct ConfiguracionApp con tags JSON apropiados
// Campos: AppName, Version, Port, Debug, DatabaseURL, Features (slice), Settings (map)
type ConfiguracionApp struct {
	// Tu código aquí - incluir tags JSON
}

// TODO: Implementa un método Validar() que verifique que los campos requeridos no estén vacíos
func (c ConfiguracionApp) Validar() error {
	// Tu código aquí
	return nil
}

// TODO: Implementa un método String() que muestre la configuración sin datos sensibles
func (c ConfiguracionApp) String() string {
	// Tu código aquí
	return ""
}

func ejercicio3() {
	fmt.Println("\n=== EJERCICIO 3: Struct Tags ===")

	config := ConfiguracionApp{
		// Tu código aquí - inicializar configuración
	}

	fmt.Println("Configuración:")
	fmt.Println(config)

	if err := config.Validar(); err != nil {
		fmt.Printf("Error de validación: %v\n", err)
	} else {
		fmt.Println("✅ Configuración válida")
	}

	// TODO: Convierte el struct a JSON y muéstralo
	// jsonData, _ := json.MarshalIndent(config, "", "  ")
	// fmt.Println("JSON:")
	// fmt.Println(string(jsonData))
}

// ==============================================
// EJERCICIO 4: Múltiple Embedding - Sistema de Vehículos
// ==============================================

// TODO: Define structs base para embedding múltiple
type Motor struct {
	// Potencia, Combustible, Cilindros
}

type Ruedas struct {
	// Cantidad, Tamaño, Tipo
}

type Identificacion struct {
	// Marca, Modelo, Año, NumeroSerie
}

// TODO: Define un struct Vehiculo que embeba los anteriores
type Vehiculo struct {
	// Tu código aquí - embedding múltiple
}

// TODO: Implementa métodos para cada struct embebido
func (m Motor) Descripcion() string {
	// Tu código aquí
	return ""
}

func (r Ruedas) Descripcion() string {
	// Tu código aquí
	return ""
}

func (i Identificacion) Descripcion() string {
	// Tu código aquí
	return ""
}

// TODO: Implementa un método ResumenCompleto() para Vehiculo
func (v Vehiculo) ResumenCompleto() string {
	// Tu código aquí - usar los métodos embebidos
	return ""
}

func ejercicio4() {
	fmt.Println("\n=== EJERCICIO 4: Múltiple Embedding ===")

	vehiculo := Vehiculo{
		// Tu código aquí - inicializar todos los campos embebidos
	}

	fmt.Println("Vehículo:")
	fmt.Println(vehiculo.ResumenCompleto())
}

// ==============================================
// EJERCICIO 5: Structs Anónimos - Procesamiento de Datos
// ==============================================

func ejercicio5() {
	fmt.Println("\n=== EJERCICIO 5: Structs Anónimos ===")

	// TODO: Crea un slice de structs anónimos que representen productos
	// Campos: ID, Nombre, Categoria, Precio, Stock
	productos := []struct {
		// Tu código aquí
	}{
		// Tu código aquí - 5 productos de ejemplo
	}

	fmt.Println("Productos:")
	for _, producto := range productos {
		// TODO: Mostrar cada producto formateado
		_ = producto // Evitar error "declared and not used"
	}

	// TODO: Crea un map de structs anónimos para estadísticas por categoría
	// Campos: TotalProductos, PrecioPromedio, StockTotal
	estadisticas := map[string]struct {
		// Tu código aquí
	}{}

	// TODO: Calcula estadísticas por categoría
	for _, producto := range productos {
		// Tu código aquí - agrupar y calcular
		_ = producto // Evitar error "declared and not used"
	}

	fmt.Println("\nEstadísticas por categoría:")
	for categoria, stats := range estadisticas {
		// TODO: Mostrar estadísticas formateadas
		_ = categoria
		_ = stats
	}
}

// ==============================================
// EJERCICIO 6: Factory Pattern - Sistema de Conexiones
// ==============================================

// TODO: Define un struct ConexionDB
type ConexionDB struct {
	// Tipo, Host, Puerto, Database, Usuario, Password, Pool, Timeout
}

// TODO: Define un Factory para crear conexiones
type DBFactory struct {
	// configuraciones predefinidas, pool de conexiones, etc.
}

// TODO: Implementa un constructor para DBFactory
func NewDBFactory() *DBFactory {
	// Tu código aquí
	return nil
}

// TODO: Implementa métodos del factory
func (f *DBFactory) CrearMySQL(database, usuario, password string) *ConexionDB {
	// Tu código aquí
	return nil
}

func (f *DBFactory) CrearPostgreSQL(database, usuario, password string) *ConexionDB {
	// Tu código aquí
	return nil
}

func (f *DBFactory) CrearRedis(database int) *ConexionDB {
	// Tu código aquí
	return nil
}

// TODO: Implementa métodos para ConexionDB
func (c *ConexionDB) Conectar() error {
	// Simular conexión
	// Tu código aquí
	return nil
}

func (c *ConexionDB) Desconectar() error {
	// Simular desconexión
	// Tu código aquí
	return nil
}

func (c *ConexionDB) String() string {
	// Tu código aquí
	return ""
}

func ejercicio6() {
	fmt.Println("\n=== EJERCICIO 6: Factory Pattern ===")

	factory := NewDBFactory()

	// Crear diferentes tipos de conexiones
	mysql := factory.CrearMySQL("tienda", "admin", "password123")
	postgres := factory.CrearPostgreSQL("analytics", "user", "secret456")
	redis := factory.CrearRedis(0)

	conexiones := []*ConexionDB{mysql, postgres, redis}

	fmt.Println("Conexiones creadas:")
	for _, conn := range conexiones {
		fmt.Println(conn)

		if err := conn.Conectar(); err != nil {
			fmt.Printf("Error conectando: %v\n", err)
		} else {
			fmt.Println("✅ Conectado")
		}
	}
}

// ==============================================
// EJERCICIO 7: Validación Avanzada - Sistema de Usuarios
// ==============================================

// TODO: Define una interfaz Validable
type Validable interface {
	// Tu código aquí
}

// TODO: Define un struct Usuario con validaciones complejas
type Usuario struct {
	// ID, Username, Email, Password, FechaNacimiento, Telefono, Direccion
	// Incluye tags de validación personalizados
	ID              int       `json:"id"`
	Username        string    `json:"username" validate:"required,min=3,max=20"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required,min=8"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Telefono        string    `json:"telefono"`
	Direccion       string    `json:"direccion"`
}

// TODO: Implementa validaciones detalladas
func (u Usuario) Validar() error {
	// Tu código aquí - validar email, password strength, edad, etc.
	return nil
}

func (u Usuario) ValidarEmail() bool {
	// Tu código aquí - regex para email
	return false
}

func (u Usuario) ValidarPassword() bool {
	// Tu código aquí - mínimo 8 chars, mayúscula, minúscula, número
	return false
}

func (u Usuario) ValidarEdad() bool {
	// Tu código aquí - mayor de 18 años
	return false
}

func ejercicio7() {
	fmt.Println("\n=== EJERCICIO 7: Validación Avanzada ===")

	usuarios := []Usuario{
		// Tu código aquí - crear 3 usuarios, algunos con errores de validación
	}

	fmt.Println("Validando usuarios:")
	for i, usuario := range usuarios {
		fmt.Printf("\nUsuario %d: %s\n", i+1, usuario.Username)

		if err := usuario.Validar(); err != nil {
			fmt.Printf("❌ Errores: %v\n", err)
		} else {
			fmt.Println("✅ Usuario válido")
		}
	}
}

// ==============================================
// EJERCICIO 8: Builder Pattern - Configuración de Servidor
// ==============================================

// TODO: Define un struct ServidorWeb con muchas opciones configurables
type ServidorWeb struct {
	// Host, Puerto, SSL, Certificado, Timeout, MaxConexiones,
	// Middleware, Rutas, Logs, CORS, etc.
}

// TODO: Define un Builder para ServidorWeb
type ServidorBuilder struct {
	// Tu código aquí
}

// TODO: Implementa el constructor del builder
func NewServidorBuilder() *ServidorBuilder {
	// Tu código aquí
	return nil
}

// TODO: Implementa métodos fluent del builder
func (b *ServidorBuilder) Host(host string) *ServidorBuilder {
	// Tu código aquí
	return b
}

func (b *ServidorBuilder) Puerto(puerto int) *ServidorBuilder {
	// Tu código aquí
	return b
}

func (b *ServidorBuilder) ConSSL(cert, key string) *ServidorBuilder {
	// Tu código aquí
	return b
}

func (b *ServidorBuilder) Timeout(timeout time.Duration) *ServidorBuilder {
	// Tu código aquí
	return b
}

func (b *ServidorBuilder) ConCORS() *ServidorBuilder {
	// Tu código aquí
	return b
}

func (b *ServidorBuilder) Build() *ServidorWeb {
	// Tu código aquí
	return nil
}

// TODO: Implementa métodos para ServidorWeb
func (s *ServidorWeb) Iniciar() error {
	// Simular inicio del servidor
	// Tu código aquí
	return nil
}

func (s *ServidorWeb) String() string {
	// Tu código aquí
	return ""
}

func ejercicio8() {
	fmt.Println("\n=== EJERCICIO 8: Builder Pattern ===")

	// TODO: Usa el builder para crear diferentes configuraciones de servidor
	servidorDev := NewServidorBuilder().
		Host("localhost").
		Puerto(8080).
		Timeout(30 * time.Second).
		Build()

	servidorProd := NewServidorBuilder().
		Host("0.0.0.0").
		Puerto(443).
		ConSSL("/path/cert.pem", "/path/key.pem").
		Timeout(60 * time.Second).
		ConCORS().
		Build()

	servidores := []*ServidorWeb{servidorDev, servidorProd}

	fmt.Println("Servidores configurados:")
	for i, servidor := range servidores {
		fmt.Printf("\nServidor %d:\n", i+1)
		fmt.Println(servidor)

		if err := servidor.Iniciar(); err != nil {
			fmt.Printf("Error iniciando: %v\n", err)
		}
	}
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func main() {
	fmt.Println("🏗️ EJERCICIOS DE STRUCTS EN GO")
	fmt.Println("==============================")

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
