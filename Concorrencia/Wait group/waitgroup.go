package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
func main()  {
	//AQUI CRIAMOS UM GRUPO DE ESPERA PARA QUE POSSAMOS DEFINIR QUANTAS FUNÇÕES COLOCAREMOS NESSE GRUPO 
	var waitGroup sync.WaitGroup
	//AQUI DEFINIMOS O NUMERO DE GOROUTINES PARA FICAR EM ""FILA""
	waitGroup.Add(2)

	//DEFINIMOS UMA GO ROUTINE E ASSIM QUE ELA ACABAR FALAMOS QUE ELA ESTÁ FINALIZADA
	go func(){
		escrever("Olá mundo")
		waitGroup.Done() //-1
	}()

	
	go func(){
		escrever("Olá goroutine")
		waitGroup.Done() // -1
	}()
	//ESTÁ FUNÇÃO ESPERA TODAS AS GOROUTINES DEFINIDAS SEREM FINALIZADAS.
	waitGroup.Wait() 
}