# 🧪 Lección 17: Testing Avanzado - TDD, Mocking y Property Testing

## 🎯 Objetivos de la Lección

Al finalizar esta lección, serás capaz de:
- Dominar Test-Driven Development (TDD) en Go
- Crear y usar mocks, stubs y test doubles efectivamente
- Implementar property-based testing con invariantes
- Escribir tests de integración robustos
- Aplicar técnicas de benchmarking y profiling
- Diseñar test suites escalables y mantenibles
- Usar testify y otras herramientas de testing avanzadas
- Implementar testing de código concurrente

---

## 🧠 Analogía: Testing como un Laboratorio de Calidad

Imagina que el **testing** en Go es como un **laboratorio de control de calidad**:

```
🔬 Laboratorio de Calidad (Test Suite)
├── 🧪 Pruebas Unitarias (Unit Tests)
│   ├── ⚗️ Tests Básicos (Simple Function Tests)
│   ├── 🔍 Tests de Comportamiento (Behavior Tests)
│   └── 🎭 Tests con Mocks (Isolated Tests)
├── 🏭 Pruebas de Integración (Integration Tests)
│   ├── 🔗 Componente a Componente
│   ├── 🌐 Sistema Completo
│   └── 📡 APIs Externas
├── 🎯 Property Testing (Property-Based Tests)
│   ├── 📏 Invariantes Matemáticas
│   ├── 🔄 Propiedades de Reversibilidad
│   └── 🛡️ Propiedades de Seguridad
├── ⚡ Benchmarks (Performance Tests)
│   ├── 📊 Tiempo de Ejecución
│   ├── 💾 Uso de Memoria
│   └── 🔄 Throughput
└── 🎪 Test Doubles (Simulaciones)
    ├── 🎭 Mocks (Verificación de Comportamiento)
    ├── 📝 Stubs (Respuestas Predefinidas)
    └── 🎨 Fakes (Implementaciones Simplificadas)
```

El **Testing Avanzado** en Go:
- **Verifica calidad** sistemáticamente
- **Previene defectos** antes de producción
- **Documenta comportamiento** esperado
- **Facilita refactoring** seguro

---

## 📚 Fundamentos del Testing en Go

### 🔧 El Ecosistema de Testing

Go incluye un ecosistema de testing robusto y extensible:

```go
package main

import (
    "fmt"
    "testing"
    "time"
)

// Tipos básicos de tests en Go
func ejemploEcosistemaTesting() {
    fmt.Println("🔧 Ecosistema de Testing en Go")
    fmt.Println("=============================")
    
    fmt.Println("📦 Built-in Testing:")
    fmt.Println("  • testing.T - Unit tests")
    fmt.Println("  • testing.B - Benchmarks")
    fmt.Println("  • testing.M - Test main")
    fmt.Println("  • testing.TB - Interface común")
    
    fmt.Println("\n🛠️ Herramientas Estándar:")
    fmt.Println("  • go test - Ejecutor de tests")
    fmt.Println("  • go test -bench - Benchmarks")
    fmt.Println("  • go test -cover - Cobertura")
    fmt.Println("  • go test -race - Race detection")
    
    fmt.Println("\n📚 Librerías Populares:")
    fmt.Println("  • testify/assert - Assertions")
    fmt.Println("  • testify/mock - Mocking")
    fmt.Println("  • testify/suite - Test suites")
    fmt.Println("  • gomock - Code generation mocks")
    fmt.Println("  • ginkgo/gomega - BDD testing")
}

// Test básico de ejemplo
func TestExample(t *testing.T) {
    // Arrange
    input := "hello"
    expected := "HELLO"
    
    // Act
    result := toUpper(input)
    
    // Assert
    if result != expected {
        t.Errorf("toUpper(%q) = %q; want %q", input, result, expected)
    }
}

func toUpper(s string) string {
    return fmt.Sprintf("%s", s) // Implementación simple para ejemplo
}
```

### 📊 Anatomía de un Test Completo

