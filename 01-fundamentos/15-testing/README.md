# 🧪 Lección 15: Testing en Go

> *"Los tests son la documentación ejecutable de tu código"* - Filosofía Go

## 🎯 Objetivos de Esta Lección

Al finalizar esta lección serás capaz de:
- ✅ Escribir **tests unitarios** efectivos en Go
- ✅ Dominar **table-driven tests** para casos múltiples
- ✅ Implementar **benchmarks** para medir performance
- ✅ Crear **mocks** y **stubs** para dependencias
- ✅ Realizar **testing de integración** y **end-to-end**
- ✅ Configurar **CI/CD** con coverage reporting
- ✅ Aplicar **TDD** (Test-Driven Development) en Go

---

## 🌟 ¿Por Qué Testing es Crucial?

El testing en Go no es solo una buena práctica, es una **filosofía fundamental**. Go incluye testing como parte del lenguaje, no como una librería externa.

### 🧠 Analogía: Testing como Control de Calidad

Imagina que estás construyendo un avión:

```mermaid
graph TD
    A[✈️ Código sin Tests] --> B[❌ Vuelo Directo]
    B --> C[💥 Crash en Producción]
    
    D[✈️ Código con Tests] --> E[🧪 Simulaciones]
    E --> F[🔧 Pruebas en Tierra]
    F --> G[🛩️ Vuelos de Prueba]
    G --> H[✅ Vuelo Comercial Seguro]
    
    style A fill:#ff9999
    style C fill:#ff0000
    style D fill:#99ff99
    style H fill:#00ff00
```

Los tests son como:
- ✅ **Simuladores de vuelo** - Tests unitarios
- ✅ **Pruebas de sistemas** - Tests de integración  
- ✅ **Vuelos de prueba** - Tests end-to-end
- ✅ **Inspecciones** - Code coverage

---

## 🏗️ Anatomía del Testing en Go

### 📁 Estructura de Tests

```
mi-proyecto/
├── calculator.go
├── calculator_test.go          # Tests unitarios
├── integration_test.go         # Tests de integración
├── example_test.go            # Examples (documentación)
├── benchmark_test.go          # Benchmarks
└── testdata/                  # Datos de prueba
    ├── input.json
    └── expected.txt
```

### 📋 Convenciones Fundamentales

1. **Archivos de test terminan en `_test.go`**
2. **Funciones de test empiezan con `Test`**
3. **Benchmarks empiezan con `Benchmark`**
4. **Examples empiezan con `Example`**
5. **Primer parámetro siempre es `*testing.T` o `*testing.B`**

---

## ✅ Tests Unitarios: Fundamentos

### 🎯 Tu Primer Test

**`calculator.go`**
```go
package calculator

import "errors"

// Add suma dos números enteros
func Add(a, b int) int {
    return a + b
}

// Divide divide dos números
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// IsEven verifica si un número es par
func IsEven(n int) bool {
    return n%2 == 0
}
```

**`calculator_test.go`**
```go
package calculator

import "testing"

// Test básico - función simple
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Test con múltiples assertions
func TestAddMultiple(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }
    
    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", 
                test.a, test.b, result, test.expected)
        }
    }
}

// Test con manejo de errores
func TestDivide(t *testing.T) {
    // Test caso exitoso
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned unexpected error: %v", err)
    }
    if result != 5 {
        t.Errorf("Divide(10, 2) = %d; want 5", result)
    }
    
    // Test caso de error
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}
```

### 🎯 Table-Driven Tests: El Patrón Go

Los **table-driven tests** son la forma idiomática de testing en Go:

