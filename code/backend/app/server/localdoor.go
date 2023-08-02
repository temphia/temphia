package server

import (
	"log"
	"net"
)

func (s *Server) localdoor() error {
	l, err := net.Listen("unix", "/tmp/temphia.sock")
	if err != nil {
		return err
	}

	s.ldListener = l

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			s.ldHandle(c)
		}

	}()

	return nil

}

func (s *Server) ldHandle(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 512)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}

	data := buf[0:nr]
	println("Server got:", string(data))
	_, err = c.Write(data)
	if err != nil {
		log.Fatal("Write: ", err)
	}

}