```go
package main

import (
    "testing"
    "time"
)

// Test completo con setup, teardown y subtests
func TestCompleteExample(t *testing.T) {
    // Setup - preparar el entorno de testing
    setup := func() *TestEnvironment {
        return &TestEnvironment{
            Database: setupTestDB(),
            Cache:    setupTestCache(),
            Logger:   setupTestLogger(),
        }
    }
    
    // Teardown - limpiar después del test
    teardown := func(env *TestEnvironment) {
        env.Database.Close()
        env.Cache.Clear()
        env.Logger.Close()
    }
    
    t.Run("UserService", func(t *testing.T) {
        env := setup()
        defer teardown(env)
        
        userService := NewUserService(env.Database, env.Cache)
        
        t.Run("CreateUser", func(t *testing.T) {
            // Given
            userData := UserData{
                Email:    "test@example.com",
                Username: "testuser",
                Age:      25,
            }
            
            // When
            user, err := userService.CreateUser(userData)
            
            // Then
            if err != nil {
                t.Fatalf("CreateUser failed: %v", err)
            }
            
            if user.Email != userData.Email {
                t.Errorf("Expected email %q, got %q", userData.Email, user.Email)
            }
            
            if user.ID == "" {
                t.Error("Expected user ID to be generated")
            }
        })
        
        t.Run("GetUser", func(t *testing.T) {
            // Test de obtener usuario
            // ... implementación
        })
        
        t.Run("UpdateUser", func(t *testing.T) {
            // Test de actualizar usuario
            // ... implementación
        })
    })
}

// Estructuras de apoyo para testing
type TestEnvironment struct {
    Database TestDB
    Cache    TestCache
    Logger   TestLogger
}

type UserData struct {
    Email    string
    Username string
    Age      int
}

type User struct {
    ID       string
    Email    string
    Username string
    Age      int
}

// Stubs para el ejemplo
type TestDB struct{}
func (db TestDB) Close() {}
func setupTestDB() TestDB { return TestDB{} }

type TestCache struct{}
func (c TestCache) Clear() {}
func setupTestCache() TestCache { return TestCache{} }

type TestLogger struct{}
func (l TestLogger) Close() {}
func setupTestLogger() TestLogger { return TestLogger{} }

type UserService struct{}
func NewUserService(db TestDB, cache TestCache) *UserService { return &UserService{} }
func (us *UserService) CreateUser(data UserData) (*User, error) {
    return &User{
        ID:       "test-id",
        Email:    data.Email,
        Username: data.Username,
        Age:      data.Age,
    }, nil
}
```

---

## 🎯 Test-Driven Development (TDD)

### 1. 🔄 El Ciclo Red-Green-Refactor

```go
package calculator

import (
    "errors"
    "testing"
)

// PASO 1: RED - Escribir test que falle
func TestCalculator_Add(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 5, 5},
        {"mixed", -3, 7, 4},
    }
    
    calc := NewCalculator()
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calc.Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                        tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// PASO 2: GREEN - Implementación mínima que pase
type Calculator struct{}

func NewCalculator() *Calculator {
    return &Calculator{}
}

func (c *Calculator) Add(a, b int) int {
    return a + b
}

// PASO 3: REFACTOR - Mejorar sin romper tests
func TestCalculator_Divide(t *testing.T) {
    calc := NewCalculator()
    
    t.Run("valid division", func(t *testing.T) {
        result, err := calc.Divide(10, 2)
        if err != nil {
            t.Fatalf("Unexpected error: %v", err)
        }
        if result != 5.0 {
            t.Errorf("Divide(10, 2) = %f; want 5.0", result)
        }
    })
    
    t.Run("division by zero", func(t *testing.T) {
        _, err := calc.Divide(10, 0)
        if err == nil {
            t.Fatal("Expected error for division by zero")
        }
        
        expectedErr := "division by zero"
        if err.Error() != expectedErr {
            t.Errorf("Expected error %q, got %q", expectedErr, err.Error())
        }
    })
}

func (c *Calculator) Divide(a, b int) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return float64(a) / float64(b), nil
}
```

### 2. 🏗️ TDD para Diseño de APIs

