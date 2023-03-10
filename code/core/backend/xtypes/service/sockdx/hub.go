package sockdx

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

const ROOM_SYSTABLE = "sys.dtable"
const ROOM_SYS_USERS = "sys.users"
const ROOM_PLUG_DEV = "plugs_dev"

const TAG_REALUSER = "sys.real_user"
const TAG_CONSOLE_CONN = "sys.console_conn"

type Hub interface {
	GetSockd() SockdCore
	GetDataSyncer() DataSyncer
	GetUserSyncer() UserSyncer
}

type DataSyncer interface {
	PushNewRow(source, tenantId, groupId, table string, data map[string]any) error
	PushUpdateRow(source, tenantId, groupId, table string, id int64, data map[string]any) error
	PushDeleteRow(source, tenantId, groupId, table string, id int64) error
}

type UserSyncer interface {
	NotifyUser(msg *entities.UserMessage) error
	NotifyMessageRead(tenantId, user string, msgIds []int) error
	NotifyMessageDelete(tenantId, user string, msgIds []int) error
}
