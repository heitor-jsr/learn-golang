package main

import "fmt"

// a cláusula defer é muito usada em go. basicamente, ela vai adiar a execução de determinado pedaço de código por algum tempo

func func1() {
	fmt.Println("executando a func1")
}

func func2() {
	fmt.Println("executando a func2")
}

func main() {
	// o defer vai basicamente dizer ao go para adiar a execução de algo até o último momento possível. no nosso caso, o ultimo momento possivel é apos a execução da func2, pois a func main não tem retorno.

	// se a função tem um retorno, o defer é executado imediatamente antes do retorno ser dado.

	// é muito usada em banco de dados.

	defer func1()
	func2()
}
