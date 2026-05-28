package middleware

import (
	"fmt"
	"net/http"
	"time"

	"coto-back/internal/logging"
)

const requestDurationHeader = "X-Request-Duration"

type responseWriter struct {
	http.ResponseWriter
	start         time.Time
	statusCode    int
	headerWritten bool
}

func newResponseWriter(w http.ResponseWriter, start time.Time) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		start:          start,
		statusCode:     http.StatusOK,
	}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *responseWriter) Write(body []byte) (int, error) {
	w.writeHeader()

	return w.ResponseWriter.Write(body)
}

func (w *responseWriter) Flush() {

	w.writeHeader()

	flusher, ok := w.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}

	flusher.Flush()
}

func (w *responseWriter) writeHeader() {
	if w.headerWritten {
		return
	}

	duration := time.Since(w.start)
	w.Header().Set(requestDurationHeader, formatRequestDuration(duration))
	w.ResponseWriter.WriteHeader(w.statusCode)
	w.headerWritten = true
}

func RequestTime(logger *logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			responseWriter := newResponseWriter(w, start)

			next.ServeHTTP(responseWriter, r)
			responseWriter.writeHeader()

			duration := time.Since(start)
			formattedDuration := formatRequestDuration(duration)

			logger.Printf(
				"%s %s completed in %s with status %d",
				r.Method,
				r.URL.Path,
				formattedDuration,
				responseWriter.statusCode,
			)
		})
	}
}

func formatRequestDuration(duration time.Duration) string {
	if duration < time.Millisecond {
		microseconds := duration.Microseconds()
		if microseconds < 1 {
			microseconds = 1
		}

		return fmt.Sprintf("%dus", microseconds)
	}

	return fmt.Sprintf("%dms", duration.Milliseconds())
}
