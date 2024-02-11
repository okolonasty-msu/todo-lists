package main

import (
	"awesomeProject/pkg/handler"
	"github.com/spf13/viper"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	router := handlers.InitRouters()
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
	//
	//srv := new(awesome.Server)
	//if err := srv.Run("8000", handlers.InitRouters()); err != nil {
	//	log.Fatalf("error occured while running server: %s", err.Error())
	//
	//}
}

func initCongig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
