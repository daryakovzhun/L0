package controller

import (
	mod "L0WB/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	row := DC.db.QueryRowx(`SELECT order_uid, track_number, entry, locale,
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

func (DC *DbController) GetAllOrders() (map[string]mod.Order, error) {
	rows, err := DC.db.Queryx(`SELECT order_uid FROM orders;`)
	if err != nil {
		return nil, err
	}

	ords := make(map[string]mod.Order)
	var order_uid string
	for rows.Next() {
		err := rows.Scan(&order_uid)
		if err != nil {
			return nil, err
		}
		ord, err := DC.GetOrder(order_uid)
		if err != nil {
			return nil, err
		}

		ords[order_uid] = ord
	}

	return ords, nil
}

func (DC *DbController) InsertDelivery(order_uid string, del *mod.Delivery) (err error) {
	query := fmt.Sprintf(`INSERT INTO delivery(
			order_uid, name, phone, zip, city, address, region, email)
			VALUES($1, :name, :phone, :zip, :city, :address, :region, :email);`, order_uid)
	_, err = DC.db.NamedExec(query, del)
	if err != nil {
		return err
	}
	return nil
}

func (DC *DbController) InsertPayment(order_uid string, pay *mod.Payment) (err error) {
	query := fmt.Sprintf(`INSERT INTO payment(
			order_uid, transaction, request_id, currency, provider, amount, 
            payment_dt, bank, delivery_cost, goods_total, custom_fee)
			VALUES($1, :transaction, :request_id, :currency, :provider,:amount,
			       :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee);`, order_uid)
	_, err = DC.db.NamedExec(query, pay)
	if err != nil {
		return err
	}
	return nil
}

func (DC *DbController) InsertItem(order_uid string, itm *mod.Item) (err error) {
	query := fmt.Sprintf(`INSERT INTO items(
			order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
			VALUES($1,:chrt_id, :track_number, :price, :rid, :name, :sale, :size, :total_price, :nm_id, :brand, :status);`, order_uid)
	_, err = DC.db.NamedExec(query, itm)
	if err != nil {
		return err
	}
	return nil
}

func (DC *DbController) InsertOrder(ord *mod.Order) (err error) {
	query := fmt.Sprintf(`INSERT INTO orders(
			order_uid, track_number, entry, locale, internal_signature, 
            customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			VALUES(:order_uid, :track_number, :entry, :locale, :internal_signature, :customer_id, 
			:delivery_service, :shardkey, :sm_id, :date_created, :oof_shard);`)
	_, err = DC.db.NamedExec(query, ord)
	if err != nil {
		return err
	}

	err = DC.InsertDelivery(ord.Order_uid, &ord.Delivery)
	if err != nil {
		return err
	}

	err = DC.InsertPayment(ord.Order_uid, &ord.Payment)
	if err != nil {
		return err
	}

	for _, itm := range ord.Items {
		err = DC.InsertItem(ord.Order_uid, &itm)
		if err != nil {
			return err
		}
	}

	return nil
}
