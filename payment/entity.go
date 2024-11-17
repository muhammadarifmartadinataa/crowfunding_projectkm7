package payment

type Transaction struct {
	ID     int `gorm:"primary_key"`
	Amount int
}