```go
func TestIsEven(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"positive even", 4, true},
        {"positive odd", 5, false},
        {"zero", 0, true},
        {"negative even", -4, true},
        {"negative odd", -3, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsEven(tt.input)
            if result != tt.expected {
                t.Errorf("IsEven(%d) = %v; want %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

### 🔧 Métodos de testing.T

```go
func TestMethods(t *testing.T) {
    // Error sin parar el test
    if condition {
        t.Error("Something went wrong")
        t.Errorf("Expected %d, got %d", expected, actual)
    }
    
    // Fatal: para el test inmediatamente
    if criticalCondition {
        t.Fatal("Critical error, stopping test")
        t.Fatalf("Critical: expected %d, got %d", expected, actual)
    }
    
    // Logging
    t.Log("This is a log message")
    t.Logf("Value: %d", value)
    
    // Skip test
    if runtime.GOOS == "windows" {
        t.Skip("Skipping on Windows")
    }
    
    // Helper (marca función como helper)
    t.Helper() // Esta función no aparecerá en stack traces
}
```

---

## 🏃 Subtests: Organización Avanzada

### 🎯 Subtests con t.Run

```go
func TestStringManipulation(t *testing.T) {
    input := "Hello, World!"
    
    t.Run("Length", func(t *testing.T) {
        if len(input) != 13 {
            t.Errorf("Expected length 13, got %d", len(input))
        }
    })
    
    t.Run("Contains", func(t *testing.T) {
        if !strings.Contains(input, "World") {
            t.Error("Should contain 'World'")
        }
    })
    
    t.Run("ToUpper", func(t *testing.T) {
        result := strings.ToUpper(input)
        expected := "HELLO, WORLD!"
        if result != expected {
            t.Errorf("Expected %q, got %q", expected, result)
        }
    })
}
```

### 🎯 Parallel Tests para Performance

```go
func TestParallelOperations(t *testing.T) {
    tests := []struct {
        name  string
        input int
    }{
        {"small", 10},
        {"medium", 100},
        {"large", 1000},
    }
    
    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Ejecuta en paralelo
            
            result := expensiveOperation(tt.input)
            if result < 0 {
                t.Errorf("Expected positive result for %d", tt.input)
            }
        })
    }
}
```

---

## 🚀 Benchmarks: Midiendo Performance

### 📊 Benchmarks Básicos

```go
// benchmark_test.go
package calculator

import (
    "testing"
    "math/rand"
)

// Benchmark simple
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}

// Benchmark con setup
func BenchmarkDivide(b *testing.B) {
    // Setup (no se mide)
    testCases := make([][2]int, 1000)
    for i := range testCases {
        testCases[i] = [2]int{rand.Intn(1000), rand.Intn(100) + 1}
    }
    
    b.ResetTimer() // Reset después del setup
    
    for i := 0; i < b.N; i++ {
        tc := testCases[i%len(testCases)]
        Divide(tc[0], tc[1])
    }
}

// Benchmark con diferentes tamaños
func BenchmarkStringBuilder(b *testing.B) {
    sizes := []int{10, 100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                var builder strings.Builder
                for j := 0; j < size; j++ {
                    builder.WriteString("a")
                }
                _ = builder.String()
            }
        })
    }
}

// Benchmark de memoria
func BenchmarkMemoryAllocation(b *testing.B) {
    b.ReportAllocs() // Reporta allocations
    
    for i := 0; i < b.N; i++ {
        slice := make([]int, 1000)
        _ = slice
    }
}
```

### 📈 Ejecutando Benchmarks

```bash
# Ejecutar todos los benchmarks
go test -bench=.

# Benchmark específico
go test -bench=BenchmarkAdd

# Con información de memoria
go test -bench=. -benchmem

# Múltiples iteraciones para precisión
go test -bench=. -count=5

# Guardar resultados para comparación
go test -bench=. > old.txt
# ... hacer cambios ...
go test -bench=. > new.txt
benchcmp old.txt new.txt
```

---

## 🎭 Mocking: Simulando Dependencias

### 🏗️ Interfaces para Testing

**`user_service.go`**
```go
package user

import (
    "errors"
    "fmt"
)

// User representa un usuario
type User struct {
    ID    int
    Name  string
    Email string
}

// UserRepository abstrae el acceso a datos
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
    FindByEmail(email string) (*User, error)
}

