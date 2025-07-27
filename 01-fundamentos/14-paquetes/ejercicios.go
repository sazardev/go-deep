// Lección 14: Paquetes y Módulos en Go
// Ejercicios prácticos para dominar la organización de código

/*
INSTRUCCIONES GENERALES:

Este archivo contiene ejercicios prácticos para aprender paquetes y módulos.
Cada ejercicio está diseñado para construir sobre el anterior.

Para completar los ejercicios:
1. Lee cada sección TODO
2. Implementa el código solicitado
3. Ejecuta los tests para validar tu implementación
4. Experimenta con diferentes variaciones

Estructura que crearemos:
biblioteca-system/
├── go.mod
├── pkg/
│   ├── models/
│   ├── validator/
│   └── logger/
├── internal/
│   ├── repository/
│   └── services/
└── cmd/
    └── main.go
*/

package main

import "fmt"

// ========================================
// Ejercicio 1: Configuración Inicial del Módulo
// ========================================

/*
TODO: Ejercicio 1 - Configuración del Proyecto

1. Crea la estructura de directorios mostrada arriba
2. Inicializa un módulo Go con: go mod init biblioteca-system
3. Crea los archivos base para cada paquete

Comandos a ejecutar:
mkdir -p pkg/{models,validator,logger}
mkdir -p internal/{repository,services}
mkdir -p cmd
go mod init biblioteca-system

Este ejercicio se hace en terminal, no en código.
*/

func ejercicio1Info() {
	fmt.Println("=== Ejercicio 1: Configuración del Proyecto ===")
	fmt.Println("Ejecuta los comandos mostrados en el comentario TODO")
	fmt.Println("✅ Ejercicio 1: Configuración completada en terminal\n")
}

// ========================================
// Ejercicio 2: Paquete Models - Definir Estructuras
// ========================================

/*
TODO: Ejercicio 2 - Crear pkg/models/book.go

Crea un archivo pkg/models/book.go con:

1. Struct Book con campos:
   - ID (público, int)
   - Title (público, string)
   - Author (público, string)
   - ISBN (público, string)
   - Available (público, bool)
   - publishedYear (privado, int)
   - category (privado, string)

2. Constructor NewBook que:
   - Acepte title, author, isbn, year, category
   - Genere ID automáticamente
   - Establezca Available como true por defecto

3. Métodos públicos:
   - GetPublishedYear() int
   - GetCategory() string
   - SetCategory(category string)
   - IsClassic() bool (libros antes de 1980)
   - GetDisplayInfo() string

4. Métodos privados:
   - generateID() int (usa timestamp o contador)

Ejemplo de uso esperado:
book := models.NewBook("Go Programming", "John Doe", "123-456", 2024, "Technology")
fmt.Println(book.GetDisplayInfo())
*/

func ejercicio2Info() {
	fmt.Println("=== Ejercicio 2: Paquete Models ===")
	fmt.Println("Crear pkg/models/book.go con struct Book y métodos")
	fmt.Println("Ver TODO arriba para detalles completos")
	fmt.Println("TODO: Implementar en pkg/models/book.go\n")
}

// ========================================
// Ejercicio 3: Paquete Models - Más Estructuras
// ========================================

/*
TODO: Ejercicio 3 - Crear pkg/models/user.go

Crea un archivo pkg/models/user.go con:

1. Struct User con campos:
   - ID (público, int)
   - Name (público, string)
   - Email (público, string)
   - registrationDate (privado, time.Time)
   - borrowedBooks (privado, []int) // IDs de libros prestados

2. Constructor NewUser que:
   - Acepte name y email
   - Genere ID automáticamente
   - Establezca registrationDate como time.Now()
   - Inicialice borrowedBooks como slice vacío

3. Métodos públicos:
   - GetRegistrationDate() time.Time
   - GetBorrowedBooks() []int (retorna copia)
   - BorrowBook(bookID int) error
   - ReturnBook(bookID int) error
   - CanBorrowMore() bool (máximo 5 libros)
   - GetMembershipDuration() time.Duration

4. Validaciones:
   - Email debe contener @
   - Name no puede estar vacío
   - No puede prestar libro ya prestado

Ejemplo de uso:
user := models.NewUser("Alice Smith", "alice@email.com")
err := user.BorrowBook(123)
*/

func ejercicio3Info() {
	fmt.Println("=== Ejercicio 3: Estructura User ===")
	fmt.Println("Crear pkg/models/user.go con struct User y métodos")
	fmt.Println("Incluir validaciones y lógica de préstamos")
	fmt.Println("TODO: Implementar en pkg/models/user.go\n")
}

// ========================================
// Ejercicio 4: Paquete Validator - Validaciones Reutilizables
// ========================================

