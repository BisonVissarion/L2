package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	host := flag.String("host", "", "Hostname or IP address")
	port := flag.String("port", "23", "Port number")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if *host == "" {
		fmt.Println("Please specify a host using the -host flag.")
		return
	}

	address := *host + ":" + *port

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	go copyToServer(os.Stdin, conn)
	copyToClient(conn, os.Stdout)
}

func copyToServer(src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func copyToClient(src io.Reader, dst io.Writer) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Connection closed.")
		os.Exit(0)
	}
}
