/*package main

import (
	"fmt"
)

func message(text string, canal chan string) {
	canal <- text
}

func main() {
	c := make(chan string, 2)

	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c)) // len -> cuantas go routines hay en el chanel, cap -> cuanta es la cantifda maxima de ese channel

	//CLOSE
	close(c) // va a cerrar el canal -  cerrar el canal cuando se va a dejar de usar

	//c <- "mensaje 3" // esto da error porque el channel esta lleno

	//RANGE - recorrer el contenido de un channel

	for message := range c {
		fmt.Println(message)
	}

	//SELECT
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}*/
