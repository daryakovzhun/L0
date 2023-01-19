package controller

import (
	mod "L0/internal/domain"
	"github.com/jmoiron/sqlx"
)

type DbController struct {
	db *sqlx.DB
}

func NewDbController(db *sqlx.DB) *DbController {
	return &DbController{db}
}

func (DC *DbController) GetDelivery(order_uid string, del *mod.Delivery) (err error) {
	row := DC.db.QueryRowx(`SELECT name, phone, zip, city, address, region, email
	FROM delivery WHERE order_uid = $1;`, order_uid)

	err = row.StructScan(del)
	if err != nil {
		return err
	}
	return nil
}

func (DC *DbController) GetPayment(order_uid string, pay *mod.Payment) (err error) {
	row := DC.db.QueryRowx(`SELECT transaction, request_id, currency, provider, amount,
 	payment_dt, bank, delivery_cost, goods_total, custom_fee
	FROM payment WHERE order_uid = $1;`, order_uid)

	err = row.StructScan(pay)
	if err != nil {
		return err
	}
	return nil
}

func (DC *DbController) GetItems(order_uid string, itms *mod.Items) (err error) {
	rows, err := DC.db.Queryx(`SELECT chrt_id, track_number, price, rid, name,
    sale, size, total_price, nm_id, brand, status
	FROM items WHERE order_uid = $1;`, order_uid)

	if err != nil {
		return err
	}

	var itm mod.Item
	for rows.Next() {
		err = rows.StructScan(&itm)
		if err != nil {
			return err
		}
		*itms = append(*itms, itm)
	}

	return nil
}

func (DC *DbController) GetOrder(order_uid string) (mod.Order, error) {
	row := DC.db.QueryRowx(`order_uid, track_number, entry, locale,
    internal_signature, customer_id, delivery_service, shardkey, sm_id,
    date_created, oof_shard
    FROM orders WHERE order_uid = $1;`, order_uid)

	var ord mod.Order
	err := row.StructScan(&ord)

	err = DC.GetDelivery(order_uid, &ord.Delivery)
	if err != nil {
		//		return nil, err   TO DO
	}

	err = DC.GetPayment(order_uid, &ord.Payment)
	if err != nil {
		//		return nil, err   TO DO
	}

	err = DC.GetItems(order_uid, &ord.Items)
	if err != nil {
		//		return nil, err   TO DO
	}

	return ord, nil
}
