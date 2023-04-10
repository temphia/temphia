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

var csess db.Session

func (sl) Db(so *config.StoreSource) (db.Session, error) {
	if csess != nil {
		return csess, nil
	}

	return NewUpperDb(so.HostPath)
}

func (sl) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return sqlite.NewTx(sqlTx)
}

func NewUpperDb(path string) (db.Session, error) {
	db, err := NewRawDB(path)
	if err != nil {
		return nil, err
	}

	sess, err := sqlite.New(db)
	if err != nil {
		return nil, err
	}

	csess = sess
	return sess, nil

}

func NewRawDB(path string) (*sql.DB, error) {

	var settings = sqlite.ConnectionURL{
		Database: path,
	}

	dvr := &sqlite3.SQLiteDriver{
		Extensions: []string{},
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			pp.Println("@connect")
			return conn.RegisterFunc("temphia_delete_record", temphiaDeleteRecord(conn), false)
		},
	}

	sql.Register("sqlite3_temphia", dvr)
	return sql.Open("sqlite3_temphia", settings.String())
}
