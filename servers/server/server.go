package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	NETWORK_TYPE = "tcp"
)

type ConntestServer struct {
	Host string
	Port int
}

func (c *ConntestServer) Init() error {
	return nil
}

func (c *ConntestServer) Start() error {
	c.startListener()
	return nil
}

func (c *ConntestServer) startListener() {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	listen, err := net.Listen(NETWORK_TYPE, addr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back34!\n"))
	conn.Write([]byte(time))

	// close conn
	conn.Close()
}
