package main

import (
	"fmt"
	"time"
)

func main() {

	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			channel1 <- "canal1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "canal2"
		}
	}()

	for {

		// veja que no código abaixo o canal 1 está sendo prejudicado pela execução do canal 2. isso porque o canal 1 está pronto para receber um dado a cada 1 segundo, mas só recebe apos o canal 2 receber a informação dele, o que leva 5 segundos. considerando que o canal 1 não depende da execução do canal 2, podemos otimizar isso com o select.

		// basicamente, o select vai validar se algum dos canais já esta pronto para receber uma informação. se estiver, ele vai executar o código vinculado ao case. se não, passa para o prox case.

		select {
		case msgChannel1 := <-channel1:
			fmt.Println(msgChannel1)
		case msgChannel2 := <-channel2:
			fmt.Println(msgChannel2)
		}
	}
}
