package main

import (
	"log"
	"net/http"

	"example.com/rest-api-recoll-mobile/config"
	"example.com/rest-api-recoll-mobile/router"

	"github.com/rs/cors"
)

func main() {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := router.SetupRouter()

	c := cors.New(cors.Options{
		// pengaturan allow origins
		// AllowedOrigins:     []string{"http://example.com"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "X-CSRF-Token"},

		//AllowCredentials: false,
		//Debug: true,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":8090", handler)
}
