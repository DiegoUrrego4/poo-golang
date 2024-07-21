package main

import (
	"composition/pkg/customer"
	"composition/pkg/invoice"
	"composition/pkg/invoiceitem"
	"fmt"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Diego", "Av Cll 100 #12-45", "3220000000"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Leche descremada", 12.34),
			invoiceitem.New(2, "Cadera de res", 54.23),
			invoiceitem.New(3, "Pan tajado", 90.00),
		),
	)

	i.SetTotal()
	fmt.Printf("%+v\n", i)

}
