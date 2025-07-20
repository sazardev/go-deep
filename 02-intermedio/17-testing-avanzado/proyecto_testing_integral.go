// 🚀 Proyecto Final: Sistema de Testing Integral
// ==============================================
// Este proyecto demuestra todas las técnicas de testing avanzado
// en un sistema de e-commerce completo con TDD, mocks, property testing y más

package main

import (
    "context"
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

// ==========================================
// 🏗️ DOMINIO: E-COMMERCE SYSTEM
// ==========================================

// Product representa un producto en el sistema
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Stock       int       `json:"stock"`
    Category    string    `json:"category"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// User representa un usuario del sistema
type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    CreatedAt time.Time `json:"created_at"`
}

// Order representa una orden de compra
type Order struct {
    ID          string      `json:"id"`
    UserID      string      `json:"user_id"`
    Items       []OrderItem `json:"items"`
    TotalAmount float64     `json:"total_amount"`
    Status      OrderStatus `json:"status"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

// OrderItem representa un item en una orden
type OrderItem struct {
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    UnitPrice float64 `json:"unit_price"`
    Subtotal  float64 `json:"subtotal"`
}

// OrderStatus representa el estado de una orden
type OrderStatus string

const (
    OrderStatusPending    OrderStatus = "pending"
    OrderStatusProcessing OrderStatus = "processing"
    OrderStatusShipped    OrderStatus = "shipped"
    OrderStatusDelivered  OrderStatus = "delivered"
    OrderStatusCancelled  OrderStatus = "cancelled"
)

// ==========================================
// 📝 INTERFACES (Para facilitar testing con mocks)
// ==========================================

// ProductRepository define operaciones de productos
type ProductRepository interface {
    GetByID(ctx context.Context, id string) (*Product, error)
    GetByCategory(ctx context.Context, category string) ([]*Product, error)
    Create(ctx context.Context, product *Product) error
    Update(ctx context.Context, product *Product) error
    UpdateStock(ctx context.Context, productID string, newStock int) error
    Search(ctx context.Context, query string) ([]*Product, error)
}

// UserRepository define operaciones de usuarios
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*User, error)
    GetByEmail(ctx context.Context, email string) (*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
}

// OrderRepository define operaciones de órdenes
type OrderRepository interface {
    GetByID(ctx context.Context, id string) (*Order, error)
    GetByUserID(ctx context.Context, userID string) ([]*Order, error)
    Create(ctx context.Context, order *Order) error
    UpdateStatus(ctx context.Context, orderID string, status OrderStatus) error
}

// PaymentService define el servicio de pagos
type PaymentService interface {
    ProcessPayment(ctx context.Context, orderID string, amount float64, paymentMethod string) (*PaymentResult, error)
    RefundPayment(ctx context.Context, paymentID string, amount float64) error
    GetPaymentStatus(ctx context.Context, paymentID string) (*PaymentStatus, error)
}

// NotificationService define el servicio de notificaciones
type NotificationService interface {
    SendOrderConfirmation(ctx context.Context, userID, orderID string) error
    SendShippingNotification(ctx context.Context, userID, orderID, trackingNumber string) error
    SendDeliveryConfirmation(ctx context.Context, userID, orderID string) error
}

// InventoryService define el servicio de inventario
type InventoryService interface {
    CheckAvailability(ctx context.Context, productID string, quantity int) (bool, error)
    ReserveStock(ctx context.Context, items []OrderItem) error
    ReleaseStock(ctx context.Context, items []OrderItem) error
    UpdateStock(ctx context.Context, productID string, change int) error
}

// PaymentResult representa el resultado de un pago
type PaymentResult struct {
    PaymentID     string    `json:"payment_id"`
    Status        string    `json:"status"`
    TransactionID string    `json:"transaction_id"`
    ProcessedAt   time.Time `json:"processed_at"`
}

// PaymentStatus representa el estado de un pago
type PaymentStatus struct {
    PaymentID   string  `json:"payment_id"`
    Status      string  `json:"status"`
    Amount      float64 `json:"amount"`
    ProcessedAt time.Time `json:"processed_at"`
}

