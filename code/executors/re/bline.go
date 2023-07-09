package re

import (
	"encoding/json"
	"net"

	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/executors/re/bind2ipc"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

// FIXME => bindingsLine could leak/open forever even event is processed
// its better not to dependent on remote exec process to close conn properly

type bindingsLine struct {
	runner *Runner
	closed bool
	conn   net.Conn
	id     string

	writeChan chan struct {
		isErr bool
		id    string
		data  []byte
	}

	bipc bind2ipc.Bind2IPC
}

func (c *bindingsLine) run() {

	go c.readLoop()

	c.writeLoop()
}

func (c *bindingsLine) readLoop() {
	mipc := modipc.NewModIPC(&c.bipc)

	for {

		if c.closed {
			return
		}

		decoder := json.NewDecoder(c.conn)
		packet := &rtypes.Packet{}

		err := decoder.Decode(packet)
		if err != nil {
			c.closed = true
			close(c.writeChan)
			return
		}

		go func() {

			senErr := func(err error) {
				c.writeChan <- struct {
					isErr bool
					id    string
					data  []byte
				}{
					isErr: true,
					data:  []byte(err.Error()),
				}
			}

			resp, err := mipc.Handle(packet.Type, lazydata.NewJsonData(packet.Data))
			if err != nil {
				senErr(err)
				return
			}

			out, err := resp.AsJsonBytes()
			if err != nil {
				senErr(err)
				return
			}

			c.writeChan <- struct {
				isErr bool
				id    string
				data  []byte
			}{
				isErr: false,
				data:  out,
			}
		}()

	}

}

func (c *bindingsLine) writeLoop() {

	for {

		if c.closed {
			return
		}

		wmsg, ok := <-c.writeChan
		if !ok {
			c.closed = true
			c.conn.Close()
			return
		}

		packet := &rtypes.Packet{
			Id:   wmsg.id,
			Data: wmsg.data,
		}

		out, err := json.Marshal(packet)
		if err != nil {
			continue
		}

		c.conn.Write(out)
		c.conn.Write([]byte("\n"))

	}

}
