package main

import "fmt"

func main() {

	var variavel1 int = 10
	//ISTO É SOMENTE UMA CÓPIA NÃO TEM RELAÇÃO ENTRE ELAS.
	var variavel2 int = variavel1
	fmt.Println(variavel2)
	variavel1++


	//O PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int

	//O & MOSTRA O ENDEREÇO DA MEMÓRIA ONDE ESTÁ O VALOR ESTÁ ARMAZENADO
	ponteiro = &variavel3
	// O * DESFAZ A REFERENCIA E PEGA O VALOR QUE ESTÁ DENTRO DO ENDEREÇO DA MEMÓRIA
	fmt.Println(variavel3, *ponteiro)


	//O PONTEIRO ACESSA O VALOR CONFORME ELE MUDA E NÃO FAZ UMA CÓPIA
	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}