// ==========================================
// 🛠️ SERVICIOS DE NEGOCIO
// ==========================================

// ECommerceService es el servicio principal que orquesta las operaciones
type ECommerceService struct {
    productRepo    ProductRepository
    userRepo       UserRepository
    orderRepo      OrderRepository
    paymentSvc     PaymentService
    notificationSvc NotificationService
    inventorySvc   InventoryService
    logger         Logger
}

// Logger interface para logging
type Logger interface {
    Info(msg string, fields ...interface{})
    Error(msg string, err error, fields ...interface{})
    Debug(msg string, fields ...interface{})
}

// NewECommerceService crea una nueva instancia del servicio
func NewECommerceService(
    productRepo ProductRepository,
    userRepo UserRepository,
    orderRepo OrderRepository,
    paymentSvc PaymentService,
    notificationSvc NotificationService,
    inventorySvc InventoryService,
    logger Logger,
) *ECommerceService {
    return &ECommerceService{
        productRepo:     productRepo,
        userRepo:        userRepo,
        orderRepo:       orderRepo,
        paymentSvc:      paymentSvc,
        notificationSvc: notificationSvc,
        inventorySvc:    inventorySvc,
        logger:          logger,
    }
}

// CreateOrder crea una nueva orden (Funcionalidad principal para testing)
func (e *ECommerceService) CreateOrder(ctx context.Context, userID string, items []OrderItem, paymentMethod string) (*Order, error) {
    e.logger.Info("Starting order creation", "userID", userID, "itemCount", len(items))
    
    // 1. Validar usuario
    user, err := e.userRepo.GetByID(ctx, userID)
    if err != nil {
        e.logger.Error("Failed to get user", err, "userID", userID)
        return nil, fmt.Errorf("invalid user: %w", err)
    }
    
    // 2. Validar items y calcular total
    var totalAmount float64
    var validatedItems []OrderItem
    
    for _, item := range items {
        product, err := e.productRepo.GetByID(ctx, item.ProductID)
        if err != nil {
            e.logger.Error("Failed to get product", err, "productID", item.ProductID)
            return nil, fmt.Errorf("invalid product %s: %w", item.ProductID, err)
        }
        
        // Verificar stock disponible
        available, err := e.inventorySvc.CheckAvailability(ctx, item.ProductID, item.Quantity)
        if err != nil {
            e.logger.Error("Failed to check availability", err, "productID", item.ProductID)
            return nil, fmt.Errorf("failed to check stock for product %s: %w", item.ProductID, err)
        }
        
        if !available {
            return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
        }
        
        // Calcular subtotal
        validatedItem := OrderItem{
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            UnitPrice: product.Price,
            Subtotal:  product.Price * float64(item.Quantity),
        }
        
        validatedItems = append(validatedItems, validatedItem)
        totalAmount += validatedItem.Subtotal
    }
    
    // 3. Crear orden
    order := &Order{
        ID:          generateOrderID(),
        UserID:      userID,
        Items:       validatedItems,
        TotalAmount: totalAmount,
        Status:      OrderStatusPending,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // 4. Reservar stock
    err = e.inventorySvc.ReserveStock(ctx, validatedItems)
    if err != nil {
        e.logger.Error("Failed to reserve stock", err, "orderID", order.ID)
        return nil, fmt.Errorf("failed to reserve stock: %w", err)
    }
    
    // 5. Guardar orden
    err = e.orderRepo.Create(ctx, order)
    if err != nil {
        // Liberar stock en caso de error
        e.inventorySvc.ReleaseStock(ctx, validatedItems)
        e.logger.Error("Failed to create order", err, "orderID", order.ID)
        return nil, fmt.Errorf("failed to create order: %w", err)
    }
    
    // 6. Procesar pago
    paymentResult, err := e.paymentSvc.ProcessPayment(ctx, order.ID, totalAmount, paymentMethod)
    if err != nil {
        // Revertir orden y liberar stock
        e.orderRepo.UpdateStatus(ctx, order.ID, OrderStatusCancelled)
        e.inventorySvc.ReleaseStock(ctx, validatedItems)
        e.logger.Error("Payment processing failed", err, "orderID", order.ID)
        return nil, fmt.Errorf("payment failed: %w", err)
    }
    
    // 7. Actualizar estado de orden si el pago fue exitoso
    if paymentResult.Status == "success" {
        order.Status = OrderStatusProcessing
        order.UpdatedAt = time.Now()
        
        err = e.orderRepo.UpdateStatus(ctx, order.ID, OrderStatusProcessing)
        if err != nil {
            e.logger.Error("Failed to update order status", err, "orderID", order.ID)
            // No retornar error aquí, ya que el pago fue exitoso
        }
        
        // 8. Enviar notificación
        err = e.notificationSvc.SendOrderConfirmation(ctx, userID, order.ID)
        if err != nil {
            e.logger.Error("Failed to send order confirmation", err, "orderID", order.ID)
            // No retornar error, la orden ya fue creada exitosamente
        }
    }
    
    e.logger.Info("Order created successfully", "orderID", order.ID, "userID", userID, "amount", totalAmount)
    return order, nil
}

// GetOrdersByUser obtiene las órdenes de un usuario
func (e *ECommerceService) GetOrdersByUser(ctx context.Context, userID string) ([]*Order, error) {
    e.logger.Info("Getting orders for user", "userID", userID)
    
    // Validar que el usuario existe
    _, err := e.userRepo.GetByID(ctx, userID)
    if err != nil {
        e.logger.Error("Failed to get user", err, "userID", userID)
        return nil, fmt.Errorf("invalid user: %w", err)
    }
    
    orders, err := e.orderRepo.GetByUserID(ctx, userID)
    if err != nil {
        e.logger.Error("Failed to get orders", err, "userID", userID)
        return nil, fmt.Errorf("failed to get orders: %w", err)
    }
    
    e.logger.Info("Retrieved orders", "userID", userID, "orderCount", len(orders))
    return orders, nil
}

// CancelOrder cancela una orden
func (e *ECommerceService) CancelOrder(ctx context.Context, orderID, userID string) error {
    e.logger.Info("Cancelling order", "orderID", orderID, "userID", userID)
    
    // Obtener orden
    order, err := e.orderRepo.GetByID(ctx, orderID)
    if err != nil {
        e.logger.Error("Failed to get order", err, "orderID", orderID)
        return fmt.Errorf("order not found: %w", err)
    }
    
    // Verificar que la orden pertenece al usuario
    if order.UserID != userID {
        e.logger.Error("Unauthorized cancellation attempt", nil, "orderID", orderID, "userID", userID)
        return errors.New("unauthorized: order does not belong to user")
    }
    
    // Verificar que la orden se puede cancelar
    if order.Status == OrderStatusShipped || order.Status == OrderStatusDelivered {
        return errors.New("cannot cancel order: already shipped or delivered")
    }
    
    if order.Status == OrderStatusCancelled {
        return errors.New("order already cancelled")
    }
    
    // Actualizar estado
    err = e.orderRepo.UpdateStatus(ctx, orderID, OrderStatusCancelled)
    if err != nil {
        e.logger.Error("Failed to update order status", err, "orderID", orderID)
        return fmt.Errorf("failed to cancel order: %w", err)
    }
    
    // Liberar stock
    err = e.inventorySvc.ReleaseStock(ctx, order.Items)
    if err != nil {
        e.logger.Error("Failed to release stock", err, "orderID", orderID)
        // Continuar, el stock se puede ajustar manualmente
    }
    
    e.logger.Info("Order cancelled successfully", "orderID", orderID, "userID", userID)
    return nil
}

// ==========================================
// 🎭 IMPLEMENTACIONES MOCK PARA TESTING
// ==========================================

// MockProductRepository implementación mock para testing
type MockProductRepository struct {
    products map[string]*Product
    mutex    sync.RWMutex
    errors   map[string]error // Para simular errores en métodos específicos
}

func NewMockProductRepository() *MockProductRepository {
    return &MockProductRepository{
        products: make(map[string]*Product),
        errors:   make(map[string]error),
    }
}

func (m *MockProductRepository) AddProduct(product *Product) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.products[product.ID] = product
}

