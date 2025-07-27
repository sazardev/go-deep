// Lección 18: Patrones Avanzados de Concurrencia en Go
// Archivo de ejercicios prácticos

package concurrencypatterns

import (
	"context"
	"sync"
	"time"
)

// Ejercicio 1: Implementa un Worker Pool
// Recibe un slice de trabajos (int) y un número de workers.
// Devuelve un slice con los resultados procesados (por ejemplo, el doble de cada trabajo).
func WorkerPool(jobs []int, numWorkers int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	// Canales para distribuir trabajos y recolectar resultados
	jobChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))

	// Enviar trabajos al canal
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	// Iniciar workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobChan {
				// Procesar: duplicar el valor
				resultChan <- job * 2
			}
		}()
	}

	// Cerrar canal de resultados cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Recolectar resultados
	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// Ejercicio 2: Implementa un Pipeline de procesamiento
// Recibe un slice de enteros y aplica dos etapas: multiplicar por 2 y luego sumar 1.
// Devuelve el resultado final como slice.
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

	// Recolectar resultados
	var results []int
	for result := range stage3 {
		results = append(results, result)
	}

	return results
}

// Ejercicio 3: Fan-Out / Fan-In
// Procesa trabajos en paralelo y recolecta los resultados en un solo canal.
func FanOutFanIn(jobs []int, numWorkers int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	// Canal de entrada
	jobChan := make(chan int, len(jobs))
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	// Fan-out: crear múltiples workers
	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker := make(chan int)
		workers[i] = worker

		go func(out chan<- int) {
			defer close(out)
			for job := range jobChan {
				// Procesar: duplicar
				out <- job * 2
			}
		}(worker)
	}

	// Fan-in: combinar resultados
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

	// Recolectar resultados
	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// Ejercicio 4: Throttling & Rate Limiting
// Limita la cantidad de trabajos concurrentes usando un canal de semáforo.
func ThrottledProcessing(jobs []int, maxConcurrent int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	// Semáforo para limitar concurrencia
	semaphore := make(chan struct{}, maxConcurrent)
	resultChan := make(chan int, len(jobs))

	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			// Adquirir semáforo
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Simular trabajo
			time.Sleep(10 * time.Millisecond)

			// Procesar
			resultChan <- j * 2
		}(job)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Recolectar resultados
	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// Ejercicio 5: Graceful Shutdown
// Procesa trabajos hasta que el contexto se cancele.
func GracefulShutdown(ctx context.Context, jobs []int) []int {
	if len(jobs) == 0 {
		return []int{}
	}

	jobChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))

	// Enviar trabajos respetando el contexto
	go func() {
		defer close(jobChan)
		for _, job := range jobs {
			select {
			case jobChan <- job:
				// Trabajo enviado
			case <-ctx.Done():
				// Contexto cancelado, detener envío
				return
			}
		}
	}()

	// Worker que respeta el contexto
	go func() {
		defer close(resultChan)
		for {
			select {
			case job, ok := <-jobChan:
				if !ok {
					return
				}

				// Procesar con timeout
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

	// Recolectar resultados hasta cancelación
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
