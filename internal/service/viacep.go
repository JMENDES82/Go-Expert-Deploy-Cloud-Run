package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ViaCEPResponse struct {
	Cep       string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func GetCityFromCEP(cep string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := client.Get(url)
	
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 || resp.StatusCode == 404 {
		return "", errors.New("cep not found")
	}

	var r ViaCEPResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	if r.Localidade == "" {
		return "", errors.New("cep not found")
	}

	return r.Localidade, nil
}
