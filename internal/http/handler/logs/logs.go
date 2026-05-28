package logs

import (
	"fmt"
	"net/http"

	"coto-back/internal/logging"
)

type Handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) Handler {

	return Handler{logger: logger}
}

func (h Handler) Stream(w http.ResponseWriter, r *http.Request) {

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming is not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	flusher.Flush()

	subscriber := h.logger.Subscribe()
	defer h.logger.Unsubscribe(subscriber)

	for {
		select {
		case <-r.Context().Done():
			return
		case message := <-subscriber:
			fmt.Fprintln(w, message)
			flusher.Flush()
		}
	}
}
