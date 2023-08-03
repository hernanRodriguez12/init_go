/*package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// CONDICIONAL IF
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("es 2")
	}

	//AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	//OR
	if valor1 == 3 || valor2 == 2 {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	//CONVERTIR TEXTO A NUMERO

	value, err := strconv.Atoi("57")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("no hay error: ", value)
	}

	// CONDICIONAL SWITCH

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("es impar")
	}

	//SWITCH SIN CONDICION
	valueSwitch := 200
	switch {
	case valueSwitch > 100:
		fmt.Println("es mayor")
	case value < 0:
		fmt.Println("es menor a 0")
	default:
		fmt.Println("no se cumple ningun case")
	}
}*/
