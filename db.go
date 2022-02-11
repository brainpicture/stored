package stored2

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
)

// DB is the main struct for handling work with fdb
type DB struct {
	fdb fdb.Database
}

// createDir would create or open an fdb directory using string name
func (db *DB) getDir(name string) (directory.DirectorySubspace, error) {
	return directory.CreateOrOpen(db.fdb, []string{name}, nil)
}

// KeyValue will create an key-value storage for low level data representation and storage
func (db *DB) KeyValue(name string) KeyValue {
	return KeyValue{
		db:        db,
		name:      name,
		directory: db.getDir(name),
	}
}

// drafts

func (db *DB) Get(obj any) any {
	
}
