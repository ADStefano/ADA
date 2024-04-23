package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/requisicoes"
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Chama a API para cadastrar um usuário no banco de dados
func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.API_URL)
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

	respostas.JSON(w, response.StatusCode, nil)
}

// Chama a API para seguir um usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	respostas.JSON(w,response.StatusCode, nil)

}

// Chama a API para parar de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-seguir", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	respostas.JSON(w,response.StatusCode, nil)

}

// Chama a API para editar o perfil do usuário
func EditarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	log.Printf("Atualizando perfil do usuário: %d", usuarioID)

	url := fmt.Sprintf("%s/usuarios/%d", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	respostas.JSON(w,response.StatusCode, nil)
}

// Chama a API para atualizar a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova": r.FormValue("nova"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}


	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	log.Printf("Atualizando senha do usuário: %d", usuarioID)

	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	log.Printf("Status code: %d", response.StatusCode)

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	respostas.JSON(w,response.StatusCode, nil)
}
