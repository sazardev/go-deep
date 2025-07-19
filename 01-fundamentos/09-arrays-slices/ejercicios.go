// 📚 Ejercicios: Arrays y Slices
// ===============================
// 
// Estos ejercicios te ayudarán a dominar arrays y slices en Go.
// Completa cada función siguiendo las especificaciones.

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// ===== EJERCICIO 1: MANIPULACIÓN BÁSICA =====

// 1.1 - Implementa una función que reciba un array de enteros y devuelva
// la suma, promedio, valor máximo y mínimo
func analizarArray(numeros [10]int) (suma int, promedio float64, maximo int, minimo int) {
	// TODO: Implementar análisis del array
	// Hint: Inicializa maximo con el primer elemento y minimo también
	// Calcula la suma iterando y encuentra max/min comparando
	return
}

// 1.2 - Crea una función que intercambie los elementos de un slice
// en las posiciones i y j, validando que los índices sean válidos
func intercambiarElementos(slice []int, i, j int) error {
	// TODO: Validar índices y intercambiar elementos
	// Retorna error si los índices están fuera de rango
	return nil
}

// 1.3 - Implementa una función que invierta un slice in-place
func invertirSlice(slice []int) {
	// TODO: Invertir el slice modificando el original
	// Hint: Intercambia elementos desde los extremos hacia el centro
}

// ===== EJERCICIO 2: ALGORITMOS DE BÚSQUEDA =====

// 2.1 - Implementa búsqueda lineal que devuelva TODOS los índices
// donde se encuentra el valor buscado
func busquedaLinealTodos(slice []int, valor int) []int {
	// TODO: Devolver slice con todos los índices donde aparece el valor
	// Si no se encuentra, devolver slice vacío
	return nil
}

// 2.2 - Implementa búsqueda binaria recursiva
func busquedaBinariaRecursiva(slice []int, valor, inicio, fin int) int {
	// TODO: Implementar búsqueda binaria usando recursión
	// Asume que el slice está ordenado
	// Retorna -1 si no encuentra el elemento
	return -1
}

// 2.3 - Encuentra el primer elemento que cumple una condición
func encontrarPrimero(slice []int, condicion func(int) bool) (int, bool) {
	// TODO: Devolver el primer elemento que cumple la condición
	// Retorna (elemento, true) si lo encuentra, (0, false) si no
	return 0, false
}

// ===== EJERCICIO 3: ALGORITMOS DE ORDENAMIENTO =====

// 3.1 - Implementa Selection Sort
func selectionSort(slice []int) {
	// TODO: Ordenar el slice usando selection sort
	// Modifica el slice original
}

// 3.2 - Implementa Insertion Sort
func insertionSort(slice []int) {
	// TODO: Ordenar el slice usando insertion sort
	// Modifica el slice original
}

// 3.3 - Implementa Merge Sort
func mergeSort(slice []int) []int {
	// TODO: Implementar merge sort
	// Retorna un nuevo slice ordenado
	return nil
}

// Función auxiliar para merge sort
func merge(izquierda, derecha []int) []int {
	// TODO: Combinar dos slices ordenados en uno ordenado
	return nil
}

// ===== EJERCICIO 4: OPERACIONES FUNCIONALES =====

// 4.1 - Implementa una función Filter genérica para slices de enteros
func filtrar(slice []int, predicado func(int) bool) []int {
	// TODO: Devolver nuevo slice con elementos que cumplan el predicado
	return nil
}

// 4.2 - Implementa una función Map para transformar elementos
func mapear(slice []int, transformacion func(int) int) []int {
	// TODO: Devolver nuevo slice con elementos transformados
	return nil
}

// 4.3 - Implementa Reduce para reducir slice a un valor
func reducir(slice []int, inicial int, operacion func(int, int) int) int {
	// TODO: Aplicar operación acumulativa a todos los elementos
	return inicial
}

// 4.4 - Implementa una función que combine filter y map
func filtrarYMapear(slice []int, predicado func(int) bool, transformacion func(int) int) []int {
	// TODO: Filtrar elementos y luego transformarlos
	return nil
}

