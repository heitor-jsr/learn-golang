package main

import (
	"fmt"
	"math"
)

// as interfaces só aceitam assinaturas de metodos, que mostram o que o metodo vai fazer. diferente de struct, que admite assinatura de dados de tipos variados, as interfaces somente aceitam metodos

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("a área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
