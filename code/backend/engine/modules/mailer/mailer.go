package mailer

import (
	"strings"

	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dailer *gomail.Dialer
	modipc *modipc.ModIPC
}

func (m *Mailer) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return m.modipc.Handle(method, args)
}

func (m *Mailer) Close() error { return nil }

type sendOpts struct {
	Address string `json:"addr,omitempty"`
	From    string `json:"from,omitempty"`
	Subject string `json:"subject,omitempty"`
	Name    string `json:"name,omitempty"`
	Body    string `json:"body,omitempty"`
}

func (m *Mailer) Send(opts *sendOpts) error {

	msg := gomail.NewMessage()

	if opts.From == "" && strings.Contains(m.dailer.Username, "@") {
		opts.From = m.dailer.Username
	}

	msg.SetHeader("From", opts.From)
	msg.SetAddressHeader("To", opts.Address, opts.Name)
	msg.SetHeader("Subject", opts.Subject)
	msg.SetBody("text/html", opts.Body)

	return m.dailer.DialAndSend(msg)
}
