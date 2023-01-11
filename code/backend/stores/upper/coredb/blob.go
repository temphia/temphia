package coredb

import (
	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/upper/db/v4"
)

func blobRefTable(sess db.Session) db.Collection {
	return dbutils.Table(sess, "blobs_references")
}

func (d *DB) AddBlob(sess db.Session, tenantId string, hash string) error {
	_, err := blobRefTable(sess).Insert(db.Cond{
		"tenant_id": tenantId,
		"blob_hash": hash,
	})
	return err
}

func (d *DB) RemoveBlob(sess db.Session, tenantId string, hash string) error {
	return blobRefTable(sess).Find(db.Cond{
		"tenant_id": tenantId,
		"blob_hash": hash,
	}).Delete()
}
