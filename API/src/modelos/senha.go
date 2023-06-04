package modelos

// Representa o corpo da requisição para atualizar a senha
type Senha struct{
	
	Nova 	string `json:"nova"`
	Atual 	string `json:"atual"`
}