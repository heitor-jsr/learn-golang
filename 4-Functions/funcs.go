package main

import "fmt"

// para declarar funções em go vc precisa tipar o argumento que ela recebe, assim como o retorno dela. ou seja, vc tem que informar ao programa qual o tipo de dado que a função vai retornar tb
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// as funções em go podem ter mais de um retorno. para definir isso, basta colocar o tipo do retorno da função entre parenteses. a quantidade de tipos colocados vai definir a quantidade de retornos da sua função.

func calcMat(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("oi")

	resultado := f("func 1")
	fmt.Println(resultado)

	// para armazenar os valores de uma func com vários retornos, vc precisa declarar a mesma quantidade de variaveis do lado esquerdo da chamada da func, onde cada uma vai receber um dos valores que é retornado pela func, como abaixo.
	// caso vc n queira armazenar todos os valores retornados pela func em uma variavel, vc pode colocar, no lugar da variavel que armazenaria o valor, um underline.
	resultSum, resultSub := calcMat(20, 15)
	fmt.Println(resultSum, resultSub)
}
