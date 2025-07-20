// Archivo: soluciones.go
// Lección 11: Structs - Soluciones Completas
// Este archivo contiene las soluciones para todos los ejercicios de structs

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ==============================================
// SOLUCIÓN EJERCICIO 1: Struct Básico
// ==============================================

type LibroSol struct {
	ID               int
	Titulo           string
	Autor            string
	Paginas          int
	Precio           float64
	Disponible       bool
	FechaPublicacion time.Time
}

var contadorLibrosSol int

func NewLibroSol(titulo, autor string, paginas int, precio float64) LibroSol {
	contadorLibrosSol++
	return LibroSol{
		ID:               contadorLibrosSol,
		Titulo:           titulo,
		Autor:            autor,
		Paginas:          paginas,
		Precio:           precio,
		Disponible:       true,
		FechaPublicacion: time.Now(),
	}
}

func (l LibroSol) String() string {
	return fmt.Sprintf("📖 %s por %s - %d páginas - $%.2f (ID: %d)",
		l.Titulo, l.Autor, l.Paginas, l.Precio, l.ID)
}

func (l *LibroSol) Descuento(porcentaje float64) {
	l.Precio = l.Precio * (1 - porcentaje)
}

func (l LibroSol) EsCaro() bool {
	return l.Precio > 25.00
}

// ==============================================
// SOLUCIÓN EJERCICIO 2: Embedding
// ==============================================

type PersonaSol struct {
	Nombre   string
	Apellido string
	Edad     int
	Email    string
}

type DireccionSol struct {
	Calle   string
	Ciudad  string
	Estado  string
	CodigoP string
	Pais    string
}

type EmpleadoSol struct {
	PersonaSol                  // Embedding anónimo
	Direccion      DireccionSol // Embedding nombrado
	Departamento   string
	Puesto         string
	SalarioMensual float64
	FechaIngreso   time.Time
}

func (p PersonaSol) NombreCompleto() string {
	return fmt.Sprintf("%s %s", p.Nombre, p.Apellido)
}

func (d DireccionSol) DireccionCompleta() string {
	return fmt.Sprintf("%s, %s, %s %s, %s",
		d.Calle, d.Ciudad, d.Estado, d.CodigoP, d.Pais)
}

func (e EmpleadoSol) CalcularSalarioAnual() float64 {
	return e.SalarioMensual * 12
}

func (e *EmpleadoSol) AumentarSalario(porcentaje float64) {
	e.SalarioMensual = e.SalarioMensual * (1 + porcentaje)
}

// ==============================================
// SOLUCIÓN EJERCICIO 3: Struct Tags
// ==============================================

type ConfiguracionAppSol struct {
	AppName     string                 `json:"app_name" validate:"required"`
	Version     string                 `json:"version" validate:"required"`
	Port        int                    `json:"port" validate:"required,min=1,max=65535"`
	Debug       bool                   `json:"debug"`
	DatabaseURL string                 `json:"database_url" validate:"required"`
	Features    []string               `json:"features"`
	Settings    map[string]interface{} `json:"settings"`
}

func (c ConfiguracionAppSol) Validar() error {
	var errores []string

	if strings.TrimSpace(c.AppName) == "" {
		errores = append(errores, "AppName es requerido")
	}

	if strings.TrimSpace(c.Version) == "" {
		errores = append(errores, "Version es requerida")
	}

	if c.Port < 1 || c.Port > 65535 {
		errores = append(errores, "Port debe estar entre 1 y 65535")
	}

	if strings.TrimSpace(c.DatabaseURL) == "" {
		errores = append(errores, "DatabaseURL es requerida")
	}

	if len(errores) > 0 {
		return errors.New(strings.Join(errores, "; "))
	}

	return nil
}

func (c ConfiguracionAppSol) String() string {
	dbURL := c.DatabaseURL
	if len(dbURL) > 20 {
		dbURL = dbURL[:20] + "..."
	}

	return fmt.Sprintf("App: %s v%s, Puerto: %d, Debug: %v, DB: %s",
		c.AppName, c.Version, c.Port, c.Debug, dbURL)
}

// ==============================================
// SOLUCIÓN EJERCICIO 4: Múltiple Embedding
// ==============================================

type Motor struct {
	Potencia    int    // HP
	Combustible string // "gasolina", "diesel", "electrico"
	Cilindros   int
}

type Ruedas struct {
	Cantidad int
	Tamaño   string // "15\"", "16\"", etc.
	Tipo     string // "aleacion", "acero"
}

type Identificacion struct {
	Marca       string
	Modelo      string
	Año         int
	NumeroSerie string
}

type Vehiculo struct {
	Identificacion // Embedding anónimo
	Motor          // Embedding anónimo
	Ruedas         // Embedding anónimo
	Color          string
	Kilometraje    int
}

