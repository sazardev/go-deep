package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// ============================================================================
// LECCIÓN 8: FUNCIONES - SOLUCIONES COMPLETAS
// ============================================================================

func main() {
	fmt.Println("🔧 SOLUCIONES DE EJERCICIOS DE FUNCIONES")
	fmt.Println("========================================")
	fmt.Println()

	solucion1()
	solucion2()
	solucion3()
	solucion4()
	solucion5()
	solucion6()
	solucion7()
	solucion8()
	solucion9()
	solucion10()
	
	// Demo adicional
	demoCompleto()
	demoPatronesAvanzados()
}

// ============================================================================
// SOLUCIÓN 1: FUNCIONES BÁSICAS
// ============================================================================

// 1.1 Función que calcula el área de un círculo
func calcularAreaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}

// 1.2 Función que determina si un número es par
func esPar(numero int) bool {
	return numero%2 == 0
}

// 1.3 Función que convierte Celsius a Fahrenheit
func celsiusAFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// 1.4 Función que encuentra el mayor de tres números
func mayor(a, b, c int) int {
	maximo := a
	if b > maximo {
		maximo = b
	}
	if c > maximo {
		maximo = c
	}
	return maximo
}

// 1.5 Función que calcula el factorial de un número
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func solucion1() {
	fmt.Println("✅ SOLUCIÓN 1: Funciones Básicas")
	fmt.Println()

	// Pruebas
	fmt.Printf("Área círculo radio 5: %.2f\n", calcularAreaCirculo(5))
	fmt.Printf("¿8 es par?: %t\n", esPar(8))
	fmt.Printf("¿7 es par?: %t\n", esPar(7))
	fmt.Printf("25°C en Fahrenheit: %.1f°F\n", celsiusAFahrenheit(25))
	fmt.Printf("Mayor entre 10, 25, 15: %d\n", mayor(10, 25, 15))
	fmt.Printf("Factorial de 5: %d\n", factorial(5))
	fmt.Printf("Factorial de 0: %d\n", factorial(0))
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 2: MÚLTIPLES VALORES DE RETORNO
// ============================================================================

// 2.1 Función que divide dos números y maneja el error
func dividirSeguro(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

// 2.2 Función que calcula cociente y residuo
func divisionCompleta(dividendo, divisor int) (cociente, residuo int) {
	if divisor == 0 {
		return 0, 0 // o podrías usar panic
	}
	cociente = dividendo / divisor
	residuo = dividendo % divisor
	return
}

// 2.3 Función que encuentra min y max en un slice
func minMax(numeros []int) (min, max int, err error) {
	if len(numeros) == 0 {
		return 0, 0, errors.New("slice vacío")
	}

	min = numeros[0]
	max = numeros[0]

	for _, num := range numeros {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max, nil
}

// 2.4 Función que valida email y extrae dominio
func validarEmail(email string) (valido bool, dominio string) {
	// Validación simple
	partes := strings.Split(email, "@")
	if len(partes) != 2 {
		return false, ""
	}

	usuario := partes[0]
	dominio = partes[1]

	if len(usuario) == 0 || len(dominio) == 0 {
		return false, ""
	}

	if !strings.Contains(dominio, ".") {
		return false, ""
	}

	return true, dominio
}

// 2.5 Función que calcula estadísticas básicas
func estadisticas(numeros []float64) (suma, promedio, minimo, maximo float64) {
	if len(numeros) == 0 {
		return 0, 0, 0, 0
	}

	suma = numeros[0]
	minimo = numeros[0]
	maximo = numeros[0]

	for i := 1; i < len(numeros); i++ {
		suma += numeros[i]
		if numeros[i] < minimo {
			minimo = numeros[i]
		}
		if numeros[i] > maximo {
			maximo = numeros[i]
		}
	}

	promedio = suma / float64(len(numeros))
	return
}

func solucion2() {
	fmt.Println("✅ SOLUCIÓN 2: Múltiples Valores de Retorno")
	fmt.Println()

	// Pruebas
	resultado, err := dividirSeguro(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", resultado)
	}

	_, err = dividirSeguro(10, 0)
	if err != nil {
		fmt.Printf("Error esperado: %v\n", err)
	}

	coc, res := divisionCompleta(17, 5)
	fmt.Printf("17 ÷ 5: cociente=%d, residuo=%d\n", coc, res)

	numeros := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max, err := minMax(numeros)
	if err == nil {
		fmt.Printf("En %v - Min: %d, Max: %d\n", numeros, min, max)
	}

	valido, dominio := validarEmail("usuario@ejemplo.com")
	fmt.Printf("Email válido: %t, dominio: %s\n", valido, dominio)

	valores := []float64{1.5, 2.8, 3.2, 1.1, 4.7}
	suma, prom, minVal, maxVal := estadisticas(valores)
	fmt.Printf("Estadísticas: suma=%.2f, promedio=%.2f, min=%.2f, max=%.2f\n", suma, prom, minVal, maxVal)
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 3: FUNCIONES VARIÁDICAS
// ============================================================================

// 3.1 Función que suma cualquier cantidad de números
func sumar(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// 3.2 Función que encuentra el promedio
func promedio(valores ...float64) float64 {
	if len(valores) == 0 {
		return 0
	}

	suma := 0.0
	for _, valor := range valores {
		suma += valor
	}
	return suma / float64(len(valores))
}

// 3.3 Función que concatena strings con separador
func concatenar(separador string, textos ...string) string {
	return strings.Join(textos, separador)
}

// 3.4 Función que imprime con formato personalizado
func imprimir(prefijo string, datos ...interface{}) {
	fmt.Print(prefijo + " ")
	for i, dato := range datos {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(dato)
	}
	fmt.Println()
}

// 3.5 Función que encuentra el máximo valor
func maximo(primero int, resto ...int) int {
	max := primero
	for _, num := range resto {
		if num > max {
			max = num
		}
	}
	return max
}

func solucion3() {
	fmt.Println("✅ SOLUCIÓN 3: Funciones Variádicas")
	fmt.Println()

	// Pruebas
	fmt.Printf("Suma: %d\n", sumar(1, 2, 3, 4, 5))
	fmt.Printf("Suma vacía: %d\n", sumar())
	fmt.Printf("Promedio: %.2f\n", promedio(1.5, 2.5, 3.5, 4.5))
	fmt.Printf("Concatenar: %s\n", concatenar(" - ", "Go", "es", "genial"))
	imprimir("[INFO]", "Usuario:", "Juan", "Edad:", 25, "Activo:", true)
	fmt.Printf("Máximo: %d\n", maximo(3, 7, 2, 9, 1, 8))

	// Pasar slice
	nums := []int{10, 20, 30}
	fmt.Printf("Suma desde slice: %d\n", sumar(nums...))
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 4: FUNCIONES COMO FIRST-CLASS CITIZENS
// ============================================================================

// 4.1 Tipo de función para operaciones matemáticas
type OperacionBinaria func(int, int) int

// 4.2 Función que aplica operación a dos números
func aplicarOperacion(a, b int, op OperacionBinaria) int {
	return op(a, b)
}

// 4.3 Función que filtra slice según predicado
func filtrar(numeros []int, predicado func(int) bool) []int {
	var resultado []int
	for _, num := range numeros {
		if predicado(num) {
			resultado = append(resultado, num)
		}
	}
	return resultado
}

// 4.4 Función que mapea elementos de un slice
func mapear(numeros []int, transformar func(int) int) []int {
	resultado := make([]int, len(numeros))
	for i, num := range numeros {
		resultado[i] = transformar(num)
	}
	return resultado
}

// 4.5 Función que reduce slice a un valor
func reducir(numeros []int, inicial int, operacion func(int, int) int) int {
	resultado := inicial
	for _, num := range numeros {
		resultado = operacion(resultado, num)
	}
	return resultado
}

func solucion4() {
	fmt.Println("✅ SOLUCIÓN 4: Funciones como First-Class Citizens")
	fmt.Println()

	// Operaciones básicas
	sumarOp := func(a, b int) int { return a + b }
	multiplicar := func(a, b int) int { return a * b }
	esMayor5 := func(n int) bool { return n > 5 }
	duplicar := func(n int) int { return n * 2 }

	// Pruebas
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Aplicar suma 5+3: %d\n", aplicarOperacion(5, 3, sumarOp))
	fmt.Printf("Aplicar multiplicación 4*7: %d\n", aplicarOperacion(4, 7, multiplicar))

	fmt.Printf("Números originales: %v\n", numeros)
	fmt.Printf("Filtrar > 5: %v\n", filtrar(numeros, esMayor5))
	fmt.Printf("Duplicar primeros 5: %v\n", mapear(numeros[:5], duplicar))
	fmt.Printf("Suma total primeros 5: %d\n", reducir(numeros[:5], 0, sumarOp))
	fmt.Printf("Producto primeros 4: %d\n", reducir(numeros[:4], 1, multiplicar))

	// Map de operaciones
	operaciones := map[string]OperacionBinaria{
		"suma":    sumarOp,
		"resta":   func(a, b int) int { return a - b },
		"mult":    multiplicar,
		"division": func(a, b int) int { return a / b },
	}

	for nombre, op := range operaciones {
		if nombre != "division" {
			fmt.Printf("%s de 12 y 4: %d\n", nombre, op(12, 4))
		} else {
			fmt.Printf("%s de 12 y 4: %d\n", nombre, op(12, 4))
		}
	}
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 5: CLOSURES
// ============================================================================

// 5.1 Función que crea un contador
func crearContador(inicial int) func() int {
	contador := inicial
	return func() int {
		contador++
		return contador
	}
}

// 5.2 Función que crea un acumulador
func crearAcumulador() func(int) int {
	total := 0
	return func(valor int) int {
		total += valor
		return total
	}
}

// 5.3 Función que crea un multiplicador
func crearMultiplicador(factor int) func(int) int {
	return func(numero int) int {
		return numero * factor
	}
}

// 5.4 Función que crea un validador de rango
func crearValidadorRango(min, max int) func(int) bool {
	return func(valor int) bool {
		return valor >= min && valor <= max
	}
}

// 5.5 Función que crea un cache simple
func crearCache() (func(string, string), func(string) (string, bool)) {
	cache := make(map[string]string)

	setter := func(clave, valor string) {
		cache[clave] = valor
	}

	getter := func(clave string) (string, bool) {
		valor, existe := cache[clave]
		return valor, existe
	}

	return setter, getter
}

func solucion5() {
	fmt.Println("✅ SOLUCIÓN 5: Closures")
	fmt.Println()

	// Pruebas
	contador := crearContador(10)
	fmt.Printf("Contador: %d, %d, %d\n", contador(), contador(), contador())

	acum := crearAcumulador()
	fmt.Printf("Acumulador: %d, %d, %d\n", acum(5), acum(3), acum(2))

	porTres := crearMultiplicador(3)
	porCinco := crearMultiplicador(5)
	fmt.Printf("Multiplicar por 3: %d, %d\n", porTres(4), porTres(7))
	fmt.Printf("Multiplicar por 5: %d, %d\n", porCinco(4), porCinco(7))

	validarEdad := crearValidadorRango(0, 120)
	validarNota := crearValidadorRango(0, 10)
	fmt.Printf("Edad 25 válida: %t, Edad 150 válida: %t\n", validarEdad(25), validarEdad(150))
	fmt.Printf("Nota 8 válida: %t, Nota 15 válida: %t\n", validarNota(8), validarNota(15))

	set, get := crearCache()
	set("usuario", "Juan")
	set("email", "juan@ejemplo.com")

	if valor, existe := get("usuario"); existe {
		fmt.Printf("Cache usuario: %s\n", valor)
	}
	if valor, existe := get("telefono"); existe {
		fmt.Printf("Cache telefono: %s\n", valor)
	} else {
		fmt.Println("telefono no encontrado en cache")
	}
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 6: DEFER, PANIC Y RECOVER
// ============================================================================

// 6.1 Función que usa defer para limpieza
func procesarArchivo(nombre string) error {
	fmt.Printf("Abriendo archivo: %s\n", nombre)

	// Simular apertura de archivo
	defer fmt.Printf("Cerrando archivo: %s\n", nombre)

	// Múltiples defers - se ejecutan en orden LIFO
	defer fmt.Println("Limpiando recursos temporales")
	defer fmt.Println("Guardando estado")

	// Simular procesamiento
	fmt.Printf("Procesando contenido de: %s\n", nombre)

	if nombre == "error.txt" {
		return errors.New("archivo corrupto")
	}

	fmt.Printf("Procesamiento exitoso de: %s\n", nombre)
	return nil
}

// 6.2 Función que maneja panic con recover
func ejecutarSeguro(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic capturado: %v", r)
		}
	}()

	fn()
	return nil
}

// 6.3 Función que puede entrar en panic
func dividirConPanic(a, b int) int {
	if b == 0 {
		panic(fmt.Sprintf("división por cero: %d / %d", a, b))
	}
	return a / b
}

// 6.4 Función con múltiples defers
func funcionConMultiplesDefers() {
	fmt.Println("Inicio de función")

	defer fmt.Println("Defer 1 - Último en ejecutar")
	defer fmt.Println("Defer 2 - Segundo en ejecutar")
	defer fmt.Println("Defer 3 - Primero en ejecutar")

	// defer con closure que captura variables
	mensaje := "Hola"
	defer func() {
		fmt.Printf("Defer con closure: %s\n", mensaje)
	}()

	mensaje = "Adiós" // Esta modificación se verá en el defer

	fmt.Println("Medio de función")
}

// 6.5 Función que cronometra ejecución con defer
func cronometrar(nombre string, fn func()) {
	inicio := time.Now()
	defer func() {
		duracion := time.Since(inicio)
		fmt.Printf("%s tomó %v\n", nombre, duracion)
	}()

	fmt.Printf("Iniciando %s...\n", nombre)
	fn()
	fmt.Printf("Completando %s\n", nombre)
}

func solucion6() {
	fmt.Println("✅ SOLUCIÓN 6: Defer, Panic y Recover")
	fmt.Println()

	// Pruebas
	fmt.Println("--- Procesando archivos ---")
	procesarArchivo("datos.txt")
	fmt.Println()

	err := procesarArchivo("error.txt")
	if err != nil {
		fmt.Printf("Error procesando archivo: %v\n", err)
	}
	fmt.Println()

	fmt.Println("--- Ejecución segura ---")
	err = ejecutarSeguro(func() {
		dividirConPanic(10, 2)
		fmt.Println("División exitosa")
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = ejecutarSeguro(func() {
		dividirConPanic(10, 0)
	})
	if err != nil {
		fmt.Printf("Error capturado: %v\n", err)
	}

	fmt.Println("\n--- Múltiples defers ---")
	funcionConMultiplesDefers()

	fmt.Println("\n--- Cronometrando ---")
	cronometrar("Operación de prueba", func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Trabajando...")
	})
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 7: RECURSIÓN
// ============================================================================

// 7.1 Fibonacci recursivo
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 7.2 Potencia recursiva
func potencia(base, exponente int) int {
	if exponente == 0 {
		return 1
	}
	if exponente == 1 {
		return base
	}

	// Optimización: dividir el exponente
	if exponente%2 == 0 {
		mitad := potencia(base, exponente/2)
		return mitad * mitad
	}
	return base * potencia(base, exponente-1)
}

// 7.3 Suma de dígitos
func sumaDigitos(n int) int {
	if n < 10 {
		return n
	}
	return (n % 10) + sumaDigitos(n/10)
}

// 7.4 Inversión de string
func invertirString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return invertirString(s[1:]) + string(s[0])
}

// 7.5 Máximo común divisor
func mcd(a, b int) int {
	if b == 0 {
		return a
	}
	return mcd(b, a%b)
}

// 7.6 Búsqueda binaria recursiva
func busquedaBinaria(arr []int, objetivo, inicio, fin int) int {
	if inicio > fin {
		return -1 // No encontrado
	}

	medio := inicio + (fin-inicio)/2

	if arr[medio] == objetivo {
		return medio
	}

	if arr[medio] > objetivo {
		return busquedaBinaria(arr, objetivo, inicio, medio-1)
	}

	return busquedaBinaria(arr, objetivo, medio+1, fin)
}

func solucion7() {
	fmt.Println("✅ SOLUCIÓN 7: Recursión")
	fmt.Println()

	// Pruebas
	fmt.Printf("Fibonacci(8): %d\n", fibonacci(8))
	fmt.Printf("Fibonacci(10): %d\n", fibonacci(10))

	fmt.Printf("2^5: %d\n", potencia(2, 5))
	fmt.Printf("3^4: %d\n", potencia(3, 4))
	fmt.Printf("5^0: %d\n", potencia(5, 0))

	fmt.Printf("Suma dígitos 123: %d\n", sumaDigitos(123))
	fmt.Printf("Suma dígitos 9876: %d\n", sumaDigitos(9876))

	fmt.Printf("Invertir 'hello': %s\n", invertirString("hello"))
	fmt.Printf("Invertir 'recursion': %s\n", invertirString("recursion"))

	fmt.Printf("MCD(48, 18): %d\n", mcd(48, 18))
	fmt.Printf("MCD(100, 25): %d\n", mcd(100, 25))

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Buscar 7: índice %d\n", busquedaBinaria(arr, 7, 0, len(arr)-1))
	fmt.Printf("Buscar 15: índice %d\n", busquedaBinaria(arr, 15, 0, len(arr)-1))
	fmt.Printf("Buscar 8: índice %d\n", busquedaBinaria(arr, 8, 0, len(arr)-1))
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 8: CALCULADORA AVANZADA
// ============================================================================

type OperacionFunc func(float64, float64) float64

type Calculadora struct {
	operaciones map[string]OperacionFunc
	historial   []string
}

func NewCalculadora() *Calculadora {
	calc := &Calculadora{
		operaciones: make(map[string]OperacionFunc),
		historial:   make([]string, 0),
	}

	// Operaciones básicas
	calc.operaciones["+"] = func(a, b float64) float64 { return a + b }
	calc.operaciones["-"] = func(a, b float64) float64 { return a - b }
	calc.operaciones["*"] = func(a, b float64) float64 { return a * b }
	calc.operaciones["/"] = func(a, b float64) float64 {
		if b == 0 {
			panic("división por cero")
		}
		return a / b
	}

	return calc
}

func (c *Calculadora) AgregarOperacion(simbolo string, op OperacionFunc) {
	c.operaciones[simbolo] = op
}

func (c *Calculadora) Calcular(operacion string, a, b float64) (resultado float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error en cálculo: %v", r)
		}
	}()

	if op, existe := c.operaciones[operacion]; existe {
		resultado = op(a, b)
		// Agregar al historial
		operacionStr := fmt.Sprintf("%.2f %s %.2f = %.2f", a, operacion, b, resultado)
		c.historial = append(c.historial, operacionStr)
		return resultado, nil
	}

	return 0, fmt.Errorf("operación '%s' no encontrada", operacion)
}

func (c *Calculadora) MostrarHistorial() {
	fmt.Println("=== Historial de Operaciones ===")
	if len(c.historial) == 0 {
		fmt.Println("Sin operaciones")
		return
	}

	for i, operacion := range c.historial {
		fmt.Printf("%d. %s\n", i+1, operacion)
	}
}

func (c *Calculadora) LimpiarHistorial() {
	c.historial = c.historial[:0]
}

func solucion8() {
	fmt.Println("✅ SOLUCIÓN 8: Calculadora Avanzada")
	fmt.Println()

	calc := NewCalculadora()

	// Agregar operaciones avanzadas
	calc.AgregarOperacion("^", func(a, b float64) float64 {
		return math.Pow(a, b)
	})
	calc.AgregarOperacion("%", func(a, b float64) float64 {
		return math.Mod(a, b)
	})

	// Pruebas
	operaciones := []struct {
		op   string
		a, b float64
	}{
		{"+", 10, 5},
		{"-", 10, 3},
		{"*", 4, 7},
		{"/", 15, 3},
		{"^", 2, 3},
		{"%", 17, 5},
		{"/", 10, 0}, // Error
	}

	for _, test := range operaciones {
		if resultado, err := calc.Calcular(test.op, test.a, test.b); err == nil {
			fmt.Printf("%.1f %s %.1f = %.2f\n", test.a, test.op, test.b, resultado)
		} else {
			fmt.Printf("%.1f %s %.1f = Error: %v\n", test.a, test.op, test.b, err)
		}
	}

	fmt.Println()
	calc.MostrarHistorial()
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 9: SISTEMA DE VALIDACIÓN
// ============================================================================

type ValidadorFunc func(interface{}) error

type Validador struct {
	reglas []ValidadorFunc
}

func NewValidador() *Validador {
	return &Validador{
		reglas: make([]ValidadorFunc, 0),
	}
}

func (v *Validador) Agregar(regla ValidadorFunc) *Validador {
	v.reglas = append(v.reglas, regla)
	return v
}

func (v *Validador) Validar(valor interface{}) error {
	for _, regla := range v.reglas {
		if err := regla(valor); err != nil {
			return err
		}
	}
	return nil
}

// Funciones de validación comunes
func ValidarRequerido() ValidadorFunc {
	return func(valor interface{}) error {
		if valor == nil {
			return errors.New("campo requerido")
		}

		switch v := valor.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return errors.New("campo requerido")
			}
		case int:
			// Los números siempre son válidos para requerido
		default:
			// Para otros tipos, verificar si es el valor cero
			if valor == nil {
				return errors.New("campo requerido")
			}
		}
		return nil
	}
}

func ValidarLongitudMinima(min int) ValidadorFunc {
	return func(valor interface{}) error {
		if str, ok := valor.(string); ok {
			if len(str) < min {
				return fmt.Errorf("longitud mínima requerida: %d", min)
			}
		}
		return nil
	}
}

func ValidarEmail() ValidadorFunc {
	return func(valor interface{}) error {
		if str, ok := valor.(string); ok {
			if !strings.Contains(str, "@") || !strings.Contains(str, ".") {
				return errors.New("formato de email inválido")
			}

			partes := strings.Split(str, "@")
			if len(partes) != 2 || len(partes[0]) == 0 || len(partes[1]) == 0 {
				return errors.New("formato de email inválido")
			}
		}
		return nil
	}
}

func ValidarRangoNumerico(min, max float64) ValidadorFunc {
	return func(valor interface{}) error {
		var num float64
		var ok bool

		switch v := valor.(type) {
		case int:
			num = float64(v)
			ok = true
		case float64:
			num = v
			ok = true
		case string:
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				num = parsed
				ok = true
			}
		}

		if !ok {
			return errors.New("valor no es numérico")
		}

		if num < min || num > max {
			return fmt.Errorf("valor debe estar entre %.2f y %.2f", min, max)
		}

		return nil
	}
}

func solucion9() {
	fmt.Println("✅ SOLUCIÓN 9: Sistema de Validación")
	fmt.Println()

	// Validador de email
	validadorEmail := NewValidador().
		Agregar(ValidarRequerido()).
		Agregar(ValidarLongitudMinima(5)).
		Agregar(ValidarEmail())

	emails := []string{"usuario@ejemplo.com", "", "invalido", "test@test.co", "a@b", "correcto@dominio.com.mx"}

	fmt.Println("=== Validación de Emails ===")
	for _, email := range emails {
		if err := validadorEmail.Validar(email); err != nil {
			fmt.Printf("❌ '%s': %v\n", email, err)
		} else {
			fmt.Printf("✅ '%s': válido\n", email)
		}
	}

	// Validador de edad
	validadorEdad := NewValidador().
		Agregar(ValidarRequerido()).
		Agregar(ValidarRangoNumerico(0, 120))

	fmt.Println("\n=== Validación de Edades ===")
	edades := []interface{}{25, -5, 150, "30", "abc", 0, 120}

	for _, edad := range edades {
		if err := validadorEdad.Validar(edad); err != nil {
			fmt.Printf("❌ %v: %v\n", edad, err)
		} else {
			fmt.Printf("✅ %v: válido\n", edad)
		}
	}
	fmt.Println()
}

// ============================================================================
// SOLUCIÓN 10: PIPELINE DE PROCESAMIENTO
// ============================================================================

type ProcesadorFunc func(interface{}) interface{}

type Pipeline struct {
	pasos []ProcesadorFunc
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		pasos: make([]ProcesadorFunc, 0),
	}
}

func (p *Pipeline) Agregar(procesador ProcesadorFunc) *Pipeline {
	p.pasos = append(p.pasos, procesador)
	return p
}

func (p *Pipeline) Procesar(entrada interface{}) interface{} {
	resultado := entrada
	for _, paso := range p.pasos {
		resultado = paso(resultado)
	}
	return resultado
}

func (p *Pipeline) ProcesarSlice(entradas []interface{}) []interface{} {
	resultados := make([]interface{}, len(entradas))
	for i, entrada := range entradas {
		resultados[i] = p.Procesar(entrada)
	}
	return resultados
}

// Procesadores comunes
func ProcesarTextoMayusculas() ProcesadorFunc {
	return func(entrada interface{}) interface{} {
		if str, ok := entrada.(string); ok {
			return strings.ToUpper(str)
		}
		return entrada
	}
}

func ProcesarTextoLimpiar() ProcesadorFunc {
	return func(entrada interface{}) interface{} {
		if str, ok := entrada.(string); ok {
			return strings.TrimSpace(str)
		}
		return entrada
	}
}

func ProcesarNumeroMultiplicar(factor float64) ProcesadorFunc {
	return func(entrada interface{}) interface{} {
		switch v := entrada.(type) {
		case float64:
			return v * factor
		case int:
			return float64(v) * factor
		}
		return entrada
	}
}

func ProcesarNumeroRaizCuadrada() ProcesadorFunc {
	return func(entrada interface{}) interface{} {
		switch v := entrada.(type) {
		case float64:
			if v >= 0 {
				return math.Sqrt(v)
			}
		case int:
			if v >= 0 {
				return math.Sqrt(float64(v))
			}
		}
		return entrada
	}
}

func ProcesarTextoAgregarPrefijo(prefijo string) ProcesadorFunc {
	return func(entrada interface{}) interface{} {
		if str, ok := entrada.(string); ok {
			return prefijo + str
		}
		return entrada
	}
}

func solucion10() {
	fmt.Println("✅ SOLUCIÓN 10: Pipeline de Procesamiento")
	fmt.Println()

	// Pipeline para texto
	pipelineTexto := NewPipeline().
		Agregar(ProcesarTextoLimpiar()).
		Agregar(ProcesarTextoMayusculas()).
		Agregar(ProcesarTextoAgregarPrefijo(">>> "))

	textos := []interface{}{"  hola mundo  ", " go es genial ", "  programación  "}
	fmt.Println("=== Pipeline de Texto ===")
	fmt.Printf("Originales: %v\n", textos)

	resultadosTexto := pipelineTexto.ProcesarSlice(textos)
	fmt.Printf("Procesados: %v\n", resultadosTexto)

	// Pipeline para números
	pipelineNumeros := NewPipeline().
		Agregar(ProcesarNumeroMultiplicar(2)).
		Agregar(ProcesarNumeroRaizCuadrada())

	numeros := []interface{}{4.0, 9.0, 16.0, 25.0, 36.0}
	fmt.Println("\n=== Pipeline de Números ===")
	fmt.Printf("Originales: %v\n", numeros)

	resultadosNum := pipelineNumeros.ProcesarSlice(numeros)
	fmt.Printf("Procesados (x2 -> √): %v\n", resultadosNum)

	// Pipeline mixto
	fmt.Println("\n=== Pipeline Individual ===")
	procesamientoComplejo := NewPipeline().
		Agregar(func(entrada interface{}) interface{} {
			fmt.Printf("Paso 1 - Entrada: %v\n", entrada)
			return entrada
		}).
		Agregar(ProcesarTextoLimpiar()).
		Agregar(func(entrada interface{}) interface{} {
			fmt.Printf("Paso 2 - Después de limpiar: %v\n", entrada)
			return entrada
		}).
		Agregar(ProcesarTextoMayusculas()).
		Agregar(func(entrada interface{}) interface{} {
			fmt.Printf("Paso 3 - Final: %v\n", entrada)
			return entrada
		})

	resultado := procesamientoComplejo.Procesar("  ejemplo de pipeline  ")
	fmt.Printf("Resultado final: %v\n", resultado)
	fmt.Println()
}

// ============================================================================
// DEMO COMPLETO
// ============================================================================

func demoCompleto() {
	fmt.Println("🎯 DEMO COMPLETO DE FUNCIONES EN GO")
	fmt.Println("==================================")

	// Mostrar todos los conceptos juntos
	fmt.Println("\n--- Funciones Básicas ---")
	fmt.Printf("Área de círculo: %.2f\n", calcularAreaCirculo(3))

	fmt.Println("\n--- Múltiples Retornos ---")
	if resultado, err := dividirSeguro(15, 4); err == nil {
		fmt.Printf("División: %.2f\n", resultado)
	}

	fmt.Println("\n--- Funciones Variádicas ---")
	fmt.Printf("Suma: %d\n", sumar(1, 2, 3, 4, 5))

	fmt.Println("\n--- Closures ---")
	contador := crearContador(0)
	fmt.Printf("Contador: %d, %d\n", contador(), contador())

	fmt.Println("\n--- Defer ---")
	defer fmt.Println("Este mensaje se imprime al final")

	fmt.Println("\n--- Recursión ---")
	fmt.Printf("Factorial de 5: %d\n", factorial(5))

	fmt.Println("Función principal terminando...")
}

// ============================================================================
// EJEMPLOS ADICIONALES Y PATRONES AVANZADOS
// ============================================================================

// Patrón Decorator con funciones
type HandlerFunc func(string) string

func decorarConLog(handler HandlerFunc) HandlerFunc {
	return func(entrada string) string {
		fmt.Printf("[LOG] Procesando: %s\n", entrada)
		resultado := handler(entrada)
		fmt.Printf("[LOG] Resultado: %s\n", resultado)
		return resultado
	}
}

func decorarConValidacion(handler HandlerFunc) HandlerFunc {
	return func(entrada string) string {
		if entrada == "" {
			return "ERROR: entrada vacía"
		}
		return handler(entrada)
	}
}

// Patrón Strategy con funciones
type EstrategiaOrdenamiento func([]string) []string

func ordenarAlfabetico(palabras []string) []string {
	// Implementación simple - en la práctica usarías sort.Strings
	resultado := make([]string, len(palabras))
	copy(resultado, palabras)
	return resultado
}

// Function composition
func componer(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func demoPatronesAvanzados() {
	fmt.Println("🎨 PATRONES AVANZADOS CON FUNCIONES")
	fmt.Println("=================================")

	// Decorator
	fmt.Println("\n--- Patrón Decorator ---")
	handlerBase := func(s string) string { return strings.ToUpper(s) }

	handlerCompleto := decorarConLog(decorarConValidacion(handlerBase))
	resultado := handlerCompleto("hola mundo")
	fmt.Printf("Resultado final: %s\n", resultado)

	// Function composition
	fmt.Println("\n--- Composición de Funciones ---")
	duplicar := func(x int) int { return x * 2 }
	sumarUno := func(x int) int { return x + 1 }

	duplicarYSumar := componer(sumarUno, duplicar)
	fmt.Printf("(5 * 2) + 1 = %d\n", duplicarYSumar(5))
}

/*
RESUMEN DE CONCEPTOS IMPLEMENTADOS:

✅ Funciones básicas con parámetros y retornos
✅ Múltiples valores de retorno e idiomas Go
✅ Funciones variádicas con ...tipo
✅ Funciones como first-class citizens
✅ Closures y captura de variables
✅ Defer, panic y recover para control de flujo
✅ Recursión con casos base y optimizaciones
✅ Patrones de diseño (Strategy, Decorator, Observer)
✅ Sistemas prácticos (Calculadora, Validación, Pipeline)
✅ Manejo de errores idiomático
✅ Programación funcional en Go

PRÓXIMOS PASOS:
1. Ejecuta cada función para ver los resultados
2. Modifica los ejemplos para experimentar
3. Implementa tus propias variaciones
4. Combina conceptos para proyectos más complejos

¡Dominar las funciones te dará las herramientas para escribir código Go elegante y mantenible! 🚀
*/