/*
TODO: Ejercicio 4 - Crear pkg/validator/validator.go

Crea un paquete de validaciones con:

1. Funciones públicas:
   - ValidateEmail(email string) error
   - ValidateISBN(isbn string) error (formato XXX-XXX o XXX-XXXX-XXX)
   - ValidateName(name string) error
   - ValidateYear(year int) error (entre 1000 y año actual)

2. Struct ValidationResult:
   - IsValid (público, bool)
   - Errors (público, []string)

3. Función ValidateBook que:
   - Acepta un Book del paquete models
   - Retorna ValidationResult
   - Valida todos los campos del libro

4. Función ValidateUser que:
   - Acepta un User del paquete models
   - Retorna ValidationResult
   - Valida nombre y email

5. Funciones privadas auxiliares:
   - isValidEmailFormat(email string) bool
   - isValidISBNFormat(isbn string) bool
   - containsOnlyLetters(s string) bool

Ejemplo de uso:
result := validator.ValidateBook(book)
if !result.IsValid {
    for _, err := range result.Errors {
        fmt.Println("Error:", err)
    }
}
*/

func ejercicio4Info() {
	fmt.Println("=== Ejercicio 4: Paquete Validator ===")
	fmt.Println("Crear pkg/validator/validator.go con validaciones")
	fmt.Println("Incluir ValidationResult y funciones de validación")
	fmt.Println("TODO: Implementar en pkg/validator/validator.go\n")
}

// ========================================
// Ejercicio 5: Paquete Logger - Sistema de Logging
// ========================================

/*
TODO: Ejercicio 5 - Crear pkg/logger/logger.go

Crea un sistema de logging con:

1. Enum de niveles (usando constantes):
   - DEBUG = 0
   - INFO = 1
   - WARN = 2
   - ERROR = 3

2. Interface Logger:
   - Debug(message string)
   - Info(message string)
   - Warn(message string)
   - Error(message string)
   - LogWithLevel(level int, message string)

3. Struct ConsoleLogger que implementa Logger:
   - includeTimestamp (privado, bool)
   - minLevel (privado, int)

4. Constructor NewConsoleLogger:
   - Acepta includeTimestamp y minLevel
   - Retorna Logger (interface, no struct)

5. Struct FileLogger que implementa Logger:
   - filename (privado, string)
   - file (privado, *os.File)

6. Constructor NewFileLogger:
   - Acepta filename
   - Abre archivo para escritura (append mode)
   - Retorna Logger y error

7. Función auxiliar privada:
   - formatMessage(level int, message string, includeTimestamp bool) string

Ejemplo de uso:
consoleLogger := logger.NewConsoleLogger(true, logger.INFO)
fileLogger, err := logger.NewFileLogger("app.log")
consoleLogger.Info("Sistema iniciado")
*/

func ejercicio5Info() {
	fmt.Println("=== Ejercicio 5: Paquete Logger ===")
	fmt.Println("Crear pkg/logger/logger.go con sistema de logging")
	fmt.Println("Implementar ConsoleLogger y FileLogger")
	fmt.Println("TODO: Implementar en pkg/logger/logger.go\n")
}

// ========================================
// Ejercicio 6: Repository Pattern - Acceso a Datos
// ========================================

/*
TODO: Ejercicio 6 - Crear internal/repository/book_repository.go

Crea un repository para manejar datos:

1. Interface BookRepository:
   - Save(book *models.Book) error
   - FindByID(id int) (*models.Book, error)
   - FindAll() ([]*models.Book, error)
   - FindByAuthor(author string) ([]*models.Book, error)
   - FindAvailable() ([]*models.Book, error)
   - Update(book *models.Book) error
   - Delete(id int) error

2. Struct InMemoryBookRepository que implementa BookRepository:
   - books (privado, map[int]*models.Book)
   - mutex (privado, sync.RWMutex)
   - logger (privado, logger.Logger)

3. Constructor NewInMemoryBookRepository:
   - Acepta logger.Logger
   - Inicializa maps y mutex
   - Retorna BookRepository (interface)

4. Implementar todos los métodos con:
   - Logging apropiado
   - Thread safety (usar mutex)
   - Validaciones
   - Manejo de errores

5. Crear internal/repository/user_repository.go similar:
   - Interface UserRepository
   - Struct InMemoryUserRepository
   - Métodos para CRUD de usuarios

Ejemplo:
repo := repository.NewInMemoryBookRepository(myLogger)
book := models.NewBook("Go Guide", "Expert", "123-456", 2024, "Tech")
err := repo.Save(book)
*/

func ejercicio6Info() {
	fmt.Println("=== Ejercicio 6: Repository Pattern ===")
	fmt.Println("Crear internal/repository/ con BookRepository y UserRepository")
	fmt.Println("Implementar pattern Repository con interfaces")
	fmt.Println("TODO: Implementar repositories en internal/repository/\n")
}

