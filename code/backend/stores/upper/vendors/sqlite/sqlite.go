package sqlite

import (
	"database/sql"

	"github.com/k0kubun/pp"
	"github.com/mattn/go-sqlite3"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

func init() {
	registry.SetStoreBuilders(store.VendorSqlite, func(so *config.StoreSource) (store.Store, error) {
		return upper.NewAdapter(&sl{})(so)
	})
}

type sl struct{}

func (sl) Db(so *config.StoreSource) (db.Session, error) {

	var settings = sqlite.ConnectionURL{
		Database: so.HostPath,
	}

	dvr := &sqlite3.SQLiteDriver{
		Extensions: []string{},
		ConnectHook: func(sc *sqlite3.SQLiteConn) error {
			pp.Println("@connect")
			return nil
		},
	}

	sql.Register("sqlite3_temphia", dvr)

	db, err := sql.Open("sqlite3_temphia", settings.String())
	if err != nil {
		return nil, err
	}

	return sqlite.New(db)
}

func (sl) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return sqlite.NewTx(sqlTx)
}
