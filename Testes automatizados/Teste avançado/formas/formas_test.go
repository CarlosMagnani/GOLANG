package test

import "testing"

func TestArea(t *testing.T){
	t.Run("Quadrado", func(t *testing.T)){
		qua := Quadrado{10, 12}
		areaEsperada := float64(120)
		areaRecebida := qua.Area()

		assert.Equal(t, areaEsperada, areaRecebida)
	}


	t.Run("Circulo", func(t *testing.T)){

	}
}