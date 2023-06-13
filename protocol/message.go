// @program:     rpc
// @file:        message.go
// @author:      ugug
// @create:      2023-06-13 08:40
// @description:

package protocol

import (
	"rpc/module"
	"rpc/util"
)

// Header 头部组成：|特定数1|消息1|序列化1|连接1|=4bytes
type Header [util.HeaderLenth]byte

// msg 二进制消息体
type msg struct {
	*Header
	info module.Info //payload
}

func NewMsg() *msg {
	header := Header([util.HeaderLenth]byte{})
	header[0] = util.CheckerNum
	return &msg{
		Header: &header,
	}
}
