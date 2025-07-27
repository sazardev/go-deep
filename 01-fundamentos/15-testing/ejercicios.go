// 🧪 Ejercicios: Testing en Go
// =============================================
// Descripción: Ejercicios prácticos para dominar testing en Go
// Dificultad: Intermedio a Avanzado
// Temas: Tests unitarios, mocking, integración, benchmarks, TDD

package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// =============================================================================
// 📦 EJERCICIO 1: Sistema de Biblioteca - Tests Unitarios Básicos
// =============================================================================

// Book representa un libro en la biblioteca
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Year     int    `json:"year"`
	Pages    int    `json:"pages"`
	Genre    string `json:"genre"`
	Language string `json:"language"`
}

// BookValidator valida datos de libros
type BookValidator struct{}

// ValidateBook valida un libro completo
func (bv *BookValidator) ValidateBook(book *Book) []string {
	var errors []string

	if book.Title == "" {
		errors = append(errors, "title cannot be empty")
	}

	if book.Author == "" {
		errors = append(errors, "author cannot be empty")
	}

	if !bv.ValidateISBN(book.ISBN) {
		errors = append(errors, "invalid ISBN format")
	}

	if book.Year < 1000 || book.Year > time.Now().Year() {
		errors = append(errors, "invalid year")
	}

	if book.Pages <= 0 {
		errors = append(errors, "pages must be positive")
	}

	if book.Genre == "" {
		errors = append(errors, "genre cannot be empty")
	}

	return errors
}

// ValidateISBN valida formato ISBN-10 o ISBN-13
func (bv *BookValidator) ValidateISBN(isbn string) bool {
	// Remove hyphens and spaces
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")

	// Check length
	if len(isbn) != 10 && len(isbn) != 13 {
		return false
	}

	// Validate ISBN-10
	if len(isbn) == 10 {
		return bv.validateISBN10(isbn)
	}

	// Validate ISBN-13
	return bv.validateISBN13(isbn)
}

func (bv *BookValidator) validateISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		if isbn[i] < '0' || isbn[i] > '9' {
			return false
		}
		digit := int(isbn[i] - '0')
		sum += digit * (10 - i)
	}

	// Check digit can be 0-9 or X
	var checkDigit int
	if isbn[9] == 'X' || isbn[9] == 'x' {
		checkDigit = 10
	} else if isbn[9] >= '0' && isbn[9] <= '9' {
		checkDigit = int(isbn[9] - '0')
	} else {
		return false
	}

	sum += checkDigit
	return sum%11 == 0
}

