package erros

import (
	"errors"
	"strings"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

func HandleMySqlError(erro error) (int, error) {

	me, ok := erro.(*mysql.MySQLError)
	if !ok {
		erro = errors.New("erro ao obter tipo de erro do banco")
	}

	var errorCode int

	switch me.Number {
	default:
		return 0, erro
	case mysqlerr.ER_DUP_ENTRY:
		if strings.Contains(erro.Error(), "nick") {
			erro = errors.New("nickname já cadastrado")
			errorCode = mysqlerr.ER_DUP_ENTRY
		} else if strings.Contains(erro.Error(), "email") {
			erro = errors.New("email já cadastrado")
			errorCode = mysqlerr.ER_DUP_ENTRY
		}
	}

	return errorCode, erro
}
