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
	Centers map[string]centerVolumeDetail `json:"centers"`
}

type centerVolumeDetail struct {
	Units       int `json:"units"`
	TotalAmount int `json:"totalAmount"`
}

type centerModelPercentageResponse struct {
	TotalUnits int                                      `json:"totalUnits"`
	Centers    map[string]centerModelPercentageByCenter `json:"centers"`
}

type centerModelPercentageByCenter struct {
	TotalUnits int                                     `json:"totalUnits"`
	Models     map[string]centerModelPercentageByModel `json:"models"`
}

type centerModelPercentageByModel struct {
	Units      int     `json:"units"`
	Percentage float64 `json:"percentage"`
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

func newCenterVolumeResponse(volumes []sales.CenterVolume) centerVolumeResponse {
	response := centerVolumeResponse{
		Centers: make(map[string]centerVolumeDetail),
	}

	for _, volume := range volumes {
		center := string(volume.DistributionCenter)
		response.Centers[center] = centerVolumeDetail{
			Units:       volume.Units,
			TotalAmount: centsToAmount(volume.TotalCents),
		}
	}

	return response
}

func newCenterModelPercentageResponse(percentages []sales.CenterModelPercentage) centerModelPercentageResponse {
	response := centerModelPercentageResponse{
		Centers: make(map[string]centerModelPercentageByCenter),
	}

	for _, percentage := range percentages {
		center := string(percentage.DistributionCenter)
		model := string(percentage.Model)
		centerDetail := response.Centers[center]

		if centerDetail.Models == nil {
			centerDetail.Models = make(map[string]centerModelPercentageByModel)
		}

		centerDetail.TotalUnits += percentage.Units
		centerDetail.Models[model] = centerModelPercentageByModel{
			Units:      percentage.Units,
			Percentage: percentage.Percentage,
		}

		response.TotalUnits += percentage.Units
		response.Centers[center] = centerDetail
	}

	return response
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
