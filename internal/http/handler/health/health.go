package health

import (
	"net/http"
	"time"

	httphelper "coto-back/internal/http"
)

var startedAt = time.Now()

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Check(w http.ResponseWriter, r *http.Request) {

	httphelper.JSON(w, http.StatusOK, map[string]any{
		"status":        "ok",
		"executionTime": time.Since(startedAt).Round(time.Second).String(),
	})
}
