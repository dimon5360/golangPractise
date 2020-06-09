package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// BUILD : Current version of build
const BUILD uint32 = 2

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

func main() {

	timestamp := time.Now()
	fmt.Printf("Application version %d.%d.%d.\n", MAJOR, MINOR, BUILD)
	fmt.Println("Time of build: ", timestamp)
	fmt.Println("Start application.")

	StartTCPServer()

}

// StartTCPServer : Start TCP server on defined port 40400
func StartTCPServer() {
	fmt.Println("Start of TCP server")

	var port string = ":40400"

	// start listening the port
	ln, _ := net.Listen("tcp", port)

	// wait accept
	conn, _ := ln.Accept()

	for {
		// wait new message
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message received: ", string(message))

		// prepare the response
		newMessage := strings.ToUpper(message)

		//send the response
		conn.Write([]byte(newMessage + "\n"))
	}
}
