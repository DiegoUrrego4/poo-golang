package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Pagando con paypal…")
}

type CreditCard struct{}

func (cc CreditCard) Pay() {
	fmt.Println("Pagando con tarjeta de crédito…")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagando con efectivo…")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digíte un número de método de pago:")
	fmt.Println("\t 1: Paypal 2: Efectivo 3: Tarjeta de crédito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Debe digitar un método válido")
	}

	if method > 3 {
		panic("Debe digitar un método válido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
