package respostas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Representa um erro da API
type ErroAPI struct{
	Erro string `json:"erro"`
}

// Retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil{
		log.Fatal(erro)
	}
}

// Trata a requisições com status code 400 ou superior
func TratarErroAPI(w http.ResponseWriter, r* http.Response){
	var erro ErroAPI

	json.NewDecoder(r.Body).Decode(&erro)
	fmt.Print(erro)
	JSON(w, r.StatusCode, erro)
}