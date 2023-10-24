package formas

import "testing"

func TestArea(t *testing.T) {
	t.Run("Retangulo com 10 e 5", func(t *testing.T) {
		ret := Retangulo{Altura: 10, Largura: 5}

		areaEsperada := float64(50)
		areaRecebida := ret.area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida é diferente da área esperada! Recebeu %f e esperava %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo com 10 e 5", func(t *testing.T) {
		cir := Circulo{raio: 10}

		areaEsperada := float64(314.1592653589793)
		areaRecebida := cir.area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida é diferente da área esperada! Recebeu %f e esperava %f", areaRecebida, areaEsperada)
		}
	})
}
