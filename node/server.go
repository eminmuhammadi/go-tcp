package node

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
)

func Server(ip string, port string, certFile string, keyFile string) error {
	// TLS Certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	// Listener
	listener, err := tls.Listen("tcp", fmt.Sprintf("%s:%s", ip, port), &tls.Config{
		/* For self-signed certificates ignored
		ClientAuth:               tls.RequireAndVerifyClientCert,
		*/
		Certificates:             []tls.Certificate{cert},
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
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

func processData(connection net.Conn) error {
	reader := bufio.NewReader(connection)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			connection.Close()
			return err
		}

		// Data => string(data)
		// Reprocess data
		connection.Write([]byte(string(data)))
	}
}
