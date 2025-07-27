package main

import "fmt"

// User struct con métodos básicos
type User struct {
Name  string
Email string
Age   int
}

// Método String (value receiver)
func (u User) String() string {
return fmt.Sprintf("User{Name: %s, Email: %s, Age: %d}", u.Name, u.Email, u.Age)
}

// Método IsAdult (value receiver)
func (u User) IsAdult() bool {
return u.Age >= 18
}

// Método UpdateEmail (pointer receiver)
func (u *User) UpdateEmail(newEmail string) {
u.Email = newEmail
}

// BankAccount struct
type BankAccount struct {
Balance float64
Owner   string
}

// Constructor
func NewBankAccount(owner string, balance float64) *BankAccount {
return &BankAccount{Owner: owner, Balance: balance}
}

// Método Deposit (pointer receiver)
func (ba *BankAccount) Deposit(amount float64) {
ba.Balance += amount
}

// Método GetBalance (value receiver)
func (ba BankAccount) GetBalance() float64 {
return ba.Balance
}

func main() {
fmt.Println("🎭 Demo Métodos en Go")
fmt.Println("=====================")

// Demo User
fmt.Println("\n=== User Demo ===")
user := User{Name: "Juan", Email: "juan@example.com", Age: 25}
fmt.Println("Usuario:", user.String())
fmt.Println("Es adulto:", user.IsAdult())

user.UpdateEmail("nuevo@email.com")
fmt.Println("Email actualizado:", user.Email)

// Demo BankAccount
fmt.Println("\n=== BankAccount Demo ===")
account := NewBankAccount("Alice", 1000.0)
fmt.Printf("Balance inicial: $%.2f\n", account.GetBalance())

account.Deposit(250.0)
fmt.Printf("Después de depositar $250: $%.2f\n", account.GetBalance())

fmt.Println("\n✅ Demo completado exitosamente!")
}
