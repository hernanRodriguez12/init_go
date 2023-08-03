/*package main

import "fmt"

func isPalindromo(texto string) {
	var textReverse string

	for i := len(texto) - 1; i >= 0; i-- {
		textReverse += string(texto[i])
	}

	if texto == textReverse {
		fmt.Println("es palindrome")
	} else {
		fmt.Println("No es palindrome")
	}

}

func main() {

	//ARRAYS -> Son inmutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//SLICE ->  no se le indica la cantidad de valores que va a tener, son mutables
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])   // imprime el elemento 0
	fmt.Println(slice[:3])  // imprime hasta el 3 elemento del slice
	fmt.Println(slice[2:4]) // imprime desde el indice 2, hasta el indice 4, el primer elemento es inclusivo y el segundo elemento es exclusivo
	fmt.Println(slice[4:])  // imprime desde el indice 4 hasta el final del slice

	//APPEND
	slice = append(slice, 7)
	fmt.Println(slice)

	//Agregar una lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // los ... indican que va a agregar los elementos uno por uno
	fmt.Println(slice)

	//RANGE
	slice2 := []string{"Hola", "que", "hace"}

	for i, valor := range slice2 {
		fmt.Println(i, valor)
	}

	for _, valor := range slice2 {
		fmt.Println(valor)
	}

	for i := range slice2 {
		fmt.Println(i)
	}

	// EJERCICIO PALINDROMO
	isPalindromo("amor a roma")
	isPalindromo("amor a romario")
}*/