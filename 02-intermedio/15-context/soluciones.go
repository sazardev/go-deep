// 💡 Soluciones de Context Package - Lección 15
// ============================================

package main

import (
	"context"
	"fmt"
	"sync"
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
// ✅ SOLUCIÓN 1: Context Básico
// ==========================================

func solucion1() {
	fmt.Println("✅ Solución 1: Context Básico")
	fmt.Println("=============================")

	// Crear context con cancelación
	ctx, cancel := context.WithCancel(context.Background())

	// Lanzar goroutine que imprime números
	go func() {
		contador := 1
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("🔴 Goroutine cancelada después de %d números: %v\n", contador-1, ctx.Err())
				return
			case <-ticker.C:
				fmt.Printf("🔢 Número: %d\n", contador)
				contador++
			}
		}
	}()

	// Esperar 1.5 segundos y cancelar
	fmt.Println("⏱️ Esperando 1.5 segundos...")
	time.Sleep(1500 * time.Millisecond)

	fmt.Println("📤 Enviando cancelación...")
	cancel()

	// Dar tiempo para que termine
	time.Sleep(300 * time.Millisecond)
	fmt.Println("✅ Solución 1 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 2: Context con Timeout
// ==========================================

func operacionLentaSol2(ctx context.Context, duracion time.Duration) error {
	fmt.Printf("🚀 Iniciando operación que durará %v\n", duracion)

	select {
	case <-time.After(duracion):
		fmt.Printf("✅ Operación completada exitosamente\n")
		return nil
	case <-ctx.Done():
		fmt.Printf("⏰ Operación cancelada por timeout: %v\n", ctx.Err())
		return ctx.Err()
	}
}

func solucion2() {
	fmt.Println("✅ Solución 2: Context con Timeout")
	fmt.Println("==================================")

	// Context con timeout de 1 segundo
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Prueba 1: Operación que debería completarse (500ms)
	fmt.Println("🎯 Prueba 1: Operación rápida (500ms)")
	err := operacionLentaSol2(ctx, 500*time.Millisecond)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}

	fmt.Println()

	// Crear nuevo context para segunda prueba
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	// Prueba 2: Operación que debería timeout (1.5s)
	fmt.Println("🎯 Prueba 2: Operación lenta (1.5s)")
	err = operacionLentaSol2(ctx2, 1500*time.Millisecond)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}

	fmt.Println("✅ Solución 2 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 3: Context Values
// ==========================================

func autenticarSol3(ctx context.Context, token string) context.Context {
	// Simular autenticación
	userID := "user-" + token[len(token)-4:] // Últimos 4 caracteres
	fmt.Printf("🔐 Usuario autenticado: %s\n", userID)

	return context.WithValue(ctx, UserIDKey, userID)
}

func autorizarSol3(ctx context.Context, recurso string) bool {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		fmt.Printf("❌ No se encontró userID en el context\n")
		return false
	}

	if userID == "user-deny" {
		fmt.Printf("🚫 Acceso denegado para %s al recurso %s\n", userID, recurso)
		return false
	}

	fmt.Printf("✅ Acceso autorizado para %s al recurso %s\n", userID, recurso)
	return true
}

func procesarRequestSol3(ctx context.Context, recurso string) {
	userID, _ := ctx.Value(UserIDKey).(string)

	if !autorizarSol3(ctx, recurso) {
		return
	}

	fmt.Printf("🔄 %s procesando recurso %s...\n", userID, recurso)
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("✅ %s completó procesamiento de %s\n", userID, recurso)
}

