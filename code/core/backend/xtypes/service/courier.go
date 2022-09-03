package service

import "context"

type MailMessage struct {
	TenantId   string
	Host       string
	Recipients string
	Subject    string
	Body       string
}

type MailSender struct {
	User     string
	Password string
}

type Courier interface {
	Start(eventbus interface{}) error
	MailTenant(ctx context.Context, msg MailMessage) error
	Mail(ctx context.Context, sender MailSender, msg MailMessage) error
}