func (bv *BookValidator) validateISBN13(isbn string) bool {
	if len(isbn) != 13 {
		return false
	}

	sum := 0
	for i := 0; i < 13; i++ {
		if isbn[i] < '0' || isbn[i] > '9' {
			return false
		}
		digit := int(isbn[i] - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	return sum%10 == 0
}

// 🎯 EJERCICIO 1: Crear tests completos para BookValidator
// Tu tarea:
// 1. Crear book_validator_test.go
// 2. Usar table-driven tests para ValidateBook
// 3. Casos edge para ValidateISBN (ISBN-10 y ISBN-13)
// 4. Tests para casos válidos e inválidos
// 5. Usar subtests para organizar mejor

// =============================================================================
// 📦 EJERCICIO 2: Sistema de Inventario - Mocking y Dependency Injection
// =============================================================================

// Product representa un producto
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

// ProductRepository define operaciones de persistencia
type ProductRepository interface {
	Save(product *Product) error
	FindByID(id int) (*Product, error)
	FindByCategory(category string) ([]*Product, error)
	UpdateStock(id int, quantity int) error
	Delete(id int) error
}

// NotificationService define servicio de notificaciones
type NotificationService interface {
	SendLowStockAlert(product *Product) error
	SendOutOfStockAlert(product *Product) error
}

// PriceCalculator calcula precios con descuentos
type PriceCalculator interface {
	CalculateDiscount(price float64, category string) float64
	CalculateTax(price float64) float64
}

// InventoryService maneja lógica de inventario
type InventoryService struct {
	repo              ProductRepository
	notification      NotificationService
	calculator        PriceCalculator
	lowStockThreshold int
}

// NewInventoryService constructor
func NewInventoryService(
	repo ProductRepository,
	notification NotificationService,
	calculator PriceCalculator,
) *InventoryService {
	return &InventoryService{
		repo:              repo,
		notification:      notification,
		calculator:        calculator,
		lowStockThreshold: 10,
	}
}

// AddProduct añade nuevo producto
func (s *InventoryService) AddProduct(product *Product) error {
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("product price must be positive")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.repo.Save(product)
}

// UpdateStock actualiza stock y envía alertas si es necesario
func (s *InventoryService) UpdateStock(productID int, quantity int) error {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	newStock := product.Stock + quantity
	if newStock < 0 {
		return errors.New("insufficient stock")
	}

	err = s.repo.UpdateStock(productID, quantity)
	if err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}

	// Check for alerts
	if newStock == 0 {
		return s.notification.SendOutOfStockAlert(product)
	} else if newStock <= s.lowStockThreshold {
		return s.notification.SendLowStockAlert(product)
	}

	return nil
}

// GetFinalPrice calcula precio final con descuentos y taxes
func (s *InventoryService) GetFinalPrice(productID int) (float64, error) {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return 0, fmt.Errorf("product not found: %w", err)
	}

	discountedPrice := s.calculator.CalculateDiscount(product.Price, product.Category)
	finalPrice := s.calculator.CalculateTax(discountedPrice)

	return finalPrice, nil
}

// 🎯 EJERCICIO 2: Crear tests con mocks
// Tu tarea:
// 1. Crear mocks manuales para todas las interfaces
// 2. Testear AddProduct con validaciones
// 3. Testear UpdateStock con diferentes escenarios de alertas
// 4. Testear GetFinalPrice con mocks que retornen valores específicos
// 5. Verificar que los mocks fueron llamados correctamente

// =============================================================================
// 📦 EJERCICIO 3: Sistema de Carritos - Concurrencia y Race Conditions
// =============================================================================

// CartItem representa un item en el carrito
type CartItem struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

// ShoppingCart representa un carrito de compras thread-safe
type ShoppingCart struct {
	mu    sync.RWMutex
	items map[int]*CartItem
	total float64
}

// NewShoppingCart constructor
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: make(map[int]*CartItem),
	}
}

// AddItem añade item al carrito
func (c *ShoppingCart) AddItem(productID int, name string, price float64, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if existingItem, exists := c.items[productID]; exists {
		existingItem.Quantity += quantity
	} else {
		c.items[productID] = &CartItem{
			ProductID: productID,
			Name:      name,
			Price:     price,
			Quantity:  quantity,
		}
	}

	c.recalculateTotal()
	return nil
}

// RemoveItem remueve item del carrito
func (c *ShoppingCart) RemoveItem(productID int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[productID]; !exists {
		return errors.New("item not found in cart")
	}

	delete(c.items, productID)
	c.recalculateTotal()
	return nil
}

// UpdateQuantity actualiza cantidad de un item
func (c *ShoppingCart) UpdateQuantity(productID int, quantity int) error {
	if quantity <= 0 {
		return c.RemoveItem(productID)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.items[productID]
	if !exists {
		return errors.New("item not found in cart")
	}

	item.Quantity = quantity
	c.recalculateTotal()
	return nil
}

// GetTotal retorna total del carrito
func (c *ShoppingCart) GetTotal() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.total
}

// GetItems retorna copia de los items
func (c *ShoppingCart) GetItems() map[int]*CartItem {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[int]*CartItem)
	for k, v := range c.items {
		items[k] = &CartItem{
			ProductID: v.ProductID,
			Name:      v.Name,
			Price:     v.Price,
			Quantity:  v.Quantity,
		}
	}
	return items
}

