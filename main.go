package main

import (
	"flag"

	network "github.com/eminmuhammadi/go-tcp/network"
)

var ip = flag.String("ip", "0.0.0.0", "--ip 0.0.0.0")
var port = flag.String("port", "8080", "--port 8080")
var data = flag.String("data", "", "--data \"Hello World\"")
var mode = flag.String("mode", "server", "--mode server/client")

func main() {
	flag.Parse()

	if *mode == "server" {
		if err := network.Server(*ip, *port); err != nil {
			panic(err)
		}
	}

	if *mode == "client" {
		if err := network.Client(*ip, *port, *data); err != nil {
			panic(err)
		}
	}
}
