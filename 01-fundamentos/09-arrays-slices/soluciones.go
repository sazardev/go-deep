// 📚 Soluciones: Arrays y Slices
// ===============================
//
// Soluciones completas y explicadas para todos los ejercicios de Arrays y Slices.

package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

// ===== SOLUCIÓN EJERCICIO 1: MANIPULACIÓN BÁSICA =====

// 1.1 - Análisis completo de un array
func analizarArray(numeros [10]int) (suma int, promedio float64, maximo int, minimo int) {
	// Inicializar con el primer elemento
	suma = numeros[0]
	maximo = numeros[0]
	minimo = numeros[0]

	// Procesar el resto de elementos
	for i := 1; i < len(numeros); i++ {
		suma += numeros[i]
		if numeros[i] > maximo {
			maximo = numeros[i]
		}
		if numeros[i] < minimo {
			minimo = numeros[i]
		}
	}

	promedio = float64(suma) / float64(len(numeros))
	return
}

// 1.2 - Intercambiar elementos con validación
func intercambiarElementos(slice []int, i, j int) error {
	if i < 0 || i >= len(slice) || j < 0 || j >= len(slice) {
		return errors.New("índices fuera de rango")
	}

	slice[i], slice[j] = slice[j], slice[i]
	return nil
}

// 1.3 - Invertir slice in-place
func invertirSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	}
}

// ===== SOLUCIÓN EJERCICIO 2: ALGORITMOS DE BÚSQUEDA =====

// 2.1 - Búsqueda lineal que devuelve todos los índices
func busquedaLinealTodos(slice []int, valor int) []int {
	var indices []int
	for i, v := range slice {
		if v == valor {
			indices = append(indices, i)
		}
	}
	return indices
}

// 2.2 - Búsqueda binaria recursiva
func busquedaBinariaRecursiva(slice []int, valor, inicio, fin int) int {
	if inicio > fin {
		return -1
	}

	medio := inicio + (fin-inicio)/2

	if slice[medio] == valor {
		return medio
	} else if slice[medio] > valor {
		return busquedaBinariaRecursiva(slice, valor, inicio, medio-1)
	} else {
		return busquedaBinariaRecursiva(slice, valor, medio+1, fin)
	}
}

// 2.3 - Encontrar primer elemento que cumple condición
func encontrarPrimero(slice []int, condicion func(int) bool) (int, bool) {
	for _, elemento := range slice {
		if condicion(elemento) {
			return elemento, true
		}
	}
	return 0, false
}

// ===== SOLUCIÓN EJERCICIO 3: ALGORITMOS DE ORDENAMIENTO =====

// 3.1 - Selection Sort
func selectionSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		// Encontrar el índice del elemento mínimo
		indiceMin := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[indiceMin] {
				indiceMin = j
			}
		}
		// Intercambiar con el elemento en posición i
		slice[i], slice[indiceMin] = slice[indiceMin], slice[i]
	}
}

// 3.2 - Insertion Sort
func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		clave := slice[i]
		j := i - 1

		// Mover elementos mayores que clave una posición adelante
		for j >= 0 && slice[j] > clave {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = clave
	}
}

// 3.3 - Merge Sort
func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	medio := len(slice) / 2
	izquierda := mergeSort(slice[:medio])
	derecha := mergeSort(slice[medio:])

	return merge(izquierda, derecha)
}

// Función auxiliar para merge sort
func merge(izquierda, derecha []int) []int {
	resultado := make([]int, 0, len(izquierda)+len(derecha))
	i, j := 0, 0

	// Combinar elementos en orden
	for i < len(izquierda) && j < len(derecha) {
		if izquierda[i] <= derecha[j] {
			resultado = append(resultado, izquierda[i])
			i++
		} else {
			resultado = append(resultado, derecha[j])
			j++
		}
	}

	// Agregar elementos restantes
	resultado = append(resultado, izquierda[i:]...)
	resultado = append(resultado, derecha[j:]...)

	return resultado
}

// ===== SOLUCIÓN EJERCICIO 4: OPERACIONES FUNCIONALES =====

