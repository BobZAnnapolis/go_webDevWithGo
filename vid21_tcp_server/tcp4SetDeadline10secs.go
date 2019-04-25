package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	msg := "\nafter starting this, go to another terminal window,type telnet localhost 8080 and start entering characters"
	msg += "\nor localhost:8080 in a browser"
	msg += "\n\nthis connection, once established, only stays open for 10 seconds"
	fmt.Println(msg)

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
	deadlineSeconds := 10
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// text by default splits at newlines - can change this
		ln := scanner.Text()
		// print the line received here
		fmt.Println(ln)
		// print this text on the connection
		fmt.Fprintf(conn, "\nI heard you say %s\n", ln)
	}
	defer conn.Close()

	// we never get here under normal operating conditions
	// we have an open stream connection
	// how does the above reader know when it's done ?
	// ctrl-c to end this program or end the telnet sessions
	fmt.Printf("Code gets here when we ctrl-c out of the program\n\tOR\ndeadline of %v seconds reached", deadlineSeconds)
}
