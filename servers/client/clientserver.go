package main

const (
	CLIENT_HOST = "localhost"
	CLIENT_PORT = 7001
)

func main() {
	Add(&ConntestClient{Host: CLIENT_HOST, Port: CLIENT_PORT})

	RunAll()
}
