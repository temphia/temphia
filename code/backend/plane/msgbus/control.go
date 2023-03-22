package msgbus

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
	"github.com/upper/db/v4"
)

func (m *MsgBus) genericPoll(currMax int64) error {

	pp.Println("@generic_poll")

	for {
		events, err := m.getEvents(currMax)
		if err != nil {
			// log here
			time.Sleep(time.Second * 2)
			continue
		}

		nextMaxId := currMax

		for _, ev := range events {
			pp.Println("@processing", ev)

			if ev.Id > nextMaxId {
				nextMaxId = ev.Id
			}

			subs := m.store.getSubs(ev.Type)

			if subs == nil {
				continue
			}

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

		err = m.setMaxId(nextMaxId)
		if err != nil {
			pp.Println("FIXME DONOT PANIC, BUT RETRY INSTEAD")
			panic(err)
		}

		pp.Println("@next_max", nextMaxId)

		currMax = nextMaxId

		time.Sleep(time.Duration(rand.Int()%5) * time.Second)
	}
}

func (m *MsgBus) watchPoll(currMax int64) error {

	return nil
}

func (m *MsgBus) getMaxId() (int64, error) {
	pp.Println("@getmax")

	kv, err := m.db.GetSystemKV("", key(m.nodeId), "xplane")
	if err != nil {
		if errors.Is(db.ErrNoMoreRows, err) {
			pp.Println("@creating_first_id")

			return 0, m.db.AddSystemKV("", &entities.SystemKV{
				Key:      key(m.nodeId),
				Type:     "xplane",
				Value:    "0",
				TenantId: "",
			})
		}

		return 0, err
	}

	return strconv.ParseInt(kv.Value, 10, 64)
}

func (m *MsgBus) setMaxId(maxid int64) error {
	return m.db.UpdateSystemKV("", key(m.nodeId), "xplane", map[string]any{
		"value": fmt.Sprintf("%d", maxid),
	})

}

func (m *MsgBus) getEvents(currMax int64) ([]*entities.SystemEvent, error) {
	return m.db.ListSystemEvent(currMax)
}

func key(nid int64) string {
	return fmt.Sprintf("xplane_max_%d", nid)
}
