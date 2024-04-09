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
	"strings"

	"github.com/gorilla/mux"
)

// Renderiza a tela de login
func CarregarLogin(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacoes
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// Renderiza a página de edição de publicação
func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.API_URL, publicacaoID)
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

	var publicacao modelos.Publicacoes
	if erro := json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "editar-publicacao.html", publicacao)
}

// Renderiza a página com os usuários que atendem o parâmetro da busca
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {

	usuario := strings.Replace(strings.ToLower(r.URL.Query().Get("usuario")), " ", "+", -1)
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.API_URL, usuario)
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

	var usuarios []modelos.Usuario
	if erro := json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	usuario, erro := modelos.BuscarDadosUsuarioCompleto(usuarioID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"],10,64)


	utils.ExecutarTemplate(w, "perfil-usuario.html", struct{
		Usuario modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario: usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})
}