// EmailService abstrae el envío de emails
type EmailService interface {
    SendWelcomeEmail(email string) error
}

// UserService maneja lógica de negocio
type UserService struct {
    repo  UserRepository
    email EmailService
}

// NewUserService constructor
func NewUserService(repo UserRepository, email EmailService) *UserService {
    return &UserService{
        repo:  repo,
        email: email,
    }
}

// RegisterUser registra un nuevo usuario
func (s *UserService) RegisterUser(name, email string) (*User, error) {
    // Validaciones
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    if email == "" {
        return nil, errors.New("email cannot be empty")
    }
    
    // Verificar si ya existe
    existing, _ := s.repo.FindByEmail(email)
    if existing != nil {
        return nil, errors.New("user already exists")
    }
    
    // Crear usuario
    user := &User{
        Name:  name,
        Email: email,
    }
    
    // Guardar
    if err := s.repo.Save(user); err != nil {
        return nil, fmt.Errorf("failed to save user: %w", err)
    }
    
    // Enviar email de bienvenida
    if err := s.email.SendWelcomeEmail(email); err != nil {
        // Log error pero no fallar
        fmt.Printf("Failed to send welcome email: %v\n", err)
    }
    
    return user, nil
}
```

### 🎭 Mocks Manuales

**`user_service_test.go`**
```go
package user

import (
    "errors"
    "testing"
)

// MockUserRepository implementa UserRepository para testing
type MockUserRepository struct {
    users       map[string]*User
    saveError   error
    findError   error
    nextID      int
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users:  make(map[string]*User),
        nextID: 1,
    }
}

func (m *MockUserRepository) Save(user *User) error {
    if m.saveError != nil {
        return m.saveError
    }
    
    user.ID = m.nextID
    m.nextID++
    m.users[user.Email] = user
    return nil
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
    if m.findError != nil {
        return nil, m.findError
    }
    
    for _, user := range m.users {
        if user.ID == id {
            return user, nil
        }
    }
    return nil, errors.New("user not found")
}

func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
    if m.findError != nil {
        return nil, m.findError
    }
    
    user, exists := m.users[email]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// SetSaveError configura error para Save
func (m *MockUserRepository) SetSaveError(err error) {
    m.saveError = err
}

// MockEmailService implementa EmailService para testing
type MockEmailService struct {
    sentEmails []string
    sendError  error
}

func NewMockEmailService() *MockEmailService {
    return &MockEmailService{
        sentEmails: make([]string, 0),
    }
}

func (m *MockEmailService) SendWelcomeEmail(email string) error {
    if m.sendError != nil {
        return m.sendError
    }
    
    m.sentEmails = append(m.sentEmails, email)
    return nil
}

func (m *MockEmailService) SetSendError(err error) {
    m.sendError = err
}

func (m *MockEmailService) GetSentEmails() []string {
    return m.sentEmails
}

// Tests usando los mocks
func TestUserService_RegisterUser_Success(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockEmail := NewMockEmailService()
    service := NewUserService(mockRepo, mockEmail)
    
    // Act
    user, err := service.RegisterUser("John Doe", "john@example.com")
    
    // Assert
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if user == nil {
        t.Fatal("Expected user, got nil")
    }
    if user.Name != "John Doe" {
        t.Errorf("Expected name 'John Doe', got %s", user.Name)
    }
    if user.Email != "john@example.com" {
        t.Errorf("Expected email 'john@example.com', got %s", user.Email)
    }
    if user.ID == 0 {
        t.Error("Expected ID to be set")
    }
    
    // Verificar que se envió email
    sentEmails := mockEmail.GetSentEmails()
    if len(sentEmails) != 1 {
        t.Errorf("Expected 1 email sent, got %d", len(sentEmails))
    }
    if sentEmails[0] != "john@example.com" {
        t.Errorf("Expected email to john@example.com, got %s", sentEmails[0])
    }
}

