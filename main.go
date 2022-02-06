package main

import (
	"flag"

	node "github.com/eminmuhammadi/go-tcp/node"
)

var ip = flag.String("ip", "0.0.0.0", "--ip 0.0.0.0")
var port = flag.String("port", "8080", "--port 8080")
var data = flag.String("data", "", "--data \"Hello World\"")
var mode = flag.String("mode", "server", "--mode server/client")
var key = flag.String("key", "", "--key \"/path/to/key\"")
var cert = flag.String("cert", "", "--cert \"/path/to/cert\"")

func main() {
	flag.Parse()

	if *mode == "server" {
		if err := node.Server(*ip, *port, *cert, *key); err != nil {
			panic(err)
		}
	}

	if *mode == "client" {
		if err := node.Client(*ip, *port, *data, *cert, *key); err != nil {
			panic(err)
		}
	}
}
