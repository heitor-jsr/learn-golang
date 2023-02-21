package main

import "fmt"

// as funções em go podem ter ponteiros como argumentos, retornos, etc

func inverter(num int) int {
	return num * -1
}

// a função abaixo recebe um ponteiro de numero. para alterar o seu valor, precisamos desreferenciar esse valor e multiplica-lo por -1. com isso, vc vai alterar o valor armazenado na variavel que é passada como argumento para a funcao
func invertedSinal(num *int) {
	*num = *num * -1
}

func main() {
	numero := 12
	inverted := inverter(numero)
	fmt.Println(inverted)

	newNum := 40
	fmt.Println(newNum)
	invertedSinal(&newNum)
	fmt.Println(newNum)

}
