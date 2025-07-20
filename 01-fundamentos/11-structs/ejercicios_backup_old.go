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
	ID               int       `json:"id"`
	Titulo           string    `json:"titulo"`
	Autor            string    `json:"autor"`
	Paginas          int       `json:"paginas"`
	Precio           float64   `json:"precio"`
	Disponible       bool      `json:"disponible"`
	FechaPublicacion time.Time `json:"fecha_publicacion"`
}

// TODO: Implementa un constructor NewLibro que reciba titulo, autor, paginas, precio
// y establezca ID automáticamente, Disponible como true y FechaPublicacion como ahora
var contadorLibros int

func NewLibro(titulo, autor string, paginas int, precio float64) Libro {
	contadorLibros++
	return Libro{
		ID:               contadorLibros,
		Titulo:           titulo,
		Autor:            autor,
		Paginas:          paginas,
		Precio:           precio,
		Disponible:       true,
		FechaPublicacion: time.Now(),
	}
}

// TODO: Implementa métodos:
// - EsCaro() bool (precio > 50)
// - AplicarDescuento(porcentaje float64)
// - String() string
func (l Libro) EsCaro() bool {
	return l.Precio > 50.0
}

func (l *Libro) AplicarDescuento(porcentaje float64) {
	l.Precio = l.Precio * (1 - porcentaje)
}

func (l Libro) String() string {
	disponible := "📚"
	if !l.Disponible {
		disponible = "📵"
	}
	return fmt.Sprintf("%s %s por %s - $%.2f", disponible, l.Titulo, l.Autor, l.Precio)
}

func ejercicio1() {
	fmt.Println("=== EJERCICIO 1: Struct Básico ===")

	// TODO: Crear 3 libros usando el constructor
	libro1 := NewLibro("El Quijote", "Cervantes", 863, 29.99)
	libro2 := NewLibro("Go Programming", "Kernighan", 380, 55.00)
	libro3 := NewLibro("Clean Code", "Martin", 464, 45.99)

	libros := []Libro{libro1, libro2, libro3}

	fmt.Println("Libros creados:")
	for _, libro := range libros {
		fmt.Printf("- %s (¿Es caro? %t)\n", libro, libro.EsCaro())
	}

	// TODO: Aplicar descuento del 15% al libro más caro
	for i := range libros {
		if libros[i].EsCaro() {
			fmt.Printf("\nAplicando 15%% descuento a: %s\n", libros[i].Titulo)
			libros[i].AplicarDescuento(0.15)
			fmt.Printf("Nuevo precio: %s\n", libros[i])
		}
	}
}

// ==============================================
// EJERCICIO 2: Embedding - Sistema de Empleados
// ==============================================

// TODO: Define struct Persona (base para embedding)
type Persona struct {
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Edad     int       `json:"edad"`
	Email    string    `json:"email"`
	Telefono string    `json:"telefono"`
	FechaNac time.Time `json:"fecha_nacimiento"`
}

// TODO: Define struct Direccion
type Direccion struct {
	Calle        string `json:"calle"`
	Numero       string `json:"numero"`
	Ciudad       string `json:"ciudad"`
	Estado       string `json:"estado"`
	CodigoPostal string `json:"codigo_postal"`
	Pais         string `json:"pais"`
}

// TODO: Define struct Empleado que embeba Persona
type Empleado struct {
	Persona                  // Embedding
	ID             int       `json:"id"`
	Departamento   string    `json:"departamento"`
	Puesto         string    `json:"puesto"`
	SalarioMensual float64   `json:"salario_mensual"`
	FechaIngreso   time.Time `json:"fecha_ingreso"`
	Direccion      Direccion `json:"direccion"`
	Activo         bool      `json:"activo"`
}

// TODO: Implementa métodos para structs
func (p Persona) NombreCompleto() string {
	return fmt.Sprintf("%s %s", p.Nombre, p.Apellido)
}

func (p Persona) CalcularEdad() int {
	return int(time.Since(p.FechaNac).Hours() / 24 / 365)
}

func (d Direccion) DireccionCompleta() string {
	return fmt.Sprintf("%s %s, %s, %s %s, %s",
		d.Calle, d.Numero, d.Ciudad, d.Estado, d.CodigoPostal, d.Pais)
}

func (e Empleado) SalarioAnual() float64 {
	return e.SalarioMensual * 12
}

