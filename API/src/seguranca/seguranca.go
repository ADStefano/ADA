package seguranca

import "golang.org/x/crypto/bcrypt"

// Recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Compara uma senha com um hash e retorna se ela são iguais
func VerificarSenha(senhaHash string, senha string, ) error{
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}