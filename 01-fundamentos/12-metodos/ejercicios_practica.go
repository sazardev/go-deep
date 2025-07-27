package main

import (
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

// TODO: Implementar método String() que retorne una representación del usuario
// func (u User) String() string {
//     // Tu código aquí
// }

// TODO: Implementar método IsAdult() que retorne true si age >= 18
// func (u User) IsAdult() bool {
//     // Tu código aquí
// }

// TODO: Implementar método UpdateEmail() que actualice el email (usar pointer receiver)
// func (u *User) UpdateEmail(newEmail string) {
//     // Tu código aquí
// }

// ========================================
// Ejercicio 2: BankAccount con operaciones bancarias
// ========================================

type BankAccount struct {
AccountNumber string
Balance       float64
Owner         string
CreatedAt     time.Time
}

// TODO: Implementar constructor NewBankAccount
// func NewBankAccount(accountNumber, owner string, initialBalance float64) *BankAccount {
//     // Tu código aquí
// }

// TODO: Implementar método Deposit (usar pointer receiver)
// func (ba *BankAccount) Deposit(amount float64) error {
//     // Tu código aquí - validar que amount > 0
// }

// TODO: Implementar método Withdraw (usar pointer receiver)
// func (ba *BankAccount) Withdraw(amount float64) error {
//     // Tu código aquí - validar amount > 0 y fondos suficientes
// }

// TODO: Implementar método GetBalance (usar value receiver)
// func (ba BankAccount) GetBalance() float64 {
//     // Tu código aquí
// }

// TODO: Implementar método Transfer entre cuentas
// func (ba *BankAccount) Transfer(to *BankAccount, amount float64) error {
//     // Tu código aquí - usar Withdraw y Deposit
// }

// ========================================
// Ejercicio 3: Calculator con Method Chaining
// ========================================

type Calculator struct {
Value float64
}

// TODO: Implementar constructor
// func NewCalculator(initial float64) *Calculator {
//     // Tu código aquí
// }

// TODO: Implementar métodos que retornen *Calculator para method chaining
// func (c *Calculator) Add(value float64) *Calculator {
//     // Tu código aquí
// }

// func (c *Calculator) Subtract(value float64) *Calculator {
//     // Tu código aquí
// }

// func (c *Calculator) Multiply(value float64) *Calculator {
//     // Tu código aquí
// }

// func (c *Calculator) Divide(value float64) *Calculator {
//     // Tu código aquí - validar division por cero
// }

// TODO: Implementar método Result() que retorne el valor final
// func (c Calculator) Result() float64 {
//     // Tu código aquí
// }

func main() {
fmt.Println("🎯 Ejercicios de Métodos en Go")
fmt.Println("==============================")

fmt.Println("\n📝 Instrucciones:")
fmt.Println("1. Descomenta los métodos TODO uno por uno")
fmt.Println("2. Implementa cada método siguiendo los comentarios")
fmt.Println("3. Ejecuta el programa para probar tu implementación")
fmt.Println("4. Los tests te dirán si tu implementación es correcta")

// Test User (descomenta cuando implementes los métodos)
/*
fmt.Println("\n=== Test User ===")
user := User{Name: "María", Email: "maria@test.com", Age: 20}
fmt.Println("Usuario:", user.String())
fmt.Println("Es adulto:", user.IsAdult())

user.UpdateEmail("maria.nueva@test.com")
fmt.Println("Email actualizado:", user.Email)
*/

// Test BankAccount (descomenta cuando implementes los métodos)
/*
fmt.Println("\n=== Test BankAccount ===")
acc1 := NewBankAccount("123", "Juan", 1000.0)
acc2 := NewBankAccount("456", "Ana", 500.0)

fmt.Printf("Balance inicial Juan: $%.2f\n", acc1.GetBalance())

acc1.Deposit(200)
fmt.Printf("Después de depositar $200: $%.2f\n", acc1.GetBalance())

err := acc1.Transfer(acc2, 150)
if err != nil {
fmt.Printf("Error: %v\n", err)
} else {
fmt.Println("✅ Transferencia exitosa")
fmt.Printf("Juan: $%.2f, Ana: $%.2f\n", acc1.GetBalance(), acc2.GetBalance())
}
*/

// Test Calculator (descomenta cuando implementes los métodos)
/*
fmt.Println("\n=== Test Calculator ===")
calc := NewCalculator(10)
result := calc.Add(5).Multiply(2).Subtract(3).Divide(2).Result()
fmt.Printf("(10 + 5) * 2 - 3 / 2 = %.2f\n", result)
*/

fmt.Println("\n✅ Comienza implementando los métodos marcados con TODO!")
}
