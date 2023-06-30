package runner

import "net"

type bindingsLine struct {
	runner *Runner
	closed bool
	conn   net.Conn
	id     string
}

func (c *bindingsLine) run() error {

	return nil
}
