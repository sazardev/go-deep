// ==============================================
// LECCIÓN 13: Goroutines y Concurrencia - Soluciones
// ==============================================

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ==============================================
// EJERCICIO 1: Primera Goroutine - SOLUCIÓN
// ==============================================

func imprimirNumeros(nombre string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("👤 %s: número %d\n", nombre, i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("✅ %s terminó\n", nombre)
}

func solucion1() {
	fmt.Println("🚀 Ejercicio 1: Primera Goroutine - SOLUCIÓN")
	fmt.Println("==============================================")

	// Ejecución secuencial
	fmt.Println("📝 Ejecución Secuencial:")
	start := time.Now()
	imprimirNumeros("Ana")
	imprimirNumeros("Bob")
	tiempoSecuencial := time.Since(start)
	fmt.Printf("⏱️ Tiempo secuencial: %v\n\n", tiempoSecuencial)

	// Ejecución concurrente
	fmt.Println("🚀 Ejecución Concurrente:")
	start = time.Now()
	go imprimirNumeros("Carlos")
	go imprimirNumeros("Diana")

	// Esperar suficiente tiempo para que terminen
	time.Sleep(600 * time.Millisecond)
	tiempoConcurrente := time.Since(start)
	fmt.Printf("⏱️ Tiempo concurrente: %v\n", tiempoConcurrente)
	fmt.Printf("🚀 Aceleración: %.2fx más rápido\n\n", float64(tiempoSecuencial)/float64(tiempoConcurrente))
}

// ==============================================
// EJERCICIO 2: WaitGroup Básico - SOLUCIÓN
// ==============================================

func trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ¡Crucial! Marcar como completado al salir

	fmt.Printf("👷 Trabajador %d iniciando\n", id)

	// Simular trabajo con tiempo aleatorio
	tiempoTrabajo := time.Duration(rand.Intn(400)+100) * time.Millisecond
	time.Sleep(tiempoTrabajo)

	fmt.Printf("✅ Trabajador %d terminó (trabajó %v)\n", id, tiempoTrabajo)
}

func solucion2() {
	fmt.Println("🔄 Ejercicio 2: WaitGroup Básico - SOLUCIÓN")
	fmt.Println("============================================")

	var wg sync.WaitGroup
	const numTrabajadores = 5

	fmt.Printf("🏗️ Lanzando %d trabajadores...\n", numTrabajadores)

	for i := 1; i <= numTrabajadores; i++ {
		wg.Add(1) // Incrementar contador antes de lanzar goroutine
		go trabajador(i, &wg)
	}

	fmt.Println("⏳ Esperando que todos terminen...")
	wg.Wait() // Bloquear hasta que todos hagan wg.Done()

	fmt.Println("🎉 ¡Todos los trabajadores terminaron!\n")
}

// ==============================================
// EJERCICIO 3: Race Condition - SOLUCIÓN
// ==============================================

var contadorRace int

func incrementarInseguro(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		// ❌ RACE CONDITION: múltiples goroutines accediendo simultáneamente
		temp := contadorRace
		temp++
		contadorRace = temp
	}
}

func solucion3() {
	fmt.Println("⚠️ Ejercicio 3: Race Condition - SOLUCIÓN")
	fmt.Println("==========================================")

	var wg sync.WaitGroup
	const numGoroutines = 10

	// Reiniciar contador
	contadorRace = 0

	fmt.Printf("🔢 Lanzando %d goroutines que incrementan 1000 veces cada una\n", numGoroutines)
	fmt.Printf("📊 Resultado esperado: %d\n", numGoroutines*1000)

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarInseguro(&wg)
	}

	wg.Wait()
	duracion := time.Since(start)

	fmt.Printf("📊 Resultado real: %d\n", contadorRace)
	fmt.Printf("⏱️ Tiempo: %v\n", duracion)

	if contadorRace != numGoroutines*1000 {
		fmt.Printf("❌ ¡Race condition detectada! Perdimos %d incrementos\n",
			numGoroutines*1000-contadorRace)
	} else {
		fmt.Println("✅ Por suerte no hubo race condition esta vez")
	}

	fmt.Println("💡 Ejecuta varias veces para ver resultados inconsistentes")
	fmt.Println("🔍 Usa 'go run -race soluciones.go' para detectar automáticamente\n")
}

