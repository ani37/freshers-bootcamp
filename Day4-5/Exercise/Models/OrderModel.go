package Models

type Order struct {
	ID         uint
	CustomerID uint
	ProductID  uint
	Quantity   uint
	Status     string
	Customer   Customer `gorm:"foreignKey:customer_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product    Product  `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *Order) TableName() string {
	return "order"
}
