package main

import (
	"concorrencia/generator"
	"fmt"
)
func main() {
	
	canal :=  multiplex(generator.Escrever2("OLÁ MUNDO"), generator.Escrever2("OLÁ MUNDO, ESTOU APRENDENDO GOLANG"))
	for i := 0; i < 10; i++ {
		fmt.Printf(<-canal)
	}
}

// EU UTILIZO TAMBÉM O CONCEITO DE GENERATOR DENTOR DO MULTIPEXADOR PARA FAZER A CRIAÇÃO DE CANAIS E
// GOROUTINES
func multiplex(canalDeEntrada, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada:
				canalDeSaida <- mensagem

			case mensagem2 := <-canalDeEntrada2:
				canalDeSaida <- mensagem2
			}
		}
	}()

	return canalDeSaida
}
