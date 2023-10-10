package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual 15")
	}

	//VARIAVEL JÁ INICIADA, FUNCIONA SOMENTE NESTE ESCOPO
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero", outroNumero)
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10", outroNumero)
	}else{
		fmt.Println("Entre 0 e -10")
	}


}