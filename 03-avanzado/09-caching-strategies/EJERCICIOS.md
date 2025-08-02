# 🧪 Ejercicios: Caching Strategies

## 📋 Instrucciones Generales

- ⏱️ **Tiempo estimado**: 6-8 horas de trabajo intensivo
- 🎯 **Objetivo**: Dominar implementación de sistemas de caché robustos
- 🧪 **Metodología**: TDD con tests de performance y concurrencia
- 📊 **Evaluación**: Funcionalidad + Performance + Código limpio

## 🎯 Ejercicio 1: LRU Cache Thread-Safe (⭐⭐)

### 📝 Descripción
Implementa un cache LRU (Least Recently Used) completamente thread-safe con soporte para TTL y métricas básicas.

### 🎯 Objetivos
- Cache LRU con tamaño máximo configurable
- Thread-safety completo (múltiples readers/writers)
- TTL (Time To Live) por item
- Métricas básicas (hits, misses, evictions)
- Limpieza automática de items expirados

### 🛠️ Implementación

```go
package cache

import (
    "container/list"
    "sync"
    "time"
)

// LRUCache implementa un cache LRU thread-safe
type LRUCache struct {
    // TODO: Implementar estructura
    // Sugerencia: usar map[string]*list.Element + doubly linked list
    // mutex para thread-safety
    // ticker para cleanup automático
}

type CacheItem struct {
    Key       string
    Value     interface{}
    ExpiresAt time.Time
}

type CacheStats struct {
    Hits      int64
    Misses    int64
    Size      int64
    Evictions int64
    HitRatio  float64
}

func NewLRUCache(maxSize int) *LRUCache {
    // TODO: Implementar constructor
    // - Inicializar estructuras de datos
    // - Iniciar goroutine para cleanup
    // - Configurar ticker cada 1 minuto
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    // TODO: Implementar
    // 1. Lock para lectura/escritura
    // 2. Verificar si existe
    // 3. Verificar TTL
    // 4. Mover al frente (LRU)
    // 5. Actualizar métricas
}

func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
    // TODO: Implementar
    // 1. Lock para escritura
    // 2. Si existe, actualizar y mover al frente
    // 3. Si no existe, agregar
    // 4. Verificar si excede maxSize
    // 5. Evict si es necesario
}

func (c *LRUCache) Delete(key string) bool {
    // TODO: Implementar eliminación
}

func (c *LRUCache) Stats() CacheStats {
    // TODO: Implementar estadísticas
}

func (c *LRUCache) Clear() {
    // TODO: Implementar limpieza completa
}

func (c *LRUCache) cleanup() {
    // TODO: Implementar limpieza automática de expirados
    // Usar ticker cada minuto
}
```

### ✅ Tests Requeridos

```go
func TestLRUCacheBasicOperations(t *testing.T) {
    // Test set/get básicos
}

func TestLRUCacheEviction(t *testing.T) {
    // Test que se evict correctamente cuando se llena
}

func TestLRUCacheTTL(t *testing.T) {
    // Test que los items expiran correctamente
}

func TestLRUCacheConcurrency(t *testing.T) {
    // Test con múltiples goroutines
    // 100 goroutines, 1000 ops cada una
}

func BenchmarkLRUCache(b *testing.B) {
    // Benchmark de operaciones
}
```

### 🎯 Criterios de Éxito
- ✅ No race conditions (usar `go test -race`)
- ✅ Hit ratio > 90% en workload típico
- ✅ Limpieza automática funciona
- ✅ Eviction LRU correcto
- ✅ Performance: > 1M ops/sec en operaciones básicas

---

## 🎯 Ejercicio 2: Multi-Level Cache (⭐⭐⭐)

### 📝 Descripción
Implementa un sistema de cache multi-nivel (L1: In-Memory, L2: Redis, L3: Database) con promotion automática y cache warming.

### 🎯 Objetivos
- Cache L1 (in-memory, rápido, pequeño)
- Cache L2 (Redis, medio, más grande)
- Cache L3 (Database/Mock, lento, completo)
- Promotion automática entre niveles
- Cache warming inteligente
- Métricas por nivel

