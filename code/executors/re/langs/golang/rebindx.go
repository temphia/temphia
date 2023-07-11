package reclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type ReBindx struct {
	port      string
	conn      net.Conn
	writechan chan []byte
	pending   map[string]chan *rtypes.Packet
	pLock     sync.Mutex
	btype     string
}

func NewReBindx(port string, btype string) *ReBindx {

	return &ReBindx{
		port:      port,
		writechan: make(chan []byte),
		pending:   make(map[string]chan *rtypes.Packet),
		pLock:     sync.Mutex{},
		btype:     btype,
	}
}

func (r *ReBindx) init() error {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", r.port))
	if err != nil {
		log.Fatalf("Failed to connect to parent process: %v", err)
	}

	// fixme => send auth

	r.conn = conn

	go r.readLoop()
	go r.writeLoop()

	return nil
}

func (r *ReBindx) readLoop() {

	reader := bufio.NewReader(r.conn)

	for {

		out, err := reader.ReadBytes('\n')
		if err != nil {
			continue
		}

		pkt := &rtypes.Packet{}
		err = json.Unmarshal(out, pkt)
		if err != nil {
			continue
		}

		r.pLock.Lock()

		pj, ok := r.pending[pkt.Id]
		if !ok {
			r.pLock.Unlock()
			continue
		}
		delete(r.pending, pkt.Id)
		r.pLock.Unlock()

		pj <- pkt

	}

}

func (r *ReBindx) writeLoop() {

	for {
		out := <-r.writechan
		r.conn.Write(out)
	}

}

func (r *ReBindx) sendPacket(pkt *rtypes.Packet) (*rtypes.Packet, error) {

	out, err := json.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	rchan := make(chan *rtypes.Packet)

	r.pLock.Lock()
	r.pending[pkt.Id] = rchan
	r.pLock.Unlock()

	r.writechan <- out

	return <-rchan, nil
}

func (r *ReBindx) sendPacket2(pkt *rtypes.Packet) error {

	out, err := json.Marshal(pkt)
	if err != nil {
		return err
	}

	r.writechan <- out

	return nil

}

// bindings

var unitByte = []byte("{}")

func (r *ReBindx) Log(msg string) {

	data, _ := json.Marshal(msg)

	r.sendPacket2(&rtypes.Packet{
		Name: "Log",
		Type: r.btype,
		Data: data,
	})

}

func (r *ReBindx) InLinks() ([]bindx.Link, error) {

	out, err := r.sendPacket(&rtypes.Packet{
		Name: "InLinks",
		Type: r.btype,
		Data: unitByte,
	})
	if err != nil {
		return nil, err
	}

	resp := []bindx.Link{}

	err = json.Unmarshal(out.Data, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReBindx) OutLinks() ([]bindx.Link, error) {
	out, err := r.sendPacket(&rtypes.Packet{
		Name: "OutLinks",
		Type: r.btype,
		Data: unitByte,
	})
	if err != nil {
		return nil, err
	}

	resp := []bindx.Link{}

	err = json.Unmarshal(out.Data, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*

	GetFileWithMeta(file string) (data []byte, version int64, err error)

	ListResources() ([]*Resource, error)
	GetResource(name string) (*Resource, error)

	LinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	LinkExecEmit(name, method string, data xtypes.LazyData) error

	NewModule(name string, data xtypes.LazyData) (int32, error)
	ModuleTicket(name string, opts xtypes.LazyData) (string, error)
	ModuleExec(mid int32, method string, data xtypes.LazyData) (xtypes.LazyData, error)
	ForkExec(method string, data []byte) error

*/
