package main

import (
	"net/http"
	"os"

	"coto-back/internal/http/router"
	"coto-back/internal/logging"
	"coto-back/internal/sales"
	"coto-back/internal/sales/fake"

	"github.com/joho/godotenv"
)

func main() {

	logger := logging.New()

	// Load .env for local development.
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	store := sales.NewStore()
	service := sales.NewService(store, fake.GenerateSales)
	service.Seed()

	r := router.New(service, logger)

	logger.Printf("App running at http://localhost%s", addr)
	logger.Printf("%v", http.ListenAndServe(addr, r))
}
