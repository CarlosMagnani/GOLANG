package main

import "fmt"

//REFERENCIA VARIAVEIS QUE ESTÃO FORA DO SEU ESCOPO
//COMO SE FOSSE UM "CALLBACK"
func closure() func()  {
	texto := "Dentro da função closure"

	funcao := func ()  {
		fmt.Println(texto)
	}

	return funcao
}

func main()  {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaonova := closure()
	funcaonova()
}