// ==============================================
// EJERCICIO 4: Mutex para Sincronización - SOLUCIÓN
// ==============================================

var (
	contadorMutex int
	mutexLock     sync.Mutex
)

func incrementarSeguro(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mutexLock.Lock() // 🔒 Bloquear acceso exclusivo
		contadorMutex++
		mutexLock.Unlock() // 🔓 Liberar acceso
	}
}

func solucion4() {
	fmt.Println("🔒 Ejercicio 4: Mutex para Sincronización - SOLUCIÓN")
	fmt.Println("====================================================")

	var wg sync.WaitGroup
	const numGoroutines = 10

	// Reiniciar contador
	contadorMutex = 0

	fmt.Printf("🔐 Lanzando %d goroutines con protección Mutex\n", numGoroutines)

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarSeguro(&wg)
	}

	wg.Wait()
	duracion := time.Since(start)

	fmt.Printf("📊 Resultado: %d (esperado: %d)\n", contadorMutex, numGoroutines*1000)
	fmt.Printf("⏱️ Tiempo: %v\n", duracion)

	if contadorMutex == numGoroutines*1000 {
		fmt.Println("✅ ¡Perfecto! Mutex eliminó la race condition")
	} else {
		fmt.Println("❌ Algo salió mal...")
	}
	fmt.Println()
}

// ==============================================
// EJERCICIO 5: Atomic Operations - SOLUCIÓN
// ==============================================

var contadorAtomic int64

func incrementarAtomico(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&contadorAtomic, 1) // ⚡ Operación atómica
	}
}

func solucion5() {
	fmt.Println("⚡ Ejercicio 5: Atomic Operations - SOLUCIÓN")
	fmt.Println("=============================================")

	var wg sync.WaitGroup
	const numGoroutines = 10

	// Reiniciar contador
	atomic.StoreInt64(&contadorAtomic, 0)

	fmt.Printf("⚡ Lanzando %d goroutines con operaciones atómicas\n", numGoroutines)

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarAtomico(&wg)
	}

	wg.Wait()
	duracion := time.Since(start)

	resultado := atomic.LoadInt64(&contadorAtomic)
	fmt.Printf("📊 Resultado: %d (esperado: %d)\n", resultado, numGoroutines*1000)
	fmt.Printf("⏱️ Tiempo: %v\n", duracion)

	if resultado == int64(numGoroutines*1000) {
		fmt.Println("✅ ¡Excelente! Operaciones atómicas = thread-safe + performance")
	}
	fmt.Println("💡 Las operaciones atómicas son más rápidas que Mutex para operaciones simples\n")
}

// ==============================================
// EJERCICIO 6: Worker Pool - SOLUCIÓN
// ==============================================

type TrabajoWP struct {
	ID   int
	Dato string
}

type ResultadoWP struct {
	TrabajoID int
	Resultado string
	Error     error
}

func worker(id int, trabajos <-chan TrabajoWP, resultados chan<- ResultadoWP, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("👷 Worker %d iniciado\n", id)

	for trabajo := range trabajos {
		fmt.Printf("👷 Worker %d procesando trabajo %d\n", id, trabajo.ID)

		// Simular procesamiento
		tiempoTrabajo := time.Duration(rand.Intn(200)+50) * time.Millisecond
		time.Sleep(tiempoTrabajo)

		resultado := ResultadoWP{
			TrabajoID: trabajo.ID,
			Resultado: fmt.Sprintf("Procesado por worker %d en %v", id, tiempoTrabajo),
			Error:     nil,
		}

		resultados <- resultado
	}

	fmt.Printf("👷 Worker %d terminando\n", id)
}