// ===== EJERCICIO 5: MATRICES MULTIDIMENSIONALES =====

// 5.1 - Crea una función que genere una matriz identidad de tamaño n
func matrizIdentidad(n int) [][]int {
	// TODO: Crear matriz n x n con 1s en la diagonal y 0s en el resto
	return nil
}

// 5.2 - Implementa multiplicación de matrices
func multiplicarMatrices(a, b [][]int) ([][]int, error) {
	// TODO: Multiplicar matrices A y B
	// Validar que las dimensiones sean compatibles
	// Retorna error si no se pueden multiplicar
	return nil, nil
}

// 5.3 - Implementa transposición de matriz
func transponerMatriz(matriz [][]int) [][]int {
	// TODO: Devolver la matriz transpuesta
	return nil
}

// 5.4 - Encuentra el elemento máximo en una matriz y su posición
func encontrarMaximoMatriz(matriz [][]int) (valor, fila, columna int) {
	// TODO: Devolver valor máximo y su posición [fila, columna]
	return
}

// ===== EJERCICIO 6: ESTRUCTURAS DE DATOS AVANZADAS =====

// 6.1 - Implementa una Cola Circular usando slice
type ColaCircular struct {
	datos     []interface{}
	frente    int
	final     int
	tamaño    int
	capacidad int
}

func NuevaColaCircular(capacidad int) *ColaCircular {
	// TODO: Inicializar cola circular
	return nil
}

func (c *ColaCircular) Enqueue(elemento interface{}) bool {
	// TODO: Agregar elemento a la cola
	// Retorna false si la cola está llena
	return false
}

func (c *ColaCircular) Dequeue() (interface{}, bool) {
	// TODO: Remover y devolver elemento del frente
	// Retorna (nil, false) si la cola está vacía
	return nil, false
}

func (c *ColaCircular) EstaLlena() bool {
	// TODO: Verificar si la cola está llena
	return false
}

func (c *ColaCircular) EstaVacia() bool {
	// TODO: Verificar si la cola está vacía
	return true
}

// 6.2 - Implementa un Buffer Deslizante para calcular promedios móviles
type BufferDeslizante struct {
	datos   []float64
	indice  int
	lleno   bool
	tamaño  int
}

func NuevoBufferDeslizante(tamaño int) *BufferDeslizante {
	// TODO: Inicializar buffer deslizante
	return nil
}

func (b *BufferDeslizante) Agregar(valor float64) {
	// TODO: Agregar valor al buffer
}

func (b *BufferDeslizante) Promedio() float64 {
	// TODO: Calcular promedio de los valores en el buffer
	return 0
}

func (b *BufferDeslizante) Maximo() float64 {
	// TODO: Encontrar valor máximo en el buffer
	return 0
}

func (b *BufferDeslizante) Minimo() float64 {
	// TODO: Encontrar valor mínimo en el buffer
	return 0
}

// ===== EJERCICIO 7: ALGORITMOS AVANZADOS =====

// 7.1 - Implementa el algoritmo de partición de Quicksort (partición de Lomuto)
func particionLomuto(slice []int, bajo, alto int) int {
	// TODO: Implementar partición de Lomuto
	// Usar el último elemento como pivote
	return 0
}

// 7.2 - Implementa búsqueda del k-ésimo elemento más pequeño
func kEsimoMenor(slice []int, k int) int {
	// TODO: Encontrar el k-ésimo elemento más pequeño sin ordenar completamente
	// Usar algoritmo Quickselect
	return 0
}

// 7.3 - Implementa algoritmo para encontrar el par de elementos con suma más cercana a un objetivo
func parSumaCercana(slice []int, objetivo int) (int, int, int) {
	// TODO: Encontrar par de elementos cuya suma esté más cerca del objetivo
	// Retorna (elemento1, elemento2, diferencia_absoluta)
	return 0, 0, math.MaxInt32
}

