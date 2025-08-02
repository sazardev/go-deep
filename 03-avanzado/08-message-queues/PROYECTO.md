# 🚀 Proyecto: Sistema de Mensajería para E-commerce Distribuido

> **Objetivo**: Construir un sistema completo de message queues para una plataforma de e-commerce que maneje millones de transacciones diarias con alta disponibilidad y confiabilidad.

## 📋 Visión General del Proyecto

Vas a implementar **"ShopStream"**, un sistema de mensajería distribuido que orqueste todos los procesos de una plataforma de e-commerce moderna. Este sistema debe manejar desde la creación de órdenes hasta las notificaciones finales, pasando por pagos, inventario, shipping y analytics.

### 🎯 Objetivos de Aprendizaje

- **🏗️ Arquitectura**: Diseñar sistemas event-driven robustos
- **⚡ Performance**: Manejar alta concurrencia y throughput
- **🔒 Reliability**: Garantizar procesamiento de mensajes críticos
- **📊 Observability**: Monitoring y alertas en tiempo real
- **🔧 Integration**: Conectar múltiples servicios y brokers
- **🌐 Scalability**: Diseño para escalar horizontalmente

---

## 🏛️ Arquitectura del Sistema

```mermaid
graph TB
    subgraph "Frontend"
        WEB[Web App]
        MOBILE[Mobile App]
        API[API Gateway]
    end
    
    subgraph "Message Brokers"
        KAFKA[Kafka Cluster]
        REDIS[Redis Streams]
        RABBITMQ[RabbitMQ]
    end
    
    subgraph "Core Services"
        ORDER[Order Service]
        PAY[Payment Service]
        INV[Inventory Service]
        SHIP[Shipping Service]
        NOTIF[Notification Service]
    end
    
    subgraph "Support Services"
        AUDIT[Audit Service]
        ANALYTICS[Analytics Service]
        FRAUD[Fraud Detection]
    end
    
    subgraph "Databases"
        POSTGRES[(PostgreSQL)]
        MONGO[(MongoDB)]
        ELASTIC[(Elasticsearch)]
    end
    
    WEB --> API
    MOBILE --> API
    API --> ORDER
    
    ORDER --> KAFKA
    KAFKA --> PAY
    KAFKA --> INV
    KAFKA --> SHIP
    KAFKA --> NOTIF
    KAFKA --> AUDIT
    
    PAY --> REDIS
    REDIS --> FRAUD
    
    NOTIF --> RABBITMQ
    
    ORDER --> POSTGRES
    ANALYTICS --> MONGO
    AUDIT --> ELASTIC
```

## 📦 Componentes del Proyecto

### 1. 🎛️ Message Router Core

El corazón del sistema que maneja routing inteligente entre brokers.

**Archivo**: `cmd/router/main.go`

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "shopstream/internal/router"
    "shopstream/internal/config"
    "shopstream/internal/monitoring"
)

func main() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Initialize monitoring
    monitor := monitoring.New(cfg.Monitoring)
    
    // Create message router
    messageRouter, err := router.New(cfg.Router, monitor)
    if err != nil {
        log.Fatalf("Failed to create router: %v", err)
    }
    
    // Start router
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    if err := messageRouter.Start(ctx); err != nil {
        log.Fatalf("Failed to start router: %v", err)
    }
    
    // Wait for shutdown signal
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    <-sigChan
    log.Println("Shutting down...")
    
    cancel()
    messageRouter.Shutdown(context.Background())
}
```

### 2. 🏪 Order Service

Maneja el ciclo de vida completo de las órdenes.

**Archivo**: `internal/services/order/service.go`

```go
package order

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    
    "shopstream/internal/models"
    "shopstream/internal/messaging"
    "shopstream/internal/repository"
)

type Service struct {
    repo       repository.OrderRepository
    publisher  messaging.Publisher
    subscriber messaging.Subscriber
}

func New(repo repository.OrderRepository, 
    publisher messaging.Publisher,
    subscriber messaging.Subscriber) *Service {
    
    return &Service{
        repo:       repo,
        publisher:  publisher,
        subscriber: subscriber,
    }
}

func (s *Service) CreateOrder(ctx context.Context, req CreateOrderRequest) (*models.Order, error) {
    // Validate order
    if err := s.validateOrder(req); err != nil {
        return nil, fmt.Errorf("invalid order: %w", err)
    }
    
    // Create order entity
    order := &models.Order{
        ID:         generateOrderID(),
        CustomerID: req.CustomerID,
        Items:      req.Items,
        Status:     models.OrderStatusPending,
        Total:      calculateTotal(req.Items),
        CreatedAt:  time.Now(),
    }
    
    // Save to database
    if err := s.repo.Create(ctx, order); err != nil {
        return nil, fmt.Errorf("failed to save order: %w", err)
    }
    
    // Publish order created event
    event := models.OrderCreatedEvent{
        OrderID:    order.ID,
        CustomerID: order.CustomerID,
        Items:      order.Items,
        Total:      order.Total,
        Timestamp:  time.Now(),
    }
    
    if err := s.publishEvent(ctx, "order.created", event); err != nil {
        // Log error but don't fail the order creation
        log.Printf("Failed to publish order created event: %v", err)
    }
    
    return order, nil
}