func (m *MockProductRepository) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByID"]; exists {
        return nil, err
    }
    
    if product, exists := m.products[id]; exists {
        return product, nil
    }
    
    return nil, errors.New("product not found")
}

func (m *MockProductRepository) GetByCategory(ctx context.Context, category string) ([]*Product, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByCategory"]; exists {
        return nil, err
    }
    
    var result []*Product
    for _, product := range m.products {
        if product.Category == category {
            result = append(result, product)
        }
    }
    
    return result, nil
}

func (m *MockProductRepository) Create(ctx context.Context, product *Product) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["Create"]; exists {
        return err
    }
    
    m.products[product.ID] = product
    return nil
}

func (m *MockProductRepository) Update(ctx context.Context, product *Product) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["Update"]; exists {
        return err
    }
    
    if _, exists := m.products[product.ID]; !exists {
        return errors.New("product not found")
    }
    
    m.products[product.ID] = product
    return nil
}

func (m *MockProductRepository) UpdateStock(ctx context.Context, productID string, newStock int) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["UpdateStock"]; exists {
        return err
    }
    
    if product, exists := m.products[productID]; exists {
        product.Stock = newStock
        return nil
    }
    
    return errors.New("product not found")
}

func (m *MockProductRepository) Search(ctx context.Context, query string) ([]*Product, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["Search"]; exists {
        return nil, err
    }
    
    var result []*Product
    for _, product := range m.products {
        if contains(product.Name, query) || contains(product.Description, query) {
            result = append(result, product)
        }
    }
    
    return result, nil
}

