package msgbus

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
	"github.com/upper/db/v4"
)

func (m *MsgBus) genericPoll(currMax int64) error {

	for {
		events, err := m.getEvents(currMax)
		if err != nil {
			// log here
			time.Sleep(time.Second * 2)
			continue
		}

		for _, ev := range events {
			subs := m.store.getSubs(ev.Type)

			msg := xplane.Message{
				Id:     ev.Id,
				Topic:  ev.Type,
				Tenant: ev.TenantId,
				Path:   "",
				Data:   ev.Data,
			}

			for _, sub := range subs {
				sub.Chan <- msg
			}
		}

		time.Sleep(time.Duration(rand.Int()%5) * time.Second)
	}
}

func (m *MsgBus) watchPoll(currMax int64) error {

	return nil
}

func (m *MsgBus) getMaxId() (int64, error) {

	kv, err := m.db.GetSystemKV("", key(m.nodeId), "xplane")
	if err != nil {
		if errors.Is(db.ErrNoMoreRows, err) {
			err := m.db.AddSystemKV("", &entities.SystemKV{
				Key:   "xplane_max",
				Type:  "xplane",
				Value: "0",
			})
			if err != nil {
				return 0, err
			}

			return 0, nil
		}

		return 0, err
	}

	return strconv.ParseInt(kv.Value, 10, 64)
}

func (m *MsgBus) getEvents(currMax int64) ([]*entities.SystemEvent, error) {
	return m.db.ListSystemEvent(currMax)
}

func key(nid int64) string {
	return fmt.Sprintf("xplane_max_%d", nid)
}
