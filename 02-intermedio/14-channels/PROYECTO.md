# 📋 Proyecto: Sistema de Monitoreo con Channels

## 🎯 Objetivo del Proyecto

Desarrollar un **sistema de monitoreo en tiempo real** que demuestre el uso avanzado de channels en Go para:
- Recolección de métricas del sistema
- Procesamiento concurrente de eventos
- Sistema de alertas
- Agregación de estadísticas
- Shutdown elegante

---

## 🏗️ Arquitectura del Sistema

```
   🏭 GENERADORES                    📊 PROCESADORES                 📈 SALIDAS
   ===============                   ================                ============

┌─ CPU Monitor ─────┐              ┌─ Event Processor ─┐           ┌─ Métricas ─┐
│ • Uso de CPU      │              │ • Validate events │           │ • Dashboard │
│ • Multi-core      │◄────────────►│ • Update metrics  │◄─────────►│ • Logs      │
│ • Every 200ms     │              │ • Route alerts    │           │ • Alerts    │
└───────────────────┘              └───────────────────┘           └─────────────┘
                                            ▲
┌─ Memory Monitor ──┐                       │
│ • RAM usage       │                       │
│ • Buffer cache    │◄──────────────────────┤
│ • Every 500ms     │                       │
└───────────────────┘                       ▼
                                   ┌─ Alert Processor ─┐
┌─ Network Monitor ─┐              │ • Critical alerts │
│ • Throughput      │              │ • Notifications   │
│ • Interfaces      │◄────────────►│ • Escalation      │
│ • Every 300ms     │              └───────────────────┘
└───────────────────┘                       ▲
                                            │
┌─ Error Generator ─┐                       │
│ • Simulate errors │                       │
│ • Error levels    │◄──────────────────────┘
│ • Every 2s        │
└───────────────────┘
```

---

## 📡 Tipos de Channels Utilizados

### 1. **Event Channel** (Buffered: 100)
```go
eventos chan Evento
```
- **Propósito**: Canal principal para todos los eventos del sistema
- **Buffer**: 100 elementos para manejar ráfagas de eventos
- **Productores**: 4 generadores (CPU, Memoria, Red, Errores)
- **Consumidores**: 1 procesador principal

### 2. **Alert Channel** (Buffered: 50)
```go
alertas chan Evento
```
- **Propósito**: Canal dedicado para alertas críticas
- **Buffer**: 50 elementos para no perder alertas importantes
- **Productores**: Event Processor (cuando detecta condiciones críticas)
- **Consumidores**: Alert Processor

### 3. **Statistics Channel** (Buffered: 10)
```go
estadisticas chan map[string]interface{}
```
- **Propósito**: Canal para enviar estadísticas agregadas
- **Buffer**: 10 elementos para estadísticas periódicas
- **Productores**: Statistics Generator (cada 3 segundos)
- **Consumidores**: Statistics Processor

### 4. **Quit Channel** (Unbuffered)
```go
quit chan bool
```
- **Propósito**: Señal de shutdown para terminación elegante
- **Buffer**: Sin buffer para señalización inmediata
- **Uso**: Context.WithCancel() para propagación de cancelación

---

## 🔄 Patrones de Concurrencia Implementados

### 1. **Fan-Out Pattern**
```
Eventos → [CPU|Memory|Network|Error] → Event Channel
```
- Múltiples generadores producen eventos independientemente
- Un solo canal central recibe todos los eventos

### 2. **Pipeline Pattern**
```
Generate → Process → Filter → Alert → Display
```
- Eventos fluyen a través de etapas de procesamiento
- Cada etapa transforma o filtra datos

### 3. **Worker Pool Pattern**
```
Events → Event Processor → [Alert|Stats] Workers
```
- Procesador principal distribuye trabajo a workers especializados

### 4. **Fan-In Pattern**
```
[Alerts|Stats|Logs] → Display Aggregator → Console
```
- Múltiples fuentes de output se combinan para mostrar

---

## 📊 Tipos de Eventos

### EventoCPU
```go
{
    Timestamp: time.Now(),
    Tipo:      EventoCPU,
    Servicio:  "sistema",
    Valor:     85.3,  // Porcentaje de uso
    Metadata: {
        "unidad": "porcentaje",
        "core":   "core-2"
    }
}
```

### EventoMemoria
```go
{
    Timestamp: time.Now(),
    Tipo:      EventoMemoria,
    Servicio:  "sistema", 
    Valor:     4096.5,  // MB utilizados
    Metadata: {
        "unidad": "MB",
        "tipo":   "RAM"
    }
}
```

### EventoRed
```go
{
    Timestamp: time.Now(),
    Tipo:      EventoRed,
    Servicio:  "networking",
    Valor:     756.2,  // Mbps throughput
    Metadata: {
        "unidad":    "Mbps",
        "interface": "eth0"
    }
}
```

### EventoError
```go
{
    Timestamp: time.Now(),
    Tipo:      EventoError,
    Servicio:  "servicio-3",
    Valor:     8.5,  // Severidad (0-10)
    Metadata: {
        "codigo": "ERR-404",
        "nivel":  "CRITICAL"
    }
}
```

---

## 🚨 Sistema de Alertas

### Condiciones de Alerta