func TestUserService_RegisterUser_ValidationErrors(t *testing.T) {
    tests := []struct {
        name          string
        userName      string
        userEmail     string
        expectedError string
    }{
        {"empty name", "", "john@example.com", "name cannot be empty"},
        {"empty email", "John Doe", "", "email cannot be empty"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Arrange
            mockRepo := NewMockUserRepository()
            mockEmail := NewMockEmailService()
            service := NewUserService(mockRepo, mockEmail)
            
            // Act
            user, err := service.RegisterUser(tt.userName, tt.userEmail)
            
            // Assert
            if err == nil {
                t.Error("Expected error, got nil")
            }
            if err.Error() != tt.expectedError {
                t.Errorf("Expected error '%s', got '%s'", tt.expectedError, err.Error())
            }
            if user != nil {
                t.Error("Expected nil user, got user")
            }
        })
    }
}

func TestUserService_RegisterUser_UserAlreadyExists(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockEmail := NewMockEmailService()
    service := NewUserService(mockRepo, mockEmail)
    
    // Crear usuario existente
    existingUser := &User{ID: 1, Name: "Jane Doe", Email: "john@example.com"}
    mockRepo.Save(existingUser)
    
    // Act
    user, err := service.RegisterUser("John Doe", "john@example.com")
    
    // Assert
    if err == nil {
        t.Error("Expected error for existing user")
    }
    if err.Error() != "user already exists" {
        t.Errorf("Expected 'user already exists', got '%s'", err.Error())
    }
    if user != nil {
        t.Error("Expected nil user")
    }
}

func TestUserService_RegisterUser_RepositoryError(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockEmail := NewMockEmailService()
    service := NewUserService(mockRepo, mockEmail)
    
    // Configurar error en repository
    mockRepo.SetSaveError(errors.New("database connection failed"))
    
    // Act
    user, err := service.RegisterUser("John Doe", "john@example.com")
    
    // Assert
    if err == nil {
        t.Error("Expected error from repository")
    }
    if user != nil {
        t.Error("Expected nil user on repository error")
    }
}
```

---

## 🔧 Testing con testify: Assertions Poderosas

### 📦 Instalación de testify

```bash
go get github.com/stretchr/testify
```

### ✅ Usando assert para Tests más Limpios

```go
package user

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

func TestUserService_WithTestify(t *testing.T) {
    // Arrange
    mockRepo := NewMockUserRepository()
    mockEmail := NewMockEmailService()
    service := NewUserService(mockRepo, mockEmail)
    
    // Act
    user, err := service.RegisterUser("John Doe", "john@example.com")
    
    // Assert con testify - más limpio
    require.NoError(t, err)                           // require para checks críticos
    require.NotNil(t, user)
    
    assert.Equal(t, "John Doe", user.Name)           // assert para comparaciones
    assert.Equal(t, "john@example.com", user.Email)
    assert.NotZero(t, user.ID)
    
    // Verificaciones de colecciones
    sentEmails := mockEmail.GetSentEmails()
    assert.Len(t, sentEmails, 1)
    assert.Contains(t, sentEmails, "john@example.com")
}

