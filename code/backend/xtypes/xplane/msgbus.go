package xplane

type Message struct {
	Id     int64
	Topic  string
	Tenant string
	Path   string
	Data   string
}

type Subscription struct {
	Id   int32
	Chan chan<- Message
}

type MsgBus interface {
	Start() error
	Stop() error

	Subscribe(topic string, ch chan Message) (int32, error)
	UnSubscribe(topic string, subid int32) error
	Submit(topic string, msg Message) (int64, error)
}
