package main

import (
	"fmt"
	"http-go/src"
	"log"
	"time"
)

func main() {
	log.Println("Main function has been called")

	cfg := src.ServerConfig{
		Address:      ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Logger:       log.Default(),
	}

	server := src.NewServer(cfg)
	fmt.Println("Server created at 8080", server)
}
