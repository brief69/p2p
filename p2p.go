package main

import (
	"fmt"
	"net"
	"sync"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
)

type P2PService struct {
	host host.Host
}

type Config struct {
	Host string
	Port int
}

func 新しいP2PService(config *Config) (*P2PService, error) {
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/%d", config.Host, config.Port)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create P2P host: %s", err.Error())
	}

	service := &P2PService{
		host: h,
	}

	h.SetStreamHandler("/p2p/1.0.0", service.handleStream)

	return service, nil
}

func (s *P2PService) Start() {
	// This is where you would start any long-running tasks associated with the service.
	// In this case, there's nothing to do, because the host automatically handles incoming connections.
}

func (s *P2PService) Stop() {
	// This is where you would stop any long-running tasks associated with the service.
	// In this case, we just need to close the host.
	s.host.Close()
}

func (s *P2PService) handleStream(stream network.Stream) {
	// これは、着信ストリームを処理する場所です。
	// 現時点では、メッセージを印刷してストリームを閉じるだけです。
	fmt.Printf("%sからストリームを受信しました！\n", stream.Conn().RemotePeer().String())
	stream.Close()
}
