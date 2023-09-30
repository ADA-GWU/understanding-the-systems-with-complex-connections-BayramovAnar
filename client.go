package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Checking command line arguments
	if len(os.Args) != 3 {
		fmt.Println("<server:port> <number>")
		os.Exit(1)
	}

	//Taking server addr and number from user input
	servAddr := os.Args[1]
	number := os.Args[2]

	//Make a connection to server
	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		fmt.Println("Can't connect to server:", err)
		os.Exit(1)
	}

	defer conn.Close()

	//Send the number in the form of the string
	_, err = conn.Write([]byte(number))
	if err != nil {
		fmt.Println("Can't send data:", err)
		os.Exit(1)
	}

	//read the response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't read response:", err)
		os.Exit(1)
	}

	//Convert the response
	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}
