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

type CourierHub interface {
	GetCourier() Courier
}

type Courier interface {
	Start(eventbus any) error
	MailTenant(ctx context.Context, msg MailMessage) error
	Mail(ctx context.Context, sender MailSender, msg MailMessage) error
}