| **Tipo** | **Condición** | **Nivel** | **Acción** |
|----------|---------------|-----------|------------|
| CPU      | > 90%         | 🔴 CRÍTICO | Alerta inmediata |
| Memoria  | > 7GB         | 🟡 WARNING | Notificación |
| Red      | < 10 Mbps     | 🟡 WARNING | Log + notificación |
| Error    | Severidad > 7 | 🔴 CRÍTICO | Escalación |

### Procesamiento de Alertas
```go
func procesadorAlertas(alertas <-chan Evento, ctx context.Context) {
    for {
        select {
        case alerta := <-alertas:
            // Determinar nivel de severidad
            // Enviar notificación
            // Log del evento
        case <-ctx.Done():
            return
        }
    }
}
```

---

## 📈 Sistema de Métricas

### Métricas Recolectadas
- **Total de eventos**: Contador atómico global
- **Eventos por tipo**: Mapa con contadores por tipo
- **Tasa de errores**: Porcentaje de eventos de error
- **Último evento**: Timestamp del evento más reciente
- **Latencia promedio**: Tiempo de procesamiento

### Agregación en Tiempo Real
```go
type Metricas struct {
    eventosTotal     int64                    // atomic
    eventosPorTipo   map[TipoEvento]int64     // mutex protected
    ultimoEvento     time.Time                // mutex protected
    eventosError     int64                    // atomic
    promedioLatencia float64                  // atomic
    mu               sync.RWMutex
}
```

---

## 🛡️ Manejo de Cancelación

### Context Pattern
```go
ctx, cancel := context.WithCancel(context.Background())

// Todas las goroutines respetan el context
func generador(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // Cleanup y terminación
            return
        case <-ticker.C:
            // Trabajo normal
        }
    }
}
```

### Shutdown Elegante
1. **Señal de cancelación**: `cancel()` en context principal
2. **Propagación**: Todas las goroutines reciben `ctx.Done()`
3. **Cleanup**: Cada componente hace su limpieza
4. **Cierre de channels**: Cierre ordenado para evitar panics
5. **Estadísticas finales**: Reporte de métricas antes de salir

---

## ⚡ Optimizaciones de Performance

### Buffer Sizing Strategy
```go
eventos := make(chan Evento, 100)        // Alto throughput
alertas := make(chan Evento, 50)         // Crítico, buffer medio  
estadisticas := make(chan Stats, 10)     // Baja frecuencia
quit := make(chan bool)                  // Unbuffered, señalización
```

### Memory Management
- **Atomic operations** para contadores de alta frecuencia
- **RWMutex** para proteger estructuras compartidas
- **Channel buffering** para reducir bloqueos
- **Context cancellation** para evitar goroutine leaks

### CPU Efficiency
- **Select statements** para multiplexación eficiente
- **Non-blocking sends** con default case cuando apropiado
- **Batch processing** en el agregador de estadísticas
- **Ticker management** con cleanup apropiado

---

## 🧪 Testing Strategy

### Test Scenarios
1. **High Load**: 1000+ eventos/segundo
2. **Error Conditions**: Simulación de fallos
3. **Graceful Shutdown**: Terminación limpia
4. **Memory Leaks**: Verificación con race detector
5. **Deadlock Detection**: Timeout en operaciones

### Performance Benchmarks
```bash
go test -bench=. -race -memprofile=mem.prof
go tool pprof mem.prof
```

---

## 🚀 Ejecución del Proyecto

### Comandos Disponibles
```bash
# Ejecutar sistema completo
go run proyecto_monitoreo.go

# Con race detector
go run -race proyecto_monitoreo.go

# Con profiling
go run proyecto_monitoreo.go -cpuprofile=cpu.prof
```

### Output Esperado
```
📡 PROYECTO: Sistema de Monitoreo con Channels
==============================================
🚀 Iniciando Sistema de Monitoreo
📊 Procesador de eventos iniciado
🚨 Procesador de alertas iniciado
📈 Procesador de estadísticas iniciado
📈 [CPU] 45.67 porcentaje
📈 [MEMORIA] 3456.78 MB
🔴 ALERTA [CPU]: CPU crítico - 92.34
==================================================
📊 ESTADÍSTICAS DEL SISTEMA
==================================================
Total eventos: 247
Eventos error: 12
Tasa de error: 4.86%
Último evento: 15:42:33
==================================================
```

---

## 🎯 Objetivos de Aprendizaje

Al completar este proyecto, habrás dominado:

✅ **Channel Fundamentals**
- Buffered vs unbuffered channels
- Channel directions (`<-chan`, `chan<-`)
- Channel closing and range operations

✅ **Concurrency Patterns**
- Fan-Out/Fan-In
- Pipeline processing
- Worker pools
- Producer-consumer

✅ **Advanced Techniques**
- Select multiplexing
- Non-blocking operations
- Context cancellation
- Graceful shutdown

✅ **Real-World Applications**
- Event-driven architecture
- Real-time monitoring
- Alert systems
- Performance optimization

---

**🎉 Este proyecto demuestra cómo los channels son la herramienta fundamental para crear sistemas concurrentes robustos y escalables en Go. ¡Domina estos patrones y estarás listo para construir aplicaciones de nivel empresarial!**
