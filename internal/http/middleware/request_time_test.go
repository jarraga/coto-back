package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"coto-back/internal/logging"
)

func TestRequestTime(t *testing.T) {

	logger := logging.New()
	middleware := RequestTime(logger)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	request := httptest.NewRequest(http.MethodGet, "/test", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}

	if response.Header().Get(requestDurationHeader) == "" {
		t.Fatal("expected request duration header")
	}
}

func TestFormatRequestDurationInMicroseconds(t *testing.T) {

	duration := formatRequestDuration(500 * time.Microsecond)

	if duration != "500us" {
		t.Fatalf("expected 500us, got %s", duration)
	}
}

func TestFormatRequestDurationInMilliseconds(t *testing.T) {

	duration := formatRequestDuration(2 * time.Millisecond)

	if duration != "2ms" {
		t.Fatalf("expected 2ms, got %s", duration)
	}
}
