package main

import (
	"economics/products/internal/config"
	"economics/products/internal/server"
)

func main() {
	config.Init()
	server.Run()
}