```go
package bankaccount

import (
    "errors"
    "testing"
    "time"
)

// Diseñando una API de cuenta bancaria con TDD

// PASO 1: Definir comportamiento esperado a través de tests
func TestBankAccount_Creation(t *testing.T) {
    t.Run("new account should have zero balance", func(t *testing.T) {
        account := NewBankAccount("ACC001", "John Doe")
        
        if account.GetBalance() != 0 {
            t.Errorf("New account balance should be 0, got %f", account.GetBalance())
        }
        
        if account.GetAccountNumber() != "ACC001" {
            t.Errorf("Expected account number ACC001, got %s", account.GetAccountNumber())
        }
        
        if account.GetOwner() != "John Doe" {
            t.Errorf("Expected owner John Doe, got %s", account.GetOwner())
        }
    })
}

func TestBankAccount_Deposit(t *testing.T) {
    account := NewBankAccount("ACC001", "John Doe")
    
    t.Run("valid deposit", func(t *testing.T) {
        err := account.Deposit(100.0)
        if err != nil {
            t.Fatalf("Unexpected error: %v", err)
        }
        
        if account.GetBalance() != 100.0 {
            t.Errorf("Expected balance 100.0, got %f", account.GetBalance())
        }
    })
    
    t.Run("negative deposit should fail", func(t *testing.T) {
        err := account.Deposit(-50.0)
        if err == nil {
            t.Fatal("Expected error for negative deposit")
        }
        
        if account.GetBalance() != 100.0 {
            t.Errorf("Balance should remain 100.0, got %f", account.GetBalance())
        }
    })
}

func TestBankAccount_Withdraw(t *testing.T) {
    account := NewBankAccount("ACC001", "John Doe")
    account.Deposit(100.0)
    
    t.Run("valid withdrawal", func(t *testing.T) {
        err := account.Withdraw(30.0)
        if err != nil {
            t.Fatalf("Unexpected error: %v", err)
        }
        
        if account.GetBalance() != 70.0 {
            t.Errorf("Expected balance 70.0, got %f", account.GetBalance())
        }
    })
    
    t.Run("withdrawal exceeding balance should fail", func(t *testing.T) {
        err := account.Withdraw(100.0)
        if err == nil {
            t.Fatal("Expected error for insufficient funds")
        }
        
        if account.GetBalance() != 70.0 {
            t.Errorf("Balance should remain 70.0, got %f", account.GetBalance())
        }
    })
}

func TestBankAccount_TransactionHistory(t *testing.T) {
    account := NewBankAccount("ACC001", "John Doe")
    
    // Realizar varias operaciones
    account.Deposit(100.0)
    account.Withdraw(30.0)
    account.Deposit(50.0)
    
    transactions := account.GetTransactionHistory()
    
    if len(transactions) != 3 {
        t.Errorf("Expected 3 transactions, got %d", len(transactions))
    }
    
    // Verificar primera transacción
    if transactions[0].Type != "DEPOSIT" {
        t.Errorf("Expected first transaction to be DEPOSIT, got %s", transactions[0].Type)
    }
    
    if transactions[0].Amount != 100.0 {
        t.Errorf("Expected first transaction amount 100.0, got %f", transactions[0].Amount)
    }
}

// PASO 2: Implementación basada en los tests
type BankAccount struct {
    accountNumber string
    owner         string
    balance       float64
    transactions  []Transaction
}

type Transaction struct {
    Type      string
    Amount    float64
    Timestamp time.Time
    Balance   float64
}

func NewBankAccount(accountNumber, owner string) *BankAccount {
    return &BankAccount{
        accountNumber: accountNumber,
        owner:         owner,
        balance:       0,
        transactions:  make([]Transaction, 0),
    }
}

func (ba *BankAccount) GetBalance() float64 {
    return ba.balance
}

func (ba *BankAccount) GetAccountNumber() string {
    return ba.accountNumber
}

func (ba *BankAccount) GetOwner() string {
    return ba.owner
}

func (ba *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    
    ba.balance += amount
    ba.addTransaction("DEPOSIT", amount)
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    
    if amount > ba.balance {
        return errors.New("insufficient funds")
    }
    
    ba.balance -= amount
    ba.addTransaction("WITHDRAWAL", amount)
    return nil
}

func (ba *BankAccount) GetTransactionHistory() []Transaction {
    // Retornar copia para evitar modificaciones externas
    result := make([]Transaction, len(ba.transactions))
    copy(result, ba.transactions)
    return result
}

func (ba *BankAccount) addTransaction(transactionType string, amount float64) {
    transaction := Transaction{
        Type:      transactionType,
        Amount:    amount,
        Timestamp: time.Now(),
        Balance:   ba.balance,
    }
    ba.transactions = append(ba.transactions, transaction)
}
```

---

## 🎭 Mocking y Test Doubles

### 1. 📝 Stubs - Respuestas Predefinidas