func (s *Service) processOrderEvents(ctx context.Context) error {
    // Listen for payment and inventory events
    paymentHandler := func(msg messaging.Message) error {
        var event models.PaymentProcessedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            return err
        }
        return s.handlePaymentProcessed(ctx, event)
    }
    
    inventoryHandler := func(msg messaging.Message) error {
        var event models.InventoryReservedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            return err
        }
        return s.handleInventoryReserved(ctx, event)
    }
    
    // Subscribe to events
    if err := s.subscriber.Subscribe(ctx, "payment.processed", paymentHandler); err != nil {
        return err
    }
    
    if err := s.subscriber.Subscribe(ctx, "inventory.reserved", inventoryHandler); err != nil {
        return err
    }
    
    return nil
}

func (s *Service) handlePaymentProcessed(ctx context.Context, event models.PaymentProcessedEvent) error {
    order, err := s.repo.GetByID(ctx, event.OrderID)
    if err != nil {
        return err
    }
    
    if event.Success {
        order.PaymentStatus = models.PaymentStatusCompleted
        order.UpdatedAt = time.Now()
    } else {
        order.Status = models.OrderStatusCancelled
        order.PaymentStatus = models.PaymentStatusFailed
        order.UpdatedAt = time.Now()
    }
    
    if err := s.repo.Update(ctx, order); err != nil {
        return err
    }
    
    // Check if order is ready to ship
    return s.checkOrderReadiness(ctx, order)
}

func (s *Service) checkOrderReadiness(ctx context.Context, order *models.Order) error {
    if order.PaymentStatus == models.PaymentStatusCompleted &&
       order.InventoryStatus == models.InventoryStatusReserved {
        
        // Order is ready to ship
        order.Status = models.OrderStatusReadyToShip
        order.UpdatedAt = time.Now()
        
        if err := s.repo.Update(ctx, order); err != nil {
            return err
        }
        
        // Publish ready to ship event
        event := models.OrderReadyToShipEvent{
            OrderID:   order.ID,
            Items:     order.Items,
            Address:   order.ShippingAddress,
            Timestamp: time.Now(),
        }
        
        return s.publishEvent(ctx, "order.ready_to_ship", event)
    }
    
    return nil
}
```

### 3. 💳 Payment Service

Procesa pagos con múltiples providers y fraud detection.

**Archivo**: `internal/services/payment/service.go`

```go
package payment

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    
    "shopstream/internal/models"
    "shopstream/internal/messaging"
    "shopstream/internal/fraud"
)

type Service struct {
    gateways   map[string]PaymentGateway
    fraud      fraud.Detector
    publisher  messaging.Publisher
    subscriber messaging.Subscriber
}

type PaymentGateway interface {
    ProcessPayment(ctx context.Context, req PaymentRequest) (*PaymentResponse, error)
    RefundPayment(ctx context.Context, transactionID string, amount float64) error
}

func New(gateways map[string]PaymentGateway,
    fraud fraud.Detector,
    publisher messaging.Publisher,
    subscriber messaging.Subscriber) *Service {
    
    service := &Service{
        gateways:   gateways,
        fraud:      fraud,
        publisher:  publisher,
        subscriber: subscriber,
    }
    
    // Start event processing
    go service.processOrderEvents(context.Background())
    
    return service
}

func (s *Service) processOrderEvents(ctx context.Context) error {
    handler := func(msg messaging.Message) error {
        var event models.OrderCreatedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            return err
        }
        return s.processPayment(ctx, event)
    }
    
    return s.subscriber.Subscribe(ctx, "order.created", handler)
}

func (s *Service) processPayment(ctx context.Context, orderEvent models.OrderCreatedEvent) error {
    // Fraud detection
    fraudResult, err := s.fraud.CheckOrder(ctx, fraud.OrderData{
        OrderID:    orderEvent.OrderID,
        CustomerID: orderEvent.CustomerID,
        Amount:     orderEvent.Total,
        Items:      orderEvent.Items,
    })
    if err != nil {
        return fmt.Errorf("fraud check failed: %w", err)
    }
    
    if fraudResult.IsFraudulent {
        return s.publishPaymentEvent(ctx, orderEvent.OrderID, false, 
            "Payment blocked due to fraud detection")
    }
    
    // Process payment based on amount
    gateway := s.selectGateway(orderEvent.Total)
    
    paymentReq := PaymentRequest{
        OrderID:       orderEvent.OrderID,
        Amount:        orderEvent.Total,
        Currency:      "USD",
        PaymentMethod: "credit_card", // This would come from order
    }
    
    response, err := gateway.ProcessPayment(ctx, paymentReq)
    if err != nil {
        return s.publishPaymentEvent(ctx, orderEvent.OrderID, false, err.Error())
    }
    
    success := response.Status == "success"
    message := response.Message
    
    return s.publishPaymentEvent(ctx, orderEvent.OrderID, success, message)
}

