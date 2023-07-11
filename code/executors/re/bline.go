package re

import (
	"bufio"
	"encoding/json"
	"net"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
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

	ctxipc  *modipc.ModIPC // 		bipc = modipc.NewModIPC(bind2ipc.New(c.runner.rootBinding))
	rootipc *modipc.ModIPC
}

func (c *bindingsLine) run() {

	go c.readLoop()

	c.writeLoop()
}

func (c *bindingsLine) readLoop() {

	reader := bufio.NewReader(c.conn)

	for {

		if c.closed {
			return
		}

		out, err := reader.ReadBytes('\n')
		if err != nil {
			pp.Println("@read_error", err.Error())
			continue
		}

		go c.handle(out)
	}

}

func (c *bindingsLine) handle(data []byte) {

	packet := &rtypes.Packet{}

	err := json.Unmarshal(data, packet)
	if err != nil {
		c.closed = true
		close(c.writeChan)
		return
	}

	sendErr := func(err error) {
		c.writeChan <- struct {
			isErr bool
			id    string
			data  []byte
		}{
			isErr: true,
			data:  []byte(err.Error()),
		}
	}

	sendOk := func(out []byte) {
		c.writeChan <- struct {
			isErr bool
			id    string
			data  []byte
		}{
			isErr: false,
			data:  out,
		}

	}

	bipc := c.ctxipc

	switch packet.Type {
	case rtypes.ROOT_BINDING_CALL:
		bipc = c.rootipc
		fallthrough

	case rtypes.CTX_BINDING_CALL:
		resp, err := bipc.Handle(packet.Type, lazydata.NewJsonData(packet.Data))
		if err != nil {
			sendErr(err)
			return
		}

		out, err := resp.AsJsonBytes()
		if err != nil {
			sendErr(err)
			return
		}
		sendOk(out)

	default:
		pp.Println("@bline_not_impl_packet", packet)
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
