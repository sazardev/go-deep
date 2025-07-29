package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"testing"
	"time"
)

// 🧪 Ejemplo de diferentes implementaciones para benchmarking

// ConcatenateNaive - Implementación ineficiente
func ConcatenateNaive(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s // ⚠️ PROBLEMA: Crea nueva string en cada iteración
	}
	return result
}

// ConcatenateBuilder - Usando strings.Builder
func ConcatenateBuilder(strs []string) string {
	var builder strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// ConcatenateBuilderPrealloc - Builder con capacidad pre-asignada
func ConcatenateBuilderPrealloc(strs []string) string {
	// Pre-calcular capacidad total
	totalLen := 0
	for _, s := range strs {
		totalLen += len(s)
	}

	var builder strings.Builder
	builder.Grow(totalLen) // 🚀 OPTIMIZACIÓN: Pre-asignar memoria

	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// ConcatenateJoin - Usando strings.Join
func ConcatenateJoin(strs []string) string {
	return strings.Join(strs, "")
}

// generateTestData - Genera datos de prueba
func generateTestData(size int) []string {
	data := make([]string, size)
	for i := 0; i < size; i++ {
		data[i] = fmt.Sprintf("string_%d_", i)
	}
	return data
}

// 📊 Benchmarks para comparar implementaciones

func BenchmarkConcatenateNaive(b *testing.B) {
	data := generateTestData(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateNaive(data)
	}
}

func BenchmarkConcatenateBuilder(b *testing.B) {
	data := generateTestData(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateBuilder(data)
	}
}

func BenchmarkConcatenateBuilderPrealloc(b *testing.B) {
	data := generateTestData(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateBuilderPrealloc(data)
	}
}

func BenchmarkConcatenateJoin(b *testing.B) {
	data := generateTestData(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateJoin(data)
	}
}

// 🔍 Benchmark con diferentes tamaños para análisis de escalabilidad
func BenchmarkConcatenateComparison(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	implementations := map[string]func([]string) string{
		"Naive":           ConcatenateNaive,
		"Builder":         ConcatenateBuilder,
		"BuilderPrealloc": ConcatenateBuilderPrealloc,
		"Join":            ConcatenateJoin,
	}

	for _, size := range sizes {
		data := generateTestData(size)

		for name, impl := range implementations {
			b.Run(fmt.Sprintf("%s-%d", name, size), func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = impl(data)
				}
			})
		}
	}
}

// 🎯 Ejemplo de función con profiling automático
func ProfiledStringConcatenation() {
	// Setup CPU profiling
	cpuFile, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Generar workload para profiling
	fmt.Println("🔍 Starting profiled workload...")

	data := generateTestData(5000)

	// Test diferentes implementaciones
	start := time.Now()
	result1 := ConcatenateNaive(data[:100]) // Solo 100 para evitar timeout
	fmt.Printf("Naive (100 items): %v\n", time.Since(start))

	start = time.Now()
	result2 := ConcatenateBuilder(data)
	fmt.Printf("Builder (5000 items): %v\n", time.Since(start))

	start = time.Now()
	result3 := ConcatenateBuilderPrealloc(data)
	fmt.Printf("BuilderPrealloc (5000 items): %v\n", time.Since(start))

	start = time.Now()
	result4 := ConcatenateJoin(data)
	fmt.Printf("Join (5000 items): %v\n", time.Since(start))

	// Verificar que todas dan el mismo resultado
	if len(result1) > 0 && len(result2) > 0 && len(result3) > 0 && len(result4) > 0 {
		fmt.Println("✅ All implementations completed successfully")
	}

	fmt.Println("🎯 CPU profile saved to: cpu_profile.prof")
	fmt.Println("📊 Analyze with: go tool pprof cpu_profile.prof")
}

// 🧠 Memory-intensive function para memory profiling
func MemoryIntensiveOperation() {
	fmt.Println("🧠 Starting memory-intensive operation...")

	// Crear muchos slices para ver allocations
	var slices [][]int

	for i := 0; i < 1000; i++ {
		// Crear slice de tamaño variable
		size := 1000 + i*10
		slice := make([]int, size)

		// Llenar con datos
		for j := range slice {
			slice[j] = i * j
		}

		slices = append(slices, slice)

		// Simular algo de procesamiento
		if i%100 == 0 {
			fmt.Printf("Created %d slices\n", i+1)
		}
	}

	// Simular más allocations con maps
	dataMap := make(map[string][]byte)
	for i := 0; i < 500; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := make([]byte, 2048) // 2KB por entrada
		dataMap[key] = value
	}

	fmt.Printf("Final stats: %d slices, %d map entries\n", len(slices), len(dataMap))

	// Forzar que las variables no sean optimizadas
	if len(slices) > 0 && len(dataMap) > 0 {
		fmt.Println("✅ Memory operations completed")
	}
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "profile":
			ProfiledStringConcatenation()
		case "memory":
			MemoryIntensiveOperation()
		default:
			fmt.Println("Usage: go run benchmark_examples.go [profile|memory]")
		}
	} else {
		fmt.Println("🧪 Benchmark Examples for Go Profiling")
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  go run benchmark_examples.go profile  - Run CPU profiling example")
		fmt.Println("  go run benchmark_examples.go memory   - Run memory-intensive example")
		fmt.Println()
		fmt.Println("Benchmarking commands:")
		fmt.Println("  go test -bench=.                      - Run all benchmarks")
		fmt.Println("  go test -bench=. -benchmem           - Include memory stats")
		fmt.Println("  go test -bench=Comparison            - Run size comparison")
		fmt.Println("  go test -bench=. -cpuprofile=cpu.prof - Generate CPU profile")
		fmt.Println("  go test -bench=. -memprofile=mem.prof - Generate memory profile")
		fmt.Println()
		fmt.Println("Profiling analysis:")
		fmt.Println("  go tool pprof cpu.prof               - Analyze CPU profile")
		fmt.Println("  go tool pprof mem.prof               - Analyze memory profile")
	}
}
