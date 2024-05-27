package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	// запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// необходимые проверки
	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status, fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, status))
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenUnknownCity(t *testing.T) {
	// запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tagil", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверка статуса
	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusBadRequest, fmt.Sprintf("expected status code: %d, got %d", http.StatusBadRequest, status))
	// проверка тела ответа
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body, fmt.Sprintf("expected response body: %s, got %s", "wrong city value", body))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// проверка статуса
	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status, fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, status))
	// проверка тела ответа
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Len(t, strings.Split(body, ","), totalCount)
}
