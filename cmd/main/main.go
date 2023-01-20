package main

import (
	"L0WB/pkg/handler"
	"L0WB/pkg/repository"
	"L0WB/pkg/service"

	"github.com/spf13/viper"
	_ "github.com/subosito/gotenv"
	"log"
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
	if err != nil {
		log.Println(err)
	}

	repos := repository.NewRepository(db)
	server := service.NewService(repos)
	handle := handler.NewHandler(server)

	server.Cache, err = server.SvControllerData.GetAllOrders()
	handle.InitRouters()
}
