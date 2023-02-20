package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("increventa 1")
		time.Sleep(time.Second)
	}
	// loop com atribuição de valor para a variavel
	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j")
		time.Sleep(time.Second)
	}

	// quando vc vai iterar sobre algo, como array, slice, string, vc deve usar o for com a clausula range. ou seja, enquanto vc estiver no "range"/tamanho da estrutura que vc está iterando, execute o código em questão.

	nomes := [3]string{"jao", "thiago", "bruno"}

	// por padrao, o primeiro argumento do for com range é o indice. por isso, caso vc n vá usa-lo, vc precisa colocar o _ antes do segundo argumento.

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// no caso de loop sob uma string, o go vai te retornar por padrao o valor das letras na tabela ASCII. para retornar a letra em si, vc precisa usar string(letra)

	for i, letra := range "LETRA" {
		fmt.Println(i, string(letra))
	}

	// vc n pode usar o for range em um struct. apenas em map, array, slice e strings.

	// vai executar o codigo infinitamente
	for {
		fmt.Println("executando infinitamente")
		time.Sleep(time.Second)
	}
}
