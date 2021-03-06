package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// TODO: rename vars to lowercase to private access

// BUILD : Current version of build
const BUILD uint32 = 1

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

// StartTCPClient : Start TCP client
func StartTCPClient() {

	var ip string = "0.0.0.0:40400"
	fmt.Println("Start of TCP client")

	// connect to host
	conn, _ := net.Dial("tcp", ip)

	// for {
	// reading data from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	// Send to socket
	fmt.Fprintf(conn, text+"\n")
	// wait an answer
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
	// }
}
