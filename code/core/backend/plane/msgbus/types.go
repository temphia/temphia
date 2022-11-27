package msgbus

type IMsgBus interface {
	Subscribe(topic string, callback func(msg Message))
	Submit(topic string, msg Message)
	SubscribeTenant(tenant, topic string, callback func(msg Message))
	SubmitTenant(tenant, topic string, msg Message)
}

type Message struct {
	Id     int64
	Topic  string
	Tenant string
	Data   []byte
}

type Subscription struct {
	Id   int32
	Path []string
	Func func(msg Message)
}
