package main

import (
	"fmt"
	"main/server"
	"time"
)

// TODO: rename vars to lowercase to private access

// BUILD : Current version of build
const BUILD uint32 = 3

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

// main : entry point app
func main() {
	PrintSwVersion()

	server.StartHTTPServer("8080")

	// TODO: now client is unused
	// client.StartTCPClient()
}

// PrintSwVersion : print out SW version and date of build
func PrintSwVersion() {

	timestamp := time.Now()
	fmt.Printf("Client application version %d.%d.%d.\n", MAJOR, MINOR, BUILD)
	fmt.Println("Time of build: ", timestamp)
	fmt.Println("Start application.")
}
