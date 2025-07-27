package main

import (
	"errors"
	"testing"
)

// ========================================
// Tests para Sistema de Notificaciones (Mocking)
// ========================================

// Mocks mejorados para testing
type EmailSenderMockAdvanced struct {
	calls      []EmailCall
	shouldFail bool
	callCount  int
}

func (m *EmailSenderMockAdvanced) SendEmail(to, subject, body string) error {
	m.callCount++
	m.calls = append(m.calls, EmailCall{To: to, Subject: subject, Body: body})
	if m.shouldFail {
		return errors.New("email service unavailable")
	}
	return nil
}

func (m *EmailSenderMockAdvanced) GetCalls() []EmailCall {
	return m.calls
}

func (m *EmailSenderMockAdvanced) SetShouldFail(fail bool) {
	m.shouldFail = fail
}

func (m *EmailSenderMockAdvanced) GetCallCount() int {
	return m.callCount
}

func (m *EmailSenderMockAdvanced) Reset() {
	m.calls = []EmailCall{}
	m.callCount = 0
	m.shouldFail = false
}

type SMSSenderMockAdvanced struct {
	calls      []SMSCall
	shouldFail bool
	callCount  int
}

func (m *SMSSenderMockAdvanced) SendSMS(to, message string) error {
	m.callCount++
	m.calls = append(m.calls, SMSCall{To: to, Message: message})
	if m.shouldFail {
		return errors.New("SMS service unavailable")
	}
	return nil
}

type PushSenderMockAdvanced struct {
	calls      []PushCall
	shouldFail bool
	callCount  int
}

func (m *PushSenderMockAdvanced) SendPushNotification(userID, title, body string) error {
	m.callCount++
	m.calls = append(m.calls, PushCall{UserID: userID, Title: title, Body: body})
	if m.shouldFail {
		return errors.New("push service unavailable")
	}
	return nil
}