```go
package userservice

import (
    "errors"
    "testing"
)

// Interface que queremos mockear
type EmailService interface {
    SendEmail(to, subject, body string) error
    ValidateEmail(email string) bool
}

type UserRepository interface {
    GetUserByEmail(email string) (*User, error)
    SaveUser(user *User) error
}

// Implementación del servicio que usa dependencias
type UserService struct {
    emailService EmailService
    userRepo     UserRepository
}

func NewUserService(emailService EmailService, userRepo UserRepository) *UserService {
    return &UserService{
        emailService: emailService,
        userRepo:     userRepo,
    }
}

func (us *UserService) RegisterUser(email, name string) error {
    // Validar email
    if !us.emailService.ValidateEmail(email) {
        return errors.New("invalid email format")
    }
    
    // Verificar si usuario ya existe
    existingUser, _ := us.userRepo.GetUserByEmail(email)
    if existingUser != nil {
        return errors.New("user already exists")
    }
    
    // Crear nuevo usuario
    user := &User{
        Email: email,
        Name:  name,
    }
    
    // Guardar usuario
    if err := us.userRepo.SaveUser(user); err != nil {
        return err
    }
    
    // Enviar email de bienvenida
    if err := us.emailService.SendEmail(email, "Welcome!", "Welcome to our service!"); err != nil {
        return err
    }
    
    return nil
}

// STUBS - Implementaciones simples para testing
type EmailServiceStub struct {
    ValidateEmailResult bool
    SendEmailError      error
    SendEmailCalls      []EmailCall
}

type EmailCall struct {
    To      string
    Subject string
    Body    string
}

func (e *EmailServiceStub) SendEmail(to, subject, body string) error {
    e.SendEmailCalls = append(e.SendEmailCalls, EmailCall{
        To:      to,
        Subject: subject,
        Body:    body,
    })
    return e.SendEmailError
}

func (e *EmailServiceStub) ValidateEmail(email string) bool {
    return e.ValidateEmailResult
}

type UserRepositoryStub struct {
    GetUserResult *User
    GetUserError  error
    SaveUserError error
    SavedUsers    []*User
}

func (ur *UserRepositoryStub) GetUserByEmail(email string) (*User, error) {
    return ur.GetUserResult, ur.GetUserError
}

func (ur *UserRepositoryStub) SaveUser(user *User) error {
    if ur.SaveUserError == nil {
        ur.SavedUsers = append(ur.SavedUsers, user)
    }
    return ur.SaveUserError
}

// Tests usando stubs
func TestUserService_RegisterUser(t *testing.T) {
    t.Run("successful registration", func(t *testing.T) {
        // Arrange
        emailStub := &EmailServiceStub{
            ValidateEmailResult: true,
            SendEmailError:      nil,
        }
        userRepoStub := &UserRepositoryStub{
            GetUserResult: nil, // Usuario no existe
            SaveUserError: nil,
        }
        
        userService := NewUserService(emailStub, userRepoStub)
        
        // Act
        err := userService.RegisterUser("test@example.com", "Test User")
        
        // Assert
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
        
        // Verificar que se guardó el usuario
        if len(userRepoStub.SavedUsers) != 1 {
            t.Errorf("Expected 1 saved user, got %d", len(userRepoStub.SavedUsers))
        }
        
        // Verificar que se envió el email
        if len(emailStub.SendEmailCalls) != 1 {
            t.Errorf("Expected 1 email sent, got %d", len(emailStub.SendEmailCalls))
        }
        
        emailCall := emailStub.SendEmailCalls[0]
        if emailCall.To != "test@example.com" {
            t.Errorf("Expected email to test@example.com, got %s", emailCall.To)
        }
    })
    
    t.Run("invalid email format", func(t *testing.T) {
        // Arrange
        emailStub := &EmailServiceStub{
            ValidateEmailResult: false, // Email inválido
        }
        userRepoStub := &UserRepositoryStub{}
        
        userService := NewUserService(emailStub, userRepoStub)
        
        // Act
        err := userService.RegisterUser("invalid-email", "Test User")
        
        // Assert
        if err == nil {
            t.Fatal("Expected error for invalid email")
        }
        
        if err.Error() != "invalid email format" {
            t.Errorf("Expected 'invalid email format', got %s", err.Error())
        }
        
        // Verificar que no se guardó usuario ni se envió email
        if len(userRepoStub.SavedUsers) != 0 {
            t.Errorf("Expected 0 saved users, got %d", len(userRepoStub.SavedUsers))
        }
        
        if len(emailStub.SendEmailCalls) != 0 {
            t.Errorf("Expected 0 emails sent, got %d", len(emailStub.SendEmailCalls))
        }
    })
    
    t.Run("user already exists", func(t *testing.T) {
        // Arrange
        existingUser := &User{Email: "test@example.com", Name: "Existing User"}
        
        emailStub := &EmailServiceStub{
            ValidateEmailResult: true,
        }
        userRepoStub := &UserRepositoryStub{
            GetUserResult: existingUser, // Usuario ya existe
        }
        
        userService := NewUserService(emailStub, userRepoStub)
        
        // Act
        err := userService.RegisterUser("test@example.com", "Test User")
        
        // Assert
        if err == nil {
            t.Fatal("Expected error for existing user")
        }
        
        if err.Error() != "user already exists" {
            t.Errorf("Expected 'user already exists', got %s", err.Error())
        }
    })
}

// Estructura User para el ejemplo
type User struct {
    Email string
    Name  string
}
```

### 2. 🎭 Mocks con Verificación de Comportamiento

