package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conecta com o banco de dados
func Conn()(*sql.DB, error){

	db, erro := sql.Open("mysql", config.DB_String)
	if erro != nil{
		fmt.Println("Erro ao conectar com o banco de dados!")
		return nil, erro
	}

	if erro := db.Ping(); erro != nil{
		fmt.Println("Erro com a conexão com o banco de dados!")
		return nil, erro
	}

	return db, nil
}