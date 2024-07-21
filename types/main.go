package main

import "fmt"

// course estructura de cursos
type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// Declaraciones de alias ⇾ nos permite crear un identificador que hace referencia a un tipo existente, heredando sus métodos
type myAlias = course // El tipo real de mi alias es tipo course

// Definición de tipo ⇾ Acá obtenemos el conjunto de campos de mi tipo base, pero no heredamos el conjunto de métodos de mi nuevo tipo
type newCourse course

// Acá creando un nuevo tipo podemos añadir métodos
type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	//c := newCourse{name: "Go"}
	//c.Print()
	//fmt.Printf("El tipo es: %T\n", c)
	var b newBool = false // El nuevo tipo almacena los mismos valores que el valor original
	fmt.Println(b.String())
}
