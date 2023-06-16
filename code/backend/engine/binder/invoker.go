package binder

import (
	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type InvokerBindings struct {
	job     *job.Job
	corehub store.CoreHub
	handle  *handle.Handle
}

func NewInvoker(handle *handle.Handle) InvokerBindings {
	return InvokerBindings{
		job:     handle.Job,
		corehub: handle.Deps.Corehub,
		handle:  handle,
	}
}

func (b *InvokerBindings) Name() string { return b.job.Invoker.Type() }

func (b *InvokerBindings) UserContext() *claim.UserContext { return b.job.Invoker.UserContext() }

func (b *InvokerBindings) UserInfo() (*entities.UserInfo, error) {
	uctx := b.job.Invoker.UserContext()
	if uctx == nil {
		return nil, easyerr.Error("empty invoker user")
	}

	ruser, err := b.corehub.GetUserByID(b.handle.Namespace, uctx.UserID)
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

func (b *InvokerBindings) ExecMethod(method string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return b.job.Invoker.ExecuteMethod(method, data)
}

func (b *InvokerBindings) UserMessage(opts *bindx.UserMessage) error {

	uctx := b.job.Invoker.UserContext()
	if uctx == nil {
		return easyerr.Error("empty invoker user")
	}

	_, err := b.corehub.AddUserMessage(&entities.UserMessage{
		Title:        opts.Title,
		Read:         false,
		Type:         "message",
		Contents:     opts.Contents,
		UserId:       uctx.UserID,
		FromUser:     "",
		FromPlug:     b.handle.PlugId,
		FromAgent:    b.handle.AgentId,
		PlugCallback: "",
		Encrypted:    false,
		CreatedAt:    nil,
		TenantId:     b.handle.Namespace,
		WarnLevel:    0,
	})

	return err
}
