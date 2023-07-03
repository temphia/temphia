package re

import (
	"fmt"
	"net"

	"github.com/k0kubun/pp"
	"github.com/tidwall/gjson"
)

func (r *Runner) startServer() error {

	// fixme => dynamic port

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		buffer := make([]byte, 1024)

		bytesRead, err := conn.Read(buffer)
		if err != nil {
			return err
		}

		data := buffer[:bytesRead]

		mtype := gjson.GetBytes(data, "type").String()
		mtoken := gjson.GetBytes(data, "token").String()

		noAuthFunc := func() {
			conn.Write([]byte(`401: bad auth`))
			conn.Close()
		}

		switch mtype {
		case "control_auth":
			if mtoken != r.token {
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
	}

	r.clineLock.Lock()
	r.controlLine = &line
	r.clineLock.Unlock()

	err := line.run()
	if err != nil {
		pp.Println(err)
	}

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
	}

	r.blinesLock.Lock()
	r.blines[id] = line
	r.blinesLock.Unlock()

	err := line.run()
	if err != nil {
		pp.Println(err)
	}

	r.blinesLock.Lock()
	delete(r.blines, id)
	r.blinesLock.Unlock()

}
