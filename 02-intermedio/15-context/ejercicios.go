// 🎯 Ejercicios de Context Package - Lección 15
// ===============================================

package main

import (
	"context"
	"fmt"
	"time"
)

// Tipos para keys de context
type contextKey string

const (
	UserIDKey    contextKey = "userID"
	RequestIDKey contextKey = "requestID"
	SessionIDKey contextKey = "sessionID"
	TenantIDKey  contextKey = "tenantID"
)

// ==========================================
// 📝 EJERCICIO 1: Context Básico
// ==========================================
// Objetivo: Crear y usar un context básico con cancelación

func ejercicio1() {
	fmt.Println("📝 Ejercicio 1: Context Básico")
	fmt.Println("==============================")

	// TODO: Crea un context con cancelación
	// var ctx context.Context
	// var cancel context.CancelFunc

	// TODO: Lanza una goroutine que imprima números cada 200ms
	// La goroutine debe detenerse cuando reciba la cancelación
	// go func() {
	//     // Tu código aquí
	// }()

	// TODO: Espera 1.5 segundos y luego cancela el context

	// TODO: Da tiempo para que la goroutine se detenga

	fmt.Println("✅ Ejercicio 1 completado\n")
}

// ==========================================
// 📝 EJERCICIO 2: Context con Timeout
// ==========================================
// Objetivo: Implementar una operación con timeout

func operacionLenta2(ctx context.Context, duracion time.Duration) error {
	// TODO: Implementa una operación que:
	// - Simule trabajo durante 'duracion'
	// - Se cancele si el context expira
	// - Retorne el error apropiado

	return nil
}

func ejercicio2() {
	fmt.Println("📝 Ejercicio 2: Context con Timeout")
	fmt.Println("===================================")

	// TODO: Crea un context con timeout de 1 segundo

	// TODO: Prueba operacionLenta2 con diferentes duraciones:
	// - 500ms (debería completarse)
	// - 1.5s (debería timeout)

	fmt.Println("✅ Ejercicio 2 completado\n")
}

// ==========================================
// 📝 EJERCICIO 3: Context Values
// ==========================================
// Objetivo: Propagar valores a través del context

func autenticar3(ctx context.Context, token string) context.Context {
	// TODO: Simula autenticación y agrega userID al context
	// El userID debe ser "user-" + los últimos 4 caracteres del token

	return ctx
}

func autorizar3(ctx context.Context, recurso string) bool {
	// TODO: Verifica que existe un userID en el context
	// Retorna true si el userID no es "user-deny"

	return false
}

func procesarRequest3(ctx context.Context, recurso string) {
	// TODO: Implementa la lógica para:
	// - Verificar autorización usando autorizar3
	// - Si está autorizado, simular procesamiento de 300ms
	// - Imprimir mensajes informativos con el userID
}

func ejercicio3() {
	fmt.Println("📝 Ejercicio 3: Context Values")
	fmt.Println("==============================")

	ctx := context.Background()

	// TODO: Prueba con diferentes tokens:
	tokens := []string{"abc-1234", "xyz-deny", "def-5678"}

	for _, token := range tokens {
		// TODO: Autentica y procesa cada request
		fmt.Printf("🔑 Procesando token: %s\n", token)
		// Tu código aquí
	}

	fmt.Println("✅ Ejercicio 3 completado\n")
}

// ==========================================
// 📝 EJERCICIO 4: Context Deadline
// ==========================================
// Objetivo: Usar deadline específico para una operación

func procesamientoPorLotes4(ctx context.Context, elementos []string) []string {
	// TODO: Procesa elementos de la lista uno por uno
	// - Cada elemento toma 100ms en procesarse
	// - Si el context expira, detén el procesamiento
	// - Retorna los elementos procesados hasta el momento

	var procesados []string
	return procesados
}

func ejercicio4() {
	fmt.Println("📝 Ejercicio 4: Context Deadline")
	fmt.Println("================================")

	elementos := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	// TODO: Crea un context con deadline de 450ms desde ahora

	// TODO: Procesa los elementos y muestra cuántos se completaron

	fmt.Println("✅ Ejercicio 4 completado\n")
}

