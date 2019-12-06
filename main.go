package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	proto := new(Proto)

	log.Println("registering configtxlator tool...")
	if err := rpc.Register(proto); err != nil {
		log.Fatalln("register proto tools error:", err)
	}

	compute := new(Compute)
	if err := rpc.Register(compute); err != nil {
		log.Fatalln("register compute tools error:", err)
	}

	rpc.HandleHTTP()

	log.Println("listening on port 1234...")
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	if err := http.Serve(l, nil); err != nil {
		log.Fatalln("running server error:", err)
	}
}
