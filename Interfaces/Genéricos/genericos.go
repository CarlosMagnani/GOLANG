package main

import "fmt"

//DESSA FORMA DEFININOS OS TIPOS GENERICOS
//RESUMIDAMENTE É UM "ANY" DO TS
func genericFunction(interf interface{}) {
	fmt.Println(interf)
}
func main() {
	genericFunction("string")
	genericFunction(1)
	genericFunction(true)

	fmt.Println(1, 2, "string", false, true, float64(12345))


	//DESSA FORMA O MAP ACEITA QUALQUER COISA
	mapa := map[interface {}]interface{}{
		1: "string",
		float64(100): true,
		"String": "String",
		true: float64(12),
	}

	fmt.Println(mapa)
}