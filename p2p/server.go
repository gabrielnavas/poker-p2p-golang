package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

type ServerConfig struct {
	Version    string
	ListenAddr string
}

type Server struct {
	ServerConfig

	handler  Handler
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer
	delPeer  chan *Peer
	msgCh    chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	dh := NewDefaultHandler()
	return &Server{
		handler:      dh,
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
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
		peer.Send([]byte(s.Version))

		go s.handleConn(peer)
	}
}

func (s *Server) readMessageLoop(peer *Peer) {
	buf := make([]byte, 1024)
	for {
		n, err := peer.conn.Read(buf)
		if err != nil {
			break
		}

		s.msgCh <- &Message{
			From:    peer.conn.RemoteAddr(),
			Payload: bytes.NewReader(buf[:n]),
		}
	}
}

func (s *Server) removePeer(peer *Peer) {
	s.delPeer <- peer
}

func (s *Server) handleConn(peer *Peer) {
	s.readMessageLoop(peer)
	s.removePeer(peer)
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
		case peer := <-s.delPeer:
			delete(s.peers, peer.conn.RemoteAddr())
			fmt.Printf("player disconnected %s\n", peer.conn.RemoteAddr())
		case peer := <-s.addPeer:
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
			s.peers[peer.conn.RemoteAddr()] = peer
		case message := <-s.msgCh:
			if err := s.handler.HandleMessage(message); err != nil {
				panic(err)
			}
		}
	}
}