func solucion6() {
	fmt.Println("🏗️ Ejercicio 6: Worker Pool - SOLUCIÓN")
	fmt.Println("======================================")

	const numWorkers = 3
	const numTrabajos = 15

	// Crear canales con buffer adecuado
	trabajos := make(chan TrabajoWP, numTrabajos)
	resultados := make(chan ResultadoWP, numTrabajos)

	var wg sync.WaitGroup

	// Lanzar workers
	fmt.Printf("🚀 Lanzando %d workers\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, trabajos, resultados, &wg)
	}

	// Enviar trabajos
	fmt.Printf("📤 Enviando %d trabajos\n", numTrabajos)
	for i := 1; i <= numTrabajos; i++ {
		trabajo := TrabajoWP{
			ID:   i,
			Dato: fmt.Sprintf("Datos del trabajo %d", i),
		}
		trabajos <- trabajo
	}
	close(trabajos) // ¡Importante! Cerrar para que workers terminen

	// Cerrar canal de resultados cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Recoger resultados
	fmt.Println("\n📊 Resultados:")
	contador := 0
	for resultado := range resultados {
		contador++
		fmt.Printf("  ✅ Trabajo %d: %s\n", resultado.TrabajoID, resultado.Resultado)
	}

	fmt.Printf("\n🎉 Procesados %d trabajos usando %d workers\n", contador, numWorkers)
	fmt.Println("💡 Patrón Worker Pool: eficiente para procesar muchas tareas similares\n")
}

// ==============================================
// EJERCICIO 7: Pipeline de Procesamiento - SOLUCIÓN
// ==============================================

func generarNumeros(numeros chan<- int) {
	defer close(numeros)

	fmt.Println("🔢 Generando números del 1 al 20...")
	for i := 1; i <= 20; i++ {
		numeros <- i
		time.Sleep(10 * time.Millisecond) // Simular generación gradual
	}
	fmt.Println("✅ Generación completada")
}

func elevarCuadrado(numeros <-chan int, cuadrados chan<- int) {
	defer close(cuadrados)

	fmt.Println("📐 Elevando números al cuadrado...")
	for num := range numeros {
		cuadrado := num * num
		fmt.Printf("  %d² = %d\n", num, cuadrado)
		cuadrados <- cuadrado
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("✅ Elevación completada")
}

func filtrarPares(cuadrados <-chan int, pares chan<- int) {
	defer close(pares)

	fmt.Println("🔍 Filtrando números pares...")
	for cuadrado := range cuadrados {
		if cuadrado%2 == 0 {
			fmt.Printf("  %d es par ✓\n", cuadrado)
			pares <- cuadrado
		}
	}
	fmt.Println("✅ Filtrado completado")
}

func sumarTodos(pares <-chan int, resultado chan<- int) {
	defer close(resultado)

	fmt.Println("➕ Sumando números pares...")
	suma := 0
	count := 0
	for par := range pares {
		suma += par
		count++
		fmt.Printf("  Sumando %d (total parcial: %d)\n", par, suma)
	}

	fmt.Printf("✅ Suma completada: %d números, total = %d\n", count, suma)
	resultado <- suma
}

func solucion7() {
	fmt.Println("🔄 Ejercicio 7: Pipeline de Procesamiento - SOLUCIÓN")
	fmt.Println("===================================================")

	// Crear canales para el pipeline
	numeros := make(chan int)
	cuadrados := make(chan int)
	pares := make(chan int)
	resultado := make(chan int)

	fmt.Println("🏭 Iniciando pipeline de 4 etapas:")
	fmt.Println("   1. Generar números")
	fmt.Println("   2. Elevar al cuadrado")
	fmt.Println("   3. Filtrar pares")
	fmt.Println("   4. Sumar todos\n")

	// Lanzar pipeline: cada etapa es una goroutine
	go generarNumeros(numeros)
	go elevarCuadrado(numeros, cuadrados)
	go filtrarPares(cuadrados, pares)
	go sumarTodos(pares, resultado)

	// Obtener resultado final
	suma := <-resultado

	fmt.Printf("\n🎯 Resultado final del pipeline: %d\n", suma)
	fmt.Println("💡 Pipeline Pattern: procesa datos en etapas concurrentes\n")
}

// ==============================================
// EJERCICIO 8: Context y Cancelación - SOLUCIÓN
// ==============================================

func trabajadorConContext(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("👷 Worker %d iniciado con context\n", id)

	contador := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 Worker %d cancelado después de %d iteraciones: %v\n",
				id, contador, ctx.Err())
			return
		default:
			// Simular trabajo
			contador++
			fmt.Printf("👷 Worker %d trabajando... (iteración %d)\n", id, contador)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func solucion8() {
	fmt.Println("🎯 Ejercicio 8: Context y Cancelación - SOLUCIÓN")
	fmt.Println("================================================")

	// Crear context con timeout de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Liberar recursos del context

	var wg sync.WaitGroup
	const numWorkers = 3

	fmt.Printf("⏰ Lanzando %d workers con timeout de 2 segundos\n\n", numWorkers)

	// Lanzar workers con context
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go trabajadorConContext(ctx, i, &wg)
	}

	// Esperar a que todos terminen (por timeout o completación)
	wg.Wait()

	fmt.Println("\n✅ Todos los workers terminaron")
	fmt.Println("💡 Context Pattern: control elegante del ciclo de vida de goroutines\n")
}

