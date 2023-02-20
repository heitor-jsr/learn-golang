package main

import "fmt"

// funções recursivas são funções que chamam ela mesma. são funções que, para funcionar, precisam chamar ela mesma recursivamente. OBS: por serem recursivas, vc precisa especificar um ponto de parada para uma função recursiva. é a partir desse momento que a função vai parar de se chamar e encerrar a sua execução.

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {
	pos := uint(10)
	fmt.Println(fibonacci(pos))
}
