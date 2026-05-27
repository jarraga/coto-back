package main

import (
	"log"
	"net/http"
	"os"

	"coto-back/internal/http/router"
	"coto-back/internal/sales"
	"coto-back/internal/sales/fake"

	"github.com/joho/godotenv"
)

func main() {

	// lectura de .env para desarrollo
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	store := sales.NewStore()
	service := sales.NewService(store, fake.GenerateSales)
	service.Seed()

	r := router.New(service)

	log.Printf("App corriendo en http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
