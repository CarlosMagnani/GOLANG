package main

import "fmt"

//"HERANÇA" EM GO.
type pessoa struct {
	altura float32
	idade int
	nome string
	sobrenome string
}

//PASSANDO OS TIPOS DE PESSOA PARA ESTUDANTE, PASSANDO APENAS A STRUCT PARA CHAMAR SOMENTE O TIPO
type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main()  {
	p1 := pessoa{1.67, 19, "Carlos", "Magnani"}

	e1 := estudante{p1, "ADS", "UNIP"}

	fmt.Println(e1)
}