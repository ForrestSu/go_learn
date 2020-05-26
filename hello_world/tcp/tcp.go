package main

import (
	"fmt"
	"net"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// 处理一个 socket 连接
func recvSocket(serverConn net.Conn) {
	fmt.Printf("recv %x\n", serverConn)
	fmt.Printf("来自 %v 的连接\n", serverConn.RemoteAddr())

	buf := make([]byte, 5)
	for {
		n, err := serverConn.Read(buf)
		//3.2 数据读尽、读取错误 关闭 socket 连接
		if n == 0 || err != nil {
			fmt.Println("n == ", n, err)
			break
		}
		data := string(buf[0:n])
		fmt.Printf("n = %d , data = <%s>\n", n, data)
		nBytes, _ := serverConn.Write([]byte("<= hello\n"))
		fmt.Printf("resp nBytes = %d\n", nBytes)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	panicOnError(err)
	defer func() {
		err := listen.Close()
		if err != nil {
			fmt.Println("Close error", err)
		}
	}()

	for {
		conn, err := listen.Accept()
		// 当fd 耗尽, Accept 会返回错误，一般采取重试即可
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		go recvSocket(conn)
	}
}
