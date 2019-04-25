package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	fmt.Println("\nafter starting this, go to another terminal window,type telnet localhost 8080 and start entering characters\nor localhost:8080 in a browser")
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// text by default splits at newlines - can change this
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get here under normal operating conditions
	// we have an open stream connection
	// how does the above reader know when it's done ?
	// ctrl-c to end this program or end the telnet sessions
	fmt.Println("Code gets here when we ctrl-c out of the program")
}
