package main

import (
	"fmt"
	"net"
)

func main() {
	connectionListener, err := net.Listen("tcp", ":5500")

	if err == nil {
		fmt.Println("Error with tcp connection to port 5500:", err)
	}

	for {
		connection, err := connectionListener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		fmt.Println("New connection")
		go handleUserConnection(connection)
	}

}

func handleUserConnection(connection net.Conn) {
	// Make sure connection is closed once user is disconnected
	defer connection.Close()

}
