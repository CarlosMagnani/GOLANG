package main

import (
	"fmt"
	"time"
)

func main()  {
	canal1, canal2 := make(chan string), make(chan string)


	go func ()  {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal1"
		}
	}()

	go func ()  {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal2"
		}
	}()


	for {
		// UTILIZANDO A CLAUSULA SELECT EU NÃO PRECISO NECESSARIAMENTE ESPERAR QUE A MENSAGEM DO CANAL 2 SEJA RECEBIDA PARA EU COMPLETAR A DO CANAL 1
		// DESSA FORMA EU ESTOU DIZENDO QUE CASO A MENSAGEM DO CANAL1 JÁ ESTEJA PRONTA PRINTA NA HORA E O MESMO VALE PARA O CANAL2
		// ASSIM EU CONSIGO DESACOPLAR A DEPENDENCIA DE ESPERA DOS CANAIS PARA PRINTAR AS MENSAGENS 
		select {
		case mensagem := <-canal1:
			fmt.Println(mensagem)

		case mensagem2 := <- canal2:
			fmt.Println(mensagem2)
		}
	}
}