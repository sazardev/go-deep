// 🔟 EJERCICIO 10: Algoritmos de Cadenas con Slices
// =================================================
//
// Demostración completa del Ejercicio 10: algoritmos KMP y LCS

package main

import (
	"fmt"
	"strings"
	"time"
)

// ===== ALGORITMO KMP (Knuth-Morris-Pratt) =====

// buscarPatronKMP encuentra todas las ocurrencias de un patrón en un texto
// usando el algoritmo KMP para eficiencia O(n+m)
func buscarPatronKMP(texto, patron string) []int {
	if len(patron) == 0 {
		return []int{}
	}

	lps := construirTablaLPS(patron)
	var indices []int

	i, j := 0, 0 // i para texto, j para patrón

	for i < len(texto) {
		if patron[j] == texto[i] {
			i++
			j++
		}

		if j == len(patron) {
			indices = append(indices, i-j)
			j = lps[j-1]
		} else if i < len(texto) && patron[j] != texto[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return indices
}

// construirTablaLPS construye la tabla de Longest Proper Prefix que es también Suffix
func construirTablaLPS(patron string) []int {
	lps := make([]int, len(patron))
	longitud := 0
	i := 1

	for i < len(patron) {
		if patron[i] == patron[longitud] {
			longitud++
			lps[i] = longitud
			i++
		} else {
			if longitud != 0 {
				longitud = lps[longitud-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// ===== ALGORITMO LCS (Longest Common Subsequence) =====

// subsecuenciaComunMasLarga encuentra la subsecuencia común más larga
// entre dos cadenas usando programación dinámica
func subsecuenciaComunMasLarga(s1, s2 string) string {
	m, n := len(s1), len(s2)

	// Crear tabla DP
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Llenar tabla DP
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruir LCS
	var lcs []byte
	i, j := m, n

	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = append([]byte{s1[i-1]}, lcs...)
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}

// Función auxiliar para encontrar el máximo
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ===== ALGORITMOS ADICIONALES CON SLICES =====

// encontrarTodasLasSubcadenas encuentra todas las subcadenas de una cadena
func encontrarTodasLasSubcadenas(s string) []string {
	var subcadenas []string
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			subcadenas = append(subcadenas, s[i:j])
		}
	}

	return subcadenas
}

// palindromoMasLargo encuentra el palíndromo más largo en una cadena
func palindromoMasLargo(s string) string {
	if len(s) == 0 {
		return ""
	}

	inicio, longitud := 0, 1

	for i := 0; i < len(s); i++ {
		// Palíndromos de longitud impar
		izq, der := i, i
		for izq >= 0 && der < len(s) && s[izq] == s[der] {
			if der-izq+1 > longitud {
				inicio = izq
				longitud = der - izq + 1
			}
			izq--
			der++
		}

		// Palíndromos de longitud par
		izq, der = i, i+1
		for izq >= 0 && der < len(s) && s[izq] == s[der] {
			if der-izq+1 > longitud {
				inicio = izq
				longitud = der - izq + 1
			}
			izq--
			der++
		}
	}

	return s[inicio : inicio+longitud]
}

// distanciaLevenshtein calcula la distancia de edición entre dos cadenas
func distanciaLevenshtein(s1, s2 string) int {
	m, n := len(s1), len(s2)

	// Crear matriz DP
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Inicializar base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Llenar matriz DP
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

// anagramasSonIguales verifica si dos cadenas son anagramas
func anagramasSonIguales(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq := make(map[rune]int)

	// Contar frecuencias de s1
	for _, char := range s1 {
		freq[char]++
	}

	// Decrementar con frecuencias de s2
	for _, char := range s2 {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}

	// Verificar que todas las frecuencias sean 0
	for _, count := range freq {
		if count != 0 {
			return false
		}
	}

	return true
}

// ===== FUNCIONES DE DEMOSTRACIÓN =====

func main() {
	fmt.Println("🔟 EJERCICIO 10: ALGORITMOS DE CADENAS CON SLICES")
	fmt.Println("=================================================")

	demoKMP()
	demoLCS()
	demoAlgoritmosAdicionales()
	demoBenchmarks()
}

func demoKMP() {
	fmt.Println("\n🔍 ALGORITMO KMP (Knuth-Morris-Pratt)")
	fmt.Println("=====================================")

	// Ejemplo básico
	texto := "ABABDABACDABABCABCABCABCABC"
	patron := "ABCABC"

	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Patrón: %s\n", patron)

	// Construir y mostrar tabla LPS
	lps := construirTablaLPS(patron)
	fmt.Printf("Tabla LPS: %v\n", lps)
	fmt.Printf("Explicación LPS:\n")
	for i, val := range lps {
		fmt.Printf("  LPS[%d] = %d (prefijo '%s')\n", i, val, patron[:i+1])
	}

	// Buscar ocurrencias
	ocurrencias := buscarPatronKMP(texto, patron)
	fmt.Printf("\nOcurrencias encontradas: %v\n", ocurrencias)

	// Mostrar ocurrencias visualmente
	for i, pos := range ocurrencias {
		antes := texto[:pos]
		match := texto[pos : pos+len(patron)]
		despues := texto[pos+len(patron):]
		fmt.Printf("Ocurrencia %d (pos %d): %s[%s]%s\n", i+1, pos, antes, match, despues)
	}

	// Ejemplos adicionales
	fmt.Println("\n📚 Ejemplos Adicionales:")

	ejemplos := []struct {
		texto, patron string
	}{
		{"abababcababa", "abab"},
		{"AABAACAADAABAABA", "AABA"},
		{"hello world hello", "hello"},
		{"mississippi", "issi"},
	}

	for _, ej := range ejemplos {
		ocur := buscarPatronKMP(ej.texto, ej.patron)
		fmt.Printf("'%s' en '%s': %v\n", ej.patron, ej.texto, ocur)
	}
}

func demoLCS() {
	fmt.Println("\n🧬 ALGORITMO LCS (Longest Common Subsequence)")
	fmt.Println("=============================================")

	// Ejemplo básico
	s1 := "ABCDGH"
	s2 := "AEDFHR"

	fmt.Printf("Cadena 1: %s\n", s1)
	fmt.Printf("Cadena 2: %s\n", s2)

	lcs := subsecuenciaComunMasLarga(s1, s2)
	fmt.Printf("LCS: '%s' (longitud: %d)\n", lcs, len(lcs))

	// Mostrar cómo se forma la LCS
	fmt.Printf("Formación de LCS:\n")
	fmt.Printf("  %s\n", s1)
	fmt.Printf("  %s\n", s2)
	fmt.Printf("  LCS: %s\n", lcs)

	// Ejemplos prácticos
	fmt.Println("\n📚 Ejemplos Prácticos:")

	ejemplos := []struct {
		s1, s2, descripcion string
	}{
		{"AGGTAB", "GXTXAYB", "Secuencias de ADN"},
		{"PROGRAMMING", "ALGORITHM", "Palabras técnicas"},
		{"ABCDEFG", "ACBDEGF", "Secuencias similares"},
		{"HUMAN", "CHIMPANZEE", "Especies relacionadas"},
		{"1234567", "1357924", "Secuencias numéricas"},
	}

	for _, ej := range ejemplos {
		lcs := subsecuenciaComunMasLarga(ej.s1, ej.s2)
		similitud := float64(len(lcs)) / float64(max(len(ej.s1), len(ej.s2))) * 100
		fmt.Printf("%s:\n", ej.descripcion)
		fmt.Printf("  '%s' ∩ '%s' = '%s' (%.1f%% similitud)\n",
			ej.s1, ej.s2, lcs, similitud)
	}
}

func demoAlgoritmosAdicionales() {
	fmt.Println("\n🔧 ALGORITMOS ADICIONALES CON SLICES")
	fmt.Println("====================================")

	// Demo 1: Todas las subcadenas
	fmt.Println("📝 1. Todas las subcadenas:")
	texto := "abc"
	subcadenas := encontrarTodasLasSubcadenas(texto)
	fmt.Printf("Subcadenas de '%s': %v\n", texto, subcadenas)
	fmt.Printf("Total: %d subcadenas\n", len(subcadenas))

	// Demo 2: Palíndromo más largo
	fmt.Println("\n🔄 2. Palíndromo más largo:")
	textos := []string{"babad", "cbbd", "racecar", "abcdef", "aabbaa"}

	for _, t := range textos {
		palindromo := palindromoMasLargo(t)
		fmt.Printf("'%s' → '%s'\n", t, palindromo)
	}

	// Demo 3: Distancia de Levenshtein
	fmt.Println("\n📏 3. Distancia de Levenshtein:")
	pares := []struct {
		s1, s2 string
	}{
		{"kitten", "sitting"},
		{"saturday", "sunday"},
		{"programming", "algorithm"},
		{"hello", "hola"},
	}

	for _, par := range pares {
		distancia := distanciaLevenshtein(par.s1, par.s2)
		fmt.Printf("'%s' ↔ '%s': distancia = %d\n", par.s1, par.s2, distancia)
	}

	// Demo 4: Anagramas
	fmt.Println("\n🔀 4. Detección de anagramas:")
	paresAnagramas := []struct {
		s1, s2 string
	}{
		{"listen", "silent"},
		{"evil", "live"},
		{"a gentleman", "elegant man"},
		{"hello", "world"},
	}

	for _, par := range paresAnagramas {
		esAnagrama := anagramasSonIguales(par.s1, par.s2)
		resultado := "❌"
		if esAnagrama {
			resultado = "✅"
		}
		fmt.Printf("%s '%s' ↔ '%s'\n", resultado, par.s1, par.s2)
	}
}

func demoBenchmarks() {
	fmt.Println("\n⏱️ BENCHMARKS DE RENDIMIENTO")
	fmt.Println("============================")

	// Benchmark KMP vs búsqueda nativa
	fmt.Println("🔍 Comparación KMP vs strings.Index:")

	textoGrande := strings.Repeat("ABCAB", 10000) + "TARGETPATTERN" + strings.Repeat("XYZXYZ", 10000)
	patron := "TARGETPATTERN"

	fmt.Printf("Tamaño del texto: %d caracteres\n", len(textoGrande))
	fmt.Printf("Patrón: '%s'\n", patron)

	// Benchmark KMP
	inicio := time.Now()
	resultadoKMP := buscarPatronKMP(textoGrande, patron)
	tiempoKMP := time.Since(inicio)

	// Benchmark strings.Index
	inicio = time.Now()
	resultadoNativo := strings.Index(textoGrande, patron)
	tiempoNativo := time.Since(inicio)

	fmt.Printf("KMP: encontrado en %v (tiempo: %v)\n", resultadoKMP, tiempoKMP)
	fmt.Printf("strings.Index: encontrado en %d (tiempo: %v)\n", resultadoNativo, tiempoNativo)

	// Benchmark LCS
	fmt.Println("\n🧬 Benchmark LCS:")
	s1 := strings.Repeat("ABCDEFGH", 100)
	s2 := strings.Repeat("ACDFHIJK", 100)

	fmt.Printf("Comparando cadenas de %d y %d caracteres\n", len(s1), len(s2))

	inicio = time.Now()
	lcs := subsecuenciaComunMasLarga(s1, s2)
	tiempoLCS := time.Since(inicio)

	fmt.Printf("LCS longitud: %d (tiempo: %v)\n", len(lcs), tiempoLCS)

	// Análisis de complejidad
	fmt.Println("\n📊 Análisis de Complejidad:")
	fmt.Printf("KMP: O(n + m) donde n=%d, m=%d\n", len(textoGrande), len(patron))
	fmt.Printf("LCS: O(n * m) donde n=%d, m=%d\n", len(s1), len(s2))

	fmt.Println("\n🎉 ¡Demostración del Ejercicio 10 completada!")
	fmt.Println("\n💡 Conceptos demostrados:")
	fmt.Println("   ✅ Algoritmo KMP para búsqueda eficiente de patrones")
	fmt.Println("   ✅ Tabla LPS (Longest Proper Prefix Suffix)")
	fmt.Println("   ✅ Algoritmo LCS con programación dinámica")
	fmt.Println("   ✅ Algoritmos adicionales: palíndromos, Levenshtein, anagramas")
	fmt.Println("   ✅ Análisis de rendimiento y complejidad")
	fmt.Println("   ✅ Aplicaciones prácticas en procesamiento de texto")
}
