package main

import (
	"fmt"
	"testing"
	"time"
)

// 🧪 Ejercicio 4: Cache-friendly vs Cache-unfriendly data structures

// 🐌 Array of Structs (AoS) - Cache unfriendly para operaciones selectivas
type Point3D struct {
	X, Y, Z float64
	Active  bool
	_       [7]byte // Padding para alignment
}

type PointsAoS []Point3D

// 🚀 Struct of Arrays (SoA) - Cache friendly
type PointsSoA struct {
	X, Y, Z []float64
	Active  []bool
}

// 🛠️ Crear datos de prueba para AoS
func CreateAoSData(size int) PointsAoS {
	points := make(PointsAoS, size)
	for i := range points {
		points[i] = Point3D{
			X:      float64(i),
			Y:      float64(i * 2),
			Z:      float64(i * 3),
			Active: i%2 == 0, // 50% activos
		}
	}
	return points
}

// 🛠️ Crear datos de prueba para SoA
func CreateSoAData(size int) PointsSoA {
	points := PointsSoA{
		X:      make([]float64, size),
		Y:      make([]float64, size),
		Z:      make([]float64, size),
		Active: make([]bool, size),
	}

	for i := 0; i < size; i++ {
		points.X[i] = float64(i)
		points.Y[i] = float64(i * 2)
		points.Z[i] = float64(i * 3)
		points.Active[i] = i%2 == 0
	}
	return points
}

// 🐌 Procesar solo puntos activos con AoS
func (points PointsAoS) ProcessActivePoints() float64 {
	var sum float64
	for i := range points {
		if points[i].Active {
			sum += points[i].X + points[i].Y + points[i].Z
		}
	}
	return sum
}

// 🚀 Procesar solo puntos activos con SoA
func (points PointsSoA) ProcessActivePoints() float64 {
	var sum float64
	for i := range points.Active {
		if points.Active[i] {
			sum += points.X[i] + points.Y[i] + points.Z[i]
		}
	}
	return sum
}

// 📊 Benchmarks para comparar cache performance
func BenchmarkAoSProcessing1K(b *testing.B) {
	points := CreateAoSData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

func BenchmarkSoAProcessing1K(b *testing.B) {
	points := CreateSoAData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

func BenchmarkAoSProcessing10K(b *testing.B) {
	points := CreateAoSData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

func BenchmarkSoAProcessing10K(b *testing.B) {
	points := CreateSoAData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

func BenchmarkAoSProcessing100K(b *testing.B) {
	points := CreateAoSData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

func BenchmarkSoAProcessing100K(b *testing.B) {
	points := CreateSoAData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = points.ProcessActivePoints()
	}
}

// 🎯 Demostración de la diferencia de cache performance
func DemoCachePerformance() {
	fmt.Println("⚡ Cache Locality Performance Demo")
	fmt.Println("=================================")

	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		fmt.Printf("\n📊 Tamaño: %d puntos\n", size)

		// Crear datos
		aosData := CreateAoSData(size)
		soaData := CreateSoAData(size)

		// Benchmark AoS
		iterations := 1000
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = aosData.ProcessActivePoints()
		}
		aosDuration := time.Since(start)

		// Benchmark SoA
		start = time.Now()
		for i := 0; i < iterations; i++ {
			_ = soaData.ProcessActivePoints()
		}
		soaDuration := time.Since(start)

		fmt.Printf("🐌 AoS (Array of Structs): %v\n", aosDuration)
		fmt.Printf("🚀 SoA (Struct of Arrays): %v\n", soaDuration)

		if aosDuration > soaDuration {
			speedup := float64(aosDuration) / float64(soaDuration)
			fmt.Printf("⚡ SoA es %.2fx más rápido\n", speedup)
		}
	}
}
