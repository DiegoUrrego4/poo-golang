package main

import "fmt"

func main() {
	// De esta forma creamos una instancia de 'Curso'
	Go := Course{
		Name:    "Go desde cero (2023)",
		Price:   12.34,
		IsFree:  false,
		UserIds: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	// Podríamos crear una instancia sin especificar el nombre del campo,
	// pero debe ir en el mismo orden que tiene la estructura
	Java := Course{
		"Java desde cero",
		14.56,
		false,
		[]uint{1, 67, 89, 4},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	// También podríamos crear un curso con solo algunas propiedades, pero en este caso
	// DEBEMOS especificar el nombre del campo o campos que vamos a usar
	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	// Si llegáramos a crear una instancia sin valores en sus propiedades, automáticamente quedarían con su zero value,
	// sin embargo, nosotros podríamos asignarle valores a sus campos fuera de la instanciación
	js := Course{}
	js.Name = "Curso de JavaScript"
	js.UserIds = []uint{56, 7}

	// Y de esta forma extraemos propiedades de nuestra instancia
	fmt.Println(Go.Name)
	fmt.Println(Java.Price)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)

	Go.PrintClasses()
	Go.ChangePrice(69.45)
	fmt.Println(Go.Price)
}
