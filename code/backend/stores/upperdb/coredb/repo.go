package coredb

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) RepoNew(tenantId string, data *entities.Repo) error {
	_, err := d.repoTable().Insert(data)
	return err
}

func (d *DB) RepoUpdate(tenantId string, id int64, data map[string]interface{}) error {
	return d.repoTable().Find(db.Cond{
		"id":        id,
		"tenant_id": tenantId,
	}).Update(data)
}

func (d *DB) RepoGet(tenantId string, id int64) (*entities.Repo, error) {
	data := &entities.Repo{}

	err := d.repoTable().Find(db.Cond{
		"id":        id,
		"tenant_id": tenantId,
	}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) RepoDel(tenantId string, id int64) error {
	return d.repoTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Delete()
}

func (d *DB) RepoList(tenantId string) ([]*entities.Repo, error) {
	datas := make([]*entities.Repo, 0)

	err := d.repoTable().Find(db.Cond{"tenant_id": tenantId}).All(&datas)
	if err != nil {
		return nil, nil
	}

	return datas, nil
}

// private

func (d *DB) repoTable() db.Collection {
	return dbutils.Table(d.session, "repos")
}