// 4.1 - Función Filter
func filtrar(slice []int, predicado func(int) bool) []int {
	var resultado []int
	for _, elemento := range slice {
		if predicado(elemento) {
			resultado = append(resultado, elemento)
		}
	}
	return resultado
}

// 4.2 - Función Map
func mapear(slice []int, transformacion func(int) int) []int {
	resultado := make([]int, len(slice))
	for i, elemento := range slice {
		resultado[i] = transformacion(elemento)
	}
	return resultado
}

// 4.3 - Función Reduce
func reducir(slice []int, inicial int, operacion func(int, int) int) int {
	resultado := inicial
	for _, elemento := range slice {
		resultado = operacion(resultado, elemento)
	}
	return resultado
}

// 4.4 - Combinar filter y map
func filtrarYMapear(slice []int, predicado func(int) bool, transformacion func(int) int) []int {
	var resultado []int
	for _, elemento := range slice {
		if predicado(elemento) {
			resultado = append(resultado, transformacion(elemento))
		}
	}
	return resultado
}

// ===== SOLUCIÓN EJERCICIO 5: MATRICES MULTIDIMENSIONALES =====

// 5.1 - Matriz identidad
func matrizIdentidad(n int) [][]int {
	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, n)
		matriz[i][i] = 1 // Diagonal principal = 1
	}
	return matriz
}

// 5.2 - Multiplicación de matrices
func multiplicarMatrices(a, b [][]int) ([][]int, error) {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil, errors.New("dimensiones incompatibles para multiplicación")
	}

	filasA, colsA := len(a), len(a[0])
	colsB := len(b[0])

	// Crear matriz resultado
	resultado := make([][]int, filasA)
	for i := range resultado {
		resultado[i] = make([]int, colsB)
	}

	// Multiplicación matriz
	for i := 0; i < filasA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				resultado[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return resultado, nil
}

// 5.3 - Transposición de matriz
func transponerMatriz(matriz [][]int) [][]int {
	if len(matriz) == 0 {
		return [][]int{}
	}

	filas, cols := len(matriz), len(matriz[0])
	transpuesta := make([][]int, cols)

	for i := range transpuesta {
		transpuesta[i] = make([]int, filas)
	}

	for i := 0; i < filas; i++ {
		for j := 0; j < cols; j++ {
			transpuesta[j][i] = matriz[i][j]
		}
	}

	return transpuesta
}

// 5.4 - Encontrar máximo en matriz
func encontrarMaximoMatriz(matriz [][]int) (valor, fila, columna int) {
	if len(matriz) == 0 || len(matriz[0]) == 0 {
		return 0, -1, -1
	}

	valor = matriz[0][0]
	fila, columna = 0, 0

	for i := range matriz {
		for j, v := range matriz[i] {
			if v > valor {
				valor = v
				fila = i
				columna = j
			}
		}
	}

	return
}

// ===== SOLUCIÓN EJERCICIO 6: ESTRUCTURAS DE DATOS AVANZADAS =====

// 6.1 - Cola Circular completa
type ColaCircular struct {
	datos     []interface{}
	frente    int
	final     int
	tamaño    int
	capacidad int
}

func NuevaColaCircular(capacidad int) *ColaCircular {
	return &ColaCircular{
		datos:     make([]interface{}, capacidad),
		frente:    0,
		final:     0,
		tamaño:    0,
		capacidad: capacidad,
	}
}

func (c *ColaCircular) Enqueue(elemento interface{}) bool {
	if c.EstaLlena() {
		return false
	}

	c.datos[c.final] = elemento
	c.final = (c.final + 1) % c.capacidad
	c.tamaño++
	return true
}

func (c *ColaCircular) Dequeue() (interface{}, bool) {
	if c.EstaVacia() {
		return nil, false
	}

	elemento := c.datos[c.frente]
	c.datos[c.frente] = nil // Limpiar referencia
	c.frente = (c.frente + 1) % c.capacidad
	c.tamaño--
	return elemento, true
}

func (c *ColaCircular) EstaLlena() bool {
	return c.tamaño == c.capacidad
}

