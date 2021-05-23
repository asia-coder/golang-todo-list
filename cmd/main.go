package main

import (
	"github.com/asiaCoder/todo-app"
	"github.com/asiaCoder/todo-app/pkg/handler"
	"github.com/asiaCoder/todo-app/pkg/repository"
	"github.com/asiaCoder/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка при чтение конфигурационного файла: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при чтение переменных окружения: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err = srv.Run(viper.GetString("http.port"), handlers.InitRoutes())

	if err != nil {
		log.Fatalf("Ошибка при работе http сервера: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
