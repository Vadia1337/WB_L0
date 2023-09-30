package store

type Store interface {
	Order() OrderRepository
	Delivery() DeliveryRepository
	Payment() PaymentRepository
	Item() ItemRepository
}