func (e *Empleado) AumentarSalario(porcentaje float64) {
	e.SalarioMensual += e.SalarioMensual * porcentaje
}

func (e Empleado) TiempoEnEmpresa() time.Duration {
	return time.Since(e.FechaIngreso)
}

func ejercicio2() {
	fmt.Println("\n=== EJERCICIO 2: Embedding ===")

	// TODO: Crear empleado con embedding
	empleado := Empleado{
		Persona: Persona{
			Nombre:   "Ana",
			Apellido: "García",
			Edad:     28,
			Email:    "ana.garcia@empresa.com",
			Telefono: "+34 600-123-456",
			FechaNac: time.Date(1995, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		ID:             1001,
		Departamento:   "Desarrollo",
		Puesto:         "Senior Developer",
		SalarioMensual: 4500.0,
		FechaIngreso:   time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
		Direccion: Direccion{
			Calle:        "Av. Principal",
			Numero:       "123",
			Ciudad:       "Madrid",
			Estado:       "Madrid",
			CodigoPostal: "28001",
			Pais:         "España",
		},
		Activo: true,
	}

	// El embedding permite acceso directo a métodos de Persona
	fmt.Printf("Empleado: %s\n", empleado.NombreCompleto())
	fmt.Printf("Email: %s\n", empleado.Email) // Acceso directo por embedding
	fmt.Printf("Puesto: %s en %s\n", empleado.Puesto, empleado.Departamento)
	fmt.Printf("Salario anual: €%.2f\n", empleado.SalarioAnual())
	fmt.Printf("Dirección: %s\n", empleado.Direccion.DireccionCompleta())

	// Aumentar salario
	empleado.AumentarSalario(0.10) // 10% aumento
	fmt.Printf("Nuevo salario anual: €%.2f\n", empleado.SalarioAnual())
}

// ==============================================
// EJERCICIO 3: Struct Tags - Configuración JSON
// ==============================================

// TODO: Define struct ConfiguracionApp con tags JSON y validación
type ConfiguracionApp struct {
	AppName     string                 `json:"app_name" validate:"required"`
	Version     string                 `json:"version" validate:"required"`
	Port        int                    `json:"port" validate:"min=1,max=65535"`
	Debug       bool                   `json:"debug"`
	DatabaseURL string                 `json:"database_url" validate:"required"`
	Features    []string               `json:"features"`
	Settings    map[string]interface{} `json:"settings"`
	Environment string                 `json:"environment" validate:"oneof=dev test prod"`
	MaxUsers    int                    `json:"max_users" validate:"min=1"`
}

// TODO: Implementa validación personalizada
func (c ConfiguracionApp) Validar() error {
	if c.AppName == "" {
		return fmt.Errorf("app_name es requerido")
	}
	if c.Version == "" {
		return fmt.Errorf("version es requerido")
	}
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("port debe estar entre 1 y 65535")
	}
	if c.DatabaseURL == "" {
		return fmt.Errorf("database_url es requerido")
	}
	return nil
}

func (c ConfiguracionApp) String() string {
	return fmt.Sprintf("%s v%s (Puerto: %d, Debug: %t)",
		c.AppName, c.Version, c.Port, c.Debug)
}

func ejercicio3() {
	fmt.Println("\n=== EJERCICIO 3: Struct Tags ===")

	// TODO: Crear configuración con todos los campos
	config := ConfiguracionApp{
		AppName:     "EcommerceAPI",
		Version:     "2.1.0",
		Port:        8080,
		Debug:       true,
		DatabaseURL: "postgres://user:pass@localhost:5432/ecommerce",
		Features:    []string{"auth", "payment", "inventory", "analytics"},
		Settings: map[string]interface{}{
			"cache_ttl":       3600,
			"max_connections": 100,
			"ssl_enabled":     true,
		},
		Environment: "dev",
		MaxUsers:    10000,
	}

	fmt.Printf("Configuración: %s\n", config)

	// Validar configuración
	if err := config.Validar(); err != nil {
		fmt.Printf("❌ Error de validación: %v\n", err)
	} else {
		fmt.Println("✅ Configuración válida")
	}

	// TODO: Mostrar serialización JSON (simulada)
	fmt.Printf("Features: %v\n", config.Features)
	fmt.Printf("Settings: %v\n", config.Settings)
}

// ==============================================
// EJERCICIO 4: Múltiple Embedding - Vehículo
// ==============================================

// TODO: Define structs para embedding múltiple
type Motor struct {
	Tipo        string  `json:"tipo"`     // "gasolina", "diesel", "híbrido", "eléctrico"
	Potencia    int     `json:"potencia"` // HP
	Cilindros   int     `json:"cilindros"`
	Combustible string  `json:"combustible"`
	Eficiencia  float64 `json:"eficiencia"` // km/litro
}

type Transmision struct {
	Tipo     string `json:"tipo"` // "manual", "automática", "CVT"
	Marchas  int    `json:"marchas"`
	Traccion string `json:"traccion"` // "FWD", "RWD", "AWD"
}

type Carroceria struct {
	Tipo     string `json:"tipo"` // "sedán", "SUV", "hatchback"
	Puertas  int    `json:"puertas"`
	Asientos int    `json:"asientos"`
	Color    string `json:"color"`
}

// TODO: Define struct Vehiculo con múltiple embedding
type Vehiculo struct {
	Motor               // Embedding
	Transmision         // Embedding
	Carroceria          // Embedding
	Marca       string  `json:"marca"`
	Modelo      string  `json:"modelo"`
	Año         int     `json:"año"`
	Kilometraje int     `json:"kilometraje"`
	Precio      float64 `json:"precio"`
	NumeroSerie string  `json:"numero_serie"`
}

// TODO: Implementa métodos para structs embebidos
func (m Motor) Descripcion() string {
	return fmt.Sprintf("Motor %s de %d HP (%d cil.) - %.1f km/l",
		m.Tipo, m.Potencia, m.Cilindros, m.Eficiencia)
}

func (t Transmision) Descripcion() string {
	return fmt.Sprintf("Transmisión %s de %d marchas (%s)",
		t.Tipo, t.Marchas, t.Traccion)
}

func (c Carroceria) Descripcion() string {
	return fmt.Sprintf("%s %s de %d puertas y %d asientos",
		c.Tipo, c.Color, c.Puertas, c.Asientos)
}

func (v Vehiculo) DescripcionCompleta() string {
	return fmt.Sprintf("%s %s %d\n  %s\n  %s\n  %s\n  Km: %d - $%.2f",
		v.Marca, v.Modelo, v.Año,
		v.Motor.Descripcion(),
		v.Transmision.Descripcion(),
		v.Carroceria.Descripcion(),
		v.Kilometraje, v.Precio)
}

func ejercicio4() {
	fmt.Println("\n=== EJERCICIO 4: Múltiple Embedding ===")

	// TODO: Crear vehículo con múltiple embedding
	vehiculo := Vehiculo{
		Motor: Motor{
			Tipo:        "híbrido",
			Potencia:    200,
			Cilindros:   4,
			Combustible: "gasolina",
			Eficiencia:  22.5,
		},
		Transmision: Transmision{
			Tipo:     "automática",
			Marchas:  8,
			Traccion: "AWD",
		},
		Carroceria: Carroceria{
			Tipo:     "SUV",
			Puertas:  5,
			Asientos: 7,
			Color:    "azul metalizado",
		},
		Marca:       "Toyota",
		Modelo:      "Highlander Hybrid",
		Año:         2024,
		Kilometraje: 15000,
		Precio:      45000.00,
		NumeroSerie: "TH2024ABC123456",
	}

	fmt.Printf("Vehículo:\n%s\n", vehiculo.DescripcionCompleta())

	// Acceso directo a campos embebidos
	fmt.Printf("\nDetalles técnicos:\n")
	fmt.Printf("- Potencia: %d HP\n", vehiculo.Potencia)
	fmt.Printf("- Transmisión: %s\n", vehiculo.Motor.Tipo) // Especificamos el embedding
	fmt.Printf("- Asientos: %d\n", vehiculo.Asientos)
}

// ==============================================
// EJERCICIO 5: Structs Anónimos - Inventario
// ==============================================

func ejercicio5() {
	fmt.Println("\n=== EJERCICIO 5: Structs Anónimos ===")

	// TODO: Crear slice de structs anónimos para productos
	productos := []struct {
		ID        int     `json:"id"`
		Nombre    string  `json:"nombre"`
		Categoria string  `json:"categoria"`
		Precio    float64 `json:"precio"`
		Stock     int     `json:"stock"`
		Activo    bool    `json:"activo"`
	}{
		{1, "Laptop Gaming", "Electrónicos", 1299.99, 5, true},
		{2, "Mouse Inalámbrico", "Electrónicos", 29.99, 50, true},
		{3, "Silla Ergonómica", "Muebles", 299.99, 10, true},
		{4, "Escritorio de Pie", "Muebles", 399.99, 8, true},
		{5, "Monitor 4K", "Electrónicos", 449.99, 12, true},
	}

	fmt.Println("Productos:")
	for _, producto := range productos {
		estado := "✅"
		if !producto.Activo || producto.Stock == 0 {
			estado = "❌"
		}
		fmt.Printf("%s %s (%s) - $%.2f [Stock: %d]\n",
			estado, producto.Nombre, producto.Categoria, producto.Precio, producto.Stock)
	}

	// TODO: Crea un map de structs anónimos para estadísticas por categoría
	estadisticas := map[string]struct {
		TotalProductos  int
		PrecioPromedio  float64
		StockTotal      int
		ValorInventario float64
	}{}

	// TODO: Calcula estadísticas por categoría
	for _, producto := range productos {
		stats := estadisticas[producto.Categoria]
		stats.TotalProductos++
		stats.StockTotal += producto.Stock
		stats.ValorInventario += producto.Precio * float64(producto.Stock)

		// Recalcular promedio
		total := float64(0)
		count := 0
		for _, p := range productos {
			if p.Categoria == producto.Categoria {
				total += p.Precio
				count++
			}
		}
		stats.PrecioPromedio = total / float64(count)

		estadisticas[producto.Categoria] = stats
	}

	fmt.Println("\nEstadísticas por categoría:")
	for categoria, stats := range estadisticas {
		fmt.Printf("📊 %s:\n", categoria)
		fmt.Printf("   Productos: %d\n", stats.TotalProductos)
		fmt.Printf("   Precio promedio: $%.2f\n", stats.PrecioPromedio)
		fmt.Printf("   Stock total: %d unidades\n", stats.StockTotal)
		fmt.Printf("   Valor inventario: $%.2f\n\n", stats.ValorInventario)
	}
}

// ==============================================
// EJERCICIO 6: Factory Pattern - Conexiones DB
// ==============================================

// TODO: Define structs para Factory Pattern
type ConexionDB struct {
	Tipo        string            `json:"tipo"`
	Host        string            `json:"host"`
	Puerto      int               `json:"puerto"`
	Database    string            `json:"database"`
	Usuario     string            `json:"usuario"`
	Password    string            `json:"-"` // No serializar
	Opciones    map[string]string `json:"opciones"`
	Conectado   bool              `json:"conectado"`
	UltimaConex time.Time         `json:"ultima_conexion"`
}

type DBFactory struct {
	Conexiones map[string]*ConexionDB `json:"conexiones"`
	Default    string                 `json:"default"`
}

// TODO: Implementa constructor y métodos de factory
func NewDBFactory() *DBFactory {
	return &DBFactory{
		Conexiones: make(map[string]*ConexionDB),
	}
}

func (f *DBFactory) CrearMySQL(nombre, host string, puerto int, db, user, pass string) *ConexionDB {
	conn := &ConexionDB{
		Tipo:     "mysql",
		Host:     host,
		Puerto:   puerto,
		Database: db,
		Usuario:  user,
		Password: pass,
		Opciones: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "true",
		},
	}
	f.Conexiones[nombre] = conn
	if f.Default == "" {
		f.Default = nombre
	}
	return conn
}

