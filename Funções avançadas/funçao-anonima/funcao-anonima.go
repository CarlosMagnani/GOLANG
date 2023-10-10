package main

import "fmt"

func main() {

	//FUNÇÃO ANONIMA É EXECUTADA LOGO APÓS SER DECLARADA
	retorno := func(texto string) string{
		return fmt.Sprintf("RECEBIDO -> %s", texto)
	}("Passando parâmetro")
	
	fmt.Println(retorno)

}