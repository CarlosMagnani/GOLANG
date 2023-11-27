package main

import (
	"fmt"
	"teste/enderecos"
)

func teste() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Cecilia de Assis silva")

	fmt.Println(tipoEndereco)
}