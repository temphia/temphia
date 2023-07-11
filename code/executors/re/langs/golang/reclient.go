package reclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/temphia/temphia/code/executors/re/rtypes"
)

type Options struct {
	Port     string
	Token    string
	AgentId  string
	PlugId   string
	TenantId string
}

type ReRouter struct {
	opts Options

	mainConn  net.Conn
	bindConns map[string]net.Conn
	handlers  map[string]Handler
}

func New(opts Options) *ReRouter {
	return &ReRouter{
		opts:      opts,
		mainConn:  nil,
		bindConns: make(map[string]net.Conn),
		handlers:  make(map[string]Handler),
	}
}

func (r *ReRouter) Init() error {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", r.opts.Port))
	if err != nil {
		log.Fatalf("Failed to connect to parent process: %v", err)
	}

	r.mainConn = conn

	// fixme => do auth

	go r.readLoop()

	return nil
}

func (r *ReRouter) Register(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *ReRouter) Close() error {
	if r.mainConn != nil {
		return r.mainConn.Close()
	}

	return nil
}

func (r *ReRouter) readLoop() {
	reader := bufio.NewReader(r.mainConn)

	for {
		out, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("Error", err)
			continue
		}

		pkt := &rtypes.Packet{}
		err = json.Unmarshal(out, pkt)
		if err != nil {
			log.Println("Error", err)
			continue
		}

		switch pkt.Type {
		case rtypes.EVENT_PROCESS:
			go func(pkt *rtypes.Packet) {
				handler := r.handlers[pkt.Name]

				handler(Jobctx{
					data: pkt.Data,
				})

				// fixme => send resp to writeloop

			}(pkt)

		default:
			log.Println("not implemented packet:", pkt)
			continue
		}
	}

}
