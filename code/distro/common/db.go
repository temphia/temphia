package common

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/temphia/temphia/code/backend/stores/upperdb/schema"
	"github.com/temphia/temphia/code/backend/stores/upperdb/vendors/sqlite"
	"github.com/upper/db/v4"
)

func InitSQLiteDB(path string) (bool, error) {

	sess, err := sqlite.NewUpperDb(path)
	if err != nil {
		return false, err
	}

	ok, err := sess.Collection("system_events").Exists()
	if err != nil {
		if !errors.Is(err, db.ErrCollectionDoesNotExist) {
			return false, err
		}
	}

	if ok {
		return false, nil
	}

	out := schema.SQLiteSchema

	conn := sess.Driver().(*sql.DB)

	ctx, cfunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cfunc()
	_, err = conn.ExecContext(ctx, string(out))
	if err != nil {
		return false, err
	}

	return true, nil
}
