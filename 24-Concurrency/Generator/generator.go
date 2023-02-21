package main

import (
	"fmt"
	"time"
)

// o generator é uma função que, basicamente, vai encapsular uma chamada para uma goroutine e devolver um canal para a gente.

func main() {
	channel := escrever("oi")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// a função generator respeita o formato abaixo. ao chamar ela na main, vc não precisa usar a clausula go, porque é a propria função generator que vai encapsular a chamada para uma goroutirne e retornar um canal para nos

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
