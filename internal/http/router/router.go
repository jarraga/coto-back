package router

import (
	"net/http"

	"coto-back/internal/http/handler/health"
	"coto-back/internal/http/handler/logs"
	"coto-back/internal/http/handler/ops"
	saleshandler "coto-back/internal/http/handler/sales"
	mw "coto-back/internal/http/middleware"
	"coto-back/internal/logging"
	"coto-back/internal/sales"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func New(service *sales.Service, logger *logging.Logger) http.Handler {

	r := chi.NewRouter()

	// Allow Swagger Editor to execute requests against the deployed API.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://editor.swagger.io"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
		},
		ExposedHeaders: []string{"X-Request-Duration"},
		MaxAge:         300,
	}))
	r.Use(mw.RequestTime(logger))

	healthHandler := health.NewHandler()
	r.Get("/", healthHandler.Check)

	logsHandler := logs.NewHandler(logger)
	r.Get("/logs/stream", logsHandler.Stream)

	opsHandler := ops.NewHandler(service)
	r.Delete("/ops/store", opsHandler.ClearStore)
	r.Post("/ops/store/seed", opsHandler.SeedStore)

	salesHandler := saleshandler.NewHandler(service)
	r.Post("/sales", salesHandler.Create)
	r.Get("/sales/volume", salesHandler.TotalVolume)
	r.Get("/sales/volume/by-center", salesHandler.VolumeByCenter)
	r.Get("/sales/model-percentages/by-center", salesHandler.ModelPercentagesByCenter)

	PrintAvailableRoutes(r, logger)

	return r
}
