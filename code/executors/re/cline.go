package re

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type controlLine struct {
	runner *Runner
	closed bool
	conn   net.Conn

	eventChan chan struct {
		resp       chan []byte
		id         string
		packetData []byte
	}

	respChan chan struct {
		id   string
		data []byte
	}

	pendingEvents map[string]chan []byte
}

func (r *controlLine) run() {

	defer func() {
		r.conn.Close()
		r.closed = false
	}()

	go r.writeLoop()

	r.readLoop()

}

func (r *controlLine) readLoop() {

	for {
		decoder := json.NewDecoder(r.conn)
		packet := &rtypes.Packet{}

		err := decoder.Decode(packet)
		if err != nil {
			return
		}

		if packet.Type != rtypes.EVENT_RESPONSE {
			pp.Println("@not_impl_packet_type", packet)
			continue
		}

		r.respChan <- struct {
			id   string
			data []byte
		}{
			id:   packet.Id,
			data: packet.Data,
		}

	}
}

func (r *controlLine) writeLoop() {

	for {

		select {
		case ev := <-r.eventChan:
			r.pendingEvents[ev.id] = ev.resp
			r.conn.Write(ev.packetData)
		case rmsg := <-r.respChan:
			rchan := r.pendingEvents[rmsg.id]
			delete(r.pendingEvents, rmsg.id)
			rchan <- rmsg.data
		}

	}
}

// this func doesnot mutate *controline state
func (r *controlLine) process(ev *event.Request) (*event.Response, error) {

	packet := rtypes.Packet{
		Id:   ev.Id,
		Type: rtypes.EVENT_PROCESS,
		Data: ev.Data,
	}

	pdata, err := json.Marshal(&packet)
	if err != nil {
		return nil, err
	}

	rchan := make(chan []byte)

	fmt.Println("@aaa", r)
	pp.Println(r.eventChan)

	r.eventChan <- struct {
		resp       chan []byte
		id         string
		packetData []byte
	}{
		resp:       rchan,
		id:         ev.Id,
		packetData: pdata,
	}

	rdata := <-rchan

	return &event.Response{
		Payload: rdata,
	}, nil
}
