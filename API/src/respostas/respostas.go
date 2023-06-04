package respostas

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
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

// Extrai o código do erro e retorna em uma string
func ExtraiCodigoErro(erro error) string {
	re, re_erro := regexp.Compile("[0-9]{4}")
	if re_erro != nil {
		return "Erro ao fazer regex"
	}

	string_error := re.FindString(erro.Error())

	return string_error
}

// Retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {

	status_erro := ExtraiCodigoErro(erro)

	JSON(w, statusCode, struct {
		Erro       string `json:"erro"`
		ErrorCode  string `json:"errorCode"`
		StatusCode int    `json:"statusCode"`
	}{
		Erro:       erro.Error(),
		ErrorCode:  status_erro,
		StatusCode: statusCode,
	})
}
