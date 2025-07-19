# 📚 Lección 09: Arrays y Slices

## 🎯 Objetivos de Aprendizaje

Al completar esta lección, serás capaz de:

- Entender la diferencia entre arrays y slices en Go
- Crear y manipular arrays de tamaño fijo
- Trabajar con slices dinámicos eficientemente
- Aplicar operaciones avanzadas como append, copy y make
- Implementar algoritmos de búsqueda y ordenamiento
- Usar slices multidimensionales para estructuras complejas
- Optimizar el rendimiento en operaciones con colecciones

## 📖 Contenido

### 1. Arrays en Go

Los **arrays** en Go son colecciones de elementos del mismo tipo con un **tamaño fijo** determinado en tiempo de compilación.

#### Características de los Arrays:
- Tamaño fijo definido en tiempo de compilación
- Elementos del mismo tipo
- Acceso por índice (0-based)
- Tipos de valor (se copian por valor)
- Memoria contigua

#### Declaración y Inicialización

```go
// Declaración básica
var numeros [5]int                    // Array de 5 enteros, inicializado con ceros
var nombres [3]string                 // Array de 3 strings, inicializado con ""

// Declaración con inicialización
var edades = [4]int{25, 30, 22, 28}   // Array con valores específicos
var colores = [3]string{"rojo", "verde", "azul"}

// Inicialización con tamaño inferido
ciudades := [...]string{"Madrid", "Barcelona", "Valencia"}  // Go infiere el tamaño

// Inicialización con índices específicos
var diasSemana = [7]string{
    0: "Domingo",
    1: "Lunes", 
    6: "Sábado",
    // Los demás se inicializan con ""
}
```

#### Operaciones Básicas con Arrays

```go
func ejemplosArrays() {
    // Crear array
    var puntuaciones [5]float64
    
    // Asignar valores
    puntuaciones[0] = 8.5
    puntuaciones[1] = 9.2
    puntuaciones[4] = 7.8
    
    // Acceder a elementos
    fmt.Printf("Primera puntuación: %.1f\n", puntuaciones[0])
    fmt.Printf("Última puntuación: %.1f\n", puntuaciones[4])
    
    // Obtener longitud
    fmt.Printf("Número de puntuaciones: %d\n", len(puntuaciones))
    
    // Iterar sobre el array
    for i := 0; i < len(puntuaciones); i++ {
        fmt.Printf("Puntuación %d: %.1f\n", i+1, puntuaciones[i])
    }
    
    // Iterar con range
    for indice, valor := range puntuaciones {
        fmt.Printf("Índice %d: %.1f\n", indice, valor)
    }
}
```

#### Arrays Multidimensionales

```go
func arraysMultidimensionales() {
    // Matriz 3x3
    var matriz [3][3]int
    
    // Inicialización directa
    tablero := [3][3]string{
        {"X", "O", "X"},
        {"O", "X", "O"},
        {"X", "O", "X"},
    }
    
    // Asignar valores
    matriz[0][0] = 1
    matriz[1][1] = 5
    matriz[2][2] = 9
    
    // Iterar matriz
    for i := 0; i < len(matriz); i++ {
        for j := 0; j < len(matriz[i]); j++ {
            fmt.Printf("%d ", matriz[i][j])
        }
        fmt.Println()
    }
    
    // Con range
    for i, fila := range tablero {
        for j, celda := range fila {
            fmt.Printf("Posición [%d,%d]: %s\n", i, j, celda)
        }
    }
}
```

### 2. Slices en Go

Los **slices** son la estructura de datos más utilizada en Go para colecciones dinámicas. Son más flexibles que los arrays.

#### Características de los Slices:
- Tamaño dinámico
- Referencias a arrays subyacentes
- Tres componentes: puntero, longitud y capacidad
- Tipos de referencia
- Más eficientes que arrays para la mayoría de casos

#### Anatomía de un Slice

