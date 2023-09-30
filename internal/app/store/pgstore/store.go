package pgstore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                 *sql.DB
	OrderRepository    *OrderRepository
	DeliveryRepository *DeliveryRepository
	PaymentRepository  *PaymentRepository
	ItemRepository     *ItemRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Order() *OrderRepository {
	if s.OrderRepository != nil {
		return s.OrderRepository
	}

	s.OrderRepository = &OrderRepository{
		store: s,
	}

	return s.OrderRepository
}

func (s *Store) Delivery() *DeliveryRepository {
	if s.DeliveryRepository != nil {
		return s.DeliveryRepository
	}

	s.DeliveryRepository = &DeliveryRepository{
		store: s,
	}

	return s.DeliveryRepository
}

func (s *Store) Payment() *PaymentRepository {
	if s.PaymentRepository != nil {
		return s.PaymentRepository
	}

	s.PaymentRepository = &PaymentRepository{
		store: s,
	}

	return s.PaymentRepository
}

func (s *Store) Item() *ItemRepository {
	if s.ItemRepository != nil {
		return s.ItemRepository
	}

	s.ItemRepository = &ItemRepository{
		store: s,
	}

	return s.ItemRepository
}
