package requisicoes

import (
	"app/src/cookies"
	"fmt"
	"io"
	"net/http"
)

// Faz uma requisição passando o cookie no header como autenticação
func RequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil{
		return nil, erro
	}

	cookie, erro := cookies.Ler(r)
	if erro != nil{
		return nil, erro
	}

	bearerToken := fmt.Sprintf("Bearer %s", cookie["token"])

	request.Header.Add("Authorization", bearerToken)

	client := http.Client{}
	response, erro := client.Do(request)
	if erro != nil{
		return nil, erro
	}

	return response, nil
}
