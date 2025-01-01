package main

import (
	"log"
	"net/http"

	"q4/backend/db"
	"q4/backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db.InitializeDatabase()

	// Set up router
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	// Add CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
