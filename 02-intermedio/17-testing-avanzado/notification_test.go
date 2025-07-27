// 🧪 Tests Mocking: Sistema de Notificaciones
// Archivo: notification_test.go
// Ejecutar con: go test -v -run TestNotification

package main

import (
	"errors"
	"testing"
)

// ==========================================
// 🧪 MOCKING TESTS - NOTIFICACIONES
// ==========================================

func TestNotificationService_SendWelcomeNotification_Success(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	userID := "user123"
	email := "test@example.com"
	phone := "+1234567890"

	// Act
	err := service.SendWelcomeNotification(userID, email, phone)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify email was sent
	if len(emailMock.SendEmailCalls) != 1 {
		t.Errorf("Expected 1 email call, got %d", len(emailMock.SendEmailCalls))
	} else {
		emailCall := emailMock.SendEmailCalls[0]
		if emailCall.To != email {
			t.Errorf("Expected email to %s, got %s", email, emailCall.To)
		}
		if emailCall.Subject != "¡Bienvenido!" {
			t.Errorf("Expected subject '¡Bienvenido!', got %s", emailCall.Subject)
		}
	}

	// Verify SMS was sent
	if len(smsMock.SendSMSCalls) != 1 {
		t.Errorf("Expected 1 SMS call, got %d", len(smsMock.SendSMSCalls))
	} else {
		smsCall := smsMock.SendSMSCalls[0]
		if smsCall.Phone != phone {
			t.Errorf("Expected SMS to %s, got %s", phone, smsCall.Phone)
		}
	}

	// Verify push notification was sent
	if len(pushMock.SendPushCalls) != 1 {
		t.Errorf("Expected 1 push call, got %d", len(pushMock.SendPushCalls))
	} else {
		pushCall := pushMock.SendPushCalls[0]
		if pushCall.UserID != userID {
			t.Errorf("Expected push to user %s, got %s", userID, pushCall.UserID)
		}
	}
}

func TestNotificationService_SendWelcomeNotification_EmailFailure(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMock{SendEmailError: errors.New("email service down")}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")

	// Assert
	if err == nil {
		t.Error("Expected error when email service fails")
	}

	// Should still attempt to send email
	if len(emailMock.SendEmailCalls) != 1 {
		t.Errorf("Expected 1 email call even with failure, got %d", len(emailMock.SendEmailCalls))
	}

	// Should not proceed to SMS and push if email fails
	if len(smsMock.SendSMSCalls) != 0 {
		t.Errorf("Expected no SMS calls when email fails, got %d", len(smsMock.SendSMSCalls))
	}
	if len(pushMock.SendPushCalls) != 0 {
		t.Errorf("Expected no push calls when email fails, got %d", len(pushMock.SendPushCalls))
	}
}

func TestNotificationService_SendPasswordResetNotification_Success(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	userID := "user123"
	email := "test@example.com"

	// Act
	err := service.SendPasswordResetNotification(userID, email)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify only email was sent
	if len(emailMock.SendEmailCalls) != 1 {
		t.Errorf("Expected 1 email call, got %d", len(emailMock.SendEmailCalls))
	} else {
		emailCall := emailMock.SendEmailCalls[0]
		if emailCall.To != email {
			t.Errorf("Expected email to %s, got %s", email, emailCall.To)
		}
		if emailCall.Subject != "Restablecimiento de contraseña" {
			t.Errorf("Expected password reset subject, got %s", emailCall.Subject)
		}
	}

	// Should not send SMS or push for password reset
	if len(smsMock.SendSMSCalls) != 0 {
		t.Errorf("Expected no SMS calls for password reset, got %d", len(smsMock.SendSMSCalls))
	}
	if len(pushMock.SendPushCalls) != 0 {
		t.Errorf("Expected no push calls for password reset, got %d", len(pushMock.SendPushCalls))
	}
}

