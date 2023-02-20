package main

import "fmt"

func main() {
	numer := 10

	if numer > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	// if init é uma estrutura de controle if/else que realiza uma avaliação do seu código, dada uma condição, e inicializa uma variavel se a condição for preenchida.

	// a maneira abaixo é como se faz um if init. nesse caso, vc declara uma variável, atribui um valor a ela e verifica se ela é maior que 0. se for, o bloco de código abaixo é executado.

	// nos casos de if init, a variavel so fica acessivel dentro do escopo do if. fora dele, vc não consegue acessar essa variavel.

	if outroNum := numer; outroNum > 0 {
		fmt.Println("numero é maior que 0")
	}
}
