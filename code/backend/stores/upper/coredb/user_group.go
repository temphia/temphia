package coredb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddUserGroup(group *entities.UserGroup) error {
	_, err := d.userGroupTable().Insert(group)
	return err
}

func (d *DB) UpdateUserGroup(tenantId, slug string, data map[string]interface{}) error {
	delete(data, "slug")
	return d.userGroupTable().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      slug,
	}).Update(&data)
}

func (d *DB) RemoveUserGroup(tenantId string, slug string) error {

	// fixme => cascade ?

	return d.userGroupTable().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      slug,
	}).Delete()
}

func (d *DB) GetUserGroup(tenantId string, slug string) (*entities.UserGroup, error) {
	uGrp := &entities.UserGroup{}

	err := d.userGroupTable().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      slug,
	}).One(uGrp)
	if err != nil {
		return nil, err
	}
	return uGrp, nil
}

func (d *DB) ListUserGroups(tenantId string) ([]*entities.UserGroup, error) {
	groups := make([]*entities.UserGroup, 0)
	err := d.userGroupTable().Find("tenant_id", tenantId).All(&groups)
	return groups, err
}
