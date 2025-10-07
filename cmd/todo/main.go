package main

import (
	"fmt"
	"hwproject/internal"
	"hwproject/internal/server"
)

func main() {
	// TODO: Кофнигурация приложения

	fmt.Println("Server started ...")

	cfg := internal.ReadConfig()

	srv := server.NewServer(cfg)

	if err := srv.Run(); err != nil {
		panic(err)
	}

	// TODO: Кофнигурация/создание хранилища

	// TODO: Конфигурация и зпуск веб-сервиса

}
