package coredb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddUserDevice(tenantId string, user string, data *entities.UserDevice) error {
	_, err := d.userDevices().Insert(data)
	return err
}

func (d *DB) UpdateUserDevice(tenantId string, user string, id int64, data map[string]any) error {

	err := d.userDevices().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   user,
		"id":        id,
	}).Update(data)

	return err
}

func (d *DB) GetUserDevice(tenantId string, user string, id int64) (*entities.UserDevice, error) {

	data := &entities.UserDevice{}

	err := d.userDevices().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   user,
		"id":        id,
	}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) ListUserDevice(tenantId string, user string) ([]*entities.UserDevice, error) {

	data := make([]*entities.UserDevice, 0)
	err := d.userDevices().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   user,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) RemoveUserDevice(tenantId string, user string, id int64) error {

	return d.userDevices().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   user,
		"id":        id,
	}).Delete()

}
