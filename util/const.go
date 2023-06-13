// @program:     rpc
// @file:        const.go
// @author:      ugug
// @create:      2023-06-12 01:46
// @description:

package util

//以下为protocol包中的常量
const (
	HeaderLenth = 4
	CheckerNum  = 0xff
)

//消息类型
const (
	Request byte = iota
	Response
)

//序列化类型
const (
	Gob byte = iota
	protobuf
	JSON
)

//连接类型
const (
	tcp byte = iota
	udp
	kcp
)

//以上为protocol包中的常量
