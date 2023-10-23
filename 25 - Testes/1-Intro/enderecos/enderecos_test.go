package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTest := "Avenida Paulista"

	tipoEsperado := "avenida"

	tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTest)

	if tipoDeEnderecoRecebido != tipoEsperado {
		t.Errorf("O tipo de endereço recebido é diferente do tipo esperado! Esperava %s e recebeu %s", tipoEsperado, tipoDeEnderecoRecebido)
	}

}
