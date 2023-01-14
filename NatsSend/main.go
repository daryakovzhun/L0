package main

import (
	"encoding/json"

	mod "./../Models"
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "2"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		panic(err.Error())
	}

	data := mod.Order{
		Order_uid:    "b563feb7b2b84b6test",
		Track_number: "WBILMTESTTRACK",
		Entry:        "WBIL",
		Delivery: mod.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com"},
		Payment: mod.Payment{
			Transaction:   "b563feb7b2b84b6test",
			Request_id:    "",
			Currency:      "USD",
			Provider:      "wbpay",
			Amount:        1817,
			Payment_dt:    1637907727,
			Bank:          "alpha",
			Delivery_cost: 1500,
			Goods_total:   317,
			Custom_fee:    0},
		Items: mod.Items{
			{
				Chrt_id:      9934930,
				Track_number: "WBILMTESTTRACK",
				Price:        453,
				Rid:          "ab4219087a764ae0btest",
				Name:         "Mascaras",
				Sale:         30,
				Size:         "0",
				Total_price:  317,
				Nm_id:        2389212,
				Brand:        "Vivienne Sabo",
				Status:       202},
			{
				Chrt_id:      9934930,
				Track_number: "WBILMTESTTRACK",
				Price:        505,
				Rid:          "ab4219087a764ae0btest11",
				Name:         "Simba",
				Sale:         20,
				Size:         "1",
				Total_price:  1235,
				Nm_id:        2389212,
				Brand:        "Vivienne Sabo",
				Status:       202},
		},
		Locale:             "en",
		Internal_signature: "",
		Customer_id:        "test",
		Delivery_service:   "meest",
		Shardkey:           "9",
		Sm_id:              99,
		Date_created:       "2021-11-26T06:22:19Z",
		Oof_shard:          "1"}

	order, err := json.Marshal(&data)
	if err != nil {
		panic(err.Error())
	}
	// Simple Synchronous Publisher
	if err := sc.Publish("Order", order); err != nil {
		panic(err.Error())
	} // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	// sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// })

	// Unsubscribe
	// sub.Unsubscribe()

	// Close connection
	sc.Close()
}
