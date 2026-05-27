package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

var startedAt = time.Now()

type HealthHandler struct{}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (h HealthHandler) Check(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"status":        "ok",
		"executionTime": time.Since(startedAt).Round(time.Second).String(),
	})
}
