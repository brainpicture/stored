package stored

// DB is the main struct for handling work with fdb
// type DB struct {
// 	fdb fdb.Database
// }

// createDir would create or open an fdb directory using string name
// func (db *DB) getDir(name string) (directory.DirectorySubspace, error) {
// 	return directory.CreateOrOpen(db.fdb, []string{name}, nil)
// }

// Connect initialising connection to the database
/*func (db *DB) Object1Key[T any](objID ObjectID, obj T) Object[T] {
	return Object[T]{
		ID: objID,
	}
}*/

type ObjectID int
