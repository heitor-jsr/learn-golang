package main

import "fmt"

// como go não é orientado a objetos, ele possui algumas estruturas de dados que permitem o uso de métodos, assim como em POO.

type user struct {
	nome  string
	idade uint8
}

// a sintaxe do metodo é a abaixo: vc declara uma func (que funciona como method), diz a qual estrutura ela referencia (o "u" serve para vc fazer referencia aos campos da estrutura struct. p ex: u.nome, u.idade, etc), e em seguida o nome desse metodo

func (u user) salvar() {
	fmt.Println("dentro de salvar")
}

// os métodos também podem alterar um valor dentro de uma determinada estrutura, usando a logica dos ponteiros

func (u *user) alterar() {
	u.idade++
}

func main() {
	user1 := user{"user 1", 20}
	fmt.Println("main")
	user1.salvar()

	user1.alterar()

	fmt.Println(user1)
}