### 🛠️ Implementación

```go
package multilevel

import (
    "context"
    "time"
)

// CacheLevel representa un nivel de cache
type CacheLevel interface {
    Get(ctx context.Context, key string) (interface{}, bool, error)
    Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
    Delete(ctx context.Context, key string) error
    Stats() LevelStats
}

type LevelStats struct {
    Level     string
    Hits      int64
    Misses    int64
    Size      int64
    HitRatio  float64
    AvgLatency time.Duration
}

// MultiLevelCache implementa cache jerárquico
type MultiLevelCache struct {
    // TODO: Implementar
    // - Slice de CacheLevel
    // - Configuración de TTL por nivel
    // - Channel para promotion async
    // - Métricas globales
}

type MultiLevelConfig struct {
    L1TTL           time.Duration
    L2TTL           time.Duration
    L3TTL           time.Duration
    PromotionAsync  bool
    WarmingInterval time.Duration
}

func NewMultiLevelCache(levels []CacheLevel, config MultiLevelConfig) *MultiLevelCache {
    // TODO: Implementar constructor
    // - Iniciar goroutines para promotion
    // - Configurar warming automático
}

func (m *MultiLevelCache) Get(ctx context.Context, key string) (interface{}, error) {
    // TODO: Implementar búsqueda en cascada
    // 1. Buscar en L1
    // 2. Si no está, buscar en L2, promover a L1
    // 3. Si no está, buscar en L3, promover a L2 y L1
    // 4. Registrar métricas por nivel
}

func (m *MultiLevelCache) Set(ctx context.Context, key string, value interface{}) error {
    // TODO: Implementar escritura en todos los niveles
    // Con TTLs diferentes para cada nivel
}

func (m *MultiLevelCache) Warm(ctx context.Context, keys []string) error {
    // TODO: Implementar cache warming
    // Cargar desde L3 hacia L1/L2
}

func (m *MultiLevelCache) GetStats() MultiLevelStats {
    // TODO: Consolidar estadísticas de todos los niveles
}

// Implementar niveles específicos
type InMemoryLevel struct {
    cache *LRUCache // Del ejercicio anterior
}

type RedisLevel struct {
    client RedisClient // Mock o real
}

type DatabaseLevel struct {
    db DataSource // Mock para tests
}
```

### ✅ Tests Requeridos

```go
func TestMultiLevelPromotion(t *testing.T) {
    // Test que L3->L2->L1 promotion funciona
}

func TestMultiLevelCacheWarming(t *testing.T) {
    // Test cache warming desde database
}

func TestMultiLevelPerformance(t *testing.T) {
    // Test que L1 es más rápido que L2 que L3
}

func TestMultiLevelFailover(t *testing.T) {
    // Test behavior cuando un nivel falla
}
```

### 🎯 Criterios de Éxito
- ✅ L1 latency < 1ms, L2 < 10ms, L3 < 100ms
- ✅ Promotion automática funciona
- ✅ Cache warming reduce miss ratio inicial
- ✅ Graceful degradation cuando un nivel falla

---

## 🎯 Ejercicio 3: Distributed Cache con Consistent Hashing (⭐⭐⭐⭐)

### 📝 Descripción
Implementa un cache distribuido usando consistent hashing para distribución de keys, con replicación y failure handling.

### 🎯 Objetivos
- Consistent hashing para distribución
- Virtual nodes para balanceeo
- Replicación configurable (N réplicas)
- Failure detection y recovery
- Rehashing automático cuando nodos caen/se agregan

### 🛠️ Implementación