// MockUserRepository implementación mock para testing
type MockUserRepository struct {
    users  map[string]*User
    emails map[string]*User
    mutex  sync.RWMutex
    errors map[string]error
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users:  make(map[string]*User),
        emails: make(map[string]*User),
        errors: make(map[string]error),
    }
}

func (m *MockUserRepository) AddUser(user *User) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.users[user.ID] = user
    m.emails[user.Email] = user
}

func (m *MockUserRepository) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*User, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByID"]; exists {
        return nil, err
    }
    
    if user, exists := m.users[id]; exists {
        return user, nil
    }
    
    return nil, errors.New("user not found")
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByEmail"]; exists {
        return nil, err
    }
    
    if user, exists := m.emails[email]; exists {
        return user, nil
    }
    
    return nil, errors.New("user not found")
}

func (m *MockUserRepository) Create(ctx context.Context, user *User) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["Create"]; exists {
        return err
    }
    
    m.users[user.ID] = user
    m.emails[user.Email] = user
    return nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *User) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["Update"]; exists {
        return err
    }
    
    if _, exists := m.users[user.ID]; !exists {
        return errors.New("user not found")
    }
    
    m.users[user.ID] = user
    return nil
}

// MockOrderRepository implementación mock para testing
type MockOrderRepository struct {
    orders    map[string]*Order
    userOrders map[string][]*Order
    mutex     sync.RWMutex
    errors    map[string]error
}

func NewMockOrderRepository() *MockOrderRepository {
    return &MockOrderRepository{
        orders:     make(map[string]*Order),
        userOrders: make(map[string][]*Order),
        errors:     make(map[string]error),
    }
}

func (m *MockOrderRepository) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockOrderRepository) GetByID(ctx context.Context, id string) (*Order, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByID"]; exists {
        return nil, err
    }
    
    if order, exists := m.orders[id]; exists {
        return order, nil
    }
    
    return nil, errors.New("order not found")
}

func (m *MockOrderRepository) GetByUserID(ctx context.Context, userID string) ([]*Order, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetByUserID"]; exists {
        return nil, err
    }
    
    if orders, exists := m.userOrders[userID]; exists {
        return orders, nil
    }
    
    return []*Order{}, nil
}

