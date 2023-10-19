package main

import (
	"fmt"
	"math"
)

// EU CRIO NA MINHA INTERFACE SOMENTE A ASSINATURA DO MÉTODO
//TODAS AS FUNC PRECISAM TER ESSA ASSINATURA PARA PODER ESCREVER A AREA
//DESSA FORMA EU CONSIGO ABSTRAIR OS TIPOS E TRAZER MAIS FLEXIBILIDADA PARA AS STRUCTS DO GOLANG
type forma interface {
	area() float64
}

type quadrado struct {
	altura  float64
	largura float64
}

//DESSA FORMA EM GO NÃO PRECISO DE FATO EXPLICITAR QUE ESTOU IMPLEMENTANDO UMA INTERFACE ELE JÁ FAZ ISSO AUTOMATICAMENTE
func(q quadrado) area() float64 {
	return q.altura * q.largura
}


type circulo struct {
	raio float64
}


func(c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}
// aqui estou passando a minha interface falando que ela vai retornar apenas oq meu metodo area retornar
func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func main() {
	q := quadrado{10, 15}
	escreverArea(q)

	c := circulo{10}
	escreverArea(c)
}