package Models

type Product struct {
	ID       uint
	Name     string
	Price    int
	Quantity int
}

func (b *Product) TableName() string {
	return "product"
}
