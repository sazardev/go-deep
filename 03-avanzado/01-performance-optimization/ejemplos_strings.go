package main

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

// 🧪 Ejercicio 2: Optimización de concatenación de strings

// 🐌 Método 1: Concatenación naive con +
func ConcatenateNaive(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

// 🚀 Método 2: Usando strings.Builder con pre-allocación
func ConcatenateBuilder(strs []string) string {
	var builder strings.Builder

	// Pre-calcular tamaño total
	totalLen := 0
	for _, s := range strs {
		totalLen += len(s)
	}
	builder.Grow(totalLen)

	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// 🚀 Método 3: Usando strings.Join
func ConcatenateJoin(strs []string) string {
	return strings.Join(strs, "")
}

// 🛠️ Generar datos de prueba
func generateStringData(count, size int) []string {
	strs := make([]string, count)
	base := strings.Repeat("a", size)
	for i := range strs {
		strs[i] = base
	}
	return strs
}

// 📊 Benchmarks para comparación
func BenchmarkConcatenateNaive100(b *testing.B) {
	strs := generateStringData(100, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateNaive(strs)
	}
}

func BenchmarkConcatenateBuilder100(b *testing.B) {
	strs := generateStringData(100, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateBuilder(strs)
	}
}

func BenchmarkConcatenateJoin100(b *testing.B) {
	strs := generateStringData(100, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateJoin(strs)
	}
}

func BenchmarkConcatenateNaive1000(b *testing.B) {
	strs := generateStringData(1000, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateNaive(strs)
	}
}

func BenchmarkConcatenateBuilder1000(b *testing.B) {
	strs := generateStringData(1000, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateBuilder(strs)
	}
}

func BenchmarkConcatenateJoin1000(b *testing.B) {
	strs := generateStringData(1000, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatenateJoin(strs)
	}
}

// 🧠 Demostración de uso de memoria
func measureMemory(fn func()) (allocsBefore, allocsAfter uint64, duration time.Duration) {
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	start := time.Now()
	fn()
	duration = time.Since(start)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	return m1.TotalAlloc, m2.TotalAlloc, duration
}

// 🎯 Función de demostración
func DemoStringConcatenation() {
	fmt.Println("🔗 Optimización de Concatenación de Strings")
	fmt.Println("==========================================")

	testCases := []struct {
		name  string
		count int
		size  int
	}{
		{"Pequeño", 100, 10},
		{"Medio", 1000, 10},
		{"Grande", 1000, 100},
	}

	for _, tc := range testCases {
		fmt.Printf("\n📊 Caso: %s (%d strings de %d chars)\n", tc.name, tc.count, tc.size)
		strs := generateStringData(tc.count, tc.size)

		// Test método naive
		allocsBefore, allocsAfter, duration := measureMemory(func() {
			_ = ConcatenateNaive(strs)
		})
		fmt.Printf("🐌 Naive: %v, Memoria: %d bytes\n",
			duration, allocsAfter-allocsBefore)

		// Test strings.Builder
		allocsBefore, allocsAfter, duration = measureMemory(func() {
			_ = ConcatenateBuilder(strs)
		})
		fmt.Printf("🚀 Builder: %v, Memoria: %d bytes\n",
			duration, allocsAfter-allocsBefore)

		// Test strings.Join
		allocsBefore, allocsAfter, duration = measureMemory(func() {
			_ = ConcatenateJoin(strs)
		})
		fmt.Printf("🚀 Join: %v, Memoria: %d bytes\n",
			duration, allocsAfter-allocsBefore)
	}

	fmt.Println("\n🧪 Para ejecutar benchmarks:")
	fmt.Println("go test -bench=BenchmarkConcatenate -benchmem")
}
