package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// function handles the client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	//buffer to read data
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't Read:", err)
		return
	}

	input, err := strconv.Atoi(string(buffer[:n])) //converting data to int
	if err != nil {
		fmt.Println("Can't convert to int:", err)
		return
	}

	response := input * 2
	conn.Write([]byte(strconv.Itoa(response))) //send back response
}

func main() {
	//Checking command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Using: server <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	//listen for connections from that port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Can't start the server:", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Server is listening on port", port)

	//Taking in connections to handle in separate goroutines
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Can't accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
