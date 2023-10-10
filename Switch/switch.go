package main

import (
	"fmt"
)

//NÃO É NECESSÁRIO UTILIZAR O BREAK NO GOLANG, ISTO É FEITO POR PADRÃO

//PRIMEIRO TIPO DE SE FAZER UM SWITCH CASE
func diaDaSemana(numero int) string{
	switch numero{
	case 1: 
		return "Domingo"
	case 2: 
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5: 
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7: 
		return "Sabado"
	default:
		return "Numero inválido"	
	} 
}

//SEGUNDO TIPO DE SE FAZER UM SWITCH CASE
func diaDaSemana2(numero int) string  {
	var diaDaSemana string

	switch{
	case numero == 1:
		diaDaSemana = "Domingo"
		//CLÁUSULA RESPONSÁVEL POR RETORNAR O VALOR DA PRÓXIMA CONDICIONAL, "NÃO MUITO UTILIZADO"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero inválido"
	}

	return diaDaSemana
}


func main()  {
	dia := diaDaSemana(4)
	fmt.Println(dia)
}