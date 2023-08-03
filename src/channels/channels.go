/*package main

import "fmt"

// go routines no permiten enviar info entre rutinas

func say(text string, c chan<- string) { //se debe indicar si el parametro del channel es de entrada o de salida, en la izquierda para salida, en la derecha para entrada
	c <- text // el simbolo <- es para indicar que se va a agregar un dato en ese canal

}
func main() {
	c := make(chan string, 1)

	fmt.Println("Hola")
	go say("Mundo", c)

	fmt.Println(<-c)
}*/
