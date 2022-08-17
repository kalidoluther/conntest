package main

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = 7001
)

func main() {
	Add(&ConntestServer{Host: SERVER_HOST, Port: SERVER_PORT})
	RunAll()
}
