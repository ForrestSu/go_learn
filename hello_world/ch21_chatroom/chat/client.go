package chat

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	conn     net.Conn
	incoming Message
	outgoing Message
	reader   *bufio.Reader
	writer   *bufio.Writer
	exitFlag chan net.Conn
	name     string
}

func (self *Client) GetName() string {
	return self.name
}

func (self *Client) SetName(name string) {
	self.name = name
}

func (self *Client) GetIncoming() Message {
	return self.incoming
}

//发送消息
func (self *Client) SendMsg(msg string) {
	self.outgoing <- msg
}

func CreateClient(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{
		conn:     conn,
		incoming: make(Message),
		outgoing: make(Message),
		reader:   reader,
		writer:   writer,
        exitFlag: make(chan net.Conn),
	}
    client.StartPoll()
	return client
}

func (self *Client) StartPoll() {
	go self.readSync()
	go self.writeSync()
}

//同步阻塞的方式，读取数据
func (self *Client) readSync() {
	for {
		if line, _, err := self.reader.ReadLine(); err == nil {
			self.incoming <- string(line)
		} else {
			log.Printf("Read error: %s\n", err)
			self.Quit()
			return
		}
	}
}

func (self *Client) writeSync() {
	for data := range self.outgoing {
		if _, err := self.writer.WriteString(data + "\n"); err != nil {
			self.Quit()
			return
		}

		if err := self.writer.Flush(); err != nil {
			log.Printf("Write error: %s\n", err)
			self.Quit()
			return
		}
	}
}
func (self *Client) Quit() {
	self.exitFlag <- self.conn
}

func (self *Client) Close() {
	_ = self.conn.Close()
}

func main() {
	// conn := new(net.Conn)
	// client = CreateClient(*conn)
}
