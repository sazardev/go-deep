// 🔄 EJERCICIOS: ESTRUCTURAS DE CONTROL
// Nivel: Fundamentos
// Lección 7: Estructuras de Control

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ========== EJERCICIO 1: VALIDADOR DE NÚMEROS ==========

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Validador de Números ===")
	fmt.Println("Ingresa números positivos (0 para terminar):")

	var numeros []int
	var suma int

	for {
		fmt.Print("Número: ")
		var num int
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if num < 0 {
			fmt.Println("❌ Solo números positivos. Intenta de nuevo.")
			continue
		}

		numeros = append(numeros, num)
		suma += num
		fmt.Printf("✅ Número %d agregado. Total actual: %d\n", num, len(numeros))
	}

	if len(numeros) == 0 {
		fmt.Println("No se ingresaron números válidos.")
		return
	}

	// Calcular estadísticas
	promedio := float64(suma) / float64(len(numeros))
	min, max := numeros[0], numeros[0]

	for _, num := range numeros[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("\n📊 Estadísticas:\n")
	fmt.Printf("   Cantidad: %d números\n", len(numeros))
	fmt.Printf("   Suma: %d\n", suma)
	fmt.Printf("   Promedio: %.2f\n", promedio)
	fmt.Printf("   Mínimo: %d\n", min)
	fmt.Printf("   Máximo: %d\n", max)
}

// ========== EJERCICIO 2: PIEDRA, PAPEL, TIJERA ==========

