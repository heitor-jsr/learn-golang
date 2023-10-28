package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
func GetUserById(w http.ResponseWriter, r *http.Request) {
	// primeira coisa que precisamos ter em mente é que aqui nos precisamos ler o parametro de busca da rota (o id que é passado para ela vai representar o id que o cliente quer recuperar do db). para isso, basta usar o mux.Vars()
	params := mux.Vars(r)
	// ocorre que o metodo acima vai retornar o id em str, mas precisamos dele em int. para isso, basta converte-lo, conforme abaixo
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	line, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
		return
	}

	var usuario usuario
	// mesma coisa do for do metodo anterior. a unica diferença é que vai verificar se a linha com o dado retornado do banco de dados existe. se existir, vai retornar os dados para o cliente, somente se o scan não retornar nenhum erro quando comparar os dados retornados com os tipos do struct.
	if line.Next() {
		if err := line.Scan(&usuario.ID, &usuario.Name, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao retornar os dados"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao retornar os dados"))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
		return
	}
	// depois de ler o ID, como no caso anterior, aqui vamos precisar ler o corpo da nossa requisição, pra somente depois abrir a conexão com o db.
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	// mesma coisa do metodo post, vai converter o json recebido para um struct, e armazenar ele dentro da referencia de memoria da variavel inicializada acima.
	if err = json.Unmarshal(reqBody, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o json para struct"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	// o proximo passo é realizar o statement. o statement é feito em todas as operações que não envolvem consulta no banco de dados, como insert, update, delete.
	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	// uma coisa importante é que o statement vai sempre retornar dois dados: a quantidade de linhas afetadas e o id inserido ou atualizado. como nesse caso não precisamos de todas as informações, podemos ignorar o primeiro retorno dela.

	if _, err = statement.Exec(usuario.Name, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}