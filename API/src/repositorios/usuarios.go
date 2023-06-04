package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Representa um repositorio de usuários, parecido com um construtor
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuários
func RepositorioUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Cria o usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuarios) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Atualiza os dados de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuarios) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Busca por todos os usuários que contenham o mesmo nome ou nick
func (repositorio Usuarios) Buscar(nome_nick string) ([]modelos.Usuarios, error) {

	nome_nick = fmt.Sprintf("%%%s%%", nome_nick) // %nome_nick%

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR  nick LIKE ?", nome_nick, nome_nick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil
}

// Busca pelo usuário com ID correspondente
func (repositorio Usuarios) BuscarID(ID uint64) (modelos.Usuarios, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios WHERE id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Usuarios{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuarios

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Busca um usuário no banco de dados pelo email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarEmail(email string) (modelos.Usuarios, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?", email,
	)
	if erro != nil {
		return modelos.Usuarios{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuarios

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}

	return usuario, nil
}

// Faz a relação do usuário com seguidor na tabela seguidores
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"INSERT IGNORE INTO seguidores(usuario_id, seguidor_id) VALUES (?, ?)",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Apaga a relação do usuário com o seguidor na tabela seguidores
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Retorna todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuarios, error) {

	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoem
		FROM usuarios u
		JOIN seguidores s ON u.id = s.seguidor_id
		WHERE s.usuario_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil
}

// Traz todos os usuários que um determinado usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuarios, error) {

	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoem
		FROM usuarios u
		JOIN seguidores s ON u.id = s.usuario_id
		WHERE s.seguidor_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil
}

// Traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario modelos.Usuarios

	if linha.Next() {
		if erro := linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// Altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {

	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET senha = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil

}
