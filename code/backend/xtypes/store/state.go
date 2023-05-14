package store

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type StateHub interface {
	OnTargetAppChange(tenantId string, id int64, data *entities.TargetApp) error
	OnTargetHookChange(tenantId string, id int64, data *entities.TargetHook) error
	OnResourceChange(tenantId, id string, data *entities.Resource) error
	OnUserGroupChange(tenantId, id string, data *entities.UserGroup) error

	OnDataGroupChange(tenantId, gid string, data *entities.TableGroup) error
	OnDataTableChange(tenantId, gid, tid string, data *entities.Table) error
	OnDataColumnChange(tenantId, gid, tid, cid string, data *entities.Column) error

	OnTenantChange(id string, data *entities.Tenant) error
	OnDomainChange(tenantId string, id int64, data *entities.TenantDomain) error
}