// 7.4 - Implementa sliding window maximum
func maximoVentanaDeslizante(slice []int, tamaño int) []int {
	// TODO: Para cada ventana de tamaño k, encontrar el máximo
	// Retorna slice con los máximos de cada ventana
	return nil
}

// ===== EJERCICIO 8: OPTIMIZACIÓN Y RENDIMIENTO =====

// 8.1 - Implementa una función que elimine duplicados manteniendo el orden
func eliminarDuplicados(slice []int) []int {
	// TODO: Eliminar duplicados eficientemente
	// Mantener el primer occurrence de cada elemento
	return nil
}

// 8.2 - Implementa intersección de dos slices ordenados
func interseccionOrdenados(slice1, slice2 []int) []int {
	// TODO: Encontrar elementos comunes en slices ordenados
	// Usar algoritmo de dos punteros para O(n+m)
	return nil
}

// 8.3 - Implementa una función que rote un slice k posiciones a la derecha
func rotarDerecha(slice []int, k int) {
	// TODO: Rotar slice k posiciones a la derecha in-place
	// Manejar casos donde k > len(slice)
}

// 8.4 - Implementa compresión Run-Length para slice de enteros
func compresionRunLength(slice []int) []int {
	// TODO: Comprimir secuencias consecutivas
	// Retorna [valor1, cantidad1, valor2, cantidad2, ...]
	return nil
}

// ===== EJERCICIO 9: CASOS DE USO PRÁCTICOS =====

// 9.1 - Implementa un sistema de histograma para analizar frecuencias
type Histograma struct {
	bins     []int
	min      float64
	max      float64
	anchoBin float64
}

func NuevoHistograma(min, max float64, numBins int) *Histograma {
	// TODO: Inicializar histograma
	return nil
}

func (h *Histograma) AgregarValor(valor float64) {
	// TODO: Agregar valor al bin correspondiente
}

func (h *Histograma) ObtenerFrecuencias() []int {
	// TODO: Devolver frecuencias de cada bin
	return nil
}

func (h *Histograma) BinMasFrecuente() int {
	// TODO: Devolver índice del bin con mayor frecuencia
	return 0
}

// 9.2 - Implementa un analizador de tendencias para datos de series temporales
type AnalizadorTendencias struct {
	datos []float64
}

func NuevoAnalizadorTendencias() *AnalizadorTendencias {
	// TODO: Inicializar analizador
	return nil
}

func (a *AnalizadorTendencias) AgregarDato(valor float64) {
	// TODO: Agregar nuevo dato a la serie
}

func (a *AnalizadorTendencias) PromedioMovil(ventana int) []float64 {
	// TODO: Calcular promedio móvil con ventana especificada
	return nil
}

func (a *AnalizadorTendencias) DetectarTendencia() string {
	// TODO: Detectar si la tendencia es "ascendente", "descendente" o "estable"
	// Usar regresión lineal simple o comparación de promedios
	return "estable"
}

func (a *AnalizadorTendencias) EncontrarPicos() []int {
	// TODO: Encontrar índices de picos (máximos locales)
	return nil
}

// ===== EJERCICIO 10: ALGORITMOS DE CADENAS CON SLICES =====

// 10.1 - Implementa algoritmo de búsqueda de patrón KMP (Knuth-Morris-Pratt)
func buscarPatronKMP(texto, patron string) []int {
	// TODO: Encontrar todas las ocurrencias del patrón en el texto
	// Usar algoritmo KMP para eficiencia O(n+m)
	return nil
}

// Función auxiliar para KMP
func construirTablaLPS(patron string) []int {
	// TODO: Construir tabla de Longest Proper Prefix que es también Suffix
	return nil
}

// 10.2 - Implementa algoritmo para encontrar la subsecuencia común más larga
func subsecuenciaComunMasLarga(s1, s2 string) string {
	// TODO: Encontrar LCS usando programación dinámica
	return ""
}

// ===== FUNCIONES DE DEMOSTRACIÓN =====

