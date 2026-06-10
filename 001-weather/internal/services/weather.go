package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vasu1705/go-system-design/001-weather/internal/models"
	"net/http"
	"time"
)

type WeatherService struct {
	Client *http.Client
	APIKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{APIKey: apiKey, Client: &http.Client{}}
}

func (s *WeatherService) GetWeather(lat, lon float64, units string) (*models.WeatherResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	startTime := time.Now()
	defer cancel()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=%s&appid=%s", lat, lon, units, s.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	duration := time.Since(startTime)
	fmt.Printf("Request duration: %v\n", duration)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: %s", resp.Status)
	}
	var weatherResp models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}
	return &weatherResp, nil
}
