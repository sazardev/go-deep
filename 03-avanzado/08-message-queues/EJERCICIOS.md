# 🧪 Ejercicios Prácticos: Message Queues y Event Streaming

## 📋 Instrucciones Generales

Estos ejercicios están diseñados para que practiques los conceptos de message queues y event streaming en Go. Cada ejercicio incluye tests que debes hacer pasar.

### 🎯 Objetivos de Aprendizaje

- Implementar diferentes patrones de mensajería
- Manejar confiabilidad y fallos
- Optimizar performance y throughput
- Integrar con brokers reales
- Diseñar arquitecturas event-driven

---

## 🟢 Ejercicio 1: Queue Básico con Retry Logic

**Dificultad**: Principiante  
**Tiempo estimado**: 30 minutos

### 📝 Descripción

Implementa un message queue básico con mecanismo de retry automático y dead letter queue.

### 🎯 Requisitos

```go
// Implementa estas interfaces
type Queue interface {
    Produce(ctx context.Context, msg Message) error
    Consume(ctx context.Context, handler MessageHandler) error
    GetMetrics() QueueMetrics
    Close() error
}

type MessageHandler interface {
    Handle(ctx context.Context, msg Message) error
}

type Message struct {
    ID         string            `json:"id"`
    Topic      string            `json:"topic"`
    Payload    []byte            `json:"payload"`
    Headers    map[string]string `json:"headers"`
    Timestamp  time.Time         `json:"timestamp"`
    Retries    int               `json:"retries"`
    MaxRetries int               `json:"max_retries"`
}

type QueueMetrics struct {
    Produced    int64
    Consumed    int64
    Failed      int64
    InFlight    int64
    DLQMessages int64
}
```

### ✅ Criterios de Éxito

1. Mensajes se procesan en orden FIFO
2. Retry automático con backoff exponencial
3. Dead Letter Queue después de max retries
4. Métricas actualizadas correctamente
5. Graceful shutdown

### 🧪 Test Template

```go
func TestBasicQueue(t *testing.T) {
    queue := NewBasicQueue("test-queue", 100)
    defer queue.Close()
    
    // Test básico de produce/consume
    msg := Message{
        ID:         "test-1",
        Topic:      "test",
        Payload:    []byte(`{"data": "test"}`),
        MaxRetries: 3,
    }
    
    ctx := context.Background()
    err := queue.Produce(ctx, msg)
    assert.NoError(t, err)
    
    received := make(chan Message, 1)
    handler := &TestHandler{
        onHandle: func(ctx context.Context, msg Message) error {
            received <- msg
            return nil
        },
    }
    
    go queue.Consume(ctx, handler)
    
    select {
    case receivedMsg := <-received:
        assert.Equal(t, msg.ID, receivedMsg.ID)
    case <-time.After(time.Second):
        t.Fatal("Message not received")
    }
    
    metrics := queue.GetMetrics()
    assert.Equal(t, int64(1), metrics.Produced)
    assert.Equal(t, int64(1), metrics.Consumed)
}

func TestRetryMechanism(t *testing.T) {
    // TODO: Test que simule fallos y verifique retries
}

func TestDeadLetterQueue(t *testing.T) {
    // TODO: Test que verifique DLQ después de max retries
}
```

---

## 🟡 Ejercicio 2: Pub/Sub con Topic Routing

**Dificultad**: Intermedio  
**Tiempo estimado**: 45 minutos

### 📝 Descripción

Implementa un sistema pub/sub con routing basado en topics y wildcards.

### 🎯 Requisitos

```go
type PubSubBroker interface {
    Subscribe(topic string) (<-chan Message, error)
    SubscribeWithPattern(pattern string) (<-chan Message, error)
    Publish(topic string, msg Message) error
    Unsubscribe(topic string) error
    Close() error
}

// Patrones soportados:
// * = exactamente una palabra
// # = cero o más palabras
// Ejemplos:
// "user.*" matches "user.created", "user.updated"
// "order.#" matches "order.created", "order.payment.processed"
```

