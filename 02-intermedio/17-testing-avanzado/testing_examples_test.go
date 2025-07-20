// 🧪 Ejemplos de Testing Avanzado
// Archivo: testing_examples_test.go
// Ejecutar con: go test -v

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

// ==========================================
// 🧪 UNIT TESTS CON TDD
// ==========================================

// Calculator es un ejemplo simple para TDD
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// ==========================================
// 🧪 TESTS TDD PARA CALCULATOR
// ==========================================

func TestCalculator_Add(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero", 0, 5, 5},
	}

	calc := &Calculator{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := &Calculator{}

	t.Run("valid division", func(t *testing.T) {
		result, err := calc.Divide(10, 2)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if result != 5 {
			t.Errorf("Divide(10, 2) = %d; expected 5", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Fatal("Expected error for division by zero")
		}
		if !strings.Contains(err.Error(), "division by zero") {
			t.Errorf("Expected 'division by zero' error, got: %v", err)
		}
	})
}

// ==========================================
// 🎯 PROPERTY-BASED TESTS
// ==========================================

func TestCalculator_AddCommutativeProperty(t *testing.T) {
	calc := &Calculator{}
	
	// Property: a + b = b + a (commutative)
	for i := 0; i < 100; i++ {
		a := rand.Intn(1000) - 500 // -500 to 500
		b := rand.Intn(1000) - 500
		
		result1 := calc.Add(a, b)
		result2 := calc.Add(b, a)
		
		if result1 != result2 {
			t.Errorf("Commutative property failed: %d + %d = %d, but %d + %d = %d", 
				a, b, result1, b, a, result2)
		}
	}
}

func TestCalculator_AddAssociativeProperty(t *testing.T) {
	calc := &Calculator{}
	
	// Property: (a + b) + c = a + (b + c) (associative)
	for i := 0; i < 100; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)
		
		result1 := calc.Add(calc.Add(a, b), c)
		result2 := calc.Add(a, calc.Add(b, c))
		
		if result1 != result2 {
			t.Errorf("Associative property failed: (%d + %d) + %d = %d, but %d + (%d + %d) = %d", 
				a, b, c, result1, a, b, c, result2)
		}
	}
}

func TestCalculator_MultiplyByZeroProperty(t *testing.T) {
	calc := &Calculator{}
	
	// Property: any number multiplied by zero equals zero
	for i := 0; i < 50; i++ {
		n := rand.Intn(1000) - 500
		result := calc.Multiply(n, 0)
		
		if result != 0 {
			t.Errorf("Zero property failed: %d * 0 = %d, expected 0", n, result)
		}
	}
}

// ==========================================
// 🔄 MOCK INTERFACES Y TESTS
// ==========================================

// EmailService interface para mocking
type EmailService interface {
	SendEmail(to, subject, body string) error
}

// UserNotifier utiliza EmailService
type UserNotifier struct {
	emailService EmailService
}

func NewUserNotifier(emailService EmailService) *UserNotifier {
	return &UserNotifier{emailService: emailService}
}

func (n *UserNotifier) NotifyUser(userEmail, message string) error {
	return n.emailService.SendEmail(userEmail, "Notification", message)
}

// MockEmailService para testing
type MockEmailService struct {
	sentEmails []EmailCall
	shouldFail bool
	failError  error
}

type EmailCall struct {
	To      string
	Subject string
	Body    string
}