func (s *Service) selectGateway(amount float64) PaymentGateway {
    // Simple routing logic
    if amount > 1000 {
        return s.gateways["premium"] // Use premium gateway for high amounts
    }
    return s.gateways["standard"]
}

func (s *Service) publishPaymentEvent(ctx context.Context, orderID string, 
    success bool, message string) error {
    
    event := models.PaymentProcessedEvent{
        OrderID:   orderID,
        Success:   success,
        Message:   message,
        Timestamp: time.Now(),
    }
    
    eventData, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    msg := messaging.Message{
        ID:        fmt.Sprintf("payment-%s-%d", orderID, time.Now().UnixNano()),
        Topic:     "payment.processed",
        Data:      eventData,
        Timestamp: time.Now(),
    }
    
    return s.publisher.Publish(ctx, msg)
}

// Circuit Breaker for payment gateways
type CircuitBreakerGateway struct {
    gateway PaymentGateway
    breaker *CircuitBreaker
}

func NewCircuitBreakerGateway(gateway PaymentGateway) *CircuitBreakerGateway {
    return &CircuitBreakerGateway{
        gateway: gateway,
        breaker: NewCircuitBreaker(CircuitBreakerConfig{
            MaxRequests:        100,
            Interval:           time.Minute,
            Timeout:            30 * time.Second,
            FailureThreshold:   5,
            OnStateChange: func(name string, from, to State) {
                log.Printf("Circuit breaker %s changed from %s to %s", name, from, to)
            },
        }),
    }
}

func (cbg *CircuitBreakerGateway) ProcessPayment(ctx context.Context, 
    req PaymentRequest) (*PaymentResponse, error) {
    
    result, err := cbg.breaker.Execute(func() (interface{}, error) {
        return cbg.gateway.ProcessPayment(ctx, req)
    })
    
    if err != nil {
        return nil, err
    }
    
    return result.(*PaymentResponse), nil
}
```

### 4. 📦 Inventory Service

Maneja reservas de inventario con optimistic locking.

**Archivo**: `internal/services/inventory/service.go`

```go
package inventory

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "time"
    
    "shopstream/internal/models"
    "shopstream/internal/messaging"
)

type Service struct {
    db         *sql.DB
    publisher  messaging.Publisher
    subscriber messaging.Subscriber
    cache      Cache
}

type Cache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{}, ttl time.Duration)
    Delete(key string)
}

func New(db *sql.DB, publisher messaging.Publisher, 
    subscriber messaging.Subscriber, cache Cache) *Service {
    
    service := &Service{
        db:         db,
        publisher:  publisher,
        subscriber: subscriber,
        cache:      cache,
    }
    
    go service.processOrderEvents(context.Background())
    
    return service
}

func (s *Service) processOrderEvents(ctx context.Context) error {
    handler := func(msg messaging.Message) error {
        var event models.OrderCreatedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            return err
        }
        return s.reserveInventory(ctx, event)
    }
    
    return s.subscriber.Subscribe(ctx, "order.created", handler)
}

func (s *Service) reserveInventory(ctx context.Context, orderEvent models.OrderCreatedEvent) error {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    success := true
    var failedItems []string
    
    // Reserve each item
    for _, item := range orderEvent.Items {
        available, err := s.checkAvailability(ctx, tx, item.ProductID, item.Quantity)
        if err != nil {
            return err
        }
        
        if !available {
            success = false
            failedItems = append(failedItems, item.ProductID)
            continue
        }
        
        // Reserve with optimistic locking
        if err := s.reserveItem(ctx, tx, item.ProductID, item.Quantity, orderEvent.OrderID); err != nil {
            success = false
            failedItems = append(failedItems, item.ProductID)
        }
    }
    
    if success {
        if err := tx.Commit(); err != nil {
            return err
        }
        
        // Clear cache for updated items
        for _, item := range orderEvent.Items {
            s.cache.Delete(fmt.Sprintf("inventory:%s", item.ProductID))
        }
    } else {
        tx.Rollback()
    }
    
    return s.publishInventoryEvent(ctx, orderEvent.OrderID, success, failedItems)
}

func (s *Service) checkAvailability(ctx context.Context, tx *sql.Tx, 
    productID string, quantity int) (bool, error) {
    
    // Try cache first
    cacheKey := fmt.Sprintf("inventory:%s", productID)
    if cached, found := s.cache.Get(cacheKey); found {
        if stock, ok := cached.(int); ok && stock >= quantity {
            return true, nil
        }
        return false, nil
    }
    
    // Query database
    var available int
    query := `SELECT available_quantity FROM inventory WHERE product_id = $1`
    err := tx.QueryRowContext(ctx, query, productID).Scan(&available)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil
        }
        return false, err
    }
    
    // Cache the result
    s.cache.Set(cacheKey, available, 5*time.Minute)
    
    return available >= quantity, nil
}

