package main

import (
	"go-world/config"
	"go-world/server"
)

func main() {
	config.LoadConfig()

	server.Start()
}