func NewMockEmailService() *MockEmailService {
	return &MockEmailService{
		sentEmails: make([]EmailCall, 0),
	}
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
	if m.shouldFail {
		return m.failError
	}
	
	m.sentEmails = append(m.sentEmails, EmailCall{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	return nil
}

func (m *MockEmailService) SetShouldFail(fail bool, err error) {
	m.shouldFail = fail
	m.failError = err
}

func (m *MockEmailService) GetSentEmails() []EmailCall {
	return m.sentEmails
}

func (m *MockEmailService) Reset() {
	m.sentEmails = make([]EmailCall, 0)
	m.shouldFail = false
	m.failError = nil
}

// ==========================================
// 🧪 TESTS CON MOCKS
// ==========================================

func TestUserNotifier_NotifyUser_Success(t *testing.T) {
	// Arrange
	mockEmail := NewMockEmailService()
	notifier := NewUserNotifier(mockEmail)
	
	userEmail := "test@example.com"
	message := "Hello, World!"
	
	// Act
	err := notifier.NotifyUser(userEmail, message)
	
	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	
	sentEmails := mockEmail.GetSentEmails()
	if len(sentEmails) != 1 {
		t.Fatalf("Expected 1 email sent, got %d", len(sentEmails))
	}
	
	email := sentEmails[0]
	if email.To != userEmail {
		t.Errorf("Expected email to %s, got %s", userEmail, email.To)
	}
	if email.Subject != "Notification" {
		t.Errorf("Expected subject 'Notification', got %s", email.Subject)
	}
	if email.Body != message {
		t.Errorf("Expected body '%s', got %s", message, email.Body)
	}
}

func TestUserNotifier_NotifyUser_EmailServiceFailure(t *testing.T) {
	// Arrange
	mockEmail := NewMockEmailService()
	expectedError := errors.New("email service unavailable")
	mockEmail.SetShouldFail(true, expectedError)
	
	notifier := NewUserNotifier(mockEmail)
	
	// Act
	err := notifier.NotifyUser("test@example.com", "message")
	
	// Assert
	if err == nil {
		t.Fatal("Expected error when email service fails")
	}
	if err != expectedError {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
	
	sentEmails := mockEmail.GetSentEmails()
	if len(sentEmails) != 0 {
		t.Errorf("Expected no emails sent on failure, got %d", len(sentEmails))
	}
}

// ==========================================
// 🏭 TEST FACTORY PATTERN
// ==========================================

type TestUser struct {
	ID    string
	Email string
	Name  string
	Age   int
}

type UserBuilder struct {
	user *TestUser
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &TestUser{
			ID:    "default-id",
			Email: "default@example.com",
			Name:  "Default User",
			Age:   25,
		},
	}
}

func (b *UserBuilder) WithID(id string) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Build() *TestUser {
	return &TestUser{
		ID:    b.user.ID,
		Email: b.user.Email,
		Name:  b.user.Name,
		Age:   b.user.Age,
	}
}

// UserValidator para testing del builder
type UserValidator struct{}

func (v *UserValidator) ValidateUser(user *TestUser) []string {
	var errors []string
	
	if user.ID == "" {
		errors = append(errors, "ID is required")
	}
	if user.Email == "" {
		errors = append(errors, "Email is required")
	}
	if !strings.Contains(user.Email, "@") {
		errors = append(errors, "Email must contain @")
	}
	if user.Name == "" {
		errors = append(errors, "Name is required")
	}
	if user.Age < 0 {
		errors = append(errors, "Age must be non-negative")
	}
	if user.Age > 150 {
		errors = append(errors, "Age must be reasonable")
	}
	
	return errors
}

// ==========================================
// 🧪 TESTS CON BUILDER PATTERN
// ==========================================

func TestUserBuilder_DefaultValues(t *testing.T) {
	user := NewUserBuilder().Build()
	
	if user.ID != "default-id" {
		t.Errorf("Expected default ID 'default-id', got %s", user.ID)
	}
	if user.Email != "default@example.com" {
		t.Errorf("Expected default email 'default@example.com', got %s", user.Email)
	}
	if user.Name != "Default User" {
		t.Errorf("Expected default name 'Default User', got %s", user.Name)
	}
	if user.Age != 25 {
		t.Errorf("Expected default age 25, got %d", user.Age)
	}
}

func TestUserBuilder_CustomValues(t *testing.T) {
	user := NewUserBuilder().
		WithID("custom-123").
		WithEmail("custom@test.com").
		WithName("Custom User").
		WithAge(30).
		Build()
	
	if user.ID != "custom-123" {
		t.Errorf("Expected ID 'custom-123', got %s", user.ID)
	}
	if user.Email != "custom@test.com" {
		t.Errorf("Expected email 'custom@test.com', got %s", user.Email)
	}
	if user.Name != "Custom User" {
		t.Errorf("Expected name 'Custom User', got %s", user.Name)
	}
	if user.Age != 30 {
		t.Errorf("Expected age 30, got %d", user.Age)
	}
}

func TestUserValidator_ValidUser(t *testing.T) {
	validator := &UserValidator{}
	user := NewUserBuilder().
		WithID("valid-123").
		WithEmail("valid@example.com").
		WithName("Valid User").
		WithAge(25).
		Build()
	
	errors := validator.ValidateUser(user)
	
	if len(errors) != 0 {
		t.Errorf("Expected no validation errors, got: %v", errors)
	}
}

func TestUserValidator_InvalidUser(t *testing.T) {
	validator := &UserValidator{}
	
	tests := []struct {
		name          string
		user          *TestUser
		expectedError string
	}{
		{
			name: "empty ID",
			user: NewUserBuilder().WithID("").Build(),
			expectedError: "ID is required",
		},
		{
			name: "empty email",
			user: NewUserBuilder().WithEmail("").Build(),
			expectedError: "Email is required",
		},
		{
			name: "invalid email",
			user: NewUserBuilder().WithEmail("invalid-email").Build(),
			expectedError: "Email must contain @",
		},
		{
			name: "negative age",
			user: NewUserBuilder().WithAge(-1).Build(),
			expectedError: "Age must be non-negative",
		},
		{
			name: "unreasonable age",
			user: NewUserBuilder().WithAge(200).Build(),
			expectedError: "Age must be reasonable",
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errors := validator.ValidateUser(test.user)
			
			if len(errors) == 0 {
				t.Fatal("Expected validation errors, got none")
			}
			
			found := false
			for _, err := range errors {
				if strings.Contains(err, test.expectedError) {
					found = true
					break
				}
			}
			
			if !found {
				t.Errorf("Expected error containing '%s', got: %v", test.expectedError, errors)
			}
		})
	}
}

// ==========================================
// ⚡ BENCHMARK TESTS
// ==========================================

func BenchmarkCalculator_Add(b *testing.B) {
	calc := &Calculator{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Add(i, i+1)
	}
}

func BenchmarkCalculator_Multiply(b *testing.B) {
	calc := &Calculator{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Multiply(i, 2)
	}
}

func BenchmarkUserBuilder_Build(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUserBuilder().
			WithID(fmt.Sprintf("user-%d", i)).
			WithEmail(fmt.Sprintf("user%d@example.com", i)).
			WithName(fmt.Sprintf("User %d", i)).
			WithAge(20 + i%50).
			Build()
	}
}

