package main

import (
	"fmt"
	"time"
)

// em go, channel é uma estrutura que pode enviar e receber dados, sincronizando as goroutines. para cria-lo, basta usar o make(chan - palavra reservada para o tipo canal - string - tipo de dado que o canal vai suportar). isso significa que somente vai receber e enviar dados do tipo string

// os canais bloqueiam a execução do programa. ou seja, são métodos bloqueantes (como async e await em nodejs). dessa forma, apesar de a goroutine dizer para o programa continuar a execução dos códigos independete de ser finalizada determinaad função, o canal vai bloquear a execução, até que receba ou envie um dado.

// no exemplo abaixo, isso acaba causando um problema, pq o uso do canal não vai permitir a execução completa do loop for. ou seja, a partir do momento em que o canal recebe um dado, o go vai continuar a execução da aplicação e finalizar ela.

// deadlock => situação na qual o canal está aguardando receber algum dado, mas não existe mais nenhuma ocasião em que são enviados dados ao canal. isso implica que a aplicação aguarda eternamente algo que nunca vai chegar. OBS: não é erro em compilação, apenas em execução. para resolver o problema, vc deve verificar se o canal esta aberto ou fechado, e se estiver fechado, "matar" a execução do loop infinito.

// para evitar o deadlock, o recomendável é verificar se o canal está aberto ou não antes de continuar a execução do programa. no for abaixo, isso é feito com a verificação de open (pode ser dado qualquer nome pra essa variavel).

func main() {
	channel := make(chan string)

	go escrever("oi", channel)

	// vai dizer que o programa deve esperar o canal receber a informação para continuar a sua execução

	for {
		msg, open := <-channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

	//o mesmo de cima pode ser feito com for mensagem := range canal {} => isso vai dizer que, enquanto o canal estiver aberto, o go vai executar a função dentro de {}

}

func escrever(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		// a sintaxe abaixo evidencia que a informação enviada pelo text vai ser transmitida pelo canal passado como parametro
		channel <- text
		time.Sleep(time.Second)
	}
	// ao finalizar o for acima, o canal é fechado para comunicação. não recebe e nem envia mais dados.
	close(channel)
}
