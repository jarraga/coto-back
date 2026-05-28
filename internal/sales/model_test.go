package sales

import "testing"

func TestIsValidDistributionCenter(t *testing.T) {

	isValid := IsValidDistributionCenter(CenterNorth)

	if !isValid {
		t.Fatal("expected north distribution center to be valid")
	}
}

func TestUnitPriceCents(t *testing.T) {

	price, ok := UnitPriceCents(ModelSport)

	if !ok {
		t.Fatal("expected sport model to have a price")
	}

	if price != 1947400 {
		t.Fatalf("expected sport price 1947400, got %d", price)
	}
}
