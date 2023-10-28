package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"crud/server"
)

func main() {
	// todas as rotas vão ser compiladas dentro do router abaixo.

	//como todas as rotas vao se conectar ao db, é interessante criar um pacote para lidar com isso.
	router := mux.NewRouter()

	// método do pacote mux que vai lidar com uma função de requisição ao banco de dados. essa função é a getUsers, e deve ter a mesma assinatura que uma função de requisição que usa o padrão http usa. no mux também podemos especificar os metodos dentro dessa funcção.
	router.HandleFunc(("/users"), server.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