// GetItemCount retorna número total de items
func (c *ShoppingCart) GetItemCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	count := 0
	for _, item := range c.items {
		count += item.Quantity
	}
	return count
}

// Clear limpia el carrito
func (c *ShoppingCart) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[int]*CartItem)
	c.total = 0
}

// recalculateTotal recalcula el total (debe llamarse con lock)
func (c *ShoppingCart) recalculateTotal() {
	c.total = 0
	for _, item := range c.items {
		c.total += item.Price * float64(item.Quantity)
	}
}

// 🎯 EJERCICIO 3: Tests de concurrencia
// Tu tarea:
// 1. Testear operaciones básicas del carrito
// 2. Crear tests de concurrencia con goroutines
// 3. Verificar que no hay race conditions
// 4. Usar go test -race para detectar races
// 5. Benchmarks para operaciones críticas

// =============================================================================
// 📦 EJERCICIO 4: Sistema de Pagos - Integration Testing
// =============================================================================

// PaymentMethod tipos de pago
type PaymentMethod string

const (
	CreditCard   PaymentMethod = "credit_card"
	DebitCard    PaymentMethod = "debit_card"
	PayPal       PaymentMethod = "paypal"
	BankTransfer PaymentMethod = "bank_transfer"
)

// PaymentRequest solicitud de pago
type PaymentRequest struct {
	ID          string            `json:"id"`
	Amount      float64           `json:"amount"`
	Currency    string            `json:"currency"`
	Method      PaymentMethod     `json:"method"`
	CustomerID  string            `json:"customer_id"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata"`
}

// PaymentStatus estado del pago
type PaymentStatus string

const (
	StatusPending   PaymentStatus = "pending"
	StatusCompleted PaymentStatus = "completed"
	StatusFailed    PaymentStatus = "failed"
	StatusCancelled PaymentStatus = "cancelled"
)

// PaymentResponse respuesta del pago
type PaymentResponse struct {
	ID            string            `json:"id"`
	Status        PaymentStatus     `json:"status"`
	Amount        float64           `json:"amount"`
	Currency      string            `json:"currency"`
	TransactionID string            `json:"transaction_id"`
	CreatedAt     time.Time         `json:"created_at"`
	CompletedAt   *time.Time        `json:"completed_at,omitempty"`
	ErrorMessage  string            `json:"error_message,omitempty"`
	Metadata      map[string]string `json:"metadata"`
}

// PaymentGateway interface para diferentes gateways
type PaymentGateway interface {
	ProcessPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
	RefundPayment(ctx context.Context, paymentID string, amount float64) (*PaymentResponse, error)
	GetPaymentStatus(ctx context.Context, paymentID string) (*PaymentResponse, error)
}

// PaymentProcessor procesa pagos usando diferentes gateways
type PaymentProcessor struct {
	gateways map[PaymentMethod]PaymentGateway
	logger   Logger
}

// Logger interface para logging
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, err error, fields ...interface{})
	Warn(msg string, fields ...interface{})
}

// NewPaymentProcessor constructor
func NewPaymentProcessor(logger Logger) *PaymentProcessor {
	return &PaymentProcessor{
		gateways: make(map[PaymentMethod]PaymentGateway),
		logger:   logger,
	}
}

// RegisterGateway registra un gateway para un método de pago
func (p *PaymentProcessor) RegisterGateway(method PaymentMethod, gateway PaymentGateway) {
	p.gateways[method] = gateway
}

