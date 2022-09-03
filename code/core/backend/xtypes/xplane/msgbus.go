package xplane

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

const (
	EventCreateTenant string = "create_tenant"
	EventUpdateTenant string = "update_tenant"
	EventDeleteTenant string = "delete_tenant"
)

type EventBus interface {
	EmitTenantEvent(tenant string, event string, data *entities.Tenant)
	EmitUserGroupEvent(event string, data *entities.UserGroup)
	EmitSchemaChange(tenant, source, group string, data interface{})

	OnDynSchemaChange(fn func(tenant, source, group string, data interface{}))
	OnTenantChange(fn func(tenant string, event string, data *entities.Tenant))
	OnUserGroupChange(fn func(event string, data *entities.UserGroup))
}
