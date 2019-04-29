package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("\nbefore starting this, run vid21_tcp_server/tcp2.go in another termnal window\n")
	fmt.Println("\nrun this and verify the text in here gets printed in that window\n")
	// this app will talk to that one
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// above
	defer conn.Close()

	// since the other program is looking for newlines, make sure to write a LF
	fmt.Fprintln(conn, "..I am writing to you...\n")

}