// ==============================================
// EJERCICIO 9: Fan-Out/Fan-In - SOLUCIÓN
// ==============================================

func generarNumerosPrimos(numeros chan<- int) {
	defer close(numeros)

	fmt.Println("🔢 Generando números del 1 al 50 para verificar primos...")
	for i := 1; i <= 50; i++ {
		numeros <- i
	}
}

func verificarPrimos(id int, numeros <-chan int, primos chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("🔍 Worker %d iniciado para verificar primos\n", id)

	for num := range numeros {
		if esPrimoSol(num) {
			fmt.Printf("✨ Worker %d encontró primo: %d\n", id, num)
			primos <- num
		}
	}

	fmt.Printf("✅ Worker %d terminó\n", id)
}

func solucion9() {
	fmt.Println("📊 Ejercicio 9: Fan-Out/Fan-In - SOLUCIÓN")
	fmt.Println("=========================================")

	const numWorkers = 5

	numeros := make(chan int)
	primos := make(chan int, 50) // Buffer suficiente

	var wg sync.WaitGroup

	fmt.Printf("🚀 Fan-Out: distribuyendo trabajo a %d workers\n", numWorkers)

	// Fan-Out: lanzar múltiples workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go verificarPrimos(i, numeros, primos, &wg)
	}

	// Generar números
	go generarNumerosPrimos(numeros)

	// Cerrar canal de primos cuando todos terminen
	go func() {
		wg.Wait()
		close(primos)
	}()

	// Fan-In: recoger todos los resultados
	fmt.Println("\n🔍 Fan-In: recolectando números primos encontrados:")
	var primosEncontrados []int
	for primo := range primos {
		primosEncontrados = append(primosEncontrados, primo)
	}

	fmt.Printf("\n🎯 Números primos encontrados (%d total): %v\n",
		len(primosEncontrados), primosEncontrados)
	fmt.Println("💡 Fan-Out/Fan-In: distribuir trabajo y agregar resultados\n")
}

// ==============================================
// EJERCICIO 10: Monitoreo y Estadísticas - SOLUCIÓN
// ==============================================

func monitorearSistema(duracion time.Duration) {
	fmt.Println("📊 Iniciando monitoreo del sistema...")

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(duracion)

	for {
		select {
		case <-timeout:
			fmt.Println("⏰ Monitoreo completado")
			return
		case <-ticker.C:
			// Obtener estadísticas
			numGoroutines := runtime.NumGoroutine()
			numCPU := runtime.NumCPU()

			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Printf("📊 Goroutines: %d | CPUs: %d | Memoria: %.2f MB\n",
				numGoroutines, numCPU, float64(m.Alloc)/1024/1024)
		}
	}
}

func generarCargaTrabajo(wg *sync.WaitGroup) {
	defer wg.Done()

	var localWg sync.WaitGroup

	// Crear 20 goroutines que trabajen un tiempo aleatorio
	for i := 0; i < 20; i++ {
		localWg.Add(1)
		go func(id int) {
			defer localWg.Done()

			duracion := time.Duration(rand.Intn(2000)+500) * time.Millisecond
			time.Sleep(duracion)

			// Simular trabajo CPU intensivo
			sum := 0
			for j := 0; j < 1000000; j++ {
				sum += j
			}
		}(i)
	}

	localWg.Wait()
}