func ejercicio2() {
	fmt.Println("\n=== Ejercicio 2: Piedra, Papel, Tijera ===")

	opciones := []string{"piedra", "papel", "tijera"}
	victorias := map[string]int{"jugador": 0, "computadora": 0, "empates": 0}

	for {
		fmt.Println("\n🎮 Elige tu jugada:")
		fmt.Println("1. Piedra")
		fmt.Println("2. Papel")
		fmt.Println("3. Tijera")
		fmt.Println("4. Ver estadísticas")
		fmt.Println("5. Salir")

		var opcion int
		fmt.Print("Opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1, 2, 3:
			jugadaJugador := opciones[opcion-1]
			jugadaComputadora := opciones[rand.Intn(3)]

			fmt.Printf("\nTú: %s vs Computadora: %s\n", jugadaJugador, jugadaComputadora)

			switch {
			case jugadaJugador == jugadaComputadora:
				fmt.Println("🤝 ¡Empate!")
				victorias["empates"]++
			case (jugadaJugador == "piedra" && jugadaComputadora == "tijera") ||
				(jugadaJugador == "papel" && jugadaComputadora == "piedra") ||
				(jugadaJugador == "tijera" && jugadaComputadora == "papel"):
				fmt.Println("🎉 ¡Ganaste!")
				victorias["jugador"]++
			default:
				fmt.Println("💻 Ganó la computadora")
				victorias["computadora"]++
			}

		case 4:
			total := victorias["jugador"] + victorias["computadora"] + victorias["empates"]
			if total == 0 {
				fmt.Println("No hay partidas jugadas aún.")
			} else {
				fmt.Printf("\n📊 Estadísticas:\n")
				fmt.Printf("   Jugador: %d (%.1f%%)\n", victorias["jugador"], float64(victorias["jugador"])/float64(total)*100)
				fmt.Printf("   Computadora: %d (%.1f%%)\n", victorias["computadora"], float64(victorias["computadora"])/float64(total)*100)
				fmt.Printf("   Empates: %d (%.1f%%)\n", victorias["empates"], float64(victorias["empates"])/float64(total)*100)
			}

		case 5:
			fmt.Println("👋 ¡Gracias por jugar!")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}

// ========== EJERCICIO 3: ANALIZADOR DE TEXTO ==========

func ejercicio3() {
	fmt.Println("\n=== Ejercicio 3: Analizador de Texto ===")
	fmt.Print("Ingresa un texto para analizar: ")

	var texto string
	fmt.Scanln(&texto)

	if strings.TrimSpace(texto) == "" {
		fmt.Println("❌ Texto vacío")
		return
	}

	// Análisis básico
	caracteres := len(texto)
	palabras := len(strings.Fields(texto))
	lineas := strings.Count(texto, "\n") + 1

	// Frecuencia de letras
	frecuencias := make(map[rune]int)
	for _, char := range strings.ToLower(texto) {
		if char >= 'a' && char <= 'z' {
			frecuencias[char]++
		}
	}

	// Palabras más comunes
	palabrasList := strings.Fields(strings.ToLower(texto))
	frecuenciaPalabras := make(map[string]int)
	for _, palabra := range palabrasList {
		// Limpiar puntuación básica
		palabra = strings.Trim(palabra, ".,!?;:")
		if len(palabra) > 0 {
			frecuenciaPalabras[palabra]++
		}
	}

	fmt.Printf("\n📊 Análisis del texto:\n")
	fmt.Printf("   Caracteres: %d\n", caracteres)
	fmt.Printf("   Palabras: %d\n", palabras)
	fmt.Printf("   Líneas: %d\n", lineas)

	fmt.Println("\n🔤 Frecuencia de letras:")
	for letra := 'a'; letra <= 'z'; letra++ {
		if count, existe := frecuencias[letra]; existe && count > 0 {
			fmt.Printf("   %c: %d\n", letra, count)
		}
	}

	fmt.Println("\n📝 Palabras más comunes:")
	for palabra, count := range frecuenciaPalabras {
		if count > 1 {
			fmt.Printf("   %s: %d veces\n", palabra, count)
		}
	}
}

// ========== EJERCICIO 4: SISTEMA DE CALIFICACIONES ==========

type Estudiante struct {
	Nombre         string
	Calificaciones map[string][]float64
}

func ejercicio4() {
	fmt.Println("\n=== Ejercicio 4: Sistema de Calificaciones ===")

	estudiantes := make(map[string]*Estudiante)

	for {
		fmt.Println("\n📚 Sistema de Calificaciones:")
		fmt.Println("1. Agregar estudiante")
		fmt.Println("2. Agregar calificación")
		fmt.Println("3. Ver promedios de estudiante")
		fmt.Println("4. Estadísticas por materia")
		fmt.Println("5. Reporte general")
		fmt.Println("6. Salir")

		var opcion int
		fmt.Print("Opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Nombre del estudiante: ")
			var nombre string
			fmt.Scanln(&nombre)

			if _, existe := estudiantes[nombre]; existe {
				fmt.Println("❌ El estudiante ya existe")
			} else {
				estudiantes[nombre] = &Estudiante{
					Nombre:         nombre,
					Calificaciones: make(map[string][]float64),
				}
				fmt.Printf("✅ Estudiante %s agregado\n", nombre)
			}

		case 2:
			fmt.Print("Nombre del estudiante: ")
			var nombre string
			fmt.Scanln(&nombre)

			estudiante, existe := estudiantes[nombre]
			if !existe {
				fmt.Println("❌ Estudiante no encontrado")
				continue
			}

			fmt.Print("Materia: ")
			var materia string
			fmt.Scanln(&materia)

			fmt.Print("Calificación (0-10): ")
			var calificacion float64
			fmt.Scanln(&calificacion)

			if calificacion < 0 || calificacion > 10 {
				fmt.Println("❌ Calificación debe estar entre 0 y 10")
				continue
			}

			estudiante.Calificaciones[materia] = append(estudiante.Calificaciones[materia], calificacion)
			fmt.Printf("✅ Calificación %.1f agregada a %s en %s\n", calificacion, nombre, materia)

		case 3:
			fmt.Print("Nombre del estudiante: ")
			var nombre string
			fmt.Scanln(&nombre)

			estudiante, existe := estudiantes[nombre]
			if !existe {
				fmt.Println("❌ Estudiante no encontrado")
				continue
			}

			fmt.Printf("\n📊 Promedios de %s:\n", nombre)
			for materia, califs := range estudiante.Calificaciones {
				suma := 0.0
				for _, calif := range califs {
					suma += calif
				}
				promedio := suma / float64(len(califs))
				fmt.Printf("   %s: %.2f (%d calificaciones)\n", materia, promedio, len(califs))
			}

		case 4:
			fmt.Print("Materia: ")
			var materia string
			fmt.Scanln(&materia)

			var todasCalifs []float64
			estudiantesConMateria := 0

			for _, estudiante := range estudiantes {
				if califs, existe := estudiante.Calificaciones[materia]; existe {
					estudiantesConMateria++
					todasCalifs = append(todasCalifs, califs...)
				}
			}

			if len(todasCalifs) == 0 {
				fmt.Printf("❌ No hay calificaciones para %s\n", materia)
				continue
			}

			suma := 0.0
			for _, calif := range todasCalifs {
				suma += calif
			}
			promedio := suma / float64(len(todasCalifs))

			fmt.Printf("\n📊 Estadísticas de %s:\n", materia)
			fmt.Printf("   Estudiantes: %d\n", estudiantesConMateria)
			fmt.Printf("   Total calificaciones: %d\n", len(todasCalifs))
			fmt.Printf("   Promedio general: %.2f\n", promedio)

		case 5:
			if len(estudiantes) == 0 {
				fmt.Println("❌ No hay estudiantes registrados")
				continue
			}

			fmt.Println("\n📊 Reporte General:")
			for nombre, estudiante := range estudiantes {
				fmt.Printf("\n👤 %s:\n", nombre)
				if len(estudiante.Calificaciones) == 0 {
					fmt.Println("   Sin calificaciones")
					continue
				}

				for materia, califs := range estudiante.Calificaciones {
					suma := 0.0
					for _, calif := range califs {
						suma += calif
					}
					promedio := suma / float64(len(califs))

					estado := ""
					switch {
					case promedio >= 9:
						estado = "🏆 Excelente"
					case promedio >= 7:
						estado = "✅ Bueno"
					case promedio >= 6:
						estado = "⚠️ Suficiente"
					default:
						estado = "❌ Insuficiente"
					}

					fmt.Printf("   %s: %.2f %s\n", materia, promedio, estado)
				}
			}

		case 6:
			fmt.Println("👋 Saliendo del sistema")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}

// ========== EJERCICIO 5: CONVERSOR DE UNIDADES ==========

func ejercicio5() {
	fmt.Println("\n=== Ejercicio 5: Conversor de Unidades ===")

	for {
		fmt.Println("\n🔄 Conversor de Unidades:")
		fmt.Println("1. Longitud")
		fmt.Println("2. Peso")
		fmt.Println("3. Temperatura")
		fmt.Println("4. Salir")

		var categoria int
		fmt.Print("Categoría: ")
		fmt.Scanln(&categoria)

		switch categoria {
		case 1:
			convertirLongitud()
		case 2:
			convertirPeso()
		case 3:
			convertirTemperatura()
		case 4:
			fmt.Println("👋 Saliendo del conversor")
			return
		default:
			fmt.Println("❌ Categoría inválida")
		}
	}
}

func convertirLongitud() {
	fmt.Println("\n📏 Conversor de Longitud:")
	fmt.Println("1. Metros a Pies")
	fmt.Println("2. Pies a Metros")
	fmt.Println("3. Metros a Pulgadas")
	fmt.Println("4. Pulgadas a Metros")

	var opcion int
	fmt.Print("Conversión: ")
	fmt.Scanln(&opcion)

	fmt.Print("Valor: ")
	var valor float64
	fmt.Scanln(&valor)

	switch opcion {
	case 1:
		resultado := valor * 3.28084
		fmt.Printf("%.2f metros = %.2f pies\n", valor, resultado)
	case 2:
		resultado := valor / 3.28084
		fmt.Printf("%.2f pies = %.2f metros\n", valor, resultado)
	case 3:
		resultado := valor * 39.3701
		fmt.Printf("%.2f metros = %.2f pulgadas\n", valor, resultado)
	case 4:
		resultado := valor / 39.3701
		fmt.Printf("%.2f pulgadas = %.2f metros\n", valor, resultado)
	default:
		fmt.Println("❌ Opción inválida")
	}
}

func convertirPeso() {
	fmt.Println("\n⚖️ Conversor de Peso:")
	fmt.Println("1. Kilogramos a Libras")
	fmt.Println("2. Libras a Kilogramos")
	fmt.Println("3. Kilogramos a Onzas")
	fmt.Println("4. Onzas a Kilogramos")

	var opcion int
	fmt.Print("Conversión: ")
	fmt.Scanln(&opcion)

	fmt.Print("Valor: ")
	var valor float64
	fmt.Scanln(&valor)

	switch opcion {
	case 1:
		resultado := valor * 2.20462
		fmt.Printf("%.2f kg = %.2f libras\n", valor, resultado)
	case 2:
		resultado := valor / 2.20462
		fmt.Printf("%.2f libras = %.2f kg\n", valor, resultado)
	case 3:
		resultado := valor * 35.274
		fmt.Printf("%.2f kg = %.2f onzas\n", valor, resultado)
	case 4:
		resultado := valor / 35.274
		fmt.Printf("%.2f onzas = %.2f kg\n", valor, resultado)
	default:
		fmt.Println("❌ Opción inválida")
	}
}

func convertirTemperatura() {
	fmt.Println("\n🌡️ Conversor de Temperatura:")
	fmt.Println("1. Celsius a Fahrenheit")
	fmt.Println("2. Fahrenheit a Celsius")
	fmt.Println("3. Celsius a Kelvin")
	fmt.Println("4. Kelvin a Celsius")

	var opcion int
	fmt.Print("Conversión: ")
	fmt.Scanln(&opcion)

	fmt.Print("Valor: ")
	var valor float64
	fmt.Scanln(&valor)

	switch opcion {
	case 1:
		resultado := (valor * 9 / 5) + 32
		fmt.Printf("%.2f°C = %.2f°F\n", valor, resultado)
	case 2:
		resultado := (valor - 32) * 5 / 9
		fmt.Printf("%.2f°F = %.2f°C\n", valor, resultado)
	case 3:
		resultado := valor + 273.15
		fmt.Printf("%.2f°C = %.2f K\n", valor, resultado)
	case 4:
		resultado := valor - 273.15
		fmt.Printf("%.2f K = %.2f°C\n", valor, resultado)
	default:
		fmt.Println("❌ Opción inválida")
	}
}

// ========== FUNCIÓN PRINCIPAL ==========

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🔄 === LABORATORIO: ESTRUCTURAS DE CONTROL ===")
	fmt.Println("Elige un ejercicio para probar:")

	for {
		fmt.Println("\n📚 Ejercicios disponibles:")
		fmt.Println("1. Validador de números")
		fmt.Println("2. Piedra, papel, tijera")
		fmt.Println("3. Analizador de texto")
		fmt.Println("4. Sistema de calificaciones")
		fmt.Println("5. Conversor de unidades")
		fmt.Println("6. Salir")

		var opcion int
		fmt.Print("Ejercicio: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			ejercicio1()
		case 2:
			ejercicio2()
		case 3:
			ejercicio3()
		case 4:
			ejercicio4()
		case 5:
			ejercicio5()
		case 6:
			fmt.Println("🎯 ¡Excelente trabajo practicando estructuras de control!")
			fmt.Println("📚 Conceptos aplicados:")
			fmt.Println("   ✅ Condicionales if-else")
			fmt.Println("   ✅ Bucles for en todas sus variantes")
			fmt.Println("   ✅ Switch statements")
			fmt.Println("   ✅ Control de flujo con break/continue")
			fmt.Println("   ✅ Manejo de datos y estructuras")
			return
		default:
			fmt.Println("❌ Opción inválida")
		}

		fmt.Println("\nPresiona Enter para continuar...")
		fmt.Scanln()
	}
}
