# Go Expert Deploy Cloud Run

Este projeto implementa um serviço escrito em Go que, a partir de um CEP válido, obtém a cidade correspondente e então consulta a temperatura atual (em Celsius, Fahrenheit e Kelvin) através da WeatherAPI. O serviço foi projetado para ser implantado no Google Cloud Run.

## Funcionalidades

- Recebe um CEP (8 dígitos) via query string (por exemplo: `?cep=29102571`).
- Valida o CEP (verifica se tem 8 dígitos).
- Consulta a [API ViaCEP](https://viacep.com.br/) para obter a localidade (cidade).
- Com a cidade obtida, faz uma consulta à [WeatherAPI](https://www.weatherapi.com/) para obter a temperatura atual em Celsius.
- Converte a temperatura de Celsius para Fahrenheit e Kelvin.
- Retorna as temperaturas no formato JSON.

## Endpoints

- `GET /weather?cep={cep}`

### Respostas esperadas

**Em caso de sucesso (CEP encontrado e temperatura obtida):**  
- Código HTTP: `200 OK`  
- Corpo da resposta (exemplo):
  ```json
  {
    "temp_C": 30.3,
    "temp_F": 86.54,
    "temp_K": 303.3
  }
**Em caso de CEP inválido (não possui 8 dígitos):**  
- Código HTTP: `422 Unprocessable Entity`  
- Mensagem: `invalid zipcode`

**Em caso de CEP não encontrado na ViaCEP ou cidade não reconhecida pela WeatherAPI:**  
- Código HTTP: `404 Not Found`  
- Mensagem: `can not find zipcode`

## Como Executar Localmente

Para rodar o projeto localmente usando `docker-compose`, siga os passos:

1. Garanta que você tenha o Docker e o docker-compose instalados.
2. Ajuste a variável de ambiente `WEATHER_API_KEY` no arquivo `docker-compose.yml` com a sua chave da WeatherAPI.
3. Execute:
   ```bash
   docker-compose up --build


## A aplicação esta implantada no Google Cloud Run 

- ULR:  https://go-expert-deploy-cloud-run-431167321935.us-central1.run.app

## Teste realizando uma requeste diretamente no link de implantação:

-  https://go-expert-deploy-cloud-run-431167321935.us-central1.run.app/weather?cep=29102571