func (s *Service) reserveItem(ctx context.Context, tx *sql.Tx, 
    productID string, quantity int, orderID string) error {
    
    // Optimistic locking with version
    query := `
        UPDATE inventory 
        SET 
            available_quantity = available_quantity - $1,
            reserved_quantity = reserved_quantity + $1,
            version = version + 1,
            updated_at = NOW()
        WHERE product_id = $2 
          AND available_quantity >= $1
    `
    
    result, err := tx.ExecContext(ctx, query, quantity, productID)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("insufficient inventory for product %s", productID)
    }
    
    // Record reservation
    reservationQuery := `
        INSERT INTO inventory_reservations (order_id, product_id, quantity, created_at)
        VALUES ($1, $2, $3, NOW())
    `
    
    _, err = tx.ExecContext(ctx, reservationQuery, orderID, productID, quantity)
    return err
}

func (s *Service) publishInventoryEvent(ctx context.Context, orderID string, 
    success bool, failedItems []string) error {
    
    event := models.InventoryReservedEvent{
        OrderID:     orderID,
        Success:     success,
        FailedItems: failedItems,
        Timestamp:   time.Now(),
    }
    
    eventData, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    msg := messaging.Message{
        ID:        fmt.Sprintf("inventory-%s-%d", orderID, time.Now().UnixNano()),
        Topic:     "inventory.reserved",
        Data:      eventData,
        Timestamp: time.Now(),
    }
    
    return s.publisher.Publish(ctx, msg)
}

// Background job to release expired reservations
func (s *Service) startReservationCleaner(ctx context.Context) {
    ticker := time.NewTicker(5 * time.Minute)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            if err := s.cleanupExpiredReservations(ctx); err != nil {
                log.Printf("Failed to cleanup expired reservations: %v", err)
            }
        case <-ctx.Done():
            return
        }
    }
}

func (s *Service) cleanupExpiredReservations(ctx context.Context) error {
    query := `
        SELECT order_id, product_id, quantity 
        FROM inventory_reservations 
        WHERE created_at < NOW() - INTERVAL '30 minutes'
          AND status = 'reserved'
    `
    
    rows, err := s.db.QueryContext(ctx, query)
    if err != nil {
        return err
    }
    defer rows.Close()
    
    for rows.Next() {
        var orderID, productID string
        var quantity int
        
        if err := rows.Scan(&orderID, &productID, &quantity); err != nil {
            continue
        }
        
        // Release reservation
        if err := s.releaseReservation(ctx, orderID, productID, quantity); err != nil {
            log.Printf("Failed to release reservation for order %s: %v", orderID, err)
        }
    }
    
    return nil
}
```

### 5. 📊 Analytics Service

Procesa eventos para generar insights de negocio.

**Archivo**: `internal/services/analytics/service.go`

```go
package analytics

import (
    "context"
    "encoding/json"
    "time"
    
    "shopstream/internal/models"
    "shopstream/internal/messaging"
)

type Service struct {
    subscriber messaging.Subscriber
    storage    AnalyticsStorage
    aggregator *EventAggregator
}

type AnalyticsStorage interface {
    StoreEvent(ctx context.Context, event AnalyticsEvent) error
    GetMetrics(ctx context.Context, query MetricsQuery) (*MetricsResult, error)
}

type EventAggregator struct {
    buffer    chan AnalyticsEvent
    batchSize int
    storage   AnalyticsStorage
}

func New(subscriber messaging.Subscriber, storage AnalyticsStorage) *Service {
    service := &Service{
        subscriber: subscriber,
        storage:    storage,
        aggregator: NewEventAggregator(storage, 1000),
    }
    
    go service.processEvents(context.Background())
    go service.aggregator.Start(context.Background())
    
    return service
}

func (s *Service) processEvents(ctx context.Context) error {
    // Process multiple event types
    eventHandlers := map[string]func(messaging.Message) error{
        "order.created":       s.handleOrderCreated,
        "payment.processed":   s.handlePaymentProcessed,
        "inventory.reserved":  s.handleInventoryReserved,
        "order.shipped":       s.handleOrderShipped,
        "user.behavior":       s.handleUserBehavior,
    }
    
    for topic, handler := range eventHandlers {
        if err := s.subscriber.Subscribe(ctx, topic, handler); err != nil {
            return err
        }
    }
    
    return nil
}

