package order

import "time"

const (
	Pending string = "pending"
	Cancel  string = "cancel"
	Sucess  string = "success"
)

type Order struct {
	ID          uint `gorm:"primaryKey" json:"order_id"`
	ServiceCode string
	AccountId   string
	AccountZone string
	TotalPrice  uint
	Price       uint
	Fee         uint
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Order) TableName() string {
	return "orders"
}
