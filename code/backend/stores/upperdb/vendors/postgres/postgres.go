package postgres

import (
	"database/sql"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upperdb"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func init() {
	registry.SetStoreBuilders(store.VendorPostgres, NewVendor)
}

func NewVendor(opts store.BuilderOptions) (store.Store, error) {
	return upperdb.NewAdapter(&pq{})(opts)
}

type pq struct{}

func (pq) Db(conf *config.StoreConfig) (db.Session, error) {
	var settings = postgresql.ConnectionURL{
		Database: conf.Target,
		User:     conf.User,
		Password: conf.Password,
		Socket:   "",
		Options:  make(map[string]string),
		Host:     conf.HostPath + ":" + conf.Port,
	}

	return postgresql.Open(settings)
}

func (pq) NewTx(sqlTx *sql.Tx) (dbutils.Tx, error) {
	return postgresql.NewTx(sqlTx)
}
