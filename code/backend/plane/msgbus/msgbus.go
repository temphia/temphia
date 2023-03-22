package msgbus

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type MsgBus struct {
	store   SubStore
	db      store.SystemOps
	nodeId  int64
	generic bool
}

func New(nodeId int64, db store.SystemOps) *MsgBus {
	return &MsgBus{
		store: SubStore{
			topics:     make(map[string][]xplane.Subscription),
			mlock:      sync.Mutex{},
			subCounter: 0,
		},
		db:      db,
		nodeId:  nodeId,
		generic: true,
	}
}

func (m *MsgBus) Start() error {

	currMax, err := m.getMaxId()
	if err != nil {
		return err
	}

	if m.generic {
		return m.genericPoll(currMax)
	}
	return m.watchPoll(currMax)
}

func (m *MsgBus) Stop() error {
	return nil
}

func (m *MsgBus) Submit(topic string, msg xplane.Message) (int64, error) {
	return 0, m.db.AddSystemEvent(&entities.SystemEvent{
		Type:      topic,
		Data:      msg.Data,
		ExtraMeta: entities.JsonStrMap{},
		TenantId:  msg.Tenant,
	})
}

func (m *MsgBus) Subscribe(topic string, ch chan xplane.Message) (int32, error) {
	return m.store.addSub(topic, ch), nil
}

func (m *MsgBus) UnSubscribe(topic string, subid int32) error {
	m.store.removeSub(topic, subid)
	return nil
}
