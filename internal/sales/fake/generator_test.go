package fake

import (
	"testing"

	"coto-back/internal/sales"
)

func TestGenerateSales(t *testing.T) {
	generatedSales := GenerateSales()

	if len(generatedSales) < minSeedSalesCount || len(generatedSales) > maxSeedSalesCount {
		t.Fatalf("expected sales count between %d and %d, got %d", minSeedSalesCount, maxSeedSalesCount, len(generatedSales))
	}
}

func TestRandomModel(t *testing.T) {
	model := RandomModel()
	_, ok := sales.UnitPriceCents(model)

	if !ok {
		t.Fatalf("expected valid model, got %s", model)
	}
}

func TestRandomDistributionCenter(t *testing.T) {
	center := RandomDistributionCenter()

	if !sales.IsValidDistributionCenter(center) {
		t.Fatalf("expected valid distribution center, got %s", center)
	}
}

func TestRandomBetween(t *testing.T) {
	value := randomBetween(1, 5)

	if value < 1 || value > 5 {
		t.Fatalf("expected value between 1 and 5, got %d", value)
	}
}