func (c *ColaCircular) EstaVacia() bool {
	return c.tamaño == 0
}

// 6.2 - Buffer Deslizante para promedios móviles
type BufferDeslizante struct {
	datos  []float64
	indice int
	lleno  bool
	tamaño int
}

func NuevoBufferDeslizante(tamaño int) *BufferDeslizante {
	return &BufferDeslizante{
		datos:  make([]float64, tamaño),
		indice: 0,
		lleno:  false,
		tamaño: tamaño,
	}
}

func (b *BufferDeslizante) Agregar(valor float64) {
	b.datos[b.indice] = valor
	b.indice = (b.indice + 1) % b.tamaño

	if !b.lleno && b.indice == 0 {
		b.lleno = true
	}
}

func (b *BufferDeslizante) Promedio() float64 {
	elementos := b.tamaño
	if !b.lleno {
		elementos = b.indice
	}

	if elementos == 0 {
		return 0
	}

	suma := 0.0
	for i := 0; i < elementos; i++ {
		suma += b.datos[i]
	}

	return suma / float64(elementos)
}

func (b *BufferDeslizante) Maximo() float64 {
	elementos := b.tamaño
	if !b.lleno {
		elementos = b.indice
	}

	if elementos == 0 {
		return 0
	}

	max := b.datos[0]
	for i := 1; i < elementos; i++ {
		if b.datos[i] > max {
			max = b.datos[i]
		}
	}

	return max
}

func (b *BufferDeslizante) Minimo() float64 {
	elementos := b.tamaño
	if !b.lleno {
		elementos = b.indice
	}

	if elementos == 0 {
		return 0
	}

	min := b.datos[0]
	for i := 1; i < elementos; i++ {
		if b.datos[i] < min {
			min = b.datos[i]
		}
	}

	return min
}

// ===== SOLUCIÓN EJERCICIO 7: ALGORITMOS AVANZADOS =====

// 7.1 - Partición de Lomuto para Quicksort
func particionLomuto(slice []int, bajo, alto int) int {
	pivot := slice[alto] // Último elemento como pivote
	i := bajo - 1        // Índice del elemento más pequeño

	for j := bajo; j < alto; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[alto] = slice[alto], slice[i+1]
	return i + 1
}

// 7.2 - k-ésimo elemento más pequeño (Quickselect)
func kEsimoMenor(slice []int, k int) int {
	if k <= 0 || k > len(slice) {
		return -1
	}

	return quickselect(slice, 0, len(slice)-1, k-1)
}

func quickselect(slice []int, bajo, alto, k int) int {
	if bajo == alto {
		return slice[bajo]
	}

	pivotIndex := particionLomuto(slice, bajo, alto)

	if k == pivotIndex {
		return slice[k]
	} else if k < pivotIndex {
		return quickselect(slice, bajo, pivotIndex-1, k)
	} else {
		return quickselect(slice, pivotIndex+1, alto, k)
	}
}

// 7.3 - Par con suma más cercana al objetivo
func parSumaCercana(slice []int, objetivo int) (int, int, int) {
	if len(slice) < 2 {
		return 0, 0, math.MaxInt32
	}

	menorDiferencia := math.MaxInt32
	mejorI, mejorJ := 0, 1

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			suma := slice[i] + slice[j]
			diferencia := int(math.Abs(float64(objetivo - suma)))

			if diferencia < menorDiferencia {
				menorDiferencia = diferencia
				mejorI, mejorJ = i, j
			}
		}
	}

	return slice[mejorI], slice[mejorJ], menorDiferencia
}

// 7.4 - Máximo en ventana deslizante
func maximoVentanaDeslizante(slice []int, tamaño int) []int {
	if len(slice) < tamaño || tamaño <= 0 {
		return []int{}
	}

	var resultado []int

	for i := 0; i <= len(slice)-tamaño; i++ {
		max := slice[i]
		for j := i + 1; j < i+tamaño; j++ {
			if slice[j] > max {
				max = slice[j]
			}
		}
		resultado = append(resultado, max)
	}

	return resultado
}

// ===== SOLUCIÓN EJERCICIO 8: OPTIMIZACIÓN Y RENDIMIENTO =====

