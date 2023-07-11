package reclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

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

func (r *ReBindx) Log(msg string) {

	data, _ := json.Marshal(msg)

	r.sendPacket2(&rtypes.Packet{
		Id:   "",
		Name: "bindx_log",
		Type: r.btype,
		Data: data,
	})

}
