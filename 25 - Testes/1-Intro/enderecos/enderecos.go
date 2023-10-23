package enderecos

import "strings"

// TipoEndereco retorna um tipo de endereco válido
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada"}

	primeiraPalavra := strings.ToLower(strings.Split(endereco, " ")[0])

	enderecoTemTipoValido := false

	for _, tipos := range tiposValidos {
		if tipos == primeiraPalavra {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return primeiraPalavra
	}

	return "Tipo inválido"
}
