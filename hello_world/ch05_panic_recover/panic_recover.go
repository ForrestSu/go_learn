package main

import "fmt"

func main() {
    defer func() {
        info := recover()
        fmt.Println("hello world!", info)
    }()

    go func() {
        panic("hhh")
    }()
}


