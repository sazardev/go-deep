// ==============================================
// LECCIÓN 14: Channels - Ejercicios
// ==============================================

package main

import (
	"fmt"
	"runtime"
	"time"
)

// ==============================================
// EJERCICIO 1: Primer Channel
// ==============================================
// Objetivo: Crear tu primer channel y entender la comunicación básica

func ejercicio1() {
	fmt.Println("📡 Ejercicio 1: Primer Channel")
	fmt.Println("=============================")

	// TODO: Crear un channel de strings
	// TODO: Crear una función que envíe 3 mensajes al channel
	// TODO: En main, recibir e imprimir cada mensaje
	// TODO: Observar el comportamiento de bloqueo

	fmt.Println("Nota: Observa cómo el emisor espera al receptor\n")
}

// ==============================================
// EJERCICIO 2: Channel Buffered vs Unbuffered
// ==============================================
// Objetivo: Entender la diferencia entre channels buffered y unbuffered

func ejercicio2() {
	fmt.Println("📦 Ejercicio 2: Buffered vs Unbuffered")
	fmt.Println("=====================================")

	// TODO: Crear función para probar channel unbuffered
	//       - Crear channel sin buffer
	//       - Intentar enviar 3 mensajes desde main
	//       - Observar qué pasa (debería bloquear)

	// TODO: Crear función para probar channel buffered
	//       - Crear channel con buffer de 3
	//       - Enviar 3 mensajes desde main
	//       - Observar que no bloquea
	//       - Leer los mensajes después

	fmt.Println()
}

// ==============================================
// EJERCICIO 3: Productor-Consumidor
// ==============================================
// Objetivo: Implementar el patrón productor-consumidor básico

func ejercicio3() {
	fmt.Println("🏭 Ejercicio 3: Productor-Consumidor")
	fmt.Println("===================================")

	// TODO: Crear función productor que:
	//       - Reciba un channel de enteros (send-only)
	//       - Genere números del 1 al 10
	//       - Envíe cada número al channel
	//       - Cierre el channel al terminar

	// TODO: Crear función consumidor que:
	//       - Reciba un channel de enteros (receive-only)
	//       - Use range para leer todos los números
	//       - Imprima cada número recibido

	// TODO: En main, crear channel y lanzar ambas funciones

	fmt.Println()
}

// ==============================================
// EJERCICIO 4: Select Statement Básico
// ==============================================
// Objetivo: Usar select para manejar múltiples channels

func ejercicio4() {
	fmt.Println("🎛️ Ejercicio 4: Select Statement")
	fmt.Println("==============================")

	// TODO: Crear 3 channels de strings
	// TODO: Crear 3 goroutines que envíen mensajes a cada channel
	//       con diferentes delays (100ms, 200ms, 300ms)
	// TODO: Usar select para recibir de cualquier channel que esté listo
	// TODO: Imprimir qué channel envió cada mensaje
	// TODO: Continuar hasta recibir todos los mensajes

	fmt.Println()
}

// ==============================================
// EJERCICIO 5: Select con Timeout
// ==============================================
// Objetivo: Implementar timeouts usando select

func ejercicio5() {
	fmt.Println("⏰ Ejercicio 5: Select con Timeout")
	fmt.Println("=================================")

	// TODO: Crear función que simule una operación lenta:
	//       - Reciba un channel y un delay
	//       - Espere el delay especificado
	//       - Envíe un resultado al channel

	// TODO: En main, crear channel y lanzar la operación lenta
	// TODO: Usar select con time.After() para implementar timeout de 2 segundos
	// TODO: Probar con delays de 1 segundo (éxito) y 3 segundos (timeout)

	fmt.Println()
}

// ==============================================
// EJERCICIO 6: Pipeline de Datos
// ==============================================
// Objetivo: Crear un pipeline de transformación de datos

func ejercicio6() {
	fmt.Println("🔄 Ejercicio 6: Pipeline de Datos")
	fmt.Println("================================")

	// TODO: Crear pipeline con 4 etapas:
	//       1. Generar números del 1 al 20
	//       2. Multiplicar por 2
	//       3. Filtrar solo números divisibles por 4
	//       4. Convertir a string con formato

	// TODO: Cada etapa debe ser una función que:
	//       - Reciba un channel de entrada
	//       - Retorne un channel de salida
	//       - Procese en una goroutine
	//       - Cierre el channel de salida al terminar

	// TODO: Conectar todas las etapas y mostrar resultados finales

	fmt.Println()
}

// ==============================================
// EJERCICIO 7: Worker Pool con Channels
// ==============================================
// Objetivo: Implementar worker pool usando channels

type Trabajo struct {
	ID   int
	Dato int
}

