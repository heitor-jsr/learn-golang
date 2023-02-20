package main

import "fmt"

// para recuperar a execução de algo que entrou em panico, vc precisa usar o if init, avaliando se a variavel é diferente de nulo. se for, a execução do programa é recuperada. a função recover vai armazenar o valor daquilo que a panic() retorna. se esse valor for diferente de nulo, a execução é retomada
func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("execução recuperar com sucesso")
	}
}

func aprovado(n1, n2 float32) bool {
	defer recuperar()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// a função panic interrompe o fluxo normal do seu programa, parando tudo. ela para de executar o que está executando. além disso, ela chama todas as funções com a cláusula defer, e se vc n recuperar essas funções com recover, o seu programa morre

	// panic é diferente de erros, porque eles permitem a execução do programa apos surgirem, com o devido tratamento. já o panic literalmente mata o seu programa se vc n tentar recurar ele com o recover

	panic("A MÉIDA É EXATAMENTE 6")

	// a lógica do recover para recuperar a execução da função que entra em panico se baseia na existencia de funções que possuem a clausula defer. isso porque, quando a aplicação entra em panic, o go vai executar todas as funções com a clausula defer antes de matar tudo e finalziar a execução
}

func main() {
	fmt.Println(aprovado(6, 6))
	fmt.Println("pós execução")
}
