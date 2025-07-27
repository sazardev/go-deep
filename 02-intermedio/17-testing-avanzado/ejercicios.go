// 🧪 Ejercicios: Testing Avanzado
// Lección 17: TDD, Mocking y Property Testing

package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"bytes"
	"math"
	"sort"
)

// ========================================
// Ejercicio 1: TDD - Calculadora Científica
// ========================================

// TODO: Usando TDD (Red-Green-Refactor), implementa una calculadora científica
// que soporte las siguientes operaciones:

type ScientificCalculator struct {
	history []string
	// TODO: Agregar campos necesarios (precisión, etc.)
}

// TODO: Implementar constructor
func NewScientificCalculator() *ScientificCalculator {
	return &ScientificCalculator{
		history: make([]string, 0),
	}
}

// TODO: Operaciones básicas (escribir tests primero)
func (sc *ScientificCalculator) Add(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	result := a + b
	sc.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Subtract(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	result := a - b
	sc.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Multiply(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	result := a * b
	sc.addToHistory(fmt.Sprintf("%.2f × %.2f = %.2f", a, b, result))
	return result
}

func (sc *ScientificCalculator) Divide(a, b float64) (float64, error) {
	// TODO: Implementar después de escribir el test
	// Manejar división por cero
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	result := a / b
	sc.addToHistory(fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, result))
	return result, nil
}

// TODO: Operaciones científicas (escribir tests primero)
func (sc *ScientificCalculator) Power(base, exponent float64) float64 {
	// TODO: Implementar después de escribir el test
	result := math.Pow(base, exponent)
	sc.addToHistory(fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
	return result
}

func (sc *ScientificCalculator) SquareRoot(x float64) (float64, error) {
	// TODO: Implementar después de escribir el test
	// Manejar números negativos
	if x < 0 {
		return 0, fmt.Errorf("raíz cuadrada de número negativo")
	}
	result := math.Sqrt(x)
	sc.addToHistory(fmt.Sprintf("√%.2f = %.2f", x, result))
	return result, nil
}

func (sc *ScientificCalculator) Sin(x float64) float64 {
	// TODO: Implementar después de escribir el test
	result := math.Sin(x)
	sc.addToHistory(fmt.Sprintf("sin(%.2f) = %.2f", x, result))
	return result
}

func (sc *ScientificCalculator) Cos(x float64) float64 {
	// TODO: Implementar después de escribir el test
	result := math.Cos(x)
	sc.addToHistory(fmt.Sprintf("cos(%.2f) = %.2f", x, result))
	return result
}

// TODO: Funcionalidades adicionales
func (sc *ScientificCalculator) GetHistory() []string {
	// TODO: Implementar historial de operaciones
	return sc.history
}

func (sc *ScientificCalculator) ClearHistory() {
	// TODO: Limpiar historial
	sc.history = make([]string, 0)
}

func (sc *ScientificCalculator) addToHistory(operation string) {
	sc.history = append(sc.history, operation)
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: TDD Calculadora Científica ===")
	// TODO: Crear tests en calculator_test.go siguiendo TDD
	// 1. Escribir test que falle (RED)
	// 2. Implementar código mínimo que pase (GREEN)
	// 3. Refactorizar manteniendo tests verdes (REFACTOR)
	
	calc := NewScientificCalculator()
	fmt.Printf("2 + 3 = %.2f\n", calc.Add(2, 3))
	fmt.Printf("10 - 4 = %.2f\n", calc.Subtract(10, 4))
	
	result, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}
	
	fmt.Printf("Historial: %v\n", calc.GetHistory())
}

// ==========================================
// Ejercicio 2: Mocking - Sistema de Notificaciones
// ==========================================

// TODO: Interfaces para servicios externos
type EmailSender interface {
	SendEmail(to, subject, body string) error
}

type SMSSender interface {
	SendSMS(to, message string) error
}

type PushNotificationSender interface {
	SendPushNotification(userID, title, body string) error
}

// TODO: Implementar servicio de notificaciones que use las interfaces
type NotificationService struct {
	emailSender EmailSender
	smsSender   SMSSender
	pushSender  PushNotificationSender
}

// TODO: Constructor que reciba las dependencias
func NewNotificationService(emailSender EmailSender, smsSender SMSSender, pushSender PushNotificationSender) *NotificationService {
	return &NotificationService{
		emailSender: emailSender,
		smsSender:   smsSender,
		pushSender:  pushSender,
	}
}

// TODO: Métodos para enviar diferentes tipos de notificaciones
func (ns *NotificationService) SendWelcomeNotification(userID string, email string, phone string) error {
	// TODO: Enviar email, SMS y push notification
	// Manejar errores de cada servicio
	
	if err := ns.emailSender.SendEmail(email, "Bienvenido", "Gracias por registrarte"); err != nil {
		return fmt.Errorf("error enviando email: %w", err)
	}
	
	if err := ns.smsSender.SendSMS(phone, "Bienvenido a nuestra plataforma"); err != nil {
		return fmt.Errorf("error enviando SMS: %w", err)
	}
	
	if err := ns.pushSender.SendPushNotification(userID, "Bienvenido", "¡Gracias por unirte!"); err != nil {
		return fmt.Errorf("error enviando push: %w", err)
	}
	
	return nil
}

func (ns *NotificationService) SendPasswordResetNotification(userID string, email string) error {
	// TODO: Enviar solo email
	return ns.emailSender.SendEmail(email, "Resetear Contraseña", "Haz clic para resetear tu contraseña")
}

func (ns *NotificationService) SendOrderConfirmation(userID string, email string, phone string, orderID string) error {
	// TODO: Enviar email y SMS con detalles del pedido
	emailBody := fmt.Sprintf("Tu pedido %s ha sido confirmado", orderID)
	smsBody := fmt.Sprintf("Pedido %s confirmado", orderID)
	
	if err := ns.emailSender.SendEmail(email, "Pedido Confirmado", emailBody); err != nil {
		return fmt.Errorf("error enviando email: %w", err)
	}
	
	if err := ns.smsSender.SendSMS(phone, smsBody); err != nil {
		return fmt.Errorf("error enviando SMS: %w", err)
	}
	
	return nil
}

// TODO: Crear mocks para testing
type EmailSenderMock struct {
	calls []EmailCall
	shouldFail bool
}

type EmailCall struct {
	To      string
	Subject string
	Body    string
}

func (m *EmailSenderMock) SendEmail(to, subject, body string) error {
	m.calls = append(m.calls, EmailCall{To: to, Subject: subject, Body: body})
	if m.shouldFail {
		return fmt.Errorf("mock error")
	}
	return nil
}

func (m *EmailSenderMock) GetCalls() []EmailCall {
	return m.calls
}

func (m *EmailSenderMock) SetShouldFail(fail bool) {
	m.shouldFail = fail
}

func ejercicio2() {
	fmt.Println("=== Ejercicio 2: Mocking Sistema de Notificaciones ===")
	// TODO: Crear mocks que permitan:
	// 1. Verificar que se llaman los métodos correctos
	// 2. Verificar los parámetros pasados
	// 3. Simular diferentes tipos de errores
	// 4. Contar número de llamadas
	
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushSenderMock{}
	
	service := NewNotificationService(emailMock, smsMock, pushMock)
	
	err := service.SendWelcomeNotification("user123", "test@email.com", "+1234567890")
	fmt.Printf("Error enviando notificación: %v\n", err)
	fmt.Printf("Llamadas a email: %d\n", len(emailMock.GetCalls()))
}

// Mocks adicionales para completar el ejemplo
type SMSSenderMock struct {
	calls []SMSCall
	shouldFail bool
}

type SMSCall struct {
	To      string
	Message string
}

func (m *SMSSenderMock) SendSMS(to, message string) error {
	m.calls = append(m.calls, SMSCall{To: to, Message: message})
	if m.shouldFail {
		return fmt.Errorf("mock SMS error")
	}
	return nil
}

type PushSenderMock struct {
	calls []PushCall
	shouldFail bool
}

type PushCall struct {
	UserID string
	Title  string
	Body   string
}

func (m *PushSenderMock) SendPushNotification(userID, title, body string) error {
	m.calls = append(m.calls, PushCall{UserID: userID, Title: title, Body: body})
	if m.shouldFail {
		return fmt.Errorf("mock push error")
	}
	return nil
}

// ===============================================
// Ejercicio 3: Property Testing - Lista Ordenada
// ===============================================

// TODO: Implementar una lista ordenada que mantenga los elementos ordenados
type SortedList struct {
	elements []int
}

// TODO: Constructor
func NewSortedList() *SortedList {
	return &SortedList{
		elements: make([]int, 0),
	}
}

// TODO: Métodos de la lista ordenada
func (sl *SortedList) Insert(value int) {
	// TODO: Insertar manteniendo orden
	pos := sort.SearchInts(sl.elements, value)
	sl.elements = append(sl.elements, 0)
	copy(sl.elements[pos+1:], sl.elements[pos:])
	sl.elements[pos] = value
}

func (sl *SortedList) Remove(value int) bool {
	// TODO: Remover elemento si existe
	pos := sort.SearchInts(sl.elements, value)
	if pos < len(sl.elements) && sl.elements[pos] == value {
		sl.elements = append(sl.elements[:pos], sl.elements[pos+1:]...)
		return true
	}
	return false
}

func (sl *SortedList) Contains(value int) bool {
	// TODO: Verificar si contiene el elemento
	pos := sort.SearchInts(sl.elements, value)
	return pos < len(sl.elements) && sl.elements[pos] == value
}

func (sl *SortedList) Size() int {
	// TODO: Retornar tamaño
	return len(sl.elements)
}

func (sl *SortedList) ToSlice() []int {
	// TODO: Convertir a slice ordenado
	result := make([]int, len(sl.elements))
	copy(result, sl.elements)
	return result
}

func (sl *SortedList) Min() (int, bool) {
	// TODO: Retornar elemento mínimo
	if len(sl.elements) == 0 {
		return 0, false
	}
	return sl.elements[0], true
}

func (sl *SortedList) Max() (int, bool) {
	// TODO: Retornar elemento máximo
	if len(sl.elements) == 0 {
		return 0, false
	}
	return sl.elements[len(sl.elements)-1], true
}

func ejercicio3() {
	fmt.Println("=== Ejercicio 3: Property Testing Lista Ordenada ===")
	// TODO: Crear property tests que verifiquen:
	// 1. La lista siempre está ordenada después de insertar
	// 2. El tamaño aumenta correctamente al insertar
	// 3. El tamaño disminuye correctamente al remover
	// 4. Contains funciona correctamente
	// 5. Min y Max son correctos
	// 6. ToSlice retorna slice ordenado
	
	list := NewSortedList()
	list.Insert(5)
	list.Insert(2)
	list.Insert(8)
	list.Insert(1)
	
	fmt.Printf("Lista: %v\n", list.ToSlice())
	fmt.Printf("Tamaño: %d\n", list.Size())
	
	min, hasMin := list.Min()
	max, hasMax := list.Max()
	fmt.Printf("Min: %d (existe: %t), Max: %d (existe: %t)\n", min, hasMin, max, hasMax)
	
	fmt.Printf("Contiene 5: %t\n", list.Contains(5))
	fmt.Printf("Contiene 10: %t\n", list.Contains(10))
	
	removed := list.Remove(5)
	fmt.Printf("Removido 5: %t\n", removed)
	fmt.Printf("Lista después: %v\n", list.ToSlice())
}

// =============================================
// Ejercicio 4: Integration Testing - API Client
// =============================================

// TODO: Cliente para una API REST
type APIClient struct {
	baseURL string
	timeout time.Duration
	client  *http.Client
}

// TODO: Estructuras para respuestas de la API
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TODO: Constructor del cliente
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		timeout: 30 * time.Second,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// TODO: Métodos del cliente
func (ac *APIClient) GetUser(userID int) (*User, error) {
	// TODO: GET /users/{id}
	url := fmt.Sprintf("%s/users/%d", ac.baseURL, userID)
	resp, err := ac.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
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
	// TODO: POST /users
	url := fmt.Sprintf("%s/users", ac.baseURL)
	
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	resp, err := ac.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
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
	// TODO: PUT /users/{id}
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
	
	resp, err := ac.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
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
	// TODO: DELETE /users/{id}
	url := fmt.Sprintf("%s/users/%d", ac.baseURL, userID)
	
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	
	resp, err := ac.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API error: %d", resp.StatusCode)
	}
	
	return nil
}