func (f *DBFactory) CrearPostgreSQL(nombre, host string, puerto int, db, user, pass string) *ConexionDB {
	conn := &ConexionDB{
		Tipo:     "postgresql",
		Host:     host,
		Puerto:   puerto,
		Database: db,
		Usuario:  user,
		Password: pass,
		Opciones: map[string]string{
			"sslmode":  "prefer",
			"timezone": "UTC",
		},
	}
	f.Conexiones[nombre] = conn
	return conn
}

func (f *DBFactory) CrearRedis(nombre, host string, puerto int, db string) *ConexionDB {
	conn := &ConexionDB{
		Tipo:     "redis",
		Host:     host,
		Puerto:   puerto,
		Database: db,
		Opciones: map[string]string{
			"maxRetries": "3",
			"poolSize":   "10",
		},
	}
	f.Conexiones[nombre] = conn
	return conn
}

// TODO: Métodos para ConexionDB
func (c *ConexionDB) Conectar() error {
	fmt.Printf("🔌 Conectando a %s://%s:%d/%s...\n",
		c.Tipo, c.Host, c.Puerto, c.Database)
	c.Conectado = true
	c.UltimaConex = time.Now()
	fmt.Printf("✅ Conexión %s establecida\n", c.Tipo)
	return nil
}

