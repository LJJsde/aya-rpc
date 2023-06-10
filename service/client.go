// @program:     rpc
// @file:        client.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package service

import (
	"reflect"
	"rpc/connection"
	"rpc/module"
	"time"
)

type Client struct {
	RegisterTime time.Time
	LogoutTime   time.Time
	//conn可选连接类型
	Conn interface{}
}

func ClientInit(network string, address string) {
	switch network {
	case "tcp":
		{
			connection.TestListen(nil, network, address)

		}
		//个人兴趣
	case "kcp":
		{
			connection.KCPListenInit(address)
		}
	case "udp":
		{
		}

	}
}

// Call 步骤（待优化
//1.吃进info类
//2.测量切片长度生成中转切片（待优化
//3.中转切片赋值完毕生成出去的info类
//4.吐出的info类进行序列化传至服务器，开始远程调用（in progress
//5.调用完毕将结果传入io（in progress
//6.输出结果（in progress
func (c *Client) Call(in module.Info) module.Info {
	if reflect.TypeOf(in.Args).Kind() != reflect.Slice {
		out := module.Info{
			ErrCode: true, ErrMessage: "args type error"}
		return out
	}
	ins := reflect.ValueOf(in.Args)
	len := ins.Len()
	outSlice := make([]interface{}, len)
	for i := 0; i < ins.Len(); i++ {
		outSlice[i] = in.Args[i]
	}
	out := module.Info{
		ErrCode:    false,
		ErrMessage: "",
		Method:     in.Method,
		Args:       outSlice,
		TimeStamp:  time.Now(),
	}
	return out
}
