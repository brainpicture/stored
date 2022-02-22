package types

import (
	"bytes"
)

type Int struct {
	source *int64
}

func (i Int) Pack(buf *bytes.Buffer) error {
	_, err := buf.Write([]byte("df"))
	return err
}
