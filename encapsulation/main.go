package main

import (
	"encapsulation/course"
	"fmt"
)

func main() {
	// Función constructora
	Go := course.New("Go desde cero", 0, false)
	Go.SetUserIds([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("POO con GO")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	fmt.Printf("%+v", Go)
}