func TestNotificationService_SendOrderConfirmation_Success(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	userID := "user123"
	email := "test@example.com"
	phone := "+1234567890"
	orderID := "order-456"

	// Act
	err := service.SendOrderConfirmation(userID, email, phone, orderID)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify email was sent
	if len(emailMock.SendEmailCalls) != 1 {
		t.Errorf("Expected 1 email call, got %d", len(emailMock.SendEmailCalls))
	} else {
		emailCall := emailMock.SendEmailCalls[0]
		if emailCall.To != email {
			t.Errorf("Expected email to %s, got %s", email, emailCall.To)
		}
		if emailCall.Subject != "Confirmación de pedido" {
			t.Errorf("Expected order confirmation subject, got %s", emailCall.Subject)
		}
		if !contains(emailCall.Body, orderID) {
			t.Errorf("Expected email body to contain order ID %s", orderID)
		}
	}

	// Verify SMS was sent
	if len(smsMock.SendSMSCalls) != 1 {
		t.Errorf("Expected 1 SMS call, got %d", len(smsMock.SendSMSCalls))
	} else {
		smsCall := smsMock.SendSMSCalls[0]
		if smsCall.Phone != phone {
			t.Errorf("Expected SMS to %s, got %s", phone, smsCall.Phone)
		}
		if !contains(smsCall.Message, orderID) {
			t.Errorf("Expected SMS message to contain order ID %s", orderID)
		}
	}

	// Should not send push for order confirmation
	if len(pushMock.SendPushCalls) != 0 {
		t.Errorf("Expected no push calls for order confirmation, got %d", len(pushMock.SendPushCalls))
	}
}

func TestNotificationService_SendOrderConfirmation_SMSFailure(t *testing.T) {
	// Arrange
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{SendSMSError: errors.New("SMS service down")}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	// Act
	err := service.SendOrderConfirmation("user123", "test@example.com", "+1234567890", "order-456")

	// Assert
	if err == nil {
		t.Error("Expected error when SMS service fails")
	}

	// Should still send email
	if len(emailMock.SendEmailCalls) != 1 {
		t.Errorf("Expected 1 email call even with SMS failure, got %d", len(emailMock.SendEmailCalls))
	}

	// Should attempt SMS
	if len(smsMock.SendSMSCalls) != 1 {
		t.Errorf("Expected 1 SMS call attempt, got %d", len(smsMock.SendSMSCalls))
	}
}

// ==========================================
// 🎯 MOCK VERIFICATION TESTS
// ==========================================

func TestEmailSenderMock_CallTracking(t *testing.T) {
	mock := &EmailSenderMock{}

	// Should start with no calls
	if len(mock.SendEmailCalls) != 0 {
		t.Errorf("Mock should start with 0 calls, got %d", len(mock.SendEmailCalls))
	}

	// Make some calls
	_ = mock.SendEmail("user1@test.com", "Subject 1", "Body 1")
	_ = mock.SendEmail("user2@test.com", "Subject 2", "Body 2")

	// Verify calls were tracked
	if len(mock.SendEmailCalls) != 2 {
		t.Errorf("Expected 2 calls, got %d", len(mock.SendEmailCalls))
	}

	// Verify call details
	firstCall := mock.SendEmailCalls[0]
	if firstCall.To != "user1@test.com" {
		t.Errorf("First call To expected 'user1@test.com', got '%s'", firstCall.To)
	}
	if firstCall.Subject != "Subject 1" {
		t.Errorf("First call Subject expected 'Subject 1', got '%s'", firstCall.Subject)
	}

	secondCall := mock.SendEmailCalls[1]
	if secondCall.To != "user2@test.com" {
		t.Errorf("Second call To expected 'user2@test.com', got '%s'", secondCall.To)
	}
}

func TestEmailSenderMock_ErrorSimulation(t *testing.T) {
	mock := &EmailSenderMock{
		SendEmailError: errors.New("simulated error"),
	}

	err := mock.SendEmail("test@example.com", "Subject", "Body")

	if err == nil {
		t.Error("Expected error from mock, got nil")
	}

	if err.Error() != "simulated error" {
		t.Errorf("Expected 'simulated error', got '%s'", err.Error())
	}

	// Call should still be tracked even with error
	if len(mock.SendEmailCalls) != 1 {
		t.Errorf("Expected call to be tracked even with error, got %d calls", len(mock.SendEmailCalls))
	}
}

// ==========================================
// 🧪 BENCHMARK TESTS
// ==========================================

func BenchmarkNotificationService_SendWelcomeNotification(b *testing.B) {
	emailMock := &EmailSenderMock{}
	smsMock := &SMSSenderMock{}
	pushMock := &PushNotificationSenderMock{}

	service := NewNotificationService(emailMock, smsMock, pushMock)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.SendWelcomeNotification("user123", "test@example.com", "+1234567890")
	}
}

func BenchmarkEmailSenderMock_SendEmail(b *testing.B) {
	mock := &EmailSenderMock{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mock.SendEmail("test@example.com", "Subject", "Body")
	}
}

// ==========================================
// 🛠️ HELPER FUNCTIONS
// ==========================================

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		(len(s) > len(substr) && s[len(s)-len(substr):] == substr) ||
		(len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
