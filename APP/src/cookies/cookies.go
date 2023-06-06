package cookies

import (
	"app/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Utiliza as variáveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string] string{
		"id": ID,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil{
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name: "dados",
		Value: dadosCodificados,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}
