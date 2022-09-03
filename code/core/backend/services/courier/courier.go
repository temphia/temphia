package courier

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

type Courier struct{}

func New() *Courier {
	return &Courier{}
}

func (c *Courier) MailTenant(ctx context.Context, msg service.MailMessage) error {
	pp.Printf("YOU GOT MAIL [%s] =>", msg)
	return nil
}

func (c *Courier) Mail(ctx context.Context, sender service.MailSender, msg service.MailMessage) error {
	pp.Printf("YOU GOT MAIL [%s] =>", sender.User, msg)
	return nil
}

func (c *Courier) Start(eventbus interface{}) error { return nil }
