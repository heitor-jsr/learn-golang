package main

import "fmt"

// basicamente, closure são funções que fazem referencia a variaveis que estão fora do seu corpo. ao ser criada, a função closure manter a referencia inicial de seus parametros e argumentos

func closure() func() {
	text := "dentro de closure"
	funcao := func() {
		fmt.Println(text)
	}
	return funcao
}

func main() {
	new := closure()
	new()
}
