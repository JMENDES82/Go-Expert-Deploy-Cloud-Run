package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JMENDES82/Go-Expert-Deploy-Cloud-Run/internal/handler"
)

func TestGetWeatherInvalidCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather?cep=123", nil)
	rr := httptest.NewRecorder()
	h := handler.NewWeatherHandler()
	http.HandlerFunc(h.GetWeather).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", status)
	}

	body, _ := ioutil.ReadAll(rr.Body)
	if string(body) != "invalid zipcode" {
		t.Errorf("Expected 'invalid zipcode', got '%s'", string(body))
	}
}
