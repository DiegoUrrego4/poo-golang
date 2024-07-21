package invoiceitem

// Item contains information of an invoice item
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new item
func New(id uint, product string, value float64) Item {
	return Item{
		id:      id,
		product: product,
		value:   value,
	}
}

// Value getter of Item.value
//func (i Item) Value() float64 {
//	return i.value
//}

type Items []Item // Acá creamos un nuevo tipo

func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}