```go
package paymentservice

import (
    "errors"
    "testing"
)

// Interface para external payment processor
type PaymentProcessor interface {
    ProcessPayment(amount float64, cardToken string) (*PaymentResult, error)
    RefundPayment(paymentID string, amount float64) (*RefundResult, error)
}

type PaymentResult struct {
    PaymentID string
    Status    string
    Amount    float64
}

type RefundResult struct {
    RefundID  string
    Status    string
    Amount    float64
}

// Mock con verificación de comportamiento
type PaymentProcessorMock struct {
    // Expectativas
    ExpectedProcessPaymentCalls []ProcessPaymentCall
    ExpectedRefundPaymentCalls  []RefundPaymentCall
    
    // Respuestas configuradas
    ProcessPaymentResults []ProcessPaymentResult
    RefundPaymentResults  []RefundPaymentResult
    
    // Llamadas reales
    ActualProcessPaymentCalls []ProcessPaymentCall
    ActualRefundPaymentCalls  []RefundPaymentCall
    
    // Contadores
    ProcessPaymentCallCount int
    RefundPaymentCallCount  int
}

type ProcessPaymentCall struct {
    Amount    float64
    CardToken string
}

type ProcessPaymentResult struct {
    Result *PaymentResult
    Error  error
}

type RefundPaymentCall struct {
    PaymentID string
    Amount    float64
}

type RefundPaymentResult struct {
    Result *RefundResult
    Error  error
}

func NewPaymentProcessorMock() *PaymentProcessorMock {
    return &PaymentProcessorMock{
        ExpectedProcessPaymentCalls: make([]ProcessPaymentCall, 0),
        ExpectedRefundPaymentCalls:  make([]RefundPaymentCall, 0),
        ProcessPaymentResults:       make([]ProcessPaymentResult, 0),
        RefundPaymentResults:        make([]RefundPaymentResult, 0),
        ActualProcessPaymentCalls:   make([]ProcessPaymentCall, 0),
        ActualRefundPaymentCalls:    make([]RefundPaymentCall, 0),
    }
}

// Configurar expectativas
func (m *PaymentProcessorMock) ExpectProcessPayment(amount float64, cardToken string) *PaymentProcessorMock {
    m.ExpectedProcessPaymentCalls = append(m.ExpectedProcessPaymentCalls, ProcessPaymentCall{
        Amount:    amount,
        CardToken: cardToken,
    })
    return m
}

func (m *PaymentProcessorMock) WillReturnProcessPayment(result *PaymentResult, err error) *PaymentProcessorMock {
    m.ProcessPaymentResults = append(m.ProcessPaymentResults, ProcessPaymentResult{
        Result: result,
        Error:  err,
    })
    return m
}

func (m *PaymentProcessorMock) ExpectRefundPayment(paymentID string, amount float64) *PaymentProcessorMock {
    m.ExpectedRefundPaymentCalls = append(m.ExpectedRefundPaymentCalls, RefundPaymentCall{
        PaymentID: paymentID,
        Amount:    amount,
    })
    return m
}

func (m *PaymentProcessorMock) WillReturnRefundPayment(result *RefundResult, err error) *PaymentProcessorMock {
    m.RefundPaymentResults = append(m.RefundPaymentResults, RefundPaymentResult{
        Result: result,
        Error:  err,
    })
    return m
}

// Implementar interface
func (m *PaymentProcessorMock) ProcessPayment(amount float64, cardToken string) (*PaymentResult, error) {
    m.ActualProcessPaymentCalls = append(m.ActualProcessPaymentCalls, ProcessPaymentCall{
        Amount:    amount,
        CardToken: cardToken,
    })
    
    if m.ProcessPaymentCallCount < len(m.ProcessPaymentResults) {
        result := m.ProcessPaymentResults[m.ProcessPaymentCallCount]
        m.ProcessPaymentCallCount++
        return result.Result, result.Error
    }
    
    return nil, errors.New("unexpected call to ProcessPayment")
}

func (m *PaymentProcessorMock) RefundPayment(paymentID string, amount float64) (*RefundResult, error) {
    m.ActualRefundPaymentCalls = append(m.ActualRefundPaymentCalls, RefundPaymentCall{
        PaymentID: paymentID,
        Amount:    amount,
    })
    
    if m.RefundPaymentCallCount < len(m.RefundPaymentResults) {
        result := m.RefundPaymentResults[m.RefundPaymentCallCount]
        m.RefundPaymentCallCount++
        return result.Result, result.Error
    }
    
    return nil, errors.New("unexpected call to RefundPayment")
}

// Verificar expectativas
func (m *PaymentProcessorMock) VerifyExpectations(t *testing.T) {
    // Verificar número de llamadas
    if len(m.ActualProcessPaymentCalls) != len(m.ExpectedProcessPaymentCalls) {
        t.Errorf("Expected %d ProcessPayment calls, got %d",
            len(m.ExpectedProcessPaymentCalls), len(m.ActualProcessPaymentCalls))
    }
    
    // Verificar parámetros de llamadas ProcessPayment
    for i, expected := range m.ExpectedProcessPaymentCalls {
        if i >= len(m.ActualProcessPaymentCalls) {
            t.Errorf("Missing ProcessPayment call %d", i)
            continue
        }
        
        actual := m.ActualProcessPaymentCalls[i]
        if actual.Amount != expected.Amount {
            t.Errorf("ProcessPayment call %d: expected amount %f, got %f",
                i, expected.Amount, actual.Amount)
        }
        
        if actual.CardToken != expected.CardToken {
            t.Errorf("ProcessPayment call %d: expected card token %s, got %s",
                i, expected.CardToken, actual.CardToken)
        }
    }
    
    // Verificar llamadas RefundPayment
    if len(m.ActualRefundPaymentCalls) != len(m.ExpectedRefundPaymentCalls) {
        t.Errorf("Expected %d RefundPayment calls, got %d",
            len(m.ExpectedRefundPaymentCalls), len(m.ActualRefundPaymentCalls))
    }
    
    for i, expected := range m.ExpectedRefundPaymentCalls {
        if i >= len(m.ActualRefundPaymentCalls) {
            t.Errorf("Missing RefundPayment call %d", i)
            continue
        }
        
        actual := m.ActualRefundPaymentCalls[i]
        if actual.PaymentID != expected.PaymentID {
            t.Errorf("RefundPayment call %d: expected payment ID %s, got %s",
                i, expected.PaymentID, actual.PaymentID)
        }
        
        if actual.Amount != expected.Amount {
            t.Errorf("RefundPayment call %d: expected amount %f, got %f",
                i, expected.Amount, actual.Amount)
        }
    }
}

// Servicio que usa el payment processor
type OrderService struct {
    paymentProcessor PaymentProcessor
}

func NewOrderService(paymentProcessor PaymentProcessor) *OrderService {
    return &OrderService{
        paymentProcessor: paymentProcessor,
    }
}

func (os *OrderService) ProcessOrder(amount float64, cardToken string) (string, error) {
    result, err := os.paymentProcessor.ProcessPayment(amount, cardToken)
    if err != nil {
        return "", err
    }
    
    if result.Status != "SUCCESS" {
        return "", errors.New("payment failed")
    }
    
    return result.PaymentID, nil
}

func (os *OrderService) RefundOrder(paymentID string, amount float64) error {
    result, err := os.paymentProcessor.RefundPayment(paymentID, amount)
    if err != nil {
        return err
    }
    
    if result.Status != "SUCCESS" {
        return errors.New("refund failed")
    }
    
    return nil
}

// Tests usando mocks con verificación
func TestOrderService_ProcessOrder(t *testing.T) {
    t.Run("successful payment", func(t *testing.T) {
        // Arrange
        mockProcessor := NewPaymentProcessorMock()
        mockProcessor.
            ExpectProcessPayment(100.0, "card-token-123").
            WillReturnProcessPayment(&PaymentResult{
                PaymentID: "payment-123",
                Status:    "SUCCESS",
                Amount:    100.0,
            }, nil)
        
        orderService := NewOrderService(mockProcessor)
        
        // Act
        paymentID, err := orderService.ProcessOrder(100.0, "card-token-123")
        
        // Assert
        if err != nil {
            t.Fatalf("Unexpected error: %v", err)
        }
        
        if paymentID != "payment-123" {
            t.Errorf("Expected payment ID payment-123, got %s", paymentID)
        }
        
        // Verificar que se llamó al mock correctamente
        mockProcessor.VerifyExpectations(t)
    })
    
    t.Run("payment processor error", func(t *testing.T) {
        // Arrange
        mockProcessor := NewPaymentProcessorMock()
        mockProcessor.
            ExpectProcessPayment(100.0, "invalid-token").
            WillReturnProcessPayment(nil, errors.New("invalid card token"))
        
        orderService := NewOrderService(mockProcessor)
        
        // Act
        _, err := orderService.ProcessOrder(100.0, "invalid-token")
        
        // Assert
        if err == nil {
            t.Fatal("Expected error but got none")
        }
        
        if err.Error() != "invalid card token" {
            t.Errorf("Expected 'invalid card token', got %s", err.Error())
        }
        
        mockProcessor.VerifyExpectations(t)
    })
}
```

