package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sábado"
	default:
		return "numéro inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	// falltrough é uma cláusula que vai avaliar a condição, executar o código que passou na condição, e não vai pular "para fora" da estrutura de controle. na verdade, vai empurrar o seu código para o proximo bloco de código, sem avaliar a condição que ele tá vinculado, apenas executando o código.

	switch {
	case numero == 1:
		diaDaSemana = "domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "segunda"
	case numero == 3:
		diaDaSemana = "terça"
	case numero == 4:
		diaDaSemana = "quarta"
	case numero == 5:
		diaDaSemana = "quinta"
	case numero == 6:
		diaDaSemana = "sexta"
	case numero == 7:
		diaDaSemana = "sábado"
	default:
		diaDaSemana = "numero inválido"
	}
	return diaDaSemana
}

func main() {
	dia := diaDaSemana(6)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}
