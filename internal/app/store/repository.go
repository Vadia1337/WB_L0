package store

import (
	"WB_L0/internal/app/model"
)

type OrderRepository interface {
	Create(*model.Order) error
	GetCount(uid string) error
	GetOne(uid string) (*model.Order, error)
	GetAll() ([]model.Order, error)
}

type DeliveryRepository interface {
	Create(*model.Delivery) error
}

type PaymentRepository interface {
	Create(*model.Payment) error
}

type ItemRepository interface {
	Create(*model.Item) error
}