```go
package distributed

import (
    "context"
    "hash/fnv"
    "sort"
    "sync"
    "time"
)

// ConsistentHashRing implementa consistent hashing
type ConsistentHashRing struct {
    // TODO: Implementar
    // - map[uint32]string para ring
    // - []uint32 para sorted keys
    // - int para virtual nodes
    // - sync.RWMutex para concurrency
}

// DistributedCache implementa cache distribuido
type DistributedCache struct {
    ring        *ConsistentHashRing
    nodes       map[string]CacheNode
    replicas    int
    failureDetector *FailureDetector
    
    // TODO: Agregar más campos necesarios
}

type CacheNode interface {
    Get(ctx context.Context, key string) (interface{}, error)
    Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
    Delete(ctx context.Context, key string) error
    Health() error
}

type FailureDetector struct {
    // TODO: Implementar detector de fallos
    // - Health checks periódicos
    // - Circuit breaker por nodo
    // - Callbacks para nodo down/up
}

func NewDistributedCache(nodes []CacheNode, replicas int) *DistributedCache {
    // TODO: Implementar constructor
    // - Crear ring con virtual nodes
    // - Iniciar failure detector
    // - Setup replication
}

func (d *DistributedCache) Get(ctx context.Context, key string) (interface{}, error) {
    // TODO: Implementar
    // 1. Hash key para encontrar nodos primario y réplicas
    // 2. Intentar nodo primario
    // 3. Si falla, intentar réplicas
    // 4. Si encuentra en réplica, reparar primario async
}

func (d *DistributedCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    // TODO: Implementar
    // 1. Encontrar nodos para la key
    // 2. Escribir en primario y réplicas
    // 3. Manejar fallos parciales
}

func (d *DistributedCache) AddNode(node CacheNode) error {
    // TODO: Implementar adición dinámica de nodos
    // - Agregar al ring
    // - Trigger rehashing si es necesario
}

func (d *DistributedCache) RemoveNode(nodeID string) error {
    // TODO: Implementar remoción de nodos
    // - Remover del ring
    // - Redirigir keys afectadas
}

func (r *ConsistentHashRing) GetNodes(key string, count int) []string {
    // TODO: Implementar obtención de N nodos para una key
    // Usando consistent hashing
}

func (r *ConsistentHashRing) AddNode(nodeID string, virtualNodes int) {
    // TODO: Implementar adición de nodo al ring
}

func (r *ConsistentHashRing) RemoveNode(nodeID string) {
    // TODO: Implementar remoción de nodo del ring
}
```

### ✅ Tests Requeridos

```go
func TestConsistentHashing(t *testing.T) {
    // Test distribución uniforme de keys
}

func TestNodeFailure(t *testing.T) {
    // Test behavior cuando un nodo falla
}

func TestNodeAddition(t *testing.T) {
    // Test agregar nodo dinámicamente
}

func TestReplication(t *testing.T) {
    // Test que réplicas funcionan correctamente
}

func TestHashingDistribution(t *testing.T) {
    // Test que distribución es uniforme
    // 10000 keys, variación < 20% entre nodos
}
```

### 🎯 Criterios de Éxito
- ✅ Distribución uniforme (< 20% variación)
- ✅ Automatic failover en < 1s
- ✅ Zero downtime cuando se agregan/remueven nodos
- ✅ Replicación mantiene consistencia

---

## 🎯 Ejercicio 4: Smart Cache con ML Predictions (⭐⭐⭐⭐⭐)

### 📝 Descripción
Implementa un cache "inteligente" que usa machine learning simple para predecir qué datos precargar y optimizar el tamaño dinámicamente.

### 🎯 Objetivos
- Predicción de hot keys usando análisis de patrones
- Auto-scaling del cache basado en métricas
- Cache warming predictivo
- Anomaly detection para patrones inusuales
- Dashboard en tiempo real

### 🛠️ Implementación

