package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/JMENDES82/Go-Expert-Deploy-Cloud-Run/internal/model"
	"github.com/JMENDES82/Go-Expert-Deploy-Cloud-Run/internal/service"
	"github.com/JMENDES82/Go-Expert-Deploy-Cloud-Run/internal/util"
)

type WeatherHandler struct{}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" || !isValidCEP(cep) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	city, err := service.GetCityFromCEP(cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find zipcode"))
		return
	}

	tempC, err := service.GetCurrentTemperature(city)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not find city"))	
		return
	}

	resp := model.WeatherResponse{
		TempC: tempC,
		TempF: util.CelsiusToFahrenheit(tempC),
		TempK: util.CelsiusToKelvin(tempC),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func isValidCEP(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}
