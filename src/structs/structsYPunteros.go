/*package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) String() string { // personalizar el output de consola deun struct
	return fmt.Sprintf("Tengo %d GB RAM, %d GB disco, y es una %s ", myPc.ram, myPc.disk, myPc.brand)
}

// los punteros son acceso a la memoria
func main() {
	a := 50
	b := &a // es el puntero de a, la direccion de memoria donde esta guardado a

	fmt.Println(a) // imprime 50
	fmt.Println(b) // imprime 0xc000014088

	//para acceder al valor
	fmt.Println(*b) //& es para acceder al espacio de memoria, * para acceder al valor de ese espacio de memoria -  imprime 50

	// si se modifica el valor de una variable, y otra variable a punta a ese espacio de memoria, el valor tambien va a cambiar en ese puntero

	*b = 100
	fmt.Println(a) // imprime 100 , el valor cambio porque tanto a como *b apuntan al mismo espacio de memoria

	//los punteros se usan para llevar funcniones de un paquete a otro de forma mas eficiente

	myPc := pc{ram: 16, disk: 200, brand: "msi"}

	myPc.ping()

	fmt.Println(myPc)

	// una forma de modificar datos a un struct sin punteros, aunque es un metodo que puede ser redundante
	myPc.ram = 60

	// modificando valor de un struct con punteros
	fmt.Println(myPc)   // antes de modificar el struct
	myPc.duplicateRam() // se modifica el struct con punteros
	fmt.Println(myPc)   // despues de modificar el strut

	fmt.Println(myPc)
}*/
