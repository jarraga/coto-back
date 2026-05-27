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
	UnitPriceAmount    int    `json:"unitPriceAmount"`
	TotalAmount        int    `json:"totalAmount"`
	CreatedAt          string `json:"createdAt"`
}

type totalVolumeResponse struct {
	Units       int `json:"units"`
	TotalAmount int `json:"totalAmount"`
}

type centerVolumeResponse struct {
	DistributionCenter string `json:"distributionCenter"`
	Units              int    `json:"units"`
	TotalAmount        int    `json:"totalAmount"`
}

func newSaleResponse(sale sales.Sale) saleResponse {
	return saleResponse{
		ID:                 sale.ID,
		Model:              string(sale.Model),
		DistributionCenter: string(sale.DistributionCenter),
		Units:              sale.Units,
		UnitPriceAmount:    centsToAmount(sale.UnitPriceCents),
		TotalAmount:        centsToAmount(sale.TotalCents),
		CreatedAt:          sale.CreatedAt.Format(time.RFC3339),
	}
}

func newTotalVolumeResponse(volume sales.TotalVolume) totalVolumeResponse {
	return totalVolumeResponse{
		Units:       volume.Units,
		TotalAmount: centsToAmount(volume.TotalCents),
	}
}

func newCenterVolumeResponses(volumes []sales.CenterVolume) []centerVolumeResponse {
	responses := make([]centerVolumeResponse, 0, len(volumes))

	for _, volume := range volumes {
		responses = append(responses, centerVolumeResponse{
			DistributionCenter: string(volume.DistributionCenter),
			Units:              volume.Units,
			TotalAmount:        centsToAmount(volume.TotalCents),
		})
	}

	return responses
}

func centsToAmount(cents int) int {
	return cents / 100
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
