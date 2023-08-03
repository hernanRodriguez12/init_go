// Engo no hay herencia, existe la composicion
/*package main

import "fmt"

type Person struct {
	name string
	age  uint
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person   // tiene propiedades de person - campo anonimo, solo se pone el tipo de struct
	Employee // tiene propiedades de Employee - campo anonimo
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {

	//
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pedro"
	ftEmployee.age = 25
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee) el polimorfismo no funciona en go
}*/
