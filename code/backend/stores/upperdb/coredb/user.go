package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddUser(user *entities.User, data *entities.UserData) error {
	u, err := d.userTable().Insert(user)
	if err != nil {
		return err
	}

	_, err = d.userDataTable().Insert(data)

	if err != nil {
		d.userTable().Find(u).Delete()
	}

	return err
}

func (d *DB) UpdateUser(tenantId, user string, data map[string]interface{}) error {
	return d.userTable().Find("tenant_id", tenantId, "user_id", user).Update(data)
}

func (d *DB) RemoveUser(tenantId string, username string) error {
	d.userDataTable().Find(db.Cond{"tenant_id": tenantId, "user_id": username}).Delete()
	return d.userTable().Find(db.Cond{"tenant_id": tenantId, "user_id": username}).Delete()
}

func (d *DB) GetUserByID(tenantId string, username string) (*entities.User, error) {
	usr := &entities.User{}
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"user_id":   username,
		},
	).One(usr)
	return usr, err
}

func (d *DB) GetUserByEmail(tenantId string, email string) (*entities.User, error) {
	usr := &entities.User{}
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"email":     email,
		},
	).One(usr)

	return usr, err
}

func (d *DB) ListUsers(tenantId string) ([]*entities.User, error) {
	us := make([]*entities.User, 0, 10)
	err := d.userTable().Find(
		db.Cond{
			"tenant_id": tenantId,
		},
	).All(&us)

	return us, err
}

func (d *DB) ListUsersByGroup(tenantId string, groupid string) ([]*entities.User, error) {
	us := make([]*entities.User, 0)

	cond := db.Cond{
		"tenant_id": tenantId,
		"group_id":  groupid,
	}

	err := d.userTable().Find(cond).All(&us)
	return us, err
}

func (d *DB) GetUserData(tenantId string, slug string) (*entities.UserData, error) {
	data := &entities.UserData{}

	err := d.userDataTable().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   slug,
	}).One(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) UpdateUserData(tenantId, slug string, data map[string]interface{}) error {

	return d.userDataTable().Find(db.Cond{
		"tenant_id": tenantId,
		"user_id":   slug,
	}).Update(data)
}

func (d *DB) ListUsersMulti(tenantId string, users ...string) ([]*entities.User, error) {
	uids := make([]*entities.User, 0)

	err := d.userDataTable().Find(db.Cond{
		"tenant_id":  tenantId,
		"user_id IN": users,
	}).All(&uids)
	if err != nil {
		return nil, err
	}
	return uids, nil
}
