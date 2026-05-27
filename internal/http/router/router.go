package router

import (
	"net/http"

	"coto-back/internal/http/handler"

	"github.com/go-chi/chi/v5"
)

func New() http.Handler {

	r := chi.NewRouter()

	healthHandler := handler.NewHealthHandler()
	r.Get("/", healthHandler.Check)

	return r
}
