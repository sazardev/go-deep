// ==============================================
// LECCIÓN 13: Goroutines y Concurrencia - Ejercicios
// ==============================================

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// ==============================================
// EJERCICIO 1: Primera Goroutine
// ==============================================
// Objetivo: Crear tu primera goroutine y entender la diferencia
// entre ejecución secuencial y concurrente

func ejercicio1() {
	fmt.Println("🚀 Ejercicio 1: Primera Goroutine")
	fmt.Println("=====================================")

	// TODO: Implementar una función que imprima números del 1 al 5
	// con un delay de 100ms entre cada número

	// TODO: Ejecutar la función de forma secuencial dos veces

	// TODO: Ejecutar la función de forma concurrente dos veces usando goroutines

	// TODO: Usar time.Sleep para esperar que terminen las goroutines

	fmt.Println("Nota: Observa la diferencia en el orden de salida\n")
}

// ==============================================
// EJERCICIO 2: WaitGroup Básico
// ==============================================
// Objetivo: Aprender a sincronizar goroutines con WaitGroup

func ejercicio2() {
	fmt.Println("🔄 Ejercicio 2: WaitGroup Básico")
	fmt.Println("================================")

	// TODO: Crear un WaitGroup

	// TODO: Crear una función que simule un trabajador que:
	//       - Reciba un ID de trabajador
	//       - Imprima "Trabajador X iniciando"
	//       - Espere un tiempo aleatorio (100-500ms)
	//       - Imprima "Trabajador X terminando"
	//       - Use defer wg.Done()

	// TODO: Lanzar 5 trabajadores como goroutines

	// TODO: Esperar a que todos terminen con wg.Wait()

	fmt.Println()
}

// ==============================================
// EJERCICIO 3: Race Condition
// ==============================================
// Objetivo: Experimentar con race conditions y aprender a detectarlas

var contadorInseguro int // Variable global compartida

func ejercicio3() {
	fmt.Println("⚠️ Ejercicio 3: Race Condition")
	fmt.Println("==============================")

	// TODO: Crear una función que incremente contadorInseguro 1000 veces

	// TODO: Lanzar 10 goroutines que ejecuten esta función

	// TODO: Usar WaitGroup para esperar que terminen todas

	// TODO: Imprimir el valor final del contador
	//       Resultado esperado: 10000
	//       Resultado real: ¿Será diferente?

	// TODO: Ejecutar varias veces y observar resultados inconsistentes

	fmt.Println("Nota: Ejecuta con 'go run -race ejercicios.go' para detectar race conditions")
	fmt.Println()
}

// ==============================================
// EJERCICIO 4: Mutex para Sincronización
// ==============================================
// Objetivo: Resolver race conditions usando Mutex

var (
	contadorSeguro int
	// TODO: Declarar un Mutex
)

func ejercicio4() {
	fmt.Println("🔒 Ejercicio 4: Mutex para Sincronización")
	fmt.Println("=========================================")

	// TODO: Crear una función que:
	//       - Use mutex.Lock() antes de incrementar
	//       - Incremente contadorSeguro 1000 veces
	//       - Use mutex.Unlock() después de incrementar

	// TODO: Lanzar 10 goroutines con esta función segura

	// TODO: Usar WaitGroup para sincronización

	// TODO: Imprimir el resultado final (debería ser exactamente 10000)

	fmt.Println()
}

// ==============================================
// EJERCICIO 5: Atomic Operations
// ==============================================
// Objetivo: Usar operaciones atómicas para mejor performance

var contadorAtomico int64

func ejercicio5() {
	fmt.Println("⚡ Ejercicio 5: Atomic Operations")
	fmt.Println("=================================")

	// TODO: Crear una función que use atomic.AddInt64 para incrementar
	//       contadorAtomico 1000 veces de forma segura

	// TODO: Medir el tiempo de ejecución con time.Now() y time.Since()

	// TODO: Lanzar 10 goroutines con esta función

	// TODO: Comparar el tiempo con el ejercicio anterior (Mutex)

	// TODO: Imprimir resultado y tiempo transcurrido

	fmt.Println()
}

// ==============================================
// EJERCICIO 6: Worker Pool
// ==============================================
// Objetivo: Implementar el patrón Worker Pool para procesar trabajos

type Trabajo struct {
	ID   int
	Dato string
}

type Resultado struct {
	TrabajoID int
	Resultado string
	Error     error
}

