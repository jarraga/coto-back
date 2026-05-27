package router

import (
	"net/http"

	"coto-back/internal/http/handler/health"
	saleshandler "coto-back/internal/http/handler/sales"
	"coto-back/internal/sales"

	"github.com/go-chi/chi/v5"
)

func New(store *sales.Store) http.Handler {

	r := chi.NewRouter()

	healthHandler := health.NewHandler()
	r.Get("/", healthHandler.Check)

	salesHandler := saleshandler.NewHandler(store)
	r.Post("/sales", salesHandler.Create)

	return r
}
