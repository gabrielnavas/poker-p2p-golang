package p2p

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

type ServerConfig struct {
	ListenAddr string
}

type Server struct {
	ServerConfig

	handler  Handler
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer
	msgCh    chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		handler:      *NewHandler(),
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		msgCh:        make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()
	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("game server running on port %s\n", s.ListenAddr)

	s.acceptLoop()
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// TODO: handle accept listener error
			panic(err)
		}

		// add peer
		peer := &Peer{
			conn: conn,
		}
		s.addPeer <- peer

		// send first message to peer
		peer.Send([]byte("GGPOKER V0.1=alpha"))

		go s.handleConn(conn)
	}
}

func (s *Server) sendMessageFromConnection(conn net.Conn, buf []byte, lenBuff int) {
	s.msgCh <- &Message{
		From:    conn.RemoteAddr(),
		Payload: bytes.NewReader(buf[:lenBuff]),
	}
}

func (s *Server) handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		// get buf from connection
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		go s.sendMessageFromConnection(conn, buf, n)
	}
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.listener = ln
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeer:
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
			s.peers[peer.conn.RemoteAddr()] = peer
		case message := <-s.msgCh:
			msg, _ := io.ReadAll(message.Payload)
			fmt.Printf("message from: %s -> %s\n", message.From, string(msg))
		}
	}
}