// 8.1 - Eliminar duplicados manteniendo orden
func eliminarDuplicados(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}

	seen := make(map[int]bool)
	var resultado []int

	for _, elemento := range slice {
		if !seen[elemento] {
			seen[elemento] = true
			resultado = append(resultado, elemento)
		}
	}

	return resultado
}

// 8.2 - Intersección de slices ordenados
func interseccionOrdenados(slice1, slice2 []int) []int {
	var resultado []int
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] == slice2[j] {
			// Evitar duplicados en resultado
			if len(resultado) == 0 || resultado[len(resultado)-1] != slice1[i] {
				resultado = append(resultado, slice1[i])
			}
			i++
			j++
		} else if slice1[i] < slice2[j] {
			i++
		} else {
			j++
		}
	}

	return resultado
}

// 8.3 - Rotar slice k posiciones a la derecha
func rotarDerecha(slice []int, k int) {
	n := len(slice)
	if n == 0 {
		return
	}

	k = k % n // Manejar k > n
	if k == 0 {
		return
	}

	// Usar algoritmo de reversión
	reverse(slice, 0, n-1) // Reversar todo
	reverse(slice, 0, k-1) // Reversar primeros k
	reverse(slice, k, n-1) // Reversar resto
}

func reverse(slice []int, inicio, fin int) {
	for inicio < fin {
		slice[inicio], slice[fin] = slice[fin], slice[inicio]
		inicio++
		fin--
	}
}

// 8.4 - Compresión Run-Length
func compresionRunLength(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}

	var resultado []int
	actual := slice[0]
	cuenta := 1

	for i := 1; i < len(slice); i++ {
		if slice[i] == actual {
			cuenta++
		} else {
			resultado = append(resultado, actual, cuenta)
			actual = slice[i]
			cuenta = 1
		}
	}

	// Agregar último grupo
	resultado = append(resultado, actual, cuenta)
	return resultado
}

// ===== SOLUCIÓN EJERCICIO 9: CASOS DE USO PRÁCTICOS =====

// 9.1 - Histograma completo
type Histograma struct {
	bins     []int
	min      float64
	max      float64
	anchoBin float64
}

func NuevoHistograma(min, max float64, numBins int) *Histograma {
	return &Histograma{
		bins:     make([]int, numBins),
		min:      min,
		max:      max,
		anchoBin: (max - min) / float64(numBins),
	}
}

func (h *Histograma) AgregarValor(valor float64) {
	if valor < h.min || valor >= h.max {
		return // Valor fuera de rango
	}

	binIndex := int((valor - h.min) / h.anchoBin)
	if binIndex >= len(h.bins) {
		binIndex = len(h.bins) - 1
	}

	h.bins[binIndex]++
}

func (h *Histograma) ObtenerFrecuencias() []int {
	copia := make([]int, len(h.bins))
	copy(copia, h.bins)
	return copia
}

func (h *Histograma) BinMasFrecuente() int {
	maxFrecuencia := h.bins[0]
	indiceMax := 0

	for i, frecuencia := range h.bins {
		if frecuencia > maxFrecuencia {
			maxFrecuencia = frecuencia
			indiceMax = i
		}
	}

	return indiceMax
}

// 9.2 - Analizador de tendencias
type AnalizadorTendencias struct {
	datos []float64
}

func NuevoAnalizadorTendencias() *AnalizadorTendencias {
	return &AnalizadorTendencias{
		datos: make([]float64, 0),
	}
}

func (a *AnalizadorTendencias) AgregarDato(valor float64) {
	a.datos = append(a.datos, valor)
}

func (a *AnalizadorTendencias) PromedioMovil(ventana int) []float64 {
	if len(a.datos) < ventana {
		return []float64{}
	}

	var promedios []float64
	for i := 0; i <= len(a.datos)-ventana; i++ {
		suma := 0.0
		for j := i; j < i+ventana; j++ {
			suma += a.datos[j]
		}
		promedios = append(promedios, suma/float64(ventana))
	}

	return promedios
}

