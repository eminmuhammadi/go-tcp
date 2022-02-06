# go-tcp
Transmission Control Protocol in Golang

## Installation
```bash
go build .
```

## Usage

### Server
```bash
./go-tcp --ip 127.0.0.1 --port 8080 --mode server
```

### Client
```bash
./go-tcp --ip 127.0.0.1 --port 8080 --mode client --data "Hello World!"
```