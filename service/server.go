// @program:     rpc
// @file:        server.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package service

import (
	"fmt"
	"reflect"
	"rpc/connection"
	"time"
)

type server struct {
	ClientNum  int64
	ServiceMap map[string]*service
}

func ServerInit(network string, address string) error {
	return serverInit(network, address)
}

func serverInit(network string, address string) error {
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
	return fmt.Errorf("network error input another")
}

func (s *server) Register(name string, receiver []interface{}) error {
	return s.register(name, receiver)
}

func (s *server) register(name string, receiver []interface{}) error {
	NewService := service{
		name:     name,
		args:     receiver,
		initTime: time.Now(),
	}
	var typeSlice []reflect.Type
	for i := 0; receiver[i] != nil; i++ {
		typeSlice[i] = reflect.TypeOf(receiver[i])
	}
	typeSlice = NewService.argsType
	s.ServiceMap = make(map[string]*service)
	return nil
}