func (s *Service) handleOrderCreated(msg messaging.Message) error {
    var event models.OrderCreatedEvent
    if err := json.Unmarshal(msg.Data, &event); err != nil {
        return err
    }
    
    analyticsEvent := AnalyticsEvent{
        Type:      "order_created",
        Timestamp: event.Timestamp,
        Data: map[string]interface{}{
            "order_id":    event.OrderID,
            "customer_id": event.CustomerID,
            "total":       event.Total,
            "item_count":  len(event.Items),
        },
        Dimensions: map[string]string{
            "customer_segment": s.getCustomerSegment(event.CustomerID),
            "order_size":       s.getOrderSizeCategory(event.Total),
        },
    }
    
    s.aggregator.AddEvent(analyticsEvent)
    return nil
}

func (s *Service) handlePaymentProcessed(msg messaging.Message) error {
    var event models.PaymentProcessedEvent
    if err := json.Unmarshal(msg.Data, &event); err != nil {
        return err
    }
    
    analyticsEvent := AnalyticsEvent{
        Type:      "payment_processed",
        Timestamp: event.Timestamp,
        Data: map[string]interface{}{
            "order_id": event.OrderID,
            "success":  event.Success,
        },
        Dimensions: map[string]string{
            "payment_status": func() string {
                if event.Success {
                    return "success"
                }
                return "failed"
            }(),
        },
    }
    
    s.aggregator.AddEvent(analyticsEvent)
    return nil
}

// Real-time metrics calculation
func (s *Service) calculateRealTimeMetrics(ctx context.Context) (*RealTimeMetrics, error) {
    now := time.Now()
    last24h := now.Add(-24 * time.Hour)
    
    query := MetricsQuery{
        StartTime:  last24h,
        EndTime:    now,
        GroupBy:    []string{"hour"},
        Aggregates: []string{"count", "sum", "avg"},
    }
    
    result, err := s.storage.GetMetrics(ctx, query)
    if err != nil {
        return nil, err
    }
    
    metrics := &RealTimeMetrics{
        OrdersPerHour:     result.GetSeriesData("order_created", "count"),
        RevenuePerHour:    result.GetSeriesData("order_created", "sum:total"),
        ConversionRate:    s.calculateConversionRate(result),
        AverageOrderValue: result.GetValue("order_created", "avg:total"),
        Timestamp:         now,
    }
    
    return metrics, nil
}

func NewEventAggregator(storage AnalyticsStorage, batchSize int) *EventAggregator {
    return &EventAggregator{
        buffer:    make(chan AnalyticsEvent, batchSize*2),
        batchSize: batchSize,
        storage:   storage,
    }
}

func (ea *EventAggregator) AddEvent(event AnalyticsEvent) {
    select {
    case ea.buffer <- event:
    default:
        // Buffer full, drop event (or implement overflow strategy)
        log.Println("Analytics buffer full, dropping event")
    }
}

func (ea *EventAggregator) Start(ctx context.Context) {
    ticker := time.NewTicker(10 * time.Second) // Flush every 10 seconds
    defer ticker.Stop()
    
    batch := make([]AnalyticsEvent, 0, ea.batchSize)
    
    for {
        select {
        case event := <-ea.buffer:
            batch = append(batch, event)
            if len(batch) >= ea.batchSize {
                ea.flushBatch(ctx, batch)
                batch = batch[:0]
            }
            
        case <-ticker.C:
            if len(batch) > 0 {
                ea.flushBatch(ctx, batch)
                batch = batch[:0]
            }
            
        case <-ctx.Done():
            if len(batch) > 0 {
                ea.flushBatch(ctx, batch)
            }
            return
        }
    }
}

func (ea *EventAggregator) flushBatch(ctx context.Context, batch []AnalyticsEvent) {
    for _, event := range batch {
        if err := ea.storage.StoreEvent(ctx, event); err != nil {
            log.Printf("Failed to store analytics event: %v", err)
        }
    }
    log.Printf("Flushed %d analytics events", len(batch))
}
```

### 6. 📊 Monitoring Dashboard

Sistema completo de monitoring con métricas Prometheus.

**Archivo**: `internal/monitoring/dashboard.go`

```go
package monitoring

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    messagesProcessed = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "shopstream_messages_processed_total",
            Help: "Total number of messages processed",
        },
        []string{"service", "topic", "status"},
    )
    
    messageLatency = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "shopstream_message_processing_duration_seconds",
            Help:    "Time spent processing messages",
            Buckets: prometheus.DefBuckets,
        },
        []string{"service", "topic"},
    )
    
    queueDepth = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "shopstream_queue_depth",
            Help: "Current number of messages in queue",
        },
        []string{"queue_name"},
    )
    
    orderMetrics = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "shopstream_orders_total",
            Help: "Total number of orders",
        },
        []string{"status"},
    )
    
    revenueMetrics = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "shopstream_revenue_total",
            Help: "Total revenue in USD",
        },
        []string{"currency"},
    )
)

type Monitor struct {
    registry *prometheus.Registry
    server   *http.Server
}

