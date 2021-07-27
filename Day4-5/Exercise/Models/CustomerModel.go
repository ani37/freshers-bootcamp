package Models

type Customer struct {
	ID        uint
	FirstName string
	LastName  string
	Phone     string
	Address   string
}

func (b *Customer) TableName() string {
	return "customer"
}