type Resultado struct {
	TrabajoID int
	Resultado int
}

func ejercicio7() {
	fmt.Println("👷 Ejercicio 7: Worker Pool")
	fmt.Println("===========================")

	// TODO: Crear función worker que:
	//       - Reciba channel de trabajos (receive-only)
	//       - Reciba channel de resultados (send-only)
	//       - Procese cada trabajo (calcular cuadrado del dato)
	//       - Envíe resultado al channel de resultados

	// TODO: En main:
	//       - Crear channels para trabajos y resultados
	//       - Lanzar 3 workers
	//       - Enviar 15 trabajos
	//       - Cerrar channel de trabajos
	//       - Recoger todos los resultados

	fmt.Println()
}

// ==============================================
// EJERCICIO 8: Fan-Out/Fan-In
// ==============================================
// Objetivo: Distribuir trabajo y combinar resultados

func ejercicio8() {
	fmt.Println("📊 Ejercicio 8: Fan-Out/Fan-In")
	fmt.Println("==============================")

	// TODO: Crear función que genere números del 1 al 50

	// TODO: Crear función worker que:
	//       - Calcule si un número es primo
	//       - Retorne el número si es primo, -1 si no

	// TODO: Crear función merge que:
	//       - Combine resultados de múltiples workers
	//       - Use sync.WaitGroup para sincronización

	// TODO: En main:
	//       - Distribuir números entre 5 workers (Fan-Out)
	//       - Combinar todos los primos encontrados (Fan-In)
	//       - Mostrar lista de números primos

	fmt.Println()
}

// ==============================================
// EJERCICIO 9: Quit Channel Pattern
// ==============================================
// Objetivo: Implementar cancelación elegante con quit channel

func ejercicio9() {
	fmt.Println("🛑 Ejercicio 9: Quit Channel")
	fmt.Println("===========================")

	// TODO: Crear función servidor que:
	//       - Use un ticker para generar eventos cada 200ms
	//       - Reciba un quit channel
	//       - Use select para manejar events y quit
	//       - Haga cleanup cuando reciba quit

	// TODO: En main:
	//       - Lanzar el servidor
	//       - Dejarlo correr por 3 segundos
	//       - Enviar señal de quit
	//       - Esperar a que termine limpiamente

	fmt.Println()
}

// ==============================================
// EJERCICIO 10: Channel de Channels
// ==============================================
// Objetivo: Usar channels de channels para multiplexación dinámica

func ejercicio10() {
	fmt.Println("📡 Ejercicio 10: Channel de Channels")
	fmt.Println("===================================")

	// TODO: Crear función que genere channels dinámicamente:
	//       - Retorne un channel de channels
	//       - Cada channel interno envíe una secuencia de números
	//       - Crear 3 channels internos con diferentes secuencias

	// TODO: Crear función multiplexor que:
	//       - Reciba channels dinámicamente
	//       - Combine todos los valores en un solo stream
	//       - Use select para leer de channels disponibles

	// TODO: Mostrar cómo los valores se intercalan dinámicamente

	fmt.Println()
}

// ==============================================
// FUNCIÓN PRINCIPAL Y UTILIDADES
// ==============================================

func ejecutarEjercicios() {
	fmt.Println("📡 EJERCICIOS: Channels")
	fmt.Println("======================")
	fmt.Printf("Usando Go %s con %d CPUs\n\n", runtime.Version(), runtime.NumCPU())

	ejercicios := []func(){
		ejercicio1,
		ejercicio2,
		ejercicio3,
		ejercicio4,
		ejercicio5,
		ejercicio6,
		ejercicio7,
		ejercicio8,
		ejercicio9,
		ejercicio10,
	}

	for i, ejercicio := range ejercicios {
		fmt.Printf("📝 Ejecutando ejercicio %d...\n", i+1)
		ejercicio()
		time.Sleep(500 * time.Millisecond) // Pausa entre ejercicios
	}

	fmt.Println("🎉 ¡Todos los ejercicios completados!")
	fmt.Println("\n💡 Consejos para seguir practicando:")
	fmt.Println("   - Experimenta con diferentes buffer sizes")
	fmt.Println("   - Prueba con más workers en el pool")
	fmt.Println("   - Mide performance con channels vs mutex")
	fmt.Println("   - Implementa patrones de retry con channels")
	fmt.Println("   - Crea sistemas pub/sub con channels")
}

func main() {
	ejecutarEjercicios()
}

// ==============================================
// FUNCIONES AUXILIARES (Para implementar)
// ==============================================

// TODO: Implementar función para verificar si un número es primo
func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// TODO: Implementar función para simular procesamiento
func simularProcesamiento(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
