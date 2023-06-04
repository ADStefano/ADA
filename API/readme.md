# Sumário:
* [Introdução] [#introdução]
* [sql] [#sql]
* [src] [#src]
* [Rotas da API] [#rotas-da-api]

# Introdução
  * API do backend do projeto, responsável por crias as rotas da API, autenticar o usuário, obter as variáveis de ambiente e fazer as operações de CRUD. Roda na porta 5000

# sql
  * ADA.dados.sql: Arquivo sql responsável por fazer a inserção de dados de teste no banco
  * ADA.session.sql: Arquivo sql responsável por criar o banco e as tabelas caso não exista ou recriar tudo caso já exista

# src
  ## autenticação: 
  * Cria, valida e extrai o token de autenticação das chamadas.
  ## banco: 
  * Responsável pela conexão com o banco de dados. 
  ## config:
  * Obtém as variáveis de ambiente.
  ## controllers:
  * Funções responsáveis por administrar e validar o CRUD.
  ## middlewares:
  * Faz o log e valida a autenticação do usuário.
  ## modelos:
  * Structs do projeto.
  ## repositorios:
  * Faz o CRUD.
  ## respostas:
  * Retorna as resposta da API em JSON.
  ## router:
  * Gerencia e configura as rotas da API.
  ## seguranca:
  * Administra as senhas e os hashs.

# Rotas da API:
  * ## LOGIN: 
        localhost:5000/login
  * ## PUBLICAÇÕES:
        localhost:5000/publicacoes 
        localhost:5000/publicacoes/{PublicacaoId}
        localhost:5000/publicacoes/{PublicacaoId}curtir
        localhost:5000/publicacoes/{PublicacaoId}descurtir
        localhost:5000/usuarios/{UsuarioId}/publicacoes
  * ## USUARIOS:
        localhost:5000/usuarios
        localhost:5000/usuarios/{UsuarioId}
        localhost:5000/usuarios/{UsuarioId}/seguir
        localhost:5000/usuarios/{UsuarioId}/parar-seguir
        localhost:5000/usuarios/{UsuarioId}/seguidores
        localhost:5000/usuarios/{UsuarioId}/seguindo
        localhost:5000/usuarios/{UsuarioId}/atualizar-senha



