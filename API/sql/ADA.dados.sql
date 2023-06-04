-- HASH SENHA 123456: $2a$10$srclkHQUUeXW4lflxs0vz.rHK9F34Bg60T7bmGYhMv7wcgiggf/eS

USE ADA;

INSERT INTO usuarios(nome, nick, email, senha) VALUES
("Teste 1", "Teste_1_Nick", "teste1@email.com", "$2a$10$srclkHQUUeXW4lflxs0vz.rHK9F34Bg60T7bmGYhMv7wcgiggf/eS"),
("Teste 2", "Teste_2_Nick", "teste2@email.com", "$2a$10$srclkHQUUeXW4lflxs0vz.rHK9F34Bg60T7bmGYhMv7wcgiggf/eS"),
("Teste 3", "Teste_3_Nick", "teste3@email.com", "$2a$10$srclkHQUUeXW4lflxs0vz.rHK9F34Bg60T7bmGYhMv7wcgiggf/eS");

INSERT INTO seguidores(usuario_id, seguidor_id) VALUES
(1,2),
(2,3),
(3,1);

INSERT INTO publicacoes(titulo, conteudo, autor_id, curtidas) VALUES
("Publicação do usuário 1","Essa publicação pertence ao usuário 1! Oba!", 1, 1),
("Publicação do usuário 2","Essa publicação pertence ao usuário 2! Oba!", 2, 2),
("Publicação do usuário 3","Essa publicação pertence ao usuário 3! Oba!", 3, 3);