// Test con múltiples assertions
func TestUserValidation_WithTestify(t *testing.T) {
    user := &User{
        ID:    1,
        Name:  "John Doe",
        Email: "john@example.com",
    }
    
    // Múltiples assertions en una línea
    assert.Equal(t, 1, user.ID)
    assert.NotEmpty(t, user.Name)
    assert.Contains(t, user.Email, "@")
    assert.Regexp(t, `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, user.Email)
}
```

### 🎭 Mocks Avanzados con testify/mock

```go
// MockUserRepository usando testify/mock
type MockUserRepositoryTestify struct {
    mock.Mock
}

func (m *MockUserRepositoryTestify) Save(user *User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepositoryTestify) FindByID(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepositoryTestify) FindByEmail(email string) (*User, error) {
    args := m.Called(email)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

func TestUserService_WithTestifyMocks(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepositoryTestify)
    mockEmail := NewMockEmailService()
    service := NewUserService(mockRepo, mockEmail)
    
    // Configure mock expectations
    mockRepo.On("FindByEmail", "john@example.com").Return(nil, errors.New("not found"))
    mockRepo.On("Save", mock.MatchedBy(func(u *User) bool {
        return u.Name == "John Doe" && u.Email == "john@example.com"
    })).Return(nil)
    
    // Act
    user, err := service.RegisterUser("John Doe", "john@example.com")
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, user)
    
    // Verify all expectations were met
    mockRepo.AssertExpectations(t)
}
```

---

## 🏗️ Testing de Integración

### 🔗 Setup y Teardown

```go
package integration

import (
    "database/sql"
    "os"
    "testing"
    
    _ "github.com/lib/pq"
)

var testDB *sql.DB

// TestMain se ejecuta antes de todos los tests
func TestMain(m *testing.M) {
    // Setup
    var err error
    testDB, err = sql.Open("postgres", "postgres://user:pass@localhost/testdb?sslmode=disable")
    if err != nil {
        panic(err)
    }
    
    // Crear tablas de test
    if err := setupTestTables(testDB); err != nil {
        panic(err)
    }
    
    // Ejecutar tests
    code := m.Run()
    
    // Teardown
    testDB.Close()
    
    // Exit
    os.Exit(code)
}

func setupTestTables(db *sql.DB) error {
    schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL
    );
    `
    _, err := db.Exec(schema)
    return err
}

// Helper para limpiar datos entre tests
func cleanupDatabase(t *testing.T) {
    t.Helper()
    
    tables := []string{"users"}
    for _, table := range tables {
        _, err := testDB.Exec("DELETE FROM " + table)
        if err != nil {
            t.Fatalf("Failed to cleanup table %s: %v", table, err)
        }
    }
}

func TestUserRepository_Integration(t *testing.T) {
    // Cleanup antes del test
    cleanupDatabase(t)
    
    // Test real con base de datos
    repo := NewPostgresUserRepository(testDB)
    
    user := &User{
        Name:  "Integration Test User",
        Email: "integration@test.com",
    }
    
    // Test Save
    err := repo.Save(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
    
    // Test FindByEmail
    found, err := repo.FindByEmail("integration@test.com")
    assert.NoError(t, err)
    assert.Equal(t, user.Name, found.Name)
    assert.Equal(t, user.Email, found.Email)
}
```

### 🐳 Testing con Docker

**`docker-compose.test.yml`**
```yaml
version: '3.8'
services:
  test-db:
    image: postgres:13
    environment:
      POSTGRES_DB: testdb
      POSTGRES_USER: testuser
      POSTGRES_PASSWORD: testpass
    ports:
      - "5433:5432"
    tmpfs:
      - /var/lib/postgresql/data

  test-redis:
    image: redis:6
    ports:
      - "6380:6379"
```

**Script de testing:**
```bash
#!/bin/bash
# test.sh

# Start test containers
docker-compose -f docker-compose.test.yml up -d

# Wait for services
sleep 5

# Run tests
export DATABASE_URL="postgres://testuser:testpass@localhost:5433/testdb?sslmode=disable"
export REDIS_URL="redis://localhost:6380"

go test -v ./...

# Cleanup
docker-compose -f docker-compose.test.yml down
```

---

## 📊 Coverage: Midiendo Cobertura de Tests

### 📈 Generando Reports de Coverage

```bash
# Ejecutar tests con coverage
go test -cover ./...

# Generar archivo de coverage detallado
go test -coverprofile=coverage.out ./...

# Ver coverage por función
go tool cover -func=coverage.out

# Generar reporte HTML
go tool cover -html=coverage.out -o coverage.html

# Coverage de packages específicos
go test -coverprofile=coverage.out -coverpkg=./... ./...
```

### 🎯 Configurando Coverage Goals

```go
// coverage_test.go
package main

import (
    "testing"
    "os"
    "os/exec"
    "strings"
    "strconv"
)

func TestCoverageGoal(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping coverage test in short mode")
    }
    
    // Ejecutar tests con coverage
    cmd := exec.Command("go", "test", "-coverprofile=coverage.out", "./...")
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("Failed to run coverage: %v\nOutput: %s", err, output)
    }
    
    // Analizar coverage
    cmd = exec.Command("go", "tool", "cover", "-func=coverage.out")
    output, err = cmd.Output()
    if err != nil {
        t.Fatalf("Failed to analyze coverage: %v", err)
    }
    
    // Extraer coverage total
    lines := strings.Split(string(output), "\n")
    for _, line := range lines {
        if strings.Contains(line, "total:") {
            parts := strings.Fields(line)
            if len(parts) >= 3 {
                coverageStr := strings.TrimSuffix(parts[2], "%")
                coverage, err := strconv.ParseFloat(coverageStr, 64)
                if err != nil {
                    t.Fatalf("Failed to parse coverage: %v", err)
                }
                
                minimumCoverage := 80.0
                if coverage < minimumCoverage {
                    t.Errorf("Coverage %.1f%% is below minimum %.1f%%", coverage, minimumCoverage)
                }
                
                t.Logf("Coverage: %.1f%%", coverage)
                return
            }
        }
    }
    
    t.Fatal("Could not find total coverage")
}
```

---

## 🎯 Test-Driven Development (TDD)

### 🔄 El Ciclo Red-Green-Refactor

```mermaid
graph LR
    A[🔴 Red<br/>Write failing test] --> B[🟢 Green<br/>Make it pass]
    B --> C[🔵 Refactor<br/>Improve code]
    C --> A
