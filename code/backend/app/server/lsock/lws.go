package lsock

import (
	"net"

	"github.com/temphia/temphia/code/backend/xtypes/xserver"
)

type lsockWs struct {
	parent *LSock
	sub    xserver.LSubcriber
	conn   net.Conn
}
