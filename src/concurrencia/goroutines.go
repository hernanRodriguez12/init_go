/*package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() { // la funcion main corre dentro de una goroutine
	//say("Hola")
	//go say("Mundo") // para que una funcion se ejecute de forma concurrente se agrega la palabra go

	//forma no eficiente -
	//time.Sleep(time.Second * 1)

	//otra alternativa
	var wg sync.WaitGroup // acumula un conjunto de go routines
	fmt.Println("Hola")
	// se indica que debe agregar una go routine al ciclo de waitGroup
	wg.Add(1) // se agrega una go routine al group
	go say("Mundo", &wg)
	wg.Wait() // se le dice a la go routine principal que espera a que se ejecuten las rutines internas


	// las go routines se usan mucho con funciones anonimas
}*/
