package order

import "gorm.io/gorm"

type OrderInterface interface {
	Create(payload *Order) (Order, error)
}

type OrderRepo struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (or *OrderRepo) Create(payload *Order) (Order, error) {
	err := or.db.Create(&payload).Error

	return *payload, err
}