func New(config Config) *Monitor {
    monitor := &Monitor{
        registry: prometheus.NewRegistry(),
    }
    
    // Register collectors
    monitor.registry.MustRegister(messagesProcessed)
    monitor.registry.MustRegister(messageLatency)
    monitor.registry.MustRegister(queueDepth)
    monitor.registry.MustRegister(orderMetrics)
    monitor.registry.MustRegister(revenueMetrics)
    
    // Setup HTTP server
    mux := http.NewMux()
    mux.Handle("/metrics", promhttp.HandlerFor(
        monitor.registry,
        promhttp.HandlerOpts{},
    ))
    mux.HandleFunc("/health", monitor.healthCheck)
    mux.HandleFunc("/dashboard", monitor.dashboard)
    
    monitor.server = &http.Server{
        Addr:    config.ListenAddr,
        Handler: mux,
    }
    
    return monitor
}

func (m *Monitor) RecordMessageProcessed(service, topic, status string, duration time.Duration) {
    messagesProcessed.WithLabelValues(service, topic, status).Inc()
    messageLatency.WithLabelValues(service, topic).Observe(duration.Seconds())
}

func (m *Monitor) UpdateQueueDepth(queueName string, depth float64) {
    queueDepth.WithLabelValues(queueName).Set(depth)
}

func (m *Monitor) RecordOrder(status string, amount float64) {
    orderMetrics.WithLabelValues(status).Inc()
    if status == "completed" {
        revenueMetrics.WithLabelValues("USD").Add(amount)
    }
}

