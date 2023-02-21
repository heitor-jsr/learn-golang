package main

import "fmt"

// worker pools funciona como uma fila de tarefas para ser executada, onde vc tem varios processos pegando itens dessa fila de maneira independente. na função abaixo, isso faz com que varios processos pegem um numero da fila de fibonacci e realize o calculo de maneira independente

func main() {

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	// para otimizar com worker pools, vc pode chamar a função worker quantas vezes quiser. isso vai criar um novo processo que, por sua vez, vai pegar mais um item da fila e executar uma das tarefas

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}