// ==========================================
// 📝 EJERCICIO 5: Multiple Contexts
// ==========================================
// Objetivo: Manejar múltiples contexts simultáneamente

func worker5(ctx context.Context, id int, resultChan chan<- string) {
	// TODO: Implementa un worker que:
	// - Ejecute iteraciones cada 100ms
	// - Se detenga cuando el context se cancele
	// - Envíe su resultado final al canal antes de terminar
}

func ejercicio5() {
	fmt.Println("📝 Ejercicio 5: Multiple Contexts")
	fmt.Println("=================================")

	// TODO: Crea un context principal con timeout de 1 segundo

	// TODO: Crea un context cancelable manualmente

	// TODO: Lanza 3 workers con el context principal

	// TODO: Después de 300ms, cancela manualmente uno de los contexts

	// TODO: Recopila todos los resultados

	fmt.Println("✅ Ejercicio 5 completado\n")
}

// ==========================================
// 📝 EJERCICIO 6: Context Middleware
// ==========================================
// Objetivo: Implementar un patrón middleware con context

type Middleware6 func(context.Context, func(context.Context)) context.Context

func loggingMiddleware6(ctx context.Context, next func(context.Context)) context.Context {
	// TODO: Implementa middleware que:
	// - Registre el inicio de la operación
	// - Ejecute la función siguiente
	// - Registre el tiempo total transcurrido

	return ctx
}

func authMiddleware6(ctx context.Context, next func(context.Context)) context.Context {
	// TODO: Implementa middleware que:
	// - Verifique si hay un userID en el context
	// - Si no hay userID, no ejecute 'next'
	// - Si hay userID, ejecute 'next'

	return ctx
}

func handlerFinal6(ctx context.Context) {
	// TODO: Implementa el handler final que:
	// - Simule trabajo de 200ms
	// - Pueda ser cancelado por el context
	// - Imprima el userID si existe
}

func ejercicio6() {
	fmt.Println("📝 Ejercicio 6: Context Middleware")
	fmt.Println("==================================")

	// TODO: Crea diferentes contexts de prueba:
	// - Uno con userID válido
	// - Uno sin userID
	// - Uno que se cancele después de 100ms

	// TODO: Aplica los middlewares y ejecuta el handler

	fmt.Println("✅ Ejercicio 6 completado\n")
}

// ==========================================
// 📝 EJERCICIO 7: Context Pipeline
// ==========================================
// Objetivo: Crear un pipeline de procesamiento con context

func etapa1_7(ctx context.Context, input string) (string, error) {
	// TODO: Primera etapa del pipeline (100ms)
	// Transforma input agregando prefijo "stage1-"
	return "", nil
}

func etapa2_7(ctx context.Context, input string) (string, error) {
	// TODO: Segunda etapa del pipeline (150ms)
	// Transforma input agregando prefijo "stage2-"
	return "", nil
}

func etapa3_7(ctx context.Context, input string) (string, error) {
	// TODO: Tercera etapa del pipeline (200ms)
	// Transforma input agregando prefijo "stage3-"
	return "", nil
}

func pipeline7(ctx context.Context, input string) (string, error) {
	// TODO: Ejecuta las 3 etapas en secuencia
	// Si cualquier etapa falla o se cancela, detén el pipeline

	return "", nil
}

func ejercicio7() {
	fmt.Println("📝 Ejercicio 7: Context Pipeline")
	fmt.Println("================================")

	inputs := []string{"data1", "data2", "data3"}

	// TODO: Procesa cada input con diferentes timeouts:
	// - input1: timeout 600ms (debería completarse)
	// - input2: timeout 300ms (debería fallar en etapa 2 o 3)
	// - input3: sin timeout pero cancela manualmente después de 250ms

	fmt.Println("✅ Ejercicio 7 completado\n")
}

// ==========================================
// 📝 EJERCICIO 8: Context Pool Worker
// ==========================================
// Objetivo: Implementar un pool de workers con context

type Job8 struct {
	ID   int
	Data string
}

type Result8 struct {
	JobID  int
	Result string
	Error  error
}

type WorkerPool8 struct {
	// TODO: Define los campos necesarios:
	// - Canal de jobs
	// - Canal de resultados
	// - Context para cancelación
	// - Número de workers
}

