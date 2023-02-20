package main

import "fmt"

func main() {
	// a função make vai alocar um espaço na memória do nosso pc para armazenar alguma coisa. no caso abaixo, vamos usar ela para criar um slice. a função vai receber o tipo de dado, o tamanho do slice e a quantidade maxima (opcional) de itens que o slice pode ter.

	// por baixo dos panos, a função make vai criar, em um primeiro momento, um array de 15 posições. porém, depois de fazer isso, ela nos retorna os 10 primeiros itens desse array que foi criado.

	// a mesma coisa acontece quando vc cria um slice a partir "do nada", sem referenciar um array explicitamente. nesse caso, o go cria um array interno que vai servir de referencia para o slice em questão.

	// quando o go percebe que o seu slice vai estoruar o tamanho do array interno, ele cria um novo array com o dobro do tamanho do atual, fazendo com que o slice aponte para esse novo array agora. por isso que o slice não possui um tamanho máximo

	slice3 := make([]float32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho atual do slice
	fmt.Println(cap(slice3)) // capacidade total do slice

}
