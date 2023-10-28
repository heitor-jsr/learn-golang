package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuarios struct {
	ID uint32 `json:"id"`
	Name string `json:"nome"`
	Email string `json:"email"`
}

// CreateUser cria um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// a primeira coisa a ser feita nesta função é ler o corpo da requisição que o cliente está enviando. para isso, é usado o pacote ioutil.ReadAll do Go.

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// escreve uma mensagem de resposta para o cliente
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuarios

	// converte o json recebido para um struct

	if err = json.Unmarshal(reqBody, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o json para struct"))
		return
	}
	// como não escrevemos nenhum retorno para o usuario com o reponseWriter, o go só vai dar o retorno pra gente no terminal com o fmt.
	fmt.Println(usuario)

	// até aqui nós só estamos lendo a requisição e jogando ela dentro do struct de usuario que é criado acima. o &usuario vai pegar o valor do json.Unmarshal e jogar dentro do locan na memoria em que tá armazenado o struct de usuario.

	// não precisamos do ping aqui, porque ele já foi dado no pacote do database.
	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	// para inserir um dado no banco de dados em go, é muito comum usar o padrão PREPARE STATEMENT, que busca evitar ataques como o sql injection. o sql injetction funciona assim: a pessoa manipula os dados a serem inseridos na sua tabela, e depois realiza um segundo comando sql para obter algo que ele queria, ou dropar a tabela.

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?,?)")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	// com isso, o statement está concluido e só precisamos executar ele com os valores corretos. para isso, usamos o statement.Exec(). é dentro dele que passamos os valores para substituir os valores dos ? ?.

	insert, err := statement.Exec(usuario.Name, usuario.Email)

	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	// passando desse ponto, o usuario foi inserido no banco de dados.

	// agora precisamos obter o id do usuario inserido. para isso, só usar o LastInsertId
	idInserted, err := insert.LastInsertId()

	if err != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}
	// se não ocorrer erro, retornamos a mensagem de sucesso para o usuário.
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprint("Usuário inserido com sucesso, Id: ", idInserted)))
}
