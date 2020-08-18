package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// BUILD : Current version of build
const BUILD uint32 = 2

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

// StartTCPServer : Start TCP server on localhost:40400
func StartTCPServer() {

	var port string = ":40400"
	fmt.Println("Start of TCP server on port localhost", port)

	// start listening the port
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Listening failed. Program ends.")
		return
	}

	// wait accept
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Accepting failed. Program ends.")
		return
	}

	for {
		// wait new message
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input msg. Program ends.")
			break
		}

		fmt.Print("Message received: ", string(message))

		// prepare the response
		newMessage := strings.ToUpper(message)

		//send the response
		conn.Write([]byte(newMessage + "\n"))
	}
}
