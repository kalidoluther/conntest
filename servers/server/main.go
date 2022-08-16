package main

import (
	"austeus.com/server"
	"austeus.com/service"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = 7001
)

func main() {
	service.Add(&server.ConntestServer{Host: SERVER_HOST, Port: SERVER_PORT})

	service.RunAll()
}
