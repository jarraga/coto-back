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
	service *sales.Service
}

func NewHandler(service *sales.Service) Handler {
	return Handler{service: service}
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
		// UnitPriceCents also validates the car model.
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

	sale := h.service.Create(saleToSave)

	httphelper.JSON(w, http.StatusCreated, newSaleResponse(sale))
}

func (h Handler) TotalVolume(w http.ResponseWriter, r *http.Request) {
	volume := h.service.TotalVolume()

	httphelper.JSON(w, http.StatusOK, newTotalVolumeResponse(volume))
}

func (h Handler) VolumeByCenter(w http.ResponseWriter, r *http.Request) {
	volumes := h.service.VolumeByCenter()

	httphelper.JSON(w, http.StatusOK, newCenterVolumeResponse(volumes))
}

func (h Handler) ModelPercentagesByCenter(w http.ResponseWriter, r *http.Request) {
	percentages := h.service.ModelPercentagesByCenter()

	httphelper.JSON(w, http.StatusOK, newCenterModelPercentageResponse(percentages))
}
