package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {
	//w := make(chan struct{}, 1)

	fmt.Println("Starting Sim7600 module")
	cl := &Sim7600GClient{
		SerialBaud: 115200,
		SerialPort: "COM8",
	}

	openConn(cl)
	initConn(cl)
	closeConn(cl)

	fmt.Println("Sim7600 done")
}

func openConn(cl *Sim7600GClient) {
	cl.listSerialPorts()
	cl.openSerialPort()
}

func initConn(cl *Sim7600GClient) {
	cl.ATQuery()
	cl.enableEchoVerboseAndGetInfo()
}

func closeConn(cl *Sim7600GClient) {
	cl.closeSerialPort()
}

type Sim7600GClient struct {
	SerialPort string
	SerialBaud int

	Port serial.Port
}

func NewSim7600GClient(port string, baud int) *Sim7600GClient {
	return &Sim7600GClient{
		SerialPort: port,
		SerialBaud: baud,
	}
}

func (s *Sim7600GClient) listSerialPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func (s *Sim7600GClient) openSerialPort() {
	mode := &serial.Mode{
		BaudRate: s.SerialBaud,
		DataBits: 8,
		Parity:   serial.NoParity,
	}

	port, err := serial.Open(s.SerialPort, mode)
	if err != nil {
		log.Fatal(err)
	}

	s.Port = port
}

func (s *Sim7600GClient) closeSerialPort() {
	err := s.Port.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Sim7600GClient) writeCommandLine(cmd *CommandLine) {
	data := cmd.Command()
	n, err := s.Port.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 100)
	respData := ""
	for {
		n, err := s.Port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		respData += string(buff[:n])
		if HasResultCode(respData) {
			break
		}
	}

	NewCommandResponseFromData(respData)
}

func (s *Sim7600GClient) writeCommandFromBody(body string) {
	cmdLine := NewCommandLineFromBody(body)
	s.writeCommandLine(cmdLine)
}

// sends a simple AT query
func (s *Sim7600GClient) ATQuery() {
	s.writeCommandFromBody("")
}

func (s *Sim7600GClient) enableEchoVerboseAndGetInfo() int {
	// enable echo and verbose mode
	s.writeCommandFromBody("E1V1")

	// get Model
	s.writeCommandFromBody("+CGMM")

	// get Manufacturer
	s.writeCommandFromBody("+CGMI")

	return -1
}

func (s *Sim7600GClient) getSignalStrength() int {
	s.writeCommandFromBody("+CSQ")

	return -1
}
