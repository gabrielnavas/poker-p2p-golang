package main

import "ggpoker/p2p"

func main() {
	// use netcat to test
	// like this: nc -T localhost 3000

	cfg := p2p.ServerConfig{
		ListenAddr: ":3000",
	}
	s := p2p.NewServer(cfg)
	s.Start()
}
