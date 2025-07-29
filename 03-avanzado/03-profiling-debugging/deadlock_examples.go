package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// 🐛 Ejemplo de código con deadlock para debugging

// BankAccount representa una cuenta bancaria con mutex para concurrencia
type BankAccount struct {
	mu      sync.Mutex
	balance int
	id      string
}

// NewBankAccount crea una nueva cuenta bancaria
func NewBankAccount(id string, initialBalance int) *BankAccount {
	return &BankAccount{
		balance: initialBalance,
		id:      id,
	}
}

// GetBalance retorna el balance actual (thread-safe)
func (ba *BankAccount) GetBalance() int {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	return ba.balance
}

// Transfer transfiere dinero entre cuentas - VERSIÓN CON DEADLOCK
func (ba *BankAccount) Transfer(to *BankAccount, amount int) error {
	fmt.Printf("🔄 Attempting transfer: %s -> %s, amount: %d\n", ba.id, to.id, amount)

	// ⚠️ PROBLEMA: Orden inconsistente de locks puede causar deadlock
	ba.mu.Lock()
	fmt.Printf("🔒 %s: acquired source lock\n", ba.id)

	// Simular trabajo que toma tiempo
	time.Sleep(100 * time.Millisecond)

	// ⚠️ AQUÍ PUEDE OCURRIR EL DEADLOCK
	// Si otra goroutine hace Transfer en dirección opuesta,
	// puede haber un ciclo de espera
	to.mu.Lock()
	fmt.Printf("🔒 %s: acquired destination lock\n", to.id)

	// Verificar fondos suficientes
	if ba.balance < amount {
		to.mu.Unlock()
		ba.mu.Unlock()
		return fmt.Errorf("insufficient funds: %d < %d", ba.balance, amount)
	}

	// Realizar transferencia
	ba.balance -= amount
	to.balance += amount

	fmt.Printf("✅ Transfer complete: %s(%d) -> %s(%d)\n",
		ba.id, ba.balance, to.id, to.balance)

	to.mu.Unlock()
	ba.mu.Unlock()

	return nil
}

// TransferSafe - Versión sin deadlock usando orden consistente
func (ba *BankAccount) TransferSafe(to *BankAccount, amount int) error {
	fmt.Printf("🛡️ Safe transfer: %s -> %s, amount: %d\n", ba.id, to.id, amount)

	// 🚀 SOLUCIÓN: Orden consistente de locks basado en ID
	first, second := ba, to
	if ba.id > to.id {
		first, second = to, ba
	}

	first.mu.Lock()
	fmt.Printf("🔒 Acquired first lock: %s\n", first.id)

	second.mu.Lock()
	fmt.Printf("🔒 Acquired second lock: %s\n", second.id)

	// Verificar fondos (ba es siempre el origen)
	if ba.balance < amount {
		second.mu.Unlock()
		first.mu.Unlock()
		return fmt.Errorf("insufficient funds: %d < %d", ba.balance, amount)
	}

	// Realizar transferencia
	ba.balance -= amount
	to.balance += amount

	fmt.Printf("✅ Safe transfer complete: %s(%d) -> %s(%d)\n",
		ba.id, ba.balance, to.id, to.balance)

	second.mu.Unlock()
	first.mu.Unlock()

	return nil
}

