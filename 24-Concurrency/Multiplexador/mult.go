package main

import (
	"fmt"
	"time"
)

// a ideia por tras dele é pegar dois ou mais canais e juntar em um unico canal. para isso, usamos uma func multiplexar, que recebe dois ou mais canais como entrada, e recebe apenas um.

func main() {
	channel := multiplexar(escrever("oi"), escrever("tchau"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexar(chan1, chan2 <-chan string) <-chan string {
	channel := make(chan string)

	// a lógica por traz desse padrão de canal é que o select vai verificar qual canal está disponivel para enviar uma msg, e, ao verificar, enviar a informação para um canal de saida unico. isso vai centralizar a troca de mensagem de dois canais em um unico canal

	go func() {
		for {
			select {
			case msg := <-chan1:
				channel <- msg
			case msg := <-chan2:
				channel <- msg
			}
		}
	}()
	return channel
}

func escrever(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("valor recebido: %s", text)
			time.Sleep(time.Second)
		}
	}()

	return channel
}
