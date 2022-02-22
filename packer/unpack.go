package packer

import (
	"bufio"
	"encoding/binary"
)

type Unpack struct {
	Version int
	reader  *bufio.Reader
}

func (p *Unpack) Int(i *int) error {
	data, err := binary.ReadUvarint(p.reader)
	*i = int(data)
	return err
}

func (p *Unpack) String(s *string) error {
	var len int
	err := p.Int(&len)
	if err != nil {
		return err
	}
	return err
}
