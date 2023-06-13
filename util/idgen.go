// @program:     rpc
// @file:        idgen.go
// @author:      ugug
// @create:      2023-06-12 11:54
// @description: id生成器，负责生成服务中产生的id，不负责服务id和机器id

package util

import (
	"fmt"
	"sync"
	"time"
)

//雪花版：|1|41时间|6服务号|4机器号|12序列号|=64

const (
	twepoch     int64 = 1686610911755 //2023-06-13 07:01:51.7552565 +0800 CST m=+0.002082101
	ServiceBit  int64 = 6             //自定义的服务号
	MachineBit  int64 = 4             //自定义的及机器号
	SequenceBit int64 = 12            //这个机器的这个服务上生成的id数

	MaxServiceID   int64 = (1 << ServiceBit) - 1 // 几个数值的上界
	MaxMachineID   int64 = (1 << MachineBit) - 1
	MaxSequenceNum int64 = (1 << SequenceBit) - 1

	TimeLMove    = uint(SequenceBit + MachineBit + SequenceBit) // 生成最终序号时，几个组成部分的左移位数
	ServiceLMove = uint(MachineBit + SequenceBit)
	MachineLMove = uint(SequenceBit)
)

//id生产者，考虑要不要和server整合为一个
type idProducePlace struct {
	sync.Mutex
	lastTimeStamp int64
	machineID     int64
	serviceID     int64
	sequence      int64
}

// IdProducePlaceInit 初始化一个id生产者，提供机器id和服务id
func IdProducePlaceInit(machineID int64, serviceID int64) (*idProducePlace, error) {
	//创建之前先检测两个ID是否合法
	if serviceID > MaxServiceID || serviceID < 0 {
		return nil, fmt.Errorf("invalid service id")
	}
	if machineID > MaxMachineID || machineID < 0 {
		return nil, fmt.Errorf("invalid machine id")
	}

	return &idProducePlace{
		lastTimeStamp: -1,
		machineID:     machineID,
		serviceID:     serviceID,
		sequence:      0,
	}, nil
}

//实现雪花生成id
func (p *idProducePlace) snowFlakeGenerateID() error {
	p.Lock() //并发锁
	defer p.Unlock()

	var presentTimeStamp = time.Now().UnixMilli()

	//检测是否有时钟回调
	if presentTimeStamp < p.lastTimeStamp {
		return fmt.Errorf("invaid time,check the clock")
		//此else为正常情况
	} else if presentTimeStamp == p.lastTimeStamp {
		p.sequence = (p.sequence + 1) & MaxSequenceNum
		//1ms内生产的id达到上限，进行推迟
		if p.sequence == 0 {
			presentTimeStamp = nextTime(p.lastTimeStamp)
		}
		//时间戳改变计数重置
	} else {
		p.sequence = 0
	}
	p.lastTimeStamp = presentTimeStamp

	id := (presentTimeStamp-twepoch)<<TimeLMove |
		p.serviceID<<ServiceLMove |
		p.machineID<<MachineLMove |
		p.sequence
	fmt.Println(id)

	return nil
}

func nextTime(currentTime int64) int64 {
	newTime := time.Now().UnixMilli()
	for newTime <= currentTime {
		newTime = time.Now().UnixMilli()
	}
	return newTime
}
