package packer

import (
	"bufio"
	"encoding/binary"
)

var endian = binary.LittleEndian

// Pack is a basic data packer into string
type Pack struct {
	version int
	writer  bufio.Writer
}

func (p *Pack) Version(v int) {
	p.version = v
}

func (p *Pack) Int(i int) error {
	return vInt(p.writer, i)
}

func (p *Pack) String(s string) error {
	err := p.Int(len(s))
	if err != nil {
		return err
	}
	_, err = p.writer.Write([]byte(s))
	return err
}
