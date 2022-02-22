package types

import "bytes"

type Storable interface {
	Pack(but *bytes.Buffer)
}
