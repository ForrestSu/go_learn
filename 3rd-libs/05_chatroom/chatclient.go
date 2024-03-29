package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ForrestSu/go_learn/3rd-libs/05_chatroom/chat"
)

//go:generate ./chatclient 172.17.0.2:8000

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <host:port>\n", os.Args[0])
		os.Exit(-1)
	}

	conn, err := net.Dial("tcp", os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	client := chat.CreateClient(conn)

	// coroutine-1
	go func() {
		for {
			msg := <-client.GetIncoming()
			_, _ = out.WriteString(msg + "\n")
			_ = out.Flush()
		}
	}()

	for {
		line, _, _ := in.ReadLine()
		client.SendMsg(string(line))
	}
}