```go
// Un slice tiene tres componentes:
// 1. Puntero al primer elemento del array subyacente
// 2. Longitud (len): número de elementos actuales
// 3. Capacidad (cap): número máximo de elementos sin reasignar memoria

func anatomiaSlice() {
    // Crear slice
    numeros := []int{1, 2, 3, 4, 5}
    
    fmt.Printf("Slice: %v\n", numeros)
    fmt.Printf("Longitud: %d\n", len(numeros))
    fmt.Printf("Capacidad: %d\n", cap(numeros))
    
    // Sub-slice
    subSlice := numeros[1:4]  // elementos del índice 1 al 3
    fmt.Printf("Sub-slice: %v\n", subSlice)
    fmt.Printf("Longitud sub-slice: %d\n", len(subSlice))
    fmt.Printf("Capacidad sub-slice: %d\n", cap(subSlice))
}
```

#### Creación de Slices

```go
func creacionSlices() {
    // 1. Slice literal
    frutas := []string{"manzana", "banana", "naranja"}
    
    // 2. Desde un array
    array := [5]int{1, 2, 3, 4, 5}
    slice1 := array[:]      // Todo el array
    slice2 := array[1:4]    // Elementos 1, 2, 3
    slice3 := array[:3]     // Primeros 3 elementos
    slice4 := array[2:]     // Desde el índice 2 hasta el final
    
    // 3. Con make()
    slice5 := make([]int, 5)       // longitud 5, capacidad 5
    slice6 := make([]int, 3, 10)   // longitud 3, capacidad 10
    slice7 := make([]string, 0, 5) // longitud 0, capacidad 5
    
    // 4. Slice nulo vs vacío
    var sliceNulo []int            // slice nulo (nil)
    sliceVacio := []int{}          // slice vacío pero no nil
    sliceVacio2 := make([]int, 0)  // slice vacío pero no nil
    
    fmt.Printf("Slice nulo: %v, es nil: %t\n", sliceNulo, sliceNulo == nil)
    fmt.Printf("Slice vacío: %v, es nil: %t\n", sliceVacio, sliceVacio == nil)
}
```

#### Operaciones con Slices

```go
func operacionesSlices() {
    // Append - Agregar elementos
    var numeros []int
    numeros = append(numeros, 1)              // [1]
    numeros = append(numeros, 2, 3, 4)        // [1, 2, 3, 4]
    
    // Append de otro slice
    masNumeros := []int{5, 6, 7}
    numeros = append(numeros, masNumeros...)  // [1, 2, 3, 4, 5, 6, 7]
    
    // Copy - Copiar elementos
    origen := []int{1, 2, 3, 4, 5}
    destino := make([]int, 3)
    copiados := copy(destino, origen)         // Copia 3 elementos
    fmt.Printf("Elementos copiados: %d\n", copiados)
    fmt.Printf("Destino: %v\n", destino)     // [1, 2, 3]
    
    // Insertar elemento en posición específica
    slice := []int{1, 2, 4, 5}
    pos := 2
    valor := 3
    // Insertar 3 en posición 2
    slice = append(slice[:pos], append([]int{valor}, slice[pos:]...)...)
    fmt.Printf("Después de insertar: %v\n", slice)  // [1, 2, 3, 4, 5]
    
    // Eliminar elemento
    slice = []int{1, 2, 3, 4, 5}
    posEliminar := 2
    slice = append(slice[:posEliminar], slice[posEliminar+1:]...)
    fmt.Printf("Después de eliminar: %v\n", slice)  // [1, 2, 4, 5]
}
```

### 3. Operaciones Avanzadas

#### Slicing Avanzado

```go
func slicingAvanzado() {
    datos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    // Sintaxis: slice[inicio:fin:capacidad]
    s1 := datos[2:5]     // [2, 3, 4] - longitud 3, capacidad 8
    s2 := datos[2:5:6]   // [2, 3, 4] - longitud 3, capacidad 4
    
    fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
    fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
    
    // Compartir array subyacente
    original := []int{1, 2, 3, 4, 5}
    slice1 := original[1:4]  // [2, 3, 4]
    slice2 := original[0:3]  // [1, 2, 3]
    
    // Modificar slice1 afecta original y slice2
    slice1[0] = 99  // Cambia el elemento en índice 1 del original
    fmt.Printf("Original: %v\n", original)  // [1, 99, 3, 4, 5]
    fmt.Printf("Slice2: %v\n", slice2)      // [1, 99, 3]
}
```

