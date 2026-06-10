package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vasu1705/go-system-design/001-weather/internal/services"
)

type WeatherController struct {
	*services.WeatherService
}

func NewWeatherController(ws *services.WeatherService) *WeatherController {
	return &WeatherController{WeatherService: ws}
}

func (ctrl *WeatherController) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ctrl.writeJSONError(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	queryParam := r.URL.Query()

	lat, lon, units := queryParam.Get("lat"), queryParam.Get("lon"), queryParam.Get("units")
	if lat == "" || lon == "" {
		ctrl.writeJSONError(w, http.StatusBadRequest, "Missing required query parameters: lat, lon")
		return
	}
	if units == "" {
		units = "metric"
	}
	latFloat, errLat := strconv.ParseFloat(lat, 64)
	lonFloat, errLon := strconv.ParseFloat(lon, 64)
	if errLat != nil || errLon != nil {
		ctrl.writeJSONError(w, http.StatusBadRequest, "Invalid latitude or longitude format")
		return
	}

	resp, err := ctrl.WeatherService.GetWeather(latFloat, lonFloat, units)
	if err != nil {
		ctrl.writeJSONError(w, http.StatusInternalServerError, "Failed to get weather data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (ctrl *WeatherController) writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
