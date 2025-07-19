package main

import (
	"fmt"
	"time"
)

// ==============================================
// EJERCICIO 1: Operaciones Básicas con Maps
// ==============================================
/*
Crea un programa que gestione un inventario de productos usando un map.
Implementa las siguientes funciones:
- agregarProducto(inventario, nombre, cantidad)
- actualizarCantidad(inventario, nombre, nuevaCantidad)
- eliminarProducto(inventario, nombre)
- consultarProducto(inventario, nombre) - retorna cantidad y si existe
- listarInventario(inventario) - muestra todos los productos ordenados por nombre
*/

func ejercicio1() {
	fmt.Println("=== EJERCICIO 1: Operaciones Básicas con Maps ===")

	inventario := make(map[string]int)

	// TODO: Implementar las funciones y probar con los siguientes casos:
	// - Agregar: "manzanas" (50), "peras" (30), "bananas" (25)
	// - Actualizar: "manzanas" a 45
	// - Consultar: "peras" y "kiwis" (no existe)
	// - Eliminar: "bananas"
	// - Listar inventario final

	fmt.Println("Inventario inicial:", inventario)
	// Tu código aquí...
}

// ==============================================
// EJERCICIO 2: Contador de Frecuencias
// ==============================================
/*
Implementa un analizador de texto que:
1. Cuente la frecuencia de palabras en un texto
2. Cuente la frecuencia de caracteres
3. Encuentre las N palabras más frecuentes
4. Calcule estadísticas básicas (total palabras, palabras únicas, etc.)
*/

func ejercicio2() {
	fmt.Println("\n=== EJERCICIO 2: Contador de Frecuencias ===")

	texto := `Go es un lenguaje de programación moderno y eficiente. 
	Go fue diseñado por Google para ser simple, rápido y confiable. 
	Con Go puedes crear aplicaciones web, microservicios y herramientas de línea de comandos.
	Go tiene una sintaxis clara y un sistema de tipos fuerte.`

	// TODO: Implementar las siguientes funciones:
	// - contarPalabras(texto) -> map[string]int
	// - contarCaracteres(texto) -> map[rune]int
	// - palabrasMasFrecuentes(frecuencias, n) -> []string
	// - estadisticasTexto(texto) -> struct con estadísticas

	fmt.Println("Texto a analizar:", texto[:50], "...")
	// Tu código aquí...
}

// ==============================================
// EJERCICIO 3: Sistema de Caché Simple
// ==============================================
/*
Implementa un sistema de caché con las siguientes características:
- Almacena pares clave-valor con tiempo de expiración
- Tiene un tamaño máximo (LRU - Least Recently Used)
- Métodos: Set, Get, Delete, Clear, Stats
- Limpia automáticamente elementos expirados
*/

type CacheItem struct {
	Valor      interface{}
	Expiracion time.Time
	Accesos    int
}

type Cache struct {
	// TODO: Definir campos necesarios
	// items map[string]CacheItem
	// maxSize int
	// mutex sync.RWMutex
}

func ejercicio3() {
	fmt.Println("\n=== EJERCICIO 3: Sistema de Caché Simple ===")

	// TODO: Implementar métodos del cache:
	// - NewCache(maxSize) *Cache
	// - Set(key, value, duration)
	// - Get(key) (interface{}, bool)
	// - Delete(key)
	// - Clear()
	// - Stats() (total, expirados, hits, misses)
	// - CleanExpired()

	cache := &Cache{}
	fmt.Println("Cache creado:", cache)

	// Probar funcionalidad:
	// - Agregar elementos con diferentes tiempos de expiración
	// - Consultar elementos existentes y no existentes
	// - Esperar expiración y verificar limpieza automática
	// - Mostrar estadísticas

	// Tu código aquí...
}

// ==============================================
// EJERCICIO 4: Índice Invertido
// ==============================================
/*
Implementa un índice invertido para búsquedas de texto:
- Cada palabra apunta a los documentos que la contienen
- Soporte para búsquedas AND y OR
- Ranking por relevancia (frecuencia de términos)
*/

type Documento struct {
	ID        int
	Titulo    string
	Contenido string
}

type IndiceInvertido struct {
	// TODO: Definir estructura
	// indice map[string]map[int]int // palabra -> documento_id -> frecuencia
	// documentos map[int]Documento
}

func ejercicio4() {
	fmt.Println("\n=== EJERCICIO 4: Índice Invertido ===")

	documentos := []Documento{
		{1, "Introducción a Go", "Go es un lenguaje de programación desarrollado por Google"},
		{2, "Programación Concurrente", "Go facilita la programación concurrente con goroutines"},
		{3, "Desarrollo Web con Go", "Crear aplicaciones web con Go es simple y eficiente"},
		{4, "Microservicios en Go", "Go es ideal para desarrollar microservicios escalables"},
	}

	_ = documentos // Evitar error de variable no usada por ahora

	// TODO: Implementar:
	// - NewIndiceInvertido() *IndiceInvertido
	// - AgregarDocumento(doc)
	// - BuscarAND(terminos) []int (documentos que contienen TODOS los términos)
	// - BuscarOR(terminos) []int (documentos que contienen ALGÚN término)
	// - RankingPorRelevancia(terminos) []int (ordenados por relevancia)
	// - PreprocesarTexto(texto) []string (limpiar y tokenizar)

	indice := &IndiceInvertido{}
	fmt.Println("Índice creado:", indice)

	// Probar con búsquedas:
	// - "Go programación" (AND)
	// - "web microservicios" (OR)
	// - Ranking por relevancia

	// Tu código aquí...
}

