package main

import "fmt"

func main() {

	// em go, a sintaxe para função anonima é a debaixo. vc declara a func(aqui podemos ter parametros), abre as chaves, escreve o corpo da func, e depois chama ela com ()

	func() {
		fmt.Println("oi")
	}()

	func(text string) {
		fmt.Println(text)
	}("precisa passar o parametro aqui")

	// variavel que recebe o retorno de uma func anonima. OBS: elas admitem retorno
	retorno := func(text string) string {
		return fmt.Sprintf("recebido -> %s", text)
	}("precisa passar o parametro aqui")

	fmt.Println(retorno)
}
