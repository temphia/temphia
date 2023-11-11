package sqlite

import (
	"database/sql"
	"path"

	"github.com/k0kubun/pp"
	"github.com/mattn/go-sqlite3"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upperdb"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

func init() {
	registry.SetStoreBuilders(store.VendorSqlite, NewVendor)
}

func NewVendor(opts store.BuilderOptions) (store.Store, error) {
	return upperdb.NewAdapter(&sl{})(opts)
}

type sl struct{}

var csess db.Session

func (sl) Db(so *config.StoreConfig) (db.Session, error) {
	if csess != nil {
		return csess, nil
	}

	return NewUpperDb(so.Target)
}

func (sl) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return sqlite.NewTx(sqlTx)
}

func NewUpperDb(p string) (db.Session, error) {
	db, err := NewRawDB(path.Join(p, "main.db"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

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
		Options: map[string]string{
			"_journal_mode": "WAL",
		},
	}

	dvr := &sqlite3.SQLiteDriver{
		Extensions: []string{},
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			pp.Println("@connect")
			return conn.RegisterFunc("temphia_delete_record", temphiaDeleteRecord(conn), false)
		},
	}

	sql.Register("sqlite3_temphia", dvr)
	surl := settings.String()
	pp.Println("@connecting_to", surl)
	return sql.Open("sqlite3_temphia", surl)
}