// ==============================================
// EJERCICIO 5: Agrupación de Datos
// ==============================================
/*
Implementa un sistema de agrupación y agregación de datos:
- Agrupa elementos por diferentes criterios
- Calcula estadísticas por grupo (sum, avg, min, max, count)
- Soporte para múltiples niveles de agrupación
*/

type Venta struct {
	ID         int
	Producto   string
	Categoria  string
	Vendedor   string
	Fecha      time.Time
	Cantidad   int
	PrecioUnit float64
	Region     string
}

func ejercicio5() {
	fmt.Println("\n=== EJERCICIO 5: Agrupación de Datos ===")

	ventas := []Venta{
		{1, "Laptop", "Electrónicos", "Ana", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), 2, 999.99, "Norte"},
		{2, "Mouse", "Electrónicos", "Carlos", time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC), 5, 25.50, "Sur"},
		{3, "Silla", "Muebles", "Ana", time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC), 1, 150.00, "Norte"},
		{4, "Teclado", "Electrónicos", "María", time.Date(2024, 1, 18, 0, 0, 0, 0, time.UTC), 3, 75.00, "Centro"},
		{5, "Mesa", "Muebles", "Carlos", time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC), 1, 300.00, "Sur"},
	}

	// TODO: Implementar funciones:
	// - agruparPorCategoria(ventas) map[string][]Venta
	// - agruparPorVendedor(ventas) map[string][]Venta
	// - agruparPorRegionYCategoria(ventas) map[string]map[string][]Venta
	// - calcularEstadisticas(ventas) (total, promedio, min, max)
	// - resumenPorGrupo(agrupacion) - mostrar estadísticas por grupo

	fmt.Printf("Total ventas a analizar: %d\n", len(ventas))
	// Tu código aquí...
}

// ==============================================
// EJERCICIO 6: Map Thread-Safe Personalizado
// ==============================================
/*
Implementa un map thread-safe con funcionalidades avanzadas:
- Operaciones atómicas (Get, Set, Delete)
- Operaciones por lotes (SetBatch, GetBatch)
- Iteración segura
- Estadísticas de uso (hits, misses, operaciones)
- Callbacks para eventos (onSet, onDelete, onEvict)
*/

type SafeMap struct {
	// TODO: Definir estructura
	// mu sync.RWMutex
	// data map[string]interface{}
	// stats Statistics
	// callbacks Callbacks
}

type Statistics struct {
	Gets    int64
	Sets    int64
	Deletes int64
	Hits    int64
	Misses  int64
}

type Callbacks struct {
	OnSet    func(key string, value interface{})
	OnDelete func(key string)
	OnEvict  func(key string, value interface{})
}

func ejercicio6() {
	fmt.Println("\n=== EJERCICIO 6: Map Thread-Safe Personalizado ===")

	// TODO: Implementar SafeMap con métodos:
	// - NewSafeMap() *SafeMap
	// - Set(key, value)
	// - Get(key) (interface{}, bool)
	// - Delete(key) bool
	// - SetBatch(map[string]interface{})
	// - GetBatch([]string) map[string]interface{}
	// - Keys() []string
	// - Len() int
	// - Clear()
	// - GetStats() Statistics
	// - SetCallbacks(callbacks)
	// - ForEach(func(key string, value interface{}))

	safeMap := &SafeMap{}
	fmt.Println("SafeMap creado:", safeMap)

	// Probar concurrencia:
	// - Múltiples goroutines escribiendo
	// - Múltiples goroutines leyendo
	// - Operaciones por lotes
	// - Callbacks en acción
	// - Mostrar estadísticas finales

	// Tu código aquí...
}

// ==============================================
// EJERCICIO 7: Sistema de Configuración Jerárquico
// ==============================================
/*
Implementa un sistema de configuración que:
- Soporte configuración jerárquica (env.app.database.host)
- Herencia de valores (desarrollo < staging < producción)
- Validación de tipos y rangos
- Watchers para cambios en configuración
- Serialización/deserialización (JSON, YAML)
*/

type ConfigValue struct {
	Valor        interface{}
	Tipo         string
	Requerido    bool
	ValorDefecto interface{}
	Validador    func(interface{}) error
}

type ConfigManager struct {
	// TODO: Definir estructura
	// configs map[string]map[string]ConfigValue
	// herencia []string // orden de herencia
	// watchers map[string][]func(string, interface{})
}

