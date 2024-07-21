package invoice

import (
	"composition/pkg/customer"
	"composition/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer // composición de 'uno a uno'
	items   invoiceitem.Items // composición de 'uno a muchos'
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
