package main

import (
	"fmt"
	"net"
	"sync"
)

const url = "scanme.nmap.org"

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", url, j)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				conn.Close()
				fmt.Printf("Port %d is open\n", j)
			}
		}(i)
	}
	wg.Wait()
}
