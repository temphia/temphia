package ucore

import (
	"database/sql"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/upper/db/v4"
)

type UpperVendor interface {
	Db(conf *config.StoreSource) (db.Session, error)
	NewTx(sqlTx *sql.Tx) (dbutils.Tx, error)
}

type DynDBOptions struct {
	Session    db.Session
	SharedLock service.DyndbLock
	TxnManager dbutils.TxManager
	DynGen     Zenerator
	TNS        tns.TNS
}

// upper throws timeout when it takes long to run query so get
// underlying driver and execute query directly
func GetDriver(sess db.Session) *sql.DB {
	return sess.Driver().(*sql.DB)
}
