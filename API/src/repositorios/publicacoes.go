package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func RepositorioPublicacoes(db *sql.DB) *Publicacoes{
	return &Publicacoes{db}
}

// Cria uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error){
	
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil{
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil{
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Traz uma única publicação do banco de dados	
func (repositorio Publicacoes) BuscarPorId(publicacaoID uint64) (modelos.Publicacao, error){
	
	linhas, erro := repositorio.db.Query(
		`SELECT p.*, u.nick FROM publicacoes p
		JOIN usuarios u ON p.autor_id = u.id
		WHERE p.id = ?`,
		publicacaoID,
	)
	if erro != nil{
		return modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next(){

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AtualizadaEm,
			&publicacao.AutorNick,
		); erro!= nil{
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Traz todas as publicações do usuário e dos seus seguidores
func (repositorio Publicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error){

	linhas, erro := repositorio.db.Query(
	   `SELECT DISTINCT p.*, u.nick FROM publicacoes p
		JOIN usuarios u ON p.autor_id = u.id
		JOIN seguidores s ON s.usuario_id = u.id
		WHERE u.id = ?
		OR s.seguidor_id = ?
		ORDER BY 1 DESC`,
		usuarioID, usuarioID,
	)
	if erro!= nil{
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next(){

		var publicacao modelos.Publicacao

		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AtualizadaEm,
			&publicacao.AutorNick,
		); erro!= nil{
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Traz todas as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPublicacoesUsuario(usuarioID uint64) ([]modelos.Publicacao, error){

	linhas, erro := repositorio.db.Query(
		`SELECT p.*, u.nick FROM publicacoes p
		 JOIN usuarios u ON p.autor_id = u.id
		 WHERE p.autor_id = ?`,
		 usuarioID,
	 )
	 if erro!= nil{
		 return nil, erro
	 }

	 defer linhas.Close()
 
	 var publicacoes []modelos.Publicacao
 
	 for linhas.Next(){
 
		 var publicacao modelos.Publicacao
 
		 if erro := linhas.Scan(
			 &publicacao.ID,
			 &publicacao.Titulo,
			 &publicacao.Conteudo,
			 &publicacao.AutorID,
			 &publicacao.Curtidas,
			 &publicacao.CriadaEm,
			 &publicacao.AtualizadaEm,
			 &publicacao.AutorNick,
		 ); erro!= nil{
			 return nil, erro
		 }
 
		 publicacoes = append(publicacoes, publicacao)
	 }
 
	 return publicacoes, nil
}

// Altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) (error){

	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ?, atualizadaEm = NOW() WHERE id = ?")
	if erro != nil{
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil{
		return erro
	}

	return nil
}

// Deleta um publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) (error){

	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if erro != nil{
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil
}

// Curte uma publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) (error){
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil{
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil
}

// Retira uma curtida de uma ºublicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) (error){
	statement, erro := repositorio.db.Prepare(`
	UPDATE publicacoes 
	SET curtidas = 
	CASE
	WHEN curtidas > 0 THEN curtidas - 1
	ELSE 0
	END
	WHERE id = ?`,
)
	if erro != nil{
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil
}