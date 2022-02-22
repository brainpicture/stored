package stored

import (
	"github.com/brainfucker/stored/packer"
	"github.com/brainfucker/stored/types"
)

type Int = types.Int
type String = types.String
type Pack = packer.Pack
type Unpack = packer.Unpack

type Object interface {
	Key(p Pack)
	Pack(p Pack)
	Unpack(p Unpack)
}
