package node

import (
	"crypto/tls"
	"fmt"
)

// Client
func Client(ip string, port string, data string, certFile string, keyFile string) error {
	// TLS Certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	// Connection
	connection, err := tls.Dial("tcp4", fmt.Sprintf("%s:%s", ip, port), &tls.Config{
		InsecureSkipVerify: true, // For self-signed certificates
		Certificates:       []tls.Certificate{cert},
	})

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
