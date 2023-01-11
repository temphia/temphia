package plane

type SockdRouter struct {
	//	natsConn *nats.Conn
}

func (sr *SockdRouter) Publish(tenantId, room string, tags map[string]string, rawData []byte) error {
	return nil
}

func (sr *SockdRouter) Broadcast(tenantId, room string, rawData []byte) error {
	return nil
}

func (sr *SockdRouter) SendSession(tenantId, session string, rawData []byte) error {
	return nil
}
