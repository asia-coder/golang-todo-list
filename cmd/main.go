package main

import (
	"github.com/asiaCoder/todo-app"
	"github.com/asiaCoder/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	err := srv.Run("8080", handlers.InitRoutes())

	if err != nil {
		log.Fatalf("Ошибка при работе http сервера: %s", err.Error())
	}

}
