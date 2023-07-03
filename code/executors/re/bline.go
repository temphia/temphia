package re

import (
	"net"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type bindingsLine struct {
	runner *Runner
	closed bool
	conn   net.Conn
	id     string

	binder bindx.Bindings
}

func (c *bindingsLine) run() error {

	return nil
}
