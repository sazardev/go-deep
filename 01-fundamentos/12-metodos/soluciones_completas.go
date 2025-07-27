package main

import (
"errors"
"fmt"
"time"
)

// ========================================
// Ejercicio 1: User con métodos básicos
// ========================================

type User struct {
Name  string
Email string
Age   int
}

// Método String (value receiver) - no modifica el struct
func (u User) String() string {
return fmt.Sprintf("User{Name: %s, Email: %s, Age: %d}", u.Name, u.Email, u.Age)
}

// Método IsAdult (value receiver) - solo consulta información
func (u User) IsAdult() bool {
return u.Age >= 18
}

// Método UpdateEmail (pointer receiver) - modifica el struct
func (u *User) UpdateEmail(newEmail string) {
u.Email = newEmail
}

// ========================================
// Ejercicio 2: BankAccount con operaciones bancarias
// ========================================

type BankAccount struct {
AccountNumber string
Balance       float64
Owner         string
CreatedAt     time.Time
}

// Constructor para BankAccount
func NewBankAccount(accountNumber, owner string, initialBalance float64) *BankAccount {
return &BankAccount{
AccountNumber: accountNumber,
Balance:       initialBalance,
Owner:         owner,
CreatedAt:     time.Now(),
}
}

// Método Deposit (pointer receiver) - modifica balance
func (ba *BankAccount) Deposit(amount float64) error {
if amount <= 0 {
return errors.New("monto debe ser mayor a cero")
}
ba.Balance += amount
return nil
}

// Método Withdraw (pointer receiver) - modifica balance
func (ba *BankAccount) Withdraw(amount float64) error {
if amount <= 0 {
return errors.New("monto debe ser mayor a cero")
}
if ba.Balance < amount {
return errors.New("fondos insuficientes")
}
ba.Balance -= amount
return nil
}

// Método GetBalance (value receiver) - solo consulta
func (ba BankAccount) GetBalance() float64 {
return ba.Balance
}

// Método Transfer - transferir entre cuentas
func (ba *BankAccount) Transfer(to *BankAccount, amount float64) error {
if err := ba.Withdraw(amount); err != nil {
return err
}
return to.Deposit(amount)
}

// ========================================
// Ejercicio 3: Calculator con Method Chaining
// ========================================

type Calculator struct {
Value float64
}

// Constructor para Calculator
func NewCalculator(initial float64) *Calculator {
return &Calculator{Value: initial}
}

// Method Chaining - todos retornan *Calculator
func (c *Calculator) Add(value float64) *Calculator {
c.Value += value
return c
}

func (c *Calculator) Subtract(value float64) *Calculator {
c.Value -= value
return c
}

func (c *Calculator) Multiply(value float64) *Calculator {
c.Value *= value
return c
}

func (c *Calculator) Divide(value float64) *Calculator {
if value != 0 {
c.Value /= value
}
return c
}

// Método Result - obtiene el valor final
func (c Calculator) Result() float64 {
return c.Value
}

// ========================================
// FUNCIÓN MAIN CON DEMOSTRACIONES
// ========================================

func main() {
fmt.Println("🎯 Soluciones Completas - Métodos en Go")
fmt.Println("=======================================")

// Demo User
fmt.Println("\n=== Demo User ===")
user := User{Name: "Ana García", Email: "ana@example.com", Age: 17}
fmt.Println("Usuario inicial:", user.String())
fmt.Println("Es adulto:", user.IsAdult())

user.UpdateEmail("ana.nueva@example.com")
fmt.Println("Email actualizado:", user.Email)

// Demo BankAccount
fmt.Println("\n=== Demo BankAccount ===")
account1 := NewBankAccount("001-2024", "Carlos López", 1500.0)
account2 := NewBankAccount("002-2024", "María Rodríguez", 800.0)

fmt.Printf("Cuentas iniciales:\n")
fmt.Printf("  %s: $%.2f\n", account1.Owner, account1.GetBalance())
fmt.Printf("  %s: $%.2f\n", account2.Owner, account2.GetBalance())

// Realizar operaciones
account1.Deposit(300)
fmt.Printf("Después de depositar $300 a %s: $%.2f\n", account1.Owner, account1.GetBalance())

err := account1.Transfer(account2, 200)
if err != nil {
fmt.Printf("Error en transferencia: %v\n", err)
} else {
fmt.Println("✅ Transferencia de $200 exitosa")
}

fmt.Printf("Balances finales:\n")
fmt.Printf("  %s: $%.2f\n", account1.Owner, account1.GetBalance())
fmt.Printf("  %s: $%.2f\n", account2.Owner, account2.GetBalance())

// Demo Calculator con Method Chaining
fmt.Println("\n=== Demo Calculator ===")
calc := NewCalculator(10)

result := calc.Add(5).Multiply(2).Subtract(3).Divide(2).Result()
fmt.Printf("Operación: (10 + 5) * 2 - 3 / 2 = %.2f\n", result)

// Nuevo cálculo
calc2 := NewCalculator(100)
result2 := calc2.Subtract(25).Multiply(0.5).Add(10).Result()
fmt.Printf("Operación: (100 - 25) * 0.5 + 10 = %.2f\n", result2)

// Demo de error handling
fmt.Println("\n=== Demo Error Handling ===")
testAccount := NewBankAccount("TEST", "Usuario Test", 50.0)

err = testAccount.Withdraw(100) // Intentar retirar más de lo disponible
if err != nil {
fmt.Printf("❌ Error esperado: %v\n", err)
}

err = testAccount.Deposit(-10) // Intentar depositar cantidad negativa
if err != nil {
fmt.Printf("❌ Error esperado: %v\n", err)
}

fmt.Println("\n✅ Todas las soluciones funcionando correctamente!")
fmt.Println("\n📋 Conceptos clave aprendidos:")
fmt.Println("  • Value receivers vs Pointer receivers")
fmt.Println("  • Method chaining")
fmt.Println("  • Constructors en Go")
fmt.Println("  • Error handling en métodos")
fmt.Println("  • Encapsulación y validación")
}
