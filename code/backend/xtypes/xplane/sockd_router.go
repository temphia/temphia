package xplane

type SockdRouter interface {
	Publish(tenantId, room string, tags map[string]string, rawData []byte) error
	Broadcast(tenantId, room string, rawData []byte) error
	SendSession(tenantId, session string, rawData []byte) error
}
