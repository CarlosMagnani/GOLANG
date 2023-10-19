package main

import "fmt"

var n int
//É EXECUTADA ANTES DA FUNÇÃO MAIN, PODE SER UTILIZADA PARA CARREGAR CONFIGS.
//A ORDEM NÃO IMPORTA, ASSIM COMO A MAIN A INIT SÓ PODE TER UMA.
func init() {
	n = 10
	fmt.Println("Rodando a função init")
}

func main()  {
	fmt.Println(n)
	fmt.Println("Rodando a função")
}

