package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 1000000000000000000
	fmt.Println(numero)

	//alias
	//INT32  = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)


	var numeroReal float32 = 123.32
	fmt.Println(numeroReal)

	
	var numeroReal2 float64 = 123.3232323
	fmt.Println(numeroReal2)

	numeroReal3 := 1213.1313
	fmt.Println(numeroReal3)

	//STRING
	var str string = "teste3"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//BOOLEAN
	var booleenol bool
	fmt.Println(booleenol)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}