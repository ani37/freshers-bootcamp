package Models

// Customer Customers table with the following defined fields
type Customer struct {
	ID        uint
	FirstName string `gorm:"not null"`
	LastName  string
	Phone     string
	Address   string
}

func (b *Customer) TableName() string {
	return "customer"
}
