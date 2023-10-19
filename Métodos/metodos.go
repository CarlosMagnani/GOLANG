package main

import "fmt"

///A STRUCT SERIAM AS "CLASSES"
type usuario struct {
	nome  string
	idade int
}

//Dessa forma nossa func estaria grudada na minha struct (classe)
//Os metodos podem retornar valores, receber parametros, funcionando semelhante a uma função
func(U usuario) salvar(){
	//QUANDO EU UTILIZO O U.NOME ESTOU CHAMANDO OS DADOS QUE MINHA STRUCT RECEBEU
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", U.nome)
}

func (U usuario) maiorDeIdade() bool {
	return U.idade >= 18
}

//ASSIM EU ALTERO OS VALORES DO MEU ATRIBUTO, PASSANDO COM PONTEIROS A MINHA STRUCT
func (U *usuario) fazerAniversario() {
	U.idade++
}

func main() {
	usuario1 := usuario{"User 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	maiorIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}