package saleshandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	httphelper "coto-back/internal/http"
	"coto-back/internal/sales"
)

type Handler struct {
	store *sales.Store
}

func NewHandler(store *sales.Store) Handler {
	return Handler{store: store}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var request createSaleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httphelper.Error(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	model := sales.CarModel(request.Model)
	unitPriceCents, ok := sales.UnitPriceCents(model)
	if !ok {
		// con esto ya verificamos que sea un modelo válido
		message := fmt.Sprintf("invalid model, only these values are allowed: %s", allowedModels())
		httphelper.Error(w, http.StatusBadRequest, message)
		return
	}

	center := sales.DistributionCenter(request.DistributionCenter)
	if !sales.IsValidDistributionCenter(center) {
		message := fmt.Sprintf("invalid distribution center, only these values are allowed: %s", allowedDistributionCenters())
		httphelper.Error(w, http.StatusBadRequest, message)
		return
	}

	if request.Units <= 0 {
		httphelper.Error(w, http.StatusBadRequest, "units must be greater than zero")
		return
	}

	totalCents := unitPriceCents * request.Units
	createdAt := time.Now().UTC()

	saleToSave := sales.Sale{
		Model:              model,
		DistributionCenter: center,
		Units:              request.Units,
		UnitPriceCents:     unitPriceCents,
		TotalCents:         totalCents,
		CreatedAt:          createdAt,
	}

	sale := h.store.Save(saleToSave)

	httphelper.JSON(w, http.StatusCreated, newSaleResponse(sale))
}