```

### 📝 Ejemplo Práctico: Calculadora TDD

#### 1. **Red** - Test que falla
```go
// calculator_test.go
func TestCalculator_Add(t *testing.T) {
    calc := NewCalculator()
    result := calc.Add(2, 3)
    assert.Equal(t, 5, result)
}
```

#### 2. **Green** - Implementación mínima
```go
// calculator.go
type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Add(a, b int) int {
    return 5 // Hardcoded para pasar el test
}
```

#### 3. **Refactor** - Mejorar implementación
```go
func (c *Calculator) Add(a, b int) int {
    return a + b // Implementación real
}
```

#### 4. **Red** - Más tests
```go
func TestCalculator_AddMultiple(t *testing.T) {
    calc := NewCalculator()
    
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {0, 0, 0},
        {-1, 1, 0},
    }
    
    for _, tt := range tests {
        result := calc.Add(tt.a, tt.b)
        assert.Equal(t, tt.expected, result)
    }
}
```

### 🏗️ TDD para Features Complejas

```go
// Ejemplo: Sistema de descuentos con TDD

// 1. RED - Test para descuento básico
func TestDiscountCalculator_PercentageDiscount(t *testing.T) {
    calc := NewDiscountCalculator()
    discount := calc.CalculateDiscount(100.0, PercentageDiscount{Rate: 0.1})
    assert.Equal(t, 10.0, discount)
}

// 2. GREEN - Implementación mínima
type DiscountCalculator struct{}

type PercentageDiscount struct {
    Rate float64
}

func (d PercentageDiscount) Apply(amount float64) float64 {
    return amount * d.Rate
}

func NewDiscountCalculator() *DiscountCalculator {
    return &DiscountCalculator{}
}

func (c *DiscountCalculator) CalculateDiscount(amount float64, discount PercentageDiscount) float64 {
    return discount.Apply(amount)
}

// 3. RED - Más tipos de descuento
func TestDiscountCalculator_FixedDiscount(t *testing.T) {
    calc := NewDiscountCalculator()
    discount := calc.CalculateDiscount(100.0, FixedDiscount{Amount: 15.0})
    assert.Equal(t, 15.0, discount)
}

// 4. REFACTOR - Interface común
type Discount interface {
    Apply(amount float64) float64
}

type FixedDiscount struct {
    Amount float64
}

func (d FixedDiscount) Apply(amount float64) float64 {
    if d.Amount > amount {
        return amount
    }
    return d.Amount
}

