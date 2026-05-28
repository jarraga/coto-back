package saleshandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"coto-back/internal/sales"
)

func TestCreateSale(t *testing.T) {
	handler := newTestHandler()
	body := strings.NewReader(`{"model":"Sport","distributionCenter":"north","units":2}`)
	request := httptest.NewRequest(http.MethodPost, "/sales", body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}
}

func TestCreateSaleWithInvalidModel(t *testing.T) {
	handler := newTestHandler()
	body := strings.NewReader(`{"model":"Invalid","distributionCenter":"north","units":2}`)
	request := httptest.NewRequest(http.MethodPost, "/sales", body)
	response := httptest.NewRecorder()

	handler.Create(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}

func TestTotalVolume(t *testing.T) {
	handler := newTestHandlerWithSale()
	request := httptest.NewRequest(http.MethodGet, "/sales/volume", nil)
	response := httptest.NewRecorder()

	handler.TotalVolume(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}

func TestVolumeByCenter(t *testing.T) {
	handler := newTestHandlerWithSale()
	request := httptest.NewRequest(http.MethodGet, "/sales/volume/by-center", nil)
	response := httptest.NewRecorder()

	handler.VolumeByCenter(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}

func TestModelPercentagesByCenter(t *testing.T) {
	handler := newTestHandlerWithSale()
	request := httptest.NewRequest(http.MethodGet, "/sales/model-percentages/by-center", nil)
	response := httptest.NewRecorder()

	handler.ModelPercentagesByCenter(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}
}

func newTestHandler() Handler {
	store := sales.NewStore()
	service := sales.NewService(store, nil)

	return NewHandler(service)
}

func newTestHandlerWithSale() Handler {
	handler := newTestHandler()

	handler.service.Create(sales.Sale{
		Model:              sales.ModelSport,
		DistributionCenter: sales.CenterNorth,
		Units:              2,
		UnitPriceCents:     1947400,
		TotalCents:         3894800,
	})

	return handler
}