func (m Motor) Descripcion() string {
	return fmt.Sprintf("Motor %s de %d HP con %d cilindros",
		m.Combustible, m.Potencia, m.Cilindros)
}

func (r Ruedas) Descripcion() string {
	return fmt.Sprintf("%d ruedas %s de %s",
		r.Cantidad, r.Tipo, r.Tamaño)
}

func (i Identificacion) Descripcion() string {
	return fmt.Sprintf("%s %s %d (S/N: %s)",
		i.Marca, i.Modelo, i.Año, i.NumeroSerie)
}

func (v Vehiculo) ResumenCompleto() string {
	return fmt.Sprintf("%s\n  %s\n  %s\n  Color: %s, Km: %d",
		v.Identificacion.Descripcion(),
		v.Motor.Descripcion(),
		v.Ruedas.Descripcion(),
		v.Color,
		v.Kilometraje)
}

// ==============================================
// SOLUCIÓN EJERCICIO 6: Factory Pattern
// ==============================================

type ConexionDB struct {
	Tipo      string
	Host      string
	Puerto    int
	Database  string
	Usuario   string
	Password  string
	Pool      int
	Timeout   time.Duration
	conectado bool
}

type DBFactory struct {
	configuraciones map[string]ConexionDB
}

func NewDBFactory() *DBFactory {
	return &DBFactory{
		configuraciones: map[string]ConexionDB{
			"mysql": {
				Tipo:    "mysql",
				Host:    "localhost",
				Puerto:  3306,
				Pool:    10,
				Timeout: 30 * time.Second,
			},
			"postgresql": {
				Tipo:    "postgresql",
				Host:    "localhost",
				Puerto:  5432,
				Pool:    20,
				Timeout: 45 * time.Second,
			},
			"redis": {
				Tipo:    "redis",
				Host:    "localhost",
				Puerto:  6379,
				Pool:    50,
				Timeout: 5 * time.Second,
			},
		},
	}
}

func (f *DBFactory) CrearMySQL(database, usuario, password string) *ConexionDB {
	config := f.configuraciones["mysql"]
	config.Database = database
	config.Usuario = usuario
	config.Password = password
	return &config
}

func (f *DBFactory) CrearPostgreSQL(database, usuario, password string) *ConexionDB {
	config := f.configuraciones["postgresql"]
	config.Database = database
	config.Usuario = usuario
	config.Password = password
	return &config
}

func (f *DBFactory) CrearRedis(database int) *ConexionDB {
	config := f.configuraciones["redis"]
	config.Database = fmt.Sprintf("db%d", database)
	return &config
}

func (c *ConexionDB) Conectar() error {
	fmt.Printf("🔌 Conectando a %s://%s:%d/%s\n",
		c.Tipo, c.Host, c.Puerto, c.Database)
	c.conectado = true
	return nil
}

func (c *ConexionDB) Desconectar() error {
	fmt.Printf("🔌 Desconectando de %s\n", c.Tipo)
	c.conectado = false
	return nil
}

func (c *ConexionDB) String() string {
	status := "desconectado"
	if c.conectado {
		status = "conectado"
	}
	return fmt.Sprintf("%s://%s:%d/%s [%s]",
		c.Tipo, c.Host, c.Puerto, c.Database, status)
}

// ==============================================
// SOLUCIÓN EJERCICIO 7: Validación Avanzada
// ==============================================

type Validable interface {
	Validar() error
}

type Usuario struct {
	ID              int       `json:"id" validate:"required"`
	Username        string    `json:"username" validate:"required,min=3,max=20"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"-" validate:"required,password"`
	FechaNacimiento time.Time `json:"birth_date" validate:"required"`
	Telefono        string    `json:"phone" validate:"phone"`
	Direccion       string    `json:"address"`
}

func (u Usuario) Validar() error {
	var errores []string

	if u.ID <= 0 {
		errores = append(errores, "ID debe ser mayor que 0")
	}

	if len(u.Username) < 3 || len(u.Username) > 20 {
		errores = append(errores, "Username debe tener entre 3 y 20 caracteres")
	}

	if !u.ValidarEmail() {
		errores = append(errores, "Email no tiene formato válido")
	}

	if !u.ValidarPassword() {
		errores = append(errores, "Password debe tener al menos 8 caracteres, una mayúscula, una minúscula y un número")
	}

	if !u.ValidarEdad() {
		errores = append(errores, "Debe ser mayor de 18 años")
	}

	if len(errores) > 0 {
		return errors.New(strings.Join(errores, "; "))
	}

	return nil
}

func (u Usuario) ValidarEmail() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(u.Email)
}

