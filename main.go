package stored

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

var db fdb.Database

// New is main constructor for creating connections
// example: db := stored.Connect("./fdb.cluster")
func Connect(clusterFilePath string) error {
	fdb.MustAPIVersion(510)
	fdb, err := fdb.Open(clusterFilePath, []byte("DB"))
	if err != nil {
		return err
	}
	db = fdb
	return nil
}