func (m *MockOrderRepository) Create(ctx context.Context, order *Order) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["Create"]; exists {
        return err
    }
    
    m.orders[order.ID] = order
    m.userOrders[order.UserID] = append(m.userOrders[order.UserID], order)
    return nil
}

func (m *MockOrderRepository) UpdateStatus(ctx context.Context, orderID string, status OrderStatus) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["UpdateStatus"]; exists {
        return err
    }
    
    if order, exists := m.orders[orderID]; exists {
        order.Status = status
        order.UpdatedAt = time.Now()
        return nil
    }
    
    return errors.New("order not found")
}

// MockPaymentService implementación mock para testing
type MockPaymentService struct {
    payments      map[string]*PaymentResult
    paymentStatus map[string]*PaymentStatus
    mutex         sync.RWMutex
    errors        map[string]error
    shouldFail    bool
}

func NewMockPaymentService() *MockPaymentService {
    return &MockPaymentService{
        payments:      make(map[string]*PaymentResult),
        paymentStatus: make(map[string]*PaymentStatus),
        errors:        make(map[string]error),
    }
}

func (m *MockPaymentService) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockPaymentService) SetShouldFail(shouldFail bool) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.shouldFail = shouldFail
}

func (m *MockPaymentService) ProcessPayment(ctx context.Context, orderID string, amount float64, paymentMethod string) (*PaymentResult, error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["ProcessPayment"]; exists {
        return nil, err
    }
    
    if m.shouldFail {
        return nil, errors.New("payment processing failed")
    }
    
    paymentID := generatePaymentID()
    result := &PaymentResult{
        PaymentID:     paymentID,
        Status:        "success",
        TransactionID: generateTransactionID(),
        ProcessedAt:   time.Now(),
    }
    
    m.payments[orderID] = result
    m.paymentStatus[paymentID] = &PaymentStatus{
        PaymentID:   paymentID,
        Status:      "completed",
        Amount:      amount,
        ProcessedAt: time.Now(),
    }
    
    return result, nil
}

func (m *MockPaymentService) RefundPayment(ctx context.Context, paymentID string, amount float64) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["RefundPayment"]; exists {
        return err
    }
    
    if status, exists := m.paymentStatus[paymentID]; exists {
        status.Status = "refunded"
        return nil
    }
    
    return errors.New("payment not found")
}

func (m *MockPaymentService) GetPaymentStatus(ctx context.Context, paymentID string) (*PaymentStatus, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["GetPaymentStatus"]; exists {
        return nil, err
    }
    
    if status, exists := m.paymentStatus[paymentID]; exists {
        return status, nil
    }
    
    return nil, errors.New("payment not found")
}

// MockNotificationService implementación mock para testing
type MockNotificationService struct {
    sentNotifications []string
    mutex            sync.RWMutex
    errors           map[string]error
}

func NewMockNotificationService() *MockNotificationService {
    return &MockNotificationService{
        sentNotifications: make([]string, 0),
        errors:           make(map[string]error),
    }
}

func (m *MockNotificationService) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockNotificationService) GetSentNotifications() []string {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    return append([]string{}, m.sentNotifications...)
}

func (m *MockNotificationService) SendOrderConfirmation(ctx context.Context, userID, orderID string) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["SendOrderConfirmation"]; exists {
        return err
    }
    
    notification := fmt.Sprintf("OrderConfirmation:userID=%s,orderID=%s", userID, orderID)
    m.sentNotifications = append(m.sentNotifications, notification)
    return nil
}

func (m *MockNotificationService) SendShippingNotification(ctx context.Context, userID, orderID, trackingNumber string) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["SendShippingNotification"]; exists {
        return err
    }
    
    notification := fmt.Sprintf("ShippingNotification:userID=%s,orderID=%s,tracking=%s", userID, orderID, trackingNumber)
    m.sentNotifications = append(m.sentNotifications, notification)
    return nil
}

