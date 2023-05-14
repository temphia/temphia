package store

type StateHub interface {
	OnTargetAppChange(tenantId string, id int64, data any)
	OnTargetHookChange(tenantId string, id int64, data any)
	OnResourceChange(tenantId, id string, data any)
	OnUserGroupChange(tenantId, id string, data any)

	OnDataGroupChange(tenantId, gid string, data any)
	OnDataTableChange(tenantId, gid, tid string, data any)
	OnDataColumnChange(tenantId, gid, tid, cid string, data any)

	OnTenantChange(id string, data any)
	OnDomainChange(tenantId string, id int64, data any)
}
