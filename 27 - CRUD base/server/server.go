package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

type usuario struct {
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

	var usuario usuario

	// converte o json recebido para um struct

	if err = json.Unmarshal(reqBody, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o json para struct"))
		return
	}
	// como não escrevemos nenhum retorno para o usuario com o reponseWriter, o go só vai dar o retorno pra gente no terminal com o fmt.
	// fmt.Println(usuario)

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

// GetUsers retorna todos usuários no banco de dados
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	// quando vamos apenas consultar um dado do db, não usamos prepare. nesses casos, usamos outro comando que é o query. isso vai pegar as linhas especificas da tabela para a gente.

	lines, err := db.Query("select * from usuarios")

	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}

	defer lines.Close()

// como retornamos todos os dados da nossa tabela, precisamos armazenar ele dentro de um slice de usuarios, para depois retornar esses dados para o cliente.

	var usuarios []usuario
	// ocorre que precisamos iterar sobre as linhas dos usuarios que nos sao retornadas pela query, para poder armazenar elas dentro da variável usuarios. para isso, usamos o cód a seguir.

	for lines.Next() {
		// isso vai ler os dados que estão na linha, jogar dentro da variavel usuario que é um struc de um só usuario, e depois jogar dentro de um slice de usuarios.
		var usuario usuario
		// o Scan vai ler cada linha da tabela e jogar dentro do struct, verificando se os dados que são retornados da query estão de acordo com a estrutura do struc.
		if err := lines.Scan(&usuario.ID, &usuario.Name, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao ler os dados"))
			return
		}

		// se tudo der certo, adicionamos o usuario dentro do slice de usuarios, com o método append.
		// isso vai popular o slice conforme os dados da query forem retornados corretamente.
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	// por fim, precisamos converter os dados dos usuarios retornados em JSON. o newEnconder vai codigicar os dados para JSON, permitindo que ele seja encodado e enviado para o cliente.
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao retornar os dados"))
		return
	}
}

// GetUser retorna um usuário específico do banco de dados
func GetUser(w http.ResponseWriter, r *http.Request) {
	// primeira coisa que precisamos ter em mente é que aqui nos precisamos ler o parametro de busca da rota (o id que é passado para ela vai representar o id que o cliente quer recuperar do db). para isso, basta usar o mux.Vars()
	params := mux.Vars(r)
	// ocorre que o metodo acima vai retornar o id em str, mas precisamos dele em int. para isso, basta converte-lo, conforme abaixo
	ID, err := strconv.ParseInt(params["id"], 10, 64)
}
