package client

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	NETWORK_TYPE = "tcp"
)

type ConntestClient struct {
	Host string
	Port int
}

func (c *ConntestClient) Init() error {
	return nil
}

func (c *ConntestClient) Start() error {
	c.testConnection()
	return nil
}

func (c *ConntestClient) testConnection() {
	strEcho := "Halo"
	servAddr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	conn, err := net.Dial(NETWORK_TYPE, servAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	log.Println("reply from server=", string(reply))

	conn.Close()
}
