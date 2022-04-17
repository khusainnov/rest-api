package main

import (
	"github.com/joho/godotenv"
	"github.com/khusainnov/rest-api"
	"github.com/khusainnov/rest-api/package/handler"
	"github.com/khusainnov/rest-api/package/repository"
	"github.com/khusainnov/rest-api/package/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Cannot read config due to error: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil{
		log.Fatalf("Cannot load env variables due to error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Cannot connect to DB due to error: %s", err.Error())
	}

	repos := repository.NewService(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := restapi.Server{}

	if err = s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Cannot run the server due to error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}