// DeadlockDemo demuestra el deadlock
func DeadlockDemo() {
	fmt.Println("🚨 DEADLOCK DEMO - This will hang!")
	fmt.Println("Use Ctrl+C to stop or analyze with debugger")
	fmt.Println()

	accountA := NewBankAccount("A", 1000)
	accountB := NewBankAccount("B", 1000)

	var wg sync.WaitGroup

	// Goroutine 1: A -> B
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if err := accountA.Transfer(accountB, 100); err != nil {
				fmt.Printf("❌ Transfer A->B failed: %v\n", err)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Goroutine 2: B -> A (dirección opuesta - causa deadlock)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if err := accountB.Transfer(accountA, 100); err != nil {
				fmt.Printf("❌ Transfer B->A failed: %v\n", err)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Timeout para evitar hang infinito en demo
	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("✅ All transfers completed (shouldn't reach here with deadlock)")
	case <-time.After(3 * time.Second):
		fmt.Println("⏰ Timeout reached - likely deadlock occurred!")
		fmt.Println("💡 Use 'dlv debug' to analyze the deadlock")
	}

	fmt.Printf("Final balances: A=%d, B=%d\n",
		accountA.GetBalance(), accountB.GetBalance())
}

// SafeDemo demuestra la versión sin deadlock
func SafeDemo() {
	fmt.Println("🛡️ SAFE DEMO - No deadlock!")
	fmt.Println()

	accountA := NewBankAccount("A", 1000)
	accountB := NewBankAccount("B", 1000)

	var wg sync.WaitGroup

	// Goroutine 1: A -> B usando método safe
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if err := accountA.TransferSafe(accountB, 100); err != nil {
				fmt.Printf("❌ Safe transfer A->B failed: %v\n", err)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Goroutine 2: B -> A usando método safe
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if err := accountB.TransferSafe(accountA, 100); err != nil {
				fmt.Printf("❌ Safe transfer B->A failed: %v\n", err)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("✅ All safe transfers completed!")
	fmt.Printf("Final balances: A=%d, B=%d\n",
		accountA.GetBalance(), accountB.GetBalance())
}

// 🔍 Herramientas para detectar el deadlock

// DetectGoroutines muestra información de goroutines para debugging
func DetectGoroutines() {
	fmt.Println("🔍 Goroutine debugging info:")
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())

	// En un debugger real, usarías:
	// (dlv) goroutines
	// (dlv) goroutine 1 bt
	// (dlv) goroutine 2 bt
}

// 🎯 Race Condition Example
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// UnsafeCounter para demostrar race conditions
type UnsafeCounter struct {
	value int // Sin mutex - race condition!
}

func (uc *UnsafeCounter) Increment() {
	uc.value++ // ⚠️ Race condition aquí
}

func (uc *UnsafeCounter) Get() int {
	return uc.value // ⚠️ Race condition aquí también
}

func RaceConditionDemo() {
	fmt.Println("🏁 RACE CONDITION DEMO")
	fmt.Println("Run with: go run -race deadlock_examples.go race")
	fmt.Println()

	safeCounter := &Counter{}
	unsafeCounter := &UnsafeCounter{}

	const numGoroutines = 10
	const numIncrements = 1000

	var wg sync.WaitGroup

	// Test safe counter
	fmt.Println("Testing safe counter...")
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				safeCounter.Increment()
			}
		}()
	}
	wg.Wait()

	// Test unsafe counter
	fmt.Println("Testing unsafe counter...")
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				unsafeCounter.Increment()
			}
		}()
	}
	wg.Wait()

	expectedValue := numGoroutines * numIncrements
	safeValue := safeCounter.Get()
	unsafeValue := unsafeCounter.Get()

	fmt.Printf("Expected value: %d\n", expectedValue)
	fmt.Printf("Safe counter:   %d ✅\n", safeValue)
	fmt.Printf("Unsafe counter: %d %s\n", unsafeValue,
		map[bool]string{true: "✅", false: "❌"}[unsafeValue == expectedValue])

	if unsafeValue != expectedValue {
		fmt.Println("💡 Race condition detected! Values don't match expected.")
		fmt.Println("   Run with -race flag to see detailed race reports.")
	}
}

// DebugMain - función para ejecutar demos de debugging
func DebugMain() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "deadlock":
			DeadlockDemo()
		case "safe":
			SafeDemo()
		case "race":
			RaceConditionDemo()
		case "goroutines":
			DetectGoroutines()
		default:
			fmt.Println("Usage: go run deadlock_examples.go [deadlock|safe|race|goroutines]")
		}
	} else {
		fmt.Println("🐛 Deadlock & Race Condition Examples")
		fmt.Println()
		fmt.Println("Available demos:")
		fmt.Println("  deadlock   - Demonstrate deadlock scenario")
		fmt.Println("  safe       - Demonstrate deadlock-free version")
		fmt.Println("  race       - Demonstrate race conditions")
		fmt.Println("  goroutines - Show goroutine debugging info")
		fmt.Println()
		fmt.Println("Debugging commands:")
		fmt.Println("  go run -race deadlock_examples.go race  - Detect race conditions")
		fmt.Println("  dlv debug deadlock_examples.go          - Debug with Delve")
		fmt.Println()
		fmt.Println("Delve commands for deadlock analysis:")
		fmt.Println("  (dlv) b main.DeadlockDemo")
		fmt.Println("  (dlv) c")
		fmt.Println("  (dlv) goroutines")
		fmt.Println("  (dlv) goroutine 1 bt")
		fmt.Println("  (dlv) goroutine 2 bt")
	}
}
