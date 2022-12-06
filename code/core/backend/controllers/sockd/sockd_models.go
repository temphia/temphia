package sockd

import "github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"

type PlugConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Plug     string
	Conn     sockdx.Conn
}

type UserConnOptions struct {
	TenantId string
	UserId   string
	GroupId  string
	DeviceId string
	Conn     sockdx.Conn
}

type UpdateDynRoomTagsOptions struct {
	TenantId  string
	DynSource string
	DynGroup  string
	ConnId    int64
}

type DevConnOptions struct {
	TenantId string
	UserId   string
	PlugId   string
	AgentId  string
	Conn     sockdx.Conn
}

type DataConnOptions struct {
	TenantId  string
	UserId    string
	DynSource string
	DynGroup  string
	Conn      sockdx.Conn
}
