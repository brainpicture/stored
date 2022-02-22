package types

import (
	"bytes"
)
type String string

func (s String) Pack(buf *bytes.Buffer) error {
	l := Int(len(s))
	l.Pack(buf)
	_,err := buf.WriteString(string(s))
	return err
}


