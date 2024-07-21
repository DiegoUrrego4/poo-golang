package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Printf("¡Hello %s!\n", p.Name)
}

// Hay que tener cuidado, pues se pueden presentar conflictos como en el siguiente ejemplo

// Human cuenta con un campo age al igual que Person, lo que da espacio a que go se confunda
type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age uint, children uint) Human {
	return Human{age, children}
}

type Employee struct {
	// Con esto estamos embebiendo a persona dentro de empleado
	// Los campos y métodos del tipo interno ahora harán parte del tipo externo
	Person
	// Al añadir a human estamos expuestos a que go se confunda al compartir el campo 'age'
	Human
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

// En Go NO existe la sobre escritura de métodos

// Greet Acá realmente estamos usando una ocultación de métodos
func (e Employee) Greet() {
	fmt.Printf("¡Hello %s desde empleado!\n", e.Name)
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	p := NewPerson("Diego", 28)
	p.Greet()
	fmt.Println()

	e := NewEmployee("Alejandro", 30, 3, 1000000)
	fmt.Println(e.Name, e.Person.Age) // Acá tendríamos que especificar si es la edad de la persona o del humano
	e.Person.Greet()
	e.Greet()
	e.Payroll()
}
