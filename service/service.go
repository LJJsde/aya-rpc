// @program:     rpc
// @file:        service.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package service

import (
	"fmt"
	"reflect"
	"rpc/module"
	"time"
)

//泛化的服务结构体，一个系列的服务共用一个结构体

type service struct {
	name      string         //服务名
	args      []interface{}  //服务内参数值
	argsType  []reflect.Type //服务内参数类型，和上面要一一对应
	serviceID int64          //服务唯一id
	machineID int64          //机器唯一id
	initTime  time.Time      //上线时间
}

// ServiceInit 初始化一个服务，服务id和机器id后期会直接在控制中心分配
func ServiceInit(name string, args []interface{}, argstype []reflect.Type, serviceID int64, machineID int64) (*service, error) {
	l := len(args)
	//轮询检查每个变量和类型是否录入正确
	for i := 0; i < l; i++ {
		if reflect.TypeOf(args[i]) == argstype[i] {
			continue
		} else //todo:可以的话，先强制转换，不行再报err
		{
			return nil, fmt.Errorf("wrong correspondence between type and value")
		}
	}

	return &service{
		name:      name,
		args:      args,
		argsType:  argstype,
		serviceID: serviceID,
		machineID: machineID,
		initTime:  time.Now(),
	}, nil
}

func (s *service) Hello(in module.Info) module.Info {
	var out module.Info
	out.Args[1] = "hello" + fmt.Sprintf("%v", in.Args[1])
	out.TimeStamp = time.Now()
	return out
}
