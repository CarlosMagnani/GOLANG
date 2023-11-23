package main

import (
	"fmt"
	"teste/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Cecilia de Assis silva")

	fmt.Println(tipoEndereco)
}