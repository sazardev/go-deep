// 🚨 Ejercicios: Error Handling Avanzado
// Lección 16: Robustez y Resilencia en Go

package main

import (
	"fmt"
)

// ========================================
// Ejercicio 1: Error Básico Personalizado
// ========================================

// TODO: Crea un tipo de error personalizado llamado ValidationError
// que contenga los campos: Field, Value, Message
// Implementa el método Error() que retorne un mensaje descriptivo

type ValidationError struct {
	// TODO: Agregar campos necesarios
}

// TODO: Implementar método Error()

// TODO: Función que valide una edad y retorne ValidationError si es inválida
func validateAge(age int) error {
	// TODO: Validar que la edad esté entre 0 y 150
	// Retornar ValidationError apropiado si no es válida
	return nil
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: Error Personalizado ===")

	// TODO: Probar con edad válida e inválida
	// Mostrar los mensajes de error generados
}

// ==========================================
// Ejercicio 2: Error Wrapping con Contexto
// ==========================================

// TODO: Función que simule una operación de base de datos
func queryDatabase(userID string) (string, error) {
	// TODO: Simular diferentes tipos de errores:
	// - userID vacío: error de validación
	// - userID "999": usuario no encontrado
	// - userID "error": error de conexión
	// - userID válido: retornar datos del usuario
	// Usar error wrapping para agregar contexto
	return "", nil
}

// TODO: Función de servicio que llame a queryDatabase y agregue más contexto
func getUserService(userID string) (string, error) {
	// TODO: Llamar queryDatabase y agregar contexto de servicio
	return "", nil
}

func ejercicio2() {
	fmt.Println("\n=== Ejercicio 2: Error Wrapping ===")

	// TODO: Probar con diferentes userIDs y mostrar la cadena de errores
	// Usar errors.Unwrap para mostrar errores originales
}

// ========================================
// Ejercicio 3: Error Accumulator Pattern
// ========================================

// TODO: Implementar un acumulador de errores para validación
type ErrorAccumulator struct {
	// TODO: Campo para almacenar múltiples errores
}

// TODO: Método para agregar un error al acumulador
func (ea *ErrorAccumulator) Add(err error) {
	// TODO: Implementar
}

// TODO: Método para verificar si hay errores
func (ea *ErrorAccumulator) HasErrors() bool {
	// TODO: Implementar
	return false
}

// TODO: Método Error() para satisfacer la interface error
func (ea *ErrorAccumulator) Error() string {
	// TODO: Implementar - combinar todos los errores en un mensaje
	return ""
}

// TODO: Función que valide múltiples campos de un usuario
func validateUser(name, email string, age int) error {
	// TODO: Crear ErrorAccumulator
	// TODO: Validar nombre (mínimo 2 caracteres)
	// TODO: Validar email (debe contener @)
	// TODO: Validar edad (usar validateAge del ejercicio 1)
	// TODO: Retornar acumulador si tiene errores
	return nil
}

func ejercicio3() {
	fmt.Println("\n=== Ejercicio 3: Error Accumulator ===")

	// TODO: Probar validateUser con datos inválidos
	// Mostrar todos los errores de validación juntos
}

// ===================================
// Ejercicio 4: Result Type Pattern
// ===================================

// TODO: Implementar un tipo Result genérico
type Result[T any] struct {
	// TODO: Campos para valor y error
}

// TODO: Constructor para resultado exitoso
func Ok[T any](value T) Result[T] {
	// TODO: Implementar
	return Result[T]{}
}

// TODO: Constructor para resultado con error
func Err[T any](err error) Result[T] {
	// TODO: Implementar
	return Result[T]{}
}

// TODO: Método para verificar si es exitoso
func (r Result[T]) IsOk() bool {
	// TODO: Implementar
	return false
}

// TODO: Método para extraer valor y error
func (r Result[T]) Unwrap() (T, error) {
	// TODO: Implementar
	var zero T
	return zero, nil
}

// TODO: Función que use el Result type para parsear un número
func parseNumber(s string) Result[int] {
	// TODO: Usar strconv.Atoi y retornar Result apropiado
	return Result[int]{}
}

func ejercicio4() {
	fmt.Println("\n=== Ejercicio 4: Result Type Pattern ===")

	// TODO: Probar parseNumber con strings válidos e inválidos
	// Mostrar el uso del Result type
}

// =====================================
// Ejercicio 5: Circuit Breaker Simple
// =====================================

// TODO: Estados del circuit breaker
type CircuitState int

const (
// TODO: Definir estados: Closed, Open, HalfOpen
)

// TODO: Implementar CircuitBreaker simple
type CircuitBreaker struct {
	// TODO: Campos necesarios: estado, contador de fallos, umbral
}

// TODO: Constructor
func NewCircuitBreaker(threshold int) *CircuitBreaker {
	// TODO: Implementar
	return nil
}

// TODO: Método para ejecutar operación con circuit breaker
func (cb *CircuitBreaker) Execute(operation func() error) error {
	// TODO: Implementar lógica del circuit breaker
	// - Si está Open, retornar error inmediatamente
	// - Si está Closed o HalfOpen, ejecutar operación
	// - Manejar éxitos y fallos apropiadamente
	return nil
}

// TODO: Función que simule una operación que falla a veces
func unreliableOperation() error {
	// TODO: Retornar error 70% del tiempo
	return nil
}

func ejercicio5() {
	fmt.Println("\n=== Ejercicio 5: Circuit Breaker ===")

	// TODO: Crear circuit breaker con umbral de 3
	// TODO: Probar múltiples ejecuciones y mostrar comportamiento
}

// ===================================
// Ejercicio 6: Error Middleware HTTP
// ===================================

