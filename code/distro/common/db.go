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

func InitSQLiteDB(path string) error {

	sess, err := sqlite.NewUpperDb(path)
	if err != nil {
		return err
	}

	ok, err := sess.Collection("tenants").Exists()
	if err != nil {
		if !errors.Is(err, db.ErrCollectionDoesNotExist) {
			return err
		}
	}

	if ok {
		return nil
	}

	out := schema.SQLiteSchema

	conn := sess.Driver().(*sql.DB)

	ctx, cfunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cfunc()
	_, err = conn.ExecContext(ctx, string(out))
	return err

}