func BenchmarkUserValidator_ValidateUser(b *testing.B) {
	validator := &UserValidator{}
	user := NewUserBuilder().
		WithID("benchmark-user").
		WithEmail("benchmark@example.com").
		WithName("Benchmark User").
		WithAge(30).
		Build()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.ValidateUser(user)
	}
}

// ==========================================
// 🧪 INTEGRATION TESTS
// ==========================================

func TestUserWorkflow_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	
	// Setup
	mockEmail := NewMockEmailService()
	notifier := NewUserNotifier(mockEmail)
	validator := &UserValidator{}
	
	// Test workflow: Create -> Validate -> Notify
	t.Run("Complete user workflow", func(t *testing.T) {
		// 1. Create user
		user := NewUserBuilder().
			WithID("integration-user").
			WithEmail("integration@example.com").
			WithName("Integration User").
			WithAge(28).
			Build()
		
		// 2. Validate user
		errors := validator.ValidateUser(user)
		if len(errors) != 0 {
			t.Fatalf("User validation failed: %v", errors)
		}
		
		// 3. Notify user
		welcomeMessage := fmt.Sprintf("Welcome %s! Your account has been created.", user.Name)
		err := notifier.NotifyUser(user.Email, welcomeMessage)
		if err != nil {
			t.Fatalf("Failed to notify user: %v", err)
		}
		
		// 4. Verify notification was sent
		sentEmails := mockEmail.GetSentEmails()
		if len(sentEmails) != 1 {
			t.Fatalf("Expected 1 email, got %d", len(sentEmails))
		}
		
		email := sentEmails[0]
		if email.To != user.Email {
			t.Errorf("Email sent to wrong address: expected %s, got %s", user.Email, email.To)
		}
		
		if !strings.Contains(email.Body, user.Name) {
			t.Errorf("Email body should contain user name '%s', got: %s", user.Name, email.Body)
		}
	})
}

