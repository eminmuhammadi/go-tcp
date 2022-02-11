package node

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"

	engine "github.com/eminmuhammadi/go-tcp/engine"
)

// Server
func Server(ip string, port string, certFile string, keyFile string) error {
	// TLS Certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	// Listener
	listener, err := tls.Listen("tcp4", fmt.Sprintf("%s:%s", ip, port), &tls.Config{
		Certificates: []tls.Certificate{cert},
	})

	if err != nil {
		return err
	}

	// Close
	defer listener.Close()

	// Accept new connections
	for {
		connection, err := listener.Accept()
		if err != nil {
			connection.Close()
			return err
		}

		// Close
		defer connection.Close()

		// Process data
		go processData(connection)
	}
}

// Process data
func processData(connection net.Conn) error {
	reader := bufio.NewReader(connection)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			connection.Close()
			return err
		}

		// Handle data
		connection.Write([]byte(engine.Handler(string(data))))
	}
}
