package main

import "ggpoker/p2p"

func main() {
	// use netcat to test
	// like this: nc -T localhost 3000

	cfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":3000",
	}
	s := p2p.NewServer(cfg)
	go s.Start()

	remoteCfg := p2p.ServerConfig{
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":4000",
	}
	remoteServer := p2p.NewServer(remoteCfg)
	remoteServer.Connect(cfg.ListenAddr)
	remoteServer.Start()
}
