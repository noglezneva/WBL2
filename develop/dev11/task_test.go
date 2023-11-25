package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	store := NewEventStore()
	handler := createUserEventHandler(store)

	// Создание запроса с валидным ID пользователя
	userID := 123
	body, _ := json.Marshal(userID)
	request := httptest.NewRequest(http.MethodPost, "/users/create", bytes.NewReader(body))
	response := httptest.NewRecorder()

	handler(response, request)

	// Проверка кода ответа
	if response.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, response.Code)
	}

	// Проверка, что пользователь был создан в хранилище
	userEvents := store.events[userID]
	if userEvents == nil {
		t.Errorf("User events not found in the store")
	}
}

func TestCreateUserHandler_InvalidUserID(t *testing.T) {
	store := NewEventStore()
	handler := createUserEventHandler(store)

	// Создание запроса с некорректным ID пользователя (не целое число)
	userID := "invalid"
	body, _ := json.Marshal(userID)
	request := httptest.NewRequest(http.MethodPost, "/users/create", bytes.NewReader(body))
	response := httptest.NewRecorder()

	handler(response, request)

	// Проверка кода ответа
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	// Проверка, что пользователь не был создан в хранилище
	if len(store.events) > 0 {
		t.Errorf("Unexpected user events in the store")
	}
}

func TestCreateUserHandler_EmptyRequestBody(t *testing.T) {
	store := NewEventStore()
	handler := createUserEventHandler(store)

	// Создание запроса с пустым телом запроса
	request := httptest.NewRequest(http.MethodPost, "/users/create", nil)
	response := httptest.NewRecorder()

	handler(response, request)

	// Проверка кода ответа
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	// Проверка, что пользователь не был создан в хранилище
	if len(store.events) > 0 {
		t.Errorf("Unexpected user events in the store")
	}
}
