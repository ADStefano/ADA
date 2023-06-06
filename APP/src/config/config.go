package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Rota da API
	API_URL = ""

	// Porta da API
	Port = 0

	// Utilizado para autenticar o cookie
	HashKey []byte

	// Utilizado para criptografar os dados do cookie
	BlockKey []byte
)

// Inicializa as variáveis de ambiente
func Carregar() {
	var erro  error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil{
		log.Fatal(erro)
	}

	API_URL = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))

	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
