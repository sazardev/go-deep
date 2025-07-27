package concurrencypatterns

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Función principal para demostrar los patrones
func RunDemos() {
	fmt.Println("🚀 Lección 18: Patrones Avanzados de Concurrencia")
	fmt.Println("=================================================")
	fmt.Println()

	// Ejecutar ejemplos interactivos
	demoWorkerPool()
	demoPipeline()
	demoFanOutFanIn()
	demoThrottling()
	demoGracefulShutdown()

	fmt.Println("\n✅ Todos los demos completados!")
	fmt.Println("\n💡 Ejecuta 'go test -v' para ver los tests")
	fmt.Println("💡 Ejecuta 'go test -bench=.' para ver benchmarks")
}

func demoWorkerPool() {
	fmt.Println("=== Demo: Worker Pool ===")

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 3

	fmt.Printf("Procesando %d trabajos con %d workers...\n", len(jobs), numWorkers)

	start := time.Now()
	results := WorkerPool(jobs, numWorkers)
	duration := time.Since(start)

	fmt.Printf("Resultados: %v\n", results)
	fmt.Printf("Tiempo: %v\n", duration)
	fmt.Println()
}

func demoPipeline() {
	fmt.Println("=== Demo: Pipeline ===")

	input := []int{1, 2, 3, 4, 5}
	fmt.Printf("Entrada: %v\n", input)
	fmt.Println("Pipeline: input -> multiplicar por 2 -> sumar 1")

	start := time.Now()
	results := Pipeline(input)
	duration := time.Since(start)

	fmt.Printf("Resultados: %v\n", results)
	fmt.Printf("Tiempo: %v\n", duration)
	fmt.Println()
}

func demoFanOutFanIn() {
	fmt.Println("=== Demo: Fan-Out / Fan-In ===")

	jobs := []int{1, 2, 3, 4, 5, 6}
	numWorkers := 3

	fmt.Printf("Distribuyendo %d trabajos entre %d workers...\n", len(jobs), numWorkers)

	start := time.Now()
	results := FanOutFanIn(jobs, numWorkers)
	duration := time.Since(start)

	fmt.Printf("Resultados: %v\n", results)
	fmt.Printf("Tiempo: %v\n", duration)
	fmt.Println()
}

func demoThrottling() {
	fmt.Println("=== Demo: Throttling ===")

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	maxConcurrent := 2

	fmt.Printf("Procesando %d trabajos con máximo %d concurrentes...\n",
		len(jobs), maxConcurrent)

	start := time.Now()
	results := ThrottledProcessing(jobs, maxConcurrent)
	duration := time.Since(start)

	fmt.Printf("Resultados: %v\n", results)
	fmt.Printf("Tiempo: %v (debería ser más lento por throttling)\n", duration)
	fmt.Println()
}

func demoGracefulShutdown() {
	fmt.Println("=== Demo: Graceful Shutdown ===")
	fmt.Println("Presiona Ctrl+C para cancelar el procesamiento...")

	// Configurar manejo de señales
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	jobs := make([]int, 100)
	for i := range jobs {
		jobs[i] = i + 1
	}

	// Simular trabajo con timeout automático si no se interrumpe
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Goroutine para escuchar señales
	go func() {
		<-sigChan
		fmt.Println("\n🛑 Señal recibida, iniciando shutdown...")
		cancel()
	}()

	fmt.Printf("Iniciando procesamiento de %d trabajos...\n", len(jobs))
	start := time.Now()
	results := GracefulShutdown(ctx, jobs)
	duration := time.Since(start)

	fmt.Printf("Procesados: %d/%d trabajos\n", len(results), len(jobs))
	fmt.Printf("Tiempo: %v\n", duration)

	if len(results) < len(jobs) {
		fmt.Println("✅ Shutdown fue exitoso - no todos los trabajos se completaron")
	} else {
		fmt.Println("✅ Todos los trabajos se completaron antes del timeout")
	}
	fmt.Println()
}
