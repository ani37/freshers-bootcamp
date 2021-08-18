package Models

// Order order table with the following defined fields
type Order struct {
	ID         uint
	CustomerID uint `gorm:"not null"`
	ProductID  uint `gorm:"not null"`
	Quantity   int  `gorm:"not null"`
	Status     string
	Customer   Customer `gorm:"foreignKey:customer_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product    Product  `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *Order) TableName() string {
	return "order"
}