---

## 🎯 Property-Based Testing

### 1. 📏 Testing de Invariantes Matemáticas

```go
package mathutils

import (
    "math"
    "math/rand"
    "testing"
    "time"
)

// Función para testear
func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// Property-based testing manual
func TestAdd_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: Commutatividad (a + b = b + a)
    t.Run("commutativity", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            a := rand.Intn(1000) - 500 // Números entre -500 y 500
            b := rand.Intn(1000) - 500
            
            result1 := Add(a, b)
            result2 := Add(b, a)
            
            if result1 != result2 {
                t.Errorf("Commutativity failed: Add(%d, %d) = %d, Add(%d, %d) = %d",
                    a, b, result1, b, a, result2)
            }
        }
    })
    
    // Propiedad: Asociatividad ((a + b) + c = a + (b + c))
    t.Run("associativity", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            a := rand.Intn(200) - 100
            b := rand.Intn(200) - 100
            c := rand.Intn(200) - 100
            
            result1 := Add(Add(a, b), c)
            result2 := Add(a, Add(b, c))
            
            if result1 != result2 {
                t.Errorf("Associativity failed: Add(Add(%d, %d), %d) = %d, Add(%d, Add(%d, %d)) = %d",
                    a, b, c, result1, a, b, c, result2)
            }
        }
    })
    
    // Propiedad: Elemento identidad (a + 0 = a)
    t.Run("identity element", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            a := rand.Intn(1000) - 500
            
            result := Add(a, 0)
            
            if result != a {
                t.Errorf("Identity failed: Add(%d, 0) = %d, expected %d", a, result, a)
            }
        }
    })
}

func TestMultiply_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: Distributividad (a * (b + c) = a * b + a * c)
    t.Run("distributivity", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            a := rand.Intn(50) - 25 // Números más pequeños para evitar overflow
            b := rand.Intn(50) - 25
            c := rand.Intn(50) - 25
            
            result1 := Multiply(a, Add(b, c))
            result2 := Add(Multiply(a, b), Multiply(a, c))
            
            if result1 != result2 {
                t.Errorf("Distributivity failed: %d * (%d + %d) = %d, %d * %d + %d * %d = %d",
                    a, b, c, result1, a, b, a, c, result2)
            }
        }
    })
    
    // Propiedad: Elemento absorvente (a * 0 = 0)
    t.Run("zero element", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            a := rand.Intn(1000) - 500
            
            result := Multiply(a, 0)
            
            if result != 0 {
                t.Errorf("Zero element failed: Multiply(%d, 0) = %d, expected 0", a, result)
            }
        }
    })
}

func TestAbs_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: |x| >= 0 para todo x
    t.Run("non-negative", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            x := rand.Intn(2000) - 1000
            
            result := Abs(x)
            
            if result < 0 {
                t.Errorf("Non-negative failed: Abs(%d) = %d, expected >= 0", x, result)
            }
        }
    })
    
    // Propiedad: |x| = |-x|
    t.Run("symmetry", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            x := rand.Intn(1000) - 500
            
            result1 := Abs(x)
            result2 := Abs(-x)
            
            if result1 != result2 {
                t.Errorf("Symmetry failed: Abs(%d) = %d, Abs(%d) = %d", x, result1, -x, result2)
            }
        }
    })
    
    // Propiedad: |x| = x si x >= 0
    t.Run("identity for non-negative", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            x := rand.Intn(1000) // Solo números positivos
            
            result := Abs(x)
            
            if result != x {
                t.Errorf("Identity failed: Abs(%d) = %d, expected %d", x, result, x)
            }
        }
    })
}

func TestReverse_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: Reverse(Reverse(s)) = s
    t.Run("double reverse", func(t *testing.T) {
        testStrings := generateRandomStrings(100)
        
        for _, s := range testStrings {
            result := Reverse(Reverse(s))
            
            if result != s {
                t.Errorf("Double reverse failed: Reverse(Reverse(%q)) = %q, expected %q",
                    s, result, s)
            }
        }
    })
    
    // Propiedad: len(Reverse(s)) = len(s)
    t.Run("length preservation", func(t *testing.T) {
        testStrings := generateRandomStrings(100)
        
        for _, s := range testStrings {
            result := Reverse(s)
            
            if len(result) != len(s) {
                t.Errorf("Length not preserved: len(Reverse(%q)) = %d, expected %d",
                    s, len(result), len(s))
            }
        }
    })
    
    // Propiedad: Reverse("") = ""
    t.Run("empty string", func(t *testing.T) {
        result := Reverse("")
        if result != "" {
            t.Errorf("Empty string failed: Reverse(\"\") = %q, expected \"\"", result)
        }
    })
}

// Helper para generar strings aleatorios
func generateRandomStrings(count int) []string {
    strings := make([]string, count)
    chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    
    for i := 0; i < count; i++ {
        length := rand.Intn(20) + 1 // Strings de 1 a 20 caracteres
        bytes := make([]byte, length)
        
        for j := range bytes {
            bytes[j] = chars[rand.Intn(len(chars))]
        }
        
        strings[i] = string(bytes)
    }
    
    return strings
}
```