```go
package smart

import (
    "context"
    "time"
)

// SmartCache implementa cache con ML predictions
type SmartCache struct {
    cache         Cache
    predictor     *AccessPredictor
    optimizer     *SizeOptimizer
    warmer        *PredictiveWarmer
    monitor       *MetricsMonitor
    dashboard     *Dashboard
}

// AccessPredictor predice patrones de acceso
type AccessPredictor struct {
    // TODO: Implementar
    // - Historia de accesos por key
    // - Análisis de frecuencia y seasonality
    // - Predicción basada en tiempo/día
}

type AccessPattern struct {
    Key           string
    Frequency     float64    // Accesos por hora promedio
    Seasonality   []float64  // Patrón por hora del día
    Trend         float64    // Tendencia (creciente/decreciente)
    LastAccess    time.Time
    PredictedNext time.Time
}

func NewSmartCache(baseCache Cache) *SmartCache {
    // TODO: Implementar constructor
    // - Iniciar predictor con ventana de análisis
    // - Configurar optimizer para auto-scaling
    // - Setup warmer para preloading
    // - Iniciar dashboard servidor
}

func (s *SmartCache) Get(key string) (interface{}, error) {
    // TODO: Implementar con logging para ML
    // 1. Get normal del cache
    // 2. Log acceso para predictor
    // 3. Trigger predicción si es necesario
}

func (s *SmartCache) Set(key string, value interface{}, ttl time.Duration) error {
    // TODO: Implementar con optimización dinámica
}

// AccessPredictor methods
func (a *AccessPredictor) RecordAccess(key string, timestamp time.Time) {
    // TODO: Registrar acceso para análisis
}

func (a *AccessPredictor) PredictHotKeys(window time.Duration) []string {
    // TODO: Predecir keys que serán accedidas pronto
    // Usar análisis de frecuencia y seasonality
}

func (a *AccessPredictor) AnalyzePatterns() map[string]AccessPattern {
    // TODO: Analizar patrones históricos
    // - Calcular frecuencia promedio
    // - Detectar patrones horarios
    // - Calcular tendencias
}

// SizeOptimizer auto-ajusta tamaño del cache
type SizeOptimizer struct {
    targetHitRatio    float64
    minSize          int
    maxSize          int
    adjustmentFactor float64
}

func (s *SizeOptimizer) OptimalSize(currentStats CacheStats) int {
    // TODO: Calcular tamaño óptimo basado en hit ratio
    // Si hit ratio < target -> increase size
    // Si hit ratio muy alto -> puede decrease size
}

// PredictiveWarmer precarga datos predicted
type PredictiveWarmer struct {
    cache      Cache
    dataSource DataSource
    predictor  *AccessPredictor
}

func (p *PredictiveWarmer) WarmPredictedKeys(ctx context.Context) error {
    // TODO: Precargar keys que se predice serán accedidas
    // 1. Get predicted hot keys
    // 2. Check cuáles no están en cache
    // 3. Fetch desde data source
    // 4. Store en cache con TTL apropiado
}

// Dashboard provides real-time metrics
type Dashboard struct {
    cache   *SmartCache
    server  *http.Server
    metrics *MetricsCollector
}

func (d *Dashboard) Start(port int) error {
    // TODO: Implementar HTTP server para dashboard
    // Endpoints:
    // GET /stats - estadísticas actuales
    // GET /predictions - predicciones actuales
    // GET /hot-keys - top hot keys
    // GET /patterns - patrones detectados
}

type DashboardData struct {
    CurrentStats    CacheStats                  `json:"current_stats"`
    Predictions     []string                    `json:"predicted_hot_keys"`
    HotKeys         []KeyStats                  `json:"hot_keys"`
    Patterns        map[string]AccessPattern    `json:"patterns"`
    OptimalSize     int                         `json:"optimal_size"`
    LastUpdate      time.Time                   `json:"last_update"`
}
```

### ✅ Tests Requeridos

```go
func TestAccessPrediction(t *testing.T) {
    // Test que predicciones son razonablemente precisas
    // Simular patrones conocidos y verificar predicción
}

func TestSizeOptimization(t *testing.T) {
    // Test que auto-scaling funciona
}

func TestPredictiveWarming(t *testing.T) {
    // Test que warming predictivo mejora hit ratio
}

func TestAnomalyDetection(t *testing.T) {
    // Test detección de patrones anómalos
}

func TestDashboardAPI(t *testing.T) {
    // Test endpoints del dashboard
}
```

### 🎯 Criterios de Éxito
- ✅ Predicciones 70%+ precision en patrones regulares
- ✅ Auto-scaling mantiene target hit ratio ±5%
- ✅ Predictive warming mejora hit ratio inicial 20%+
- ✅ Dashboard responsive y actualizado cada 5s

