package sockdhub

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type notifyMessage struct {
	Type string      `json:"type,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (s *SockdHub) NotifyUser(msg *entities.UserMessage) error {
	nmsg := notifyMessage{
		Type: "new",
		Data: msg,
	}

	return s.notify(msg.TenantId, msg.UserId, nmsg)
}

func (s *SockdHub) NotifyMessageRead(tenantId, user string, msgIds []int) error {
	return s.notify(tenantId, user, notifyMessage{
		Type: "read",
		Data: msgIds,
	})
}

func (s *SockdHub) NotifyMessageDelete(tenantId, user string, msgIds []int) error {
	return s.notify(tenantId, user, notifyMessage{
		Type: "delete",
		Data: msgIds,
	})
}

func (s *SockdHub) notify(tenantId, user string, nmsg notifyMessage) error {
	out, err := json.Marshal(&nmsg)
	if err != nil {
		return err
	}

	return s.sockd.SendTagged(
		tenantId,
		ROOM_SYS_USERS,
		[]string{fmt.Sprint("sys.user_", user)},
		[]int64{},
		out,
	)

}
