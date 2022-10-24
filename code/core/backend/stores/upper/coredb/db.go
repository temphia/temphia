package coredb

import (
	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/temphia/temphia/code/core/backend/stores/upper/ucore"
	"github.com/thoas/go-funk"
	"github.com/upper/db/v4"
)

type DB struct {
	session          db.Session
	returningSupport bool
	vendor           string
	txn              dbutils.TxManager
}

func New(session db.Session, vendor string) *DB {
	db := &DB{
		session:          session,
		returningSupport: false,
		vendor:           vendor,
		txn:              dbutils.TxManager{}, // fixme => send txbuilder
	}

	return db
}

func (d *DB) Execute(qstr string) error {
	return dbutils.Execute(ucore.GetDriver(d.session), qstr)
}

func (d *DB) Migrate() error {
	return nil
}

func (d *DB) GetInnerDriver() interface{} {
	return d.session
}

func (d *DB) table(name string) db.Collection {
	return d.session.Collection(name)
}

func (d *DB) userGroupTable() db.Collection {
	return d.table("user_groups")
}

func (d *DB) userGroupAuth() db.Collection {
	return d.table("user_group_auths")
}

func (d *DB) userGroupData() db.Collection {
	return d.table("user_group_datas")
}

func (d *DB) userTable() db.Collection {
	return d.table("users")
}

func (d *DB) userDataTable() db.Collection {
	return d.table("user_datas")
}

func (d *DB) userMessagesTable() db.Collection {
	return d.table("user_messages")
}

// func blueprintBlobsTable(sess db.Session) db.Collection {
// 	return dbutils.Table(sess, "blueprint_blobs")
// }

func only(data map[string]interface{}, keys ...string) bool {
	if len(data) > len(keys) {
		return false
	}

	for hkey := range data {
		if !funk.ContainsString(keys, hkey) {
			return false
		}
	}
	return true
}

func (d *DB) Ping() error {
	return d.session.Ping()
}
