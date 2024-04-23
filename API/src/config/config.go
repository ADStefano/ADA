package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	// Conexão com o banco de dados
	DB_String = ""

	// Porta onde a API vai estar rodando
	Porta = 0

	// Chave que vai ser usada para assinar o token
	Secret_Key []byte
)

func Config(){

	var erro error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro) 	
	}

	Porta, erro = strconv.Atoi(os.Getenv("PORT"))
	if erro != nil{
		Porta = 9000
	}

	DB_String = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE"),
	)

	Secret_Key = []byte(os.Getenv("SECRET_KEY"))

}