func NewWorkerPool8(numWorkers int) *WorkerPool8 {
	// TODO: Crea e inicializa un nuevo WorkerPool
	return nil
}

func (wp *WorkerPool8) Start() {
	// TODO: Inicia todos los workers
}

func (wp *WorkerPool8) worker(id int) {
	// TODO: Implementa la lógica del worker:
	// - Escucha jobs del canal
	// - Procesa cada job (simula 200ms de trabajo)
	// - Envía resultados al canal de resultados
	// - Se detiene cuando se cancela el context
}

func (wp *WorkerPool8) Submit(job Job8) {
	// TODO: Envía un job al pool
}

func (wp *WorkerPool8) Stop() {
	// TODO: Detiene el pool de workers elegantemente
}

func ejercicio8() {
	fmt.Println("📝 Ejercicio 8: Context Pool Worker")
	fmt.Println("===================================")

	// TODO: Crea un worker pool con 3 workers

	// TODO: Envía 5 jobs al pool

	// TODO: Recopila resultados por 1 segundo

	// TODO: Detén el pool elegantemente

	fmt.Println("✅ Ejercicio 8 completado\n")
}

// ==========================================
// 📝 EJERCICIO 9: Context con Select
// ==========================================
// Objetivo: Usar select con múltiples contexts y operaciones

func servicioExterno9(ctx context.Context, servicio string) (string, error) {
	// TODO: Simula llamadas a diferentes servicios externos:
	// - "rapido": 100ms
	// - "medio": 300ms
	// - "lento": 800ms
	// Debe ser cancelable por el context

	return "", nil
}

func ejercicio9() {
	fmt.Println("📝 Ejercicio 9: Context con Select")
	fmt.Println("==================================")

	// TODO: Implementa un patrón que:
	// - Haga llamadas concurrentes a 3 servicios
	// - Use el primer resultado que llegue
	// - Cancele las operaciones restantes
	// - Tenga un timeout global de 500ms

	servicios := []string{"rapido", "medio", "lento"}

	fmt.Println("✅ Ejercicio 9 completado\n")
}

// ==========================================
// 📝 EJERCICIO 10: Context Composition
// ==========================================
// Objetivo: Combinar múltiples contexts en uno compuesto

type CompositeContext10 struct {
	// TODO: Define la estructura para un context compuesto
	// que combine múltiples contexts
}

func NewCompositeContext10(contexts ...context.Context) *CompositeContext10 {
	// TODO: Crea un context que se cancele cuando
	// cualquiera de los contexts padre se cancele
	return nil
}

func (cc *CompositeContext10) Done() <-chan struct{} {
	// TODO: Implementa el método Done
	return nil
}

func (cc *CompositeContext10) Err() error {
	// TODO: Implementa el método Err
	return nil
}

func (cc *CompositeContext10) Deadline() (time.Time, bool) {
	// TODO: Retorna el deadline más próximo
	return time.Time{}, false
}

func (cc *CompositeContext10) Value(key interface{}) interface{} {
	// TODO: Busca el valor en todos los contexts padre
	return nil
}

func ejercicio10() {
	fmt.Println("📝 Ejercicio 10: Context Composition")
	fmt.Println("====================================")

	// TODO: Crea múltiples contexts:
	// - Uno con valores (userID, requestID)
	// - Uno con timeout de 2 segundos
	// - Uno cancelable manualmente

	// TODO: Combínalos en un CompositeContext

	// TODO: Demuestra que:
	// - Los valores se pueden acceder
	// - El timeout funciona
	// - La cancelación manual funciona

	fmt.Println("✅ Ejercicio 10 completado\n")
}

// ==========================================
// 🏃‍♂️ FUNCIÓN PRINCIPAL
// ==========================================

func main() {
	fmt.Println("🎯 Ejercicios de Context Package")
	fmt.Println("=================================")
	fmt.Println("💡 Completa cada ejercicio implementando el código faltante")
	fmt.Println("🔧 Usa las funciones TODO como guía")
	fmt.Println()

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

	fmt.Println("🎉 ¡Todos los ejercicios completados!")
	fmt.Println("👀 Revisa las soluciones en soluciones.go")
}
