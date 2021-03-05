package main

import (
	"github.com/taark/crawler/src/crawler"
	"github.com/taark/crawler/src/server"
	"os"
)

func main() {
	// получение port из .env
	port := os.Getenv("PORT")

	//инициализация сервера
	s := server.New(
		server.WithCrawler(crawler.Scan), // передача в сервер обработик-crawler
		server.WithPort(port),            //установка порта
	)

	err := s.Run() //старт сервера
	if err != nil {
		panic(err) //все плохо, сервер падает. Например, если порт занят
	}

}
