package mailer

import (
	"crypto/tls"

	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"gopkg.in/gomail.v2"
)

var _ etypes.ModuleBuilder = (*MailerBuilder)(nil)

func NewBuilder(app any) (etypes.ModuleBuilder, error) {
	return &MailerBuilder{
		app: app.(xtypes.App),
	}, nil
}

type MailerBuilder struct {
	app xtypes.App
}

type ArgOptions struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (l MailerBuilder) Instance(opts etypes.ModuleOptions) (etypes.Module, error) {
	mailer := &Mailer{}
	mailer.modipc = modipc.NewModIPC(mailer)

	aopts := ArgOptions{}
	err := opts.Args.AsObject(aopts)
	if err != nil {
		return nil, err
	}

	// fixme => extract options from opts.Resource

	d := gomail.NewDialer(aopts.Host, aopts.Port, aopts.Username, aopts.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mailer.dailer = d

	return mailer, nil
}
