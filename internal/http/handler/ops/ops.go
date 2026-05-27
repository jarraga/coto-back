package ops

import (
	"net/http"

	httphelper "coto-back/internal/http"
	"coto-back/internal/sales"
)

type Handler struct {
	service *sales.Service
}

func NewHandler(service *sales.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) ClearStore(w http.ResponseWriter, r *http.Request) {

	h.service.Clear()

	httphelper.JSON(w, http.StatusOK, map[string]string{
		"status": "store cleared",
	})
}

func (h Handler) SeedStore(w http.ResponseWriter, r *http.Request) {

	h.service.Seed()

	httphelper.JSON(w, http.StatusOK, map[string]string{
		"status": "store seeded",
	})
}
