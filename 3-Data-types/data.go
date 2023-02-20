package main

import (
	"errors"
	"fmt"
)

func main() {
	// go suporta 4 tipos de dados inteiros: int8, int16, int32 e int64. a divisão é feita de acordo com o tamanho, em bites, que cada numero inteiro ocupa na memoria. quanto maior ou menor o numero, a classe deles é alterada.
	// se vc n especificar o tipoe de int que a variavel vai ser, o go vai utilizar a arquitetura do seu computador como base para definir o tipo de int que vai ser criado. ex: se o seu pc for 32bits, vai criar um int32, se for 64, vai criar um int64
	var num int16 = 100

	fmt.Println(num)

	var num2 uint32 = 23443

	fmt.Println(num2)

	// uint é um outro tipo de numero inteiro, que referencia um inteiro sem sinal. assim, ele n admite que vc utilize um sinal nele.

	// int32 = rune. normalmente esse alias é usado para tratar de numeros inteiros que representam caracteres da tabela asc
	var num3 rune = 1234
	fmt.Println(num3)

	// alias para o uint8. é a mesma coisa
	var num4 byte = 4
	fmt.Println(num4)

	// os numeros reais só admitem dois tipos: float32 e flaot64. a sintaxe dele é sempre com ponto e nunca com virgula. se vc utilizar a inferencia de tipos nesses caras, com o :=, o go vai utilizar a mesma lógica de se basear na arquitetura do seu pc para definir se a variavel é float32 ou float64.

	//para definir strings vc tem que usar sempre aspas duplas. usar aspas simples vai fazer com que o go leve em consideração o tipo char

	// as varivaeis char são representadas por um unico caractere, e te retornam o numero referente ao caractere em questão na tabela asc. se vc colocar mais de um caractere, o go retorna um erro de compilação. o tipo desse char é int.
	char := 'b'

	fmt.Println(char)

	// valor zero é o valor que é atribuido para uma variavel quando vc cria ela, mas não atribui um valor a ela. para as strings, o valor zero é a string vazia. para numeros, o valor 0 é 0. para erro, o valor zero é nil (null). o valor 0 do bool é false. os valores zero também são conhecidos como valores inciais

	var text int
	fmt.Println(text)

	var bol bool = true
	fmt.Println(bol)

	// em go, o erro é um tipo de dado proprio. para atribuir um valor a ele, vc tem que atribuir um valor do tipo error. esses valores são criados a partir de um  pacote chamado errors. como se ve abaixo, para criar um novo erro, basta usar o errors.New("nome do erro aqui")
	var erro error = errors.New("erro interno")
	fmt.Println(erro)

}
