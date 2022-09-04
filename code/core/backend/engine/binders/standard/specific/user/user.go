package user

import (
	"errors"

	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	ErrEmptyCurrentUser = errors.New("EMPTY CURRENT USER")
)

type Binding struct {
	chub       store.CoreHub
	handle     *handle.Handle
	iuserCache *job.InvokeUser
}

func New(handle *handle.Handle) Binding {
	return Binding{
		handle: handle,
		chub:   handle.Deps.Corehub,
	}
}

func (ub *Binding) ListUsers(group string) ([]string, error) {
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

func (ub *Binding) MessageUser(group, user, message string, encrypted bool) error {

	_, err := ub.chub.AddUserMessage(&entities.UserMessage{
		Id:           0,
		Title:        "plug message",
		Read:         false,
		Type:         "message",
		Contents:     message,
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

func (ub *Binding) MessageCurrentUser(title, message string, callback bool) error {
	ub.loadInvokeUser()
	if ub.iuserCache == nil {
		return ErrEmptyCurrentUser
	}

	_, err := ub.chub.AddUserMessage(&entities.UserMessage{
		Title:        title,
		Read:         false,
		Type:         "message",
		Contents:     message,
		UserId:       ub.iuserCache.UserId,
		FromUser:     "",
		FromPlug:     ub.handle.PlugId,
		FromAgent:    ub.handle.AgentId,
		PlugCallback: "",
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     ub.handle.Namespace,
		WarnLevel:    0,
	})

	return err
}

func (ub *Binding) CurrentUser() (*entities.UserInfo, error) {
	ub.loadInvokeUser()
	if ub.iuserCache == nil {
		return nil, ErrEmptyCurrentUser
	}

	ruser, err := ub.chub.GetUserByID(ub.handle.Namespace, ub.iuserCache.UserId)
	if err != nil {
		return nil, err
	}

	return &entities.UserInfo{
		UserId:    ruser.UserId,
		FullName:  ruser.FullName,
		Bio:       ruser.Bio,
		PublicKey: ruser.PublicKey,
		Email:     ruser.Email,
		GroupId:   ruser.GroupID,
	}, nil
}

// private

func (ub *Binding) loadInvokeUser() {
	if ub.iuserCache != nil {
		return
	}

	invoker := ub.handle.Job.Invoker
	invoker.CurrentUser()
	ub.iuserCache = invoker.CurrentUser()
}
