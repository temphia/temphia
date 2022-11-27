package msgbus

import "sync"

var (
	_ IMsgBus = (*MessageBus)(nil)
)

type MessageBus struct {
	globalTopics BusRoom
	tenantTopics map[string]*BusRoom
	ttLock       sync.RWMutex
}

func New() *MessageBus {
	return &MessageBus{
		globalTopics: BusRoom{
			topics: make(map[string][]Subscription),
			mlock:  sync.RWMutex{},
		},
		tenantTopics: make(map[string]*BusRoom),
		ttLock:       sync.RWMutex{},
	}
}

func (mb *MessageBus) Subscribe(topic string, callback func(msg Message)) {
	mb.globalTopics.AddSub(topic, callback)
}

func (mb *MessageBus) Submit(topic string, msg Message) {
	mb.globalTopics.SendSub(topic, msg)
}

func (mb *MessageBus) SubscribeTenant(tenant, topic string, callback func(msg Message)) {

	mb.ttLock.Lock()
	troom := mb.tenantTopics[tenant]
	if troom == nil {
		troom = &BusRoom{
			topics:     make(map[string][]Subscription),
			mlock:      sync.RWMutex{},
			subCounter: 0,
		}
		mb.tenantTopics[tenant] = troom
	}
	mb.ttLock.Unlock()

	troom.AddSub(topic, callback)
}

func (mb *MessageBus) SubmitTenant(tenant, topic string, msg Message) {

	mb.ttLock.RLock()
	troom := mb.tenantTopics[tenant]
	if troom == nil {
		mb.ttLock.RUnlock()
		return
	}
	mb.ttLock.RUnlock()

	troom.SendSub(topic, msg)
}
