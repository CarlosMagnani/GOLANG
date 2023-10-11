package main

import "fmt"

func funcao1() {
	fmt.Println("FUNCAO 1")
}

func func2() {
	fmt.Println("FUNCAO 2")
}

func alunoAprovado(n1, n2 float32)bool {
	//EVITANDO REPETIÇÃO DE CÓDIGO PODEMO UTILIZAR O DEFER PARA ANTES DE REALIZAR O RETORNO EXECUTAR OQUE ESTAVAMOS ESPERANDO
	//MUITO UTIL PARA MANIPULAÇÃO DE BANCO DE DADOS, PARA FAZER O FECHAMENTO DE CONEXAO DE DADOS CASO SUCESSO OU ERRO
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Calculando as médias")
	media := (n1 + n2 ) / 2

	if media >= 6 {
		return true 
	}

	return false
}

func main() {
	//CLAUSULÁ DEFER SIGNIFICA PARA O GO QUE ELE DEVE ADIAR A EXECUÇÃO, ATÉ O ULTIMO MOMENTO.
	//SIMILAR AO AWAIT DO JS
	defer funcao1()
	func2()
	fmt.Println(alunoAprovado(7,8))
}