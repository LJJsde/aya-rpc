// @program:     rpc
// @file:        conn.go
// @author:      ugug
// @create:      2023-06-11 00:13
// @description:

package connection

import (
	"crypto/sha1"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"net"
	"rpc/service"
	"testing"
)

// TestListen 建立tcp链接
func TestListen(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())
}

// KCPListenStart 建立kcp链接
func KCPListenStart() {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	if listener, err := kcp.ListenWithOptions("127.0.0.1:12345", block, 10, 3); err == nil {
		// spin-up the client
		go service.Client()
		for {
			s, err := listener.AcceptKCP()
			if err != nil {
				log.Fatal(err)
			}
			go service.HandleEcho(s)
		}
	} else {
		log.Fatal(err)
	}
}