#### Redimensionamiento y Capacidad

```go
func redimensionamientoCapacidad() {
    // Entender el crecimiento de capacidad
    var slice []int
    
    for i := 0; i < 10; i++ {
        slice = append(slice, i)
        fmt.Printf("len: %d, cap: %d, slice: %v\n", 
                   len(slice), cap(slice), slice)
    }
    
    // Preasignar capacidad para mejor rendimiento
    slice2 := make([]int, 0, 100)  // Capacidad inicial de 100
    for i := 0; i < 50; i++ {
        slice2 = append(slice2, i)
        // No habrá reasignaciones hasta llenar la capacidad
    }
    
    // Reducir slice manteniendo capacidad
    slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Printf("Antes: len=%d, cap=%d\n", len(slice3), cap(slice3))
    
    slice3 = slice3[:5]  // Reduce longitud pero mantiene capacidad
    fmt.Printf("Después: len=%d, cap=%d\n", len(slice3), cap(slice3))
    
    // Forzar nueva asignación para liberar memoria
    slice4 := make([]int, len(slice3))
    copy(slice4, slice3)
    slice3 = slice4  // Ahora slice3 tiene capacidad exacta
    fmt.Printf("Optimizado: len=%d, cap=%d\n", len(slice3), cap(slice3))
}
```

### 4. Algoritmos Comunes

#### Búsqueda

```go
func algoritmosBusqueda() {
    datos := []int{3, 7, 1, 9, 4, 2, 8, 5, 6}
    
    // Búsqueda lineal
    buscarLineal := func(slice []int, valor int) int {
        for i, v := range slice {
            if v == valor {
                return i
            }
        }
        return -1
    }
    
    indice := buscarLineal(datos, 9)
    fmt.Printf("Búsqueda lineal - Índice de 9: %d\n", indice)
    
    // Búsqueda binaria (requiere slice ordenado)
    datosOrdenados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    buscarBinario := func(slice []int, valor int) int {
        izquierda, derecha := 0, len(slice)-1
        
        for izquierda <= derecha {
            medio := (izquierda + derecha) / 2
            
            if slice[medio] == valor {
                return medio
            } else if slice[medio] < valor {
                izquierda = medio + 1
            } else {
                derecha = medio - 1
            }
        }
        return -1
    }
    
    indice = buscarBinario(datosOrdenados, 6)
    fmt.Printf("Búsqueda binaria - Índice de 6: %d\n", indice)
}
```

#### Ordenamiento

```go
func algoritmosOrdenamiento() {
    // Bubble Sort
    bubbleSort := func(slice []int) {
        n := len(slice)
        for i := 0; i < n-1; i++ {
            for j := 0; j < n-i-1; j++ {
                if slice[j] > slice[j+1] {
                    slice[j], slice[j+1] = slice[j+1], slice[j]
                }
            }
        }
    }
    
    datos1 := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Printf("Antes de Bubble Sort: %v\n", datos1)
    bubbleSort(datos1)
    fmt.Printf("Después de Bubble Sort: %v\n", datos1)
    
    // Quick Sort
    var quickSort func([]int, int, int)
    quickSort = func(slice []int, bajo, alto int) {
        if bajo < alto {
            pi := particion(slice, bajo, alto)
            quickSort(slice, bajo, pi-1)
            quickSort(slice, pi+1, alto)
        }
    }
    
    datos2 := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Printf("Antes de Quick Sort: %v\n", datos2)
    quickSort(datos2, 0, len(datos2)-1)
    fmt.Printf("Después de Quick Sort: %v\n", datos2)
}

func particion(slice []int, bajo, alto int) int {
    pivot := slice[alto]
    i := bajo - 1
    
    for j := bajo; j < alto; j++ {
        if slice[j] < pivot {
            i++
            slice[i], slice[j] = slice[j], slice[i]
        }
    }
    slice[i+1], slice[alto] = slice[alto], slice[i+1]
    return i + 1
}
```

