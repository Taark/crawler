package main

import (
	"github.com/taark/crawler/src/crawler"
	"github.com/taark/crawler/src/server"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	s := server.New(
		server.WithCrawler(crawler.Scan),
		server.WithPort(port),
	)

	err := s.Run()
	if err != nil {
		panic(err)
	}

}
