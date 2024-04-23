package modelos

import (
	"errors"
	"log"
	"strings"
	"time"
)

type Publicacao struct {
	ID           uint64    `json:"id,omitempty"`
	Titulo       string    `json:"titulo,omitempty"`
	Conteudo     string    `json:"conteudo,omitempty"`
	AutorID      uint64    `json:"autorId,omitempty"`
	AutorNick    string    `json:"autornick,omitempty"`
	Curtidas     uint64    `json:"curtidas"`
	CriadaEm     time.Time `json:"criadaEm,omitempty"`
	AtualizadaEm time.Time `json:"atualizadaEm,omitempty"`
}

// Vai validar e formatar os campos necessários das publicações
func (Publicacao *Publicacao) Preparar() error {
	if erro := Publicacao.validar(); erro != nil {
		return erro
	}

	Publicacao.formatar()
	return nil
}

func (Publicacao *Publicacao) validar() error {

	log.Printf("Titulo da publicação: %s\n", Publicacao.Titulo)
	log.Printf("Conteudo da publicação: %s", Publicacao.Conteudo)

	if Publicacao.Titulo == "" {
		return errors.New("o título é obrigatório e não pode estar vazio")
	}

	if Publicacao.Conteudo == "" {
		return errors.New("o conteudo é obrigatório e não pode estar vazio")
	}

	return nil
}

func (Publicacao *Publicacao) formatar() {
	Publicacao.Titulo = strings.TrimSpace(Publicacao.Titulo)
	Publicacao.Conteudo = strings.TrimSpace(Publicacao.Conteudo)
}
