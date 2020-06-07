package main

import "fmt"

// BUILD : Current version of build
const BUILD uint32 = 1

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

func main() {

	fmt.Printf("Application version %d.%d.%d.\n", MAJOR, MINOR, BUILD)
	fmt.Println("Start application")
}
