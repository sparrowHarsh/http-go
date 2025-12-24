package main

import (
<<<<<<< HEAD
	"http-go/src"
	"log"
	"os"
=======
	"fmt"
	"http-go/src"
	"log"
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
	"time"
)

func main() {
<<<<<<< HEAD
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
=======
	log.Println("Main function has been called")

	cfg := src.ServerConfig{
		Address:      ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Logger:       log.Default(),
	}

	server := src.NewServer(cfg)
	fmt.Println("Server created at 8080", server)
>>>>>>> 9b2dcd4e58e15f4cde100f8bd0f9b2d3ce939485
}
