package distro

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/postfinance/single"
	"github.com/upper/db/v4"

	"github.com/temphia/temphia/code/backend/app/config"

	"github.com/temphia/temphia/code/backend/xtypes"

	"github.com/temphia/temphia/code/backend/stores/upperdb/schema"
	"github.com/temphia/temphia/code/backend/stores/upperdb/vendors/sqlite"
)

func (d *DistroApp) Run() error {
	return d.app.Run()
}

func (d *DistroApp) GetApp() xtypes.App {
	return d.app
}

func (d *DistroApp) Configd() config.Confd {
	return d.confd
}

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

func IsRunning(folder string) bool {

	s, err := single.New("instance", single.WithLockPath(folder))
	if err != nil {
		panic(err)
	}

	err = s.Lock()
	if err != nil {
		if errors.Is(err, single.ErrAlreadyRunning) {
			return true
		}

		panic(err)
	}

	defer s.Unlock()

	return false

}