func (a *AnalizadorTendencias) DetectarTendencia() string {
	if len(a.datos) < 2 {
		return "estable"
	}

	// Calcular pendiente usando regresión lineal simple
	n := float64(len(a.datos))
	sumaX, sumaY, sumaXY, sumaX2 := 0.0, 0.0, 0.0, 0.0

	for i, y := range a.datos {
		x := float64(i)
		sumaX += x
		sumaY += y
		sumaXY += x * y
		sumaX2 += x * x
	}

	pendiente := (n*sumaXY - sumaX*sumaY) / (n*sumaX2 - sumaX*sumaX)

	if pendiente > 0.1 {
		return "ascendente"
	} else if pendiente < -0.1 {
		return "descendente"
	}
	return "estable"
}

func (a *AnalizadorTendencias) EncontrarPicos() []int {
	if len(a.datos) < 3 {
		return []int{}
	}

	var picos []int
	for i := 1; i < len(a.datos)-1; i++ {
		if a.datos[i] > a.datos[i-1] && a.datos[i] > a.datos[i+1] {
			picos = append(picos, i)
		}
	}

	return picos
}

// ===== SOLUCIÓN EJERCICIO 10: ALGORITMOS DE CADENAS CON SLICES =====

// 10.1 - Algoritmo KMP para búsqueda de patrones
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

// Construir tabla LPS (Longest Proper Prefix que es también Suffix)
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

