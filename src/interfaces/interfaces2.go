// Go no implementa interfaces de manera explicita, no se usa ninguna palabrra reservada
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
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "es un full time Employee"
}

func (tEmployee TemporallyEmployee) getMessage() string {
	return "es un temporally Employee"
}

type TemporallyEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pedro"
	ftEmployee.age = 25
	ftEmployee.id = 5
	GetMessage(ftEmployee)

	tEmployee := TemporallyEmployee{}
	tEmployee.name = "Juan"
	tEmployee.age = 25
	tEmployee.id = 5
	tEmployee.taxRate = 10
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(tEmployee)
	//GetMessage(ftEmployee) el polimorfismo no funciona en go
}*/