### 2. 🔄 Property Testing para Estructuras de Datos

```go
package datastructures

import (
    "math/rand"
    "sort"
    "testing"
    "time"
)

// Stack implementation
type Stack struct {
    items []int
}

func NewStack() *Stack {
    return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}

func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack) Size() int {
    return len(s.items)
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

// Property-based testing para Stack
func TestStack_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: Push luego Pop retorna el mismo elemento
    t.Run("push-pop invariant", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            stack := NewStack()
            value := rand.Intn(1000)
            
            stack.Push(value)
            result, ok := stack.Pop()
            
            if !ok {
                t.Errorf("Pop failed after Push(%d)", value)
                continue
            }
            
            if result != value {
                t.Errorf("Push-Pop invariant failed: Push(%d), Pop() = %d", value, result)
            }
        }
    })
    
    // Propiedad: Size aumenta en 1 después de Push
    t.Run("push increases size", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            stack := NewStack()
            
            // Crear stack con tamaño aleatorio
            initialSize := rand.Intn(50)
            for j := 0; j < initialSize; j++ {
                stack.Push(rand.Intn(100))
            }
            
            sizeBefore := stack.Size()
            stack.Push(rand.Intn(100))
            sizeAfter := stack.Size()
            
            if sizeAfter != sizeBefore+1 {
                t.Errorf("Size invariant failed: size before %d, size after %d", sizeBefore, sizeAfter)
            }
        }
    })
    
    // Propiedad: Pop en stack vacío retorna false
    t.Run("pop empty stack", func(t *testing.T) {
        for i := 0; i < 50; i++ {
            stack := NewStack()
            _, ok := stack.Pop()
            
            if ok {
                t.Error("Pop on empty stack should return false")
            }
        }
    })
    
    // Propiedad: LIFO (Last In, First Out)
    t.Run("LIFO property", func(t *testing.T) {
        for i := 0; i < 50; i++ {
            stack := NewStack()
            values := generateRandomInts(rand.Intn(20) + 1)
            
            // Push todos los valores
            for _, v := range values {
                stack.Push(v)
            }
            
            // Pop todos los valores y verificar orden
            for j := len(values) - 1; j >= 0; j-- {
                result, ok := stack.Pop()
                if !ok {
                    t.Errorf("Pop failed at index %d", j)
                    break
                }
                
                if result != values[j] {
                    t.Errorf("LIFO violated: expected %d, got %d at position %d", values[j], result, j)
                }
            }
        }
    })
}

// Queue implementation
type Queue struct {
    items []int
}

func NewQueue() *Queue {
    return &Queue{items: make([]int, 0)}
}

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func (q *Queue) Size() int {
    return len(q.items)
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

// Property-based testing para Queue
func TestQueue_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: FIFO (First In, First Out)
    t.Run("FIFO property", func(t *testing.T) {
        for i := 0; i < 50; i++ {
            queue := NewQueue()
            values := generateRandomInts(rand.Intn(20) + 1)
            
            // Enqueue todos los valores
            for _, v := range values {
                queue.Enqueue(v)
            }
            
            // Dequeue todos los valores y verificar orden
            for j := 0; j < len(values); j++ {
                result, ok := queue.Dequeue()
                if !ok {
                    t.Errorf("Dequeue failed at index %d", j)
                    break
                }
                
                if result != values[j] {
                    t.Errorf("FIFO violated: expected %d, got %d at position %d", values[j], result, j)
                }
            }
        }
    })
    
    // Propiedad: Size consistency
    t.Run("size consistency", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            queue := NewQueue()
            operations := rand.Intn(50) + 10
            expectedSize := 0
            
            for j := 0; j < operations; j++ {
                if rand.Float32() < 0.6 || queue.IsEmpty() { // 60% enqueue, 40% dequeue
                    queue.Enqueue(rand.Intn(100))
                    expectedSize++
                } else {
                    _, ok := queue.Dequeue()
                    if ok {
                        expectedSize--
                    }
                }
                
                if queue.Size() != expectedSize {
                    t.Errorf("Size inconsistent: expected %d, got %d", expectedSize, queue.Size())
                }
            }
        }
    })
}

// Sorted list property testing
func TestSortedList_Properties(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    
    // Propiedad: Lista ordenada permanece ordenada
    t.Run("sorted list remains sorted", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            // Generar lista aleatoria
            list := generateRandomInts(rand.Intn(50) + 1)
            
            // Ordenar
            sort.Ints(list)
            
            // Verificar que está ordenada
            if !isSorted(list) {
                t.Errorf("List not sorted after sort.Ints: %v", list)
            }
        }
    })
    
    // Propiedad: Inserción en lista ordenada mantiene orden
    t.Run("insertion maintains order", func(t *testing.T) {
        for i := 0; i < 100; i++ {
            // Crear lista ordenada
            list := generateRandomInts(rand.Intn(20) + 1)
            sort.Ints(list)
            
            // Insertar elemento aleatorio
            newElement := rand.Intn(200) - 100
            newList := insertSorted(list, newElement)
            
            // Verificar que sigue ordenada
            if !isSorted(newList) {
                t.Errorf("List not sorted after insertion: %v", newList)
            }
            
            // Verificar que el elemento está presente
            found := false
            for _, v := range newList {
                if v == newElement {
                    found = true
                    break
                }
            }
            if !found {
                t.Errorf("Inserted element %d not found in list %v", newElement, newList)
            }
        }
    })
}

// Helper functions
func generateRandomInts(count int) []int {
    ints := make([]int, count)
    for i := range ints {
        ints[i] = rand.Intn(200) - 100 // Números entre -100 y 100
    }
    return ints
}

func isSorted(list []int) bool {
    for i := 1; i < len(list); i++ {
        if list[i] < list[i-1] {
            return false
        }
    }
    return true
}

func insertSorted(list []int, element int) []int {
    // Encontrar posición de inserción
    insertPos := 0
    for i, v := range list {
        if v > element {
            insertPos = i
            break
        }
        insertPos = i + 1
    }
    
    // Crear nueva lista con elemento insertado
    newList := make([]int, len(list)+1)
    copy(newList, list[:insertPos])
    newList[insertPos] = element
    copy(newList[insertPos+1:], list[insertPos:])
    
    return newList
}
```

