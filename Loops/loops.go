package main

import (
	"fmt"
	"time"
)

func main() {

	//FOR BEM PARECIDO COM O "WHILE..." VAI SER EXECUTADO ENQUANTO A CONDICIONAL NÃO FOR ATENDIDA
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
		// time.Sleep(time.Second)
	}

	//FOR CONVENCIONAL
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		j++
		
		time.Sleep(time.Second)
	}


	//FOR PARECIDO COM O FOROF, FOREACH. UM FOR DE ITERAÇÃO
	nomes := []string{"João", "Lucas", "Claudio"}

	for index, nome := range nomes{
		fmt.Println(index, nome)
	}

	//ITERANDO SOBRE PALAVRAS
	for index, letra := range "PALAVRA"{
		fmt.Println(index, letra)
	}


	// ITERANDO POR "ARRAY DE OBJETOS"
	usuario := map[string]string {
		"nome": "Carlos",
		"sobrenome": "Magnani",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//LOOP INFINITO
	for{

	}
}