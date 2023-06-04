package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error){

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["UsuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	
	return token.SignedString([]byte(config.Secret_Key))
}

// Verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil{
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return nil
	}

	return errors.New("token inválido")
}

// Extrai o token do request
func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// Extrai o ID do usuário do token
func ExtrairUsuarioID (r *http.Request) (uint64, error){

	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil{
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f",permissoes["UsuarioId"]), 10, 64)
		if erro != nil{
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, nil

}

// Valida a chave de verificação do token
func retornarChaveVerificacao(token *jwt.Token)(interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.Secret_Key, nil
}