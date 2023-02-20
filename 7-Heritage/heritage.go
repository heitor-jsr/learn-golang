package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// como go não é uma linguagem feita para ser orientada a objetos, o mais perto que chegamos de herança é o caso abaixo, onde temos um struct mais genérico sendo atribuido para outro mais especifico, sem atribuir um tipo a ele. isso vai fazer com que todos os campos do struct mais genérico sejam "herdados" pelo struct mias especifico

type estudante struct {
	pessoa
	curso string
	facul string
}

func main() {

	// para ter um estudante, vc precisa ter criado uma pessoa, vez que o estudante herda da pessoa os campos informados.
	p1 := pessoa{"jao", "de tal", 22, 178}
	e1 := estudante{p1, "eng", "usp"}

	fmt.Println(e1)
}
