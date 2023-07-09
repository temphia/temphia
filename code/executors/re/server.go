package re

import (
	"fmt"
	"net"

	"github.com/k0kubun/pp"
	"github.com/tidwall/gjson"
)

func (r *Runner) startServer() error {
	pp.WithLineInfo = true

	// fixme => dynamic port

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	r.listener = listener
	return nil
}

func (r *Runner) acceptLoop() {

	defer r.listener.Close()

	for {

		conn, err := r.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		buffer := make([]byte, 1024)

		bytesRead, err := conn.Read(buffer)
		if err != nil {
			pp.Println("@first_packet_error", err.Error())
			conn.Write([]byte(`500: first packet read error`))
			conn.Close()
			continue
		}

		data := buffer[:bytesRead]

		pp.Println(string(data))

		mtype := gjson.GetBytes(data, "type").String()
		mtoken := gjson.GetBytes(data, "token").String()

		noAuthFunc := func() {
			conn.Write([]byte(`401: bad auth`))
			conn.Close()
		}

		switch mtype {
		case "control_auth":
			if mtoken != r.token && mtoken != "superman" {
				noAuthFunc()
				continue
			}

			go r.handleControl(conn)

		case "bind_auth":
			if mtoken != r.token {
				noAuthFunc()
				continue
			}

			go r.handleBind(conn, "FIXME")

		default:
			noAuthFunc()
			continue
		}

	}

}

func (r *Runner) handleControl(conn net.Conn) {

	line := controlLine{
		runner: r,
		closed: false,
		conn:   conn,
		eventChan: make(chan struct {
			resp       chan []byte
			id         string
			packetData []byte
		}),
		respChan: make(chan struct {
			id   string
			data []byte
		}),
		pendingEvents: make(map[string]chan []byte),
	}

	r.clineLock.Lock()
	r.controlLine = &line
	r.clineLock.Unlock()

	line.run()

	r.clineLock.Lock()
	r.controlLine = nil
	r.clineLock.Unlock()

}

func (r *Runner) handleBind(conn net.Conn, id string) {

	line := &bindingsLine{
		runner: r,
		closed: false,
		conn:   conn,
		id:     id,
		writeChan: make(chan struct {
			isErr bool
			id    string
			data  []byte
		}),
	}

	line.run()

}