// TODO: Tipo para errores HTTP con código de estado
type HTTPError struct {
	// TODO: Campos: Code, Message, StatusCode
}

// TODO: Implementar método Error()

// TODO: Constructores para errores comunes
func NewBadRequest(message string) HTTPError {
	// TODO: Implementar
	return HTTPError{}
}

func NewNotFound(resource string) HTTPError {
	// TODO: Implementar
	return HTTPError{}
}

func NewInternalError(err error) HTTPError {
	// TODO: Implementar
	return HTTPError{}
}

// TODO: Simulación de handler HTTP que puede fallar
func simulateHTTPHandler(endpoint string) error {
	// TODO: Simular diferentes errores según el endpoint:
	// - "/invalid": BadRequest
	// - "/missing": NotFound
	// - "/error": InternalError
	// - otros: éxito (nil)
	return nil
}

func ejercicio6() {
	fmt.Println("\n=== Ejercicio 6: HTTP Error Handling ===")

	// TODO: Probar diferentes endpoints y mostrar errores HTTP apropiados
}

// =====================================
// Ejercicio 7: Error Metrics y Logging
// =====================================

// TODO: Sistema simple de métricas de errores
type ErrorMetrics struct {
	// TODO: Contadores por tipo de error
}

// TODO: Constructor
func NewErrorMetrics() *ErrorMetrics {
	// TODO: Implementar
	return nil
}

// TODO: Método para registrar un error
func (em *ErrorMetrics) RecordError(errorType string) {
	// TODO: Implementar
}

// TODO: Método para obtener estadísticas
func (em *ErrorMetrics) GetStats() map[string]int {
	// TODO: Implementar
	return nil
}

// TODO: Logger que use las métricas
type ErrorLogger struct {
	// TODO: Campo para métricas
}

// TODO: Método para loggear error con métricas
func (el *ErrorLogger) LogError(err error, service string) {
	// TODO: Determinar tipo de error y registrar en métricas
	// TODO: Loggear el error con formato estructurado
}

func ejercicio7() {
	fmt.Println("\n=== Ejercicio 7: Error Metrics ===")

	// TODO: Crear logger con métricas
	// TODO: Loggear varios errores de diferentes tipos
	// TODO: Mostrar estadísticas finales
}

// =======================================
// Ejercicio 8: Retry con Backoff Exponencial
// =======================================

// TODO: Función que reinente una operación con backoff exponencial
func retryWithBackoff(operation func() error, maxRetries int) error {
	// TODO: Implementar retry con backoff exponencial
	// TODO: Usar time.Sleep con delays crecientes: 1s, 2s, 4s, 8s...
	// TODO: Retornar el último error si todos los reintentos fallan
	return nil
}

// TODO: Operación que falla las primeras N veces
var attemptCount int

func flakeyOperation() error {
	// TODO: Fallar las primeras 3 veces, luego exitoso
	return nil
}

func ejercicio8() {
	fmt.Println("\n=== Ejercicio 8: Retry con Backoff ===")

	// TODO: Probar retryWithBackoff con flakeyOperation
	// TODO: Mostrar los intentos y delays
}

// ================================
// Ejercicio 9: Error Testing Helper
// ================================

// TODO: Helper para testing que verifique tipos de error específicos
func assertErrorType[T error](t interface{}, err error, expectedType T) bool {
	// TODO: Usar errors.As para verificar el tipo
	// TODO: Retornar true si el tipo coincide
	// En un test real, esto usaría *testing.T
	return false
}

// TODO: Helper para verificar error sentinela
func assertErrorIs(t interface{}, err error, target error) bool {
	// TODO: Usar errors.Is para verificar
	return false
}

// TODO: Función de ejemplo para testear
func processInput(input string) error {
	if input == "" {
		return NewBadRequest("input cannot be empty")
	}
	if input == "not-found" {
		return NewNotFound("resource")
	}
	return nil
}

func ejercicio9() {
	fmt.Println("\n=== Ejercicio 9: Error Testing ===")

	// TODO: Probar processInput con diferentes inputs
	// TODO: Usar los helpers para verificar tipos de error
	// TODO: Mostrar resultados de las verificaciones
}

// ===================================
// Ejercicio 10: Sistema de Error Completo
// ===================================

// TODO: Implementar un sistema completo que combine:
// - Errores personalizados
// - Error wrapping
// - Circuit breaker
// - Métricas
// - Retry logic

// TODO: Servicio que use todos los componentes
type RobustService struct {
	// TODO: Campos para circuit breaker, métricas, etc.
}

// TODO: Constructor
func NewRobustService() *RobustService {
	// TODO: Implementar
	return nil
}

// TODO: Método principal que demuestre robustez
func (rs *RobustService) ProcessRequest(requestID string) error {
	// TODO: Implementar usando:
	// - Circuit breaker para protección
	// - Retry con backoff
	// - Error wrapping con contexto
	// - Logging con métricas
	return nil
}

func ejercicio10() {
	fmt.Println("\n=== Ejercicio 10: Sistema Robusto Completo ===")

	// TODO: Crear servicio robusto
	// TODO: Procesar múltiples requests, algunos exitosos, otros fallidos
	// TODO: Mostrar métricas y comportamiento del circuit breaker
}

// ===============================
// Función principal de ejercicios
// ===============================

func main() {
	fmt.Println("🚨 EJERCICIOS: ERROR HANDLING AVANZADO")
	fmt.Println("=====================================")

	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()
	ejercicio9()
	ejercicio10()

	fmt.Println("\n🎯 ¡Completa todos los ejercicios para dominar el error handling avanzado!")
	fmt.Println("💡 Recuerda: El buen error handling es la diferencia entre código amateur y profesional")
}
