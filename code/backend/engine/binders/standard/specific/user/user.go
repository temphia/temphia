package user

import (
	"errors"
	"fmt"

	"github.com/temphia/temphia/code/backend/engine/binders/standard/handle"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/invoker"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	ErrEmptyCurrentUser = errors.New("EMPTY CURRENT USER")
)

type Binding struct {
	chub       store.CoreHub
	handle     *handle.Handle
	iuserCache *invoker.User
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

func (ub *Binding) MessageCurrentUser(opts *bindx.UserMessage) error {
	ub.loadInvokeUser()
	if ub.iuserCache == nil {
		return ErrEmptyCurrentUser
	}

	_, err := ub.chub.AddUserMessage(&entities.UserMessage{
		Title:        opts.Title,
		Read:         false,
		Type:         "message",
		Contents:     opts.Contents,
		UserId:       ub.iuserCache.Id,
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

	ruser, err := ub.chub.GetUserByID(ub.handle.Namespace, ub.iuserCache.Id)
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

	// invoker := ub.handle.Job.Invoker
	// ub.iuserCache = invoker.User()
}
