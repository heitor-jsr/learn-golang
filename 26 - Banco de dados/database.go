package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", urlConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// if init abaixo. vai simplesmente jogar o valor do ping dentro de err, e se ele for diferente
	// de nil, o go da o panic no err. se for igual a nil, significa que a conexao foi estabelecida
	// e nao existe erro algum dentro da variavel.

	// ou seja, vai verificar se a variavel inicializada respeita a condição e, se sim, segue o codigo
	// ou da o panic.

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Conectado com sucesso!")
}
