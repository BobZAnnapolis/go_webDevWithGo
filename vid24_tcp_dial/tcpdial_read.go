package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	fmt.Println("\nbefore starting this, run vid21_tcp_server/tcp.go in another termnal window\n")
	// this app will talk to that one
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// above
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))

}
