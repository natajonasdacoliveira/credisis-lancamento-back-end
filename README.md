# credisis-lancamento-back-end
Atividade 2 

# Instalando

Instale o golang https://golang.org/

git clone https://github.com/natajonasdacoliveira/credisis-unidades-federativas

cd /credisis-unidades-federativas

# Troubleshooting
Caso necessário, instalar os pacotes necessários com o comando go get.

Exemplo : go get github.com/labstack/echo/v4

# Banco de dados
Acesse o arquivo db/db.go e modifique de acordo com o banco utilizado.

"root:root@tcp(127.0.0.1:3306)/conta_corrente"

Em ordem: nome, senha, IP, nome do banco.

# API
Podem ser encontradas no arquivo server.go, na função main.

# SQL

drop database if exists conta_corrente;
create database conta_corrente;
use conta_corrente;

create table cooperado(
	idCooperado int auto_increment primary key,
    conta varchar(200) not null unique,
    saldo double not null,
    nome varchar(200) not null
);

create table lancamento(
	idLancamento int auto_increment primary key,
    debito double not null default 0,
    credito double not null default 0,
    tipoLancamento varchar(200) not null,
    descricao varchar(200) not null,
    createdAt dateTime not null,
    sistema varchar(200) not null,
    
	idCooperadoOrigemFK int not null,
    foreign key(idCooperadoOrigemFK) references cooperado(idCooperado) on delete no action,
    
    idCooperadoDestinoFK int not null,
	foreign key(idCooperadoDestinoFK) references cooperado(idCooperado) on delete no action
);

insert into cooperado values(null, '0000110-7', 800, 'Anderson');
insert into cooperado values(null, '0000120-5', 1000, 'Felipe');

# POSTMAN API JSON

{
	"info": {
		"_postman_id": "75c27803-5ca2-455f-99c3-1178a28f44f2",
		"name": "api-credisis-desafio-conta-corrente",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Lancamento",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"credito\": \"0\",\r\n    \"debito\": \"150\",\r\n    \"tipoLancamento\": \"Débito\",\r\n    \"descricao\": \"Transferência para Felipe\",\r\n    \"sistema\": \"Mobile\",\r\n    \"idCooperadoOrigemFK\": \"1\",\r\n    \"idCooperadoDestinoFK\": \"2\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/lancamento",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"lancamento"
					]
				}
			},
			"response": []
		},
		{
			"name": "localhost:1323/cooperado/saldo",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"idCooperado\": \"2\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:1323/cooperado/saldo",
					"host": [
						"localhost"
					],
					"port": "1323",
					"path": [
						"cooperado",
						"saldo"
					]
				}
			},
			"response": []
		}
	]
}