### 5. Slices Multidimensionales

```go
func slicesMultidimensionales() {
    // Slice de slices (matriz dinámica)
    matriz := make([][]int, 3)
    
    // Inicializar cada fila
    for i := range matriz {
        matriz[i] = make([]int, 4)  // 3x4 matriz
    }
    
    // Llenar matriz
    valor := 1
    for i := 0; i < len(matriz); i++ {
        for j := 0; j < len(matriz[i]); j++ {
            matriz[i][j] = valor
            valor++
        }
    }
    
    // Mostrar matriz
    for i, fila := range matriz {
        fmt.Printf("Fila %d: %v\n", i, fila)
    }
    
    // Matriz irregular (jagged array)
    triangulo := make([][]int, 5)
    for i := range triangulo {
        triangulo[i] = make([]int, i+1)  // Fila i tiene i+1 elementos
        for j := range triangulo[i] {
            triangulo[i][j] = i + j
        }
    }
    
    fmt.Println("\nMatriz triangular:")
    for i, fila := range triangulo {
        fmt.Printf("Fila %d: %v\n", i, fila)
    }
}
```

### 6. Optimización de Rendimiento

```go
func optimizacionRendimiento() {
    // 1. Preasignar capacidad cuando se conoce el tamaño
    n := 10000
    
    // Ineficiente - múltiples reasignaciones
    var slice1 []int
    for i := 0; i < n; i++ {
        slice1 = append(slice1, i)
    }
    
    // Eficiente - capacidad preasignada
    slice2 := make([]int, 0, n)
    for i := 0; i < n; i++ {
        slice2 = append(slice2, i)
    }
    
    // Más eficiente - asignar directamente
    slice3 := make([]int, n)
    for i := 0; i < n; i++ {
        slice3[i] = i
    }
    
    // 2. Evitar copias innecesarias
    datos := make([]int, 1000000)
    
    // Ineficiente - copia todo el slice
    procesarIneficiente := func(slice []int) {
        copia := make([]int, len(slice))
        copy(copia, slice)
        // Procesar copia...
    }
    
    // Eficiente - trabajar con el slice original
    procesarEficiente := func(slice []int) {
        // Procesar directamente el slice...
        for i, v := range slice {
            datos[i] = v * 2
        }
    }
    
    procesarEficiente(datos)
    
    // 3. Reutilizar slices
    buffer := make([]byte, 1024)  // Buffer reutilizable
    
    procesarDatos := func(datos [][]byte) {
        for _, data := range datos {
            // Limpiar buffer
            for i := range buffer {
                buffer[i] = 0
            }
            // Usar buffer para procesar data
            copy(buffer, data)
            // Procesar...
        }
    }
    
    // Uso del buffer reutilizable
    lotesDatos := [][]byte{
        []byte("datos1"),
        []byte("datos2"), 
        []byte("datos3"),
    }
    procesarDatos(lotesDatos)
}
```

### 7. Patrones Comunes

#### Filter, Map, Reduce

```go
func patronesFuncionales() {
    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter - Filtrar elementos
    filtrar := func(slice []int, predicado func(int) bool) []int {
        var resultado []int
        for _, v := range slice {
            if predicado(v) {
                resultado = append(resultado, v)
            }
        }
        return resultado
    }
    
    pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
    fmt.Printf("Números pares: %v\n", pares)
    
    // Map - Transformar elementos
    mapear := func(slice []int, funcion func(int) int) []int {
        resultado := make([]int, len(slice))
        for i, v := range slice {
            resultado[i] = funcion(v)
        }
        return resultado
    }
    
    cuadrados := mapear(numeros, func(n int) int { return n * n })
    fmt.Printf("Cuadrados: %v\n", cuadrados)
    
    // Reduce - Reducir a un valor
    reducir := func(slice []int, inicial int, funcion func(int, int) int) int {
        resultado := inicial
        for _, v := range slice {
            resultado = funcion(resultado, v)
        }
        return resultado
    }
    
    suma := reducir(numeros, 0, func(acc, n int) int { return acc + n })
    producto := reducir(numeros, 1, func(acc, n int) int { return acc * n })
    
    fmt.Printf("Suma: %d\n", suma)
    fmt.Printf("Producto: %d\n", producto)
}
```

