package packer

import (
	"bufio"
	"encoding/binary"
)

type Integer interface {
	int | int64 | uint64 | int32 | uint32
}

func vInt[V Integer](writer bufio.Writer, v V) error {
	buf8 := make([]byte, 8)
	l := binary.PutUvarint(buf8, uint64(v))
	_, err := writer.Write(buf8[:l])
	return err
}
