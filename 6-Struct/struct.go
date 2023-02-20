package main

import "fmt"

// a forma de criar o struct é a abaixo. ele é um tipo em go.

type user struct {
	nome     string
	idade    uint8
	endereco address
}

type address struct {
	stree  string
	number uint8
}

func main() {
	// em go, struct nada mais é do que uma coleção de campos, onde cada um tem um nome e um tipo. ele é usado quando vc tem várias informações que precisam ser salvas, de tipos diferentes, e que são relativas a uma unica coisa/pessoa.
	fmt.Println("Struct")

	// criar um user e passando as suas informações. se vc n passar nenhum valor para o struct, ele vai atribuir o valor 0 para cada um dos tipos de informações que estão inseridas nele.
	var u user
	fmt.Println(u)

	// populando o user

	u.nome = "heitor"
	u.idade = 20

	fmt.Println(u)

	endereco := address{"rua de tal", 2}

	// inferencia de tipo para struct
	u2 := user{"vinicius", 21, endereco}

	fmt.Println(u2)

	// para atribuir valor a apenas um dos campos, é necessário informar qual deles vc vai querer atribuir o valor em questão, como abaixo.
	u3 := user{idade: 23}
	fmt.Println(u3)

	// nada impede que vc tenha um aninhamento de structs.
}