#### Chunk y Batch Processing

```go
func patronesChunk() {
    datos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
    
    // Dividir slice en chunks
    chunk := func(slice []int, tamaño int) [][]int {
        var chunks [][]int
        for i := 0; i < len(slice); i += tamaño {
            fin := i + tamaño
            if fin > len(slice) {
                fin = len(slice)
            }
            chunks = append(chunks, slice[i:fin])
        }
        return chunks
    }
    
    chunks := chunk(datos, 4)
    fmt.Println("Datos divididos en chunks:")
    for i, c := range chunks {
        fmt.Printf("Chunk %d: %v\n", i+1, c)
    }
    
    // Procesar en lotes
    procesarLote := func(lote []int) int {
        suma := 0
        for _, v := range lote {
            suma += v
        }
        return suma
    }
    
    fmt.Println("\nProcesamiento por lotes:")
    for i, lote := range chunks {
        resultado := procesarLote(lote)
        fmt.Printf("Lote %d - Suma: %d\n", i+1, resultado)
    }
}
```

### 8. Casos de Uso Avanzados

#### Buffer Circular

```go
type BufferCircular struct {
    datos    []interface{}
    inicio   int
    tamaño   int
    capacidad int
}

func NuevoBufferCircular(capacidad int) *BufferCircular {
    return &BufferCircular{
        datos:     make([]interface{}, capacidad),
        capacidad: capacidad,
    }
}

func (b *BufferCircular) Agregar(elemento interface{}) {
    b.datos[(b.inicio+b.tamaño)%b.capacidad] = elemento
    if b.tamaño < b.capacidad {
        b.tamaño++
    } else {
        b.inicio = (b.inicio + 1) % b.capacidad
    }
}

func (b *BufferCircular) Obtener(indice int) interface{} {
    if indice >= b.tamaño {
        return nil
    }
    return b.datos[(b.inicio+indice)%b.capacidad]
}

func (b *BufferCircular) ToSlice() []interface{} {
    resultado := make([]interface{}, b.tamaño)
    for i := 0; i < b.tamaño; i++ {
        resultado[i] = b.Obtener(i)
    }
    return resultado
}

func demoBufferCircular() {
    buffer := NuevoBufferCircular(5)
    
    // Agregar elementos
    for i := 1; i <= 8; i++ {
        buffer.Agregar(i)
        fmt.Printf("Después de agregar %d: %v\n", i, buffer.ToSlice())
    }
}
```

#### Stack y Queue

```go
// Stack usando slice
type Stack []interface{}

func (s *Stack) Push(elemento interface{}) {
    *s = append(*s, elemento)
}

func (s *Stack) Pop() interface{} {
    if len(*s) == 0 {
        return nil
    }
    indice := len(*s) - 1
    elemento := (*s)[indice]
    *s = (*s)[:indice]
    return elemento
}

func (s *Stack) Peek() interface{} {
    if len(*s) == 0 {
        return nil
    }
    return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

// Queue usando slice
type Queue []interface{}

func (q *Queue) Enqueue(elemento interface{}) {
    *q = append(*q, elemento)
}

func (q *Queue) Dequeue() interface{} {
    if len(*q) == 0 {
        return nil
    }
    elemento := (*q)[0]
    *q = (*q)[1:]
    return elemento
}

func (q *Queue) Front() interface{} {
    if len(*q) == 0 {
        return nil
    }
    return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
    return len(*q) == 0
}

func demoStackQueue() {
    // Demo Stack
    var stack Stack
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    fmt.Printf("Stack peek: %v\n", stack.Peek())
    fmt.Printf("Stack pop: %v\n", stack.Pop())
    fmt.Printf("Stack pop: %v\n", stack.Pop())
    
    // Demo Queue
    var queue Queue
    queue.Enqueue("A")
    queue.Enqueue("B") 
    queue.Enqueue("C")
    
    fmt.Printf("Queue front: %v\n", queue.Front())
    fmt.Printf("Queue dequeue: %v\n", queue.Dequeue())
    fmt.Printf("Queue dequeue: %v\n", queue.Dequeue())
}
```

