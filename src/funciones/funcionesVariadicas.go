//

/*package main

import "fmt"

func suma(values ...int) int { // la funcion toma el parametro como un slice
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// funcion retorno con nombre, puedo colocar variables a los retornos, y go infiere que variables debe devolver sin especificarlas en el return
func getValues(a int) (doble int, triple int, cuadruple int) {
	doble = 2 * a
	triple = 3 * a
	cuadruple = 4 * a
	return

}

func main() {
	fmt.Println(suma(2, 3, 5, 7))
	printNames("Juan", "Pedro", "Luis")
	fmt.Println(getValues(2))
} */
