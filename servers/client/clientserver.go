package main

import (
	"austeus.com/client"
	"austeus.com/service"
)

const (
	CLIENT_HOST = "localhost"
	CLIENT_PORT = 7001
)

func main() {
	service.Add(&client.ConntestClient{Host: CLIENT_HOST, Port: CLIENT_PORT})

	service.RunAll()
}