func (c *ConexionDB) Desconectar() {
	c.Conectado = false
	fmt.Printf("🔌 Desconectado de %s\n", c.Tipo)
}

func (c *ConexionDB) String() string {
	estado := "❌ desconectado"
	if c.Conectado {
		estado = "✅ conectado"
	}
	return fmt.Sprintf("%s://%s:%d/%s [%s]",
		c.Tipo, c.Host, c.Puerto, c.Database, estado)
}

func ejercicio6() {
	fmt.Println("\n=== EJERCICIO 6: Factory Pattern ===")

	// TODO: Usar factory para crear diferentes conexiones
	factory := NewDBFactory()

	// Crear diferentes tipos de conexiones
	mysql := factory.CrearMySQL("main", "localhost", 3306, "ecommerce", "root", "secret")
	postgres := factory.CrearPostgreSQL("analytics", "db.empresa.com", 5432, "analytics", "analyst", "pass123")
	redis := factory.CrearRedis("cache", "cache.empresa.com", 6379, "0")

	// Usar las conexiones creadas
	_ = mysql
	_ = postgres
	_ = redis

	fmt.Println("Conexiones creadas:")
	for nombre, conn := range factory.Conexiones {
		fmt.Printf("  %s: %s\n", nombre, conn)
	}

	fmt.Printf("\nConexión por defecto: %s\n", factory.Default)

	// Conectar todas
	fmt.Println("\nEstableciendo conexiones:")
	for _, conn := range factory.Conexiones {
		conn.Conectar()
	}

	fmt.Println("\nEstado final:")
	for nombre, conn := range factory.Conexiones {
		fmt.Printf("  %s: %s\n", nombre, conn)
	}
}

