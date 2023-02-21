package main

import (
	"fmt"
	"time"
)

// concorrencia é diferente de paralelismo. enquanto o paralelismo permite a execução de duas tarefas ao mesmo tempo em processadores com mais de um nucleo, a concorrencia permite a execução de duas tarefas ao mesmo tempo, inclusive em processadores com um unico nucleo (uma tarefa n espera a outra acabar - elas revezam o tempo de uso do nucleo do processador).

func main() {
	go escrever("oi") // nesse cesnário, a segunda chamada de função nunca é executada. uma maneira de contornar isso é usando uma goroutine. uma goroutine é um método que é invocado da maneira ao lado. isso faz com que o Go executa a função da linha 11, mas informa ao Go para não esperar ela ser finalizada para prosseguir com a execução do programa

	// basicamente, a goroutine vai dizer para vc executar algo, mas não esperar a execução ser finalizada para que a sua aplicação continue a ser executada.
	escrever("tchau")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
