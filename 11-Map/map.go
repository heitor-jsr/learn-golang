package main

import "fmt"

func main() {
	// maps em go são estruturas que contém um par de chave:valor. as chaves devem ser únicas dentro de Go. tanto as chaves quanto os valores dos maps possuem um tipo homogeneo (não mudam).

	// OBS: os maps devem ser criados com a sintaxe :=; ou seja, devem ser inicializados na sua criação, caso contrário o Go vai apontar um erro.

	// é possível percorrer o map com o for, passando pelas suas chaves e seus valores de uma unica vez, não sendo necessário realizar uma dupla iteração (como no javascript)

	// para deletar uma chave:valor de um map, basta usar a func delete(nomeDoMap, chave)

	aprovados := make(map[int]string) // uma forma de criar

	users := map[string]string{
		"nome":      "heitor",
		"sobrenome": "scalon",
	} // outra forma de criar maps. o que está dentro das chaves é o tipo do valor da chave, e o que está fora é o tipo do valor dos valores.

	// acessar o valor de uma chave única
	fmt.Println(users["nome"])

	fmt.Println(users)

	aprovados[12345678900] = "heitor"
	aprovados[12345678300] = "vinicius"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// forma mais fácil de criar um map:

	funcsESalarios := map[string]float64{
		"José":   1900.00,
		"João":   2500.00,
		"Heitor": 5255.55,
	}

	// adiciona uma chave:

	funcsESalarios["Vinicius"] = 2250.00

	fmt.Println(funcsESalarios)

	// maps aninhados:

	funcsPorSetor := map[string]map[string]float64{
		// veja que a primeira chave representa uma str, e o seu valor é um map, que por sua vez tem como chave uma str e como valor um float
		"Tecnologia": {
			"Heitor":   5520.00,
			"Vinicius": 5500.00,
		},
		"Juridico": {
			"Felipe": 5000.00,
		},
		"Contábil": {
			"Amanda": 5000.00,
		},
	}

	for setor := range funcsPorSetor {
		fmt.Printf("Funcionários do setor de %s\n", setor)
		for nome, salario := range funcsPorSetor[setor] {
			fmt.Printf("%s tem o salário de %.2f\n", nome, salario)
		}
	}

	// map que recebe um map como valor para a sua chave:

	users3 := map[string]map[string]string{
		"nome": {
			"nome":      "heitor",
			"sobrenome": "ribeiro",
		},
	}

	fmt.Println(users3)
}
