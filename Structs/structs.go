package main

import "fmt"

//MODELO DE USUARIO -- CLASSE
type usuario struct {
	nome string
	idade int8
	enderecoUsuario endereco
}

type endereco struct {
	bairro string
	cep int
}
func main()  {
	
	//PRIMEIRA FORMA DE ATRIBUIR VALOR
	var u usuario 
	u.nome = "Carlos"
	u.idade = 19
	fmt.Println(u)
	example := endereco{"Gramadão", 13211883}
	//SEGUNDA FORMA DE ATRIBUIR VALOR
	user := usuario{"Carlos", 19, example}
	fmt.Println(user)

	//TERCEIRA FORMA DE ATRIBUIR VALORES
	user2 := usuario{nome:"Davi"}
	fmt.Println(user2)
}