package main

import "fmt"

// Una interfaz es un conjunto de firmas de métodos
// y en GO se fomenta el uso de interfaces simples, a menudo de un solo método
// También las buenas prácticas indican que el nombre de la interfaz deberá ser el nombre de nuestro método
// añadiéndole 'ER' al final

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

// Pero, ¿y si queremos pasar un tipo que cubra varias estructuras?
// Para eso existen las interfaces que combinan interfaces. La única condición es que
// las interfaces que usemos no pueden poseer nombres de métodos disyuntos, es decir con el mismo nombre
// Ejemplo : Greet() en ambas interfaces

type GreeterByer interface {
	Greeter
	Byer
}

// Cualquier tipo válido que implemente el método greet estará implementando la interfaz Greeter
// Ejemplos

type Person struct {
	Name string
}

// Con solo hacer esto, Person ya está implementando la interfaz greeter

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adiós soy %s\n", p.Name)
}

func (p Person) String() string {
	return "Hola, soy una persona y mi nombre es " + p.Name
}

// Recordemos que esto se puede aplicar a CUALQUIER TIPO

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Adiós soy %s\n", t)
}

// Usando la interfaz GreeterByer puedo evitarme usar métodos aparte como los siguientes:
/*
func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
	}
}
*/
/*
func ByeAll(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", b, b)
	}
}
*/

// Y podríamos usar un solo método para evitar repetir código:

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	// Funcionaría con cualquiera de los dos:
	p := Person{Name: "Clara"}
	//t := Text("Daisy")
	// Al ambas cumplir con el contrato de Greeter, las dos variables son aceptadas por el método GreetAll
	//GreetAll(p, t)
	//ByeAll(p, t)
	//All(p, t)

	fmt.Println(p)
}
