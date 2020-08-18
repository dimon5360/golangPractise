package main

import (
	"fmt"
	"main/client"
	"main/server"
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
	fmt.Printf("Client application version %d.%d.%d.\n", MAJOR, MINOR, BUILD)
	fmt.Println("Time of build: ", timestamp)
	fmt.Println("Start application.")

	server.StartTCPServer()

	client.StartTCPClient()
}
