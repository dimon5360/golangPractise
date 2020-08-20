package server

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
)

// TODO: rename vars to lowercase to private access

// BUILD : Current version of build
const BUILD uint32 = 3

// MINOR : Minor version of application
const MINOR uint32 = 0

// MAJOR : Major version of application
const MAJOR uint32 = 0

// StartTCPServer : Start TCP server on localhost:40400
func StartTCPServer(port string) {

	fmt.Println("Start of TCP server on port localhost:", port)

	// start listening the port
	ln, err := net.Listen("tcp", ":"+port)
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
	fmt.Println("Accepting succed. Wait input msg.")

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

type state struct {
	*sync.Mutex                   // inherit locked methods
	Vals        map[string]string // map ids to values
}

// State : what is it??
var State = &state{&sync.Mutex{}, map[string]string{}}

// StartHTTPServer : function starts HTTP server
func StartHTTPServer(port string) {

	fmt.Println("Start of TCP server on port localhost:", port)

	err := http.ListenAndServe("0.0.0.0:"+port, http.HandlerFunc(handle))
	if err != nil {
		fmt.Println(err)
	}
}

// handle : function-handler to process input HTTP requests
func handle(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Got input request")
	switch req.Method {
	case "POST":
		postHandler(rw, req)

	case "GET":
		if req.URL.String() == "/form" {
			formHandler(rw, req)
			return
		}
		getHandler(rw, req)
	}
}

func getHandler(rw http.ResponseWriter, req *http.Request) {
	State.Lock()
	defer State.Unlock()            // ensure the lock is removed after leaving the the function
	id := req.URL.Query().Get("id") // keep id from request
	val := State.Vals[id]
	delete(State.Vals, id)
	rw.Write(([]byte("got: " + val)))
}

func postHandler(rw http.ResponseWriter, req *http.Request) {
	State.Lock()
	defer State.Unlock()      // ensure the lock is removed after leaving the the function
	id := req.FormValue("id") // keep id from request
	State.Vals[id] = req.FormValue("val")
	rw.Write(([]byte("go to http://0.0.0.0:8080/?id=42")))
}

var form = `<html>
    <body>
        <form action="/" method="POST">
            ID: <input name="id" value="42" /><br />
            Val: <input name="val" /><br />
            <input type="submit" value="submit"/>
        </form>
    </body>
</html>`

func formHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(form))
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// type state struct {
// 	*sync.Mutex                   // inherits locking methods
// 	Vals        map[string]string // map ids to values
// }

// var State = &state{&sync.Mutex{}, map[string]string{}}

// func get(rw http.ResponseWriter, req *http.Request) {
// 	State.Lock()
// 	defer State.Unlock()            // ensure the lock is removed after leaving the the function
// 	id := req.URL.Query().Get("id") // if you need other types, take a look at strconv package
// 	val := State.Vals[id]
// 	delete(State.Vals, id)
// 	rw.Write([]byte("got: " + val))
// }

// func post(rw http.ResponseWriter, req *http.Request) {
// 	State.Lock()
// 	defer State.Unlock()
// 	id := req.FormValue("id")
// 	State.Vals[id] = req.FormValue("val")
// 	rw.Write([]byte("go to http://localhost:8080/?id=42"))
// }

// var form = `<html>
//     <body>
//         <form action="/" method="POST">
//             ID: <input name="id" value="42" /><br />
//             Val: <input name="val" /><br />
//             <input type="submit" value="submit"/>
//         </form>
//     </body>
// </html>`

// func formHandler(rw http.ResponseWriter, req *http.Request) {
// 	rw.Write([]byte(form))
// }

// // for real routing take a look at gorilla/mux package
// func handler(rw http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case "POST":
// 		post(rw, req)
// 	case "GET":
// 		if req.URL.String() == "/form" {
// 			formHandler(rw, req)
// 			return
// 		}
// 		get(rw, req)
// 	}
// }

// func main() {
// 	fmt.Println("go to http://localhost:8080/form")
// 	// thats the default webserver of the net/http package, but you may
// 	// create custom servers as well
// 	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
