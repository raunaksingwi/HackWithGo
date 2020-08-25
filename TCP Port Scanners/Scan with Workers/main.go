package main

import (
	"fmt"
	"net"
	"sort"
)

const url = "scanme.nmap.org"

func worker(ports <-chan int, results chan<- int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", url, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
		} else {
			conn.Close()
			results <- port
		}
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
