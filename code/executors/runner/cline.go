package runner

import (
	"net"

	"github.com/k0kubun/pp"
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
	buffer := make([]byte, 1024)

	for {
		// Read data into the buffer
		bytesRead, err := r.conn.Read(buffer)
		if err != nil {
			return
		}

		data := buffer[:bytesRead]
		pp.Println(string(data))

	}
}

func (r *controlLine) writeLoop() {

}
