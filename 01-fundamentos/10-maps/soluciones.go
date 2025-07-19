// Archivo: soluciones.go
// Este archivo contiene las soluciones completas para todos los ejercicios de mapas.
// Para ejecutar: go run soluciones.go

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// ========================================
// SOLUCIONES A LOS EJERCICIOS DE MAPAS
// ========================================

// Estructuras para la solución de caching
type CacheItemSol struct {
	value     interface{}
	expiresAt time.Time
}

type CacheSol struct {
	data map[string]CacheItemSol
	mu   sync.RWMutex
}

func NewCacheSol() *CacheSol {
	return &CacheSol{
		data: make(map[string]CacheItemSol),
	}
}

func (c *CacheSol) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheItemSol{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

func (c *CacheSol) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.expiresAt) {
		delete(c.data, key)
		return nil, false
	}

	return item.value, true
}

func (c *CacheSol) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]CacheItemSol)
}

// Estructuras para la solución del índice invertido
type IndexSol struct {
	index map[string][]string
	mu    sync.RWMutex
}

func NewIndexSol() *IndexSol {
	return &IndexSol{
		index: make(map[string][]string),
	}
}

func (idx *IndexSol) AddDocument(docID, content string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	words := strings.Fields(strings.ToLower(content))
	for _, word := range words {
		found := false
		for _, doc := range idx.index[word] {
			if doc == docID {
				found = true
				break
			}
		}
		if !found {
			idx.index[word] = append(idx.index[word], docID)
		}
	}
}

func (idx *IndexSol) Search(term string) []string {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	return idx.index[strings.ToLower(term)]
}

// Estructura para datos agrupados
type GroupedDataSol struct {
	groups map[string][]interface{}
	mu     sync.RWMutex
}

func NewGroupedDataSol() *GroupedDataSol {
	return &GroupedDataSol{
		groups: make(map[string][]interface{}),
	}
}

func (gd *GroupedDataSol) AddItem(key string, item interface{}) {
	gd.mu.Lock()
	defer gd.mu.Unlock()
	gd.groups[key] = append(gd.groups[key], item)
}

func (gd *GroupedDataSol) GetGroup(key string) []interface{} {
	gd.mu.RLock()
	defer gd.mu.RUnlock()
	return gd.groups[key]
}

// Estructura para el mapa thread-safe
type SafeMapSol struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewSafeMapSol() *SafeMapSol {
	return &SafeMapSol{
		data: make(map[string]interface{}),
	}
}

func (sm *SafeMapSol) Set(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMapSol) Get(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMapSol) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMapSol) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// Estructura para el sistema de configuración
type ConfigSol struct {
	config map[string]interface{}
	mu     sync.RWMutex
}

func NewConfigSol() *ConfigSol {
	return &ConfigSol{
		config: make(map[string]interface{}),
	}
}

func (c *ConfigSol) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.config[key] = value
}

func (c *ConfigSol) GetString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, exists := c.config[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func (c *ConfigSol) GetInt(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, exists := c.config[key]; exists {
		if i, ok := value.(int); ok {
			return i
		}
	}
	return 0
}

func (c *ConfigSol) GetBool(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if value, exists := c.config[key]; exists {
		if b, ok := value.(bool); ok {
			return b
		}
	}
	return false
}

// Estructura para el router HTTP
type RouterSol struct {
	routes map[string]func(string) string
	mu     sync.RWMutex
}

func NewRouterSol() *RouterSol {
	return &RouterSol{
		routes: make(map[string]func(string) string),
	}
}

func (r *RouterSol) AddRoute(path string, handler func(string) string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[path] = handler
}

func (r *RouterSol) HandleRequest(path, params string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if handler, exists := r.routes[path]; exists {
		return handler(params)
	}
	return "404 - Ruta no encontrada"
}

func (r *RouterSol) ListRoutes() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	routes := make([]string, 0, len(r.routes))
	for route := range r.routes {
		routes = append(routes, route)
	}
	return routes
}