---

## 🎯 Resumen de la Lección

### ✅ Conceptos Clave Aprendidos

1. **🧪 Testing Ecosystem**: Built-in testing, herramientas y librerías
2. **🔄 TDD Cycle**: Red-Green-Refactor para diseño dirigido por tests
3. **🎭 Test Doubles**: Stubs, mocks y fakes para aislamiento
4. **🎯 Property Testing**: Verificación de invariantes matemáticas
5. **🏗️ Integration Testing**: Testing de sistemas completos
6. **⚡ Benchmarking**: Medición de performance y memory
7. **🧪 Test Organization**: Suites, setup/teardown y helpers
8. **🔧 Advanced Techniques**: Testing concurrente y edge cases

### 🏆 Logros Desbloqueados

- [ ] 🥇 **Testing Rookie**: Primeros tests unitarios
- [ ] 🥈 **TDD Practitioner**: Desarrollo dirigido por tests
- [ ] 🥉 **Mock Master**: Dominio de test doubles
- [ ] 🏅 **Property Tester**: Testing basado en propiedades
- [ ] 🎖️ **Integration Expert**: Tests de sistema completo
- [ ] 🏆 **Testing Wizard**: Maestría completa en testing

### 📚 Próximos Pasos

En la **Lección 18: Performance y Optimization**, aprenderemos:
- Profiling y benchmarking avanzado
- Optimización de memoria y CPU
- Técnicas de caching
- Optimization patterns

---

**🎉 ¡Felicitaciones! Has dominado el testing avanzado en Go. Ahora puedes crear test suites robustas que garantizan la calidad del código.**

*Recuerda: "Tests are the safety net that allows you to change code with confidence."* 🧪