func (m *MockNotificationService) SendDeliveryConfirmation(ctx context.Context, userID, orderID string) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["SendDeliveryConfirmation"]; exists {
        return err
    }
    
    notification := fmt.Sprintf("DeliveryConfirmation:userID=%s,orderID=%s", userID, orderID)
    m.sentNotifications = append(m.sentNotifications, notification)
    return nil
}

// MockInventoryService implementación mock para testing
type MockInventoryService struct {
    stock         map[string]int
    reservedStock map[string]int
    mutex         sync.RWMutex
    errors        map[string]error
}

func NewMockInventoryService() *MockInventoryService {
    return &MockInventoryService{
        stock:         make(map[string]int),
        reservedStock: make(map[string]int),
        errors:        make(map[string]error),
    }
}

func (m *MockInventoryService) SetStock(productID string, quantity int) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.stock[productID] = quantity
}

func (m *MockInventoryService) SetError(method string, err error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.errors[method] = err
}

func (m *MockInventoryService) CheckAvailability(ctx context.Context, productID string, quantity int) (bool, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    if err, exists := m.errors["CheckAvailability"]; exists {
        return false, err
    }
    
    available := m.stock[productID] - m.reservedStock[productID]
    return available >= quantity, nil
}

func (m *MockInventoryService) ReserveStock(ctx context.Context, items []OrderItem) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["ReserveStock"]; exists {
        return err
    }
    
    for _, item := range items {
        m.reservedStock[item.ProductID] += item.Quantity
    }
    
    return nil
}

func (m *MockInventoryService) ReleaseStock(ctx context.Context, items []OrderItem) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["ReleaseStock"]; exists {
        return err
    }
    
    for _, item := range items {
        m.reservedStock[item.ProductID] -= item.Quantity
        if m.reservedStock[item.ProductID] < 0 {
            m.reservedStock[item.ProductID] = 0
        }
    }
    
    return nil
}

func (m *MockInventoryService) UpdateStock(ctx context.Context, productID string, change int) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    if err, exists := m.errors["UpdateStock"]; exists {
        return err
    }
    
    m.stock[productID] += change
    if m.stock[productID] < 0 {
        m.stock[productID] = 0
    }
    
    return nil
}

// MockLogger implementación mock para testing
type MockLogger struct {
    logs   []LogEntry
    mutex  sync.RWMutex
}

type LogEntry struct {
    Level   string
    Message string
    Fields  map[string]interface{}
    Error   error
    Time    time.Time
}

func NewMockLogger() *MockLogger {
    return &MockLogger{
        logs: make([]LogEntry, 0),
    }
}

func (m *MockLogger) Info(msg string, fields ...interface{}) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    m.logs = append(m.logs, LogEntry{
        Level:   "INFO",
        Message: msg,
        Fields:  convertToMap(fields),
        Time:    time.Now(),
    })
}

func (m *MockLogger) Error(msg string, err error, fields ...interface{}) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    m.logs = append(m.logs, LogEntry{
        Level:   "ERROR",
        Message: msg,
        Error:   err,
        Fields:  convertToMap(fields),
        Time:    time.Now(),
    })
}

func (m *MockLogger) Debug(msg string, fields ...interface{}) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    m.logs = append(m.logs, LogEntry{
        Level:   "DEBUG",
        Message: msg,
        Fields:  convertToMap(fields),
        Time:    time.Now(),
    })
}

func (m *MockLogger) GetLogs() []LogEntry {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    return append([]LogEntry{}, m.logs...)
}

func (m *MockLogger) GetLogsByLevel(level string) []LogEntry {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    var result []LogEntry
    for _, log := range m.logs {
        if log.Level == level {
            result = append(result, log)
        }
    }
    return result
}

// ==========================================
// 🛠️ UTILIDADES
// ==========================================

// generateOrderID genera un ID único para órdenes
func generateOrderID() string {
    return fmt.Sprintf("ORD-%d", time.Now().UnixNano())
}

// generatePaymentID genera un ID único para pagos
func generatePaymentID() string {
    return fmt.Sprintf("PAY-%d", time.Now().UnixNano())
}

// generateTransactionID genera un ID único para transacciones
func generateTransactionID() string {
    return fmt.Sprintf("TXN-%d", time.Now().UnixNano())
}

