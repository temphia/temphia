package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
	"gitlab.com/mr_balloon/golib/hmap"
)

// perm
func (d *DB) AddPerm(data *entities.Permission) error {
	return nil
}

func (d *DB) UpdatePerm(data hmap.H) error {
	return nil
}

func (d *DB) GetPerm(tenantId string, id int64) (*entities.Permission, error) {
	return nil, nil
}

func (d *DB) RemovePerm(data *entities.Permission) error {
	return nil
}

// role
func (d *DB) AddRole(data *entities.Role) error {
	return nil
}

func (d *DB) GetRole(tenantId string, id int64) (*entities.Role, error) {
	return nil, nil
}

func (d *DB) UpdateRole(data hmap.H) error {
	return nil
}

func (d *DB) RemoveRole(data *entities.Role) error {
	return nil
}

// user role
func (d *DB) AddUserRole(data *entities.UserRole) error {
	return nil
}

func (d *DB) RemoveUserRole(tenantId, user, role string) error {
	return nil
}

// query

func (d *DB) ListAllPerm(tenantId string) ([]*entities.Permission, error) {
	return nil, nil
}

func (d *DB) ListAllRole(tenantId string) ([]*entities.Permission, error) {
	return nil, nil
}

func (d *DB) ListAllUserRole(tenantId string) ([]*entities.Permission, error) {
	return nil, nil
}

func (d *DB) ListAllUserPerm(tenantId string) ([]*entities.Permission, error) {
	return nil, nil
}

func (d *DB) ListUserPerm(tenantId string, userId, objType, objsub string) ([]*entities.Permission, error) {
	roles := make([]*entities.UserRole, 0)
	err := d.table("user_roles").Find("tenant_id", tenantId, "user_id", userId).Select("role_id").All(&roles)
	if err != nil {
		return nil, err
	}

	roleIds := make([]string, 0, len(roles))
	for _, roleId := range roles {
		roleIds = append(roleIds, roleId.RoleId)
	}

	perms := make([]*entities.Permission, 0)

	err = d.table("user_perms").Find(db.Cond{
		"tenant_id":       tenantId,
		"object_type":     objType,
		"object_sub_type": objType,
		"role_id IN":      roleIds,
	}).All(&perms)

	return perms, err
}
