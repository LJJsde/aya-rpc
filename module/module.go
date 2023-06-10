package module

import (
	"github.com/xtaci/kcp-go/v5"
	"time"
)

type Client struct {
	RegisterTime time.Time
	LogoutTime   time.Time
	Conn         kcp.Listener
}

type Server struct {
	ClientNum int64
}

type Info struct {
	TimeStamp time.Time
	Method    []interface{}
	Args      []interface{}
}
