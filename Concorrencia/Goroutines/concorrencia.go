package concorrencia

import (
	"fmt"
	"time"
)

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
//UTILIZANDO A PALAVRA RESERVADA GO MOSTRAMOS PARA O GOLANG QUE ELE NÃO PRECISA ESPERAR A FUNÇÃO ACABAR PARA PROCESSAR A PRÓXIMA
//ASSIM AS DUAS CONCORREM ENTRE SI PARA SEREM EXECUTADAS
//NÃO É NECESSÁRIO COLOCAR A PALAVRA GO EM TODAS AS FUNCS NESSE CONTEXTO POIS SE NÃO O PROJETO NEM EXECUTARÁ
func main()  {
	go escrever("Olá mundo")
	escrever("Olá goroutine")
}