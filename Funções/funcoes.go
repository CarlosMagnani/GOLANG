package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}


func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)


	var f = func(txt string) string{
		fmt.Println("Funcao:", txt)
		return txt
	}
	f("HelloWorld")

	result := f("Testew")
	fmt.Println(result)


	resultadoSoma, resultadoSubtracao := calculosMatematicos(10,15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}