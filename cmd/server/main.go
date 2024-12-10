package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/JMENDES82/Go-Expert-Deploy-Cloud-Run/internal/handler"
)

func main() {
	_ = godotenv.Load() // Carrega variáveis de ambiente do .env, se existir.

	router := mux.NewRouter()

	weatherHandler := handler.NewWeatherHandler()
	router.HandleFunc("/weather", weatherHandler.GetWeather).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
