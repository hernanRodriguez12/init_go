// permiten intanciar las clases y convertirlas en objetos
/*package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {

	// forma 1 de crear un constructor de un struct
	e := Employee{}
	fmt.Printf("%v\n", e)

	// forma 2 de crear un constructor de un struct
	e2 := Employee{id: 1, name: "Juan", vacation: true}
	fmt.Printf("%v\n", e2)

	// forma 3 de crear un constructor de un struct
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 4
	e3.name = "Pedro"

	fmt.Printf("%v\n", *e3)

	// forma 4 puede ser la mas recomendada
	e4 := NewEmployee(2, "Luis", true)
	fmt.Printf("%v\n", *e4)

}*/