func ejercicio7() {
	fmt.Println("\n=== EJERCICIO 7: Sistema de Configuración Jerárquico ===")

	// TODO: Implementar ConfigManager:
	// - NewConfigManager() *ConfigManager
	// - SetDefault(path, value, validator)
	// - Set(env, path, value)
	// - Get(env, path) interface{}
	// - GetWithType[T](env, path) (T, error)
	// - SetInheritance(envs []string)
	// - Watch(path, callback)
	// - Validate(env) error
	// - LoadFromJSON(env, jsonData)
	// - ToJSON(env) string

	cm := &ConfigManager{}
	fmt.Println("ConfigManager creado:", cm)

	// Configuraciones de ejemplo:
	configData := map[string]map[string]interface{}{
		"default": {
			"app.name":       "MiApp",
			"app.version":    "1.0.0",
			"database.host":  "localhost",
			"database.port":  5432,
			"database.ssl":   false,
			"cache.ttl":      300,
			"api.rate_limit": 1000,
		},
		"desarrollo": {
			"database.host": "dev-db.local",
			"debug":         true,
		},
		"produccion": {
			"database.host":  "prod-db.empresa.com",
			"database.ssl":   true,
			"api.rate_limit": 10000,
			"debug":          false,
		},
	}

	fmt.Println("Configuraciones a cargar:", configData)
	// Tu código aquí...
}

// ==============================================
// EJERCICIO 8: Router HTTP Básico
// ==============================================
/*
Implementa un router HTTP que:
- Registra rutas con diferentes métodos (GET, POST, PUT, DELETE)
- Soporte para parámetros de ruta (/users/:id)
- Middleware para autenticación, logging, CORS
- Grupos de rutas (/api/v1/...)
- Manejo de errores personalizado
*/

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

type Context struct {
	Method     HTTPMethod
	Path       string
	Params     map[string]string
	Query      map[string]string
	Headers    map[string]string
	Body       string
	StatusCode int
	Response   string
}

type Handler func(*Context) error
type Middleware func(Handler) Handler

type Router struct {
	// TODO: Definir estructura
	// routes map[HTTPMethod]map[string]Handler
	// middlewares []Middleware
	// groups map[string]*Router
}

func ejercicio8() {
	fmt.Println("\n=== EJERCICIO 8: Router HTTP Básico ===")

	// TODO: Implementar Router:
	// - NewRouter() *Router
	// - Use(middleware Middleware)
	// - Get/Post/Put/Delete(path, handler)
	// - Group(prefix) *Router
	// - Handle(method, path, body, headers) (*Context, error)
	// - extractParams(pattern, path) map[string]string
	// - applyMiddlewares(handler) Handler

	router := &Router{}
	fmt.Println("Router creado:", router)

	// Middlewares de ejemplo:
	loggingMiddleware := func(next Handler) Handler {
		return func(ctx *Context) error {
			fmt.Printf("[%s] %s\n", ctx.Method, ctx.Path)
			return next(ctx)
		}
	}

	authMiddleware := func(next Handler) Handler {
		return func(ctx *Context) error {
			if token := ctx.Headers["Authorization"]; token == "" {
				ctx.StatusCode = 401
				ctx.Response = "Unauthorized"
				return fmt.Errorf("no authorization header")
			}
			return next(ctx)
		}
	}

	// Handlers de ejemplo:
	getUsersHandler := func(ctx *Context) error {
		ctx.StatusCode = 200
		ctx.Response = `[{"id": 1, "name": "Juan"}, {"id": 2, "name": "Ana"}]`
		return nil
	}

	getUserHandler := func(ctx *Context) error {
		userID := ctx.Params["id"]
		ctx.StatusCode = 200
		ctx.Response = fmt.Sprintf(`{"id": %s, "name": "Usuario %s"}`, userID, userID)
		return nil
	}

	createUserHandler := func(ctx *Context) error {
		ctx.StatusCode = 201
		ctx.Response = `{"id": 3, "name": "Nuevo Usuario", "created": true}`
		return nil
	}

	fmt.Println("Middlewares y handlers definidos")
	fmt.Println("loggingMiddleware:", loggingMiddleware != nil)
	fmt.Println("authMiddleware:", authMiddleware != nil)
	fmt.Println("getUsersHandler:", getUsersHandler != nil)
	fmt.Println("getUserHandler:", getUserHandler != nil)
	fmt.Println("createUserHandler:", createUserHandler != nil)

	// Configurar rutas:
	// router.Use(loggingMiddleware)
	// router.Get("/users", getUsersHandler)
	// router.Get("/users/:id", getUserHandler)
	//
	// apiV1 := router.Group("/api/v1")
	// apiV1.Use(authMiddleware)
	// apiV1.Post("/users", createUserHandler)

	// Probar requests:
	// - GET /users
	// - GET /users/123
	// - POST /api/v1/users (sin auth)
	// - POST /api/v1/users (con auth)

	// Tu código aquí...
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func main() {
	fmt.Println("🗺️ EJERCICIOS DE MAPS EN GO")
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
