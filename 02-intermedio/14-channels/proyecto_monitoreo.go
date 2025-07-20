// ==============================================
// LECCIÓN 14: Channels - Proyecto Sistema de Monitoreo
// ==============================================

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
// SISTEMA DE MONITOREO EN TIEMPO REAL
// ==============================================

// Tipos de eventos del sistema
type TipoEvento int

const (
	EventoCPU TipoEvento = iota
	EventoMemoria
	EventoRed
	EventoDisco
	EventoError
)

func (t TipoEvento) String() string {
	switch t {
	case EventoCPU:
		return "CPU"
	case EventoMemoria:
		return "MEMORIA"
	case EventoRed:
		return "RED"
	case EventoDisco:
		return "DISCO"
	case EventoError:
		return "ERROR"
	default:
		return "DESCONOCIDO"
	}
}

// Estructura de evento
type Evento struct {
	Timestamp time.Time
	Tipo      TipoEvento
	Servicio  string
	Valor     float64
	Metadata  map[string]string
}

// Métricas del sistema
type Metricas struct {
	eventosTotal     int64
	eventosPorTipo   map[TipoEvento]int64
	ultimoEvento     time.Time
	eventosError     int64
	promedioLatencia float64
	mu               sync.RWMutex
}

// Sistema de monitoreo
type SistemaMonitoreo struct {
	eventos      chan Evento
	alertas      chan Evento
	estadisticas chan map[string]interface{}
	quit         chan bool
	metricas     *Metricas
	ctx          context.Context
	cancel       context.CancelFunc
}

// ==============================================
// GENERADORES DE EVENTOS
// ==============================================

func generadorCPU(eventos chan<- Evento, ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🔄 Generador CPU terminando...")
			return
		case <-ticker.C:
			uso := rand.Float64() * 100
			evento := Evento{
				Timestamp: time.Now(),
				Tipo:      EventoCPU,
				Servicio:  "sistema",
				Valor:     uso,
				Metadata: map[string]string{
					"unidad": "porcentaje",
					"core":   fmt.Sprintf("core-%d", rand.Intn(4)),
				},
			}

			select {
			case eventos <- evento:
			case <-ctx.Done():
				return
			}
		}
	}
}

func generadorMemoria(eventos chan<- Evento, ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("💾 Generador Memoria terminando...")
			return
		case <-ticker.C:
			uso := 2048 + rand.Float64()*6144 // 2-8 GB
			evento := Evento{
				Timestamp: time.Now(),
				Tipo:      EventoMemoria,
				Servicio:  "sistema",
				Valor:     uso,
				Metadata: map[string]string{
					"unidad": "MB",
					"tipo":   "RAM",
				},
			}

			select {
			case eventos <- evento:
			case <-ctx.Done():
				return
			}
		}
	}
}

func generadorRed(eventos chan<- Evento, ctx context.Context) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🌐 Generador Red terminando...")
			return
		case <-ticker.C:
			throughput := rand.Float64() * 1000 // Mbps
			evento := Evento{
				Timestamp: time.Now(),
				Tipo:      EventoRed,
				Servicio:  "networking",
				Valor:     throughput,
				Metadata: map[string]string{
					"unidad":    "Mbps",
					"interface": fmt.Sprintf("eth%d", rand.Intn(3)),
				},
			}

			select {
			case eventos <- evento:
			case <-ctx.Done():
				return
			}
		}
	}
}

func generadorErrores(eventos chan<- Evento, ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("❌ Generador Errores terminando...")
			return
		case <-ticker.C:
			// Generar error ocasional
			if rand.Float64() < 0.3 {
				severidad := rand.Float64() * 10
				evento := Evento{
					Timestamp: time.Now(),
					Tipo:      EventoError,
					Servicio:  fmt.Sprintf("servicio-%d", rand.Intn(5)),
					Valor:     severidad,
					Metadata: map[string]string{
						"codigo": fmt.Sprintf("ERR-%03d", rand.Intn(999)),
						"nivel":  obtenerNivelError(severidad),
					},
				}

				select {
				case eventos <- evento:
				case <-ctx.Done():
					return
				}
			}
		}
	}
}

func obtenerNivelError(severidad float64) string {
	switch {
	case severidad < 3:
		return "INFO"
	case severidad < 6:
		return "WARNING"
	case severidad < 8:
		return "ERROR"
	default:
		return "CRITICAL"
	}
}

// ==============================================
// PROCESADORES DE EVENTOS
// ==============================================

