package coredb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

// auth

func (d *DB) AddUserGroupAuth(data *entities.UserGroupAuth) error {
	_, err := d.userGroupAuth().Insert(data)
	return err
}

func (d *DB) UpdateUserGroupAuth(tenantId string, gslug string, id int64, data map[string]interface{}) error {
	return d.userGroupAuth().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Update(data)
}

func (d *DB) ListUserGroupAuth(tenantId string, gslug string) ([]*entities.UserGroupAuth, error) {
	data := make([]*entities.UserGroupAuth, 0)
	err := d.userGroupAuth().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) GetUserGroupAuth(tenantId string, gslug string, id int64) (*entities.UserGroupAuth, error) {
	data := &entities.UserGroupAuth{}
	err := d.userGroupAuth().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) RemoveUserGroupAuth(tenantId, gslug string, id int64) error {
	return d.userGroupAuth().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Delete()
}

// hook

func (d *DB) AddUserGroupHook(data *entities.UserGroupHook) error {
	_, err := d.userGroupHook().Insert(data)
	return err
}

func (d *DB) UpdateUserGroupHook(tenantId string, gslug string, id int64, data map[string]interface{}) error {
	return d.userGroupHook().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Update(data)
}

func (d *DB) ListUserGroupHook(tenantId string, gslug string) ([]*entities.UserGroupHook, error) {
	data := make([]*entities.UserGroupHook, 0)
	err := d.userGroupHook().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) GetUserGroupHook(tenantId string, gslug string, id int64) (*entities.UserGroupHook, error) {
	data := &entities.UserGroupHook{}

	err := d.userGroupHook().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) RemoveUserGroupHook(tenantId, gslug string, id int64) error {
	return d.userGroupHook().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Delete()
}

// plug

func (d *DB) AddUserGroupPlug(data *entities.UserGroupPlug) error {
	_, err := d.userGroupPlug().Insert(data)
	return err
}

func (d *DB) UpdateUserGroupPlug(tenantId string, gslug string, id int64, data map[string]interface{}) error {
	return d.userGroupPlug().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Update(data)
}

func (d *DB) ListUserGroupPlug(tenantId string, gslug string) ([]*entities.UserGroupPlug, error) {
	data := make([]*entities.UserGroupPlug, 0)
	err := d.userGroupPlug().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) GetUserGroupPlug(tenantId string, gslug string, id int64) (*entities.UserGroupPlug, error) {
	data := &entities.UserGroupPlug{}

	err := d.userGroupPlug().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) RemoveUserGroupPlug(tenantId, gslug string, id int64) error {
	return d.userGroupPlug().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Delete()
}

// data

func (d *DB) AddUserGroupData(data *entities.UserGroupData) error {
	_, err := d.userGroupData().Insert(data)
	return err
}

func (d *DB) UpdateUserGroupData(tenantId string, gslug string, id int64, data map[string]interface{}) error {
	return d.userGroupData().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Update(data)
}

func (d *DB) ListUserGroupData(tenantId string, gslug string) ([]*entities.UserGroupData, error) {
	data := make([]*entities.UserGroupData, 0)
	err := d.userGroupData().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) GetUserGroupData(tenantId string, gslug string, id int64) (*entities.UserGroupData, error) {
	data := &entities.UserGroupData{}
	err := d.userGroupData().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) RemoveUserGroupData(tenantId, gslug string, id int64) error {
	return d.userGroupData().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_group": gslug,
		"id":         id,
	}).Delete()
}
