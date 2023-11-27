package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Quadrado struct {
	altura  float64
	largura float64
}

// DESSA FORMA EM GO NÃO PRECISO DE FATO EXPLICITAR QUE ESTOU IMPLEMENTANDO UMA INTERFACE ELE JÁ FAZ ISSO AUTOMATICAMENTE
func (q Quadrado) area() float64 {
	return q.altura * q.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// aqui estou passando a minha interface falando que ela vai retornar apenas oq meu metodo area retornar
func escreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func main() {
	q := Quadrado{10, 15}
	escreverArea(q)

	c := Circulo{10}
	escreverArea(c)
}