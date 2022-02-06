# go-tcp
Transmission Control Protocol in Golang over secure channels

## Installation
```bash
go build .
```

## Generate a certificate
```bash
bash gen-cert.sh
```

## Usage

### Server
```bash
./go-tcp --ip 127.0.0.1 --port 8080 --mode server --key certs/server.key --cert certs/server.pem
```

### Client
```bash
./go-tcp --ip 127.0.0.1 --port 8080 --mode client --data "Hello World!" --key certs/client.key --cert certs/client.pem
```