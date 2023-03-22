package xplane

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

const (
	EventCreateTenant string = "create_tenant"
	EventUpdateTenant string = "update_tenant"
	EventDeleteTenant string = "delete_tenant"
)

type EventBus interface {
	EmitTenantEvent(tenant string, event string, data *entities.Tenant)
	EmitUserGroupEvent(event string, data *entities.UserGroup)
	EmitSchemaChange(tenant, source, group string, data any)

	OnDynSchemaChange(fn func(tenant, source, group string, data any))
	OnTenantChange(fn func(tenant string, event string, data *entities.Tenant))
	OnUserGroupChange(fn func(event string, data *entities.UserGroup))
}

type Message struct {
	Id     int64
	Topic  string
	Tenant string
	Path   string
	Data   string
}

type Subscription struct {
	Id   int32
	Chan chan<- Message
}

type MsgBus interface {
	Subscribe(topic string, ch chan Message) (int32, error)
	UnSubscribe(topic string, subid int32) error
	Submit(topic string, msg Message) (int64, error)
}