// ========================================
// Ejercicio 7: Service Layer - Lógica de Negocio
// ========================================

/*
TODO: Ejercicio 7 - Crear internal/services/library_service.go

Crea la capa de servicios con lógica de negocio:

1. Struct LibraryService:
   - bookRepo (privado, repository.BookRepository)
   - userRepo (privado, repository.UserRepository)
   - validator (privado, validator.Validator) // Crear interface Validator
   - logger (privado, logger.Logger)

2. Constructor NewLibraryService:
   - Acepta todas las dependencias
   - Retorna *LibraryService

3. Métodos de gestión de libros:
   - AddBook(title, author, isbn string, year int, category string) (*models.Book, error)
   - GetBook(id int) (*models.Book, error)
   - GetAllBooks() ([]*models.Book, error)
   - GetBooksByAuthor(author string) ([]*models.Book, error)
   - GetAvailableBooks() ([]*models.Book, error)
   - UpdateBook(book *models.Book) error
   - RemoveBook(id int) error

4. Métodos de gestión de usuarios:
   - RegisterUser(name, email string) (*models.User, error)
   - GetUser(id int) (*models.User, error)
   - GetAllUsers() ([]*models.User, error)

5. Métodos de préstamos:
   - BorrowBook(userID, bookID int) error
   - ReturnBook(userID, bookID int) error
   - GetUserBorrowedBooks(userID int) ([]*models.Book, error)

6. Cada método debe:
   - Validar entrada usando validator
   - Usar logger para eventos importantes
   - Manejar errores apropiadamente
   - Implementar lógica de negocio

Ejemplo:
service := services.NewLibraryService(bookRepo, userRepo, validator, logger)
book, err := service.AddBook("Go Patterns", "John Doe", "123-456", 2024, "Programming")
*/

func ejercicio7Info() {
	fmt.Println("=== Ejercicio 7: Service Layer ===")
	fmt.Println("Crear internal/services/library_service.go")
	fmt.Println("Implementar lógica de negocio y coordinación")
	fmt.Println("TODO: Implementar en internal/services/library_service.go\n")
}

// ========================================
// Ejercicio 8: Aplicación Principal - Integrando Todo
// ========================================

/*
TODO: Ejercicio 8 - Crear cmd/main.go

Crea la aplicación principal que integra todos los paquetes:

1. Función main que:
   - Inicializa logger (console y file)
   - Crea repositories
   - Crea validator
   - Crea service layer
   - Ejecuta demostración completa

2. Función setupDependencies que:
   - Retorna todos los componentes configurados
   - Maneja errores de inicialización

3. Función demonstrateLibrarySystem que:
   - Registra algunos usuarios
   - Añade algunos libros
   - Realiza préstamos
   - Muestra reportes
   - Demuestra manejo de errores

4. Función printSystemReport que:
   - Muestra estadísticas del sistema
   - Lista libros disponibles
   - Lista usuarios registrados
   - Muestra préstamos activos

Estructura esperada del main:
```go
func main() {
    logger, bookRepo, userRepo, validator, service := setupDependencies()

    demonstrateLibrarySystem(service, logger)

    printSystemReport(service, logger)
}
```

5. Crear también cmd/cli/main.go con interfaz de línea de comandos básica.
*/

func ejercicio8Info() {
	fmt.Println("=== Ejercicio 8: Aplicación Principal ===")
	fmt.Println("Crear cmd/main.go integrando todos los paquetes")
	fmt.Println("Demostrar funcionalidad completa del sistema")
	fmt.Println("TODO: Implementar en cmd/main.go\n")
}

// ========================================
// Ejercicio 9: Testing Cross-Package
// ========================================

/*
TODO: Ejercicio 9 - Crear tests para validar integración

Crea tests que validen la integración entre paquetes:

1. pkg/models/book_test.go:
   - Tests unitarios para Book
   - Tests de validación de campos
   - Tests de métodos públicos

2. pkg/validator/validator_test.go:
   - Tests para todas las funciones de validación
   - Tests con casos edge
   - Tests de integración con models

3. internal/services/library_service_test.go:
   - Tests de integración usando mocks
   - Tests de flujos completos
   - Tests de manejo de errores

4. Crear pkg/mocks/ con mocks para testing:
   - MockBookRepository
   - MockUserRepository
   - MockLogger
   - MockValidator

5. Tests de integración en integration_test.go:
   - Tests end-to-end
   - Tests de múltiples operaciones
   - Tests de concurrencia

Comandos para ejecutar:
go test ./...
go test -v ./pkg/...
go test -v ./internal/...
go test -race ./...
*/

func ejercicio9Info() {
	fmt.Println("=== Ejercicio 9: Testing Cross-Package ===")
	fmt.Println("Crear tests para validar integración entre paquetes")
	fmt.Println("Incluir tests unitarios, de integración y mocks")
	fmt.Println("TODO: Implementar tests en todos los paquetes\n")
}

