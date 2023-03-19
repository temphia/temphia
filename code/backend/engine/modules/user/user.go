package user

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Binding struct {
	chub   store.CoreHub
	handle *handle.Handle
}

func New(handle *handle.Handle) Binding {
	return Binding{
		handle: handle,
		chub:   handle.Deps.Corehub,
	}
}

func (ub *Binding) ListUser(group string) ([]string, error) {
	users, err := ub.chub.ListUsersByGroup(ub.handle.Namespace, group)
	if err != nil {
		return nil, err
	}

	ustrs := make([]string, 0, len(users))

	for _, u := range users {
		ustrs = append(ustrs, u.UserId)
	}

	return ustrs, nil
}

func (ub *Binding) MessageUser(group, user string, opts *bindx.UserMessage) error {

	if opts.UsingCurrentUser {
		return easyerr.NotImpl()
	}

	_, err := ub.chub.AddUserMessage(&entities.UserMessage{
		Id:           0,
		Title:        fmt.Sprintf("Plug message: %s", opts.Title),
		Read:         false,
		Type:         "message",
		Contents:     opts.Contents,
		UserId:       user,
		FromUser:     "",
		FromPlug:     ub.handle.PlugId,
		FromAgent:    ub.handle.AgentId,
		PlugCallback: "",
		WarnLevel:    1,
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     ub.handle.Namespace,
	})

	return err
}

func (ub *Binding) GetUser(group, name string) (*entities.UserInfo, error) {
	user, err := ub.chub.GetUserByID(ub.handle.Namespace, name)
	if err != nil {
		return nil, err
	}

	usr := &entities.UserInfo{
		UserId:    user.UserId,
		FullName:  user.FullName,
		Bio:       user.Bio,
		PublicKey: user.PublicKey,
		Email:     user.Email,
		GroupId:   user.GroupID,
	}

	return usr, nil
}