func (m *Monitor) healthCheck(w http.ResponseWriter, r *http.Request) {
    health := HealthStatus{
        Status:    "healthy",
        Timestamp: time.Now(),
        Services:  make(map[string]ServiceHealth),
    }
    
    // Check service health
    services := []string{"order", "payment", "inventory", "shipping"}
    for _, service := range services {
        health.Services[service] = m.checkServiceHealth(service)
    }
    
    // Overall health
    allHealthy := true
    for _, serviceHealth := range health.Services {
        if serviceHealth.Status != "healthy" {
            allHealthy = false
            break
        }
    }
    
    if !allHealthy {
        health.Status = "degraded"
        w.WriteHeader(http.StatusServiceUnavailable)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(health)
}

func (m *Monitor) dashboard(w http.ResponseWriter, r *http.Request) {
    dashboard := DashboardData{
        Timestamp: time.Now(),
        Metrics:   m.getCurrentMetrics(),
        Charts:    m.getChartData(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(dashboard)
}

func (m *Monitor) getCurrentMetrics() MetricsSummary {
    return MetricsSummary{
        TotalOrders:       m.getCounterValue(orderMetrics, "created"),
        CompletedOrders:   m.getCounterValue(orderMetrics, "completed"),
        TotalRevenue:      m.getCounterValue(revenueMetrics, "USD"),
        MessagesPerSecond: m.calculateMessagesPerSecond(),
        AverageLatency:    m.calculateAverageLatency(),
    }
}

type HealthStatus struct {
    Status    string                    `json:"status"`
    Timestamp time.Time                 `json:"timestamp"`
    Services  map[string]ServiceHealth  `json:"services"`
}

type ServiceHealth struct {
    Status      string        `json:"status"`
    Latency     time.Duration `json:"latency"`
    ErrorRate   float64       `json:"error_rate"`
    LastChecked time.Time     `json:"last_checked"`
}

type DashboardData struct {
    Timestamp time.Time      `json:"timestamp"`
    Metrics   MetricsSummary `json:"metrics"`
    Charts    []ChartData    `json:"charts"`
}

type MetricsSummary struct {
    TotalOrders       float64 `json:"total_orders"`
    CompletedOrders   float64 `json:"completed_orders"`
    TotalRevenue      float64 `json:"total_revenue"`
    MessagesPerSecond float64 `json:"messages_per_second"`
    AverageLatency    float64 `json:"average_latency"`
}

type ChartData struct {
    Name   string      `json:"name"`
    Type   string      `json:"type"`
    Data   []DataPoint `json:"data"`
    Labels []string    `json:"labels"`
}

type DataPoint struct {
    Timestamp time.Time `json:"timestamp"`
    Value     float64   `json:"value"`
}
```

## 🧪 Testing Strategy

### Unit Tests

**Archivo**: `internal/services/order/service_test.go`

```go
package order_test

import (
    "context"
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    
    "shopstream/internal/models"
    "shopstream/internal/services/order"
    "shopstream/internal/mocks"
)

func TestCreateOrder(t *testing.T) {
    // Setup mocks
    mockRepo := &mocks.OrderRepository{}
    mockPublisher := &mocks.Publisher{}
    mockSubscriber := &mocks.Subscriber{}
    
    service := order.New(mockRepo, mockPublisher, mockSubscriber)
    
    // Setup expectations
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Order")).
        Return(nil)
    mockPublisher.On("Publish", mock.Anything, mock.AnythingOfType("messaging.Message")).
        Return(nil)
    
    // Test data
    req := order.CreateOrderRequest{
        CustomerID: "customer-123",
        Items: []models.Item{
            {ProductID: "product-1", Quantity: 2, Price: 10.00},
        },
    }
    
    // Execute
    result, err := service.CreateOrder(context.Background(), req)
    
    // Verify
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, req.CustomerID, result.CustomerID)
    assert.Equal(t, models.OrderStatusPending, result.Status)
    assert.Equal(t, 20.00, result.Total)
    
    mockRepo.AssertExpectations(t)
    mockPublisher.AssertExpectations(t)
}
```

### Integration Tests

**Archivo**: `test/integration/order_flow_test.go`

```go
package integration_test

import (
    "context"
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type OrderFlowTestSuite struct {
    suite.Suite
    testEnv *TestEnvironment
}

func (suite *OrderFlowTestSuite) SetupSuite() {
    suite.testEnv = NewTestEnvironment()
    suite.testEnv.Start()
}

func (suite *OrderFlowTestSuite) TearDownSuite() {
    suite.testEnv.Stop()
}

func (suite *OrderFlowTestSuite) TestCompleteOrderFlow() {
    ctx := context.Background()
    
    // Create order
    order := &models.Order{
        CustomerID: "test-customer",
        Items: []models.Item{
            {ProductID: "test-product", Quantity: 1, Price: 100.00},
        },
    }
    
    createdOrder, err := suite.testEnv.OrderService.CreateOrder(ctx, order)
    suite.NoError(err)
    
    // Wait for processing
    suite.Eventually(func() bool {
        order, _ := suite.testEnv.OrderService.GetOrder(ctx, createdOrder.ID)
        return order != nil && order.Status == models.OrderStatusReadyToShip
    }, 10*time.Second, 100*time.Millisecond)
    
    // Verify final state
    finalOrder, err := suite.testEnv.OrderService.GetOrder(ctx, createdOrder.ID)
    suite.NoError(err)
    suite.Equal(models.OrderStatusReadyToShip, finalOrder.Status)
    suite.Equal(models.PaymentStatusCompleted, finalOrder.PaymentStatus)
    suite.Equal(models.InventoryStatusReserved, finalOrder.InventoryStatus)
}

func TestOrderFlowTestSuite(t *testing.T) {
    suite.Run(t, new(OrderFlowTestSuite))
}
```

### Load Tests

**Archivo**: `test/load/order_load_test.go`

```go
package load_test

import (
    "context"
    "fmt"
    "sync"
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
)

func TestOrderServiceLoad(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping load test in short mode")
    }
    
    testEnv := NewTestEnvironment()
    testEnv.Start()
    defer testEnv.Stop()
    
    const (
        numGoroutines = 100
        ordersPerGoroutine = 100
        totalOrders = numGoroutines * ordersPerGoroutine
    )
    
    var wg sync.WaitGroup
    results := make(chan error, totalOrders)
    
    start := time.Now()
    
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            for j := 0; j < ordersPerGoroutine; j++ {
                order := &models.Order{
                    CustomerID: fmt.Sprintf("customer-%d-%d", workerID, j),
                    Items: []models.Item{
                        {ProductID: "load-test-product", Quantity: 1, Price: 50.00},
                    },
                }
                
                _, err := testEnv.OrderService.CreateOrder(context.Background(), order)
                results <- err
            }
        }(i)
    }
    
    wg.Wait()
    close(results)
    
    duration := time.Since(start)
    
    // Collect results
    var errors []error
    for err := range results {
        if err != nil {
            errors = append(errors, err)
        }
    }
    
    // Verify results
    errorRate := float64(len(errors)) / float64(totalOrders) * 100
    ordersPerSecond := float64(totalOrders) / duration.Seconds()
    
    t.Logf("Load test results:")
    t.Logf("  Total orders: %d", totalOrders)
    t.Logf("  Duration: %v", duration)
    t.Logf("  Orders/sec: %.2f", ordersPerSecond)
    t.Logf("  Error rate: %.2f%%", errorRate)
    
    assert.Less(t, errorRate, 1.0, "Error rate should be less than 1%")
    assert.Greater(t, ordersPerSecond, 100.0, "Should process at least 100 orders/sec")
}
```

## 📈 Performance Requirements

### 🎯 Throughput Targets

- **Orders**: 1,000 órdenes/segundo
- **Messages**: 50,000 mensajes/segundo
- **Analytics**: 100,000 eventos/segundo

### ⚡ Latency Targets

- **Order Creation**: < 100ms p99
- **Payment Processing**: < 500ms p95
- **Message Routing**: < 10ms p95
- **Analytics Processing**: < 1s (batch)

### 📊 Reliability Targets

- **Uptime**: 99.9%
- **Message Delivery**: 99.99% (at-least-once)
- **Data Consistency**: Eventually consistent
- **Recovery Time**: < 5 minutes

## 🚀 Deployment Strategy

### 🐳 Docker Compose

**Archivo**: `docker-compose.yml`

```yaml
version: '3.8'

