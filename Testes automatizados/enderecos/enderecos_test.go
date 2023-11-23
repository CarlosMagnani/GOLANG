package enderecos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type cenariosDeTeste struct {
	enderecoInserido string
	expectedEndereco string
}

func TestTipoDeEndereco(t *testing.T) {


	cenariosDeTeste := []cenariosDeTeste {
		{"Rua abc", "RUA"},
		{"Avenida das Rosas", "AVENIDA"},
		{"Estrada dos Imigrantes", "ESTRADA"},
		// {"LALALAL", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		receiveEndereco := TipoDeEndereco(cenario.enderecoInserido)
		assert.Equal(t, cenario.expectedEndereco, receiveEndereco)
	}
}