func (ac *APIClient) ListUsers() ([]User, error) {
	// TODO: GET /users
	url := fmt.Sprintf("%s/users", ac.baseURL)
	resp, err := ac.client.Get(url)
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
	// TODO: Crear integration tests que:
	// 1. Usen un servidor HTTP de test
	// 2. Simulen diferentes respuestas de la API
	// 3. Verifiquen timeout y retry logic
	// 4. Testen casos de error (404, 500, etc.)
	// 5. Verifiquen serialización/deserialización JSON
	
	// Ejemplo básico (en un test real usarías httptest.Server)
	client := NewAPIClient("https://jsonplaceholder.typicode.com")
	
	user, err := client.GetUser(1)
	if err != nil {
		fmt.Printf("Error obteniendo usuario: %v\n", err)
	} else {
		fmt.Printf("Usuario obtenido: %+v\n", user)
	}
}

// ============================================
// Main function para ejecutar ejercicios
// ============================================

func main() {
	fmt.Println("🧪 Ejercicios de Testing Avanzado")
	fmt.Println("==================================")
	fmt.Println()

	ejercicio1()
	fmt.Println()

	ejercicio2()
	fmt.Println()

	ejercicio3()
	fmt.Println()

	ejercicio4()
	fmt.Println()

	fmt.Println("💡 Para completar los ejercicios:")
	fmt.Println("   1. Implementa los TODOs siguiendo TDD")
	fmt.Println("   2. Crea archivos *_test.go para cada ejercicio")
	fmt.Println("   3. Escribe property tests y benchmarks")
	fmt.Println("   4. Usa mocking para aislar dependencias")
	fmt.Println("   5. Practica integration testing con servidores de test")
}
