package main

import "fmt"

func main() {
	// os ponteiros nada mais são do que uma variável que recebe o endereço na memória de uma variável em específico (ele n guarda o valor, mas o endereço de memoria para o valor em questão). ou seja, o ponteiro é uma referencia de memória

	var num int = 100

	var pointer *int

	num++

	pointer = &num

	num++

	fmt.Println(pointer)
	// para acessar o valor que o endereço de memória do ponteiro aponta, basta colocar o asteristico na sua frente. esse processo é chamado de desreferenciação
	fmt.Println(*pointer)

}