func (u Usuario) ValidarPassword() bool {
	if len(u.Password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(u.Password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(u.Password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(u.Password)

	return hasUpper && hasLower && hasDigit
}

func (u Usuario) ValidarEdad() bool {
	años := time.Since(u.FechaNacimiento).Hours() / 24 / 365.25
	return años >= 18
}

// ==============================================
// SOLUCIÓN EJERCICIO 8: Builder Pattern
// ==============================================

type ServidorWeb struct {
	Host          string
	Puerto        int
	SSL           bool
	Certificado   string
	ClavePrivada  string
	Timeout       time.Duration
	MaxConexiones int
	CORS          bool
	LogLevel      string
	Middleware    []string
	rutasCount    int
}

type ServidorBuilder struct {
	servidor *ServidorWeb
}

func NewServidorBuilder() *ServidorBuilder {
	return &ServidorBuilder{
		servidor: &ServidorWeb{
			Host:          "localhost",
			Puerto:        8080,
			SSL:           false,
			Timeout:       30 * time.Second,
			MaxConexiones: 100,
			CORS:          false,
			LogLevel:      "info",
			Middleware:    []string{},
			rutasCount:    0,
		},
	}
}

func (b *ServidorBuilder) Host(host string) *ServidorBuilder {
	b.servidor.Host = host
	return b
}

func (b *ServidorBuilder) Puerto(puerto int) *ServidorBuilder {
	b.servidor.Puerto = puerto
	return b
}

func (b *ServidorBuilder) ConSSL(cert, key string) *ServidorBuilder {
	b.servidor.SSL = true
	b.servidor.Certificado = cert
	b.servidor.ClavePrivada = key
	return b
}

func (b *ServidorBuilder) Timeout(timeout time.Duration) *ServidorBuilder {
	b.servidor.Timeout = timeout
	return b
}

func (b *ServidorBuilder) ConCORS() *ServidorBuilder {
	b.servidor.CORS = true
	return b
}

func (b *ServidorBuilder) Build() *ServidorWeb {
	return b.servidor
}

func (s *ServidorWeb) Iniciar() error {
	protocol := "HTTP"
	if s.SSL {
		protocol = "HTTPS"
	}

	fmt.Printf("🚀 Iniciando servidor %s en %s:%d\n", protocol, s.Host, s.Puerto)
	fmt.Printf("⚙️ Timeout: %v, Max Conexiones: %d\n", s.Timeout, s.MaxConexiones)
	fmt.Printf("🔧 CORS: %v, Log Level: %s\n", s.CORS, s.LogLevel)

	return nil
}

func (s *ServidorWeb) String() string {
	protocol := "http"
	if s.SSL {
		protocol = "https"
	}

	return fmt.Sprintf("%s://%s:%d (Timeout: %v, CORS: %v)",
		protocol, s.Host, s.Puerto, s.Timeout, s.CORS)
}

// ==============================================
// FUNCIONES DE DEMOSTRACIÓN
// ==============================================

func solucion1() {
	fmt.Println("=== SOLUCIÓN 1: Struct Básico ===")

	libro1 := NewLibroSol("El Quijote", "Cervantes", 863, 29.99)
	libro2 := NewLibroSol("Go Programming", "Kernighan", 380, 45.50)

	fmt.Println("Libros creados:")
	fmt.Println(libro1)
	fmt.Println(libro2)

	libro1.Descuento(0.10)
	libro2.Descuento(0.15)

	fmt.Println("\nDespués de descuentos:")
	fmt.Println(libro1)
	fmt.Println(libro2)

	fmt.Printf("\n¿El Quijote es caro? %v\n", libro1.EsCaro())
	fmt.Printf("¿Go Programming es caro? %v\n", libro2.EsCaro())
}

func solucion2() {
	fmt.Println("\n=== SOLUCIÓN 2: Embedding ===")

	empleado := EmpleadoSol{
		PersonaSol: PersonaSol{
			Nombre:   "Ana",
			Apellido: "García",
			Edad:     30,
			Email:    "ana@empresa.com",
		},
		Direccion: DireccionSol{
			Calle:   "Av. Principal 123",
			Ciudad:  "Madrid",
			Estado:  "Madrid",
			CodigoP: "28001",
			Pais:    "España",
		},
		Departamento:   "Desarrollo",
		Puesto:         "Senior Developer",
		SalarioMensual: 4500.00,
		FechaIngreso:   time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
	}

	fmt.Printf("Empleado: %s\n", empleado.NombreCompleto())
	fmt.Printf("Dirección: %s\n", empleado.Direccion.DireccionCompleta())
	fmt.Printf("Salario anual: $%.2f\n", empleado.CalcularSalarioAnual())

	empleado.AumentarSalario(0.10)
	fmt.Printf("Nuevo salario anual: $%.2f\n", empleado.CalcularSalarioAnual())
}

func solucion3() {
	fmt.Println("\n=== SOLUCIÓN 3: Struct Tags ===")

	config := ConfiguracionAppSol{
		AppName:     "MiApp",
		Version:     "1.2.3",
		Port:        8080,
		Debug:       true,
		DatabaseURL: "postgres://user:pass@localhost:5432/mydb",
		Features:    []string{"auth", "metrics", "logging"},
		Settings: map[string]interface{}{
			"max_connections": 100,
			"timeout":         30,
			"ssl_enabled":     true,
		},
	}

	fmt.Println("Configuración:")
	fmt.Println(config)

	if err := config.Validar(); err != nil {
		fmt.Printf("Error de validación: %v\n", err)
	} else {
		fmt.Println("✅ Configuración válida")
	}

	jsonData, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println("JSON:")
	fmt.Println(string(jsonData))
}

func solucion4() {
	fmt.Println("\n=== SOLUCIÓN 4: Múltiple Embedding ===")

	vehiculo := Vehiculo{
		Identificacion: Identificacion{
			Marca:       "Toyota",
			Modelo:      "Camry",
			Año:         2023,
			NumeroSerie: "ABC123456789",
		},
		Motor: Motor{
			Potencia:    200,
			Combustible: "gasolina",
			Cilindros:   4,
		},
		Ruedas: Ruedas{
			Cantidad: 4,
			Tamaño:   "17\"",
			Tipo:     "aleacion",
		},
		Color:       "Azul Metalizado",
		Kilometraje: 15000,
	}

	fmt.Println("Vehículo:")
	fmt.Println(vehiculo.ResumenCompleto())
}

func solucion5() {
	fmt.Println("\n=== SOLUCIÓN 5: Structs Anónimos ===")

	productos := []struct {
		ID        int
		Nombre    string
		Categoria string
		Precio    float64
		Stock     int
	}{
		{1, "Laptop Gaming", "Electrónicos", 1299.99, 5},
		{2, "Mouse Inalámbrico", "Electrónicos", 29.99, 50},
		{3, "Silla Ergonómica", "Muebles", 299.99, 10},
		{4, "Escritorio", "Muebles", 199.99, 8},
		{5, "Monitor 4K", "Electrónicos", 399.99, 12},
	}

	fmt.Println("Productos:")
	for _, producto := range productos {
		fmt.Printf("- %s (%s): $%.2f (Stock: %d)\n",
			producto.Nombre, producto.Categoria, producto.Precio, producto.Stock)
	}

	estadisticas := map[string]struct {
		TotalProductos int
		PrecioPromedio float64
		StockTotal     int
	}{}

	for _, producto := range productos {
		stats := estadisticas[producto.Categoria]
		stats.TotalProductos++
		stats.PrecioPromedio = (stats.PrecioPromedio*float64(stats.TotalProductos-1) + producto.Precio) / float64(stats.TotalProductos)
		stats.StockTotal += producto.Stock
		estadisticas[producto.Categoria] = stats
	}

	fmt.Println("\nEstadísticas por categoría:")
	for categoria, stats := range estadisticas {
		fmt.Printf("- %s: %d productos, precio promedio $%.2f, stock total %d\n",
			categoria, stats.TotalProductos, stats.PrecioPromedio, stats.StockTotal)
	}
}

func solucion6() {
	fmt.Println("\n=== SOLUCIÓN 6: Factory Pattern ===")

	factory := NewDBFactory()

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

func solucion7() {
	fmt.Println("\n=== SOLUCIÓN 7: Validación Avanzada ===")

	usuarios := []Usuario{
		{
			ID:              1,
			Username:        "ana_garcia",
			Email:           "ana@ejemplo.com",
			Password:        "MiPassword123",
			FechaNacimiento: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			Telefono:        "+34612345678",
			Direccion:       "Calle Principal 123",
		},
		{
			ID:              2,
			Username:        "xy",                                        // ERROR: muy corto
			Email:           "email-invalido",                            // ERROR: formato inválido
			Password:        "123",                                       // ERROR: muy simple
			FechaNacimiento: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC), // ERROR: menor de edad
		},
		{
			ID:              3,
			Username:        "carlos_lopez",
			Email:           "carlos@test.com",
			Password:        "SecurePass123",
			FechaNacimiento: time.Date(1985, 8, 22, 0, 0, 0, 0, time.UTC),
			Telefono:        "+34687654321",
		},
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

func solucion8() {
	fmt.Println("\n=== SOLUCIÓN 8: Builder Pattern ===")

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

func main() {
	fmt.Println("🏗️ SOLUCIONES DE EJERCICIOS - STRUCTS")
	fmt.Println("=====================================")

	solucion1()
	solucion2()
	solucion3()
	solucion4()
	solucion5()
	solucion6()
	solucion7()
	solucion8()

	fmt.Println("\n=== FIN DE LAS SOLUCIONES ===")
}