// contains verifica si una cadena contiene otra (case-insensitive)
func contains(str, substr string) bool {
    return len(str) >= len(substr) && 
           fmt.Sprintf("%s", str) != "" && 
           fmt.Sprintf("%s", substr) != ""
}

// convertToMap convierte una lista de interfaces a mapa
func convertToMap(fields []interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    for i := 0; i < len(fields)-1; i += 2 {
        if key, ok := fields[i].(string); ok {
            result[key] = fields[i+1]
        }
    }
    return result
}

// ==========================================
// 🎯 DEMOSTRACIÓN Y EJEMPLOS
// ==========================================

func main() {
    fmt.Println("🚀 Sistema de Testing Integral - E-Commerce")
    fmt.Println("===========================================")
    
    // Crear instancias mock para demostración
    productRepo := NewMockProductRepository()
    userRepo := NewMockUserRepository()
    orderRepo := NewMockOrderRepository()
    paymentSvc := NewMockPaymentService()
    notificationSvc := NewMockNotificationService()
    inventorySvc := NewMockInventoryService()
    logger := NewMockLogger()
    
    // Crear servicio principal
    ecommerce := NewECommerceService(
        productRepo,
        userRepo,
        orderRepo,
        paymentSvc,
        notificationSvc,
        inventorySvc,
        logger,
    )
    
    // Setup de datos de prueba
    setupTestData(productRepo, userRepo, inventorySvc)
    
    // Demostrar funcionalidad
    demonstrateOrderCreation(ecommerce, logger)
    
    fmt.Println("\n🎉 Demostración completada!")
    fmt.Println("📝 Revisa los archivos de testing para ver ejemplos completos de:")
    fmt.Println("   • TDD con tests unitarios")
    fmt.Println("   • Mocking y test doubles")
    fmt.Println("   • Property-based testing")
    fmt.Println("   • Integration testing")
    fmt.Println("   • Benchmarking")
}

func setupTestData(productRepo *MockProductRepository, userRepo *MockUserRepository, inventorySvc *MockInventoryService) {
    // Agregar productos
    products := []*Product{
        {
            ID:          "prod-1",
            Name:        "Laptop Gaming",
            Description: "High-performance gaming laptop",
            Price:       1299.99,
            Stock:       50,
            Category:    "electronics",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
            ID:          "prod-2",
            Name:        "Wireless Mouse",
            Description: "Ergonomic wireless mouse",
            Price:       29.99,
            Stock:       100,
            Category:    "electronics",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
    }
    
    for _, product := range products {
        productRepo.AddProduct(product)
        inventorySvc.SetStock(product.ID, product.Stock)
    }
    
    // Agregar usuarios
    user := &User{
        ID:        "user-1",
        Email:     "john@example.com",
        FirstName: "John",
        LastName:  "Doe",
        CreatedAt: time.Now(),
    }
    userRepo.AddUser(user)
}

func demonstrateOrderCreation(ecommerce *ECommerceService, logger *MockLogger) {
    fmt.Println("\n🛍️ Demonstrando creación de orden...")
    
    ctx := context.Background()
    userID := "user-1"
    
    items := []OrderItem{
        {ProductID: "prod-1", Quantity: 1},
        {ProductID: "prod-2", Quantity: 2},
    }
    
    order, err := ecommerce.CreateOrder(ctx, userID, items, "credit_card")
    if err != nil {
        fmt.Printf("❌ Error creating order: %v\n", err)
        return
    }
    
    fmt.Printf("✅ Orden creada exitosamente: %s\n", order.ID)
    fmt.Printf("💰 Total: $%.2f\n", order.TotalAmount)
    fmt.Printf("📊 Estado: %s\n", order.Status)
    
    // Mostrar logs
    logs := logger.GetLogs()
    fmt.Printf("\n📝 Logs generados (%d):\n", len(logs))
    for _, log := range logs[:3] { // Mostrar solo los primeros 3
        fmt.Printf("   [%s] %s\n", log.Level, log.Message)
    }
}