services:
  # Message Brokers
  kafka:
    image: confluentinc/cp-kafka:latest
    environment:
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"

  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"

  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"

  # Databases
  postgres:
    image: postgres:13
    environment:
      POSTGRES_DB: shopstream
      POSTGRES_USER: shopstream
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"

  mongodb:
    image: mongo:4.4
    ports:
      - "27017:27017"

  # Services
  order-service:
    build:
      context: .
      dockerfile: cmd/order-service/Dockerfile
    environment:
      - DATABASE_URL=postgres://shopstream:password@postgres:5432/shopstream
      - KAFKA_BROKERS=kafka:9092
      - REDIS_URL=redis:6379
    depends_on:
      - postgres
      - kafka
      - redis

  payment-service:
    build:
      context: .
      dockerfile: cmd/payment-service/Dockerfile
    environment:
      - KAFKA_BROKERS=kafka:9092
      - REDIS_URL=redis:6379
    depends_on:
      - kafka
      - redis

  # Monitoring
  prometheus:
    image: prom/prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./monitoring/prometheus.yml:/etc/prometheus/prometheus.yml

  grafana:
    image: grafana/grafana
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
```

### ☸️ Kubernetes Deployment

**Archivo**: `k8s/order-service.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: order-service
  template:
    metadata:
      labels:
        app: order-service
    spec:
      containers:
      - name: order-service
        image: shopstream/order-service:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: url
        - name: KAFKA_BROKERS
          value: "kafka:9092"
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: order-service
spec:
  selector:
    app: order-service
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP
```

## 📋 Entregables del Proyecto

### ✅ Funcionalidad Core

1. **✅ Message Router**: Routing inteligente entre brokers
2. **✅ Order Service**: Gestión completa de órdenes
3. **✅ Payment Service**: Procesamiento con multiple gateways
4. **✅ Inventory Service**: Reservas con optimistic locking
5. **✅ Shipping Service**: Gestión de envíos
6. **✅ Analytics Service**: Procesamiento de eventos en tiempo real

### ✅ Infraestructura

1. **✅ Multi-Broker Support**: Kafka, Redis, RabbitMQ
2. **✅ Database Integration**: PostgreSQL, MongoDB
3. **✅ Caching Layer**: Redis con cache-aside pattern
4. **✅ Monitoring**: Prometheus + Grafana
5. **✅ Health Checks**: Endpoints completos de salud
6. **✅ Circuit Breakers**: Para servicios externos

### ✅ Calidad

1. **✅ Unit Tests**: 90%+ coverage
2. **✅ Integration Tests**: Flujos end-to-end
3. **✅ Load Tests**: Performance validation
4. **✅ Documentation**: API docs + architecture
5. **✅ Logging**: Structured logging con correlation IDs
6. **✅ Tracing**: Distributed tracing setup

## 🎯 Criterios de Evaluación

### 🏆 Excelente (90-100%)

- Todas las funcionalidades implementadas
- Performance targets alcanzados
- Zero downtime deployment
- Comprehensive monitoring
- Clean, maintainable code

### 🥇 Muy Bueno (80-89%)

- Core functionality completa
- Good performance
- Basic monitoring
- Most tests passing
- Good code quality

### 🥈 Bueno (70-79%)

- Basic functionality works
- Acceptable performance
- Some tests
- Working deployment

### 🥉 Necesita Mejora (<70%)

- Incomplete functionality
- Performance issues
- Missing tests
- Deployment problems

---

## 💡 Tips para el Éxito

### 🎯 Enfoque por Fases

1. **Fase 1**: Message queue básico + Order service
2. **Fase 2**: Payment + Inventory services
3. **Fase 3**: Analytics + Monitoring
4. **Fase 4**: Performance optimization
5. **Fase 5**: Production readiness

### 🔧 Herramientas Recomendadas

- **IDE**: VS Code con Go extension
- **Testing**: Testify, Gomega
- **Mocking**: Go Mock, Testify Mock
- **Benchmarking**: Go built-in benchmarks
- **Profiling**: pprof, go-torch
- **Load Testing**: k6, Apache Bench

### 📚 Recursos Adicionales

- [Kafka Go Client Documentation](https://github.com/segmentio/kafka-go)
- [Redis Go Client](https://github.com/go-redis/redis)
- [RabbitMQ Go Client](https://github.com/streadway/amqp)
- [Prometheus Go Client](https://github.com/prometheus/client_golang)

---

¡Buena suerte construyendo **ShopStream**! 🚀 Este proyecto te dará experiencia real en sistemas distribuidos y message queues a escala enterprise.
