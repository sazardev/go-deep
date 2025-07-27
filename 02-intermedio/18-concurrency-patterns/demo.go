package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ===============================================
// Patrones de Concurrencia Implementados
// ===============================================

// WorkerPool procesa trabajos usando múltiples workers
func WorkerPool(jobs []int, numWorkers int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	jobChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))

	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobChan {
				resultChan <- job * 2
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// Pipeline procesa datos en etapas secuenciales
func Pipeline(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	// Etapa 1: generar números
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for _, n := range input {
			stage1 <- n
		}
	}()

	// Etapa 2: multiplicar por 2
	stage2 := make(chan int)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			stage2 <- n * 2
		}
	}()

	// Etapa 3: sumar 1
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			stage3 <- n + 1
		}
	}()

	var results []int
	for result := range stage3 {
		results = append(results, result)
	}

	return results
}

// FanOutFanIn distribuye trabajo y recolecta resultados
func FanOutFanIn(jobs []int, numWorkers int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	jobChan := make(chan int, len(jobs))
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker := make(chan int)
		workers[i] = worker

		go func(out chan<- int) {
			defer close(out)
			for job := range jobChan {
				out <- job * 2
			}
		}(worker)
	}

	resultChan := make(chan int)
	var wg sync.WaitGroup

	for _, worker := range workers {
		wg.Add(1)
		go func(input <-chan int) {
			defer wg.Done()
			for result := range input {
				resultChan <- result
			}
		}(worker)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// ThrottledProcessing limita la concurrencia
func ThrottledProcessing(jobs []int, maxConcurrent int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	semaphore := make(chan struct{}, maxConcurrent)
	resultChan := make(chan int, len(jobs))

	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			time.Sleep(10 * time.Millisecond)
			resultChan <- j * 2
		}(job)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// GracefulShutdown procesa con cancelación elegante
func GracefulShutdown(ctx context.Context, jobs []int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	jobChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))

	go func() {
		defer close(jobChan)
		for _, job := range jobs {
			select {
			case jobChan <- job:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		defer close(resultChan)
		for {
			select {
			case job, ok := <-jobChan:
				if !ok {
					return
				}

				select {
				case <-time.After(50 * time.Millisecond):
					resultChan <- job * 2
				case <-ctx.Done():
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	var results []int
	for {
		select {
		case result, ok := <-resultChan:
			if !ok {
				return results
			}
			results = append(results, result)
		case <-ctx.Done():
			return results
		}
	}
}

// ===============================================
// Demos Interactivos
// ===============================================

func main() {
	fmt.Println("🚀 Lección 18: Patrones Avanzados de Concurrencia")
	fmt.Println("=================================================")
	fmt.Println()

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
