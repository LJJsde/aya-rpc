// @program:     rpc
// @file:        code.go
// @author:      ugug
// @create:      2023-06-11 05:27
// @description:

package Protocol

import (
	"bytes"
	"encoding/gob"
	"rpc/module"
)

// CodeProcessor (暂定）序列化接口
type CodeProcessor interface {
	Decode([]byte, interface{}) error
	Encode(interface{}) (error, []byte)
}

type CodeProcess struct {
	Processor CodeProcessor
}

//TODO：自己写个JSON序列化版本加上（虽然现在这个麻烦点好像也行

func (c *CodeProcess) Decode(b []byte) (error, module.Info) {
	var i module.Info
	d := gob.NewDecoder(bytes.NewReader(b))
	if err := d.Decode(i); err != nil {
		return err, i
	}
	return nil, i
}

func (c CodeProcess) Encode(i module.Info) (error, []byte) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	if err := e.Encode(i); err != nil {
		return err, nil
	}
	return nil, b.Bytes()
}
