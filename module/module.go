// @program:     rpc
// @file:        module.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package module

import (
	"time"
)

// Info 客户端送出和服务端收到的，非二进制的信息的结构体
type Info struct {
	//T表示服务端函数执行错误，F表示正常
	ErrCode bool
	//放置具体方法返回的错误信息
	ErrMessage string
	TimeStamp  time.Time
	//喜欢我interface切片小王子吗
	Method []interface{}
	Args   []interface{}
}
