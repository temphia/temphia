package logx

import "github.com/rs/zerolog"

type Provider interface {
	Proxy
	Sink
}

type Log string

type QueryRequest struct {
	Count   int
	From    string
	To      string
	Cursor  string
	Filters map[string]string
}

type Proxy interface {
	Query(tenantId string, req QueryRequest) ([]Log, error)
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
