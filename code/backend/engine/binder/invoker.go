package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (b *Binder) Name() string {
	return b.Job.Invoker.Type()
}

func (b *Binder) UserContext() *claim.UserContext {
	// fixme => static user context ?

	if b.Job != nil {
		return b.Job.Invoker.UserContext()
	}

	return nil
}

func (b *Binder) UserInfo() (*entities.UserInfo, error) {
	uctx := b.Job.Invoker.UserContext()
	if uctx == nil {
		return nil, easyerr.Error("empty invoker user")
	}

	ruser, err := b.Deps.Corehub.GetUserByID(b.Namespace, uctx.UserID)
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

func (b *Binder) ExecMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.Job.Invoker.ExecuteMethod(method, data)
}

func (b *Binder) UserMessage(opts *bindx.UserMessage) error {

	uctx := b.Job.Invoker.UserContext()
	if uctx == nil {
		return easyerr.Error("empty invoker user")
	}

	corehub := b.Deps.Corehub

	_, err := corehub.AddUserMessage(&entities.UserMessage{
		Title:        opts.Title,
		Read:         false,
		Type:         "message",
		Contents:     opts.Contents,
		UserId:       uctx.UserID,
		FromUser:     "",
		FromPlug:     b.PlugId,
		FromAgent:    b.AgentId,
		PlugCallback: "",
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     b.Namespace,
		WarnLevel:    0,
	})

	return err
}
