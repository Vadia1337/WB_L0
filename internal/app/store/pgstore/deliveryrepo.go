package pgstore

import "WB_L0/internal/app/model"

type DeliveryRepository struct {
	store *Store
}

func (r *DeliveryRepository) Create(d *model.Delivery) error {
	return r.store.db.QueryRow(
		"INSERT INTO deliveries (name, phone, zip, city, address, region, email)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	).Scan(&d.Id)
}
