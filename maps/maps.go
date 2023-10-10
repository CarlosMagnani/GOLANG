package main

import "fmt"

func main() {

	//DENTRO DO [TIPO CHAVE]TIPO VALOR
	//TODOS TEM QUE SER DO MESMO TIPO SE NÃO DA PROBLEMA
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",  
	}
	fmt.Println(usuario["nome"])

	//ANINHANDO OS MAPS ('OBJETOS')
	usuario2 := map[string]map[string]string {
		"pessoa": {
			"nome": "Carlos",
			"sobrenome": "Magnani",
		},
	}
	fmt.Println(usuario2)

}