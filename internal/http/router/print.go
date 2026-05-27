package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PrintAvailableRoutes(routes chi.Routes) {

	fmt.Println()
	fmt.Println("Available endpoints")
	fmt.Println("-------------------")
	fmt.Printf("%-8s %s\n", "METHOD", "PATH")

	err := chi.Walk(routes, func(method string, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {

		fmt.Printf("%-8s %s\n", method, path)
		return nil
	})
	if err != nil {

		log.Printf("error printing available routes: %v", err)
	}

	fmt.Println()
}
