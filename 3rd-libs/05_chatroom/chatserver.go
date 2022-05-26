package main

import (
	"fmt"
	"os"

	"github.com/ForrestSu/go_learn/3rd-libs/05_chatroom/chat"
)

//go:generate ./chatserver  172.17.0.2:8000

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <host:port>\n", os.Args[0])
		os.Exit(-1)
	}

	server := chat.CreateServer()
	fmt.Printf("Running on %s\n", os.Args[1])
	server.Start(os.Args[1])
}