func ejercicio6() {
	fmt.Println("🏗️ Ejercicio 6: Worker Pool")
	fmt.Println("===========================")

	const numWorkers = 3
	const numTrabajos = 15

	// TODO: Crear canales para trabajos y resultados

	// TODO: Crear función worker que:
	//       - Reciba trabajos de un canal
	//       - Procese cada trabajo (simular con time.Sleep)
	//       - Envíe resultados a otro canal
	//       - Use WaitGroup para sincronización

	// TODO: Lanzar workers como goroutines

	// TODO: Enviar trabajos al canal

	// TODO: Cerrar canal de trabajos

	// TODO: Recoger todos los resultados

	fmt.Println()
}

// ==============================================
// EJERCICIO 7: Pipeline de Procesamiento
// ==============================================
// Objetivo: Crear un pipeline de transformación de datos

func ejercicio7() {
	fmt.Println("🔄 Ejercicio 7: Pipeline de Procesamiento")
	fmt.Println("=========================================")

	// TODO: Crear un pipeline con 4 etapas:
	//       1. Generar números del 1 al 20
	//       2. Elevar al cuadrado
	//       3. Filtrar solo los pares
	//       4. Calcular la suma total

	// TODO: Cada etapa debe ser una goroutine separada

	// TODO: Usar canales para comunicar entre etapas

	// TODO: Imprimir el resultado final

	fmt.Println()
}

// ==============================================
// EJERCICIO 8: Context y Cancelación
// ==============================================
// Objetivo: Usar context para controlar y cancelar goroutines

func ejercicio8() {
	fmt.Println("🎯 Ejercicio 8: Context y Cancelación")
	fmt.Println("====================================")

	// TODO: Crear un context con timeout de 2 segundos

	// TODO: Crear una función que:
	//       - Reciba un context
	//       - Ejecute un bucle infinito
	//       - Use select con ctx.Done() para cancelación
	//       - Imprima mensajes de trabajo cada 200ms

	// TODO: Lanzar varias goroutines con esta función

	// TODO: Observar cómo se cancelan automáticamente después del timeout

	fmt.Println()
}

// ==============================================
// EJERCICIO 9: Fan-Out/Fan-In
// ==============================================
// Objetivo: Distribuir trabajo a múltiples workers y agregar resultados

func ejercicio9() {
	fmt.Println("📊 Ejercicio 9: Fan-Out/Fan-In")
	fmt.Println("==============================")

	// TODO: Crear función que genere números del 1 al 50

	// TODO: Crear función worker que calcule si un número es primo

	// TODO: Usar 5 workers para procesar en paralelo (Fan-Out)

	// TODO: Recoger todos los números primos encontrados (Fan-In)

	// TODO: Imprimir la lista de números primos

	fmt.Println()
}

// ==============================================
// EJERCICIO 10: Monitoreo y Estadísticas
// ==============================================
// Objetivo: Monitorear goroutines y recoger estadísticas

func ejercicio10() {
	fmt.Println("📊 Ejercicio 10: Monitoreo y Estadísticas")
	fmt.Println("=========================================")

	// TODO: Crear función que monitoree:
	//       - Número de goroutines activas (runtime.NumGoroutine())
	//       - Uso de memoria (runtime.ReadMemStats())
	//       - Número de CPUs (runtime.NumCPU())

	// TODO: Crear función que genere carga de trabajo:
	//       - Lance 100 goroutines
	//       - Cada una haga trabajo durante tiempo aleatorio

	// TODO: Mostrar estadísticas cada 100ms durante 3 segundos

	// TODO: Observar cómo cambian las métricas

	fmt.Println()
}

// ==============================================
// FUNCIÓN PRINCIPAL Y UTILIDADES
// ==============================================

func ejecutarEjercicios() {
	fmt.Println("🚀 EJERCICIOS: Goroutines y Concurrencia")
	fmt.Println("=========================================")
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
	fmt.Println("   - Ejecuta con -race para detectar race conditions")
	fmt.Println("   - Experimenta con diferentes números de workers")
	fmt.Println("   - Mide performance con diferentes enfoques")
	fmt.Println("   - Prueba con datasets más grandes")
}

func main() {
	ejecutarEjercicios()
}

// ==============================================
// FUNCIONES AUXILIARES (Para implementar)
// ==============================================

// TODO: Implementar función para simular trabajo CPU-intensivo
func simularTrabajoCPU(duracion time.Duration) {
	// Simular trabajo que use CPU
}

// TODO: Implementar función para simular trabajo I/O
func simularTrabajoIO(duracion time.Duration) {
	// Simular operación de I/O
}

// TODO: Implementar función para generar números aleatorios
func generarNumeroAleatorio(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// TODO: Implementar función para verificar si un número es primo
func esPrimo(n int) bool {
	// Implementar algoritmo para verificar primos
	return false
}

// TODO: Implementar función para calcular estadísticas de memoria
func obtenerEstadisticasMemoria() (uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc, m.Sys
}
