package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)

	// o código abaixo vai causar um deadlock, porque o envio da informação ao canal vai bloquear a execução do programa, fazendo com que ele fique eternamente esperando algum lugar receber essa informação. ou seja, o programa nunca vai chegar na var msg.

	// para contonar isso podemos usar o buffer ao declarar um canal. o buffer só vai bloquear a execução do programa quando o canal atingir sua capacidade máxima. a capacidade máxima é definida pelo buffer.

	channel <- "oi mundo"

	msg := <-channel

	fmt.Println(msg)
}
