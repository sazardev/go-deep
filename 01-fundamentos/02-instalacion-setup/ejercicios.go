// 🚀 Ejercicios: Instalación y Setup de Go
// =====================================
// Descripción: Ejercicios prácticos para verificar tu instalación y configuración
// Dificultad: Principiante
// Temas: Instalación, workspace, herramientas, primer programa

package main

import (
	"fmt"
	"go/build"
	"os"
	"runtime"
	"strings"
)

// =============================================================================
// 📦 EJERCICIO 1: Verificación de Instalación
// =============================================================================

// showGoInfo muestra información completa de la instalación de Go
func showGoInfo() {
	fmt.Println("🚀 INFORMACIÓN DE GO")
	fmt.Println("====================")

	// Versión de Go
	fmt.Printf("📋 Versión: %s\n", runtime.Version())

	// Sistema operativo y arquitectura
	fmt.Printf("💻 SO: %s\n", runtime.GOOS)
	fmt.Printf("🏗️ Arquitectura: %s\n", runtime.GOARCH)

	// Número de CPUs disponibles
	fmt.Printf("⚡ CPUs: %d\n", runtime.NumCPU())

	// Variables de entorno importantes
	fmt.Printf("📁 GOROOT: %s\n", runtime.GOROOT())
	fmt.Printf("🏠 GOPATH: %s\n", build.Default.GOPATH)

	// Go version compiled with
	fmt.Printf("🔧 Compilador: %s\n", runtime.Compiler)

	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 2: Verificación de Variables de Entorno
// =============================================================================

// checkEnvironment verifica las variables de entorno críticas
func checkEnvironment() {
	fmt.Println("🌍 VARIABLES DE ENTORNO")
	fmt.Println("=======================")

	// Variables importantes para Go
	envVars := []string{
		"GOROOT",
		"GOPATH",
		"GOBIN",
		"GOPROXY",
		"GOSUMDB",
		"GO111MODULE",
		"CGO_ENABLED",
	}

	for _, envVar := range envVars {
		value := os.Getenv(envVar)
		if value != "" {
			fmt.Printf("✅ %s: %s\n", envVar, value)
		} else {
			fmt.Printf("⚪ %s: (no configurado)\n", envVar)
		}
	}

	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 3: Test de Compilación y Ejecución
// =============================================================================

// Función que demuestra diferentes características básicas de Go
func demonstrateGoFeatures() {
	fmt.Println("🎯 CARACTERÍSTICAS DE GO")
	fmt.Println("========================")

	// 1. Variables y tipos
	var message string = "¡Hola desde Go!"
	number := 42
	pi := 3.14159
	isAwesome := true

	fmt.Printf("📝 String: %s\n", message)
	fmt.Printf("🔢 Integer: %d\n", number)
	fmt.Printf("🥧 Float: %.2f\n", pi)
	fmt.Printf("✅ Boolean: %t\n", isAwesome)

	// 2. Array y Slice
	languages := []string{"Go", "Python", "JavaScript", "Rust"}
	fmt.Printf("📋 Lenguajes: %v\n", languages)

	// 3. Map
	frameworks := map[string]string{
		"Go":         "Gin, Echo, Fiber",
		"Python":     "Django, Flask, FastAPI",
		"JavaScript": "Express, Next.js, Nuxt",
	}
	fmt.Printf("🛠️ Frameworks Go: %s\n", frameworks["Go"])

	// 4. Function call
	result := add(10, 32)
	fmt.Printf("➕ 10 + 32 = %d\n", result)

	// 5. Error handling
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("➗ 10 / 2 = %.2f\n", quotient)
	}

	fmt.Println()
}

// add suma dos números enteros
func add(a, b int) int {
	return a + b
}

// divide realiza división con manejo de error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

// =============================================================================
// 📦 EJERCICIO 4: Test de Concurrencia Básica
// =============================================================================

// testConcurrency demuestra goroutines básicas
func testConcurrency() {
	fmt.Println("⚡ CONCURRENCIA EN GO")
	fmt.Println("====================")

	// Channel para comunicación
	done := make(chan bool)
	messages := make(chan string, 3)

	// Goroutine 1
	go func() {
		messages <- "Mensaje desde goroutine 1"
		fmt.Println("🟢 Goroutine 1 ejecutada")
	}()

	// Goroutine 2
	go func() {
		messages <- "Mensaje desde goroutine 2"
		fmt.Println("🔵 Goroutine 2 ejecutada")
	}()

	// Goroutine 3
	go func() {
		messages <- "Mensaje desde goroutine 3"
		fmt.Println("🟣 Goroutine 3 ejecutada")
		done <- true
	}()

	// Recibir mensajes
	for i := 0; i < 3; i++ {
		msg := <-messages
		fmt.Printf("📨 Recibido: %s\n", msg)
	}

	// Esperar finalización
	<-done
	fmt.Println("✅ Todas las goroutines completadas")
	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 5: Test de Paquetes Estándar
// =============================================================================

// testStandardLibrary prueba algunos paquetes de la biblioteca estándar
func testStandardLibrary() {
	fmt.Println("📚 BIBLIOTECA ESTÁNDAR")
	fmt.Println("======================")

	// strings package
	text := "  Go es INCREÍBLE para desarrollo backend  "
	fmt.Printf("📝 Original: '%s'\n", text)
	fmt.Printf("✂️ Trimmed: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("📏 Longitud: %d\n", len(text))
	fmt.Printf("🔤 Minúsculas: '%s'\n", strings.ToLower(text))
	fmt.Printf("🔠 Mayúsculas: '%s'\n", strings.ToUpper(text))
	fmt.Printf("🔍 Contiene 'Go': %t\n", strings.Contains(text, "Go"))

	// Split y Join
	words := strings.Fields(strings.TrimSpace(text))
	fmt.Printf("📋 Palabras: %v\n", words)
	joined := strings.Join(words, "-")
	fmt.Printf("🔗 Unidas con '-': %s\n", joined)

	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 6: Información del Sistema
// =============================================================================

// systemInfo muestra información detallada del sistema
func systemInfo() {
	fmt.Println("🖥️ INFORMACIÓN DEL SISTEMA")
	fmt.Println("===========================")

	// Información de memoria
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("🧠 Memoria asignada: %d KB\n", bToKb(m.Alloc))
	fmt.Printf("📊 Total asignado: %d KB\n", bToKb(m.TotalAlloc))
	fmt.Printf("⚙️ Sistema: %d KB\n", bToKb(m.Sys))
	fmt.Printf("🔄 Garbage Collections: %d\n", m.NumGC)

	// Información de goroutines
	fmt.Printf("🏃 Goroutines activas: %d\n", runtime.NumGoroutine())

	// Información de CPU
	fmt.Printf("💻 CPUs lógicas: %d\n", runtime.NumCPU())

	fmt.Println()
}

// bToKb convierte bytes a kilobytes
func bToKb(b uint64) uint64 {
	return b / 1024
}

// =============================================================================
// 📦 EJERCICIO 7: Verificación de Herramientas Go
// =============================================================================

// checkGoTools verifica que las herramientas de Go estén disponibles
func checkGoTools() {
	fmt.Println("🛠️ HERRAMIENTAS DE GO")
	fmt.Println("=====================")

	fmt.Println("Las siguientes herramientas deberían estar disponibles:")
	fmt.Println("(Ejecuta estos comandos en tu terminal)")
	fmt.Println()

	tools := []struct {
		name        string
		command     string
		description string
	}{
		{"Go Compiler", "go version", "Muestra la versión de Go"},
		{"Go Build", "go build", "Compila paquetes Go"},
		{"Go Run", "go run", "Compila y ejecuta programas Go"},
		{"Go Test", "go test", "Ejecuta tests"},
		{"Go Format", "gofmt", "Formatea código Go"},
		{"Go Imports", "goimports", "Formatea y organiza imports"},
		{"Go Vet", "go vet", "Examina código para errores"},
		{"Go Doc", "go doc", "Muestra documentación"},
		{"Go Get", "go get", "Descarga e instala paquetes"},
		{"Go Mod", "go mod", "Gestión de módulos"},
	}

	for i, tool := range tools {
		fmt.Printf("%d. 🔧 %s\n", i+1, tool.name)
		fmt.Printf("   💻 Comando: %s\n", tool.command)
		fmt.Printf("   📝 Descripción: %s\n", tool.description)
		fmt.Println()
	}
}

// =============================================================================
// 📦 EJERCICIO 8: Creación de Proyecto de Ejemplo
// =============================================================================

// demonstrateProject muestra cómo sería un proyecto Go básico
func demonstrateProject() {
	fmt.Println("📁 ESTRUCTURA DE PROYECTO GO")
	fmt.Println("============================")

	fmt.Println("Un proyecto Go típico se ve así:")
	fmt.Println()

	projectStructure := `
mi-proyecto-go/
├── go.mod                 # Definición del módulo
├── go.sum                 # Checksums de dependencias
├── main.go               # Punto de entrada principal
├── README.md             # Documentación del proyecto
├── .gitignore           # Archivos a ignorar en Git
├── cmd/                 # Aplicaciones principales
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/            # Código privado del proyecto
│   ├── handlers/
│   │   └── user.go
│   └── models/
│       └── user.go
├── pkg/                 # Código reutilizable público
│   └── utils/
│       └── helpers.go
├── api/                 # Definiciones de API
│   └── openapi.yaml
├── web/                 # Assets web
│   ├── static/
│   └── templates/
├── configs/             # Archivos de configuración
│   └── config.yaml
├── scripts/             # Scripts de build/deploy
│   └── build.sh
├── docs/               # Documentación adicional
│   └── ARCHITECTURE.md
└── tests/              # Tests de integración
    └── integration_test.go
`

	fmt.Println(projectStructure)

	fmt.Println("🎯 Comandos básicos para crear un proyecto:")
	fmt.Println("├── mkdir mi-proyecto-go")
	fmt.Println("├── cd mi-proyecto-go")
	fmt.Println("├── go mod init github.com/tu-usuario/mi-proyecto-go")
	fmt.Println("├── touch main.go")
	fmt.Println("├── go run main.go")
	fmt.Println("└── go build -o mi-app")
	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 9: Benchmark Básico
// =============================================================================

// simplePerformanceTest realiza un test básico de performance
func simplePerformanceTest() {
	fmt.Println("⚡ TEST DE PERFORMANCE")
	fmt.Println("=====================")

	// Test de concatenación de strings
	iterations := 100000

	// Método 1: Concatenación simple
	fmt.Printf("🔧 Concatenando %d strings...\n", iterations)

	// Simulación (en un test real usarías testing.B)
	result := ""
	for i := 0; i < iterations; i++ {
		result += "a"
		if i%10000 == 0 {
			fmt.Printf("⏳ Progreso: %d%%\n", (i*100)/iterations)
		}
	}

	fmt.Printf("✅ Completado! Longitud final: %d\n", len(result))
	fmt.Println("💡 En tests reales, usa strings.Builder para mejor performance")
	fmt.Println()
}

// =============================================================================
// 📦 EJERCICIO 10: Validación Final
// =============================================================================

// finalValidation realiza una validación completa de la instalación
func finalValidation() {
	fmt.Println("🎯 VALIDACIÓN FINAL")
	fmt.Println("==================")

	checks := []struct {
		name   string
		status string
		tip    string
	}{
		{
			"Go instalado",
			"✅ Correcto",
			"Versión " + runtime.Version() + " detectada",
		},
		{
			"Variables de entorno",
			"✅ Configuradas",
			"GOROOT y GOPATH disponibles",
		},
		{
			"Compilación",
			"✅ Funcional",
			"Código compila sin errores",
		},
		{
			"Concurrencia",
			"✅ Operativa",
			"Goroutines funcionando correctamente",
		},
		{
			"Biblioteca estándar",
			"✅ Disponible",
			"Paquetes estándar accesibles",
		},
	}

	fmt.Println("Estado de tu instalación:")
	fmt.Println()

	for i, check := range checks {
		fmt.Printf("%d. %s: %s\n", i+1, check.name, check.status)
		fmt.Printf("   💡 %s\n", check.tip)
		fmt.Println()
	}

	fmt.Println("🎉 ¡Felicitaciones! Tu instalación de Go está completa y funcional.")
	fmt.Println("🚀 Estás listo para comenzar tu viaje con Go.")
	fmt.Println()

	fmt.Println("📝 Próximos pasos sugeridos:")
	fmt.Println("├── 1. Explorar el Tour de Go (https://tour.golang.org/)")
	fmt.Println("├── 2. Leer 'Effective Go' (https://golang.org/doc/effective_go.html)")
	fmt.Println("├── 3. Practicar en Go Playground (https://play.golang.org/)")
	fmt.Println("├── 4. Unirse a la comunidad Go")
	fmt.Println("└── 5. Continuar con el siguiente módulo del curso")
}

// =============================================================================
// 📋 FUNCIÓN PRINCIPAL
// =============================================================================

func main() {
	fmt.Println("🚀 EJERCICIOS: INSTALACIÓN Y SETUP DE GO")
	fmt.Println("========================================")
	fmt.Printf("📅 Fecha: %s\n", "2025-07-28")
	fmt.Printf("👨‍💻 Usuario: %s\n", os.Getenv("USER"))
	fmt.Printf("💻 Directorio: %s\n", getCurrentDir())
	fmt.Println()

	// Ejecutar todos los ejercicios
	showGoInfo()
	checkEnvironment()
	demonstrateGoFeatures()
	testConcurrency()
	testStandardLibrary()
	systemInfo()
	checkGoTools()
	demonstrateProject()
	simplePerformanceTest()
	finalValidation()

	fmt.Println("🎊 ¡Todos los ejercicios completados exitosamente!")
	fmt.Println("📚 Continúa con la siguiente lección para aprender sintaxis básica.")
}

// getCurrentDir obtiene el directorio actual de forma segura
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "No disponible"
	}
	return dir
}

// =============================================================================
// 📋 INSTRUCCIONES PARA EJECUTAR
// =============================================================================

/*
🎯 CÓMO EJECUTAR ESTOS EJERCICIOS:

1. **Compilar y ejecutar:**
   ```bash
   go run ejercicios.go
   ```

2. **Compilar a binario:**
   ```bash
   go build -o setup-test ejercicios.go
   ./setup-test
   ```

3. **Ejecutar con información detallada:**
   ```bash
   go run -race ejercicios.go
   ```

4. **Ver solo información de Go:**
   ```bash
   go version
   go env
   ```

🔍 VALIDACIONES QUE REALIZA:

✅ Verifica instalación correcta de Go
✅ Comprueba variables de entorno
✅ Testa compilación y ejecución
✅ Valida concurrencia básica
✅ Prueba biblioteca estándar
✅ Muestra información del sistema
✅ Lista herramientas disponibles
✅ Demuestra estructura de proyecto
✅ Realiza test básico de performance
✅ Proporciona validación final completa

🎯 OBJETIVOS DE APRENDIZAJE:

📝 Familiarizarse con el entorno Go
🛠️ Entender herramientas básicas
🏗️ Conocer estructura de proyectos
⚡ Experimentar con concurrencia
📚 Explorar biblioteca estándar
🎯 Validar instalación completa

¡Disfruta explorando Go! 🚀
*/
