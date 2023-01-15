package db

import (
	"database/sql"
	"errors"
	"fmt"

	mod "Models"
	_ "github.com/lib/pq"
)

type DataBase struct {
	User      string
	Password  string
	Dbname    string
	Sslmode   string
	_executer *sql.DB
}

func (Db *DataBase) Connect() (err error) {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", Db.User, Db.Password, Db.Dbname, Db.Sslmode)
	Db._executer, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return nil
}

func (Db *DataBase) GetOrderIdByOrder(ord mod.Order, del_id int, pay_id int) (int, error) {
	Row, err := Db._executer.QueryRow(`SELECT id FROM Orders
	Where order_uid = $1
	And track_number = $2
	And entry = $3
	And delivery_id = $4
	And payment_id =  $5
	And locale = $6
	And internal_signature = $7
	And customer_id = $8
	And delivery_service = $9
	And shardkey = $10
	And sm_id = $11
	And oof_shard = $12
	And date_created = $13;`,
		ord.Order_uid,
		ord.Track_number,
		ord.Entry,
		del_id,
		pay_id,
		ord.Locale,
		ord.Internal_signature,
		ord.Customer_id,
		ord.Delivery_service,
		ord.Shardkey,
		ord.Sm_id,
		ord.Oof_shard,
		ord.Date_created)

	var id int
	if Row.Next() {
		err := Row.Scan(&id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("Null Select")
}
