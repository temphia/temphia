package plane

import (
	"github.com/tidwall/buntdb"
)

type Locker struct {
	db *buntdb.DB
}

func NewLocker() *Locker {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		panic(err)
	}

	return &Locker{
		db: db,
	}
}
