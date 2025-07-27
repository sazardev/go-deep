// Ejemplos detallados de Patrones de Concurrencia
package concurrencypatterns

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ===============================================
// Ejemplo 1: Worker Pool Pattern
// ===============================================

// Job representa un trabajo a procesar
type Job struct {
	ID   int
	Data int
}

// Result representa el resultado de un trabajo
type Result struct {
	JobID int
	Value int
	Error error
}

// WorkerPoolExample demuestra el patrón worker pool
func WorkerPoolExample() {
	fmt.Println("=== Worker Pool Pattern ===")

	// Crear canales
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)

	// Número de workers
	numWorkers := 3

	// Iniciar workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Enviar trabajos
	numJobs := 10
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j, Data: j * 2}
		}
		close(jobs)
	}()

	// Recolectar resultados en una goroutine separada
	go func() {
		wg.Wait()
		close(results)
	}()

	// Procesar resultados
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Job %d error: %v\n", result.JobID, result.Error)
		} else {
			fmt.Printf("Job %d result: %d\n", result.JobID, result.Value)
		}
	}
}

// worker procesa trabajos del canal
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		// Simular trabajo
		time.Sleep(100 * time.Millisecond)

		// Procesar (ejemplo: duplicar el valor)
		result := Result{
			JobID: job.ID,
			Value: job.Data * 2,
		}

		results <- result
	}
	fmt.Printf("Worker %d finished\n", id)
}

// ===============================================
// Ejemplo 2: Pipeline Pattern
// ===============================================

// PipelineExample demuestra el patrón pipeline
func PipelineExample() {
	fmt.Println("\n=== Pipeline Pattern ===")

	// Pipeline: números -> duplicar -> sumar 1 -> imprimir
	numbers := []int{1, 2, 3, 4, 5}

	// Etapa 1: generar números
	numbersChan := generator(numbers)

	// Etapa 2: duplicar
	doubledChan := doubler(numbersChan)

	// Etapa 3: sumar 1
	incrementedChan := incrementer(doubledChan)

	// Etapa 4: consumir resultados
	for result := range incrementedChan {
		fmt.Printf("Pipeline result: %d\n", result)
	}
}

// generator envía números a un canal
func generator(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			out <- n
		}
	}()
	return out
}

// doubler duplica números del canal de entrada
func doubler(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// incrementer suma 1 a números del canal de entrada
func incrementer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 1
		}
	}()
	return out
}

// ===============================================
// Ejemplo 3: Fan-Out / Fan-In Pattern
// ===============================================

// FanOutFanInExample demuestra distribuir trabajo y recolectar resultados
func FanOutFanInExample() {
	fmt.Println("\n=== Fan-Out / Fan-In Pattern ===")

	// Datos de entrada
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Fan-out: distribuir trabajo a múltiples workers
	in := make(chan int, len(input))
	for _, v := range input {
		in <- v
	}
	close(in)

	// Crear múltiples workers (fan-out)
	numWorkers := 3
	workers := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = squareWorker(in)
	}

	// Fan-in: recolectar resultados de todos los workers
	results := fanIn(workers...)

	// Procesar resultados
	for result := range results {
		fmt.Printf("Fan-out/Fan-in result: %d\n", result)
	}
}

// squareWorker eleva al cuadrado números del canal
func squareWorker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			// Simular trabajo pesado
			time.Sleep(50 * time.Millisecond)
			out <- n * n
		}
	}()
	return out
}

// fanIn combina múltiples canales en uno solo
func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Función para copiar de input a output
	multiplex := func(input <-chan int) {
		defer wg.Done()
		for value := range input {
			out <- value
		}
	}

	// Iniciar una goroutine por cada canal de entrada
	wg.Add(len(inputs))
	for _, input := range inputs {
		go multiplex(input)
	}

	// Cerrar el canal de salida cuando todos terminen
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// ===============================================
// Ejemplo 4: Rate Limiting / Throttling
// ===============================================

// RateLimitingExample demuestra limitación de tasa
func RateLimitingExample() {
	fmt.Println("\n=== Rate Limiting / Throttling ===")

	// Limitar a 2 operaciones concurrentes máximo
	semaphore := make(chan struct{}, 2)

	// Trabajos a procesar
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go func(jobID int) {
			defer wg.Done()

			// Adquirir semáforo (throttling)
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Liberar semáforo

			fmt.Printf("Processing job %d (goroutines: %d)\n", jobID, runtime.NumGoroutine())

			// Simular trabajo
			time.Sleep(200 * time.Millisecond)

			fmt.Printf("Job %d completed\n", jobID)
		}(job)
	}

	wg.Wait()
}

// ===============================================
// Ejemplo 5: Graceful Shutdown
// ===============================================

// GracefulShutdownExample demuestra parada elegante
func GracefulShutdownExample() {
	fmt.Println("\n=== Graceful Shutdown ===")

	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Canal para trabajos
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Enviar trabajos
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i:
				fmt.Printf("Sent job %d\n", i)
			case <-ctx.Done():
				fmt.Println("Stopped sending jobs due to context cancellation")
				return
			}
		}
	}()

	// Worker que respeta el contexto
	go func() {
		defer close(results)
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					fmt.Println("No more jobs, worker stopping")
					return
				}

				// Simular trabajo
				select {
				case <-time.After(500 * time.Millisecond):
					results <- job * 2
					fmt.Printf("Processed job %d\n", job)
				case <-ctx.Done():
					fmt.Printf("Job %d cancelled due to context\n", job)
					return
				}

			case <-ctx.Done():
				fmt.Println("Worker stopping due to context cancellation")
				return
			}
		}
	}()

	// Recolectar resultados hasta que se cancele el contexto
	for {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("Results channel closed")
				return
			}
			fmt.Printf("Got result: %d\n", result)
		case <-ctx.Done():
			fmt.Printf("Shutdown complete: %v\n", ctx.Err())
			return
		}
	}
}

// ===============================================
// Ejemplo 6: Select con Timeout
// ===============================================

// TimeoutExample demuestra uso de select con timeout
func TimeoutExample() {
	fmt.Println("\n=== Select with Timeout ===")

	// Canal que podría tardar mucho
	slowChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		slowChan <- "slow result"
	}()

	// Select con timeout
	select {
	case result := <-slowChan:
		fmt.Printf("Received: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Operation took too long")
	}
}

// ===============================================
// Función principal para ejecutar ejemplos
// ===============================================

func RunAllExamples() {
	fmt.Println("🚀 Ejemplos de Patrones de Concurrencia")
	fmt.Println("========================================")

	WorkerPoolExample()
	PipelineExample()
	FanOutFanInExample()
	RateLimitingExample()
	GracefulShutdownExample()
	TimeoutExample()

	fmt.Println("\n✅ Todos los ejemplos completados!")
}
