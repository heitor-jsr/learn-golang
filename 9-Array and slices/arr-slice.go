package main

import (
	"fmt"
)

func main() {
	// arrays em go são estruturas homogêneas (armazenam um mesmo tipo de dado) e estáticas (são estruturas fixas - permanecem com o mesmo tamanho/não é possível modificá-lo). além disso, são estruturas indexadas, ou seja, vc acessa os elementos do array de acordo com o indíce de cada elemento dentro da estrutura.
	// se vc n definir o conteúdo de um array, o go vai popular ele com os 0 do tipo que vc definiu para os dados que vão estar dentro do array.

	var arr1 [5]int

	arr1[0], arr1[1], arr1[2] = 7, 3, 9

	fmt.Println(arr1)

	// para popular o array criado basta usar a sintaxe abaixo. OBS: se tentar incluir um elemento a mais no array (um elemento que vai extrapolar o numero de itens definido na sua criação), o Go vai apontar um erro de compilação

	// se vc declarar um arr da maneira como abaixo, o go vai fixar o tamanho do arr de acordo com a quantidade de intes que cv informou que ele tem

	arr2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr2)

	// já o slice é uma estrutura que é baseada no arr, mas com uma maior flexibilidade. para cria-lo, a sintaxe é a debaixo. veja que vc n precisa atribuir um tamanho para o slice

	s1 := []int{1, 2, 3}

	// como o nome sugere, o slice representa uma fatia, um pedaço de algo. com isso, é possível que tenhamos um array com 20 posições, sendo que existem 4 slices que apontam para 5 posições distintas desse array. com isso, apesar de os slieces representarem posições (e consequentemente valores) diferentes do array, eles apontam para o mesmo lugar na memória em que o array está armazenado.

	s2 := arr1[1:3]
	//slice cirado que aponta para as posições de indice 1 a 3 do array a1, excluindo, entretanto, a posição final- representa, portanto, o 2 e o 3.
	// OBS: esse s2 não gerou um novo array! ele apenas utilizou o seu proprio ponteiro para apontar para o primeiro elemento do array que ele referencia, e ir até o tamanho final definido pelo slice. dessa forma, se no array original o valor de algum desses indices mudar, o valor retornado pelo slice tambem vai mudar

	s3 := arr1[:2]
	// vai do inicio do array a1 até o elemento de index 2.

	s4 := arr1[3:]
	// vai do index 3 do array a1 até o final dele.

	// vc pode criar um sliece a partir de outroo: s5 := s3[:1]

	fmt.Println(s1, s2, s3, s4)

	// a função append é exclusiva dos slices em go. ela vai adicionar um item no slice em qeustão, e te retornar um novo slice com esse item na posição desejada. se vc atribuir esse retorno para o slice que está sendo alterado, ele vai receber um valor novo.

	s3 = append(s3, 8)
	fmt.Println(s3)

}
