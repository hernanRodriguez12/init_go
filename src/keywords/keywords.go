/*package main

import "fmt"

func main() {
	//DEFER -> ejecuta la ultima funcion antes de que la ejecucion termine, lo ideal es solo usar un defer por funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//CONTINUE Y BREAK ->  Se usan dentro del ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("es 2")
			continue // se usa para que continue una ejecucion
		}

		if i == 6 {
			fmt.Println("es 6")
			break // se usa para que INTERRUMPA una ejecucion
		}
	}
}*/
