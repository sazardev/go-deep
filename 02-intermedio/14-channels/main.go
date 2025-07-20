// ==============================================
// LECCIÓN 14: Channels - Ejecutor Principal
// ==============================================

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("📡 LECCIÓN 14: Channels")
	fmt.Println("======================")
	fmt.Printf("Go %s | CPUs: %d\n\n", runtime.Version(), runtime.NumCPU())

	fmt.Println("📝 Archivos disponibles:")
	fmt.Println("   📋 ejercicios.go  - Plantillas para practicar")
	fmt.Println("   ✅ soluciones.go  - Implementaciones completas")
	fmt.Println("")
	fmt.Println("🚀 Para ejecutar:")
	fmt.Println("   go run ejercicios.go   # Plantillas de ejercicios")
	fmt.Println("   go run soluciones.go   # Soluciones completas")
	fmt.Println("")
	fmt.Println("💡 Ejecutando soluciones por defecto...")

	// Ejecutar las soluciones por defecto
	ejecutarSoluciones()
}
