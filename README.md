# pos-golang-apis

Repositório para os exercícios de criação das apis em golang da pós.

Nos itens a seguir, serão inseridas anotações e comentários anotados durante as aulas.


## Anotações Gerais

Para carregar as configurações, vamos utilizar o pacote viper. Ele é muito famoso na comunidade.

Geração de UUID será com a biblioteca uuid do google.

## JWT

Utilizamos o módulo de JWT que faz parte do midleware do chi. O token é dividido em três partes, que são separados por ".":

- A primeira parte define qual algoritmo está sendo utilizado. É o "header" do token.

- A segunda parte são os dados que você está passando (payload).

- A terceira e última parte é a assinatura do token. Ela pode ser baseada em chave ou certificado.

## Bancos de dados

Para o desenvolvimento utilizaremos o gorm e o SQLite.

Exemplos de query no SQLite:

```bash

wander@bsnote283:~/repos/CURSO-GO/07apis/pos-golang-apis$ sqlite3 cmd/server/test.db
SQLite version 3.37.2 2022-01-06 13:25:41
Enter ".help" for usage hints.
sqlite> 
sqlite> .databases
main: /home/wander/repos/CURSO-GO/07apis/pos-golang-apis/cmd/server/test.db r/w
sqlite> 
sqlite> .tables
products  users   
sqlite> 
sqlite> select * from users;
314f49a2-aa40-4dfa-88be-b4a9b6ccf382|John Doe|j@j.com|$2a$10$4bTq35A0hxnsKnOjtNmTyOE1md2TCVxRjGa3dnGv0Uh8aaY7VGVX2
sqlite> 
sqlite> .exit
wander@bsnote283:~/repos/CURSO-GO/07apis/pos-golang-apis$ 


```

## Entitys

As entidades ajudam a manter a consistência da aplicação.
As regras de negócios ficam nas entidades.
As senhas serão criptografadas utilizando o pacote golang.org/x/crypto/bcrypt

## Testes

A importância de fazer esses testes é que estamos implementando uma api e, caso dê um erro lá no final, não saberemos se o erro é na implementação da API ou na entidade. Dessa forma, quanto mais pudermos testar agora, reduzimos os riscos lá na frente.

Para executar todos os testes, basta ir na raiz do projeto e executar o seguinte comando:

```bash

wander@bsnote283:~/repos/CURSO-GO/07apis/pos-golang-apis$ go test ./...
?   	github.com/wandermaia/pos-golang-apis/cmd/server	[no test files]
?   	github.com/wandermaia/pos-golang-apis/configs	[no test files]
?   	github.com/wandermaia/pos-golang-apis/internal/dto	[no test files]
?   	github.com/wandermaia/pos-golang-apis/internal/infra/webserver/handlers	[no test files]
?   	github.com/wandermaia/pos-golang-apis/pkg/entity	[no test files]
ok  	github.com/wandermaia/pos-golang-apis/internal/entity	0.232s
ok  	github.com/wandermaia/pos-golang-apis/internal/infra/database	0.131s
wander@bsnote283:~/repos/CURSO-GO/07apis/pos-golang-apis$ 

```
Onde os três pontos indicam que ele vai pesquisar recursivamente nas subpastas a procura de arquivos de teste.

Para os testes de URL, utilizamos uma extensão do vscode chamada "REST Client".


## MidleWares

Midlewares importantes

- r.Use(middleware.Logger) -  para log
- r.Use(middleware.Recoverer) - recuperar o servidor
- r.Use(middleware.RealIP) - passa o IP real da requisição 

## Documentação

Para documentação, utilizaremos o openApi, através da biblioteca do swaggo.

Essa biblioteca converte as Go para o Swagger Documentation 2.0.

É necessário baixar essa biblioteca através do comando

```bash

go install github.com/swaggo/swag/cmd/swag@latest

```
> **OBSERVAÇÃO:**
>
> Os binários instalados pelo Go ficam na pasta bin do seu GOPATH. Por isso, é necessário adicionar no path para que ele funcione de qualquer local
> 
> O export pode ser criado no arquivo .bashrc conforme abaixo:
>
>```bash
> export GOPATH=$HOME/go
> export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
>
>```

Ao executar o comando `swag init -g cmd/server/main.go` na raiz do projeto, será gerada uma pasta chamada *docs*, contendo três arquivos: docs.go, que contém o código para abrir o swagger e swagger.json e swagger.yaml, onde temos a documentação do padrão api.

O pacote swag já traz alguns plugins para trabalhar com os principais frameworks, incluindo o go-chi.

## Referências


Viper

https://github.com/spf13/viper

Google uuid

https://github.com/google/uuid

Go bcrypt

https://pkg.go.dev/golang.org/x/crypto/bcrypt

Gorm

https://gorm.io/index.html

jwt io

https://jwt.io/

Golang chi

https://github.com/go-chi/chi

Swag

https://github.com/swaggo/swag

OPENAPI INITIATIVE

https://www.openapis.org/
