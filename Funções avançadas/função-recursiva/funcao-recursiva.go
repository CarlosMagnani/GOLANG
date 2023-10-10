package main

import "fmt"

//FUNCS RECURSIVAS DEPENDEM DELAS MESMAS PARA CONTINUAR SENDO EXECUTADAS, CHAMAM ELAS MESMAS

//É NECESSÁRIO MOSTRAR PARA O GO QUANDO PARAR, POR CONTA DO 'STACKEROVERFLOW' ONDE SOBRECARREGA O SISTEMA
func fibonnaci(posicao uint)uint  {
	//CONDIÇÃO DE PARADA
	if posicao <= 1{
		return posicao
	}

	return fibonnaci(posicao - 2) + fibonnaci(posicao - 1)
}
func main()  {
	posicao := uint(10)
	
	for i := uint(0); i < posicao; i++{
		fmt.Println(fibonnaci(i))
	}
	fmt.Println(fibonnaci(posicao))
}