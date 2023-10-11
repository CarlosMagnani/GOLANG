package main

import "fmt"

//FUNÇÃO QUE RECUPERA UMA OUTRA FUNÇÃO QUE ESTÁ EM PÂNICO
//CASO NÃO TENHA NENHUM PANICO ELA É IGNORADA
func recuperandoFunao()  {
	if r := recover(); r != nil{
		fmt.Println("Execução recuperada com SUCESSO")
	}	
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperandoFunao()

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//O PANIC PARA A EXECUÇÃO MAS ANTES ELE CHAMA AS FUNÇÕES QUE TEM O DEFER
	//NÃO É A MELHOR MANEIRA DE SE TRATAR ERROS NO SEU PROGRAMA MAS É UTILZIADO EM BIBLIOTECAS EXTERNAS
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6,6))
	fmt.Println("Pós Execução")
}