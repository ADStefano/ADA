package modelos

// Contém o id e o token do usuário autenticado
type Token struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
