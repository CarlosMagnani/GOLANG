package Generator

import (
	"fmt"
	"time"
)

// A FUNÇÃO DE UM PADRÃO GENERATOR É UM FUNÇÃO EM QUE ENCAPSULA UMA GO ROUTINE E DEVOLVE UM CANAL PARA NÓS

//ESTAMOS RETORNANDO UM CANAL DIRECIONAL, DE UMA MÃO SÓ
func Escrever2(texto string) <-chan string {
	canal := make(chan string)

	go func(){
		//AQUI FAZEMOS UM LOOP PARA MANDAR AS MENSAGENS PARA O CANAL
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}();

	
	return canal
}
// A VANTAGEM DE UTILIZAR ESSE PARAMETRÔ É NÃO SER NECESSARIO CHAMAR ELA UTILIZANDO UMA GOROUTINE