// 10.2 - Subsecuencia común más larga (LCS)
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
	var lcs strings.Builder
	i, j := m, n

	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs.WriteByte(s1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// Reversar resultado
	resultado := lcs.String()
	return reverseString(resultado)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ===== FUNCIONES DE DEMOSTRACIÓN =====

func main() {
	fmt.Println("✅ SOLUCIONES: Arrays y Slices")
	fmt.Println("===============================")

	demostrarSoluciones()
}

func demostrarSoluciones() {
	fmt.Println("\n🚀 Ejecutando demostraciones de soluciones...")

	// Demo 1: Análisis de Array
	fmt.Println("\n1️⃣ ANÁLISIS DE ARRAY:")
	numeros := [10]int{5, 2, 8, 1, 9, 3, 7, 4, 6, 10}
	suma, promedio, maximo, minimo := analizarArray(numeros)
	fmt.Printf("Array: %v\n", numeros)
	fmt.Printf("📊 Suma: %d | Promedio: %.2f | Max: %d | Min: %d\n", suma, promedio, maximo, minimo)

	// Demo 2: Búsqueda Avanzada
	fmt.Println("\n2️⃣ BÚSQUEDAS:")
	slice := []int{1, 3, 5, 3, 7, 3, 9}
	indices := busquedaLinealTodos(slice, 3)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("🔍 Índices donde aparece 3: %v\n", indices)

	ordenado := []int{1, 3, 5, 7, 9, 11, 13, 15}
	indice := busquedaBinariaRecursiva(ordenado, 7, 0, len(ordenado)-1)
	fmt.Printf("🎯 Búsqueda binaria de 7 en %v: índice %d\n", ordenado, indice)

	// Demo 3: Ordenamiento
	fmt.Println("\n3️⃣ ALGORITMOS DE ORDENAMIENTO:")
	datos := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", datos)

	// Selection Sort
	copia := make([]int, len(datos))
	copy(copia, datos)
	selectionSort(copia)
	fmt.Printf("🔢 Selection Sort: %v\n", copia)

	// Merge Sort
	resultado := mergeSort(datos)
	fmt.Printf("🔀 Merge Sort: %v\n", resultado)

	// Demo 4: Operaciones Funcionales
	fmt.Println("\n4️⃣ OPERACIONES FUNCIONALES:")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pares := filtrar(nums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("➡️ Pares: %v\n", pares)

	cuadrados := mapear(nums[:5], func(n int) int { return n * n })
	fmt.Printf("📐 Cuadrados: %v\n", cuadrados)

	suma2 := reducir(nums[:5], 0, func(a, b int) int { return a + b })
	fmt.Printf("➕ Suma 1-5: %d\n", suma2)

	// Demo 5: Matrices
	fmt.Println("\n5️⃣ OPERACIONES CON MATRICES:")
	identidad := matrizIdentidad(3)
	fmt.Printf("🆔 Matriz identidad 3x3:\n")
	imprimirMatriz(identidad)

	matriz1 := [][]int{{1, 2}, {3, 4}}
	matriz2 := [][]int{{5, 6}, {7, 8}}
	producto, _ := multiplicarMatrices(matriz1, matriz2)
	fmt.Printf("\n✖️ Multiplicación:\n")
	fmt.Printf("A = ")
	imprimirMatriz(matriz1)
	fmt.Printf("B = ")
	imprimirMatriz(matriz2)
	fmt.Printf("A×B = ")
	imprimirMatriz(producto)

	// Demo 6: Estructuras Avanzadas
	fmt.Println("\n6️⃣ ESTRUCTURAS DE DATOS:")

	// Cola Circular
	fmt.Printf("🔄 Cola Circular:\n")
	cola := NuevaColaCircular(3)
	for i := 1; i <= 4; i++ {
		if cola.Enqueue(i) {
			fmt.Printf("✅ Enqueued %d\n", i)
		} else {
			fmt.Printf("❌ Cola llena, no se pudo encolar %d\n", i)
		}
	}

	for !cola.EstaVacia() {
		if elemento, ok := cola.Dequeue(); ok {
			fmt.Printf("⬅️ Dequeued: %v\n", elemento)
		}
	}

	// Buffer Deslizante
	fmt.Printf("\n📊 Buffer Deslizante:\n")
	buffer := NuevoBufferDeslizante(3)
	valores := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	for _, v := range valores {
		buffer.Agregar(v)
		fmt.Printf("Agregado %.1f | Promedio: %.2f | Min: %.1f | Max: %.1f\n",
			v, buffer.Promedio(), buffer.Minimo(), buffer.Maximo())
	}

	// Demo 7: Algoritmos Avanzados
	fmt.Println("\n7️⃣ ALGORITMOS AVANZADOS:")

	// k-ésimo menor
	nums2 := []int{3, 6, 2, 8, 1, 4, 9, 7, 5}
	k := 3
	kMenor := kEsimoMenor(make([]int, len(nums2)), k) // Copia para no modificar original
	copy(nums2, []int{3, 6, 2, 8, 1, 4, 9, 7, 5})
	fmt.Printf("🎯 %d-ésimo menor en %v: %d\n", k, nums2, kMenor)

	// Par suma cercana
	objetivo := 10
	elem1, elem2, diff := parSumaCercana(nums2, objetivo)
	fmt.Printf("🎪 Par más cercano a suma %d: %d + %d (diff: %d)\n", objetivo, elem1, elem2, diff)

	// Demo 8: Optimización
	fmt.Println("\n8️⃣ OPTIMIZACIÓN:")

	// Eliminar duplicados
	conDups := []int{1, 2, 2, 3, 4, 4, 4, 5}
	sinDups := eliminarDuplicados(conDups)
	fmt.Printf("🔄 Original: %v\n", conDups)
	fmt.Printf("✨ Sin duplicados: %v\n", sinDups)

	// Rotación
	rot := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("🔄 Antes rotación: %v\n", rot)
	rotarDerecha(rot, 3)
	fmt.Printf("➡️ Después rotar 3: %v\n", rot)

	// Demo 9: Casos Prácticos
	fmt.Println("\n9️⃣ CASOS DE USO PRÁCTICOS:")

	// Histograma
	fmt.Printf("📊 Histograma:\n")
	hist := NuevoHistograma(0.0, 10.0, 5)
	datosHist := []float64{1.2, 3.4, 5.6, 7.8, 2.1, 4.3, 6.5, 8.7, 1.9, 3.8}

	for _, v := range datosHist {
		hist.AgregarValor(v)
	}

	freqs := hist.ObtenerFrecuencias()
	fmt.Printf("Datos: %v\n", datosHist)
	fmt.Printf("Frecuencias: %v\n", freqs)
	fmt.Printf("Bin más frecuente: %d\n", hist.BinMasFrecuente())

	// Demo 10: Algoritmos de Cadenas
	fmt.Println("\n🔟 ALGORITMOS DE CADENAS:")

	// KMP
	texto := "ABABDABACDABABCABCABCABCABC"
	patron := "ABCABC"
	ocurrencias := buscarPatronKMP(texto, patron)
	fmt.Printf("🔍 Patrón '%s' en '%s'\n", patron, texto)
	fmt.Printf("📍 Ocurrencias en posiciones: %v\n", ocurrencias)

	// LCS
	s1, s2 := "ABCDGH", "AEDFHR"
	lcs := subsecuenciaComunMasLarga(s1, s2)
	fmt.Printf("🧬 LCS de '%s' y '%s': '%s'\n", s1, s2, lcs)

	fmt.Println("\n🎉 ¡Todas las demostraciones completadas exitosamente!")
}

// Función auxiliar para imprimir matrices
func imprimirMatriz(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Printf("%v\n", fila)
	}
}

// ===== BENCHMARKS =====

func demoBenchmarks() {
	fmt.Println("\n⏱️ BENCHMARKS DE RENDIMIENTO")
	fmt.Println("============================")

	// Benchmark de ordenamiento
	tamaños := []int{1000, 5000, 10000}

	for _, n := range tamaños {
		fmt.Printf("\n📏 Tamaño: %d elementos\n", n)

		datos := make([]int, n)
		for i := range datos {
			datos[i] = n - i // Peor caso para algunos algoritmos
		}

		// Selection Sort
		copia := make([]int, len(datos))
		copy(copia, datos)
		inicio := fmt.Sprintf("%v", datos[:5])

		timeStart := fmt.Sprintf("Selection Sort iniciado con %s...", inicio)
		fmt.Printf("🔢 %s\n", timeStart)
		selectionSort(copia)
		fmt.Printf("✅ Selection Sort completado\n")

		// Merge Sort
		fmt.Printf("🔀 Merge Sort iniciado...\n")
		mergeSort(datos)
		fmt.Printf("✅ Merge Sort completado\n")

		// Go sort nativo
		copy(copia, datos)
		fmt.Printf("⚡ Go Sort nativo iniciado...\n")
		sort.Ints(copia)
		fmt.Printf("✅ Go Sort completado\n")
	}
}

// ===== TESTS DE VALIDACIÓN =====

func validarTodasLasSoluciones() {
	fmt.Println("\n🧪 VALIDACIÓN DE SOLUCIONES")
	fmt.Println("============================")

	tests := []struct {
		nombre string
		test   func() bool
	}{
		{"Análisis de Array", testAnalizarArraySol},
		{"Búsqueda Lineal Todos", testBusquedaLinealTodosSol},
		{"Búsqueda Binaria Recursiva", testBusquedaBinariaSol},
		{"Selection Sort", testSelectionSortSol},
		{"Merge Sort", testMergeSortSol},
		{"Operaciones Funcionales", testOperacionesFuncionalesSol},
		{"Matriz Identidad", testMatrizIdentidadSol},
		{"Multiplicación Matrices", testMultiplicacionMatricesSol},
		{"Cola Circular", testColaCircularSol},
		{"Buffer Deslizante", testBufferDeslizanteSol},
		{"Eliminar Duplicados", testEliminarDuplicadosSol},
		{"Rotación", testRotacionSol},
		{"Histograma", testHistogramaSol},
		{"KMP", testKMPSol},
	}

	pasados := 0
	for _, test := range tests {
		if test.test() {
			fmt.Printf("✅ %s: PASADO\n", test.nombre)
			pasados++
		} else {
			fmt.Printf("❌ %s: FALLIDO\n", test.nombre)
		}
	}

	fmt.Printf("\n📊 Resultado Final: %d/%d tests pasados (%.1f%%)\n",
		pasados, len(tests), float64(pasados)/float64(len(tests))*100)
}

// Tests individuales
func testAnalizarArraySol() bool {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	suma, promedio, maximo, minimo := analizarArray(array)
	return suma == 55 && promedio == 5.5 && maximo == 10 && minimo == 1
}

func testBusquedaLinealTodosSol() bool {
	slice := []int{1, 3, 5, 3, 7, 3, 9}
	indices := busquedaLinealTodos(slice, 3)
	return len(indices) == 3 && indices[0] == 1 && indices[1] == 3 && indices[2] == 5
}

func testBusquedaBinariaSol() bool {
	slice := []int{1, 3, 5, 7, 9, 11, 13}
	indice := busquedaBinariaRecursiva(slice, 7, 0, len(slice)-1)
	return indice == 3
}

func testSelectionSortSol() bool {
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(slice)
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

func testMergeSortSol() bool {
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	resultado := mergeSort(slice)
	for i := 1; i < len(resultado); i++ {
		if resultado[i] < resultado[i-1] {
			return false
		}
	}
	return len(resultado) == len(slice)
}

func testOperacionesFuncionalesSol() bool {
	numeros := []int{1, 2, 3, 4, 5}
	pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
	suma := reducir(numeros, 0, func(a, b int) int { return a + b })
	cuadrados := mapear([]int{1, 2, 3}, func(n int) int { return n * n })
	return len(pares) == 2 && suma == 15 && len(cuadrados) == 3 && cuadrados[2] == 9
}

func testMatrizIdentidadSol() bool {
	matriz := matrizIdentidad(3)
	if len(matriz) != 3 || len(matriz[0]) != 3 {
		return false
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j && matriz[i][j] != 1 {
				return false
			}
			if i != j && matriz[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func testMultiplicacionMatricesSol() bool {
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{5, 6}, {7, 8}}
	resultado, err := multiplicarMatrices(a, b)
	if err != nil {
		return false
	}
	// Resultado esperado: [[19, 22], [43, 50]]
	return resultado[0][0] == 19 && resultado[0][1] == 22 &&
		resultado[1][0] == 43 && resultado[1][1] == 50
}

func testColaCircularSol() bool {
	cola := NuevaColaCircular(3)

	// Test enqueue
	for i := 1; i <= 3; i++ {
		if !cola.Enqueue(i) {
			return false
		}
	}

	// Cola debe estar llena
	if !cola.EstaLlena() || cola.Enqueue(4) {
		return false
	}

	// Test dequeue
	for i := 1; i <= 3; i++ {
		elemento, ok := cola.Dequeue()
		if !ok || elemento != i {
			return false
		}
	}

	return cola.EstaVacia()
}

func testBufferDeslizanteSol() bool {
	buffer := NuevoBufferDeslizante(3)
	buffer.Agregar(1.0)
	buffer.Agregar(2.0)
	buffer.Agregar(3.0)

	promedio := buffer.Promedio()
	minimo := buffer.Minimo()
	maximo := buffer.Maximo()

	return promedio == 2.0 && minimo == 1.0 && maximo == 3.0
}

func testEliminarDuplicadosSol() bool {
	slice := []int{1, 2, 2, 3, 4, 4, 4, 5}
	resultado := eliminarDuplicados(slice)
	esperado := []int{1, 2, 3, 4, 5}

	if len(resultado) != len(esperado) {
		return false
	}

	for i, v := range esperado {
		if resultado[i] != v {
			return false
		}
	}

	return true
}

func testRotacionSol() bool {
	slice := []int{1, 2, 3, 4, 5}
	rotarDerecha(slice, 2)
	esperado := []int{4, 5, 1, 2, 3}

	for i, v := range esperado {
		if slice[i] != v {
			return false
		}
	}

	return true
}

func testHistogramaSol() bool {
	hist := NuevoHistograma(0.0, 10.0, 5)
	hist.AgregarValor(1.0)
	hist.AgregarValor(2.0)
	hist.AgregarValor(2.5)

	freqs := hist.ObtenerFrecuencias()
	return freqs[0] == 1 && freqs[1] == 2 // Primer bin: 1 valor, segundo bin: 2 valores
}

func testKMPSol() bool {
	texto := "ABABCABABA"
	patron := "ABAB"
	indices := buscarPatronKMP(texto, patron)
	return len(indices) == 2 && indices[0] == 0 && indices[1] == 6
}