func solucion3() {
	fmt.Println("✅ Solución 3: Context Values")
	fmt.Println("=============================")

	ctx := context.Background()
	tokens := []string{"abc-1234", "xyz-deny", "def-5678"}

	for i, token := range tokens {
		fmt.Printf("🔑 Procesando token: %s\n", token)

		// Autenticar y obtener context con userID
		authCtx := autenticarSol3(ctx, token)

		// Procesar request
		recurso := fmt.Sprintf("recurso-%d", i+1)
		procesarRequestSol3(authCtx, recurso)

		fmt.Println()
	}

	fmt.Println("✅ Solución 3 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 4: Context Deadline
// ==========================================

func procesamientoPorLotesSol4(ctx context.Context, elementos []string) []string {
	var procesados []string

	fmt.Printf("📦 Iniciando procesamiento de %d elementos\n", len(elementos))

	for i, elemento := range elementos {
		select {
		case <-ctx.Done():
			fmt.Printf("⏰ Procesamiento interrumpido en elemento %d: %v\n", i+1, ctx.Err())
			return procesados
		default:
			// Procesar elemento (100ms)
			time.Sleep(100 * time.Millisecond)
			procesado := "processed-" + elemento
			procesados = append(procesados, procesado)
			fmt.Printf("  ✅ Elemento %d/%d: %s -> %s\n", i+1, len(elementos), elemento, procesado)
		}
	}

	fmt.Printf("🎉 Procesamiento completo: %d elementos\n", len(procesados))
	return procesados
}

func solucion4() {
	fmt.Println("✅ Solución 4: Context Deadline")
	fmt.Println("===============================")

	elementos := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	// Context con deadline de 450ms (debería procesar ~4 elementos)
	deadline := time.Now().Add(450 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("⏰ Deadline establecido: %v\n", deadline.Format("15:04:05.000"))

	// Procesar elementos
	start := time.Now()
	procesados := procesamientoPorLotesSol4(ctx, elementos)
	duracion := time.Since(start)

	fmt.Printf("📊 Resultados:\n")
	fmt.Printf("  - Elementos procesados: %d/%d\n", len(procesados), len(elementos))
	fmt.Printf("  - Tiempo transcurrido: %v\n", duracion)
	fmt.Printf("  - Elementos: %v\n", procesados)

	fmt.Println("✅ Solución 4 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 5: Multiple Contexts
// ==========================================

func workerSol5(ctx context.Context, id int, resultChan chan<- string) {
	defer close(resultChan)

	iteraciones := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	fmt.Printf("👷 Worker %d iniciado\n", id)

	for {
		select {
		case <-ctx.Done():
			resultado := fmt.Sprintf("Worker %d terminado después de %d iteraciones: %v",
				id, iteraciones, ctx.Err())
			fmt.Printf("🔴 %s\n", resultado)
			resultChan <- resultado
			return
		case <-ticker.C:
			iteraciones++
			fmt.Printf("⚡ Worker %d - iteración %d\n", id, iteraciones)
		}
	}
}

func solucion5() {
	fmt.Println("✅ Solución 5: Multiple Contexts")
	fmt.Println("================================")

	// Context principal con timeout de 1 segundo
	mainCtx, mainCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer mainCancel()

	// Context cancelable manualmente
	manualCtx, manualCancel := context.WithCancel(context.Background())
	defer manualCancel()

	// Canales para resultados
	results := make([]chan string, 3)
	for i := range results {
		results[i] = make(chan string, 1)
	}

	// Lanzar workers
	go workerSol5(mainCtx, 1, results[0])
	go workerSol5(mainCtx, 2, results[1])
	go workerSol5(manualCtx, 3, results[2]) // Este usará el context manual

	// Cancelar worker 3 después de 300ms
	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("📤 Cancelando worker 3 manualmente...")
		manualCancel()
	}()

	// Recopilar resultados
	var wg sync.WaitGroup
	wg.Add(len(results))

	for i, resultChan := range results {
		go func(id int, ch <-chan string) {
			defer wg.Done()
			if result := <-ch; result != "" {
				fmt.Printf("📊 Resultado %d: %s\n", id+1, result)
			}
		}(i, resultChan)
	}

	wg.Wait()
	fmt.Println("✅ Solución 5 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 6: Context Middleware
// ==========================================

type MiddlewareSol6 func(context.Context, func(context.Context)) context.Context

func loggingMiddlewareSol6(ctx context.Context, next func(context.Context)) context.Context {
	start := time.Now()
	requestID := "req-" + fmt.Sprintf("%d", start.UnixNano()%10000)
	ctx = context.WithValue(ctx, RequestIDKey, requestID)

	fmt.Printf("🔍 [%s] Request iniciado\n", requestID)

	next(ctx)

	duration := time.Since(start)
	fmt.Printf("📊 [%s] Request completado en %v\n", requestID, duration)

	return ctx
}

func authMiddlewareSol6(ctx context.Context, next func(context.Context)) context.Context {
	requestID, _ := ctx.Value(RequestIDKey).(string)

	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		fmt.Printf("🔐 [%s] Usuario autenticado: %s\n", requestID, userID)
		next(ctx)
	} else {
		fmt.Printf("❌ [%s] Usuario no autenticado - bloqueando request\n", requestID)
	}

	return ctx
}

func handlerFinalSol6(ctx context.Context) {
	requestID, _ := ctx.Value(RequestIDKey).(string)
	userID, _ := ctx.Value(UserIDKey).(string)

	fmt.Printf("💼 [%s] Ejecutando lógica de negocio para %s\n", requestID, userID)

	// Simular trabajo que puede ser cancelado
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("🎉 [%s] Lógica de negocio completada\n", requestID)
	case <-ctx.Done():
		fmt.Printf("⏰ [%s] Lógica cancelada: %v\n", requestID, ctx.Err())
	}
}

func aplicarMiddlewares(ctx context.Context, middlewares []MiddlewareSol6, handler func(context.Context)) {
	var processedHandler func(context.Context) = handler

	// Aplicar middlewares en orden inverso
	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		currentHandler := processedHandler

		processedHandler = func(ctx context.Context) {
			middleware(ctx, currentHandler)
		}
	}

	processedHandler(ctx)
}

func solucion6() {
	fmt.Println("✅ Solución 6: Context Middleware")
	fmt.Println("=================================")

	middlewares := []MiddlewareSol6{
		loggingMiddlewareSol6,
		authMiddlewareSol6,
	}

	// Test 1: Context con userID válido
	fmt.Println("🎯 Test 1: Usuario autenticado")
	ctx1 := context.WithValue(context.Background(), UserIDKey, "user-123")
	aplicarMiddlewares(ctx1, middlewares, handlerFinalSol6)

	fmt.Println()

	// Test 2: Context sin userID
	fmt.Println("🎯 Test 2: Usuario no autenticado")
	ctx2 := context.Background()
	aplicarMiddlewares(ctx2, middlewares, handlerFinalSol6)

	fmt.Println()

	// Test 3: Context que se cancela
	fmt.Println("🎯 Test 3: Context con cancelación")
	ctx3, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	ctx3 = context.WithValue(ctx3, UserIDKey, "user-456")
	aplicarMiddlewares(ctx3, middlewares, handlerFinalSol6)

	time.Sleep(200 * time.Millisecond) // Esperar para ver el timeout
	fmt.Println("✅ Solución 6 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 7: Context Pipeline
// ==========================================

func etapa1Sol7(ctx context.Context, input string) (string, error) {
	select {
	case <-time.After(100 * time.Millisecond):
		result := "stage1-" + input
		fmt.Printf("✅ Etapa 1: %s -> %s\n", input, result)
		return result, nil
	case <-ctx.Done():
		fmt.Printf("❌ Etapa 1 cancelada: %v\n", ctx.Err())
		return "", ctx.Err()
	}
}

func etapa2Sol7(ctx context.Context, input string) (string, error) {
	select {
	case <-time.After(150 * time.Millisecond):
		result := "stage2-" + input
		fmt.Printf("✅ Etapa 2: %s -> %s\n", input, result)
		return result, nil
	case <-ctx.Done():
		fmt.Printf("❌ Etapa 2 cancelada: %v\n", ctx.Err())
		return "", ctx.Err()
	}
}

func etapa3Sol7(ctx context.Context, input string) (string, error) {
	select {
	case <-time.After(200 * time.Millisecond):
		result := "stage3-" + input
		fmt.Printf("✅ Etapa 3: %s -> %s\n", input, result)
		return result, nil
	case <-ctx.Done():
		fmt.Printf("❌ Etapa 3 cancelada: %v\n", ctx.Err())
		return "", ctx.Err()
	}
}

func pipelineSol7(ctx context.Context, input string) (string, error) {
	fmt.Printf("🔄 Iniciando pipeline para: %s\n", input)

	// Etapa 1
	result1, err := etapa1Sol7(ctx, input)
	if err != nil {
		return "", err
	}

	// Etapa 2
	result2, err := etapa2Sol7(ctx, result1)
	if err != nil {
		return "", err
	}

	// Etapa 3
	result3, err := etapa3Sol7(ctx, result2)
	if err != nil {
		return "", err
	}

	fmt.Printf("🎉 Pipeline completado: %s\n", result3)
	return result3, nil
}

func solucion7() {
	fmt.Println("✅ Solución 7: Context Pipeline")
	fmt.Println("===============================")

	inputs := []string{"data1", "data2", "data3"}

	// Input 1: timeout 600ms (debería completarse - total ~450ms)
	fmt.Println("🎯 Procesando data1 con timeout 600ms")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel1()

	start := time.Now()
	result, err := pipelineSol7(ctx1, inputs[0])
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error: %v (después de %v)\n", err, duration)
	} else {
		fmt.Printf("✅ Resultado: %s (en %v)\n", result, duration)
	}

	fmt.Println()

	// Input 2: timeout 300ms (debería fallar en etapa 2 o 3)
	fmt.Println("🎯 Procesando data2 con timeout 300ms")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel2()

	start = time.Now()
	result, err = pipelineSol7(ctx2, inputs[1])
	duration = time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error: %v (después de %v)\n", err, duration)
	} else {
		fmt.Printf("✅ Resultado: %s (en %v)\n", result, duration)
	}

	fmt.Println()

	// Input 3: cancelación manual después de 250ms
	fmt.Println("🎯 Procesando data3 con cancelación manual")
	ctx3, cancel3 := context.WithCancel(context.Background())

	// Cancelar después de 250ms
	go func() {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("📤 Enviando cancelación manual...")
		cancel3()
	}()

	start = time.Now()
	result, err = pipelineSol7(ctx3, inputs[2])
	duration = time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error: %v (después de %v)\n", err, duration)
	} else {
		fmt.Printf("✅ Resultado: %s (en %v)\n", result, duration)
	}

	fmt.Println("✅ Solución 7 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 8: Context Pool Worker
// ==========================================

type JobSol8 struct {
	ID   int
	Data string
}

type ResultSol8 struct {
	JobID  int
	Result string
	Error  error
}

type WorkerPoolSol8 struct {
	jobs       chan JobSol8
	results    chan ResultSol8
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	numWorkers int
}

func NewWorkerPoolSol8(numWorkers int) *WorkerPoolSol8 {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerPoolSol8{
		jobs:       make(chan JobSol8, numWorkers*2),
		results:    make(chan ResultSol8, numWorkers*2),
		ctx:        ctx,
		cancel:     cancel,
		numWorkers: numWorkers,
	}
}

func (wp *WorkerPoolSol8) Start() {
	fmt.Printf("🏭 Iniciando worker pool con %d workers\n", wp.numWorkers)

	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPoolSol8) worker(id int) {
	defer wp.wg.Done()
	fmt.Printf("👷 Worker %d iniciado\n", id)

	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				fmt.Printf("👷 Worker %d: canal cerrado, terminando\n", id)
				return
			}

			fmt.Printf("⚡ Worker %d procesando job %d: %s\n", id, job.ID, job.Data)

			// Simular trabajo de 200ms
			select {
			case <-time.After(200 * time.Millisecond):
				result := fmt.Sprintf("processed_%s_by_worker_%d", job.Data, id)
				wp.results <- ResultSol8{
					JobID:  job.ID,
					Result: result,
					Error:  nil,
				}
				fmt.Printf("✅ Worker %d completó job %d\n", id, job.ID)

			case <-wp.ctx.Done():
				wp.results <- ResultSol8{
					JobID:  job.ID,
					Result: "",
					Error:  wp.ctx.Err(),
				}
				fmt.Printf("❌ Worker %d: job %d cancelado\n", id, job.ID)
				return
			}

		case <-wp.ctx.Done():
			fmt.Printf("👷 Worker %d: cancelado\n", id)
			return
		}
	}
}

func (wp *WorkerPoolSol8) Submit(job JobSol8) {
	select {
	case wp.jobs <- job:
		fmt.Printf("📤 Job %d enviado al pool\n", job.ID)
	case <-wp.ctx.Done():
		fmt.Printf("❌ No se pudo enviar job %d: pool cerrado\n", job.ID)
	}
}

func (wp *WorkerPoolSol8) Stop() {
	fmt.Println("🛑 Deteniendo worker pool...")

	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
	wp.cancel()

	fmt.Println("✅ Worker pool detenido")
}

func solucion8() {
	fmt.Println("✅ Solución 8: Context Pool Worker")
	fmt.Println("==================================")

	// Crear worker pool con 3 workers
	pool := NewWorkerPoolSol8(3)
	pool.Start()

	// Goroutine para recopilar resultados
	go func() {
		for result := range pool.results {
			if result.Error != nil {
				fmt.Printf("❌ Job %d falló: %v\n", result.JobID, result.Error)
			} else {
				fmt.Printf("📊 Job %d resultado: %s\n", result.JobID, result.Result)
			}
		}
	}()

	// Enviar 5 jobs
	for i := 1; i <= 5; i++ {
		job := JobSol8{
			ID:   i,
			Data: fmt.Sprintf("task-%d", i),
		}
		pool.Submit(job)
	}

	// Recopilar resultados por 1 segundo
	fmt.Println("⏱️ Recopilando resultados por 1 segundo...")
	time.Sleep(1 * time.Second)

	// Detener pool
	pool.Stop()

	// Dar tiempo para que se muestren los últimos resultados
	time.Sleep(200 * time.Millisecond)
	fmt.Println("✅ Solución 8 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 9: Context con Select
// ==========================================

func servicioExternoSol9(ctx context.Context, servicio string) (string, error) {
	var duracion time.Duration

	switch servicio {
	case "rapido":
		duracion = 100 * time.Millisecond
	case "medio":
		duracion = 300 * time.Millisecond
	case "lento":
		duracion = 800 * time.Millisecond
	default:
		duracion = 500 * time.Millisecond
	}

	fmt.Printf("🌐 Consultando servicio '%s' (durará %v)\n", servicio, duracion)

	select {
	case <-time.After(duracion):
		resultado := fmt.Sprintf("respuesta_de_%s", servicio)
		fmt.Printf("✅ Servicio '%s' respondió: %s\n", servicio, resultado)
		return resultado, nil
	case <-ctx.Done():
		fmt.Printf("❌ Servicio '%s' cancelado: %v\n", servicio, ctx.Err())
		return "", ctx.Err()
	}
}

func solucion9() {
	fmt.Println("✅ Solución 9: Context con Select")
	fmt.Println("=================================")

	servicios := []string{"rapido", "medio", "lento"}

	// Context con timeout global de 500ms
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Canal para resultados
	resultChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	// Lanzar llamadas concurrentes
	fmt.Println("🚀 Iniciando llamadas concurrentes a servicios...")

	for _, servicio := range servicios {
		go func(s string) {
			result, err := servicioExternoSol9(ctx, s)
			if err != nil {
				select {
				case errorChan <- err:
				default: // Canal lleno, ignorar
				}
			} else {
				select {
				case resultChan <- result:
				default: // Canal lleno, ignorar
				}
			}
		}(servicio)
	}

	// Esperar el primer resultado o timeout
	start := time.Now()

	select {
	case result := <-resultChan:
		duration := time.Since(start)
		fmt.Printf("🎉 Primer resultado obtenido: %s (en %v)\n", result, duration)

		// Cancelar operaciones restantes
		fmt.Println("📤 Cancelando operaciones restantes...")
		cancel()

	case <-ctx.Done():
		duration := time.Since(start)
		fmt.Printf("⏰ Timeout global alcanzado después de %v: %v\n", duration, ctx.Err())

	case err := <-errorChan:
		duration := time.Since(start)
		fmt.Printf("❌ Error en servicio después de %v: %v\n", duration, err)
	}

	// Dar tiempo para que se muestren las cancelaciones
	time.Sleep(200 * time.Millisecond)
	fmt.Println("✅ Solución 9 completada\n")
}

// ==========================================
// ✅ SOLUCIÓN 10: Context Composition
// ==========================================

type CompositeContextSol10 struct {
	contexts []context.Context
	done     chan struct{}
	err      error
	once     sync.Once
}

func NewCompositeContextSol10(contexts ...context.Context) *CompositeContextSol10 {
	cc := &CompositeContextSol10{
		contexts: contexts,
		done:     make(chan struct{}),
	}

	// Monitorear todos los contexts
	for _, ctx := range contexts {
		go func(c context.Context) {
			<-c.Done()
			cc.cancel(c.Err())
		}(ctx)
	}

	return cc
}

func (cc *CompositeContextSol10) cancel(err error) {
	cc.once.Do(func() {
		cc.err = err
		close(cc.done)
	})
}

func (cc *CompositeContextSol10) Done() <-chan struct{} {
	return cc.done
}

func (cc *CompositeContextSol10) Err() error {
	select {
	case <-cc.done:
		return cc.err
	default:
		return nil
	}
}

func (cc *CompositeContextSol10) Deadline() (time.Time, bool) {
	var earliest time.Time
	var hasDeadline bool

	for _, ctx := range cc.contexts {
		if deadline, ok := ctx.Deadline(); ok {
			if !hasDeadline || deadline.Before(earliest) {
				earliest = deadline
				hasDeadline = true
			}
		}
	}

	return earliest, hasDeadline
}

func (cc *CompositeContextSol10) Value(key interface{}) interface{} {
	for _, ctx := range cc.contexts {
		if value := ctx.Value(key); value != nil {
			return value
		}
	}
	return nil
}

func operacionComplicadaSol10(ctx context.Context, nombre string) {
	userID := ctx.Value(UserIDKey)
	requestID := ctx.Value(RequestIDKey)

	fmt.Printf("🔄 [%v] [%v] Iniciando %s\n", requestID, userID, nombre)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	iteracion := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🔴 [%v] %s cancelada en iteración %d: %v\n",
				requestID, nombre, iteracion, ctx.Err())
			return
		case <-ticker.C:
			iteracion++
			fmt.Printf("⚡ [%v] %s - iteración %d\n", requestID, nombre, iteracion)

			if iteracion >= 8 {
				fmt.Printf("✅ [%v] %s completada exitosamente\n", requestID, nombre)
				return
			}
		}
	}
}

