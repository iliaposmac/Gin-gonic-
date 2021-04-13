package main

import (
	"os"

	"github.com/iliaposmac/todo-app"
	"github.com/iliaposmac/todo-app/pkg/handler"
	"github.com/iliaposmac/todo-app/pkg/repository"
	"github.com/iliaposmac/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter((new(logrus.JSONFormatter)))

	if err := initConfig(); err != nil {
		logrus.Fatalln("Can not read config file", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalln("Can not get env variables", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		logrus.Fatalln("Can not connect to datanase", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHanlder(services)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to run server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
