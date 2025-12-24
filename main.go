package main

import (
	"http-go/src"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[HTTP] ", log.LstdFlags)

	server := src.NewServer(src.ServerConfig{
		Address:      ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Logger:       logger,
	})

	// // Register routes
	// server.Router.GET("/", func(req *src.HttpRequest) *src.HttpResponse {
	// 	return src.NewHttpResponse(200, "Hello, World!")
	// })

	// server.Router.GET("/health", func(req *src.HttpRequest) *src.HttpResponse {
	// 	return src.NewHttpResponse(200, `{"status": "ok"}`)
	// })

	server.ListenConnection()
}