func procesadorEventos(sistema *SistemaMonitoreo) {
	fmt.Println("📊 Procesador de eventos iniciado")

	for {
		select {
		case <-sistema.ctx.Done():
			fmt.Println("📊 Procesador de eventos terminando...")
			return

		case evento := <-sistema.eventos:
			// Actualizar métricas
			atomic.AddInt64(&sistema.metricas.eventosTotal, 1)

			sistema.metricas.mu.Lock()
			if sistema.metricas.eventosPorTipo == nil {
				sistema.metricas.eventosPorTipo = make(map[TipoEvento]int64)
			}
			sistema.metricas.eventosPorTipo[evento.Tipo]++
			sistema.metricas.ultimoEvento = evento.Timestamp
			sistema.metricas.mu.Unlock()

			// Procesar según tipo
			switch evento.Tipo {
			case EventoCPU:
				if evento.Valor > 90 {
					enviarAlerta(sistema.alertas, evento, "CPU crítico")
				}
			case EventoMemoria:
				if evento.Valor > 7000 { // > 7GB
					enviarAlerta(sistema.alertas, evento, "Memoria alta")
				}
			case EventoRed:
				if evento.Valor < 10 {
					enviarAlerta(sistema.alertas, evento, "Conectividad baja")
				}
			case EventoError:
				atomic.AddInt64(&sistema.metricas.eventosError, 1)
				if evento.Valor > 7 {
					enviarAlerta(sistema.alertas, evento, "Error crítico")
				}
			}

			// Log del evento
			fmt.Printf("📈 [%s] %s: %.2f %s\n",
				evento.Timestamp.Format("15:04:05"),
				evento.Tipo,
				evento.Valor,
				evento.Metadata["unidad"])
		}
	}
}

func enviarAlerta(alertas chan<- Evento, evento Evento, descripcion string) {
	evento.Metadata["alerta"] = descripcion

	select {
	case alertas <- evento:
		// Alerta enviada
	default:
		// Canal de alertas lleno, alerta perdida
		fmt.Printf("⚠️ Canal de alertas lleno, alerta perdida: %s\n", descripcion)
	}
}

func procesadorAlertas(alertas <-chan Evento, ctx context.Context) {
	fmt.Println("🚨 Procesador de alertas iniciado")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🚨 Procesador de alertas terminando...")
			return

		case alerta := <-alertas:
			nivel := "🟡"
			if alerta.Valor > 8 {
				nivel = "🔴"
			}

			fmt.Printf("%s ALERTA [%s]: %s - %.2f\n",
				nivel,
				alerta.Tipo,
				alerta.Metadata["alerta"],
				alerta.Valor)
		}
	}
}

func generadorEstadisticas(sistema *SistemaMonitoreo) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	fmt.Println("📊 Generador de estadísticas iniciado")

	for {
		select {
		case <-sistema.ctx.Done():
			fmt.Println("📊 Generador de estadísticas terminando...")
			return

		case <-ticker.C:
			stats := calcularEstadisticas(sistema.metricas)

			select {
			case sistema.estadisticas <- stats:
			case <-sistema.ctx.Done():
				return
			}
		}
	}
}

func calcularEstadisticas(metricas *Metricas) map[string]interface{} {
	metricas.mu.RLock()
	defer metricas.mu.RUnlock()

	total := atomic.LoadInt64(&metricas.eventosTotal)
	errores := atomic.LoadInt64(&metricas.eventosError)

	stats := map[string]interface{}{
		"eventos_total":    total,
		"eventos_error":    errores,
		"tasa_error":       float64(errores) / float64(total) * 100,
		"ultimo_evento":    metricas.ultimoEvento.Format("15:04:05"),
		"eventos_por_tipo": metricas.eventosPorTipo,
	}

	return stats
}

func procesadorEstadisticas(estadisticas <-chan map[string]interface{}, ctx context.Context) {
	fmt.Println("📈 Procesador de estadísticas iniciado")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("📈 Procesador de estadísticas terminando...")
			return

		case stats := <-estadisticas:
			separador := strings.Repeat("=", 50)
			fmt.Println("\n" + separador)
			fmt.Println("📊 ESTADÍSTICAS DEL SISTEMA")
			fmt.Println(separador)
			fmt.Printf("Total eventos: %v\n", stats["eventos_total"])
			fmt.Printf("Eventos error: %v\n", stats["eventos_error"])
			fmt.Printf("Tasa de error: %.2f%%\n", stats["tasa_error"])
			fmt.Printf("Último evento: %v\n", stats["ultimo_evento"])

			if porTipo, ok := stats["eventos_por_tipo"].(map[TipoEvento]int64); ok {
				fmt.Println("\nEventos por tipo:")
				for tipo, count := range porTipo {
					fmt.Printf("  %s: %d\n", tipo, count)
				}
			}
			fmt.Println(separador + "\n")
		}
	}
}

