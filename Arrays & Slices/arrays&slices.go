package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ARRAYS

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1,2,3,4,5}
	fmt.Println(array2)

	//ARRAT SEM TAMANHO FIXO
	array3 := [...]int{1,2,3,4,5,6,7,8}
	fmt.Println(array3)
	//FIM ARRAY/////////////////////////


	//SLICES
	slice := []int{10,11,12,13,14,15}
	fmt.Println(slice)

	//ADIÇÃO DE VALORES DENTRO DO SLICE
	slice = append(slice, 18)

	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array3), slice)




	//ARRAYS INTERNOS
	//OS SLICES AUMENTA A CAPACIDADE AO ESTOURAR A CAPACIDADE MÁXIMA SE ADAPTANDO AO TAMANHO RECEBIDO. SENDO ASSIM INFINTO.
	//AO NÃO PASSAR A CAPACIDADE ELE DEFINE COMO TAMANHO INICIAL O VALOR DE LEN E AO ESTOURAR ELE DOBRA ESSE VALOR.
	fmt.Println("------------------")
	slice3 := make([]int, 10, 15)
	fmt.Println(len(slice3), cap(slice3))
}