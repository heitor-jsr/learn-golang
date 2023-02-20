package main

import "fmt"

func main() {
	som := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	rest := 10 % 2

	// vc n pode realizar operações com dados de tipos diferentes em go. p. ex., vc n pode tentar somar um int8 com int16.

	fmt.Println(som, sub, div, mult, rest)

	//Atribuição
	var variavel1 string = "str"
	// esse tipo de atribuição inicializa também a var, além de atribuir. esse tipo de atribuição também realiza uma inferência de tipo, caso o usuario não o informe.
	var2 := "str2"
	fmt.Println(variavel1, var2)

	// Relacionais
	// sempre vão retornar um booleano (true/false)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!false)

	// Unários - operadores que somente interagem com uma var por vez

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	fmt.Println(numero)
	numero *= 3
	numero /= 2
	numero %= 3
	fmt.Println(numero)

	// Em go não temos o operador ternário. assim, precisamos usar o if/else

}