func main() {
	fmt.Println("🧪 EJERCICIOS: Arrays y Slices")
	fmt.Println("===============================")
	
	// Aquí puedes probar tus implementaciones
	demostrarEjercicios()
}

func demostrarEjercicios() {
	fmt.Println("\n📊 Ejecutando demostraciones...")
	
	// Demo Ejercicio 1: Análisis de array
	fmt.Println("\n1. Análisis de Array:")
	numeros := [10]int{5, 2, 8, 1, 9, 3, 7, 4, 6, 10}
	suma, promedio, maximo, minimo := analizarArray(numeros)
	fmt.Printf("Array: %v\n", numeros)
	fmt.Printf("Suma: %d, Promedio: %.2f, Max: %d, Min: %d\n", suma, promedio, maximo, minimo)
	
	// Demo Ejercicio 2: Búsqueda
	fmt.Println("\n2. Búsqueda Lineal:")
	slice := []int{1, 3, 5, 3, 7, 3, 9}
	indices := busquedaLinealTodos(slice, 3)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Índices donde aparece 3: %v\n", indices)
	
	// Demo Ejercicio 3: Ordenamiento
	fmt.Println("\n3. Algoritmos de Ordenamiento:")
	datos := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", datos)
	
	copia := make([]int, len(datos))
	copy(copia, datos)
	selectionSort(copia)
	fmt.Printf("Selection Sort: %v\n", copia)
	
	copia = make([]int, len(datos))
	copy(copia, datos)
	insertionSort(copia)
	fmt.Printf("Insertion Sort: %v\n", copia)
	
	resultado := mergeSort(datos)
	fmt.Printf("Merge Sort: %v\n", resultado)
	
	// Demo Ejercicio 4: Operaciones Funcionales
	fmt.Println("\n4. Operaciones Funcionales:")
	numeros2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	pares := filtrar(numeros2, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Números pares: %v\n", pares)
	
	cuadrados := mapear(numeros2, func(n int) int { return n * n })
	fmt.Printf("Cuadrados: %v\n", cuadrados)
	
	suma2 := reducir(numeros2, 0, func(a, b int) int { return a + b })
	fmt.Printf("Suma total: %d\n", suma2)
	
	// Demo Ejercicio 5: Matrices
	fmt.Println("\n5. Matrices:")
	identidad := matrizIdentidad(3)
	fmt.Printf("Matriz identidad 3x3:\n")
	imprimirMatriz(identidad)
	
	// Demo Ejercicio 6: Estructuras Avanzadas
	fmt.Println("\n6. Cola Circular:")
	cola := NuevaColaCircular(5)
	for i := 1; i <= 3; i++ {
		cola.Enqueue(i)
		fmt.Printf("Enqueue %d\n", i)
	}
	
	for !cola.EstaVacia() {
		elemento, _ := cola.Dequeue()
		fmt.Printf("Dequeue: %v\n", elemento)
	}
	
	// Demo Ejercicio 7: Buffer Deslizante
	fmt.Println("\n7. Buffer Deslizante:")
	buffer := NuevoBufferDeslizante(3)
	valores := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	
	for _, v := range valores {
		buffer.Agregar(v)
		fmt.Printf("Agregado %.1f - Promedio: %.2f\n", v, buffer.Promedio())
	}
	
	// Demo Ejercicio 8: Optimización
	fmt.Println("\n8. Eliminación de Duplicados:")
	conDuplicados := []int{1, 2, 2, 3, 4, 4, 4, 5}
	sinDuplicados := eliminarDuplicados(conDuplicados)
	fmt.Printf("Original: %v\n", conDuplicados)
	fmt.Printf("Sin duplicados: %v\n", sinDuplicados)
	
	// Demo Ejercicio 9: Histograma
	fmt.Println("\n9. Histograma:")
	hist := NuevoHistograma(0.0, 10.0, 5)
	datosRandom := []float64{1.2, 3.4, 5.6, 7.8, 2.1, 4.3, 6.5, 8.7, 1.9}
	
	for _, v := range datosRandom {
		hist.AgregarValor(v)
	}
	
	frecuencias := hist.ObtenerFrecuencias()
	fmt.Printf("Frecuencias por bin: %v\n", frecuencias)
	fmt.Printf("Bin más frecuente: %d\n", hist.BinMasFrecuente())
	
	fmt.Println("\n✅ Demo completada. Implementa las funciones para ver resultados reales.")
}

// Función auxiliar para imprimir matrices
func imprimirMatriz(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Printf("%v\n", fila)
	}
}

