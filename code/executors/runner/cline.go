package runner

import (
	"encoding/json"
	"net"

	"github.com/temphia/temphia/code/executors/runner/rtypes"
)

type controlLine struct {
	runner *Runner
	closed bool
	conn   net.Conn
}

func (r *controlLine) run() error {

	defer func() {
		r.conn.Close()
		r.closed = false
	}()

	go r.writeLoop()

	r.readLoop()

	return nil

}

func (r *controlLine) readLoop() {

	for {
		decoder := json.NewDecoder(r.conn)
		packet := &rtypes.Packet{}

		err := decoder.Decode(packet)
		if err != nil {
			return
		}

	}
}

func (r *controlLine) writeLoop() {

}
