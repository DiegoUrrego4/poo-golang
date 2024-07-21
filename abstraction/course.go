package main

import "fmt"

// Las estructuras en GO son como las 'clases' en lenguajes como Java y C#

// Course sería la representación de una 'clase', pero acá son 'colecciones de campos'
type Course struct {
	// atributos
	// Si inicia con mayúscula, son importados, si no, no lo son
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
