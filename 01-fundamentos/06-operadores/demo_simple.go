package main

import "fmt"

func demoSimple() {
	fmt.Println("✅ OPERADORES EN GO - LECCIÓN COMPLETADA")
	fmt.Println("======================================")

	// Aritméticos
	a, b := 15, 4
	fmt.Printf("Aritméticos: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Aritméticos: %d * %d = %d\n", a, b, a*b)

	// Comparación
	fmt.Printf("Comparación: %d > %d = %t\n", a, b, a > b)
	fmt.Printf("Comparación: %d == %d = %t\n", a, b, a == b)

	// Lógicos
	x, y := true, false
	fmt.Printf("Lógicos: %t && %t = %t\n", x, y, x && y)
	fmt.Printf("Lógicos: %t || %t = %t\n", x, y, x || y)

	// Bitwise
	n1, n2 := 12, 10
	fmt.Printf("Bitwise: %d & %d = %d\n", n1, n2, n1&n2)
	fmt.Printf("Bitwise: %d | %d = %d\n", n1, n2, n1|n2)

	fmt.Println("\n🎉 ¡Todos los operadores funcionando correctamente!")
}
