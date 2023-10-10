package main

import "fmt"

//RETORNO NOMEADO CONSISTE EM DECLARAR A VARIAVEL QUE VAI SER RETORNADA DIRETAMENTE NO ESCOPO DA FUNÇÃO
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 + n2
	return
}

func main()  {
	soma, subtracao := calculosMatematicos(10, 20)

	fmt.Println(soma, subtracao)
}