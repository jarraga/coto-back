package sales

import "testing"

func TestNewStore(t *testing.T) {
	store := NewStore()

	if store == nil {
		t.Fatal("expected store instance, got nil")
	}

	if store.nextID != 1 {
		t.Fatalf("expected next ID 1, got %d", store.nextID)
	}

	if store.sales == nil {
		t.Fatal("expected sales slice to be initialized")
	}
}

func TestStoreSave(t *testing.T) {
	store := NewStore()
	sale := Sale{
		Model:              ModelSedan,
		DistributionCenter: CenterNorth,
		Units:              1,
		UnitPriceCents:     800000,
		TotalCents:         800000,
	}

	store.Save(sale)

	if len(store.sales) != 1 {
		t.Fatalf("expected 1 stored sale, got %d", len(store.sales))
	}
}

func TestStoreFindAll(t *testing.T) {
	store := NewStore()

	store.Save(Sale{
		Model:              ModelSUV,
		DistributionCenter: CenterSouth,
		Units:              2,
		UnitPriceCents:     950000,
		TotalCents:         1900000,
	})

	sales := store.FindAll()
	if len(sales) != 1 {
		t.Fatalf("expected 1 sale, got %d", len(sales))
	}
}

func TestStoreClear(t *testing.T) {
	store := NewStore()

	store.Save(Sale{
		Model:              ModelSport,
		DistributionCenter: CenterWest,
		Units:              1,
		UnitPriceCents:     1947400,
		TotalCents:         1947400,
	})

	store.Clear()

	sales := store.FindAll()
	if len(sales) != 0 {
		t.Fatalf("expected empty store after clear, got %d sales", len(sales))
	}

	savedSale := store.Save(Sale{
		Model:              ModelOffroad,
		DistributionCenter: CenterEast,
		Units:              1,
		UnitPriceCents:     1250000,
		TotalCents:         1250000,
	})

	if savedSale.ID != 1 {
		t.Fatalf("expected ID to restart at 1 after clear, got %d", savedSale.ID)
	}
}