func main() {
	fmt.Println("=== SOLUCIONES DE EJERCICIOS - MAPAS ===")

	// Ejercicio 1: Operaciones básicas con mapas
	fmt.Println("\n1. Operaciones básicas con mapas:")
	estudiantes := map[string]int{
		"Juan":  85,
		"María": 92,
		"Pedro": 78,
		"Ana":   96,
	}

	// Agregar estudiante
	estudiantes["Luis"] = 88

	// Calcular promedio
	total := 0
	for _, nota := range estudiantes {
		total += nota
	}
	promedio := float64(total) / float64(len(estudiantes))

	// Encontrar mejor estudiante
	mejorEstudiante := ""
	mejorNota := 0
	for nombre, nota := range estudiantes {
		if nota > mejorNota {
			mejorNota = nota
			mejorEstudiante = nombre
		}
	}

	fmt.Printf("Promedio: %.2f\n", promedio)
	fmt.Printf("Mejor estudiante: %s con nota %d\n", mejorEstudiante, mejorNota)

	// Ejercicio 2: Contador de frecuencias
	fmt.Println("\n2. Contador de frecuencias:")
	texto := "Go es genial Go es rápido Go es fácil de aprender"
	palabras := strings.Fields(strings.ToLower(texto))

	frecuencias := make(map[string]int)
	for _, palabra := range palabras {
		frecuencias[palabra]++
	}

	fmt.Println("Frecuencias de palabras:")
	for palabra, freq := range frecuencias {
		fmt.Printf("%s: %d\n", palabra, freq)
	}

	// Ejercicio 3: Sistema de caché
	fmt.Println("\n3. Sistema de caché:")
	cache := NewCacheSol()

	// Agregar elementos al caché
	cache.Set("user:1", "Juan Pérez", 2*time.Second)
	cache.Set("user:2", "María González", 3*time.Second)

	// Obtener elementos
	if value, exists := cache.Get("user:1"); exists {
		fmt.Printf("Cache hit: %v\n", value)
	}

	// Esperar a que expire
	time.Sleep(3 * time.Second)
	if _, exists := cache.Get("user:1"); !exists {
		fmt.Println("Cache miss: elemento expirado")
	}

	// Ejercicio 4: Índice invertido
	fmt.Println("\n4. Índice invertido:")
	index := NewIndexSol()

	index.AddDocument("doc1", "Go es un lenguaje de programación")
	index.AddDocument("doc2", "Python es fácil de aprender")
	index.AddDocument("doc3", "Go es rápido y eficiente")

	resultados := index.Search("go")
	fmt.Printf("Documentos que contienen 'go': %v\n", resultados)

	// Ejercicio 5: Agrupación de datos
	fmt.Println("\n5. Agrupación de datos:")
	productos := []map[string]interface{}{
		{"nombre": "Laptop", "categoria": "Electrónicos", "precio": 1200},
		{"nombre": "Mouse", "categoria": "Electrónicos", "precio": 25},
		{"nombre": "Silla", "categoria": "Muebles", "precio": 150},
		{"nombre": "Mesa", "categoria": "Muebles", "precio": 300},
	}

	grouped := NewGroupedDataSol()
	for _, producto := range productos {
		categoria := producto["categoria"].(string)
		grouped.AddItem(categoria, producto)
	}

	electronicos := grouped.GetGroup("Electrónicos")
	fmt.Printf("Productos electrónicos: %d\n", len(electronicos))

	// Ejercicio 6: Mapa thread-safe
	fmt.Println("\n6. Mapa thread-safe:")
	safeMap := NewSafeMapSol()

	var wg sync.WaitGroup

	// Escritores concurrentes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key-%d", id), fmt.Sprintf("value-%d", id))
		}(i)
	}

	wg.Wait()
	fmt.Printf("Elementos en el mapa: %d\n", safeMap.Len())

	// Ejercicio 7: Sistema de configuración
	fmt.Println("\n7. Sistema de configuración:")
	config := NewConfigSol()

	config.Set("app.name", "Mi Aplicación")
	config.Set("app.port", 8080)
	config.Set("app.debug", true)

	fmt.Printf("Nombre: %s\n", config.GetString("app.name"))
	fmt.Printf("Puerto: %d\n", config.GetInt("app.port"))
	fmt.Printf("Debug: %t\n", config.GetBool("app.debug"))

	// Ejercicio 8: Router HTTP simple
	fmt.Println("\n8. Router HTTP simple:")
	router := NewRouterSol()

	router.AddRoute("/", func(params string) string {
		return "Página de inicio"
	})

	router.AddRoute("/users", func(params string) string {
		return "Lista de usuarios"
	})

	router.AddRoute("/user", func(params string) string {
		return fmt.Sprintf("Usuario: %s", params)
	})

	fmt.Println(router.HandleRequest("/", ""))
	fmt.Println(router.HandleRequest("/users", ""))
	fmt.Println(router.HandleRequest("/user", "123"))
	fmt.Println(router.HandleRequest("/nonexistent", ""))

	fmt.Printf("Rutas disponibles: %v\n", router.ListRoutes())

	fmt.Println("\n=== FIN DE LAS SOLUCIONES ===")
}