func solucion10() {
	fmt.Println("📊 Ejercicio 10: Monitoreo y Estadísticas - SOLUCIÓN")
	fmt.Println("====================================================")

	// Iniciar monitoreo en background
	go monitorearSistema(3 * time.Second)

	var wg sync.WaitGroup

	fmt.Println("🏗️ Generando carga de trabajo con múltiples oleadas de goroutines...")

	// Generar 3 oleadas de carga de trabajo
	for oleada := 1; oleada <= 3; oleada++ {
		fmt.Printf("\n🌊 Oleada %d: lanzando 20 goroutines\n", oleada)
		wg.Add(1)
		go generarCargaTrabajo(&wg)

		time.Sleep(800 * time.Millisecond) // Espacio entre oleadas
	}

	fmt.Println("\n⏳ Esperando que termine toda la carga de trabajo...")
	wg.Wait()

	// Dar tiempo para que el monitoreo termine
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n✅ Todas las tareas completadas")
	fmt.Println("💡 Observa cómo varían las métricas con la carga de trabajo\n")
}

// ==============================================
// FUNCIONES AUXILIARES
// ==============================================

func esPrimoSol(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func ejecutarSoluciones() {
	fmt.Println("🚀 SOLUCIONES: Goroutines y Concurrencia")
	fmt.Println("=========================================")
	fmt.Printf("Usando Go %s con %d CPUs\n", runtime.Version(), runtime.NumCPU())
	fmt.Printf("Iniciando con %d goroutines activas\n\n", runtime.NumGoroutine())

	// Inicializar semilla para números aleatorios
	rand.Seed(time.Now().UnixNano())

	soluciones := []struct {
		nombre  string
		funcion func()
	}{
		{"Primera Goroutine", solucion1},
		{"WaitGroup Básico", solucion2},
		{"Race Condition", solucion3},
		{"Mutex para Sincronización", solucion4},
		{"Atomic Operations", solucion5},
		{"Worker Pool", solucion6},
		{"Pipeline de Procesamiento", solucion7},
		{"Context y Cancelación", solucion8},
		{"Fan-Out/Fan-In", solucion9},
		{"Monitoreo y Estadísticas", solucion10},
	}

	for i, solucion := range soluciones {
		fmt.Printf("📝 Ejecutando solución %d: %s\n", i+1, solucion.nombre)
		fmt.Println(strings.Repeat("=", 50))

		inicio := time.Now()
		solucion.funcion()
		duracion := time.Since(inicio)

		fmt.Printf("⏱️ Tiempo de ejecución: %v\n", duracion)
		fmt.Printf("📊 Goroutines activas: %d\n", runtime.NumGoroutine())

		if i < len(soluciones)-1 {
			fmt.Println("\n" + strings.Repeat("-", 60) + "\n")
			time.Sleep(300 * time.Millisecond) // Pausa entre ejercicios
		}
	}

	fmt.Println("\n🎉 ¡Todas las soluciones ejecutadas exitosamente!")
	fmt.Println("\n🎓 Conceptos dominados:")
	fmt.Println("   ✅ Goroutines básicas")
	fmt.Println("   ✅ WaitGroup para sincronización")
	fmt.Println("   ✅ Race conditions y su detección")
	fmt.Println("   ✅ Mutex para protección")
	fmt.Println("   ✅ Operaciones atómicas")
	fmt.Println("   ✅ Worker Pool pattern")
	fmt.Println("   ✅ Pipeline pattern")
	fmt.Println("   ✅ Context para control")
	fmt.Println("   ✅ Fan-Out/Fan-In pattern")
	fmt.Println("   ✅ Monitoreo de rendimiento")

	fmt.Println("\n💡 Próximo paso: Lección 14 - Channels")
	fmt.Println("   🔗 Comunicación entre goroutines")
	fmt.Println("   📡 Select statement")
	fmt.Println("   🎯 Patrones avanzados con channels")
}

func main() {
	ejecutarSoluciones()
}
