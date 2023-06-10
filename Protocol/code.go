package Protocol

import (
	"bytes"
	"encoding/gob"
)

// CodeProcessor (暂定）序列化接口
type CodeProcessor interface {
	Decode([]byte, interface{}) error
	Encode(interface{}) (error, []byte)
}

type CodeProcess struct{}

func (c *CodeProcess) Decode(b []byte, in interface{}) (error, out interface{}) {
	d := gob.NewDecoder(bytes.NewReader(b))
	if err := d.Decode(b); err != nil {
		return err, nil
	}
	return nil, out
}

func (c CodeProcess) Encode(i interface{}) (error, []byte) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	if err := e.Encode(i); err != nil {
		return err, nil
	}
	return nil, b.Bytes()
}