---

## 🎯 Ejercicio 5: Cache Performance Benchmark Suite (⭐⭐⭐)

### 📝 Descripción
Crea una suite completa de benchmarks para evaluar performance de diferentes estrategias de cache bajo varios escenarios.

### 🎯 Objetivos
- Benchmarks comparativos entre implementaciones
- Tests de carga con diferentes patrones
- Memory profiling y leak detection
- Performance regression tests
- Automated performance CI

### 🛠️ Implementación

```go
package benchmarks

import (
    "sync"
    "testing"
    "time"
)

// BenchmarkSuite ejecuta tests comprehensivos
type BenchmarkSuite struct {
    caches map[string]Cache
    scenarios []BenchmarkScenario
}

type BenchmarkScenario struct {
    Name           string
    Duration       time.Duration
    Concurrency    int
    ReadWriteRatio float64 // 0.8 = 80% reads, 20% writes
    KeyDistribution KeyDistribution
    DataSize       int
}

type KeyDistribution interface {
    NextKey() string
}

// Implementar diferentes distribuciones
type UniformDistribution struct {
    keyCount int
    current  int
}

type ZipfianDistribution struct {
    // TODO: Implementar distribución Zipfian (realista)
    // 80% de accesos van a 20% de keys
}

type HotspotDistribution struct {
    // TODO: Implementar distribución con hotspots
    // Simula patrones reales donde algunas keys son muy populares
}

func NewBenchmarkSuite() *BenchmarkSuite {
    return &BenchmarkSuite{
        caches: make(map[string]Cache),
        scenarios: []BenchmarkScenario{
            {
                Name:           "Uniform_80Read_20Write",
                Duration:       time.Minute,
                Concurrency:    100,
                ReadWriteRatio: 0.8,
                KeyDistribution: &UniformDistribution{keyCount: 10000},
                DataSize:       1024, // 1KB values
            },
            {
                Name:           "Zipfian_90Read_10Write",
                Duration:       time.Minute,
                Concurrency:    100,
                ReadWriteRatio: 0.9,
                KeyDistribution: &ZipfianDistribution{},
                DataSize:       4096, // 4KB values
            },
            // TODO: Agregar más scenarios
        },
    }
}

func (b *BenchmarkSuite) AddCache(name string, cache Cache) {
    b.caches[name] = cache
}

func (b *BenchmarkSuite) RunAll(t *testing.B) {
    for cacheName, cache := range b.caches {
        for _, scenario := range b.scenarios {
            name := fmt.Sprintf("%s_%s", cacheName, scenario.Name)
            t.Run(name, func(b *testing.B) {
                b.RunParallel(func(pb *testing.PB) {
                    b.runScenario(cache, scenario, pb)
                })
            })
        }
    }
}

func (b *BenchmarkSuite) runScenario(cache Cache, scenario BenchmarkScenario, pb *testing.PB) {
    // TODO: Implementar ejecución de scenario
    // 1. Setup workload según distribución
    // 2. Ejecutar reads/writes según ratio
    // 3. Medir latencias, throughput
    // 4. Verificar correctness
}

// Benchmark específicos
func BenchmarkCacheGet(b *testing.B) {
    cache := NewLRUCache(10000)
    
    // Prepopulate
    for i := 0; i < 1000; i++ {
        cache.Set(fmt.Sprintf("key-%d", i), i, time.Hour)
    }
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            key := fmt.Sprintf("key-%d", rand.Intn(1000))
            cache.Get(key)
        }
    })
}

func BenchmarkCacheSet(b *testing.B) {
    cache := NewLRUCache(10000)
    
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        i := 0
        for pb.Next() {
            key := fmt.Sprintf("key-%d", i)
            cache.Set(key, i, time.Hour)
            i++
        }
    })
}

func BenchmarkCacheMixed(b *testing.B) {
    // TODO: Benchmark con mix de operations
    // 80% reads, 20% writes
}

// Memory benchmarks
func BenchmarkCacheMemoryUsage(b *testing.B) {
    // TODO: Benchmark memory usage per item
    // Usar b.ReportAllocs()
}

// Latency distribution benchmark
func BenchmarkCacheLatencyDistribution(b *testing.B) {
    // TODO: Medir P50, P95, P99 latencias
    // Usar histogram para distribución
}

// Contention benchmark
func BenchmarkCacheContention(b *testing.B) {
    // TODO: Test performance con high contention
    // Múltiples goroutines accediendo same keys
}
```

