package pgstore

import (
	"WB_L0/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(o *model.Order) error {
	return r.store.db.QueryRow(
		"INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, "+
			"internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
		o.OrderUid,
		o.TrackNumber,
		o.Entry,
		o.Delivery.Id,
		o.Payment.Id,
		o.Locale,
		o.InternalSignature,
		o.CustomerId,
		o.DeliveryService,
		o.ShardKey,
		o.SmId,
		o.DateCreated,
		o.OofShard,
	).Scan(&o.Id)
}

func (r *OrderRepository) GetOne(uid string) (*model.Order, error) {

	var outOrder model.Order

	rows, err := r.store.db.Query(
		"SELECT * FROM orders"+
			" LEFT JOIN deliveries d ON d.id = orders.delivery_id"+
			" LEFT JOIN payments p ON p.id = orders.payment_id"+
			" INNER JOIN items i ON orders.track_number = i.track_number"+
			" WHERE order_uid = $1",
		uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item model.Item
		var order model.Order
		useless := new(interface{})
		err := rows.Scan(&order.Id, &order.OrderUid, &order.TrackNumber, &order.Entry, useless, useless, &order.Locale,
			&order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.ShardKey, &order.SmId,
			&order.DateCreated, &order.OofShard, &order.Delivery.Id, &order.Delivery.Name, &order.Delivery.Phone,
			&order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
			&order.Delivery.Email, &order.Payment.Id, &order.Payment.Transaction, &order.Payment.RequestId,
			&order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDt,
			&order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee,
			&item.Id, &item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size,
			&item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			return nil, err
		}
		if outOrder.OrderUid == order.OrderUid {
			outOrder.Items = append(outOrder.Items, item)
		} else {
			order.Items = append(order.Items, item)
			outOrder = order
		}
	}

	return &outOrder, nil
}

func (r *OrderRepository) GetCount(uid string) error {
	var id int
	return r.store.db.QueryRow(
		"SELECT id FROM orders WHERE order_uid = $1",
		uid,
	).Scan(&id)
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order

	rows, err := r.store.db.Query(
		"SELECT * FROM orders" +
			" LEFT JOIN deliveries d ON d.id = orders.delivery_id" +
			" LEFT JOIN payments p ON p.id = orders.payment_id" +
			" INNER JOIN items i ON orders.track_number = i.track_number")
	if err != nil {
		return nil, err
	}

	var lastUid string
	for rows.Next() {
		var order model.Order
		var item model.Item
		useless := new(interface{})
		err := rows.Scan(&order.Id, &order.OrderUid, &order.TrackNumber, &order.Entry, useless, useless, &order.Locale,
			&order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.ShardKey, &order.SmId,
			&order.DateCreated, &order.OofShard, &order.Delivery.Id, &order.Delivery.Name, &order.Delivery.Phone,
			&order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
			&order.Delivery.Email, &order.Payment.Id, &order.Payment.Transaction, &order.Payment.RequestId,
			&order.Payment.Currency, &order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDt,
			&order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal, &order.Payment.CustomFee,
			&item.Id, &item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size,
			&item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			return orders, err
		}

		if lastUid == order.OrderUid {
			orders[len(orders)-1].Items = append(orders[len(orders)-1].Items, item)
		} else {
			order.Items = append(order.Items, item)
			orders = append(orders, order)
		}

		lastUid = order.OrderUid
	}

	return orders, nil
}
