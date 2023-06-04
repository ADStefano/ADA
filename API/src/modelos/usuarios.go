package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Representa um usuário
type Usuarios struct {
	
	ID 			uint64 		`json:"id,omitempty"`
	Nome 		string 		`json:"nome,omitempty"`
	Nick 		string 		`json:"nick,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	Senha 		string 		`json:"senha,omitempty"`
	CriadoEm 	time.Time 	`json:"criadoem,omitempty"`
}

// Método para validar e formatar os dados do usuário
func(usuario *Usuarios) Preparar(etapa string) error{

	if erro:= usuario.validar(etapa); erro != nil{
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	return nil
}

func(usuario *Usuarios) validar(etapa string) error{
	if usuario.Nome == ""{
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if usuario.Email == ""{
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil{
		 	return errors.New("o e-mail inserido é inválido")
	} 
	if usuario.Nick == ""{
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == ""{
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func(usuario *Usuarios) formatar(etapa string) error{
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro"{
		senhaHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil{
			return erro
		}

		usuario.Senha = string(senhaHash)
	}

	return nil

}