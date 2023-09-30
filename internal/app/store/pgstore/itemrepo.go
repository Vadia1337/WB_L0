package pgstore

import "WB_L0/internal/app/model"

type ItemRepository struct {
	store *Store
}

func (r *ItemRepository) Create(i *model.Item) error {
	return r.store.db.QueryRow(
		"INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, "+
			"total_price, nm_id, brand, status)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		i.ChrtId,
		i.TrackNumber,
		i.Price,
		i.Rid,
		i.Name,
		i.Sale,
		i.Size,
		i.TotalPrice,
		i.NmId,
		i.Brand,
		i.Status,
	).Scan(&i.Id)
}
