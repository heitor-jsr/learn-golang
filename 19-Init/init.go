package main

import "fmt"

// a função init é uma função que é executada antes mesmo da função main. vc pode ter uma func init por arquivo, diferente da main, que deve ser uma por pacote.

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
