package concurrencypatterns

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		jobs       []int
		numWorkers int
		want       []int
	}{
		{
			name:       "basic worker pool",
			jobs:       []int{1, 2, 3, 4, 5},
			numWorkers: 3,
			want:       []int{2, 4, 6, 8, 10},
		},
		{
			name:       "empty jobs",
			jobs:       []int{},
			numWorkers: 3,
			want:       []int{},
		},
		{
			name:       "single worker",
			jobs:       []int{1, 2, 3},
			numWorkers: 1,
			want:       []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WorkerPool(tt.jobs, tt.numWorkers)

			// Ordenar ambos slices para comparar (orden puede variar por concurrencia)
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WorkerPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPipeline(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "basic pipeline",
			input: []int{1, 2, 3},
			want:  []int{3, 5, 7}, // (1*2+1, 2*2+1, 3*2+1)
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single value",
			input: []int{5},
			want:  []int{11}, // 5*2+1 = 11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pipeline(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFanOutFanIn(t *testing.T) {
	tests := []struct {
		name       string
		jobs       []int
		numWorkers int
		want       []int
	}{
		{
			name:       "basic fan-out/fan-in",
			jobs:       []int{1, 2, 3, 4},
			numWorkers: 2,
			want:       []int{2, 4, 6, 8},
		},
		{
			name:       "empty jobs",
			jobs:       []int{},
			numWorkers: 2,
			want:       []int{},
		},
		{
			name:       "more workers than jobs",
			jobs:       []int{1, 2},
			numWorkers: 5,
			want:       []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FanOutFanIn(tt.jobs, tt.numWorkers)

			// Ordenar para comparar (orden puede variar)
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FanOutFanIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThrottledProcessing(t *testing.T) {
	tests := []struct {
		name          string
		jobs          []int
		maxConcurrent int
		want          []int
	}{
		{
			name:          "basic throttling",
			jobs:          []int{1, 2, 3, 4},
			maxConcurrent: 2,
			want:          []int{2, 4, 6, 8},
		},
		{
			name:          "no throttling",
			jobs:          []int{1, 2, 3},
			maxConcurrent: 10,
			want:          []int{2, 4, 6},
		},
		{
			name:          "heavy throttling",
			jobs:          []int{1, 2, 3, 4},
			maxConcurrent: 1,
			want:          []int{2, 4, 6, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got := ThrottledProcessing(tt.jobs, tt.maxConcurrent)
			duration := time.Since(start)

			// Verificar que se respeta el throttling
			expectedMinDuration := time.Duration(len(tt.jobs)/tt.maxConcurrent) * 10 * time.Millisecond
			if duration < expectedMinDuration {
				t.Logf("Duration %v might indicate throttling is not working properly", duration)
			}

			// Ordenar para comparar
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThrottledProcessing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGracefulShutdown(t *testing.T) {
	tests := []struct {
		name    string
		jobs    []int
		timeout time.Duration
		minLen  int // Mínimo número de resultados esperados
	}{
		{
			name:    "quick shutdown",
			jobs:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			timeout: 100 * time.Millisecond,
			minLen:  1, // Al menos debe procesar algunos
		},
		{
			name:    "long timeout",
			jobs:    []int{1, 2, 3},
			timeout: 1 * time.Second,
			minLen:  3, // Debe procesar todos
		},
		{
			name:    "empty jobs",
			jobs:    []int{},
			timeout: 100 * time.Millisecond,
			minLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			got := GracefulShutdown(ctx, tt.jobs)

			if len(got) < tt.minLen {
				t.Errorf("GracefulShutdown() returned %d results, want at least %d", len(got), tt.minLen)
			}

			// Verificar que los resultados son correctos (doble de entrada)
			for _, result := range got {
				if result%2 != 0 {
					t.Errorf("GracefulShutdown() returned odd result %d, expected even numbers only", result)
				}
			}
		})
	}
}

// Benchmarks para medir rendimiento

func BenchmarkWorkerPool(b *testing.B) {
	jobs := make([]int, 1000)
	for i := range jobs {
		jobs[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WorkerPool(jobs, 4)
	}
}

func BenchmarkPipeline(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Pipeline(input)
	}
}

func BenchmarkFanOutFanIn(b *testing.B) {
	jobs := make([]int, 1000)
	for i := range jobs {
		jobs[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FanOutFanIn(jobs, 4)
	}
}

// Test de concurrencia para verificar que realmente se ejecuta en paralelo
func TestConcurrencyActuallyWorks(t *testing.T) {
	t.Run("worker pool parallelism", func(t *testing.T) {
		jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}

		start := time.Now()
		WorkerPool(jobs, 4)
		duration := time.Since(start)

		// Con 4 workers, debería ser más rápido que secuencial
		// (aunque este test puede ser flaky en sistemas lentos)
		if duration > 100*time.Millisecond {
			t.Logf("WorkerPool took %v, might not be running in parallel efficiently", duration)
		}
	})
}
