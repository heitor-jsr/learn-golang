package main

import "fmt"

// função de retorno noemado vai te retornar os valores em variaveis nomeadas. como se ve abaixo, a função calc vai retornar dois valores, um de nome soma, e outro de nome sub. veja que no corpo da funcao vc n precisa criar as variaveis, pois elas ja foram criadas na declaração da função (na sua assinatura)

func calc(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := calc(6, 2)

	fmt.Println(sum, sub)
}