### ✅ Performance Targets

```go
// Performance regression tests
func TestPerformanceRegression(t *testing.T) {
    benchmarks := []struct {
        name string
        target time.Duration
        test func() time.Duration
    }{
        {
            name: "Get_Latency_P99",
            target: 1 * time.Millisecond,
            test: func() time.Duration {
                // TODO: Medir P99 latency
            },
        },
        {
            name: "Set_Throughput",
            target: 1000000, // 1M ops/sec
            test: func() time.Duration {
                // TODO: Medir throughput
            },
        },
    }
    
    for _, bench := range benchmarks {
        result := bench.test()
        if result > bench.target {
            t.Errorf("%s: %v > target %v", bench.name, result, bench.target)
        }
    }
}
```

### 🎯 Criterios de Éxito
- ✅ LRU Cache: > 1M get/sec, < 1ms P99 latency
- ✅ Multi-level: L1 < 1ms, L2 < 10ms, L3 < 100ms
- ✅ Distributed: < 5ms P99 con replicación
- ✅ Zero memory leaks en stress tests 24h
- ✅ Automated benchmarks en CI

---

## 🎯 Ejercicio 6: Production-Ready Cache System (⭐⭐⭐⭐⭐)

### 📝 Descripción
Integra todos los componentes anteriores en un sistema de cache production-ready con monitoring, alerting, y operational tools.

### 🎯 Objetivos
- Sistema completo de cache para producción
- Health checks y monitoring comprehensive
- Alerting inteligente basado en métricas
- Operational tools (CLI, dashboard, debugging)
- Configuration management y feature flags
- Circuit breakers y graceful degradation

### 🛠️ Implementación