func (c *DiscountCalculator) CalculateDiscount(amount float64, discount Discount) float64 {
    return discount.Apply(amount)
}
```

---

## 🚀 Testing en CI/CD

### 🔧 GitHub Actions para Go

**`.github/workflows/test.yml`**
```yaml
name: Tests

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: testdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.24
    
    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-
    
    - name: Install dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable
    
    - name: Generate coverage report
      run: go tool cover -html=coverage.out -o coverage.html
    
    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
        flags: unittests
        name: codecov-umbrella
    
    - name: Check coverage threshold
      run: |
        COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        echo "Coverage: $COVERAGE%"
        if (( $(echo "$COVERAGE < 80" | bc -l) )); then
          echo "Coverage $COVERAGE% is below minimum 80%"
          exit 1
        fi

  lint:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v3
      with:
        go-version: 1.24
    
    - name: golangci-lint
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

  security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v3
      with:
        go-version: 1.24
    
    - name: Run Gosec Security Scanner
      uses: securecodewarrior/github-action-gosec@master
      with:
        args: './...'
```

---

## 🛠️ Herramientas Avanzadas de Testing

### 🎯 Fuzzing (Go 1.18+)

```go
// fuzz_test.go
package calculator

import (
    "testing"
    "unicode/utf8"
)

func FuzzValidateEmail(f *testing.F) {
    // Seed corpus
    f.Add("test@example.com")
    f.Add("user.name@domain.co.uk")
    f.Add("invalid-email")
    
    f.Fuzz(func(t *testing.T, email string) {
        // El test no debe hacer panic
        result := ValidateEmail(email)
        
        // Propiedades que siempre deben cumplirse
        if !utf8.ValidString(email) {
            t.Skip("Invalid UTF-8")
        }
        
        if result && !strings.Contains(email, "@") {
            t.Errorf("ValidateEmail(%q) = true, but email doesn't contain @", email)
        }
    })
}

// Ejecutar fuzzing
// go test -fuzz=FuzzValidateEmail
```

### 🔍 Property-Based Testing

```go
// property_test.go
package calculator

import (
    "testing"
    "testing/quick"
)

func TestAddCommutative(t *testing.T) {
    f := func(a, b int) bool {
        return Add(a, b) == Add(b, a)
    }
    
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func TestAddAssociative(t *testing.T) {
    f := func(a, b, c int) bool {
        return Add(Add(a, b), c) == Add(a, Add(b, c))
    }
    
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}
```

### 🎭 Contract Testing

```go
// contract_test.go
package api

import (
    "testing"
    "encoding/json"
    "net/http/httptest"
)

// Test que valida contratos de API
func TestUserAPIContract(t *testing.T) {
    router := SetupRouter()
    
    tests := []struct {
        name           string
        method         string
        url           string
        body          string
        expectedCode  int
        expectedShape interface{}
    }{
        {
            name:          "Get Users",
            method:        "GET",
            url:          "/users",
            expectedCode:  200,
            expectedShape: []User{},
        },
        {
            name:          "Create User",
            method:        "POST",
            url:          "/users",
            body:         `{"name":"Test","email":"test@example.com"}`,
            expectedCode:  201,
            expectedShape: User{},
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
            req.Header.Set("Content-Type", "application/json")
            
            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)
            
            assert.Equal(t, tt.expectedCode, w.Code)
            
            // Validate response structure
            var response interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            
            // Validate shape matches expected
            validateResponseShape(t, response, tt.expectedShape)
        })
    }
}
```

---

## 🎯 Mejores Prácticas de Testing

### ✅ **Principios FIRST**

- **🚀 Fast** - Tests rápidos que se ejecuten frecuentemente
- **🔒 Independent** - Tests que no dependan unos de otros
- **🔄 Repeatable** - Mismos resultados en cualquier entorno
- **✅ Self-Validating** - Pass/fail claro, sin interpretación
- **⏰ Timely** - Escritos junto con el código de producción

### 📋 **Estructura AAA**

```go
func TestUserService_RegisterUser(t *testing.T) {
    // Arrange - Configurar test data y mocks
    mockRepo := NewMockUserRepository()
    service := NewUserService(mockRepo)
    userName := "John Doe"
    userEmail := "john@example.com"
    
    // Act - Ejecutar la función bajo test
    user, err := service.RegisterUser(userName, userEmail)
    
    // Assert - Verificar resultados
    assert.NoError(t, err)
    assert.Equal(t, userName, user.Name)
    assert.Equal(t, userEmail, user.Email)
}
```

### 🎯 **Qué Testear vs Qué No Testear**

#### ✅ **TESTEAR:**
- Lógica de negocio
- Casos edge y errores
- Interfaces públicas
- Integraciones críticas
- Algoritmos complejos

#### ❌ **NO TESTEAR:**
- Getters/setters simples
- Frameworks externos
- Configuración trivial
- Código autogenerado

### 🚀 **Performance de Tests**

```go
// Optimizar tests lentos
func TestExpensiveOperation(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping expensive test in short mode")
    }
    
    // Test costoso aquí
}

