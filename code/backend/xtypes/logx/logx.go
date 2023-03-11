package logx

import "github.com/rs/zerolog"

type Provider interface {
	Proxy
	Sink
}

type Log string

type Proxy interface {
	Query(from, to, tenantId string, filters map[string]string) ([]Log, error)
}

type Message struct {
	Index    string
	TenantId string
	Payload  string
}

type Sink interface {
	Send(node_id int64, messages []Message) error
	LastLogId() (int64, error)
}

type Service interface {
	GetEngineLogger() *zerolog.Logger
	GetAppLogger() *zerolog.Logger
	GetSiteLogger(tenantId, domain string) zerolog.Logger
	GetServiceLogger(service string) zerolog.Logger
	GetLogProxy() Proxy
}
