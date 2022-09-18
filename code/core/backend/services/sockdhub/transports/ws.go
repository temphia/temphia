package transports

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type ConnWS struct {
	conn net.Conn
	id   int64
}

func NewConnWS(c *gin.Context, id int64) (*ConnWS, error) {

	conn, err := upgradeWS(c)
	if err != nil {
		return nil, err
	}

	return &ConnWS{
		conn: conn,
		id:   id,
	}, nil

}

func (c *ConnWS) Id() int64 {
	return c.id
}

func (c *ConnWS) Write(data []byte) error {
	// fixme => for do inside loop
	return wsutil.WriteServerMessage(c.conn, ws.OpText, data)
}

func (c *ConnWS) Close() error {
	return c.conn.Close()
}

func (c *ConnWS) Read() ([]byte, error) {

TOP:
	msg, op, err := wsutil.ReadClientData(c.conn)

	if err != nil {
		return nil, err
	}

	switch op {
	case ws.OpText:
		return msg, nil
	default:
		log.Println("unknown OP", op)
		goto TOP
	}

}

func upgradeWS(c *gin.Context) (net.Conn, error) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	return conn, err
}
