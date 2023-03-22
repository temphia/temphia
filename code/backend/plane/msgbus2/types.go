package msgbus

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

var _ MessageBus = (*MsgBus)(nil)

type MessageBus interface {
	Subscribe(topic string, ch chan Message) (int32, error)
	UnSubscribe(topic string, subid int32) error
	Submit(topic string, msg Message) (int64, error)
}