// ProcessPayment procesa un pago
func (p *PaymentProcessor) ProcessPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	// Validar request
	if err := p.validatePaymentRequest(req); err != nil {
		return nil, fmt.Errorf("invalid payment request: %w", err)
	}

	// Obtener gateway
	gateway, exists := p.gateways[req.Method]
	if !exists {
		return nil, fmt.Errorf("unsupported payment method: %s", req.Method)
	}

	// Log intento de pago
	p.logger.Info("Processing payment",
		"payment_id", req.ID,
		"amount", req.Amount,
		"method", req.Method,
		"customer_id", req.CustomerID)

	// Procesar pago
	response, err := gateway.ProcessPayment(ctx, req)
	if err != nil {
		p.logger.Error("Payment processing failed", err,
			"payment_id", req.ID,
			"method", req.Method)
		return nil, fmt.Errorf("payment processing failed: %w", err)
	}

	// Log resultado
	if response.Status == StatusCompleted {
		p.logger.Info("Payment completed successfully",
			"payment_id", req.ID,
			"transaction_id", response.TransactionID)
	} else {
		p.logger.Warn("Payment not completed",
			"payment_id", req.ID,
			"status", response.Status,
			"error", response.ErrorMessage)
	}

	return response, nil
}

// validatePaymentRequest valida la solicitud de pago
func (p *PaymentProcessor) validatePaymentRequest(req *PaymentRequest) error {
	if req.ID == "" {
		return errors.New("payment ID cannot be empty")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if req.Currency == "" {
		return errors.New("currency cannot be empty")
	}
	if req.CustomerID == "" {
		return errors.New("customer ID cannot be empty")
	}
	return nil
}

// 🎯 EJERCICIO 4: Integration Testing
// Tu tarea:
// 1. Crear mock gateways para diferentes métodos de pago
// 2. Simular respuestas exitosas y de error
// 3. Testear timeouts y contexts cancelados
// 4. Mock del logger para verificar logs
// 5. Tests end-to-end del flujo completo de pago

// =============================================================================
// 📦 EJERCICIO 5: Sistema de Búsqueda - Benchmarks y Performance
// =============================================================================

// SearchEngine motor de búsqueda de productos
type SearchEngine struct {
	products []Product
	index    map[string][]int // palabra -> lista de IDs de productos
}

// NewSearchEngine constructor
func NewSearchEngine(products []Product) *SearchEngine {
	engine := &SearchEngine{
		products: products,
		index:    make(map[string][]int),
	}
	engine.buildIndex()
	return engine
}

// buildIndex construye índice de búsqueda
func (s *SearchEngine) buildIndex() {
	for i, product := range s.products {
		words := s.extractWords(product.Name + " " + product.Category)
		for _, word := range words {
			word = strings.ToLower(word)
			s.index[word] = append(s.index[word], i)
		}
	}
}

// extractWords extrae palabras de un texto
func (s *SearchEngine) extractWords(text string) []string {
	words := strings.Fields(text)
	var cleaned []string
	for _, word := range words {
		// Remove punctuation and convert to lowercase
		clean := strings.ToLower(strings.Trim(word, ".,!?;:()[]{}\"'"))
		if len(clean) > 0 {
			cleaned = append(cleaned, clean)
		}
	}
	return cleaned
}

// Search busca productos por query
func (s *SearchEngine) Search(query string) []Product {
	words := s.extractWords(query)
	if len(words) == 0 {
		return []Product{}
	}

	// Get product IDs for each word
	var candidateIDs [][]int
	for _, word := range words {
		if ids, exists := s.index[word]; exists {
			candidateIDs = append(candidateIDs, ids)
		}
	}

	if len(candidateIDs) == 0 {
		return []Product{}
	}

	// Find intersection of all candidate lists
	resultIDs := s.intersectLists(candidateIDs)

	// Convert IDs to products
	var results []Product
	for _, id := range resultIDs {
		if id < len(s.products) {
			results = append(results, s.products[id])
		}
	}

	return results
}

// intersectLists encuentra la intersección de múltiples listas
func (s *SearchEngine) intersectLists(lists [][]int) []int {
	if len(lists) == 0 {
		return []int{}
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// Sort lists by length (smallest first)
	sort.Slice(lists, func(i, j int) bool {
		return len(lists[i]) < len(lists[j])
	})

	result := make([]int, len(lists[0]))
	copy(result, lists[0])

	for i := 1; i < len(lists); i++ {
		result = s.intersectTwo(result, lists[i])
		if len(result) == 0 {
			break
		}
	}

	return result
}

// intersectTwo encuentra intersección de dos listas ordenadas
func (s *SearchEngine) intersectTwo(list1, list2 []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		if list1[i] == list2[j] {
			result = append(result, list1[i])
			i++
			j++
		} else if list1[i] < list2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

// SearchByCategory busca productos por categoría
func (s *SearchEngine) SearchByCategory(category string) []Product {
	var results []Product
	category = strings.ToLower(category)

	for _, product := range s.products {
		if strings.ToLower(product.Category) == category {
			results = append(results, product)
		}
	}

	return results
}

// SearchByPriceRange busca productos en rango de precio
func (s *SearchEngine) SearchByPriceRange(minPrice, maxPrice float64) []Product {
	var results []Product

	for _, product := range s.products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			results = append(results, product)
		}
	}

	return results
}

// FuzzySearch búsqueda aproximada usando distancia de Levenshtein
func (s *SearchEngine) FuzzySearch(query string, maxDistance int) []Product {
	query = strings.ToLower(query)
	var results []Product

	for _, product := range s.products {
		name := strings.ToLower(product.Name)
		if s.levenshteinDistance(query, name) <= maxDistance {
			results = append(results, product)
		}
	}

	return results
}

// levenshteinDistance calcula distancia de Levenshtein
func (s *SearchEngine) levenshteinDistance(s1, s2 string) int {
	len1, len2 := len(s1), len(s2)
	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len1][len2]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}

