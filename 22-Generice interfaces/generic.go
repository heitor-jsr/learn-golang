package main

import "fmt"

// em go existe um tipo genérico. ele é criado a partir de uma interface genérica, permitindo que sua aplicação aceite qualquer coisa. para criar esse tipo gnérico, basta criar uma interface sem nada dentro, pois ela vai aceitar qualquer coisa. no código abaixo, temos uma função generica que recebe uma interface vazia como parametro. isso implica que todo e qualquer metodo atende essa interface, vez que ela não estabelece restrições.

// com isso, ao chamar a func generic, vc pode passar parametros de diversos tipos: numero, str, float, etc. isso implica em uma ausencia de tipagem naquilo que a func vai receber e retornar.

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic(1)
	generic("eqw")
	generic(1.2)
	generic(true)
}