// ==============================================
// SISTEMA PRINCIPAL
// ==============================================

func NewSistemaMonitoreo() *SistemaMonitoreo {
	ctx, cancel := context.WithCancel(context.Background())

	return &SistemaMonitoreo{
		eventos:      make(chan Evento, 100),
		alertas:      make(chan Evento, 50),
		estadisticas: make(chan map[string]interface{}, 10),
		quit:         make(chan bool),
		metricas: &Metricas{
			eventosPorTipo: make(map[TipoEvento]int64),
		},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *SistemaMonitoreo) Iniciar() {
	fmt.Println("🚀 Iniciando Sistema de Monitoreo")
	fmt.Println("==================================")

	// Lanzar generadores de eventos
	go generadorCPU(s.eventos, s.ctx)
	go generadorMemoria(s.eventos, s.ctx)
	go generadorRed(s.eventos, s.ctx)
	go generadorErrores(s.eventos, s.ctx)

	// Lanzar procesadores
	go procesadorEventos(s)
	go procesadorAlertas(s.alertas, s.ctx)
	go generadorEstadisticas(s)
	go procesadorEstadisticas(s.estadisticas, s.ctx)

	fmt.Println("✅ Todos los componentes iniciados")
}

func (s *SistemaMonitoreo) Detener() {
	fmt.Println("\n🛑 Iniciando shutdown del sistema...")

	// Cancelar context para señalar a todas las goroutines
	s.cancel()

	// Dar tiempo para que terminen limpiamente
	time.Sleep(500 * time.Millisecond)

	// Cerrar channels
	close(s.eventos)
	close(s.alertas)
	close(s.estadisticas)
	close(s.quit)

	fmt.Println("✅ Sistema detenido correctamente")
}

func (s *SistemaMonitoreo) EstadisticasFinales() {
	total := atomic.LoadInt64(&s.metricas.eventosTotal)
	errores := atomic.LoadInt64(&s.metricas.eventosError)

	separador := strings.Repeat("=", 50)
	fmt.Println("\n" + separador)
	fmt.Println("📋 ESTADÍSTICAS FINALES")
	fmt.Println(separador)
	fmt.Printf("Total de eventos procesados: %d\n", total)
	fmt.Printf("Total de errores: %d\n", errores)

	if total > 0 {
		fmt.Printf("Tasa de error: %.2f%%\n", float64(errores)/float64(total)*100)
	}

	s.metricas.mu.RLock()
	fmt.Println("\nDistribución por tipo:")
	for tipo, count := range s.metricas.eventosPorTipo {
		porcentaje := float64(count) / float64(total) * 100
		fmt.Printf("  %s: %d (%.1f%%)\n", tipo, count, porcentaje)
	}
	s.metricas.mu.RUnlock()

	fmt.Println(separador)
}

// ==============================================
// FUNCIÓN PRINCIPAL DEL PROYECTO
// ==============================================

func EjecutarProyectoMonitoreo() {
	fmt.Println("📡 PROYECTO: Sistema de Monitoreo con Channels")
	fmt.Println("==============================================")
	fmt.Println("Este proyecto demuestra el uso avanzado de channels en Go:")
	fmt.Println("✅ Múltiples productores y consumidores")
	fmt.Println("✅ Fan-Out/Fan-In patterns")
	fmt.Println("✅ Buffered channels para performance")
	fmt.Println("✅ Context para cancelación elegante")
	fmt.Println("✅ Select para multiplexación")
	fmt.Println("✅ Quit channels para shutdown")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	// Crear y configurar sistema
	sistema := NewSistemaMonitoreo()

	// Iniciar sistema
	sistema.Iniciar()

	// Ejecutar por 15 segundos
	fmt.Println("⏱️ Sistema ejecutándose por 15 segundos...")
	tiempo := time.NewTimer(15 * time.Second)

	<-tiempo.C

	// Detener sistema
	sistema.Detener()

	// Mostrar estadísticas finales
	sistema.EstadisticasFinales()

	fmt.Println("\n💡 Conceptos demostrados:")
	fmt.Println("   📡 Channels buffered y unbuffered")
	fmt.Println("   🎛️ Select para multiplexación")
	fmt.Println("   🔄 Productores y consumidores concurrentes")
	fmt.Println("   📊 Agregación de datos en tiempo real")
	fmt.Println("   🚨 Sistema de alertas con channels")
	fmt.Println("   🛑 Shutdown elegante con context")
	fmt.Println("   ⚡ Performance con buffering estratégico")

	fmt.Printf("\n📊 Goroutines finales: %d\n", runtime.NumGoroutine())
}
