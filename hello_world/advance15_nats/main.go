package main

import (
    "fmt"
    "github.com/nats-io/nats.go"
    "time"
)

func main() {
    // Connect to a server
    nc, _ := nats.Connect("nats://192.168.5.134:4222")

    // Simple Publisher
    err := nc.Publish("foo", []byte("Hello World"))
    if err != nil {
        fmt.Println("error", err)
        return
    }

    // Simple Async Subscriber
    sub, err := nc.Subscribe("foo", func(m *nats.Msg) {
        fmt.Printf("Received a message: %s\n", string(m.Data))
    })
    fmt.Println("sub = ", sub)
    go func() {
        for {
            fmt.Println("pub...")
            nc.Publish("foo", []byte("Hello World"))
            time.Sleep(1 * time.Second)
        }
    }()
    // keep
    for {
        time.Sleep(1 * time.Second)
    }
}
