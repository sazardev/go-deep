// 🧪 Soluciones: Testing Avanzado
// Lección 17: TDD, Mocking y Property Testing
// Ejecutar tests con: go test -v
// Ejecutar benchmarks con: go test -bench=.

package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

// ========================================
// Ejercicio 1: TDD - Calculadora Científica
// ========================================

type ScientificCalculator struct {
	history []string
}

func NewScientificCalculator() *ScientificCalculator {
	return &ScientificCalculator{
		history: make([]string, 0),
	}
}

func (sc *ScientificCalculator) Add(a, b float64) float64 {
	result := a + b
	sc.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Subtract(a, b float64) float64 {
	result := a - b
	sc.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Multiply(a, b float64) float64 {
	result := a * b
	sc.addToHistory(fmt.Sprintf("%.2f × %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	sc.addToHistory(fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, result))
	return result, nil
}

func (sc *ScientificCalculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	sc.addToHistory(fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
	return result
}

func (sc *ScientificCalculator) SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number")
	}
	result := math.Sqrt(x)
	sc.addToHistory(fmt.Sprintf("√%.2f = %.2f", x, result))
	return result, nil
}

func (sc *ScientificCalculator) Sin(x float64) float64 {
	result := math.Sin(x)
	sc.addToHistory(fmt.Sprintf("sin(%.2f) = %.2f", x, result))
	return result
}

func (sc *ScientificCalculator) Cos(x float64) float64 {
	result := math.Cos(x)
	sc.addToHistory(fmt.Sprintf("cos(%.2f) = %.2f", x, result))
	return result
}

func (sc *ScientificCalculator) GetHistory() []string {
	result := make([]string, len(sc.history))
	copy(result, sc.history)
	return result
}

func (sc *ScientificCalculator) ClearHistory() {
	sc.history = make([]string, 0)
}

func (sc *ScientificCalculator) addToHistory(operation string) {
	sc.history = append(sc.history, operation)
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: TDD Calculadora Científica ===")

	calc := NewScientificCalculator()

	// Demonstración de uso
	fmt.Printf("Suma: %.2f\n", calc.Add(5, 3))
	fmt.Printf("Resta: %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("Multiplicación: %.2f\n", calc.Multiply(6, 7))

	div, err := calc.Divide(15, 3)
	if err != nil {
		fmt.Printf("Error en división: %v\n", err)
	} else {
		fmt.Printf("División: %.2f\n", div)
	}

	fmt.Printf("Potencia: %.2f\n", calc.Power(2, 3))

	sqrt, err := calc.SquareRoot(16)
	if err != nil {
		fmt.Printf("Error en raíz cuadrada: %v\n", err)
	} else {
		fmt.Printf("Raíz cuadrada: %.2f\n", sqrt)
	}

	fmt.Printf("Seno: %.4f\n", calc.Sin(math.Pi/2))
	fmt.Printf("Coseno: %.4f\n", calc.Cos(0))

	fmt.Println("\nHistorial de operaciones:")
	for i, op := range calc.GetHistory() {
		fmt.Printf("%d. %s\n", i+1, op)
	}
}

// ==========================================
// Ejercicio 2: Mocking - Sistema de Notificaciones
// ==========================================

type EmailSender interface {
	SendEmail(to, subject, body string) error
}

type SMSSender interface {
	SendSMS(phone, message string) error
}

type PushNotificationSender interface {
	SendPushNotification(userID, title, message string) error
}

type NotificationService struct {
	emailSender EmailSender
	smsSender   SMSSender
	pushSender  PushNotificationSender
}

func NewNotificationService(emailSender EmailSender, smsSender SMSSender, pushSender PushNotificationSender) *NotificationService {
	return &NotificationService{
		emailSender: emailSender,
		smsSender:   smsSender,
		pushSender:  pushSender,
	}
}

func (ns *NotificationService) SendWelcomeNotification(userID string, email string, phone string) error {
	// Enviar email de bienvenida
	if err := ns.emailSender.SendEmail(email, "¡Bienvenido!", "Gracias por registrarte en nuestro servicio."); err != nil {
		return fmt.Errorf("failed to send welcome email: %w", err)
	}

	// Enviar SMS de bienvenida
	if err := ns.smsSender.SendSMS(phone, "¡Bienvenido! Tu cuenta ha sido creada exitosamente."); err != nil {
		return fmt.Errorf("failed to send welcome SMS: %w", err)
	}

	// Enviar push notification
	if err := ns.pushSender.SendPushNotification(userID, "¡Bienvenido!", "Tu cuenta está lista para usar."); err != nil {
		return fmt.Errorf("failed to send welcome push notification: %w", err)
	}

	return nil
}

func (ns *NotificationService) SendPasswordResetNotification(userID string, email string) error {
	subject := "Restablecimiento de contraseña"
	body := "Haz clic en el enlace para restablecer tu contraseña."

	return ns.emailSender.SendEmail(email, subject, body)
}

func (ns *NotificationService) SendOrderConfirmation(userID string, email string, phone string, orderID string) error {
	// Email de confirmación
	emailSubject := "Confirmación de pedido"
	emailBody := fmt.Sprintf("Tu pedido %s ha sido confirmado.", orderID)
	if err := ns.emailSender.SendEmail(email, emailSubject, emailBody); err != nil {
		return fmt.Errorf("failed to send order confirmation email: %w", err)
	}

	// SMS de confirmación
	smsMessage := fmt.Sprintf("Pedido %s confirmado. Te notificaremos cuando esté listo.", orderID)
	if err := ns.smsSender.SendSMS(phone, smsMessage); err != nil {
		return fmt.Errorf("failed to send order confirmation SMS: %w", err)
	}

	return nil
}

// Mocks para testing
type EmailSenderMock struct {
	SendEmailCalls []EmailCall
	SendEmailError error
}

type EmailCall struct {
	To      string
	Subject string
	Body    string
}

func (m *EmailSenderMock) SendEmail(to, subject, body string) error {
	m.SendEmailCalls = append(m.SendEmailCalls, EmailCall{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	return m.SendEmailError
}

type SMSSenderMock struct {
	SendSMSCalls []SMSCall
	SendSMSError error
}

type SMSCall struct {
	Phone   string
	Message string
}

func (m *SMSSenderMock) SendSMS(phone, message string) error {
	m.SendSMSCalls = append(m.SendSMSCalls, SMSCall{
		Phone:   phone,
		Message: message,
	})
	return m.SendSMSError
}

type PushNotificationSenderMock struct {
	SendPushCalls []PushCall
	SendPushError error
}

type PushCall struct {
	UserID  string
	Title   string
	Message string
}

func (m *PushNotificationSenderMock) SendPushNotification(userID, title, message string) error {
	m.SendPushCalls = append(m.SendPushCalls, PushCall{
		UserID:  userID,
		Title:   title,
		Message: message,
	})
	return m.SendPushError
}

func ejercicio2() {
	fmt.Println("=== Ejercicio 2: Mocking Sistema de Notificaciones ===")

	// Crear mocks
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	// Crear servicio con mocks
	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Test de bienvenida
	err := service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")
	if err != nil {
		fmt.Printf("Error enviando notificación de bienvenida: %v\n", err)
		return
	}

	fmt.Printf("✅ Notificación de bienvenida enviada\n")
	fmt.Printf("Emails enviados: %d\n", len(emailMock.SendEmailCalls))
	fmt.Printf("SMS enviados: %d\n", len(smsMock.SendSMSCalls))
	fmt.Printf("Push notifications enviadas: %d\n", len(pushMock.SendPushCalls))

	// Verificar calls
	if len(emailMock.SendEmailCalls) > 0 {
		fmt.Printf("Primer email: To=%s, Subject=%s\n",
			emailMock.SendEmailCalls[0].To, emailMock.SendEmailCalls[0].Subject)
	}
}

// ===============================================
// Ejercicio 3: Property Testing - Lista Ordenada
// ===============================================

type SortedList struct {
	items []int
}

func NewSortedList() *SortedList {
	return &SortedList{
		items: make([]int, 0),
	}
}

func (sl *SortedList) Insert(value int) {
	// Búsqueda binaria para encontrar posición de inserción
	pos := sort.SearchInts(sl.items, value)

	// Insertar en la posición correcta
	sl.items = append(sl.items, 0)
	copy(sl.items[pos+1:], sl.items[pos:])
	sl.items[pos] = value
}

func (sl *SortedList) Remove(value int) bool {
	pos := sort.SearchInts(sl.items, value)
	if pos < len(sl.items) && sl.items[pos] == value {
		sl.items = append(sl.items[:pos], sl.items[pos+1:]...)
		return true
	}
	return false
}

func (sl *SortedList) Contains(value int) bool {
	pos := sort.SearchInts(sl.items, value)
	return pos < len(sl.items) && sl.items[pos] == value
}

func (sl *SortedList) Size() int {
	return len(sl.items)
}

func (sl *SortedList) ToSlice() []int {
	result := make([]int, len(sl.items))
	copy(result, sl.items)
	return result
}

func (sl *SortedList) Min() (int, bool) {
	if len(sl.items) == 0 {
		return 0, false
	}
	return sl.items[0], true
}

func (sl *SortedList) Max() (int, bool) {
	if len(sl.items) == 0 {
		return 0, false
	}
	return sl.items[len(sl.items)-1], true
}

func ejercicio3() {
	fmt.Println("=== Ejercicio 3: Property Testing Lista Ordenada ===")

	sl := NewSortedList()

	// Property test: insertar mantiene orden
	fmt.Println("Testing: Inserción mantiene orden")
	values := []int{5, 2, 8, 1, 9, 3}
	for _, v := range values {
		sl.Insert(v)
		slice := sl.ToSlice()
		if !sort.IntsAreSorted(slice) {
			fmt.Printf("❌ Lista no ordenada después de insertar %d: %v\n", v, slice)
			return
		}
	}
	fmt.Printf("✅ Lista mantiene orden: %v\n", sl.ToSlice())

	// Property test: Contains funciona correctamente
	fmt.Println("Testing: Contains funciona correctamente")
	for _, v := range values {
		if !sl.Contains(v) {
			fmt.Printf("❌ Contains(%d) debería ser true\n", v)
			return
		}
	}
	if sl.Contains(99) {
		fmt.Printf("❌ Contains(99) debería ser false\n")
		return
	}
	fmt.Printf("✅ Contains funciona correctamente\n")

	// Property test: Min y Max son correctos
	fmt.Println("Testing: Min y Max")
	min, minOk := sl.Min()
	max, maxOk := sl.Max()
	if !minOk || !maxOk {
		fmt.Printf("❌ Min o Max falló en lista no vacía\n")
		return
	}
	if min != 1 || max != 9 {
		fmt.Printf("❌ Min=%d (esperado 1), Max=%d (esperado 9)\n", min, max)
		return
	}
	fmt.Printf("✅ Min=%d, Max=%d correctos\n", min, max)

	// Property test: Remove funciona correctamente
	fmt.Println("Testing: Remove")
	originalSize := sl.Size()
	removed := sl.Remove(5)
	if !removed {
		fmt.Printf("❌ Remove(5) debería retornar true\n")
		return
	}
	if sl.Size() != originalSize-1 {
		fmt.Printf("❌ Tamaño incorrecto después de remove\n")
		return
	}
	if sl.Contains(5) {
		fmt.Printf("❌ Lista aún contiene 5 después de remove\n")
		return
	}
	fmt.Printf("✅ Remove funciona correctamente. Nueva lista: %v\n", sl.ToSlice())
}

// ================================================
// Ejercicio 4: Integration Testing - API Client
// ================================================

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (ac *APIClient) GetUser(userID int) (*User, error) {
	url := fmt.Sprintf("%s/users/%d", ac.baseURL, userID)
	resp, err := ac.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ac *APIClient) CreateUser(request CreateUserRequest) (*User, error) {
	url := fmt.Sprintf("%s/users", ac.baseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := ac.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ac *APIClient) UpdateUser(userID int, request CreateUserRequest) (*User, error) {
	url := fmt.Sprintf("%s/users/%d", ac.baseURL, userID)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("user not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ac *APIClient) DeleteUser(userID int) error {
	url := fmt.Sprintf("%s/users/%d", ac.baseURL, userID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("user not found")
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API error: %d", resp.StatusCode)
	}

	return nil
}

func (ac *APIClient) ListUsers() ([]User, error) {
	url := fmt.Sprintf("%s/users", ac.baseURL)
	resp, err := ac.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func ejercicio4() {
	fmt.Println("=== Ejercicio 4: Integration Testing API Client ===")

	// Para demostración, crear un cliente (en tests reales usaríamos httptest.Server)
	client := NewAPIClient("https://jsonplaceholder.typicode.com")

	// Demonstrar funcionalidad (estos podrían fallar sin conexión)
	fmt.Println("Cliente API creado para: https://jsonplaceholder.typicode.com")
	fmt.Println("En tests reales, usaríamos httptest.Server para simular la API")

	// Ejemplo de test structure (sin llamada real)
	fmt.Println("Estructura de test de integración:")
	fmt.Println("1. Crear httptest.Server con handlers simulados")
	fmt.Println("2. Crear APIClient apuntando al test server")
	fmt.Println("3. Hacer llamadas y verificar requests/responses")
	fmt.Println("4. Testear casos de error (404, 500, timeout)")
}

// ================================================
// Ejercicio 5: Benchmark Testing - Algoritmos de Búsqueda
// ================================================

func LinearSearch(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearch(slice []int, target int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func BinarySearchRecursive(slice []int, target int) int {
	return binarySearchHelper(slice, target, 0, len(slice)-1)
}

func binarySearchHelper(slice []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if slice[mid] == target {
		return mid
	} else if slice[mid] < target {
		return binarySearchHelper(slice, target, mid+1, right)
	} else {
		return binarySearchHelper(slice, target, left, mid-1)
	}
}

func BubbleSort(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)

	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

func QuickSort(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(slice []int, low, high int) {
	if low < high {
		pi := partition(slice, low, high)
		quickSortHelper(slice, low, pi-1)
		quickSortHelper(slice, pi+1, high)
	}
}

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

func MergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func ejercicio5() {
	fmt.Println("=== Ejercicio 5: Benchmark Testing Algoritmos ===")

	// Crear datos de prueba
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	sortedData := make([]int, len(data))
	copy(sortedData, data)
	sort.Ints(sortedData)

	target := sortedData[500] // Elemento en el medio

	// Comparar búsquedas
	fmt.Printf("Búsqueda en slice de %d elementos:\n", len(data))

	start := time.Now()
	linearResult := LinearSearch(sortedData, target)
	linearTime := time.Since(start)

	start = time.Now()
	binaryResult := BinarySearch(sortedData, target)
	binaryTime := time.Since(start)

	start = time.Now()
	binaryRecResult := BinarySearchRecursive(sortedData, target)
	binaryRecTime := time.Since(start)

	fmt.Printf("Linear Search: índice %d, tiempo %v\n", linearResult, linearTime)
	fmt.Printf("Binary Search: índice %d, tiempo %v\n", binaryResult, binaryTime)
	fmt.Printf("Binary Search Recursive: índice %d, tiempo %v\n", binaryRecResult, binaryRecTime)

	// Comparar algoritmos de ordenamiento
	unsortedData := make([]int, 100) // Datos más pequeños para bubble sort
	for i := range unsortedData {
		unsortedData[i] = rand.Intn(100)
	}

	fmt.Printf("\nOrdenamiento de slice de %d elementos:\n", len(unsortedData))

	start = time.Now()
	bubbleResult := BubbleSort(unsortedData)
	bubbleTime := time.Since(start)

	start = time.Now()
	quickResult := QuickSort(unsortedData)
	quickTime := time.Since(start)

	start = time.Now()
	mergeResult := MergeSort(unsortedData)
	mergeTime := time.Since(start)

	fmt.Printf("Bubble Sort: tiempo %v\n", bubbleTime)
	fmt.Printf("Quick Sort: tiempo %v\n", quickTime)
	fmt.Printf("Merge Sort: tiempo %v\n", mergeTime)

	// Verificar que están ordenados
	fmt.Printf("Bubble Sort correcto: %v\n", sort.IntsAreSorted(bubbleResult))
	fmt.Printf("Quick Sort correcto: %v\n", sort.IntsAreSorted(quickResult))
	fmt.Printf("Merge Sort correcto: %v\n", sort.IntsAreSorted(mergeResult))
}

// ===================================================
// Ejercicio 6: Test Suites - Sistema de Inventario
// ===================================================

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	products map[string]*Product
}

type Order struct {
	ID       string
	Products map[string]int // ProductID -> Quantity
	Total    float64
	Status   string
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]*Product),
	}
}

func (inv *Inventory) AddProduct(product Product) error {
	if product.ID == "" {
		return errors.New("product ID cannot be empty")
	}

	if product.Price < 0 {
		return errors.New("product price cannot be negative")
	}

	if product.Quantity < 0 {
		return errors.New("product quantity cannot be negative")
	}

	if _, exists := inv.products[product.ID]; exists {
		return errors.New("product already exists")
	}

	inv.products[product.ID] = &Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return nil
}

func (inv *Inventory) UpdateStock(productID string, quantity int) error {
	product, exists := inv.products[productID]
	if !exists {
		return errors.New("product not found")
	}

	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	product.Quantity = quantity
	return nil
}

func (inv *Inventory) GetProduct(productID string) (*Product, error) {
	product, exists := inv.products[productID]
	if !exists {
		return nil, errors.New("product not found")
	}

	// Retornar copia para evitar modificaciones externas
	return &Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (inv *Inventory) ProcessOrder(order Order) error {
	// Verificar que todos los productos existan y tengan stock suficiente
	for productID, quantity := range order.Products {
		product, exists := inv.products[productID]
		if !exists {
			return fmt.Errorf("product %s not found", productID)
		}

		if product.Quantity < quantity {
			return fmt.Errorf("insufficient stock for product %s: have %d, need %d",
				productID, product.Quantity, quantity)
		}
	}

	// Actualizar stock
	for productID, quantity := range order.Products {
		inv.products[productID].Quantity -= quantity
	}

	return nil
}

func (inv *Inventory) GetLowStockProducts(threshold int) []Product {
	var lowStockProducts []Product

	for _, product := range inv.products {
		if product.Quantity <= threshold {
			lowStockProducts = append(lowStockProducts, Product{
				ID:       product.ID,
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity,
			})
		}
	}

	return lowStockProducts
}

func ejercicio6() {
	fmt.Println("=== Ejercicio 6: Test Suites Sistema de Inventario ===")

	inv := NewInventory()

	// Agregar productos
	products := []Product{
		{ID: "P001", Name: "Laptop", Price: 999.99, Quantity: 10},
		{ID: "P002", Name: "Mouse", Price: 29.99, Quantity: 50},
		{ID: "P003", Name: "Keyboard", Price: 79.99, Quantity: 2},
	}

	for _, p := range products {
		if err := inv.AddProduct(p); err != nil {
			fmt.Printf("Error agregando producto %s: %v\n", p.ID, err)
			return
		}
	}
	fmt.Printf("✅ Agregados %d productos al inventario\n", len(products))

	// Procesar pedido
	order := Order{
		ID: "ORD001",
		Products: map[string]int{
			"P001": 2,
			"P002": 5,
		},
		Status: "pending",
	}

	if err := inv.ProcessOrder(order); err != nil {
		fmt.Printf("Error procesando pedido: %v\n", err)
		return
	}
	fmt.Printf("✅ Pedido %s procesado exitosamente\n", order.ID)

	// Verificar stock actualizado
	laptop, _ := inv.GetProduct("P001")
	mouse, _ := inv.GetProduct("P002")
	fmt.Printf("Stock actualizado - Laptop: %d, Mouse: %d\n", laptop.Quantity, mouse.Quantity)

	// Productos con stock bajo
	lowStock := inv.GetLowStockProducts(5)
	fmt.Printf("Productos con stock bajo (≤5): %d\n", len(lowStock))
	for _, p := range lowStock {
		fmt.Printf("  - %s: %d unidades\n", p.Name, p.Quantity)
	}
}

// =====================================================
// Ejercicio 7: Testify Framework - API de Autenticación
// =====================================================

type AuthService struct {
	users  map[string]string // username -> password
	tokens map[string]TokenData
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type TokenData struct {
	Username  string
	ExpiresAt time.Time
	IsValid   bool
}

func NewAuthService() *AuthService {
	return &AuthService{
		users:  make(map[string]string),
		tokens: make(map[string]TokenData),
	}
}

func (as *AuthService) RegisterUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password cannot be empty")
	}

	if _, exists := as.users[username]; exists {
		return errors.New("user already exists")
	}

	as.users[username] = password
	return nil
}

func (as *AuthService) Login(request LoginRequest) (*Token, error) {
	password, exists := as.users[request.Username]
	if !exists {
		return nil, errors.New("user not found")
	}

	if password != request.Password {
		return nil, errors.New("invalid password")
	}

	// Generar tokens
	accessToken, err := generateToken()
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken()
	if err != nil {
		return nil, err
	}

	// Almacenar token data
	as.tokens[accessToken] = TokenData{
		Username:  request.Username,
		ExpiresAt: time.Now().Add(1 * time.Hour),
		IsValid:   true,
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    3600, // 1 hora en segundos
	}, nil
}

func (as *AuthService) ValidateToken(token string) (bool, error) {
	tokenData, exists := as.tokens[token]
	if !exists {
		return false, nil
	}

	if !tokenData.IsValid {
		return false, nil
	}

	if time.Now().After(tokenData.ExpiresAt) {
		return false, nil
	}

	return true, nil
}

func (as *AuthService) RefreshToken(refreshToken string) (*Token, error) {
	// En una implementación real, validaríamos el refresh token
	// Por simplicidad, asumimos que el refresh token es válido

	newAccessToken, err := generateToken()
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := generateToken()
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    3600,
	}, nil
}

func (as *AuthService) Logout(token string) error {
	tokenData, exists := as.tokens[token]
	if !exists {
		return errors.New("invalid token")
	}

	tokenData.IsValid = false
	as.tokens[token] = tokenData
	return nil
}

func generateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ejercicio7() {
	fmt.Println("=== Ejercicio 7: Testify Framework Autenticación ===")

	authService := NewAuthService()

	// Registrar usuario
	err := authService.RegisterUser("testuser", "password123")
	if err != nil {
		fmt.Printf("Error registrando usuario: %v\n", err)
		return
	}
	fmt.Printf("✅ Usuario registrado exitosamente\n")

	// Login
	loginReq := LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	token, err := authService.Login(loginReq)
	if err != nil {
		fmt.Printf("Error en login: %v\n", err)
		return
	}
	fmt.Printf("✅ Login exitoso. Token: %s...\n", token.AccessToken[:16])

	// Validar token
	valid, err := authService.ValidateToken(token.AccessToken)
	if err != nil {
		fmt.Printf("Error validando token: %v\n", err)
		return
	}
	fmt.Printf("✅ Token válido: %v\n", valid)

	// Logout
	err = authService.Logout(token.AccessToken)
	if err != nil {
		fmt.Printf("Error en logout: %v\n", err)
		return
	}
	fmt.Printf("✅ Logout exitoso\n")

	// Validar token después del logout
	valid, err = authService.ValidateToken(token.AccessToken)
	if err != nil {
		fmt.Printf("Error validando token: %v\n", err)
		return
	}
	fmt.Printf("✅ Token válido después del logout: %v\n", valid)
}

// ============================================
// Main function para ejecutar ejercicios
// ============================================

func main() {
	fmt.Println("🧪 Soluciones de Testing Avanzado")
	fmt.Println("==================================")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	ejercicio1()
	fmt.Println()

	ejercicio2()
	fmt.Println()

	ejercicio3()
	fmt.Println()

	ejercicio4()
	fmt.Println()

	ejercicio5()
	fmt.Println()

	ejercicio6()
	fmt.Println()

	ejercicio7()
	fmt.Println()

	fmt.Println("🎉 Todas las soluciones completadas!")
	fmt.Println("💡 Para ejecutar tests:")
	fmt.Println("   go test -v")
	fmt.Println("   go test -bench=.")
	fmt.Println("   go test -cover")
}
