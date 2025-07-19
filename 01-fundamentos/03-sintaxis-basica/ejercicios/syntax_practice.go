package main

import "fmt"

// TODO: Agrega documentación para esta constante
const MaxAttempts = 3

// TODO: Agrega documentación para esta variable
var applicationName = "Go Syntax Practice"

// TODO: Documenta esta función
func greetUser(name string) {
    fmt.Printf("¡Hola, %s! Bienvenido a %s\n", name, applicationName)
}

// TODO: Documenta esta función
func CalculateScore(correct, total int) float64 {
    if total == 0 {
        return 0
    }
    return float64(correct) / float64(total) * 100
}

// TODO: Agrega defer statement apropiado
func main() {
    fmt.Println("=== Ejercicio de Sintaxis Go ===")
    
    greetUser("Estudiante")
    
    score := CalculateScore(8, 10)
    fmt.Printf("Tu puntuación: %.1f%%\n", score)
    
    // TODO: Agregar un mensaje de despedida usando defer
}
