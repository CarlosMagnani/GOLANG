package main

import "fmt"

//NESTA FUNÇÃO ELA RECEBE N NÚMEROS DO TIPO DEFINIDO, NO FINAL É CONVERTIDO PARA UM SLICE
//UMA FUNÇÃO SÓ PODE TER UM PARAMETRO VARIÁTICO E NECESSARIAMENTE TAMBÉM PRECISA SER O ULTIMO PARAMÊTRO RECEBIDO
func soma(numeros ...int) int{
	total := 0

	for _, numero := range numeros{
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int)  {
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main()  {
	totalSoma := soma(1,2,3,4,5,6,7,8,9)
	fmt.Println(totalSoma)

	escrever("roi", 12,13,4,5)

}