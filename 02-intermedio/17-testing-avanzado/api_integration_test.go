package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// ========================================
// Integration Tests para API Client
// ========================================

func TestAPIClient_GetUser_Integration(t *testing.T) {
	// Crear servidor de prueba
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar método y path
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/users/1" {
			t.Errorf("Expected path /users/1, got %s", r.URL.Path)
		}

		// Respuesta simulada
		user := User{
			ID:    1,
			Name:  "Test User",
			Email: "test@example.com",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	// Test
	client := NewAPIClient(server.URL)
	user, err := client.GetUser(1)

	if err != nil {
		t.Errorf("GetUser should not return error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", user.ID)
	}

	if user.Name != "Test User" {
		t.Errorf("Expected user name 'Test User', got '%s'", user.Name)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected user email 'test@example.com', got '%s'", user.Email)
	}
}

func TestAPIClient_GetUser_NotFound(t *testing.T) {
	// Servidor que simula usuario no encontrado
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "User not found"}`))
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	user, err := client.GetUser(999)

	if err == nil {
		t.Error("GetUser should return error for 404 response")
	}

	if user != nil {
		t.Error("GetUser should return nil user for 404 response")
	}
}

func TestAPIClient_CreateUser_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar método
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.URL.Path != "/users" {
			t.Errorf("Expected path /users, got %s", r.URL.Path)
		}

		// Verificar Content-Type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}

		// Leer y verificar request body
		var request CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("Error decoding request body: %v", err)
		}

		if request.Name != "New User" {
			t.Errorf("Expected name 'New User', got '%s'", request.Name)
		}

		if request.Email != "newuser@example.com" {
			t.Errorf("Expected email 'newuser@example.com', got '%s'", request.Email)
		}

		// Respuesta
		user := User{
			ID:    123,
			Name:  request.Name,
			Email: request.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	request := CreateUserRequest{
		Name:  "New User",
		Email: "newuser@example.com",
	}

	user, err := client.CreateUser(request)

	if err != nil {
		t.Errorf("CreateUser should not return error: %v", err)
	}

	if user.ID != 123 {
		t.Errorf("Expected user ID 123, got %d", user.ID)
	}

	if user.Name != "New User" {
		t.Errorf("Expected user name 'New User', got '%s'", user.Name)
	}
}

func TestAPIClient_UpdateUser_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		if r.URL.Path != "/users/1" {
			t.Errorf("Expected path /users/1, got %s", r.URL.Path)
		}

		var request CreateUserRequest
		json.NewDecoder(r.Body).Decode(&request)

		user := User{
			ID:    1,
			Name:  request.Name,
			Email: request.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	request := CreateUserRequest{
		Name:  "Updated User",
		Email: "updated@example.com",
	}

	user, err := client.UpdateUser(1, request)

	if err != nil {
		t.Errorf("UpdateUser should not return error: %v", err)
	}

	if user.Name != "Updated User" {
		t.Errorf("Expected updated name 'Updated User', got '%s'", user.Name)
	}
}

func TestAPIClient_DeleteUser_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		if r.URL.Path != "/users/1" {
			t.Errorf("Expected path /users/1, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	err := client.DeleteUser(1)

	if err != nil {
		t.Errorf("DeleteUser should not return error: %v", err)
	}
}

func TestAPIClient_ListUsers_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/users" {
			t.Errorf("Expected path /users, got %s", r.URL.Path)
		}

		users := []User{
			{ID: 1, Name: "User 1", Email: "user1@example.com"},
			{ID: 2, Name: "User 2", Email: "user2@example.com"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	users, err := client.ListUsers()

	if err != nil {
		t.Errorf("ListUsers should not return error: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0].ID != 1 {
		t.Errorf("Expected first user ID 1, got %d", users[0].ID)
	}
}

// ========================================
// Tests de Timeout y Error Handling
// ========================================

func TestAPIClient_Timeout(t *testing.T) {
	// Servidor que tarda mucho en responder
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond) // Más que el timeout del cliente
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Cliente con timeout muy corto
	client := NewAPIClient(server.URL)
	client.timeout = 10 * time.Millisecond
	client.client.Timeout = 10 * time.Millisecond

	_, err := client.GetUser(1)

	if err == nil {
		t.Error("GetUser should return timeout error")
	}
}

func TestAPIClient_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal server error"}`))
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	_, err := client.GetUser(1)

	if err == nil {
		t.Error("GetUser should return error for 500 response")
	}
}

