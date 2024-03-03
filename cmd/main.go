package main

import (
	awesome "awesomeProject"
	"awesomeProject/pkg/handler"
	"awesomeProject/pkg/repository"
	"awesomeProject/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initCongig(); err != nil {
		log.Fatalf("error initializing config")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(awesome.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initCongig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
