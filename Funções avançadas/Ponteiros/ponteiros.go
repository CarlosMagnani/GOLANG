package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int)  {
	//AGORA EU RECEBO O ENDEREÇO DE MEMÓRIA DA MINHA VARIAVEL E ALTERANDO ELA DIRETAMENTE
	//NÃO É NECESSARIO UM RETORNO POIS EU ALTERO DIRETAMENTE O VALOR NA MEMÓRIA
	*numero = *numero * -1
}

func main() {
	//DESSA FORMA ESTOU ENVIANDO APENAS UMA CÓPIA DA MINHA VARIAVEL E NÃO ALTERANDO O VALOR DA MESMA
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)


	novoNumero := 40
	fmt.Println(novoNumero)
	//PARA CONSEGUIR PASSAR VARIAVEIS COM O ENDEREÇO DE MEMÓRIA É PRECISO COLOCAR O & NA FRENTE.
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}