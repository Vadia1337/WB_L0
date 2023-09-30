package pgstore

import "WB_L0/internal/app/model"

type PaymentRepository struct {
	store *Store
}

func (r *PaymentRepository) Create(p *model.Payment) error {
	return r.store.db.QueryRow(
		"INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, "+
			"bank, delivery_cost, goods_total, custom_fee)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		p.Transaction,
		p.RequestId,
		p.Currency,
		p.Provider,
		p.Amount,
		p.PaymentDt,
		p.Bank,
		p.DeliveryCost,
		p.GoodsTotal,
		p.CustomFee,
	).Scan(&p.Id)
}
