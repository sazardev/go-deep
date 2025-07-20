// ==============================================
// PROYECTO: Sistema de Procesamiento de Datos Concurrente
// ==============================================
// Simulador de sistema de análisis de datos en tiempo real
// que procesa streams de información usando patrones de concurrencia

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ==============================================
// ESTRUCTURAS DE DATOS
// ==============================================

// Evento representa un evento de datos entrante
type Evento struct {
	ID        int64     `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserID    int       `json:"user_id"`
	Action    string    `json:"action"`
	Data      string    `json:"data"`
	Priority  int       `json:"priority"` // 1=alta, 2=media, 3=baja
}

// EventoProcesado representa un evento después de procesamiento
type EventoProcesado struct {
	EventoOriginal Evento        `json:"evento_original"`
	Procesado      time.Time     `json:"procesado"`
	Resultado      string        `json:"resultado"`
	ProcessorID    int           `json:"processor_id"`
	TiempoProceso  time.Duration `json:"tiempo_proceso"`
}

// Estadisticas del sistema
type Estadisticas struct {
	EventosProcesados  int64         `json:"eventos_procesados"`
	EventosDescartados int64         `json:"eventos_descartados"`
	TiempoPromedio     time.Duration `json:"tiempo_promedio"`
	ErroresTotales     int64         `json:"errores_totales"`
	ProcessorsActivos  int           `json:"processors_activos"`
}

// ==============================================
// GENERADOR DE EVENTOS (Producer)
// ==============================================

type GeneradorEventos struct {
	eventosGenerados int64
	activo           bool
	mu               sync.RWMutex
}

func NewGeneradorEventos() *GeneradorEventos {
	return &GeneradorEventos{
		activo: true,
	}
}

func (g *GeneradorEventos) GenerarEventos(ctx context.Context, output chan<- Evento, eventosPerSec int) {
	ticker := time.NewTicker(time.Second / time.Duration(eventosPerSec))
	defer ticker.Stop()

	actions := []string{"click", "view", "purchase", "search", "logout", "login", "share"}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🛑 Generador de eventos deteniéndose...")
			return
		case <-ticker.C:
			// Generar evento aleatorio
			evento := Evento{
				ID:        atomic.AddInt64(&g.eventosGenerados, 1),
				Timestamp: time.Now(),
				UserID:    rand.Intn(10000),
				Action:    actions[rand.Intn(len(actions))],
				Data:      fmt.Sprintf("data_%d", rand.Intn(1000)),
				Priority:  rand.Intn(3) + 1,
			}

			select {
			case output <- evento:
				// Evento enviado exitosamente
			default:
				// Canal lleno, descartar evento
				fmt.Printf("⚠️ Evento %d descartado - canal lleno\n", evento.ID)
			}
		}
	}
}

func (g *GeneradorEventos) GetEventosGenerados() int64 {
	return atomic.LoadInt64(&g.eventosGenerados)
}

// ==============================================
// PROCESADOR DE EVENTOS (Worker)
// ==============================================

type ProcesadorEventos struct {
	id                int
	eventosProcesados int64
	errores           int64
	tiempoTotal       int64 // en nanosegundos
	activo            bool
}

func NewProcesadorEventos(id int) *ProcesadorEventos {
	return &ProcesadorEventos{
		id:     id,
		activo: true,
	}
}

func (p *ProcesadorEventos) ProcesarEventos(ctx context.Context, input <-chan Evento, output chan<- EventoProcesado, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("🔧 Procesador %d iniciado\n", p.id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 Procesador %d deteniéndose...\n", p.id)
			return
		case evento, ok := <-input:
			if !ok {
				fmt.Printf("✅ Procesador %d terminó - canal cerrado\n", p.id)
				return
			}

			inicio := time.Now()

			// Simular procesamiento complejo
			resultado := p.procesarEvento(evento)

			tiempoProceso := time.Since(inicio)

			// Crear evento procesado
			eventoProcesado := EventoProcesado{
				EventoOriginal: evento,
				Procesado:      time.Now(),
				Resultado:      resultado,
				ProcessorID:    p.id,
				TiempoProceso:  tiempoProceso,
			}

			// Enviar resultado
			select {
			case output <- eventoProcesado:
				atomic.AddInt64(&p.eventosProcesados, 1)
				atomic.AddInt64(&p.tiempoTotal, int64(tiempoProceso))
			case <-ctx.Done():
				return
			}
		}
	}
}

func (p *ProcesadorEventos) procesarEvento(evento Evento) string {
	// Simular tiempo de procesamiento basado en prioridad
	var tiempoProceso time.Duration
	switch evento.Priority {
	case 1: // Alta prioridad - procesamiento rápido
		tiempoProceso = time.Duration(rand.Intn(50)) * time.Millisecond
	case 2: // Media prioridad
		tiempoProceso = time.Duration(rand.Intn(100)+50) * time.Millisecond
	case 3: // Baja prioridad
		tiempoProceso = time.Duration(rand.Intn(200)+100) * time.Millisecond
	}

	time.Sleep(tiempoProceso)

	// Simular ocasional error (5% probabilidad)
	if rand.Float32() < 0.05 {
		atomic.AddInt64(&p.errores, 1)
		return fmt.Sprintf("ERROR: Falló procesamiento de %s", evento.Action)
	}

	// Procesamiento exitoso
	return fmt.Sprintf("Procesado: %s para usuario %d con prioridad %d",
		evento.Action, evento.UserID, evento.Priority)
}

func (p *ProcesadorEventos) GetEstadisticas() (int64, int64, time.Duration) {
	procesados := atomic.LoadInt64(&p.eventosProcesados)
	errores := atomic.LoadInt64(&p.errores)
	tiempoTotal := atomic.LoadInt64(&p.tiempoTotal)

	var tiempoPromedio time.Duration
	if procesados > 0 {
		tiempoPromedio = time.Duration(tiempoTotal / procesados)
	}

	return procesados, errores, tiempoPromedio
}

// ==============================================
// AGREGADOR DE RESULTADOS (Consumer)
// ==============================================

type AgregadorResultados struct {
	resultadosProcesados  int64
	resultadosDescartados int64
	estadisticas          map[string]int64
	mu                    sync.RWMutex
}

func NewAgregadorResultados() *AgregadorResultados {
	return &AgregadorResultados{
		estadisticas: make(map[string]int64),
	}
}

func (a *AgregadorResultados) AgregarResultados(ctx context.Context, input <-chan EventoProcesado, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("📊 Agregador de resultados iniciado")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🛑 Agregador deteniéndose...")
			return
		case resultado, ok := <-input:
			if !ok {
				fmt.Println("✅ Agregador terminó - canal cerrado")
				return
			}

			// Procesar resultado
			a.procesarResultado(resultado)
		}
	}
}

func (a *AgregadorResultados) procesarResultado(resultado EventoProcesado) {
	a.mu.Lock()
	defer a.mu.Unlock()

	atomic.AddInt64(&a.resultadosProcesados, 1)

	// Actualizar estadísticas por acción
	action := resultado.EventoOriginal.Action
	a.estadisticas[action]++

	// Log cada 100 resultados
	if a.resultadosProcesados%100 == 0 {
		fmt.Printf("📈 Resultados procesados: %d\n", a.resultadosProcesados)
	}
}

func (a *AgregadorResultados) GetEstadisticas() (int64, map[string]int64) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	// Crear copia del mapa para evitar race conditions
	estadisticasCopia := make(map[string]int64)
	for k, v := range a.estadisticas {
		estadisticasCopia[k] = v
	}

	return atomic.LoadInt64(&a.resultadosProcesados), estadisticasCopia
}

// ==============================================
// MONITOR DEL SISTEMA
// ==============================================

type MonitorSistema struct {
	generador    *GeneradorEventos
	procesadores []*ProcesadorEventos
	agregador    *AgregadorResultados
}

func NewMonitorSistema(generador *GeneradorEventos, procesadores []*ProcesadorEventos, agregador *AgregadorResultados) *MonitorSistema {
	return &MonitorSistema{
		generador:    generador,
		procesadores: procesadores,
		agregador:    agregador,
	}
}

func (m *MonitorSistema) Monitorear(ctx context.Context, intervalo time.Duration) {
	ticker := time.NewTicker(intervalo)
	defer ticker.Stop()

	fmt.Println("📊 Monitor del sistema iniciado")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🛑 Monitor deteniéndose...")
			return
		case <-ticker.C:
			m.mostrarEstadisticas()
		}
	}
}

func (m *MonitorSistema) mostrarEstadisticas() {
	separador := strings.Repeat("=", 60)
	fmt.Println("\n" + separador)
	fmt.Println("📊 ESTADÍSTICAS DEL SISTEMA")
	fmt.Println(separador)

	// Estadísticas del generador
	eventosGenerados := m.generador.GetEventosGenerados()
	fmt.Printf("🔢 Eventos generados: %d\n", eventosGenerados)

	// Estadísticas de procesadores
	var totalProcesados, totalErrores int64
	var tiempoPromedioTotal time.Duration

	fmt.Println("\n👷 Procesadores:")
	for _, procesador := range m.procesadores {
		procesados, errores, tiempoPromedio := procesador.GetEstadisticas()
		totalProcesados += procesados
		totalErrores += errores
		tiempoPromedioTotal += tiempoPromedio

		fmt.Printf("  Procesador %d: %d procesados, %d errores, %v tiempo promedio\n",
			procesador.id, procesados, errores, tiempoPromedio)
	}

	// Estadísticas del agregador
	resultados, estadisticasAcciones := m.agregador.GetEstadisticas()
	fmt.Printf("\n📈 Resultados agregados: %d\n", resultados)

	// Estadísticas por acción
	fmt.Println("\n📋 Estadísticas por acción:")
	for accion, cantidad := range estadisticasAcciones {
		fmt.Printf("  %s: %d\n", accion, cantidad)
	}

	// Métricas del sistema
	fmt.Printf("\n🖥️ Sistema: %d goroutines activas, %d CPUs\n",
		runtime.NumGoroutine(), runtime.NumCPU())

	// Throughput
	if eventosGenerados > 0 {
		eficiencia := float64(totalProcesados) / float64(eventosGenerados) * 100
		fmt.Printf("⚡ Eficiencia: %.1f%% (%d/%d)\n", eficiencia, totalProcesados, eventosGenerados)
	}

	fmt.Println(separador)
}

// ==============================================
// SISTEMA PRINCIPAL
// ==============================================

type SistemaProcesamiento struct {
	generador    *GeneradorEventos
	procesadores []*ProcesadorEventos
	agregador    *AgregadorResultados
	monitor      *MonitorSistema

	canalEventos    chan Evento
	canalResultados chan EventoProcesado
}

func NewSistemaProcesamiento(numProcesadores int, bufferSize int) *SistemaProcesamiento {
	// Crear componentes
	generador := NewGeneradorEventos()
	agregador := NewAgregadorResultados()

	// Crear procesadores
	procesadores := make([]*ProcesadorEventos, numProcesadores)
	for i := 0; i < numProcesadores; i++ {
		procesadores[i] = NewProcesadorEventos(i + 1)
	}

	// Crear monitor
	monitor := NewMonitorSistema(generador, procesadores, agregador)

	// Crear canales
	canalEventos := make(chan Evento, bufferSize)
	canalResultados := make(chan EventoProcesado, bufferSize)

	return &SistemaProcesamiento{
		generador:       generador,
		procesadores:    procesadores,
		agregador:       agregador,
		monitor:         monitor,
		canalEventos:    canalEventos,
		canalResultados: canalResultados,
	}
}

func (s *SistemaProcesamiento) Ejecutar(duracion time.Duration, eventosPerSec int) {
	ctx, cancel := context.WithTimeout(context.Background(), duracion)
	defer cancel()

	var wg sync.WaitGroup

	fmt.Println("🚀 INICIANDO SISTEMA DE PROCESAMIENTO CONCURRENTE")
	fmt.Printf("📊 Configuración: %d procesadores, %d eventos/seg, duración %v\n\n",
		len(s.procesadores), eventosPerSec, duracion)

	// Iniciar generador de eventos
	go s.generador.GenerarEventos(ctx, s.canalEventos, eventosPerSec)

	// Iniciar procesadores
	for _, procesador := range s.procesadores {
		wg.Add(1)
		go procesador.ProcesarEventos(ctx, s.canalEventos, s.canalResultados, &wg)
	}

	// Iniciar agregador
	wg.Add(1)
	go s.agregador.AgregarResultados(ctx, s.canalResultados, &wg)

	// Iniciar monitor
	go s.monitor.Monitorear(ctx, 2*time.Second)

	// Esperar a que termine el contexto
	<-ctx.Done()

	// Cerrar canales para permitir que las goroutines terminen limpiamente
	close(s.canalEventos)

	// Esperar a que todos los procesadores terminen
	wg.Wait()
	close(s.canalResultados)

	// Mostrar estadísticas finales
	fmt.Println("\n🏁 SISTEMA DETENIDO - ESTADÍSTICAS FINALES:")
	s.monitor.mostrarEstadisticas()
}

// ==============================================
// FUNCIÓN PRINCIPAL
// ==============================================

func main() {
	fmt.Println("🏭 PROYECTO: Sistema de Procesamiento de Datos Concurrente")
	fmt.Println("==========================================================")

	// Configuración del sistema
	numProcesadores := runtime.NumCPU() // Usar todos los CPUs disponibles
	bufferSize := 1000                  // Buffer de canales
	eventosPerSec := 50                 // Eventos por segundo
	duracion := 10 * time.Second        // Duración de la simulación

	fmt.Printf("🖥️ CPUs disponibles: %d\n", runtime.NumCPU())
	fmt.Printf("⚙️ Procesadores configurados: %d\n", numProcesadores)
	fmt.Printf("📦 Buffer de canales: %d\n", bufferSize)
	fmt.Printf("🔢 Eventos por segundo: %d\n", eventosPerSec)
	fmt.Printf("⏱️ Duración: %v\n\n", duracion)

	// Configurar random seed
	rand.Seed(time.Now().UnixNano())

	// Crear y ejecutar sistema
	sistema := NewSistemaProcesamiento(numProcesadores, bufferSize)
	sistema.Ejecutar(duracion, eventosPerSec)

	fmt.Println("\n✅ Simulación completada exitosamente!")
	fmt.Println("\n💡 Conceptos demostrados:")
	fmt.Println("   🔄 Producer-Consumer pattern")
	fmt.Println("   👷 Worker Pool pattern")
	fmt.Println("   📊 Monitoring y métricas")
	fmt.Println("   🎯 Context para cancelación")
	fmt.Println("   🔒 Operaciones atómicas")
	fmt.Println("   📡 Comunicación vía channels")
	fmt.Println("   ⚡ Procesamiento en tiempo real")

	// Estadísticas finales de goroutines
	time.Sleep(100 * time.Millisecond) // Permitir cleanup
	fmt.Printf("\n📊 Goroutines finales: %d\n", runtime.NumGoroutine())
}
