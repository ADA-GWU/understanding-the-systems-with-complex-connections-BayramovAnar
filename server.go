package main

import (
	"fmt"
	"net"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Can't Read:", err)
		return
	}

	input, err := strconv.Atoi(string(buffer[:n]))
	if err != nil {
		fmt.Println("Can't convert to int:", err)
		return
	}

	response := input * 2
	conn.Write([]byte(strconv.Itoa(response)))
}

func main() {
	// don't forget to run a func here!!!
}
