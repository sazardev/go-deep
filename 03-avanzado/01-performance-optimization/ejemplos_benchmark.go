package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// 🧪 Ejercicio 1: Comparación de algoritmos de búsqueda

// 🐌 Búsqueda lineal simple
func linearSearch(data []int, target int) int {
	for i, val := range data {
		if val == target {
			return i
		}
	}
	return -1
}

// 🚀 Búsqueda binaria (requiere slice ordenado)
func binarySearch(data []int, target int) int {
	left, right := 0, len(data)-1

	for left <= right {
		mid := left + (right-left)/2
		if data[mid] == target {
			return mid
		}
		if data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 🛠️ Función para generar datos de prueba
func generateTestData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i * 2 // Números pares
	}
	return data
}

// 🛠️ Función para generar datos ordenados
func generateSortedTestData(size int) []int {
	return generateTestData(size) // Ya están ordenados
}

// 📊 Benchmarks para comparación
func BenchmarkLinearSearch1K(b *testing.B) {
	data := generateTestData(1000)
	target := data[500]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = linearSearch(data, target)
	}
}

func BenchmarkBinarySearch1K(b *testing.B) {
	data := generateSortedTestData(1000)
	target := data[500]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binarySearch(data, target)
	}
}

func BenchmarkLinearSearch10K(b *testing.B) {
	data := generateTestData(10000)
	target := data[5000]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = linearSearch(data, target)
	}
}

func BenchmarkBinarySearch10K(b *testing.B) {
	data := generateSortedTestData(10000)
	target := data[5000]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binarySearch(data, target)
	}
}

func BenchmarkLinearSearch100K(b *testing.B) {
	data := generateTestData(100000)
	target := data[50000]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = linearSearch(data, target)
	}
}

func BenchmarkBinarySearch100K(b *testing.B) {
	data := generateSortedTestData(100000)
	target := data[50000]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = binarySearch(data, target)
	}
}

// 🎯 Función principal para demos
func main() {
	fmt.Println("🔍 Comparación de Algoritmos de Búsqueda")
	fmt.Println("=======================================")

	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		data := generateSortedTestData(size)
		target := data[size/2]

		// Medir búsqueda lineal
		start := time.Now()
		result1 := linearSearch(data, target)
		linearTime := time.Since(start)

		// Medir búsqueda binaria
		start = time.Now()
		result2 := binarySearch(data, target)
		binaryTime := time.Since(start)

		fmt.Printf("\n📊 Tamaño: %d elementos\n", size)
		fmt.Printf("🐌 Búsqueda Lineal: %v (posición: %d)\n", linearTime, result1)
		fmt.Printf("🚀 Búsqueda Binaria: %v (posición: %d)\n", binaryTime, result2)

		if linearTime > 0 && binaryTime > 0 {
			speedup := float64(linearTime) / float64(binaryTime)
			fmt.Printf("⚡ Mejora: %.2fx más rápida\n", speedup)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Ejecutar demo de strings
	DemoStringConcatenation()

	fmt.Println("\n🧪 Para ejecutar benchmarks:")
	fmt.Println("go test -bench=. -benchmem")
}
