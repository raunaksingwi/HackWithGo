package main

import (
	"fmt"
	"net"
)

const (
	url = "scanme.nmap.org"
)

func main() {

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("%s:%d", url, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is not open
			continue
		} else {
			fmt.Printf("Port %d is open\n", i)
			conn.Close()
		}
	}
}
