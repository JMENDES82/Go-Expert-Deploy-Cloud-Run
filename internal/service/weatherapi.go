package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type WeatherAPIResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetCurrentTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return 0.0, errors.New("missing WEATHER_API_KEY env var")
	}
	
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	query := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, query)
	resp, err := client.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 404 {
		return 0.0, errors.New("city not found")
	}

	var w WeatherAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return 0.0, err
	}

	if w.Location.Name == "" {
		return 0.0, errors.New("city not found")
	}

	return w.Current.TempC, nil
}
