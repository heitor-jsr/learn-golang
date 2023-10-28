package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	// user:password@/dbname?charset=utf8
	urlConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConn)

	//a função de conexao te retorna sempre o ponteiro de sql e um erro.
	// MAS, esses retornos são autoexclusivos. ou seja, nunca vai retornar os dois juntos.
	// por isso, no caso de erro não podemos simplesmente retornar um log.Fatal do erro.
	// precisamos retornar o nil, que representa o valor 0 do ponteiro de sql, e o err que a função nos passa de retorno
	if err != nil {
		return nil, err
	}

	// apenas nas funções que vão trabalhar com o http que vamos usar o defer para fechar a conexão.

	if err = db.Ping(); err != nil {
		return nil, err
	}
	// finalmente, como eu preciso sempre retornar os dois resultados que a função de conexao nos retorna, eu retorno o db, que representa a conexão com o banco, e o nil, que representa o valor 0 do erro que a func nos retorna.
	return db, nil
}
