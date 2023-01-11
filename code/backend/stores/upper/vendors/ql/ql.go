package ql

import (
	"database/sql"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/ql"
)

func init() {
	registry.SetStoreBuilders(store.VendorQL, func(so *config.StoreSource) (store.Store, error) {
		return upper.NewAdapter(&qldb{})(so)
	})
}

type qldb struct{}

func (qldb) Db(conf *config.StoreSource) (db.Session, error) {
	var settings = ql.ConnectionURL{
		Database: conf.Name,
	}

	return ql.Open(settings)
}

func (qldb) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return ql.NewTx(sqlTx)
}

func (qldb) CoreSchema() string {
	return ""
}
func (qldb) DtableSchema() string {
	return ""
}
