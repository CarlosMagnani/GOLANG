package enderecos

import "strings"

//A função verifica o tipo de endereço enviado
func TipoDeEndereco(end string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	//O SPLIT FAZ VIRAR UM SLICE SEM QUE TIVER UM ESPAÇO
	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(end), " ")[0]


	endValid := false 
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			endValid = true
		}
	}

	if endValid {
		return strings.ToTitle(primeiraPalavraDoEndereco) 
	}

	return "Tipo Invalido"
}