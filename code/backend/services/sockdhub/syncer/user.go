package syncer

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
)

type UserSyncer struct {
	sockd sockdx.SockdCore
}

func NewUser(sockd sockdx.SockdCore) *UserSyncer {
	return &UserSyncer{
		sockd: sockd,
	}
}

type notifyMessage struct {
	Type string `json:"type,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (s *UserSyncer) NotifyUser(msg *entities.UserMessage) error {
	nmsg := notifyMessage{
		Type: "new",
		Data: msg,
	}

	return s.notify(msg.TenantId, msg.UserId, nmsg)
}

func (s *UserSyncer) NotifyMessageRead(tenantId, user string, msgIds []int) error {
	return s.notify(tenantId, user, notifyMessage{
		Type: "read",
		Data: msgIds,
	})
}

func (s *UserSyncer) NotifyMessageDelete(tenantId, user string, msgIds []int) error {
	return s.notify(tenantId, user, notifyMessage{
		Type: "delete",
		Data: msgIds,
	})
}

func (s *UserSyncer) notify(tenantId, user string, nmsg notifyMessage) error {
	out, err := json.Marshal(&nmsg)
	if err != nil {
		return err
	}

	return s.sockd.SendTagged(
		tenantId,
		sockdx.ROOM_SYS_USERS,
		[]string{fmt.Sprint("sys.user_", user)},
		[]int64{},
		out,
	)

}