// Ejecutar solo tests rápidos
// go test -short

// Tests en paralelo para speedup
func TestParallelOperations(t *testing.T) {
    t.Parallel()
    
    tests := []struct{...}{}
    
    for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            // Test code
        })
    }
}
```

---

## 🎉 Proyecto Práctico: Sistema de E-commerce

Vamos a aplicar todo lo aprendido en un proyecto real que incluya:

### 📁 Estructura del Proyecto Testing

```
ecommerce-testing/
├── pkg/
│   ├── models/
│   │   ├── product.go
│   │   └── product_test.go
│   ├── cart/
│   │   ├── cart.go
│   │   ├── cart_test.go
│   │   └── mocks.go
│   └── payment/
│       ├── payment.go
│       ├── payment_test.go
│       └── integration_test.go
├── internal/
│   ├── handlers/
│   │   ├── product_handler.go
│   │   └── product_handler_test.go
│   └── services/
│       ├── order_service.go
│       └── order_service_test.go
├── testdata/
│   ├── products.json
│   └── orders.json
├── integration_test.go
├── coverage.sh
└── Makefile
```

Este proyecto se desarrollará en los ejercicios, aplicando:
- ✅ Tests unitarios con table-driven tests
- ✅ Mocking de dependencias externas
- ✅ Tests de integración con base de datos
- ✅ Benchmarks para operaciones críticas
- ✅ TDD para nuevas features
- ✅ CI/CD con coverage reporting

---

## 📚 Recursos para Continuar

### 🎯 Próximos Pasos

1. **✅ Completar los ejercicios** de esta lección
2. **📖 Explorar testing avanzado** con microservicios
3. **🔗 Implementar** contract testing
4. **🚀 Configurar** CI/CD pipeline completo

### 🌐 Enlaces Útiles

- **[Testing Package Docs](https://golang.org/pkg/testing/)** - Documentación oficial
- **[Testify](https://github.com/stretchr/testify)** - Testing toolkit
- **[GoMock](https://github.com/golang/mock)** - Mock generation
- **[Ginkgo](https://github.com/onsi/ginkgo)** - BDD testing framework

---

## 🧠 Resumen de Conceptos Clave

```mermaid
mindmap
  root((Testing en Go))
    Tests Unitarios
      Table-driven tests
      Subtests
      Parallel execution
    Mocking
      Manual mocks
      Testify mocks
      Dependency injection
    Integration
      Database testing
      HTTP testing
      Docker containers
    Advanced
      Benchmarks
      Fuzzing
      Property-based
    CI/CD
      GitHub Actions
      Coverage reporting
      Quality gates
```

¡Felicitaciones! Ahora dominas el testing en Go desde lo básico hasta técnicas avanzadas. El testing es fundamental para código confiable y mantenible.

**[→ Continúa con los Ejercicios de Testing](./ejercicios.go)**
