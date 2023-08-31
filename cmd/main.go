package main

import (
	"log"

	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config: %s", err.Error())
	}

	handler := new(handler.Handler)
	handler.InitRoutes()

	server := new(auth.Server)
	if err := server.Start(viper.GetInt("server.port")); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
