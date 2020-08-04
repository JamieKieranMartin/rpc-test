package main

import (
	"bufio"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Reply struct {
	Data string
}

func main() {
   client, err := jsonrpc.Dial("tcp", "localhost:12345") //Only change this
   if err != nil {
     log.Fatal(err)
   }
   in := bufio.NewReader(os.Stdin)
   for {
     line, _, err := in.ReadLine()
     if err != nil {
       log.Fatal(err)
     }
   var reply Reply
   err = client.Call("Listener.GetLine", line, &reply)
     if err != nil {
       log.Fatal(err)
     }
   log.Printf("Reply: %v, Data: %v", reply, reply.Data)
   }
}