package network

import (
	"bufio"
	"fmt"
	"net"
)

func Server(ip string, port string) error {
	// Listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return err
	}

	// Close
	defer listener.Close()

	// Accept new connections
	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}

		go processData(connection)
	}
}

func processData(connection net.Conn) error {
	reader := bufio.NewReader(connection)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		// Data => string(data)
		// Reprocess data
		connection.Write([]byte(string(data)))
	}
}
