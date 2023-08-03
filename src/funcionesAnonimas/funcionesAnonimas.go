// son funciones que se ejecutan una sola ves
/*package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}() // los parentesis son para llamar la funcion

	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("se va a ejecutar")
		time.Sleep(5 * time.Second)
		fmt.Println("se ejecuto correctamente")
		c <- 1
	}()

	<-c

}*/
