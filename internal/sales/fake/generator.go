package fake

import (
	"math/rand"
	"time"

	"coto-back/internal/sales"
)

const (
	minSeedSalesCount = 1000
	maxSeedSalesCount = 3000
	minUnitsPerSale   = 1
	maxUnitsPerSale   = 5
)

func GenerateSales() []sales.Sale {
	salesCount := randomBetween(minSeedSalesCount, maxSeedSalesCount)
	generatedSales := make([]sales.Sale, 0, salesCount)

	for i := 0; i < salesCount; i++ {

		model := RandomModel()
		center := RandomDistributionCenter()
		units := randomBetween(minUnitsPerSale, maxUnitsPerSale)
		unitPriceCents, _ := sales.UnitPriceCents(model)

		sale := sales.Sale{
			Model:              model,
			DistributionCenter: center,
			Units:              units,
			UnitPriceCents:     unitPriceCents,
			TotalCents:         unitPriceCents * units,
			CreatedAt:          time.Now().UTC(),
		}

		generatedSales = append(generatedSales, sale)
	}

	return generatedSales
}

func RandomModel() sales.CarModel {
	value := rand.Intn(100)

	switch {
	case value < 40:
		return sales.ModelSedan
	case value < 70:
		return sales.ModelSUV
	case value < 90:
		return sales.ModelOffroad
	default:
		return sales.ModelSport
	}
}

func RandomDistributionCenter() sales.DistributionCenter {
	value := rand.Intn(100)

	switch {
	case value < 45:
		return sales.CenterNorth
	case value < 70:
		return sales.CenterSouth
	case value < 90:
		return sales.CenterEast
	default:
		return sales.CenterWest
	}
}

func randomBetween(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
