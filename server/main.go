package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

type Math int

func (m *Math) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Math)
	rpc.Register(listener)
	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
