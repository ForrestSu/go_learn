package main

import (
	. "./chat"
	"fmt"
	"os"
)

// ./chatserver  172.17.0.2:8000
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <host:port>\n", os.Args[0])
		os.Exit(-1)
	}

	server := CreateServer()
	fmt.Printf("Running on %s\n", os.Args[1])
	server.Start(os.Args[1])
}
