package course

import "fmt"

// Los identificadores exportados son todos los que estén con su inicial en letra mayúscula

type course struct {
	name    string
	price   float64
	isFree  bool
	userIds []uint
	classes map[uint]string
}

// Go NO tiene constructores, pero se puede hacer de la siguiente manera

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// Getters y setters

func (c *course) SetName(name string) {
	c.name = name
}

func (c *course) Name() string {
	return c.name
}

func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}

func (c *course) IsFree() bool {
	return c.isFree
}

func (c *course) SetUserIds(userIds []uint) {
	c.userIds = userIds
}

func (c *course) UserIds() []uint {
	return c.userIds
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) Classes() map[uint]string {
	return c.classes
}
