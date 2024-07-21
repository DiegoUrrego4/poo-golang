package main

import (
	"fmt"
	"strings"
)

func wrapper(i any) {
	//fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	// Si la interfaz vacía es de tipo string hacer algo
	// ¿Cómo lo podemos hacer? Pues usando el type assertion
	/*
		v, ok := i.(string) // Con esto pregunto si es de tipo string
		if ok {
			fmt.Println(strings.ToUpper(v))
		}
	*/

	// Podemos también usar un switch type:
	switch v := i.(type) {
	case string:
		fmt.Println(strings.ToUpper(v))
	case int:
		fmt.Println("Es int", v*2)
	case float64:
		fmt.Println("Es float64")
	case bool:
		fmt.Println("Es booleano")
	default:
		fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	}
}

func main() {
	wrapper(12)
	wrapper(14.3)
	wrapper(false)
	wrapper("Texto")
	wrapper("DieGo")
	wrapper([]int{1, 2, 3})
}
