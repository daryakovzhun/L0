package main

import (
	"L0WB/internal/domain"
	"L0WB/pkg/handler"
	"L0WB/pkg/repository"
	"L0WB/pkg/service"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	_ "github.com/subosito/gotenv"
	"io"
	"log"
	"math/rand"
	"time"
)

func init() {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	//if err := gotenv.Load(); err != nil {
	//	log.Println(err)
	//}
}

func main() {
	cfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := cfg.ConnectDB()
	checkFail("Connect DataBase", err)

	repos := repository.NewRepository(db)
	server := service.NewService(repos)
	handle := handler.NewHandler(server)

	ords, err := server.SvControllerData.GetAllOrders()
	checkFail("GetAllOrders Cache", err)
	server.Cache.SetFew(ords, 15*time.Minute)

	go startNats(server)

	handle.InitRouters()
}

func checkFail(funcname string, err error) {
	if err != nil {
		log.Println(err)
	} else {
		log.Println(funcname + ": OK")
	}
}

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}

func SubscriberNats(s *service.Service, conn stan.Conn) {
	var err error

	_, err = conn.Subscribe("NewOrder", func(msg *stan.Msg) {

		var ord domain.Order
		if err = json.Unmarshal(msg.Data, &ord); err != nil {
			log.Println(err)
			return
		}
		checkFail("InsertOrder", s.SvControllerData.InsertOrder(&ord))
		//s.Cache[ord.Order_uid] = ord
		s.Cache.Set(ord.Order_uid, ord, time.Minute)

		fmt.Printf("seq = %d [redelivered = %v] mes= %s \n", msg.Sequence, msg.Redelivered, msg.Data)

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		msg.Ack()

	}, stan.DurableName("i-will-remember"), stan.MaxInflight(100), stan.SetManualAckMode())

	if err != nil {
		log.Println(err)
	}
}

func startNats(s *service.Service) {
	if err := runNats(s); err != nil {
		log.Println(err)
	}
}

func runNats(s *service.Service) error {
	conn, err := stan.Connect(
		viper.GetString("natStreaming.clusterId"),
		viper.GetString("natStreaming.clientId"),
	)
	checkFail("Connect NATS Streaming", err)
	defer logCloser(conn)

	done := make(chan struct{})
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	SubscriberNats(s, conn)
	<-done

	return nil
}