## 🎯 Mejores Prácticas

### 1. **Elección entre Arrays y Slices**
```go
// Usar arrays cuando:
// - El tamaño es fijo y conocido
// - Se necesita rendimiento máximo
// - Se trabaja con datos pequeños

var coordenadas [3]float64  // Punto 3D

// Usar slices cuando:
// - El tamaño es dinámico
// - Se necesita flexibilidad
// - Es el caso general

var usuarios []Usuario
```

### 2. **Preasignar Capacidad**
```go
// ❌ Ineficiente
var datos []int
for i := 0; i < 10000; i++ {
    datos = append(datos, i)
}

// ✅ Eficiente
datos := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    datos = append(datos, i)
}
```

### 3. **Evitar Memory Leaks**
```go
// ❌ Puede causar memory leak
func procesar() []int {
    datos := make([]int, 1000000)
    // ... llenar datos
    return datos[0:10]  // Mantiene referencia a todo el array
}

// ✅ Crear nueva asignación
func procesar() []int {
    datos := make([]int, 1000000)
    // ... llenar datos
    resultado := make([]int, 10)
    copy(resultado, datos[0:10])
    return resultado
}
```

### 4. **Chequear Bounds**
```go
func accederSeguro(slice []int, indice int) (int, bool) {
    if indice < 0 || indice >= len(slice) {
        return 0, false
    }
    return slice[indice], true
}
```

## ⚠️ Errores Comunes

### 1. **Modificación durante iteración**
```go
// ❌ Problemático
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    if v%2 == 0 {
        slice = append(slice[:i], slice[i+1:]...)  // Modifica durante iteración
    }
}

// ✅ Correcto
slice := []int{1, 2, 3, 4, 5}
var resultado []int
for _, v := range slice {
    if v%2 != 0 {  // Mantener impares
        resultado = append(resultado, v)
    }
}
```

### 2. **Slice de slice compartido**
```go
// ❌ Problema de referencia compartida
original := []int{1, 2, 3, 4, 5}
slice1 := original[0:3]
slice2 := original[2:5]
slice1[2] = 99  // Afecta a slice2 también

// ✅ Copias independientes
slice1 := make([]int, 3)
slice2 := make([]int, 3)
copy(slice1, original[0:3])
copy(slice2, original[2:5])
```

## 🧪 Ejercicios Prácticos

Completa los ejercicios en el archivo `ejercicios.go` para practicar:

1. **Manipulación básica** de arrays y slices
2. **Algoritmos de búsqueda** y ordenamiento
3. **Operaciones de filtrado** y transformación
4. **Estructuras de datos** avanzadas
5. **Optimización** de rendimiento
6. **Matrices multidimensionales**
7. **Patrones funcionales**
8. **Buffer circular** implementado con slices

## 🔗 Próxima Lección

En la **Lección 10: Maps y Structs**, aprenderemos sobre:
- Maps (diccionarios) para asociar claves y valores
- Structs para crear tipos de datos personalizados
- Embedding y composición
- Métodos en structs
- Interfaces básicas

## 📚 Recursos Adicionales

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)
- [Go Data Structures](https://research.swtch.com/godata)

---
**Recuerda**: Los slices son una de las características más poderosas de Go. Dominar su uso eficiente es fundamental para escribir código Go idiomático y performante.
