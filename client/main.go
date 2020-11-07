package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:12345") //Only change this
	if err != nil {
		log.Fatal(err)
	}

	var reply int
	err = client.Call("Math.Add", &Args{3, 4}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %v", reply)

}