```go
package production

import (
    "context"
    "time"
)

// ProductionCacheSystem es el sistema completo
type ProductionCacheSystem struct {
    cache           *SmartCache          // Del ejercicio 4
    distributed     *DistributedCache    // Del ejercicio 3
    monitoring      *MonitoringSystem
    alerting        *AlertingSystem
    healthChecker   *HealthCheckService
    configManager   *ConfigurationManager
    circuitBreaker  *CircuitBreaker
    metricsExporter *MetricsExporter
    dashboard       *OperationalDashboard
    cli             *CLITools
}

type CacheSystemConfig struct {
    CacheType      string            `yaml:"cache_type"`      // "local", "distributed", "hybrid"
    MaxSize        int               `yaml:"max_size"`
    TTLDefault     time.Duration     `yaml:"ttl_default"`
    ReplicationFactor int            `yaml:"replication_factor"`
    MonitoringConfig  MonitoringConfig `yaml:"monitoring"`
    AlertingConfig    AlertingConfig   `yaml:"alerting"`
    CircuitBreakerConfig CBConfig     `yaml:"circuit_breaker"`
}

func NewProductionCacheSystem(config CacheSystemConfig) (*ProductionCacheSystem, error) {
    // TODO: Implementar constructor completo
    // - Validar configuración
    // - Inicializar todos los componentes
    // - Setup monitoring y alerting
    // - Start health checks
    // - Initialize circuit breakers
}

// MonitoringSystem recolecta métricas comprehensive
type MonitoringSystem struct {
    collector    *MetricsCollector
    exporter     MetricsExporter
    healthChecks map[string]HealthCheck
}

type MetricsCollector struct {
    // TODO: Implementar collector con múltiples backends
    // - Prometheus
    // - StatsD
    // - Custom metrics
}

func (m *MonitoringSystem) CollectMetrics() CacheMetrics {
    // TODO: Recolectar métricas completas
    return CacheMetrics{
        Performance: PerformanceMetrics{
            Throughput:    m.getThroughput(),
            Latency:       m.getLatencyDistribution(),
            ErrorRate:     m.getErrorRate(),
        },
        Resource: ResourceMetrics{
            MemoryUsage:   m.getMemoryUsage(),
            CPUUsage:      m.getCPUUsage(),
            NetworkIO:     m.getNetworkIO(),
        },
        Business: BusinessMetrics{
            HitRatio:      m.getHitRatio(),
            KeyDistribution: m.getKeyDistribution(),
            UserSessions:  m.getActiveSessions(),
        },
    }
}

// AlertingSystem maneja alertas inteligentes
type AlertingSystem struct {
    rules     []AlertRule
    channels  []AlertChannel
    history   AlertHistory
    ml        *AnomalyDetector
}

type AlertRule struct {
    Name        string
    Condition   string // "hit_ratio < 0.8"
    Severity    AlertSeverity
    Frequency   time.Duration
    Suppression time.Duration
}

type AlertSeverity int

const (
    SeverityInfo AlertSeverity = iota
    SeverityWarning
    SeverityError
    SeverityCritical
)

func (a *AlertingSystem) EvaluateRules(metrics CacheMetrics) []Alert {
    // TODO: Evaluar reglas contra métricas actuales
    // - Usar expression evaluator para conditions
    // - Aplicar suppression logic
    // - Detectar anomalías con ML
}

// HealthCheckService monitorea salud del sistema
type HealthCheckService struct {
    checks   map[string]HealthCheck
    results  map[string]HealthResult
    interval time.Duration
}

type HealthCheck interface {
    Name() string
    Check(ctx context.Context) HealthResult
}

type HealthResult struct {
    Status   HealthStatus
    Message  string
    Duration time.Duration
    Timestamp time.Time
}

type HealthStatus int

const (
    HealthStatusHealthy HealthStatus = iota
    HealthStatusDegraded
    HealthStatusUnhealthy
)

// Implementar health checks específicos
type CacheConnectivityCheck struct {
    cache Cache
}

func (c *CacheConnectivityCheck) Check(ctx context.Context) HealthResult {
    // TODO: Verificar conectividad del cache
    // Set/Get test key con timeout
}

type CachePerformanceCheck struct {
    cache     Cache
    threshold time.Duration
}

func (c *CachePerformanceCheck) Check(ctx context.Context) HealthResult {
    // TODO: Verificar performance está dentro de thresholds
}

// CircuitBreaker protege contra fallos en cascada
type CircuitBreaker struct {
    state        CBState
    failureCount int
    lastFailure  time.Time
    config       CBConfig
}

type CBState int

const (
    CBStateClosed CBState = iota
    CBStateOpen
    CBStateHalfOpen
)

type CBConfig struct {
    FailureThreshold int           `yaml:"failure_threshold"`
    RecoveryTimeout  time.Duration `yaml:"recovery_timeout"`
    SuccessThreshold int           `yaml:"success_threshold"`
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    // TODO: Implementar circuit breaker logic
    // - Check current state
    // - Execute function if closed/half-open
    // - Update state based on result
}

// OperationalDashboard provides ops interface
type OperationalDashboard struct {
    server  *http.Server
    system  *ProductionCacheSystem
    wsConns map[string]*websocket.Conn
}

func (d *OperationalDashboard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // TODO: Implementar endpoints operacionales
    // GET /health - health status
    // GET /metrics - current metrics
    // GET /alerts - active alerts
    // POST /admin/cache/clear - operational actions
    // GET /dashboard - real-time dashboard
    // WebSocket /ws - real-time updates
}

// CLITools provides command line interface
type CLITools struct {
    system *ProductionCacheSystem
}

func (cli *CLITools) Execute(args []string) error {
    // TODO: Implementar CLI commands
    // cache-cli stats
    // cache-cli health
    // cache-cli clear [pattern]
    // cache-cli warm [keys...]
    // cache-cli config [get|set] [key] [value]
    // cache-cli debug [key]
}

// Configuration management
type ConfigurationManager struct {
    config     CacheSystemConfig
    watchers   []ConfigWatcher
    validator  ConfigValidator
}

func (c *ConfigurationManager) WatchConfig() {
    // TODO: Watch config file/environment changes
    // Notify watchers when config changes
    // Validate new config before applying
}

func (c *ConfigurationManager) UpdateConfig(newConfig CacheSystemConfig) error {
    // TODO: Hot reload configuration
    // - Validate new config
    // - Apply changes without downtime
    // - Notify components of changes
}
```

