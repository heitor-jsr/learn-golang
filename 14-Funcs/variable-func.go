package main

import "fmt"

// funções variáticas são aquelas que podem receber N parametros, sendo 	dispensável informar a quantidade de parametros que ela vai receber
// ao criar uma função variática, go vai encapsular os valores que o usuario passar como parametro dentro de um slice.

func soma(nums ...int) int {
	total := 0

	for _, numero := range nums {
		total += numero
	}

	return total

}

// vc só pode ter um tipo de dado que vai servir como parametro variatico por função. ou seja, se o seu int é o parametro variatico da funcao abaixo, vc n pode ter um float32 como parametro variatico. OBS: o parametro vbariatico tem que ser obrigatoriamente o ultimo parametro.

func esc(text string, nums ...int) {
	for _, numero := range nums {
		fmt.Println(text, numero)
	}
}

func main() {
	totals := soma(1, 2, 3, 5, 5)
	fmt.Println(totals)

	esc("oi", 1, 2, 3, 4)
}