func solucion10() {
	fmt.Println("✅ Solución 10: Context Composition")
	fmt.Println("===================================")

	// Crear múltiples contexts

	// 1. Context con valores
	valueCtx := context.Background()
	valueCtx = context.WithValue(valueCtx, UserIDKey, "user-123")
	valueCtx = context.WithValue(valueCtx, RequestIDKey, "req-456")

	// 2. Context con timeout de 2 segundos
	timeoutCtx, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()

	// 3. Context cancelable manualmente
	cancelableCtx, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	// Crear composite context
	compositeCtx := NewCompositeContextSol10(valueCtx, timeoutCtx, cancelableCtx)

	fmt.Println("🏗️ Context compuesto creado con:")
	fmt.Printf("   👤 User ID: %v\n", compositeCtx.Value(UserIDKey))
	fmt.Printf("   📨 Request ID: %v\n", compositeCtx.Value(RequestIDKey))

	if deadline, ok := compositeCtx.Deadline(); ok {
		fmt.Printf("   ⏰ Deadline: %v\n", deadline.Format("15:04:05"))
	}

	// Lanzar operación
	go operacionComplicadaSol10(compositeCtx, "Operación Compuesta")

	// Test 1: Verificar que los valores funcionan
	fmt.Println("\n🧪 Test 1: Acceso a valores")
	if userID := compositeCtx.Value(UserIDKey); userID != nil {
		fmt.Printf("✅ UserID obtenido: %v\n", userID)
	}

	// Test 2: Simular cancelación manual después de 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("📤 Activando cancelación manual...")
		cancel2()
	}()

	// Esperar a que termine
	<-compositeCtx.Done()

	// Test 3: Verificar que el error se propaga
	fmt.Printf("🏁 Context terminado con error: %v\n", compositeCtx.Err())

	time.Sleep(300 * time.Millisecond)
	fmt.Println("✅ Solución 10 completada\n")
}

// ==========================================
// 🏃‍♂️ FUNCIÓN PRINCIPAL
// ==========================================

func main() {
	fmt.Println("💡 Soluciones de Context Package")
	fmt.Println("=================================")
	fmt.Println("✨ Implementaciones completas de todos los ejercicios")
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

	fmt.Println("🎉 ¡Todas las soluciones ejecutadas!")
	fmt.Println("🎯 Has dominado el Context package de Go")
}
