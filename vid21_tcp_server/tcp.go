package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	fmt.Println("\nafter starting this, go to another termnal window and type telnet localhost 8080\n")
	// net.Listen takes 2 parms : 1-what kind of network do you want to listen on, and, 2-what port
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	// above Listen returns a Listerner that HAS TO BE CLOSED
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day ?")
		fmt.Fprintf(conn, "%v", "Well, i hope !")

		conn.Close()
	}
}