### ✅ Production Checklist

```go
func TestProductionReadiness(t *testing.T) {
    system := setupProductionSystem(t)
    
    // Health checks
    assert.True(t, system.IsHealthy())
    
    // Performance requirements
    assert.Less(t, system.GetP99Latency(), 10*time.Millisecond)
    assert.Greater(t, system.GetThroughput(), 100000) // 100k ops/sec
    assert.Greater(t, system.GetHitRatio(), 0.9)     // 90% hit ratio
    
    // Reliability requirements
    assert.NoError(t, system.TestFailover())
    assert.NoError(t, system.TestRecovery())
    assert.NoError(t, system.TestScaling())
    
    // Operational requirements
    assert.True(t, system.HasMetrics())
    assert.True(t, system.HasAlerting())
    assert.True(t, system.HasDashboard())
    assert.True(t, system.HasCLI())
    
    // Security requirements
    assert.True(t, system.HasAuthentication())
    assert.True(t, system.HasAuthorization())
    assert.True(t, system.HasEncryption())
}
```

### 🎯 Criterios de Éxito
- ✅ 99.9% uptime bajo carga normal
- ✅ < 10ms P99 latency end-to-end
- ✅ Auto-recovery de fallos en < 30s
- ✅ Zero-downtime config updates
- ✅ Comprehensive monitoring y alerting
- ✅ Production-ready ops tools

---

## 🏆 Evaluación Final

### 📊 Rubrica de Evaluación

| Criterio | Excelente (4) | Bueno (3) | Satisfactorio (2) | Insuficiente (1) |
|----------|---------------|-----------|-------------------|------------------|
| **Funcionalidad** | Todos los ejercicios completos y funcionando | 5/6 ejercicios completos | 4/6 ejercicios completos | < 4 ejercicios |
| **Performance** | Supera todos los targets | Cumple 90% targets | Cumple 70% targets | < 70% targets |
| **Código Limpio** | Código ejemplar, documentado | Código limpio, bien estructurado | Código funcional, algunos issues | Código difícil de leer |
| **Testing** | Coverage > 90%, benchmarks completos | Coverage > 80%, buenos tests | Coverage > 60%, tests básicos | Coverage < 60% |
| **Concurrencia** | Thread-safe, optimizado | Thread-safe, correcto | Mostly thread-safe | Race conditions |

### 🎯 Entregables

1. **📁 Código Fuente**: Todos los ejercicios implementados
2. **🧪 Test Suite**: Tests comprehensivos con coverage > 80%
3. **📊 Benchmarks**: Performance benchmarks con resultados
4. **📖 Documentación**: README con instrucciones y decisiones de diseño
5. **🎥 Demo**: Video/screenshots del sistema funcionando

### ⏰ Timeline Sugerido

- **Día 1-2**: Ejercicios 1-2 (LRU + Multi-level)
- **Día 3-4**: Ejercicios 3-4 (Distributed + Smart)
- **Día 5-6**: Ejercicios 5-6 (Benchmarks + Production)
- **Día 7**: Testing, documentación, demo

### 🚀 Bonus Points

- **🔒 Security**: Implementar encryption at rest/in transit
- **📈 Monitoring**: Integración con Grafana/Prometheus real
- **🌐 API Gateway**: Cache para API responses
- **🤖 Auto-tuning**: ML para parameter optimization
- **📦 Containerización**: Docker + K8s deployment

---

¡Demuestra tu maestría en caching strategies! 💪 Este proyecto será una excelente adición a tu portfolio. 🚀