// 🎯 EJERCICIO 5: Benchmarks de performance
// Tu tarea:
// 1. Crear benchmarks para diferentes tipos de búsqueda
// 2. Comparar performance con diferentes tamaños de datos
// 3. Benchmark de construcción del índice
// 4. Benchmark de búsqueda fuzzy vs exacta
// 5. Optimizar y re-benchmarkear

// =============================================================================
// 📦 EJERCICIO 6: API REST - HTTP Testing
// =============================================================================

// APIServer servidor HTTP para la API
type APIServer struct {
	inventory *InventoryService
	processor *PaymentProcessor
	search    *SearchEngine
}

// NewAPIServer constructor
func NewAPIServer(
	inventory *InventoryService,
	processor *PaymentProcessor,
	search *SearchEngine,
) *APIServer {
	return &APIServer{
		inventory: inventory,
		processor: processor,
		search:    search,
	}
}

// ProductResponse respuesta de producto
type ProductResponse struct {
	Product *Product `json:"product"`
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
}

// SearchResponse respuesta de búsqueda
type SearchResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Query    string    `json:"query"`
}

// ErrorResponse respuesta de error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// HealthResponse respuesta de health check
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// GetHealth endpoint de health check
func (s *APIServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// En un caso real usarías json.NewEncoder(w).Encode(response)
	_ = response // Avoid unused variable error
}

// GetProduct obtiene producto por ID
func (s *APIServer) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path (simplificado)
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.writeErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// En implementación real usarías el repository
	product := &Product{
		ID:    id,
		Name:  "Sample Product",
		Price: 99.99,
		Stock: 10,
	}

	response := ProductResponse{
		Product: product,
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(response)
	_ = response // Avoid unused variable error
}

// SearchProducts busca productos
func (s *APIServer) SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		s.writeErrorResponse(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// En implementación real usarías s.search.Search(query)
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 20.99},
	}

	response := SearchResponse{
		Products: products,
		Total:    len(products),
		Query:    query,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(response)
	_ = response // Avoid unused variable error
}

