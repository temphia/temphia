package plane

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type EventBus struct {
	//natsConn               *nats.Conn
	tenantEventHandlers    []func(tenant string, event string, data *entities.Tenant)
	usergroupEventHandlers []func(event string, data *entities.UserGroup)
	schemaEventHandlers    []func(tenant, source, group string, data interface{})
}

func NewEventBus() *EventBus {

	return &EventBus{
		tenantEventHandlers:    make([]func(tenant string, event string, data *entities.Tenant), 0),
		usergroupEventHandlers: make([]func(event string, data *entities.UserGroup), 0),
		schemaEventHandlers:    make([]func(tenant string, source string, group string, data interface{}), 0),
	}
}

func (eb *EventBus) EmitTenantEvent(tenant string, event string, data *entities.Tenant) {
	for _, fn := range eb.tenantEventHandlers {
		fn(tenant, event, data)
	}
}

func (eb *EventBus) EmitUserGroupEvent(event string, data *entities.UserGroup) {
	for _, fn := range eb.usergroupEventHandlers {
		fn(event, data)
	}
}

func (eb *EventBus) EmitSchemaChange(tenant, source, group string, data interface{}) {
	for _, fn := range eb.schemaEventHandlers {
		fn(tenant, source, group, data)
	}
}

func (eb *EventBus) OnDynSchemaChange(fn func(tenant, source, group string, data interface{})) {
	eb.schemaEventHandlers = append(eb.schemaEventHandlers, fn)
}

func (eb *EventBus) OnTenantChange(fn func(tenant string, event string, data *entities.Tenant)) {
	eb.tenantEventHandlers = append(eb.tenantEventHandlers, fn)
}

func (eb *EventBus) OnUserGroupChange(fn func(event string, data *entities.UserGroup)) {
	eb.usergroupEventHandlers = append(eb.usergroupEventHandlers, fn)
}