// ===== BENCHMARKS Y PRUEBAS DE RENDIMIENTO =====

func benchmarkOrdenamiento() {
	fmt.Println("\n⏱️  BENCHMARK: Algoritmos de Ordenamiento")
	fmt.Println("==========================================")
	
	tamaños := []int{1000, 5000, 10000}
	
	for _, n := range tamaños {
		fmt.Printf("\nTamaño: %d elementos\n", n)
		
		// Generar datos aleatorios
		datos := make([]int, n)
		for i := range datos {
			datos[i] = rand.Intn(n)
		}
		
		// Benchmark Selection Sort
		copia := make([]int, len(datos))
		copy(copia, datos)
		inicio := time.Now()
		selectionSort(copia)
		tiempoSelection := time.Since(inicio)
		
		// Benchmark Insertion Sort
		copy(copia, datos)
		inicio = time.Now()
		insertionSort(copia)
		tiempoInsertion := time.Since(inicio)
		
		// Benchmark Merge Sort
		inicio = time.Now()
		mergeSort(datos)
		tiempoMerge := time.Since(inicio)
		
		// Benchmark Sort nativo de Go
		copy(copia, datos)
		inicio = time.Now()
		sort.Ints(copia)
		tiempoNativo := time.Since(inicio)
		
		fmt.Printf("Selection Sort: %v\n", tiempoSelection)
		fmt.Printf("Insertion Sort: %v\n", tiempoInsertion)
		fmt.Printf("Merge Sort:     %v\n", tiempoMerge)
		fmt.Printf("Go Sort:        %v\n", tiempoNativo)
	}
}

// ===== VALIDADORES Y TESTS =====

func validarImplementaciones() {
	fmt.Println("\n🧪 VALIDACIÓN: Implementaciones")
	fmt.Println("================================")
	
	// Validar que las funciones están implementadas
	tests := []struct {
		nombre string
		test   func() bool
	}{
		{"Análisis Array", testAnalizarArray},
		{"Búsqueda Lineal", testBusquedaLineal},
		{"Selection Sort", testSelectionSort},
		{"Operaciones Funcionales", testOperacionesFuncionales},
		{"Matriz Identidad", testMatrizIdentidad},
		{"Cola Circular", testColaCircular},
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
	
	fmt.Printf("\nResultado: %d/%d tests pasados\n", pasados, len(tests))
}

func testAnalizarArray() bool {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	suma, promedio, maximo, minimo := analizarArray(array)
	return suma == 55 && promedio == 5.5 && maximo == 10 && minimo == 1
}

func testBusquedaLineal() bool {
	slice := []int{1, 3, 5, 3, 7, 3, 9}
	indices := busquedaLinealTodos(slice, 3)
	return len(indices) == 3 && indices[0] == 1 && indices[1] == 3 && indices[2] == 5
}

func testSelectionSort() bool {
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(slice)
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

func testOperacionesFuncionales() bool {
	numeros := []int{1, 2, 3, 4, 5}
	pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
	suma := reducir(numeros, 0, func(a, b int) int { return a + b })
	return len(pares) == 2 && suma == 15
}

func testMatrizIdentidad() bool {
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

func testColaCircular() bool {
	cola := NuevaColaCircular(3)
	if cola == nil {
		return false
	}
	
	// Test básico
	cola.Enqueue(1)
	cola.Enqueue(2)
	
	elemento, ok := cola.Dequeue()
	return ok && elemento == 1
}
