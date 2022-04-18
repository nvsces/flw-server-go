package main

import (
	"log"
	"os"

	trip "github.com/nvsces/flw-server-go"
	"github.com/nvsces/flw-server-go/pkg/handler" // пакет для логирования
	"github.com/nvsces/flw-server-go/pkg/repository"
	"github.com/nvsces/flw-server-go/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

    if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

    	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

    if err!=nil{
        logrus.Fatal("failed to inizialize db: %s", err.Error())
    }

    repos:=repository.NewRepository(db)
    services:= service.NewService(repos)
    handlers:= handler.NewHandler(services)

   srv:=new(trip.Server)
   if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err !=nil {
       log.Fatalf("error accured while running http server : %s", err.Error())
   }
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}