package store

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type StateHub interface {
	OnTargetAppChange(tenantId string, id int64, data *entities.TargetApp)
	OnTargetHookChange(tenantId string, id int64, data *entities.TargetHook)
	OnResourceChange(tenantId, id string, data *entities.Resource)
	OnUserGroupChange(tenantId, id string, data *entities.UserGroup)

	OnDataGroupChange(tenantId, gid string, data *entities.TableGroup)
	OnDataTableChange(tenantId, gid, tid string, data *entities.Table)
	OnDataColumnChange(tenantId, gid, tid, cid string, data *entities.Column)

	OnTenantChange(id string, data *entities.Tenant)
	OnDomainChange(tenantId string, id int64, data *entities.TenantDomain)
}
