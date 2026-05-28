package ops

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"coto-back/internal/sales"
)

func TestClearStore(t *testing.T) {
	store := sales.NewStore()
	service := sales.NewService(store, nil)
	handler := NewHandler(service)
	request := httptest.NewRequest(http.MethodDelete, "/ops/store", nil)
	response := httptest.NewRecorder()

	store.Save(sales.Sale{
		Model:              sales.ModelSedan,
		DistributionCenter: sales.CenterNorth,
		Units:              1,
		UnitPriceCents:     800000,
		TotalCents:         800000,
	})

	handler.ClearStore(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}

func TestSeedStore(t *testing.T) {
	store := sales.NewStore()
	service := sales.NewService(store, func() []sales.Sale {
		return []sales.Sale{
			{
				Model:              sales.ModelSUV,
				DistributionCenter: sales.CenterSouth,
				Units:              1,
				UnitPriceCents:     950000,
				TotalCents:         950000,
			},
		}
	})
	handler := NewHandler(service)
	request := httptest.NewRequest(http.MethodPost, "/ops/store/seed", nil)
	response := httptest.NewRecorder()

	handler.SeedStore(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}
