// Package main demonstrates proper Go syntax and documentation.
// This program calculates scores and greets users following Go conventions.
package main

import "fmt"

// MaxAttempts defines the maximum number of retry attempts
// allowed for any operation in the application.
const MaxAttempts = 3

// applicationName holds the current application's display name.
// This variable is used throughout the application for branding.
var applicationName = "Go Syntax Practice"

// greetUser prints a personalized greeting message to the user.
// It combines the user's name with the application name for context.
func greetUser(name string) {
    fmt.Printf("¡Hola, %s! Bienvenido a %s\n", name, applicationName)
}

// CalculateScore computes the percentage score based on correct and total answers.
// Returns 0 if total is 0 to avoid division by zero.
// The result is returned as a float64 percentage (0-100).
func CalculateScore(correct, total int) float64 {
    if total == 0 {
        return 0
    }
    return float64(correct) / float64(total) * 100
}

// main is the entry point of the application.
// It demonstrates proper use of defer, function calls, and output formatting.
func main() {
    defer fmt.Println("\n¡Gracias por usar Go Syntax Practice!")
    
    fmt.Println("=== Ejercicio de Sintaxis Go ===")
    
    greetUser("Estudiante")
    
    score := CalculateScore(8, 10)
    fmt.Printf("Tu puntuación: %.1f%%\n", score)
}