func TestNotificationService_SendWelcomeNotification_Success(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMockAdvanced{}
	smsMock := &SMSSenderMockAdvanced{}
	pushMock := &PushSenderMockAdvanced{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")

	// Assert
	if err != nil {
		t.Errorf("SendWelcomeNotification should not return error, got: %v", err)
	}

	// Verify email was sent
	if emailMock.GetCallCount() != 1 {
		t.Errorf("Email should be called once, got %d calls", emailMock.GetCallCount())
	}

	emailCalls := emailMock.GetCalls()
	if len(emailCalls) != 1 {
		t.Errorf("Expected 1 email call, got %d", len(emailCalls))
	}

	if emailCalls[0].To != "test@example.com" {
		t.Errorf("Email sent to wrong address: %s", emailCalls[0].To)
	}

	if emailCalls[0].Subject != "Bienvenido" {
		t.Errorf("Wrong email subject: %s", emailCalls[0].Subject)
	}

	// Verify SMS was sent
	if smsMock.callCount != 1 {
		t.Errorf("SMS should be called once, got %d calls", smsMock.callCount)
	}

	// Verify Push was sent
	if pushMock.callCount != 1 {
		t.Errorf("Push should be called once, got %d calls", pushMock.callCount)
	}
}

func TestNotificationService_SendWelcomeNotification_EmailFails(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMockAdvanced{}
	emailMock.SetShouldFail(true)
	smsMock := &SMSSenderMockAdvanced{}
	pushMock := &PushSenderMockAdvanced{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")

	// Assert
	if err == nil {
		t.Error("SendWelcomeNotification should return error when email fails")
	}

	// Verify email was attempted
	if emailMock.GetCallCount() != 1 {
		t.Errorf("Email should be called once, got %d calls", emailMock.GetCallCount())
	}

	// Verify SMS and Push were not called due to early return
	if smsMock.callCount != 0 {
		t.Errorf("SMS should not be called when email fails, got %d calls", smsMock.callCount)
	}

	if pushMock.callCount != 0 {
		t.Errorf("Push should not be called when email fails, got %d calls", pushMock.callCount)
	}
}

func TestNotificationService_SendPasswordResetNotification(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMockAdvanced{}
	smsMock := &SMSSenderMockAdvanced{}
	pushMock := &PushSenderMockAdvanced{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendPasswordResetNotification("user123", "test@example.com")

	// Assert
	if err != nil {
		t.Errorf("SendPasswordResetNotification should not return error, got: %v", err)
	}

	// Verify only email was sent
	if emailMock.GetCallCount() != 1 {
		t.Errorf("Email should be called once, got %d calls", emailMock.GetCallCount())
	}

	if smsMock.callCount != 0 {
		t.Errorf("SMS should not be called for password reset, got %d calls", smsMock.callCount)
	}

	if pushMock.callCount != 0 {
		t.Errorf("Push should not be called for password reset, got %d calls", pushMock.callCount)
	}

	// Verify email content
	emailCalls := emailMock.GetCalls()
	if emailCalls[0].Subject != "Resetear Contraseña" {
		t.Errorf("Wrong email subject: %s", emailCalls[0].Subject)
	}
}

func TestNotificationService_SendOrderConfirmation(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMockAdvanced{}
	smsMock := &SMSSenderMockAdvanced{}
	pushMock := &PushSenderMockAdvanced{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendOrderConfirmation("user123", "test@example.com", "+1234567890", "ORD-001")

	// Assert
	if err != nil {
		t.Errorf("SendOrderConfirmation should not return error, got: %v", err)
	}

	// Verify email and SMS were sent, but not push
	if emailMock.GetCallCount() != 1 {
		t.Errorf("Email should be called once, got %d calls", emailMock.GetCallCount())
	}

	if smsMock.callCount != 1 {
		t.Errorf("SMS should be called once, got %d calls", smsMock.callCount)
	}

	if pushMock.callCount != 0 {
		t.Errorf("Push should not be called for order confirmation, got %d calls", pushMock.callCount)
	}

	// Verify order ID is included in messages
	emailCalls := emailMock.GetCalls()
	if !contains(emailCalls[0].Body, "ORD-001") {
		t.Errorf("Email should contain order ID, got: %s", emailCalls[0].Body)
	}

	smsCalls := smsMock.calls
	if !contains(smsCalls[0].Message, "ORD-001") {
		t.Errorf("SMS should contain order ID, got: %s", smsCalls[0].Message)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			func() bool {
				for i := 0; i <= len(s)-len(substr); i++ {
					if s[i:i+len(substr)] == substr {
						return true
					}
				}
				return false
			}())))
}

// ========================================
// Test para verificar comportamiento de mocks
// ========================================

func TestEmailSenderMock_VerifyCallParameters(t *testing.T) {
	mock := &EmailSenderMockAdvanced{}

	// Act
	err1 := mock.SendEmail("user1@test.com", "Subject 1", "Body 1")
	err2 := mock.SendEmail("user2@test.com", "Subject 2", "Body 2")

	// Assert
	if err1 != nil || err2 != nil {
		t.Error("Mock should not return errors by default")
	}

	calls := mock.GetCalls()
	if len(calls) != 2 {
		t.Errorf("Expected 2 calls, got %d", len(calls))
	}

	// Verify first call
	if calls[0].To != "user1@test.com" {
		t.Errorf("First call to should be 'user1@test.com', got '%s'", calls[0].To)
	}

	if calls[0].Subject != "Subject 1" {
		t.Errorf("First call subject should be 'Subject 1', got '%s'", calls[0].Subject)
	}

	// Verify second call
	if calls[1].To != "user2@test.com" {
		t.Errorf("Second call to should be 'user2@test.com', got '%s'", calls[1].To)
	}
}

func TestEmailSenderMock_ErrorSimulation(t *testing.T) {
	mock := &EmailSenderMockAdvanced{}
	mock.SetShouldFail(true)

	// Act
	err := mock.SendEmail("test@example.com", "Test", "Test body")

	// Assert
	if err == nil {
		t.Error("Mock should return error when configured to fail")
	}

	// Verify call was still recorded
	if mock.GetCallCount() != 1 {
		t.Errorf("Call should be recorded even when failing, got %d calls", mock.GetCallCount())
	}
}

// ========================================
// Benchmarks para Mocking
// ========================================

func BenchmarkNotificationService_SendWelcomeNotification(b *testing.B) {
	emailMock := &EmailSenderMockAdvanced{}
	smsMock := &SMSSenderMockAdvanced{}
	pushMock := &PushSenderMockAdvanced{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		emailMock.Reset()
		service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")
	}
}
