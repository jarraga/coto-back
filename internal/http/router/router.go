package router

import (
	"net/http"

	"coto-back/internal/http/handler/health"
	saleshandler "coto-back/internal/http/handler/sales"
	mw "coto-back/internal/http/middleware"
	"coto-back/internal/sales"

	"github.com/go-chi/chi/v5"
)

func New(service *sales.Service) http.Handler {

	r := chi.NewRouter()
	r.Use(mw.RequestTime)

	healthHandler := health.NewHandler()
	r.Get("/", healthHandler.Check)

	salesHandler := saleshandler.NewHandler(service)
	r.Post("/sales", salesHandler.Create)
	r.Get("/sales/volume", salesHandler.TotalVolume)
	r.Get("/sales/volume/by-center", salesHandler.VolumeByCenter)

	return r
}
