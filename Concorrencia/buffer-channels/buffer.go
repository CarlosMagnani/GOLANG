package main

import "fmt"

func main() {
	//AO COLOCAR O SEGUNDO PARAMETRO EU MOSTRO PARA MEUS CANAIS QUE EU DEFINO UM "BUFFER"
	//OU SEJA EU DEFINO UMA CAPACIDADE SENDO ASSIM ELE NÃO É BLOQUEANTE POIS DESSA FORMA EU NÃO PRECISO ESPERAR 
	//ALGUEM RECEBER MINHAS MENSAGENS PELOS CANAI. ASSIM EU CONSIGO TRABALHAR COM CANAIS NA MESMA FUNÇÃO
	canal := make(chan string, 2)

	canal <- "Olá Mundo"
	canal <- "Olá Mundo, programando em golang"

	msg := <-canal
	msg2 := <- canal
	fmt.Println(msg)
	fmt.Println(msg2)
}