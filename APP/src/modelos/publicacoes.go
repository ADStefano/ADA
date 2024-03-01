package modelos

import "time"

// Representa uma publicação feita por um usuário
type Publicacoes struct {
	ID           uint64    `json:"id,omitempty"`
	Titulo       string    `json:"titulo,omitempty"`
	Conteudo     string    `json:"conteudo,omitempty"`
	AutorID      uint64    `json:"autorId,omitempty"`
	AutorNick    string    `json:"autorNick,omitempty"`
	Curtidas     uint32    `json:"curtidas"`
	CriadaEm     time.Time `json:"criadaEm"`
	AtualizadaEm time.Time `json:"atualizadaEm"`
}