// ==============================================
// EJERCICIO 7: Validación Avanzada - Usuarios
// ==============================================

// TODO: Define interface para validación
type Validable interface {
	Validar() error
}

// TODO: Define un struct Usuario con validaciones complejas
type Usuario struct {
	ID              int       `json:"id"`
	Username        string    `json:"username" validate:"required,min=3,max=20"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required,min=8"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Telefono        string    `json:"telefono"`
	Direccion       string    `json:"direccion"`
	Activo          bool      `json:"activo"`
	Rol             string    `json:"rol" validate:"oneof=admin user guest"`
	UltimoLogin     time.Time `json:"ultimo_login"`
}

// TODO: Implementa validaciones detalladas
func (u Usuario) Validar() error {
	// Validar username
	if len(u.Username) < 3 || len(u.Username) > 20 {
		return fmt.Errorf("username debe tener entre 3 y 20 caracteres")
	}

	// Validar email
	if !u.ValidarEmail() {
		return fmt.Errorf("email no es válido")
	}

	// Validar password
	if !u.ValidarPassword() {
		return fmt.Errorf("password no cumple los requisitos de seguridad")
	}

	// Validar edad
	if !u.ValidarEdad() {
		return fmt.Errorf("usuario debe ser mayor de 13 años")
	}

	return nil
}

func (u Usuario) ValidarEmail() bool {
	// Validación básica de email
	email := u.Email
	return len(email) > 5 &&
		len(email) < 100 &&
		fmt.Sprintf("%s", email)[0] != '@' &&
		fmt.Sprintf("%s", email)[len(email)-1] != '@'
}

func (u Usuario) ValidarPassword() bool {
	pass := u.Password
	if len(pass) < 8 {
		return false
	}

	tieneMinuscula := false
	tieneMayuscula := false
	tieneNumero := false

	for _, char := range pass {
		if char >= 'a' && char <= 'z' {
			tieneMinuscula = true
		}
		if char >= 'A' && char <= 'Z' {
			tieneMayuscula = true
		}
		if char >= '0' && char <= '9' {
			tieneNumero = true
		}
	}

	return tieneMinuscula && tieneMayuscula && tieneNumero
}

func (u Usuario) ValidarEdad() bool {
	edad := time.Since(u.FechaNacimiento).Hours() / 24 / 365
	return edad >= 13
}

func (u Usuario) EdadActual() int {
	return int(time.Since(u.FechaNacimiento).Hours() / 24 / 365)
}

func ejercicio7() {
	fmt.Println("\n=== EJERCICIO 7: Validación Avanzada ===")

	usuarios := []Usuario{
		{
			ID:              1,
			Username:        "ana_developer",
			Email:           "ana@empresa.com",
			Password:        "SecurePass123",
			FechaNacimiento: time.Date(1995, 5, 15, 0, 0, 0, 0, time.UTC),
			Telefono:        "+34-600-123-456",
			Direccion:       "Madrid, España",
			Activo:          true,
			Rol:             "admin",
			UltimoLogin:     time.Now().AddDate(0, 0, -1),
		},
		{
			ID:              2,
			Username:        "xy",                                        // Muy corto
			Email:           "invalid-email",                             // Sin @ ni dominio
			Password:        "123",                                       // Muy simple
			FechaNacimiento: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), // Menor de edad
			Activo:          false,
			Rol:             "user",
		},
		{
			ID:              3,
			Username:        "carlos_manager",
			Email:           "carlos@empresa.com",
			Password:        "MyComplexPass456",
			FechaNacimiento: time.Date(1988, 10, 22, 0, 0, 0, 0, time.UTC),
			Telefono:        "+34-610-987-654",
			Direccion:       "Barcelona, España",
			Activo:          true,
			Rol:             "user",
			UltimoLogin:     time.Now().AddDate(0, 0, -3),
		},
	}

	fmt.Println("Validando usuarios:")
	for i, usuario := range usuarios {
		fmt.Printf("\nUsuario %d: %s\n", i+1, usuario.Username)
		fmt.Printf("  Email: %s\n", usuario.Email)
		fmt.Printf("  Edad: %d años\n", usuario.EdadActual())
		fmt.Printf("  Rol: %s\n", usuario.Rol)

		if err := usuario.Validar(); err != nil {
			fmt.Printf("  ❌ Errores: %v\n", err)
		} else {
			fmt.Printf("  ✅ Usuario válido\n")
		}
	}
}

// ==============================================
// EJERCICIO 8: Builder Pattern - Servidor Web
// ==============================================

// TODO: Define structs para Builder Pattern
type ServidorWeb struct {
	Host          string        `json:"host"`
	Puerto        int           `json:"puerto"`
	HTTPS         bool          `json:"https"`
	Timeout       time.Duration `json:"timeout"`
	MaxConexiones int           `json:"max_conexiones"`
	CORS          bool          `json:"cors"`
	LogLevel      string        `json:"log_level"`
	Middlewares   []string      `json:"middlewares"`
	TLSCert       string        `json:"tls_cert,omitempty"`
	TLSKey        string        `json:"tls_key,omitempty"`
	StaticPath    string        `json:"static_path"`
	DatabaseURL   string        `json:"database_url"`
}

type ServidorBuilder struct {
	servidor *ServidorWeb
}

// TODO: Implementa constructor del builder
func NewServidorBuilder() *ServidorBuilder {
	return &ServidorBuilder{
		servidor: &ServidorWeb{
			Host:          "localhost",
			Puerto:        8080,
			HTTPS:         false,
			Timeout:       30 * time.Second,
			MaxConexiones: 100,
			CORS:          false,
			LogLevel:      "info",
			Middlewares:   []string{},
		},
	}
}

// TODO: Implementa métodos del builder (fluent interface)
func (b *ServidorBuilder) Host(host string) *ServidorBuilder {
	b.servidor.Host = host
	return b
}

func (b *ServidorBuilder) Puerto(puerto int) *ServidorBuilder {
	b.servidor.Puerto = puerto
	return b
}

func (b *ServidorBuilder) HTTPS(cert, key string) *ServidorBuilder {
	b.servidor.HTTPS = true
	b.servidor.TLSCert = cert
	b.servidor.TLSKey = key
	return b
}

func (b *ServidorBuilder) Timeout(timeout time.Duration) *ServidorBuilder {
	b.servidor.Timeout = timeout
	return b
}

func (b *ServidorBuilder) MaxConexiones(max int) *ServidorBuilder {
	b.servidor.MaxConexiones = max
	return b
}

func (b *ServidorBuilder) CORS(enable bool) *ServidorBuilder {
	b.servidor.CORS = enable
	return b
}

func (b *ServidorBuilder) LogLevel(level string) *ServidorBuilder {
	b.servidor.LogLevel = level
	return b
}

func (b *ServidorBuilder) AgregarMiddleware(middleware string) *ServidorBuilder {
	b.servidor.Middlewares = append(b.servidor.Middlewares, middleware)
	return b
}

func (b *ServidorBuilder) StaticFiles(path string) *ServidorBuilder {
	b.servidor.StaticPath = path
	return b
}

func (b *ServidorBuilder) Database(url string) *ServidorBuilder {
	b.servidor.DatabaseURL = url
	return b
}

func (b *ServidorBuilder) Build() *ServidorWeb {
	return b.servidor
}

// TODO: Métodos para ServidorWeb
func (s *ServidorWeb) URL() string {
	protocolo := "http"
	if s.HTTPS {
		protocolo = "https"
	}
	return fmt.Sprintf("%s://%s:%d", protocolo, s.Host, s.Puerto)
}

func (s *ServidorWeb) Iniciar() {
	protocolo := "HTTP"
	if s.HTTPS {
		protocolo = "HTTPS"
	}

	fmt.Printf("🚀 Iniciando servidor %s\n", protocolo)
	fmt.Printf("   URL: %s\n", s.URL())
	fmt.Printf("   Timeout: %v\n", s.Timeout)
	fmt.Printf("   Max conexiones: %d\n", s.MaxConexiones)
	fmt.Printf("   CORS: %t\n", s.CORS)
	fmt.Printf("   Log level: %s\n", s.LogLevel)

	if len(s.Middlewares) > 0 {
		fmt.Printf("   Middlewares: %v\n", s.Middlewares)
	}

	if s.StaticPath != "" {
		fmt.Printf("   Archivos estáticos: %s\n", s.StaticPath)
	}

	if s.DatabaseURL != "" {
		fmt.Printf("   Base de datos: %s\n", s.DatabaseURL[:20]+"...")
	}

	fmt.Println("   ✅ Servidor listo para recibir conexiones")
}

func (s *ServidorWeb) String() string {
	return fmt.Sprintf("%s (Timeout: %v, Max: %d)",
		s.URL(), s.Timeout, s.MaxConexiones)
}

func ejercicio8() {
	fmt.Println("\n=== EJERCICIO 8: Builder Pattern ===")

	// TODO: Crear servidores con builder pattern
	fmt.Println("📋 Configurando servidores con Builder Pattern...")

	// Servidor de desarrollo
	servidorDev := NewServidorBuilder().
		Host("localhost").
		Puerto(3000).
		Timeout(10 * time.Second).
		CORS(true).
		LogLevel("debug").
		AgregarMiddleware("cors").
		AgregarMiddleware("logger").
		AgregarMiddleware("auth").
		StaticFiles("./public").
		Database("postgres://dev:dev@localhost:5432/app_dev").
		Build()

	// Servidor de producción
	servidorProd := NewServidorBuilder().
		Host("0.0.0.0").
		Puerto(443).
		HTTPS("/etc/ssl/cert.pem", "/etc/ssl/key.pem").
		Timeout(30 * time.Second).
		MaxConexiones(500).
		CORS(false).
		LogLevel("warn").
		AgregarMiddleware("security").
		AgregarMiddleware("rateLimit").
		AgregarMiddleware("auth").
		StaticFiles("/var/www/static").
		Database("postgres://prod_user:secret@db.empresa.com:5432/app_prod").
		Build()

	fmt.Println("\n🔧 Servidor de Desarrollo:")
	fmt.Printf("   %s\n", servidorDev)
	servidorDev.Iniciar()

	fmt.Println("\n🏭 Servidor de Producción:")
	fmt.Printf("   %s\n", servidorProd)
	servidorProd.Iniciar()
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
	fmt.Println("📚 Verifica tus soluciones con 'go run soluciones.go'")
	fmt.Println("🎯 Continúa con el proyecto: 'go run proyecto_ecommerce.go'")
}
