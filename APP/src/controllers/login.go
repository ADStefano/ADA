package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.API_URL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	var Token modelos.Token

	if erro = json.NewDecoder(response.Body).Decode(&Token); erro != nil{
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, Token.ID, Token.Token); erro != nil{
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w,http.StatusOK, nil)
}
