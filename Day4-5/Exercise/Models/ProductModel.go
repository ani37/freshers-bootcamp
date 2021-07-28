package Models

// Product product table with the following defined fields
type Product struct {
	ID       uint
	Name     string `gorm:"not null"`
	Price    int    `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}

func (b *Product) TableName() string {
	return "product"
}
