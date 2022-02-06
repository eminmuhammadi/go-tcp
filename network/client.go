package network

import (
	"fmt"
	"net"
)

func Client(ip string, port string, data string) error {
	// Listener
	listener, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return err
	}

	// Connection
	connection, err := net.DialTCP("tcp", nil, listener)
	if err != nil {
		return err
	}

	// Send
	if _, err := connection.Write([]byte(data + "\n")); err != nil {
		return err
	}

	// Receive
	reply := make([]byte, 1024) // TODO: Data buffer

	if _, err := connection.Read(reply); err != nil {
		return err
	}

	// Data print
	println(string(reply))

	// Close
	if err := connection.Close(); err != nil {
		return err
	}

	return nil
}
