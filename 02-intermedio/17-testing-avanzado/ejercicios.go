// 🧪 Ejercicios: Testing Avanzado
// Lección 17: TDD, Mocking y Property Testing

package main

import (
	"fmt"
)

// ========================================
// Ejercicio 1: TDD - Calculadora Científica
// ========================================

// TODO: Usando TDD (Red-Green-Refactor), implementa una calculadora científica
// que soporte las siguientes operaciones:

type ScientificCalculator struct {
	// TODO: Agregar campos necesarios (historial, precisión, etc.)
}

// TODO: Implementar constructor
func NewScientificCalculator() *ScientificCalculator {
	return nil
}

// TODO: Operaciones básicas (escribir tests primero)
func (sc *ScientificCalculator) Add(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

func (sc *ScientificCalculator) Subtract(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

func (sc *ScientificCalculator) Multiply(a, b float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

func (sc *ScientificCalculator) Divide(a, b float64) (float64, error) {
	// TODO: Implementar después de escribir el test
	// Manejar división por cero
	return 0, nil
}

// TODO: Operaciones científicas (escribir tests primero)
func (sc *ScientificCalculator) Power(base, exponent float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

func (sc *ScientificCalculator) SquareRoot(x float64) (float64, error) {
	// TODO: Implementar después de escribir el test
	// Manejar números negativos
	return 0, nil
}

func (sc *ScientificCalculator) Sin(x float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

func (sc *ScientificCalculator) Cos(x float64) float64 {
	// TODO: Implementar después de escribir el test
	return 0
}

// TODO: Funcionalidades adicionales
func (sc *ScientificCalculator) GetHistory() []string {
	// TODO: Implementar historial de operaciones
	return nil
}

func (sc *ScientificCalculator) ClearHistory() {
	// TODO: Limpiar historial
}

func ejercicio1() {
	fmt.Println("=== Ejercicio 1: TDD Calculadora Científica ===")
	// TODO: Crear tests en calculator_test.go siguiendo TDD
	// 1. Escribir test que falle (RED)
	// 2. Implementar código mínimo que pase (GREEN)
	// 3. Refactorizar manteniendo tests verdes (REFACTOR)
}

// ==========================================
// Ejercicio 2: Mocking - Sistema de Notificaciones
// ==========================================

// TODO: Interfaces para servicios externos
type EmailSender interface {
	// TODO: Definir métodos para envío de email
}

type SMSSender interface {
	// TODO: Definir métodos para envío de SMS
}

type PushNotificationSender interface {
	// TODO: Definir métodos para notificaciones push
}

// TODO: Implementar servicio de notificaciones que use las interfaces
type NotificationService struct {
	// TODO: Agregar campos para las dependencias
}

// TODO: Constructor que reciba las dependencias
func NewNotificationService(emailSender EmailSender, smsSender SMSSender, pushSender PushNotificationSender) *NotificationService {
	return nil
}

// TODO: Métodos para enviar diferentes tipos de notificaciones
func (ns *NotificationService) SendWelcomeNotification(userID string, email string, phone string) error {
	// TODO: Enviar email, SMS y push notification
	// Manejar errores de cada servicio
	return nil
}

func (ns *NotificationService) SendPasswordResetNotification(userID string, email string) error {
	// TODO: Enviar solo email
	return nil
}

func (ns *NotificationService) SendOrderConfirmation(userID string, email string, phone string, orderID string) error {
	// TODO: Enviar email y SMS con detalles del pedido
	return nil
}

// TODO: Crear mocks para testing
// Ejemplo de estructura para el mock:
type EmailSenderMock struct {
	// TODO: Campos para capturar llamadas y configurar respuestas
}

// TODO: Implementar métodos del mock con verificación

func ejercicio2() {
	fmt.Println("=== Ejercicio 2: Mocking Sistema de Notificaciones ===")
	// TODO: Crear mocks que permitan:
	// 1. Verificar que se llaman los métodos correctos
	// 2. Verificar los parámetros pasados
	// 3. Simular diferentes tipos de errores
	// 4. Contar número de llamadas
}

// ===============================================
// Ejercicio 3: Property Testing - Lista Ordenada
// ===============================================

// TODO: Implementar una lista ordenada que mantenga los elementos ordenados
type SortedList struct {
	// TODO: Campos necesarios
}

// TODO: Constructor
func NewSortedList() *SortedList {
	return nil
}

// TODO: Métodos de la lista ordenada
func (sl *SortedList) Insert(value int) {
	// TODO: Insertar manteniendo orden
}

func (sl *SortedList) Remove(value int) bool {
	// TODO: Remover elemento si existe
	return false
}

func (sl *SortedList) Contains(value int) bool {
	// TODO: Verificar si contiene el elemento
	return false
}

func (sl *SortedList) Size() int {
	// TODO: Retornar tamaño
	return 0
}

func (sl *SortedList) ToSlice() []int {
	// TODO: Convertir a slice ordenado
	return nil
}

func (sl *SortedList) Min() (int, bool) {
	// TODO: Retornar elemento mínimo
	return 0, false
}

func (sl *SortedList) Max() (int, bool) {
	// TODO: Retornar elemento máximo
	return 0, false
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
}

// =============================================
// Ejercicio 4: Integration Testing - API Client
// =============================================

// TODO: Cliente para una API REST
type APIClient struct {
	// TODO: Campos para configuración (base URL, timeout, etc.)
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
	return nil
}

// TODO: Métodos del cliente
func (ac *APIClient) GetUser(userID int) (*User, error) {
	// TODO: GET /users/{id}
	return nil, nil
}

func (ac *APIClient) CreateUser(request CreateUserRequest) (*User, error) {
	// TODO: POST /users
	return nil, nil
}

func (ac *APIClient) UpdateUser(userID int, request CreateUserRequest) (*User, error) {
	// TODO: PUT /users/{id}
	return nil, nil
}

func (ac *APIClient) DeleteUser(userID int) error {
	// TODO: DELETE /users/{id}
	return nil
}

func (ac *APIClient) ListUsers() ([]User, error) {
	// TODO: GET /users
	return nil, nil
}

func ejercicio4() {
	fmt.Println("=== Ejercicio 4: Integration Testing API Client ===")
	// TODO: Crear integration tests que:
	// 1. Usen un servidor HTTP de test
	// 2. Simulen diferentes respuestas de la API
	// 3. Verifiquen timeout y retry logic
	// 4. Testen casos de error (404, 500, etc.)
	// 5. Verifiquen serialización/deserialización JSON
}

// ================================================
// Ejercicio 5: Benchmark Testing - Algoritmos de Búsqueda
// ================================================

// TODO: Implementar diferentes algoritmos de búsqueda
func LinearSearch(slice []int, target int) int {
	// TODO: Búsqueda lineal
	return -1
}

func BinarySearch(slice []int, target int) int {
	// TODO: Búsqueda binaria (slice debe estar ordenado)
	return -1
}

func BinarySearchRecursive(slice []int, target int) int {
	// TODO: Búsqueda binaria recursiva
	return -1
}

// TODO: Implementar algoritmos de ordenamiento para benchmarks
func BubbleSort(slice []int) []int {
	// TODO: Bubble sort
	return nil
}

func QuickSort(slice []int) []int {
	// TODO: Quick sort
	return nil
}

func MergeSort(slice []int) []int {
	// TODO: Merge sort
	return nil
}

func ejercicio5() {
	fmt.Println("=== Ejercicio 5: Benchmark Testing Algoritmos ===")
	// TODO: Crear benchmarks que:
	// 1. Comparen performance de búsqueda lineal vs binaria
	// 2. Comparen diferentes algoritmos de ordenamiento
	// 3. Midan uso de memoria
	// 4. Testen con diferentes tamaños de datos
	// 5. Incluyan benchmarks de CPU y memoria
}

// ===================================================
// Ejercicio 6: Test Suites - Sistema de Inventario
// ===================================================

// TODO: Sistema de inventario con múltiples componentes
type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	// TODO: Campos necesarios (productos, índices, etc.)
}

type Order struct {
	ID       string
	Products map[string]int // ProductID -> Quantity
	Total    float64
	Status   string
}

// TODO: Implementar sistema de inventario
func NewInventory() *Inventory {
	return nil
}

func (inv *Inventory) AddProduct(product Product) error {
	// TODO: Agregar producto al inventario
	return nil
}

func (inv *Inventory) UpdateStock(productID string, quantity int) error {
	// TODO: Actualizar stock de producto
	return nil
}

func (inv *Inventory) GetProduct(productID string) (*Product, error) {
	// TODO: Obtener producto por ID
	return nil, nil
}

func (inv *Inventory) ProcessOrder(order Order) error {
	// TODO: Procesar pedido (verificar stock, actualizar inventario)
	return nil
}

func (inv *Inventory) GetLowStockProducts(threshold int) []Product {
	// TODO: Obtener productos con stock bajo
	return nil
}

func ejercicio6() {
	fmt.Println("=== Ejercicio 6: Test Suites Sistema de Inventario ===")
	// TODO: Crear suite de tests que incluya:
	// 1. Setup/Teardown para cada test
	// 2. Tests unitarios para cada método
	// 3. Tests de integración para flujos completos
	// 4. Tests de edge cases (stock negativo, productos duplicados)
	// 5. Tests de concurrencia (múltiples operaciones simultáneas)
}

// =====================================================
// Ejercicio 7: Testify Framework - API de Autenticación
// =====================================================

// TODO: Sistema de autenticación usando testify
type AuthService struct {
	// TODO: Campos necesarios
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

// TODO: Implementar servicio de autenticación
func NewAuthService() *AuthService {
	return nil
}

func (as *AuthService) Login(request LoginRequest) (*Token, error) {
	// TODO: Validar credenciales y generar token
	return nil, nil
}

func (as *AuthService) ValidateToken(token string) (bool, error) {
	// TODO: Validar token JWT
	return false, nil
}

func (as *AuthService) RefreshToken(refreshToken string) (*Token, error) {
	// TODO: Renovar token usando refresh token
	return nil, nil
}

func (as *AuthService) Logout(token string) error {
	// TODO: Invalidar token
	return nil
}

func ejercicio7() {
	fmt.Println("=== Ejercicio 7: Testify Framework Autenticación ===")
	// TODO: Usar testify para:
	// 1. assert.* para verificaciones simples
	// 2. require.* para verificaciones que deben parar el test
	// 3. mock.* para crear mocks avanzados
	// 4. suite.* para organizar tests en suites
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

	ejercicio5()
	fmt.Println()

	ejercicio6()
	fmt.Println()

	ejercicio7()
	fmt.Println()

	fmt.Println("💡 Para completar los ejercicios:")
	fmt.Println("   1. Implementa los TODOs siguiendo TDD")
	fmt.Println("   2. Crea archivos *_test.go para cada ejercicio")
	fmt.Println("   3. Ejecuta: go test -v")
	fmt.Println("   4. Para benchmarks: go test -bench=.")
}
