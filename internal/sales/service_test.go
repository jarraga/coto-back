package sales

import "testing"

func TestServiceCreate(t *testing.T) {

	store := NewStore()
	service := NewService(store, nil)

	service.Create(Sale{
		Model:              ModelSedan,
		DistributionCenter: CenterNorth,
		Units:              1,
		UnitPriceCents:     800000,
		TotalCents:         800000,
	})

	if len(store.sales) != 1 {
		t.Fatalf("expected 1 stored sale, got %d", len(store.sales))
	}
}

func TestServiceClear(t *testing.T) {

	store := NewStore()
	service := NewService(store, nil)

	store.Save(Sale{
		Model:              ModelSUV,
		DistributionCenter: CenterSouth,
		Units:              1,
		UnitPriceCents:     950000,
		TotalCents:         950000,
	})

	service.Clear()

	if len(store.sales) != 0 {
		t.Fatalf("expected empty store, got %d sales", len(store.sales))
	}
}

func TestServiceSeed(t *testing.T) {

	store := NewStore()
	service := NewService(store, func() []Sale {
		return []Sale{
			{
				Model:              ModelSport,
				DistributionCenter: CenterWest,
				Units:              1,
				UnitPriceCents:     1947400,
				TotalCents:         1947400,
			},
		}
	})

	service.Seed()

	if len(store.sales) != 1 {
		t.Fatalf("expected 1 seeded sale, got %d", len(store.sales))
	}
}

func TestServiceTotalVolume(t *testing.T) {

	store := NewStore()
	service := NewService(store, nil)

	store.Save(Sale{
		Model:              ModelSedan,
		DistributionCenter: CenterNorth,
		Units:              2,
		UnitPriceCents:     800000,
		TotalCents:         1600000,
	})

	volume := service.TotalVolume()

	if volume.TotalCents != 1600000 {
		t.Fatalf("expected total cents 1600000, got %d", volume.TotalCents)
	}
}

func TestServiceVolumeByCenter(t *testing.T) {

	store := NewStore()
	service := NewService(store, nil)

	store.Save(Sale{
		Model:              ModelOffroad,
		DistributionCenter: CenterEast,
		Units:              3,
		UnitPriceCents:     1250000,
		TotalCents:         3750000,
	})

	volumes := service.VolumeByCenter()

	if volumes[2].TotalCents != 3750000 {
		t.Fatalf("expected east total cents 3750000, got %d", volumes[2].TotalCents)
	}
}

func TestServiceModelPercentagesByCenter(t *testing.T) {

	store := NewStore()
	service := NewService(store, nil)

	store.Save(Sale{
		Model:              ModelSUV,
		DistributionCenter: CenterSouth,
		Units:              1,
		UnitPriceCents:     950000,
		TotalCents:         950000,
	})

	percentages := service.ModelPercentagesByCenter()

	if percentages[5].Percentage != 100 {
		t.Fatalf("expected south SUV percentage 100, got %.2f", percentages[5].Percentage)
	}
}