// ========================================
// Ejercicio 10: Gestión de Dependencias y Versionado
// ========================================

/*
TODO: Ejercicio 10 - Dependencias Externas y Versionado

1. Añadir dependencias externas útiles:
   go get github.com/google/uuid        # Para IDs únicos
   go get github.com/stretchr/testify   # Para testing avanzado
   go get gopkg.in/yaml.v3             # Para configuración
   go get github.com/gorilla/mux        # Para HTTP routing (opcional)

2. Modificar los modelos para usar UUID:
   - Cambiar ID de int a string
   - Usar uuid.New() para generar IDs
   - Actualizar todos los métodos relacionados

3. Crear pkg/config/ para gestión de configuración:
   - Struct Config con parámetros del sistema
   - Carga desde archivo YAML
   - Variables de entorno override
   - Validación de configuración

4. Añadir tags de versión al proyecto:
   git tag v0.1.0
   git tag v0.2.0
   git tag v1.0.0

5. Crear ejemplo de uso como módulo externo:
   - Crear directorio example/
   - go mod init example-usage
   - Importar biblioteca-system como dependencia
   - Demostrar uso de la API pública

6. Optimizar go.mod:
   - Ejecutar go mod tidy
   - Verificar go mod verify
   - Revisar go.sum

Estructura final esperada:
biblioteca-system/
├── go.mod (con dependencias)
├── go.sum
├── configs/
│   ├── config.yaml
│   └── test.yaml
├── example/
│   ├── go.mod
│   └── main.go
└── [resto de la estructura]
*/

func ejercicio10Info() {
	fmt.Println("=== Ejercicio 10: Dependencias y Versionado ===")
	fmt.Println("Añadir dependencias externas y gestión de versiones")
	fmt.Println("Crear configuración y ejemplo de uso externo")
	fmt.Println("TODO: Implementar gestión completa de dependencias\n")
}

// ========================================
// Ejercicio Bonus: API HTTP y Deploy
// ========================================

/*
TODO: Ejercicio Bonus - API REST y Deploy

Si quieres llevar el proyecto al siguiente nivel:

1. Crear pkg/api/ con endpoints REST:
   - GET /books
   - POST /books
   - GET /books/{id}
   - PUT /books/{id}
   - DELETE /books/{id}
   - POST /users/{id}/borrow/{bookId}
   - POST /users/{id}/return/{bookId}

2. Añadir middleware:
   - Logging de requests
   - CORS
   - Rate limiting
   - Authentication básica

3. Crear Dockerfile:
   - Multi-stage build
   - Imagen mínima
   - Security best practices

4. Crear docker-compose.yml:
   - App container
   - PostgreSQL container
   - Redis container (cache)

5. Deploy a la nube:
   - Heroku, Railway, o DigitalOcean
   - Variables de entorno
   - Health checks
   - Monitoring básico

Este ejercicio bonus te prepara para desarrollo real con Go!
*/

func ejercicioBonusInfo() {
	fmt.Println("=== Ejercicio Bonus: API REST y Deploy ===")
	fmt.Println("Crear API HTTP completa y deploy a producción")
	fmt.Println("Incluir Docker, middleware y monitoring")
	fmt.Println("TODO: Implementar API REST completa\n")
}

// ========================================
// Función principal para ejecutar ejemplos
// ========================================

func main() {
	fmt.Println("🎪 Ejercicios: Paquetes y Módulos en Go")
	fmt.Println("=====================================")
	fmt.Println()

	ejercicio1Info()
	ejercicio2Info()
	ejercicio3Info()
	ejercicio4Info()
	ejercicio5Info()
	ejercicio6Info()
	ejercicio7Info()
	ejercicio8Info()
	ejercicio9Info()
	ejercicio10Info()
	ejercicioBonusInfo()

	fmt.Println("🎯 Resumen de Ejercicios:")
	fmt.Println("1. ✅ Configuración del proyecto")
	fmt.Println("2. 📚 Paquete models (Book, User)")
	fmt.Println("3. ✔️  Paquete validator")
	fmt.Println("4. 📝 Paquete logger")
	fmt.Println("5. 🗄️  Repository pattern")
	fmt.Println("6. ⚙️  Service layer")
	fmt.Println("7. 🚀 Aplicación principal")
	fmt.Println("8. 🧪 Testing cross-package")
	fmt.Println("9. 📦 Gestión de dependencias")
	fmt.Println("10. 🌐 Bonus: API REST y deploy")
	fmt.Println()
	fmt.Println("💡 Cada ejercicio construye sobre el anterior")
	fmt.Println("🔗 Al final tendrás un sistema completo de biblioteca")
	fmt.Println("🎯 Perfecto para tu portfolio!")
}
