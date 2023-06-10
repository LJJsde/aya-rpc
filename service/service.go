// @program:     rpc
// @file:        service.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package service

import (
	"fmt"
	"rpc/module"
	"time"
)

type HelloService struct{}

func (s *HelloService) Hello(in module.Info) module.Info {
	var out module.Info
	out.Args[1] = "hello" + fmt.Sprintf("%v", in.Args[1])
	out.TimeStamp = time.Now()
	return out
}