### ✅ Criterios de Éxito

1. Subscribers reciben solo mensajes de topics relevantes
2. Wildcards funcionan correctamente (* y #)
3. Múltiples subscribers por topic
4. No bloqueo si un subscriber es lento
5. Cleanup automático de subscribers inactivos

### 🧪 Test Template

```go
func TestTopicRouting(t *testing.T) {
    broker := NewPubSubBroker()
    defer broker.Close()
    
    // Subscribe to specific topic
    userCh, err := broker.Subscribe("user.created")
    assert.NoError(t, err)
    
    // Subscribe with wildcard
    orderCh, err := broker.SubscribeWithPattern("order.*")
    assert.NoError(t, err)
    
    // Publish messages
    msg1 := Message{ID: "1", Topic: "user.created", Payload: []byte("user")}
    msg2 := Message{ID: "2", Topic: "order.created", Payload: []byte("order")}
    
    broker.Publish("user.created", msg1)
    broker.Publish("order.created", msg2)
    
    // Verify routing
    select {
    case received := <-userCh:
        assert.Equal(t, "1", received.ID)
    case <-time.After(time.Second):
        t.Fatal("User message not received")
    }
    
    select {
    case received := <-orderCh:
        assert.Equal(t, "2", received.ID)
    case <-time.After(time.Second):
        t.Fatal("Order message not received")
    }
}
```

---

## 🔴 Ejercicio 3: Sistema de Pedidos E-commerce

**Dificultad**: Avanzado  
**Tiempo estimado**: 90 minutos

### 📝 Descripción

Implementa un sistema completo de procesamiento de pedidos usando message queues con orquestación de servicios.

### 🎯 Requisitos

```go
// Servicios que debes implementar
type OrderService interface {
    CreateOrder(ctx context.Context, order Order) error
    GetOrder(ctx context.Context, orderID string) (*Order, error)
}

type PaymentService interface {
    ProcessPayment(ctx context.Context, paymentReq PaymentRequest) error
}

type InventoryService interface {
    ReserveInventory(ctx context.Context, items []Item) error
    ReleaseInventory(ctx context.Context, items []Item) error
}

type ShippingService interface {
    CreateShipment(ctx context.Context, shipmentReq ShippingRequest) error
}

type OrderOrchestrator interface {
    ProcessOrder(ctx context.Context, orderID string) error
}

// Events que deben ser publicados
type OrderCreated struct {
    OrderID   string    `json:"order_id"`
    Items     []Item    `json:"items"`
    Total     float64   `json:"total"`
    Timestamp time.Time `json:"timestamp"`
}

type PaymentProcessed struct {
    OrderID   string `json:"order_id"`
    Success   bool   `json:"success"`
    Amount    float64 `json:"amount"`
}

type InventoryReserved struct {
    OrderID string `json:"order_id"`
    Success bool   `json:"success"`
    Items   []Item `json:"items"`
}
```

### 🎯 Flujo Esperado

1. **Order Created** → Payment Processing + Inventory Check (paralelo)
2. **Payment Success + Inventory Success** → Shipping
3. **Any Failure** → Compensation (rollback)

### ✅ Criterios de Éxito

1. Flujo completo funciona correctamente
2. Manejo de fallos con compensación
3. Procesamiento paralelo donde sea posible
4. Idempotencia en todos los servicios
5. Métricas de cada paso del proceso

### 🧪 Test Template

```go
func TestOrderWorkflow(t *testing.T) {
    // Setup
    broker := SetupTestBroker()
    defer broker.Close()
    
    orderService := NewOrderService(broker)
    paymentService := NewPaymentService(broker)
    inventoryService := NewInventoryService(broker)
    shippingService := NewShippingService(broker)
    orchestrator := NewOrderOrchestrator(broker, 
        paymentService, inventoryService, shippingService)
    
    // Create order
    order := Order{
        ID:         "order-123",
        CustomerID: "customer-456",
        Items: []Item{
            {ProductID: "product-1", Quantity: 2, Price: 10.00},
        },
        Total: 20.00,
    }
    
    ctx := context.Background()
    err := orderService.CreateOrder(ctx, order)
    assert.NoError(t, err)
    
    // Process workflow
    err = orchestrator.ProcessOrder(ctx, order.ID)
    assert.NoError(t, err)
    
    // Verify final state
    time.Sleep(time.Second) // Wait for async processing
    
    finalOrder, err := orderService.GetOrder(ctx, order.ID)
    assert.NoError(t, err)
    assert.Equal(t, "shipped", finalOrder.Status)
}

func TestOrderWorkflowWithPaymentFailure(t *testing.T) {
    // TODO: Test compensation cuando falla el pago
}

func TestOrderWorkflowWithInventoryFailure(t *testing.T) {
    // TODO: Test compensation cuando falla inventory
}
```

---

## ⚡ Ejercicio 4: High-Performance Message Broker

**Dificultad**: Experto  
**Tiempo estimado**: 120 minutos

### 📝 Descripción

Implementa un message broker de alta performance con batching, connection pooling y sharding.

### 🎯 Requisitos

```go
type HighPerformanceBroker interface {
    CreateTopic(name string, partitions int) error
    Produce(topic string, key string, message []byte) error
    ProduceBatch(topic string, messages []KeyedMessage) error
    Subscribe(topic string, groupID string, handler MessageHandler) error
    GetMetrics() BrokerMetrics
    Shutdown(ctx context.Context) error
}

type KeyedMessage struct {
    Key     string
    Value   []byte
    Headers map[string]string
}

type BrokerMetrics struct {
    MessagesPerSecond  float64
    BytesPerSecond     float64
    AvgLatency         time.Duration
    ErrorRate          float64
    ActiveConnections  int
    TopicMetrics       map[string]TopicMetrics
}

type TopicMetrics struct {
    Partitions      int
    MessagesProduced int64
    MessagesConsumed int64
    BytesProduced   int64
    BytesConsumed   int64
}
```

### ✅ Criterios de Éxito

1. **Performance**: 10,000+ mensajes/segundo
2. **Latency**: < 1ms promedio para produce
3. **Batching**: Agrupa mensajes automáticamente
4. **Partitioning**: Distribución por key hash
5. **Load Balancing**: Entre consumers del mismo group
6. **Monitoring**: Métricas detalladas en tiempo real

### 🧪 Performance Benchmark

```go
func BenchmarkHighThroughput(b *testing.B) {
    broker := NewHighPerformanceBroker()
    defer broker.Shutdown(context.Background())
    
    err := broker.CreateTopic("benchmark", 4)
    require.NoError(b, err)
    
    b.ResetTimer()
    b.SetParallelism(10)
    
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            key := fmt.Sprintf("key-%d", rand.Intn(1000))
            message := []byte("benchmark message payload")
            
            err := broker.Produce("benchmark", key, message)
            if err != nil {
                b.Error(err)
            }
        }
    })
    
    b.StopTimer()
    
    metrics := broker.GetMetrics()
    b.Logf("Messages/sec: %.2f", metrics.MessagesPerSecond)
    b.Logf("Avg Latency: %v", metrics.AvgLatency)
}

func TestLoadBalancing(t *testing.T) {
    // TODO: Test que múltiples consumers reciban mensajes balanceados
}

func TestPartitioning(t *testing.T) {
    // TODO: Test que messages con mismo key van a misma partición
}
```

---

## 🔗 Ejercicio 5: Integración Multi-Broker

**Dificultad**: Avanzado  
**Tiempo estimado**: 75 minutos

### 📝 Descripción

Implementa una abstracción que permita trabajar con múltiples message brokers (Redis, RabbitMQ, in-memory) de manera transparente.

### 🎯 Requisitos

```go
type MultiBroker interface {
    AddBroker(name string, broker MessageBroker) error
    SetRoutingRule(pattern string, brokerName string) error
    Produce(topic string, message Message) error
    Subscribe(topic string, handler MessageHandler) error
    GetBrokerHealth() map[string]HealthStatus
}

type MessageBroker interface {
    Produce(topic string, message Message) error
    Subscribe(topic string, handler MessageHandler) error
    Health() HealthStatus
    Close() error
}

type HealthStatus struct {
    Status    string                 `json:"status"` // "healthy", "degraded", "unhealthy"
    Latency   time.Duration          `json:"latency"`
    Errors    int64                  `json:"errors"`
    LastCheck time.Time              `json:"last_check"`
    Details   map[string]interface{} `json:"details"`
}
```

### ✅ Criterios de Éxito

1. **Failover**: Si un broker falla, usar backup automáticamente
2. **Load Balancing**: Distribuir carga entre brokers disponibles
3. **Health Monitoring**: Verificación continua de salud
4. **Routing Rules**: Envío a brokers específicos por pattern
5. **Graceful Degradation**: Funciona aunque algunos brokers fallen

### 🧪 Test Template

```go
func TestMultiBrokerFailover(t *testing.T) {
    multiBroker := NewMultiBroker()
    
    // Add brokers
    primaryBroker := NewMemoryBroker()
    backupBroker := NewMemoryBroker()
    
    multiBroker.AddBroker("primary", primaryBroker)
    multiBroker.AddBroker("backup", backupBroker)
    
    // Set routing (primary first, backup on failure)
    multiBroker.SetRoutingRule("*", "primary,backup")
    
    // Subscribe to both brokers
    received := make(chan Message, 10)
    handler := &TestHandler{
        onHandle: func(ctx context.Context, msg Message) error {
            received <- msg
            return nil
        },
    }
    
    multiBroker.Subscribe("test.topic", handler)
    
    // Send message (should go to primary)
    msg := Message{ID: "test-1", Topic: "test.topic"}
    err := multiBroker.Produce("test.topic", msg)
    assert.NoError(t, err)
    
    // Simulate primary broker failure
    primaryBroker.Close()
    
    // Send another message (should go to backup)
    msg2 := Message{ID: "test-2", Topic: "test.topic"}
    err = multiBroker.Produce("test.topic", msg2)
    assert.NoError(t, err)
    
    // Verify both messages received
    assert.Eventually(t, func() bool {
        return len(received) == 2
    }, time.Second*5, time.Millisecond*100)
}
```

---

## 📊 Ejercicio 6: Monitoring y Alertas

**Dificultad**: Intermedio  
**Tiempo estimado**: 60 minutos

### 📝 Descripción

Implementa un sistema completo de monitoring para message queues con métricas Prometheus y alertas.

### 🎯 Requisitos

```go
type QueueMonitor interface {
    RecordProduced(queue string, size int)
    RecordConsumed(queue string, duration time.Duration, success bool)
    RecordQueueDepth(queue string, depth int)
    GetCurrentMetrics() MetricsSnapshot
    SetupAlerts(rules []AlertRule) error
}

type AlertRule struct {
    Name        string        `json:"name"`
    Metric      string        `json:"metric"`
    Condition   string        `json:"condition"` // ">", "<", "=="
    Threshold   float64       `json:"threshold"`
    Duration    time.Duration `json:"duration"`
    Callback    AlertCallback `json:"-"`
}

type AlertCallback func(alert Alert)

type Alert struct {
    Rule        AlertRule     `json:"rule"`
    Value       float64       `json:"value"`
    Timestamp   time.Time     `json:"timestamp"`
    Severity    string        `json:"severity"`
    Description string        `json:"description"`
}

type MetricsSnapshot struct {
    QueueDepths      map[string]int     `json:"queue_depths"`
    ProcessingRates  map[string]float64 `json:"processing_rates"`
    ErrorRates       map[string]float64 `json:"error_rates"`
    AvgLatencies     map[string]time.Duration `json:"avg_latencies"`
    Timestamp        time.Time          `json:"timestamp"`
}
```

### ✅ Criterios de Éxito

1. Métricas Prometheus exportadas correctamente
2. Alertas se disparan según condiciones
3. Dashboard con métricas en tiempo real
4. Historial de alertas
5. Health check endpoint

### 🧪 Test Template

```go
func TestMetricsCollection(t *testing.T) {
    monitor := NewQueueMonitor()
    
    // Simulate queue activity
    monitor.RecordProduced("orders", 100)
    monitor.RecordConsumed("orders", time.Millisecond*50, true)
    monitor.RecordConsumed("orders", time.Millisecond*75, false)
    monitor.RecordQueueDepth("orders", 25)
    
    metrics := monitor.GetCurrentMetrics()
    
    assert.Equal(t, 25, metrics.QueueDepths["orders"])
    assert.True(t, metrics.ProcessingRates["orders"] > 0)
    assert.True(t, metrics.ErrorRates["orders"] > 0)
    assert.True(t, metrics.AvgLatencies["orders"] > 0)
}

func TestAlerting(t *testing.T) {
    monitor := NewQueueMonitor()
    
    alertTriggered := false
    alertRule := AlertRule{
        Name: "High Queue Depth",
        Metric: "queue_depth",
        Condition: ">",
        Threshold: 100,
        Duration: time.Second,
        Callback: func(alert Alert) {
            alertTriggered = true
        },
    }
    
    monitor.SetupAlerts([]AlertRule{alertRule})
    
    // Trigger alert condition
    monitor.RecordQueueDepth("orders", 150)
    
    // Wait for alert
    assert.Eventually(t, func() bool {
        return alertTriggered
    }, time.Second*2, time.Millisecond*100)
}
```

---

## 🎯 Ejercicios Bonus

### 🌟 Bonus 1: Event Sourcing

Implementa un event store con snapshots y replay de eventos.

### 🌟 Bonus 2: CQRS Pattern

Separa commands y queries usando message queues.

### 🌟 Bonus 3: Saga Orchestration

Implementa el patrón Saga para transacciones distribuidas.

### 🌟 Bonus 4: Message Deduplication

Sistema para detectar y eliminar mensajes duplicados.

### 🌟 Bonus 5: Geographic Distribution

Message queues que funcionen across múltiples regiones.

---

## 📋 Guía de Evaluación

### ✅ Criterios de Evaluación

| Aspecto            | Peso | Descripción                     |
| ------------------ | ---- | ------------------------------- |
| **Funcionalidad**  | 40%  | Todos los tests pasan           |
| **Performance**    | 20%  | Benchmarks cumplen requisitos   |
| **Código Limpio**  | 15%  | Arquitectura clara y mantenible |
| **Error Handling** | 15%  | Manejo robusto de errores       |
| **Testing**        | 10%  | Tests comprehensivos            |

### 🎯 Puntuación

- **90-100%**: 🏆 Expert Level
- **80-89%**: 🥇 Advanced 
- **70-79%**: 🥈 Intermediate
- **60-69%**: 🥉 Beginner
- **<60%**: 📚 Needs More Practice

---

## 💡 Tips para el Éxito

1. **🧪 Test First**: Escribe tests antes que implementación
2. **📊 Measure Everything**: Performance es crítico en message queues
3. **🔄 Think Async**: Diseña para operaciones no bloqueantes
4. **🎯 Start Simple**: Implementa funcionalidad básica primero
5. **📚 Read Documentation**: Entiende bien los brokers que usas
6. **🔒 Handle Failures**: Los sistemas distribuidos fallan, prepárate
7. **⚡ Optimize Later**: Haz que funcione, luego optimiza

¡Buena suerte con los ejercicios! 🚀
