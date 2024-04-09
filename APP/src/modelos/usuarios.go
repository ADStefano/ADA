package modelos

import (
	"app/src/config"
	"app/src/requisicoes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Representa um usuário da aplicação
type Usuario struct {
	ID          uint          `json:"id"`
	Nome        string        `json:"nome"`
	Nick        string        `json:"nick"`
	Email       string        `json:"email"`
	CriadoEm    time.Time     `json:"criadoEm"`
	Seguidores  []Usuario     `json:"seguidores"`
	Seguindo    []Usuario     `json:"seguindo"`
	Publicacoes []Publicacoes `json:"publicacoes"`
}

// Retorna o struct de usuário completamente preenchido
func BuscarDadosUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {

	log.Printf("Buscando dados do perfil do usuário: %d", usuarioID)

	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacoes)

	go BuscarDadosBaseUsuario(canalUsuario, usuarioID, r)
	go BuscarDadosSeguidores(canalSeguidores, usuarioID, r)
	go BuscarDadosSeguindo(canalSeguindo, usuarioID, r)
	go BuscarDadosPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacoes
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar usuário")
			}
			usuario = usuarioCarregado
		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar seguidores")
			}
			seguidores = seguidoresCarregados
		case seguindoCarregado := <-canalSeguindo:
			if seguindoCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar quem o usuário segue")
			}
			seguindo = seguindoCarregado
		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar quem o usuário segue")
			}
			publicacoes = publicacoesCarregadas
		}

	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil

}

// Retorna os dados base do usuário
func BuscarDadosBaseUsuario(canalUsuario chan<- Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canalUsuario <- Usuario{}
		return
	}

	defer response.Body.Close()

	var usuario Usuario
	if erro := json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canalUsuario <- Usuario{}
		return
	}

	canalUsuario <- usuario

}

// Retorna os seguidores do usuário
func BuscarDadosSeguidores(canalSeguidores chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canalSeguidores <- nil
		return
	}

	defer response.Body.Close()

	var seguidores []Usuario
	if erro := json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canalSeguidores <- nil
		return
	}

	if seguidores == nil{
		canalSeguidores <- make([]Usuario, 0)
		return
	}

	canalSeguidores <- seguidores

}

// Retorna quem o usuário segue
func BuscarDadosSeguindo(canalSeguindo chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canalSeguindo <- nil
		return
	}

	defer response.Body.Close()

	var seguindo []Usuario
	if erro := json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canalSeguindo <- nil
		return
	}

	if seguindo == nil{
		canalSeguindo <- make([]Usuario, 0)
		return
	}

	canalSeguindo <- seguindo

}

// Retorna as publicações do usuário
func BuscarDadosPublicacoes(canalPublicacoes chan<- []Publicacoes, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.API_URL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canalPublicacoes <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacoes
	if erro := json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canalPublicacoes <- nil
		return
	}

	if publicacoes == nil{
		canalPublicacoes <- make([]Publicacoes, 0)
		return
	}

	canalPublicacoes <- publicacoes

}
