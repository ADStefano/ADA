package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/requisicoes"
	"app/src/respostas"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Renderiza a tela de login
func CarregarLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// Renderiza a tela de cadastro
func CarregarCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// Renderiza a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publicacoes", config.API_URL)

	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarErroAPI(w, response)
		return
	}

	var publicacoes []modelos.Publicacoes

	if erro := json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct{
		Publicacoes []modelos.Publicacoes
		UsuarioID uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID: usuarioID,
	})
}