func TestAPIClient_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"invalid": json}`)) // JSON inválido
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	_, err := client.GetUser(1)

	if err == nil {
		t.Error("GetUser should return error for invalid JSON")
	}
}

// ========================================
// Tests de Comportamiento de Red
// ========================================

func TestAPIClient_NetworkError(t *testing.T) {
	// Cliente apuntando a un servidor que no existe
	client := NewAPIClient("http://localhost:99999")

	_, err := client.GetUser(1)

	if err == nil {
		t.Error("GetUser should return network error")
	}
}

func TestAPIClient_UserAgent_Headers(t *testing.T) {
	headerCalled := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar headers personalizados si los hubiera
		headerCalled = true

		user := User{ID: 1, Name: "Test", Email: "test@example.com"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	_, err := client.GetUser(1)

	if err != nil {
		t.Errorf("GetUser should not return error: %v", err)
	}

	if !headerCalled {
		t.Error("Handler should have been called")
	}
}

// ========================================
// Test de Flujo Completo (End-to-End)
// ========================================

func TestAPIClient_CompleteUserFlow(t *testing.T) {
	users := make(map[int]User)
	nextID := 1

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodPost:
			if r.URL.Path == "/users" {
				var request CreateUserRequest
				json.NewDecoder(r.Body).Decode(&request)

				user := User{
					ID:    nextID,
					Name:  request.Name,
					Email: request.Email,
				}
				users[nextID] = user
				nextID++

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(user)
			}

		case http.MethodGet:
			if r.URL.Path == "/users" {
				var userList []User
				for _, user := range users {
					userList = append(userList, user)
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(userList)
			} else {
				// GET /users/{id}
				var userID int
				fmt.Sscanf(r.URL.Path, "/users/%d", &userID)

				if user, exists := users[userID]; exists {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			}

		case http.MethodDelete:
			var userID int
			fmt.Sscanf(r.URL.Path, "/users/%d", &userID)

			if _, exists := users[userID]; exists {
				delete(users, userID)
				w.WriteHeader(http.StatusNoContent)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)

	// 1. Crear usuario
	createRequest := CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, err := client.CreateUser(createRequest)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if createdUser.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", createdUser.ID)
	}

	// 2. Obtener usuario creado
	fetchedUser, err := client.GetUser(createdUser.ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if fetchedUser.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", fetchedUser.Name)
	}

	// 3. Listar usuarios
	userList, err := client.ListUsers()
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	if len(userList) != 1 {
		t.Errorf("Expected 1 user in list, got %d", len(userList))
	}

	// 4. Eliminar usuario
	err = client.DeleteUser(createdUser.ID)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// 5. Verificar que el usuario fue eliminado
	_, err = client.GetUser(createdUser.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user")
	}

	// 6. Verificar lista vacía
	userList, err = client.ListUsers()
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	if len(userList) != 0 {
		t.Errorf("Expected 0 users in list, got %d", len(userList))
	}
}

// ========================================
// Benchmarks para API Client
// ========================================

func BenchmarkAPIClient_GetUser(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, Name: "Test", Email: "test@example.com"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.GetUser(1)
	}
}

func BenchmarkAPIClient_CreateUser(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request CreateUserRequest
		json.NewDecoder(r.Body).Decode(&request)

		user := User{ID: 1, Name: request.Name, Email: request.Email}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)
	request := CreateUserRequest{Name: "Test", Email: "test@example.com"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.CreateUser(request)
	}
}
