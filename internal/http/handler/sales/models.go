package saleshandler

import (
	"strings"
	"time"

	"coto-back/internal/sales"
)

type createSaleRequest struct {
	Model              string `json:"model"`
	DistributionCenter string `json:"distributionCenter"`
	Units              int    `json:"units"`
}

type saleResponse struct {
	ID                 int    `json:"id"`
	Model              string `json:"model"`
	DistributionCenter string `json:"distributionCenter"`
	Units              int    `json:"units"`
	UnitPriceCents     int    `json:"unitPriceCents"`
	TotalCents         int    `json:"totalCents"`
	CreatedAt          string `json:"createdAt"`
}

func newSaleResponse(sale sales.Sale) saleResponse {
	return saleResponse{
		ID:                 sale.ID,
		Model:              string(sale.Model),
		DistributionCenter: string(sale.DistributionCenter),
		Units:              sale.Units,
		UnitPriceCents:     sale.UnitPriceCents,
		TotalCents:         sale.TotalCents,
		CreatedAt:          sale.CreatedAt.Format(time.RFC3339),
	}
}

func allowedModels() string {
	models := sales.Models()
	values := make([]string, 0, len(models))

	for _, model := range models {
		values = append(values, string(model))
	}

	return strings.Join(values, ", ")
}

func allowedDistributionCenters() string {
	centers := sales.DistributionCenters()
	values := make([]string, 0, len(centers))

	for _, center := range centers {
		values = append(values, string(center))
	}

	return strings.Join(values, ", ")
}
