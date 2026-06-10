package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vasu1705/go-system-design/001-weather/internal/controllers"
	"github.com/vasu1705/go-system-design/001-weather/internal/services"
)

func main() {
	apiKey := "YOUR_OPENWEATHER_API_KEY"
	weatherService := services.NewWeatherService(apiKey)

	weatherController := controllers.NewWeatherController(weatherService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/weather", weatherController.GetWeatherHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,   // Max time to read request headers/body
		WriteTimeout: 10 * time.Second,  // Max time to write response payload
		IdleTimeout:  120 * time.Second, // Max time keep-alive connections stay idle
	}

	log.Println("Server is spinning up on http://localhost:8080...")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %v", err)
	}
}
