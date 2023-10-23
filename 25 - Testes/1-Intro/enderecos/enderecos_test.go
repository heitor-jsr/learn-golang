package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rua Paulista", "rua"},
		{"Estrada Paulista", "estrada"},
		{"rodovia Paulista", "Tipo inválido"},
		{"alameda Paulista", "Tipo inválido"},
		{"praça Paulista", "Tipo inválido"},
		{"quadra Paulista", "Tipo inválido"},
		{"quadra Paulista", "quadra"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo de endereço recebido é diferente do tipo esperado! Esperava %s e recebeu %s", cenario.enderecoEsperado, tipoDeEnderecoRecebido)
		}
	}

}
