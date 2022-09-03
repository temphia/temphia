package sockdhub

import "github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"

const ROOM_SYSTABLE = "sys.dtable"
const ROOM_SYS_USERS = "sys.users"
const ROOM_PLUG_DEV = "plugs_dev"

const TAG_REALUSER = "sys.real_user"
const TAG_CONSOLE_CONN = "sys.console_conn"

type SockdHub struct {
	sockd sockdx.SockdCore
}

func New(sockd sockdx.SockdCore) *SockdHub {
	return &SockdHub{
		sockd: sockd,
	}
}
