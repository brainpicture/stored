package stored

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

// Connect is main constructor for creating connections
// example: db := stored.Connect("./fdb.cluster")
func Connect(clusterFilePath string) *DB {
	fdb.MustAPIVersion(510)
	conn := DB{
		fdb: fdb.MustOpen(clusterFilePath, []byte("DB")),
	}
	return &conn
}
