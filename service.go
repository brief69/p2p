package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
)

type P2PService1 struct {
	host host.Host
}

type Config1 struct {
	Host string
	Port int
}

func NewP2PService(config *Config) (*P2PService, error) {
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/%d", config.Host, config.Port)),
	)
	if err != nil {
		return nil, fmt.Errorf("P2Pホストの作成に失敗しました: %s", err.Error())
	}

	service := &P2PService{
		host: h,
	}

	h.SetStreamHandler("/p2p/1.0.0", service.handleStream)

	return service, nil
}

func (s *P2PService) StartService() {
	// これは、サービスに関連する長時間実行のタスクを開始する場所です。
	// この場合、ホストが自動的に着信接続を処理するため、何もする必要はありません。
}

func (s *P2PService) Terminate() {
	// これは、サービスに関連する長時間実行のタスクを停止する場所です。
	// この場合、ホストを閉じるだけで済みます。
	s.host.Close()
}

func (s *P2PService) handleIncomingStream(stream network.Stream) {
	// これは、着信ストリームを処理する場所です。
	// 現時点では、メッセージを印刷してストリームを閉じるだけです。
	fmt.Printf("%sからストリームを受信しました！\n", stream.Conn().RemotePeer().Pretty())
	stream.Close()
}
