package main

import (
	"log"

	"github.com/ilgiz-ayupov/auth-app"
	"github.com/ilgiz-ayupov/auth-app/pkg/handler"
	"github.com/ilgiz-ayupov/auth-app/pkg/repository"
	"github.com/ilgiz-ayupov/auth-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация файла настроек
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config: %s", err.Error())
	}

	// Подключение к базе данных
	db, err := repository.OpenSqliteDB(&repository.Config{
		Path: viper.GetString("sqlite.path"),
	})
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}
	defer db.Close()

	// Запуск сервера
	repos := repository.InitRepository(db)
	services := service.InitService(repos)
	handlers := handler.InitHandler(services)

	server := new(auth.Server)
	if err := server.Start(viper.GetInt("server.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
