package main

import (
	"fmt"
	"sync"
	"time"
)

// para criar o waitGroup basta usar a forma padrão abaixo. no método add não precisa referenciar as goroutines, apenas dizer a quantidade delas que o waitgroup vai abranger. depois disso, basta criar as suas goroutines e, quando ela for finalizada, vc chama o metodo Done do waitgroup para diminuir um do contador informado no metodo add. o metodo Wait informa para a func main esperar todas as goroutines serem finalizadas para poder ser encerrada.

// sem o waitGroup, o go vai encerrar a execução da func main sem sequer iniciar as execuções das goroutines.

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("oi")
		waitGroup.Done()

	}()

	go func() {
		escrever("tchau")
		waitGroup.Done()

	}()

	waitGroup.Wait()

}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
