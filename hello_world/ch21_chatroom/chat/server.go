package chat

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
)

type Message chan string

const (
	MaxClientSize = 50
)

// 保存连接信息
type ClientTable map[net.Conn]*Client

type Server struct {
	uuid        int32
	listener    net.Listener
	clients     ClientTable
	availTokens int32
	pending     chan net.Conn
	quiting     chan net.Conn
	incoming    Message
	outgoing    Message
}

func (self *Server) generateToken() {
	atomic.AddInt32(&(self.availTokens), 1)
}

func (self *Server) loadToken() int32 {
	return atomic.LoadInt32(&(self.availTokens))
}

func (self *Server) takeToken() {
	atomic.AddInt32(&(self.availTokens), -1)
}

func CreateServer() *Server {
	server := &Server{
		uuid:        0,
		clients:     make(ClientTable, MaxClientSize),
		availTokens: MaxClientSize,
		pending:     make(chan net.Conn),
		quiting:     make(chan net.Conn),
		incoming:    make(Message),
		outgoing:    make(Message),
	}
	server.listen()
	return server
}

// 创建一个 listen 协程
func (self *Server) listen() {
	go func() {
		for {
			select {
			case msg := <-self.incoming:
				// 收到一个用户发送消息
				self.broadcast(msg)
			case conn := <-self.pending:
				//把新连接加入聊天室
				self.join(conn)
			case conn := <-self.quiting:
				// 离开聊天室
				self.leave(conn)
			}
		}
	}()
}

//TODO: 仅在 listen 协程里面操作map，保证线程安全
func (self *Server) broadcast(message string) {
	log.Printf("Broadcasting message: %s\n", message)
	for _, client := range self.clients {
		client.SendMsg(message)
	}
}

//TODO: 仅在 listen 协程里面操作map，保证线程安全
func (self *Server) join(conn net.Conn) {
	client := CreateClient(conn)
	name := self.getUniqName()
	client.SetName(name)
	self.clients[conn] = client
	log.Printf("Auto assigned name for conn %p: %s\n", conn, name)

	/// 启动一个新协程，负责当前 conn 的消息收客户端消息，以及向客户端发送消息
	go func() {
		for {
			select {
			case msg := <-client.incoming:
				log.Printf("Got message: %s from client %s\n", msg, client.GetName())
				if strings.HasPrefix(msg, ":") {
					log.Println("执行一些命令!")
				}
				// fallthrough to normal message if it is not parsable or executable
				self.incoming <- fmt.Sprintf("%s says: %s", client.GetName(), msg)
			case connNow := <-client.exitFlag:
				log.Printf("Client %s is quiting\n", client.GetName())
				// 发消息给chan, 让listen 协程进行处理即可
				self.quiting <- connNow
			}
		}
	}()
}

func (self *Server) getUniqName() string {
	atomic.AddInt32(&(self.uuid), 1)
	return "svr-" + strconv.Itoa((int)(self.uuid))
}

func (self *Server) leave(conn net.Conn) {
	if conn != nil {
		_ = conn.Close()
		delete(self.clients, conn)
	}
	self.generateToken()
}

func (self *Server) Start(bindAddr string) {
	var err error
	self.listener, err = net.Listen("tcp", bindAddr)
	if err != nil {
		log.Println("fail to bind:", bindAddr, err)
		return
	}

	for {
		conn, err1 := self.listener.Accept()
		if err1 != nil {
			log.Println(err1)
			return
		}

		if self.loadToken() > 0 {
			log.Printf("Exceeded the maximum limit of available connections %d.\n", MaxClientSize)
			_ = conn.Close()
			continue
		}
		// 把conn 通过消息发送给listen协程处理
		log.Printf("A new connection %v kicks\n", conn)
		self.takeToken()
		self.pending <- conn
	}
}

// FIXME: need to figure out if this is the correct approach to gracefully
// terminate a server.
func (self *Server) Stop() {
	_ = self.listener.Close()
}
