package respostas

import (
	"api/src/erros"
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {

	errorCode, erro := erros.HandleMySqlError(erro)

	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
		ErrorCode int `json:"errorCode"`
	}{
		Erro: erro.Error(),
		ErrorCode: errorCode,
	})
}
