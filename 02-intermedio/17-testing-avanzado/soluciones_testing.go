// 🧪 Soluciones: Testing Avanzado
// Lección 17: TDD, Mocking y Property Testing

// +build ignore

package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
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