// writeErrorResponse escribe respuesta de error
func (s *APIServer) writeErrorResponse(w http.ResponseWriter, message string, code int) {
	response := ErrorResponse{
		Success: false,
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// json.NewEncoder(w).Encode(response)
	_ = response // Avoid unused variable error
}

// 🎯 EJERCICIO 6: HTTP Testing
// Tu tarea:
// 1. Crear tests para cada endpoint usando httptest
// 2. Testear diferentes códigos de respuesta
// 3. Validar headers y contenido JSON
// 4. Tests de middleware y autenticación
// 5. Tests de integración end-to-end

// =============================================================================
// 📦 EJERCICIO 7: Calculator Avanzada - TDD (Test-Driven Development)
// =============================================================================

// 🎯 EJERCICIO 7: Implementa una calculadora usando TDD
// Tu tarea:
// 1. Escribe PRIMERO los tests para cada operación
// 2. Implementa el código mínimo para pasar
// 3. Refactoriza manteniendo los tests verdes
// 4. Añade operaciones avanzadas: potencia, raíz, factorial
// 5. Manejo de errores para división por cero, números negativos

// AdvancedCalculator calculadora con operaciones avanzadas
type AdvancedCalculator struct {
	history []CalculationHistory
}

// CalculationHistory historial de cálculos
type CalculationHistory struct {
	Operation string    `json:"operation"`
	Result    float64   `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

// NewAdvancedCalculator constructor
func NewAdvancedCalculator() *AdvancedCalculator {
	return &AdvancedCalculator{
		history: make([]CalculationHistory, 0),
	}
}

// Add suma dos números
func (c *AdvancedCalculator) Add(a, b float64) float64 {
	result := a + b
	c.addToHistory(fmt.Sprintf("%.2f + %.2f", a, b), result)
	return result
}

// Subtract resta dos números
func (c *AdvancedCalculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addToHistory(fmt.Sprintf("%.2f - %.2f", a, b), result)
	return result
}

// Multiply multiplica dos números
func (c *AdvancedCalculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addToHistory(fmt.Sprintf("%.2f * %.2f", a, b), result)
	return result
}

// Divide divide dos números
func (c *AdvancedCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	c.addToHistory(fmt.Sprintf("%.2f / %.2f", a, b), result)
	return result, nil
}

// Power calcula a^b
func (c *AdvancedCalculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	c.addToHistory(fmt.Sprintf("%.2f ^ %.2f", base, exponent), result)
	return result
}

// SquareRoot calcula raíz cuadrada
func (c *AdvancedCalculator) SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("square root of negative number")
	}
	result := math.Sqrt(n)
	c.addToHistory(fmt.Sprintf("√%.2f", n), result)
	return result, nil
}

// Factorial calcula factorial
func (c *AdvancedCalculator) Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("factorial of negative number")
	}
	if n > 20 {
		return 0, errors.New("factorial too large")
	}

	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}

	c.addToHistory(fmt.Sprintf("%d!", n), float64(result))
	return result, nil
}

// GetHistory retorna historial de cálculos
func (c *AdvancedCalculator) GetHistory() []CalculationHistory {
	return c.history
}

// ClearHistory limpia el historial
func (c *AdvancedCalculator) ClearHistory() {
	c.history = make([]CalculationHistory, 0)
}

// addToHistory añade cálculo al historial
func (c *AdvancedCalculator) addToHistory(operation string, result float64) {
	c.history = append(c.history, CalculationHistory{
		Operation: operation,
		Result:    result,
		Timestamp: time.Now(),
	})
}

// =============================================================================
// 📦 EJERCICIO 8: Worker Pool - Testing de Concurrencia Avanzada
// =============================================================================

// Job representa un trabajo a procesar
type Job struct {
	ID     int
	Data   interface{}
	Result chan JobResult
}

// JobResult resultado de un trabajo
type JobResult struct {
	ID       int
	Output   interface{}
	Error    error
	Duration time.Duration
}

// WorkerPool pool de workers
type WorkerPool struct {
	workerCount int
	jobQueue    chan Job
	quit        chan bool
	wg          sync.WaitGroup
}

// NewWorkerPool constructor
func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		jobQueue:    make(chan Job, queueSize),
		quit:        make(chan bool),
	}
}

// Start inicia el worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Stop detiene el worker pool
func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
}

// SubmitJob envía trabajo al pool
func (wp *WorkerPool) SubmitJob(job Job) {
	wp.jobQueue <- job
}

// worker procesa trabajos
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case job := <-wp.jobQueue:
			start := time.Now()

			// Simular procesamiento
			var output interface{}
			var err error

			switch data := job.Data.(type) {
			case int:
				// Simular cálculo intensivo
				time.Sleep(time.Millisecond * time.Duration(data))
				output = data * 2
			case string:
				// Simular procesamiento de string
				time.Sleep(time.Millisecond * 10)
				output = strings.ToUpper(data)
			default:
				err = errors.New("unsupported data type")
			}

			duration := time.Since(start)

			result := JobResult{
				ID:       job.ID,
				Output:   output,
				Error:    err,
				Duration: duration,
			}

			job.Result <- result

		case <-wp.quit:
			return
		}
	}
}

// 🎯 EJERCICIO 8: Tests de concurrencia avanzada
// Tu tarea:
// 1. Testear que el worker pool procesa trabajos correctamente
// 2. Verificar que múltiples workers funcionan en paralelo
// 3. Testear shutdown graceful del pool
// 4. Verificar que no hay leaks de goroutines
// 5. Benchmarks de throughput con diferentes configuraciones

// =============================================================================
// 📋 INSTRUCCIONES PARA COMPLETAR LOS EJERCICIOS
// =============================================================================

/*
🎯 GUÍA DE IMPLEMENTACIÓN:

1. **ESTRUCTURA DE ARCHIVOS:**
   ```
   15-testing/
   ├── book_validator_test.go      # Ejercicio 1
   ├── inventory_service_test.go   # Ejercicio 2
   ├── shopping_cart_test.go       # Ejercicio 3
   ├── payment_processor_test.go   # Ejercicio 4
   ├── search_engine_test.go       # Ejercicio 5
   ├── api_server_test.go         # Ejercicio 6
   ├── calculator_test.go         # Ejercicio 7 (TDD)
   ├── worker_pool_test.go        # Ejercicio 8
   ├── integration_test.go        # Tests de integración
   └── coverage.sh               # Script de coverage
   ```

2. **COMANDOS ÚTILES:**
   ```bash
   # Ejecutar todos los tests
   go test -v ./...

   # Tests con race detection
   go test -race ./...

   # Coverage completo
   go test -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out

   # Solo tests rápidos
   go test -short ./...

   # Benchmarks
   go test -bench=. -benchmem ./...

   # Tests específicos
   go test -run TestBookValidator ./...
   ```

3. **CRITERIOS DE EVALUACIÓN:**
   ✅ Tests unitarios completos con table-driven tests
   ✅ Uso correcto de mocks y dependency injection
   ✅ Tests de concurrencia sin race conditions
   ✅ Cobertura de código > 85%
   ✅ Benchmarks implementados correctamente
   ✅ Tests de integración funcionales
   ✅ Aplicación correcta de TDD
   ✅ Manejo adecuado de errores y edge cases

4. **BONUS TRACKS:**
   🚀 Implementar property-based testing
   🚀 Crear contract tests para APIs
   🚀 Configurar CI/CD con GitHub Actions
   🚀 Añadir fuzzing tests
   🚀 Implementar mutation testing

¡Feliz testing! 🧪✨
*/

func main() {
	fmt.Println("🧪 Ejercicios de Testing en Go")
	fmt.Println("==============================")
	fmt.Println("Implementa los tests en archivos separados")
	fmt.Println("Usa 'go test -v ./...' para ejecutar todos los tests")
	fmt.Println("¡Que la fuerza del testing esté contigo! 🚀")
}
