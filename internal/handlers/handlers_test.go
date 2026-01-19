package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortener/internal/storage"
)

// 1. Test the Home/Root Route
func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	RootHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Welcome") {
		t.Errorf("Expected welcome message, got %s", rr.Body.String())
	}
}

// 2. Test Shorten Route (Success)
func TestShortenHandler_Success(t *testing.T) {
	body := []byte(`{"url": "https://google.com"}`)
	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	ShortenHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}

// 3. Test Shorten Route (Wrong Method - handles the 'if r.Method != POST' block)
func TestShortenHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/shorten", nil)
	rr := httptest.NewRecorder()

	ShortenHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405, got %d", rr.Code)
	}
}

// 4. Test Shorten Route (Bad JSON - handles the 'if err != nil' block)
func TestShortenHandler_InvalidJSON(t *testing.T) {
	body := []byte(`{"url": "broken-json"`) // Missing closing bracket
	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	ShortenHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", rr.Code)
	}
}

// 5. Test Redirect Route (Success)
func TestRedirectHandler_Success(t *testing.T) {
	id := storage.SaveURL("https://apple.com")
	req := httptest.NewRequest("GET", "/r/"+id, nil)
	rr := httptest.NewRecorder()

	RedirectHandler(rr, req)

	if rr.Code != http.StatusFound { // 302 Found
		t.Errorf("Expected 302, got %d", rr.Code)
	}
}

// 6. Test Redirect Route (ID Not Found - handles the 'if !exists' block)
func TestRedirectHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/r/invalidID999", nil)
	rr := httptest.NewRecorder()

	RedirectHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", rr.Code)
	}
}

// 7. Test Batch Route (Success)
func TestBatchHandler_Success(t *testing.T) {
	body := []byte(`["https://github.com", "https://go.dev"]`)
	req := httptest.NewRequest("POST", "/batch", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	BatchHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}

// 8. Test Batch Route (Invalid JSON)
func TestBatchHandler_InvalidJSON(t *testing.T) {
	body := []byte(`["broken-array"`)
	req := httptest.NewRequest("POST", "/batch", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	BatchHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", rr.Code)
	}
}