// ==========================================
// 🎯 TABLE-DRIVEN TESTS
// ==========================================

func TestCalculator_AllOperations(t *testing.T) {
	calc := &Calculator{}
	
	tests := []struct {
		name        string
		operation   string
		a, b        int
		expected    int
		expectError bool
	}{
		{"add positive", "add", 5, 3, 8, false},
		{"add negative", "add", -5, -3, -8, false},
		{"subtract positive", "subtract", 10, 3, 7, false},
		{"subtract negative", "subtract", 5, 10, -5, false},
		{"multiply positive", "multiply", 4, 3, 12, false},
		{"multiply by zero", "multiply", 5, 0, 0, false},
		{"divide valid", "divide", 10, 2, 5, false},
		{"divide by zero", "divide", 10, 0, 0, true},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result int
			var err error
			
			switch test.operation {
			case "add":
				result = calc.Add(test.a, test.b)
			case "subtract":
				result = calc.Subtract(test.a, test.b)
			case "multiply":
				result = calc.Multiply(test.a, test.b)
			case "divide":
				result, err = calc.Divide(test.a, test.b)
			default:
				t.Fatalf("Unknown operation: %s", test.operation)
			}
			
			if test.expectError {
				if err == nil {
					t.Errorf("Expected error for %s(%d, %d), got none", test.operation, test.a, test.b)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s(%d, %d): %v", test.operation, test.a, test.b, err)
				}
				if result != test.expected {
					t.Errorf("%s(%d, %d) = %d; expected %d", test.operation, test.a, test.b, result, test.expected)
				}
			}
		})
	}
}

// ==========================================
// 🔍 REFLECTION-BASED TESTS
// ==========================================

func TestCalculator_MethodsExist(t *testing.T) {
	calc := &Calculator{}
	calcType := reflect.TypeOf(calc)
	
	expectedMethods := []string{"Add", "Subtract", "Multiply", "Divide"}
	
	for _, methodName := range expectedMethods {
		method, exists := calcType.MethodByName(methodName)
		if !exists {
			t.Errorf("Method %s does not exist", methodName)
			continue
		}
		
		// Check method signature
		if method.Type.NumIn() != 3 { // receiver + 2 parameters
			t.Errorf("Method %s should have 2 parameters, got %d", methodName, method.Type.NumIn()-1)
		}
	}
}

// ==========================================
// ⏱️ TIMEOUT TESTS
// ==========================================

func TestWithTimeout(t *testing.T) {
	// Simulate a function that might hang
	slowFunction := func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}
	
	// Test with sufficient timeout
	t.Run("sufficient timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			done <- slowFunction()
		}()
		
		select {
		case err := <-done:
			if err != nil {
				t.Errorf("Function failed: %v", err)
			}
		case <-ctx.Done():
			t.Error("Function timed out")
		}
	})
	
	// Test with insufficient timeout
	t.Run("insufficient timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			done <- slowFunction()
		}()
		
		select {
		case <-done:
			t.Error("Function should have timed out")
		case <-ctx.Done():
			// This is expected
		}
	})
}

// ==========================================
// 🔄 CLEANUP TESTS
// ==========================================

func TestWithCleanup(t *testing.T) {
	// Setup
	tempData := make(map[string]string)
	tempData["test"] = "value"
	
	// Register cleanup
	t.Cleanup(func() {
		// This will run after the test completes
		for key := range tempData {
			delete(tempData, key)
		}
	})
	
	// Test logic
	if tempData["test"] != "value" {
		t.Error("Setup failed")
	}
	
	// No need to manually cleanup - t.Cleanup will handle it
}

// ==========================================
// 📊 HELPER FUNCTIONS
// ==========================================

func assertEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func assertContains(t *testing.T, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("Expected '%s' to contain '%s'", str, substr)
	}
}
