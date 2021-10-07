package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARRAY")

	// 1- ARRAY
	var array1 [5]string //array com 5 posições e todos devem ser do mesmo tipo
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//2 - ARRAY
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// 3 - ARRAY
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// 1 - SLICE
	fmt.Println("**********************************")
	fmt.Println("SLICE")

	slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(slice)

	slice = append(slice, 70) // ACRECENTA ALGO NO SLICE
	fmt.Println(slice)

	// 2 - SLICE

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// 1 - ARRAYS INTERNOS

	fmt.Println("************************")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(len(slice3)) //TAMANHO
	fmt.Println(cap(slice3)) //CAPACIDADE

	// 2 SLICE

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
